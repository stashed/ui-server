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

package shared

import (
	"errors"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	resourcemetrics "kmodules.xyz/resource-metrics"
)

const (
	LastAppliedConfiguration = "kubectl.kubernetes.io/last-applied-configuration"
	Timeout                  = 4 * time.Second
)

func GetDBVersion(obj map[string]interface{}) (string, error) {
	val, found, err := unstructured.NestedFieldCopy(obj, "spec", "version")
	if err != nil {
		return "", err
	}
	if !found {
		return "", errors.New("version field can't be found")
	}
	return val.(string), nil
}

func GetDatabaseType(obj map[string]interface{}) (string, error) {
	mode, err := resourcemetrics.Mode(obj)
	if err != nil {
		return "", err
	}
	return mode, nil
}

func GetDatabaseStatus(obj map[string]interface{}) (string, error) {
	val, found, err := unstructured.NestedFieldCopy(obj, "status", "phase")
	if err != nil {
		return "", err
	}
	if !found {
		return "", errors.New("version field can't be found")
	}
	return val.(string), nil
}
