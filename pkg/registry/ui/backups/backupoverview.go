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
	"time"

	"stash.appscode.dev/apimachinery/apis/ui"
	uiv1alpha1 "stash.appscode.dev/apimachinery/apis/ui/v1alpha1"

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
var _ rest.Storage = &BackupOverviewStorage{}

func NewBackupOverviewStorage(kc client.Client, a authorizer.Authorizer) *BackupOverviewStorage {
	return &BackupOverviewStorage{
		kc: kc,
		a:  a,
		gr: schema.GroupResource{
			Group:    ui.GroupName,
			Resource: uiv1alpha1.ResourceBackupOverviews,
		},
		convertor: rest.NewDefaultTableConvertor(schema.GroupResource{
			Group:    ui.GroupName,
			Resource: uiv1alpha1.ResourceBackupOverviews,
		}),
	}
}

func (r *BackupOverviewStorage) GroupVersionKind(_ schema.GroupVersion) schema.GroupVersionKind {
	return uiv1alpha1.SchemeGroupVersion.WithKind(uiv1alpha1.ResourceKindBackupOverview)
}

func (r *BackupOverviewStorage) NamespaceScoped() bool {
	return true
}

func (r *BackupOverviewStorage) New() runtime.Object {
	return &uiv1alpha1.BackupOverview{}
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

	// Todo: sent real data
	return &uiv1alpha1.BackupOverview{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		Spec: uiv1alpha1.BackupOverviewSpec{
			Schedule:           "Every 30 min",
			LastBackupTime:     &metav1.Time{Time: time.Now()},
			UpcomingBackupTime: &metav1.Time{Time: time.Now()},
			BackupStorage:      "GCS-Bucket",
			DataSize:           "10Gi",
			NumberOfSnapshots:  100,
			DataIntegrity:      true,
			DataDirectory:      "/data/db",
		},
	}, nil
}

func (r *BackupOverviewStorage) NewList() runtime.Object {
	return &uiv1alpha1.BackupOverviewList{}
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
		Verb:      "get",
		Namespace: ns,
		APIGroup:  r.gr.Group,
		Resource:  r.gr.Resource,
		Name:      "",
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

	//todo: Implement list logic

	boList := make([]uiv1alpha1.BackupOverview, 0)
	bo := uiv1alpha1.BackupOverview{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "mg-sh",
			Namespace: ns,
		},
		Spec: uiv1alpha1.BackupOverviewSpec{
			Schedule:           "Every 30 min",
			LastBackupTime:     &metav1.Time{Time: time.Now()},
			UpcomingBackupTime: &metav1.Time{Time: time.Now()},
			BackupStorage:      "GCS-Bucket",
			DataSize:           "10Gi",
			NumberOfSnapshots:  100,
			DataIntegrity:      true,
			DataDirectory:      "/data/db",
		},
	}
	boList = append(boList, bo)

	res := uiv1alpha1.BackupOverviewList{
		TypeMeta: metav1.TypeMeta{},
		ListMeta: metav1.ListMeta{},
		Items:    boList,
	}
	res.ListMeta.SelfLink = ""
	return &res, nil
}

func (r *BackupOverviewStorage) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return r.convertor.ConvertToTable(ctx, object, tableOptions)
}
