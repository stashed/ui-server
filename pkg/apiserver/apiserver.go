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

package apiserver

import (
	"context"
	"fmt"

	stashinstall "stash.appscode.dev/apimachinery/apis/stash/install"
	"stash.appscode.dev/apimachinery/apis/ui"
	uiinstall "stash.appscode.dev/apimachinery/apis/ui/install"
	uiv1alpha1 "stash.appscode.dev/apimachinery/apis/ui/v1alpha1"
	"stash.appscode.dev/ui-server/pkg/registry/ui/backups"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
	"k8s.io/klog/v2/klogr"
	"kmodules.xyz/authorizer/rbac"
	cu "kmodules.xyz/client-go/client"
	appcataloginstall "kmodules.xyz/custom-resources/apis/appcatalog/install"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

var (
	// Scheme defines methods for serializing and deserializing API objects.
	Scheme = runtime.NewScheme()
	// Codecs provides methods for retrieving codecs and serializers for specific
	// versions and content types.
	Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
	uiinstall.Install(Scheme)
	stashinstall.Install(Scheme)
	appcataloginstall.Install(Scheme)

	_ = clientgoscheme.AddToScheme(Scheme)

	// we need to add the options to empty v1
	// TODO fix the server code to avoid this
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})

	// TODO: keep the generic API server from wanting this
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

// ExtraConfig holds custom apiserver config
type ExtraConfig struct {
	ClientConfig *restclient.Config
}

// Config defines the config for the apiserver
type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}

// UIServer contains state for a Kubernetes cluster master/api server.
type UIServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
	Manager          manager.Manager
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}

// CompletedConfig embeds a private pointer that cannot be instantiated outside of this package.
type CompletedConfig struct {
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{
		cfg.GenericConfig.Complete(),
		&cfg.ExtraConfig,
	}

	c.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	return CompletedConfig{&c}
}

// New returns a new instance of UIServer from the given config.
func (c completedConfig) New(ctx context.Context) (*UIServer, error) {
	genericServer, err := c.GenericConfig.New("stash-ui-server", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	// ctrl.SetLogger(...)
	log.SetLogger(klogr.New()) // nolint:staticcheck

	mgr, err := manager.New(c.ExtraConfig.ClientConfig, manager.Options{
		Scheme:                 Scheme,
		Metrics:                metricsserver.Options{BindAddress: ""},
		HealthProbeBindAddress: "",
		LeaderElection:         false,
		LeaderElectionID:       "5b87adeb.ui.stash.appscode.com",
		NewClient:              cu.NewClient,
		Client: client.Options{
			Cache: &client.CacheOptions{
				DisableFor: []client.Object{
					&core.Pod{},
				},
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("unable to start manager, reason: %v", err)
	}
	ctrlClient := mgr.GetClient()

	rbacAuthorizer := rbac.NewForManagerOrDie(ctx, mgr)

	s := &UIServer{
		GenericAPIServer: genericServer,
		Manager:          mgr,
	}

	{
		apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(ui.GroupName, Scheme, metav1.ParameterCodec, Codecs)

		v1alpha1storage := map[string]rest.Storage{}

		v1alpha1storage[uiv1alpha1.ResourceBackupOverviews] = backups.NewBackupOverviewStorage(ctrlClient, rbacAuthorizer)

		apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage

		if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
			return nil, err
		}
	}

	return s, nil
}
