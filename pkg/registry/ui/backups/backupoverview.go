/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package backups

import (
	"context"
	"errors"
	"fmt"
	"time"

	stashapi "stash.appscode.dev/apimachinery/apis/stash"
	stashv1alpha1 "stash.appscode.dev/apimachinery/apis/stash/v1alpha1"
	stashv1beta1 "stash.appscode.dev/apimachinery/apis/stash/v1beta1"
	"stash.appscode.dev/apimachinery/apis/ui"
	uiapi "stash.appscode.dev/apimachinery/apis/ui/v1alpha1"

	"github.com/lnquy/cron"
	rcron "github.com/robfig/cron/v3"
	"gomodules.xyz/pointer"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type BackupOverviewStorage struct {
	kc        client.Client
	a         authorizer.Authorizer
	gr        schema.GroupResource
	convertor rest.TableConvertor
}

var _ rest.GroupVersionKindProvider = &BackupOverviewStorage{}
var _ rest.Scoper = &BackupOverviewStorage{}
var _ rest.Getter = &BackupOverviewStorage{}
var _ rest.Lister = &BackupOverviewStorage{}

func NewBackupOverviewStorage(kc client.Client, a authorizer.Authorizer) *BackupOverviewStorage {
	return &BackupOverviewStorage{
		kc: kc,
		a:  a,
		gr: schema.GroupResource{
			Group:    stashapi.GroupName,
			Resource: stashv1beta1.ResourcePluralBackupConfiguration,
		},
		convertor: rest.NewDefaultTableConvertor(schema.GroupResource{
			Group:    ui.GroupName,
			Resource: uiapi.ResourceBackupOverviews,
		}),
	}
}

func (r *BackupOverviewStorage) GroupVersionKind(_ schema.GroupVersion) schema.GroupVersionKind {
	return uiapi.SchemeGroupVersion.WithKind(uiapi.ResourceKindBackupOverview)
}

func (r *BackupOverviewStorage) NamespaceScoped() bool {
	return true
}

func (r *BackupOverviewStorage) New() runtime.Object {
	return &uiapi.BackupOverview{}
}

func (r *BackupOverviewStorage) NewList() runtime.Object {
	return &uiapi.BackupOverviewList{}
}

func (r *BackupOverviewStorage) Get(ctx context.Context, name string, _ *metav1.GetOptions) (runtime.Object, error) {
	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		return nil, apierrors.NewBadRequest("missing namespace")
	}

	user, ok := apirequest.UserFrom(ctx)
	if !ok {
		return nil, apierrors.NewBadRequest("missing user info")
	}

	attrs := authorizer.AttributesRecord{
		User:      user,
		Verb:      "get",
		Namespace: ns,
		APIGroup:  r.gr.Group,
		Resource:  r.gr.Resource,
		Name:      name,
	}
	decision, why, err := r.a.Authorize(ctx, attrs)
	if err != nil {
		return nil, apierrors.NewInternalError(err)
	}
	if decision != authorizer.DecisionAllow {
		return nil, apierrors.NewForbidden(r.gr, name, errors.New(why))
	}
	backupConfig := &stashv1beta1.BackupConfiguration{}
	if err := r.kc.Get(ctx, client.ObjectKey{Name: name, Namespace: ns}, backupConfig); err != nil {
		return nil, fmt.Errorf("failed to get BackupConfiguration, reason: %v", err)
	}
	return r.getBackupOverview(ctx, backupConfig.DeepCopy())
}

func (r *BackupOverviewStorage) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		return nil, apierrors.NewBadRequest("missing namespace")
	}

	user, ok := apirequest.UserFrom(ctx)
	if !ok {
		return nil, apierrors.NewBadRequest("missing user info")
	}

	attrs := authorizer.AttributesRecord{
		User:      user,
		Verb:      "list",
		Namespace: ns,
		APIGroup:  r.gr.Group,
		Resource:  r.gr.Resource,
	}
	decision, why, err := r.a.Authorize(ctx, attrs)
	if err != nil {
		return nil, apierrors.NewInternalError(err)
	}
	if decision != authorizer.DecisionAllow {
		return nil, apierrors.NewForbidden(r.gr, "", errors.New(why))
	}

	opts := client.ListOptions{Namespace: ns}
	if options != nil {
		if options.LabelSelector != nil && !options.LabelSelector.Empty() {
			opts.LabelSelector = options.LabelSelector
		}
		if options.FieldSelector != nil && !options.FieldSelector.Empty() {
			opts.FieldSelector = options.FieldSelector
		}
		opts.Limit = options.Limit
		opts.Continue = options.Continue
	}

	backupCfgList := stashv1beta1.BackupConfigurationList{}
	if err := r.kc.List(ctx, &backupCfgList, &opts); err != nil {
		return nil, err
	}

	backupOverviews := make([]uiapi.BackupOverview, 0, len(backupCfgList.Items))
	for _, c := range backupCfgList.Items {
		bo, err := r.getBackupOverview(ctx, c.DeepCopy())
		if err != nil {
			return nil, err
		}
		backupOverviews = append(backupOverviews, *bo)
	}
	result := &uiapi.BackupOverviewList{
		TypeMeta: metav1.TypeMeta{},
		ListMeta: backupCfgList.ListMeta,
		Items:    backupOverviews,
	}
	return result, nil
}

func (r *BackupOverviewStorage) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return r.convertor.ConvertToTable(ctx, object, tableOptions)
}

func (r *BackupOverviewStorage) getBackupOverview(ctx context.Context, cfg *stashv1beta1.BackupConfiguration) (*uiapi.BackupOverview, error) {
	repo, err := getRepository(ctx, r.kc, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to get Repository, reason: %v", err)
	}

	exprDesc, _ := cron.NewDescriptor()
	desc, err := exprDesc.ToDescription(cfg.Spec.Schedule, cron.Locale_en)
	if err != nil {
		return nil, err
	}

	sched, err := rcron.NewParser(rcron.Minute | rcron.Hour | rcron.Dom | rcron.Month | rcron.Dow).Parse(cfg.Spec.Schedule)
	if err != nil {
		return nil, err
	}

	backupOverview := &uiapi.BackupOverview{
		ObjectMeta: *cfg.ObjectMeta.DeepCopy(),
		Spec: uiapi.BackupOverviewSpec{
			Schedule:           fmt.Sprintf("%s(%s)", cfg.Spec.Schedule, desc),
			LastBackupTime:     repo.Status.LastBackupTime,
			UpcomingBackupTime: &metav1.Time{Time: sched.Next(time.Now())},
			Repository:         repo.Name,
			DataSize:           repo.Status.TotalSize,
			NumberOfSnapshots:  repo.Status.SnapshotCount,
			DataIntegrity:      pointer.Bool(repo.Status.Integrity),
		},
	}
	if cfg.Spec.Paused {
		backupOverview.Spec.Status = uiapi.BackupStatusPaused
	} else {
		backupOverview.Spec.Status = uiapi.BackupStatusActive
	}
	backupOverview.SelfLink = ""
	backupOverview.ManagedFields = nil

	return backupOverview, nil
}

// Helper function to get the Repository for a Stash BackupConfiguration object
func getRepository(ctx context.Context, kc client.Client, backupConfig *stashv1beta1.BackupConfiguration) (*stashv1alpha1.Repository, error) {
	repoKey := client.ObjectKey{Name: backupConfig.Spec.Repository.Name, Namespace: backupConfig.Spec.Repository.Namespace}
	if repoKey.Namespace == "" {
		repoKey.Namespace = backupConfig.Namespace
	}

	repo := &stashv1alpha1.Repository{}
	if err := kc.Get(ctx, repoKey, repo); err != nil {
		return nil, err
	}
	return repo, nil
}
