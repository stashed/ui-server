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
	"fmt"
	"time"

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
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type BackupOverviewStorage struct {
	kc client.Client
	a  authorizer.Authorizer
	gr schema.GroupResource
}

var _ rest.GroupVersionKindProvider = &BackupOverviewStorage{}
var _ rest.Scoper = &BackupOverviewStorage{}
var _ rest.Creater = &BackupOverviewStorage{}
var _ rest.Storage = &BackupOverviewStorage{}

func NewBackupOverviewStorage(kc client.Client, a authorizer.Authorizer) *BackupOverviewStorage {
	return &BackupOverviewStorage{
		kc: kc,
		a:  a,
		gr: schema.GroupResource{
			Group:    ui.GroupName,
			Resource: uiapi.ResourceBackupOverviews,
		},
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

func (r *BackupOverviewStorage) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, _ *metav1.CreateOptions) (runtime.Object, error) {
	in := obj.(*uiapi.BackupOverview)
	if in.Request == nil {
		return nil, apierrors.NewBadRequest("missing apirequest")
	}
	req := in.Request

	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		return nil, apierrors.NewBadRequest("missing namespace")
	}

	//user, ok := apirequest.UserFrom(ctx)
	//if !ok {
	//	return nil, apierrors.NewBadRequest("missing user info")
	//}

	//attrs := authorizer.AttributesRecord{
	//	User:      user,
	//	Verb:      "create",
	//	Namespace: ns,
	//	APIGroup:  r.gr.Group,
	//	Resource:  r.gr.Resource,
	//	Name:      in.Request.Ref.Name,
	//}
	//decision, why, err := r.a.Authorize(ctx, attrs)
	//if err != nil {
	//	return nil, apierrors.NewInternalError(err)
	//}
	//if decision != authorizer.DecisionAllow {
	//	return nil, apierrors.NewForbidden(r.gr, in.Request.Ref.Name, errors.New(why))
	//}

	rid, err := kmapi.ExtractResourceID(r.kc.RESTMapper(), req.Resource)
	if err != nil {
		return nil, apierrors.NewInternalError(err)
	}
	gvr := rid.GroupVersionResource()

	ab, err := getAppBinding(ctx, r.kc, gvr, client.ObjectKey{Name: req.Ref.Name, Namespace: ns})
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
		Repository:         repo.Name,
		DataSize:           repo.Status.TotalSize,
		NumberOfSnapshots:  repo.Status.SnapshotCount,
		DataIntegrity:      pointer.Bool(repo.Status.Integrity),
	}
	if backupConfig.Spec.Paused {
		in.Response.Status = uiapi.BackupStatusPaused
	} else {
		in.Response.Status = uiapi.BackupStatusActive
	}
	return in, nil
}
