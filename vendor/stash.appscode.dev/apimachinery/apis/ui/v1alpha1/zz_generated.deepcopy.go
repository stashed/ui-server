//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupOverview) DeepCopyInto(out *BackupOverview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupOverview.
func (in *BackupOverview) DeepCopy() *BackupOverview {
	if in == nil {
		return nil
	}
	out := new(BackupOverview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupOverview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupOverviewList) DeepCopyInto(out *BackupOverviewList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupOverview, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupOverviewList.
func (in *BackupOverviewList) DeepCopy() *BackupOverviewList {
	if in == nil {
		return nil
	}
	out := new(BackupOverviewList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupOverviewList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupOverviewSpec) DeepCopyInto(out *BackupOverviewSpec) {
	*out = *in
	if in.LastBackupTime != nil {
		in, out := &in.LastBackupTime, &out.LastBackupTime
		*out = (*in).DeepCopy()
	}
	if in.UpcomingBackupTime != nil {
		in, out := &in.UpcomingBackupTime, &out.UpcomingBackupTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupOverviewSpec.
func (in *BackupOverviewSpec) DeepCopy() *BackupOverviewSpec {
	if in == nil {
		return nil
	}
	out := new(BackupOverviewSpec)
	in.DeepCopyInto(out)
	return out
}
