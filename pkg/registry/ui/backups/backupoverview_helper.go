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

	stashv1alpha1 "stash.appscode.dev/apimachinery/apis/stash/v1alpha1"
	stashv1beta1 "stash.appscode.dev/apimachinery/apis/stash/v1beta1"

	"k8s.io/apimachinery/pkg/runtime/schema"
	appcatalog "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	kubedbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Helper function to get the associate AppBinding for a specific KubeDB database
func getAppBinding(ctx context.Context, kc client.Client, gvr schema.GroupVersionResource, key client.ObjectKey) (*appcatalog.AppBinding, error) {
	labels, err := getOffshootLabels(ctx, kc, gvr, key)
	if err != nil {
		return nil, err
	}

	abList := &appcatalog.AppBindingList{}
	opts := &client.ListOptions{Namespace: key.Namespace}
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

// Helper function to get the BackupConfiguration for a Database AppBinding
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

// Helper function to get the KubeDB Database OffShoot labels
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
