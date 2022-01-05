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
	apierrors "k8s.io/apimachinery/pkg/api/errors"
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

func (r *BackupOverviewStorage) Get(ctx context.Context, name string, opts *metav1.GetOptions) (runtime.Object, error) {
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

	cfgList := &stashv1beta1.BackupConfigurationList{}
	if err := r.kc.List(ctx, cfgList); err != nil {
		return nil, err
	}

	backupConfig := stashv1beta1.BackupConfiguration{}
	for _, c := range cfgList.Items {
		if c.Spec.Target != nil && c.Spec.Target.Ref.Name == name {
			backupConfig = c
			break
		}
	}

	if backupConfig.Spec.Target == nil {
		return nil, apierrors.NewInternalError(fmt.Errorf("no BackupConfiguration is found for the Database %v/%v", ns, name))
	}
	fmt.Println(backupConfig)

	repo := &stashv1alpha1.Repository{}
	repoKey := client.ObjectKey{Name: backupConfig.Spec.Repository.Name, Namespace: backupConfig.Spec.Repository.Namespace}
	if repoKey.Namespace == "" {
		repoKey.Namespace = backupConfig.Namespace
	}
	if err := r.kc.Get(ctx, repoKey, repo); err != nil {
		return nil, err
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

	return &uiapi.BackupOverview{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		Spec: uiapi.BackupOverviewSpec{
			Schedule:           fmt.Sprintf("%s(%s)", backupConfig.Spec.Schedule, desc),
			LastBackupTime:     repo.Status.LastBackupTime,
			UpcomingBackupTime: &metav1.Time{Time: sched.Next(time.Now())},
			BackupStorage:      repo.Name,
			DataSize:           repo.Status.TotalSize,
			NumberOfSnapshots:  repo.Status.SnapshotCount,
			DataIntegrity:      pointer.Bool(repo.Status.Integrity),
			DataDirectory:      "unknown",
		},
	}, nil
}

//func (r *BackupOverviewStorage) NewList() runtime.Object {
//	return &uiapi.BackupOverviewList{}
//}

//func (r *BackupOverviewStorage) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
//	ns, ok := apirequest.NamespaceFrom(ctx)
//	if !ok {
//		return nil, apierrors.NewBadRequest("missing namespace")
//	}
//
//	user, ok := apirequest.UserFrom(ctx)
//	if !ok {
//		return nil, apierrors.NewBadRequest("missing user info")
//	}
//
//	attrs := authorizer.AttributesRecord{
//		User:      user,
//		Verb:      "get",
//		Namespace: ns,
//		APIGroup:  r.gr.Group,
//		Resource:  r.gr.Resource,
//		Name:      "",
//	}
//
//	decision, why, err := r.a.Authorize(ctx, attrs)
//	if err != nil {
//		return nil, apierrors.NewInternalError(err)
//	}
//	if decision != authorizer.DecisionAllow {
//		return nil, apierrors.NewForbidden(r.gr, "", errors.New(why))
//	}
//
//	opts := client.ListOptions{Namespace: ns}
//	if options != nil {
//		if options.LabelSelector != nil && !options.LabelSelector.Empty() {
//			opts.LabelSelector = options.LabelSelector
//		}
//		if options.FieldSelector != nil && !options.FieldSelector.Empty() {
//			opts.FieldSelector = options.FieldSelector
//		}
//		opts.Limit = options.Limit
//		opts.Continue = options.Continue
//	}
//
//	//todo: Implement list logic
//
//	boList := make([]uiapi.BackupOverview, 0)
//	bo := uiapi.BackupOverview{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      "mg-sh",
//			Namespace: ns,
//		},
//		Spec: uiapi.BackupOverviewSpec{
//			Schedule:           "Every 30 min",
//			LastBackupTime:     &metav1.Time{Time: time.Now()},
//			UpcomingBackupTime: &metav1.Time{Time: time.Now()},
//			BackupStorage:      "GCS-Bucket",
//			DataSize:           "10Gi",
//			NumberOfSnapshots:  100,
//			DataIntegrity:      true,
//			DataDirectory:      "/data/db",
//		},
//	}
//	boList = append(boList, bo)
//
//	res := uiapi.BackupOverviewList{
//		TypeMeta: metav1.TypeMeta{},
//		ListMeta: metav1.ListMeta{},
//		Items:    boList,
//	}
//	res.ListMeta.SelfLink = ""
//	return &res, nil
//}

func (r *BackupOverviewStorage) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return r.convertor.ConvertToTable(ctx, object, tableOptions)
}
