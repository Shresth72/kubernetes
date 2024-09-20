//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
func (in *KlusterStatus) DeepCopyInto(out *KlusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterStatus.
func (in *KlusterStatus) DeepCopy() *KlusterStatus {
	if in == nil {
		return nil
	}
	out := new(KlusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kluster) DeepCopyInto(out *Kluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta

	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)

	out.Status = in.Status

	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kluster.
func (in *Kluster) DeepCopy() *Kluster {
	if in == nil {
		return nil
	}
	out := new(Kluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterList) DeepCopyInto(out *KlusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterList.
func (in *KlusterList) DeepCopy() *KlusterList {
	if in == nil {
		return nil
	}
	out := new(KlusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KlusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterSpec) DeepCopyInto(out *KlusterSpec) {
	*out = *in
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]NodePool, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterSpec.
func (in *KlusterSpec) DeepCopy() *KlusterSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePool) DeepCopyInto(out *NodePool) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePool.
func (in *NodePool) DeepCopy() *NodePool {
	if in == nil {
		return nil
	}
	out := new(NodePool)
	in.DeepCopyInto(out)
	return out
}
