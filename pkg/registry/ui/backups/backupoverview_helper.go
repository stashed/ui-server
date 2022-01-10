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

	stashv1alpha1 "stash.appscode.dev/apimachinery/apis/stash/v1alpha1"
	stashv1beta1 "stash.appscode.dev/apimachinery/apis/stash/v1beta1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

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
