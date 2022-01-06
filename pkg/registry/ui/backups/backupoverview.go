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

	stashv1alpha1 "stash.appscode.dev/apimachinery/apis/stash/v1alpha1"
	stashv1beta1 "stash.appscode.dev/apimachinery/apis/stash/v1beta1"
	"stash.appscode.dev/apimachinery/apis/ui"
	uiapi "stash.appscode.dev/apimachinery/apis/ui/v1alpha1"

	"github.com/lnquy/cron"
	rcron "github.com/robfig/cron/v3"
	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	appcatalog "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	kubedbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
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

//var _ rest.Getter = &BackupOverviewStorage{}
var _ rest.Creater = &BackupOverviewStorage{}

//var _ rest.Lister = &BackupOverviewStorage{}
var _ rest.Storage = &BackupOverviewStorage{}

func NewBackupOverviewStorage(kc client.Client, a authorizer.Authorizer) *BackupOverviewStorage {
	return &BackupOverviewStorage{
		kc: kc,
		a:  a,
		gr: schema.GroupResource{
			Group:    ui.GroupName,
			Resource: uiapi.ResourceBackupOverviews,
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

func (r *BackupOverviewStorage) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	in := obj.(*uiapi.BackupOverview)

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
		Verb:      "create",
		Namespace: ns,
		APIGroup:  r.gr.Group,
		Resource:  r.gr.Resource,
		Name:      in.Name,
	}
	decision, why, err := r.a.Authorize(ctx, attrs)
	if err != nil {
		return nil, apierrors.NewInternalError(err)
	}
	if decision != authorizer.DecisionAllow {
		return nil, apierrors.NewForbidden(r.gr, in.Name, errors.New(why))
	}
	gvr := schema.GroupVersionResource{
		Group:    in.Request.Group,
		Version:  in.Request.Version,
		Resource: in.Request.Resource,
	}
	ab, err := getAppBinding(ctx, r.kc, gvr, client.ObjectKey{Name: in.Name, Namespace: in.Namespace})
	if err != nil {
		return nil, fmt.Errorf("failed to get AppBinding, reason: %v", err)
	}
	backupConfig, err := getBackupConfig(ctx, r.kc, ab)
	if err != nil {
		return nil, fmt.Errorf("failed to get BackupConfiguration, reason: %v", err)
	}
	repo, err := getRepository(ctx, r.kc, backupConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get Repository, reason: %v", err)
	}

	exprDesc, _ := cron.NewDescriptor()
	desc, err := exprDesc.ToDescription(backupConfig.Spec.Schedule, cron.Locale_en)
	if err != nil {
		return nil, err
	}

	sched, err := rcron.NewParser(rcron.Minute | rcron.Hour | rcron.Dom | rcron.Month | rcron.Dow).Parse(backupConfig.Spec.Schedule)
	if err != nil {
		return nil, err
	}

	in.Response = uiapi.BackupOverviewResponse{
		Schedule:           fmt.Sprintf("%s(%s)", backupConfig.Spec.Schedule, desc),
		LastBackupTime:     repo.Status.LastBackupTime,
		UpcomingBackupTime: &metav1.Time{Time: sched.Next(time.Now())},
		BackupStorage:      repo.Name,
		DataSize:           repo.Status.TotalSize,
		NumberOfSnapshots:  repo.Status.SnapshotCount,
		DataIntegrity:      pointer.Bool(repo.Status.Integrity),
		DataDirectory:      "unknown",
	}
	return in, nil
}

func getAppBinding(ctx context.Context, kc client.Client, gvr schema.GroupVersionResource, key client.ObjectKey) (*appcatalog.AppBinding, error) {
	labels, err := getOffshootLabels(ctx, kc, gvr, key)
	if err != nil {
		return nil, err
	}

	abList := &appcatalog.AppBindingList{}
	opts := &client.ListOptions{Namespace: core.NamespaceAll}
	selector := client.MatchingLabels(labels)
	selector.ApplyToList(opts)
	if err := kc.List(ctx, abList, opts); err != nil {
		return nil, err
	}

	if len(abList.Items) != 1 {
		return nil, fmt.Errorf("expect one AppBinding but got %v", len(abList.Items))
	}
	return &abList.Items[0], nil
}

func getBackupConfig(ctx context.Context, kc client.Client, ab *appcatalog.AppBinding) (*stashv1beta1.BackupConfiguration, error) {
	cfgList := &stashv1beta1.BackupConfigurationList{}
	if err := kc.List(ctx, cfgList, client.InNamespace(ab.Namespace)); err != nil {
		return nil, err
	}
	for _, cfg := range cfgList.Items {
		if cfg.Spec.Target != nil && cfg.Spec.Target.Ref.Name == ab.Name {
			return &cfg, nil
		}
	}
	return nil, errors.New("no BackupConfiguration is found for the given Database")
}

func getOffshootLabels(ctx context.Context, kc client.Client, gvr schema.GroupVersionResource, key client.ObjectKey) (map[string]string, error) {
	switch gvr.Resource {
	case kubedbapi.ResourcePluralMongoDB:
		db := &kubedbapi.MongoDB{}
		if err := kc.Get(ctx, key, db); err != nil {
			return nil, err
		}
		return db.OffshootSelectors(), nil
	case kubedbapi.ResourcePluralElasticsearch:
		db := &kubedbapi.Elasticsearch{}
		if err := kc.Get(ctx, key, db); err != nil {
			return nil, err
		}
		return db.OffshootSelectors(), nil
	case kubedbapi.ResourcePluralPostgres:
		db := &kubedbapi.Postgres{}
		if err := kc.Get(ctx, key, db); err != nil {
			return nil, err
		}
		return db.OffshootSelectors(), nil
	case kubedbapi.ResourcePluralMySQL:
		db := &kubedbapi.MySQL{}
		if err := kc.Get(ctx, key, db); err != nil {
			return nil, err
		}
		return db.OffshootSelectors(), nil
	case kubedbapi.ResourcePluralMariaDB:
		db := &kubedbapi.MariaDB{}
		if err := kc.Get(ctx, key, db); err != nil {
			return nil, err
		}
		return db.OffshootSelectors(), nil
	case kubedbapi.ResourcePluralRedis:
		db := &kubedbapi.Redis{}
		if err := kc.Get(ctx, key, db); err != nil {
			return nil, err
		}
		return db.OffshootSelectors(), nil
	default:
		return nil, fmt.Errorf("database type with GVR %v is not supported", gvr.String())
	}
}

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

func (r *BackupOverviewStorage) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return r.convertor.ConvertToTable(ctx, object, tableOptions)
}
