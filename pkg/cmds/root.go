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

package cmds

import (
	"os"

	"github.com/spf13/cobra"
	v "gomodules.xyz/x/version"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:               "ui-operator [command]",
		Short:             `stash-ui-server by AppsCode`,
		DisableAutoGenTag: true,
	}

	rootCmd.AddCommand(v.NewCmdVersion())
	ctx := genericapiserver.SetupSignalContext()
	rootCmd.AddCommand(NewCmdRun(ctx, os.Stdout, os.Stderr))

	return rootCmd
}
