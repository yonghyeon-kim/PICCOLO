//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevRes) DeepCopyInto(out *DevRes) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevRes.
func (in *DevRes) DeepCopy() *DevRes {
	if in == nil {
		return nil
	}
	out := new(DevRes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevRes) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevResInfo) DeepCopyInto(out *DevResInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevResInfo.
func (in *DevResInfo) DeepCopy() *DevResInfo {
	if in == nil {
		return nil
	}
	out := new(DevResInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevResList) DeepCopyInto(out *DevResList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevRes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevResList.
func (in *DevResList) DeepCopy() *DevResList {
	if in == nil {
		return nil
	}
	out := new(DevResList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevResList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevResSpec) DeepCopyInto(out *DevResSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevResSpec.
func (in *DevResSpec) DeepCopy() *DevResSpec {
	if in == nil {
		return nil
	}
	out := new(DevResSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevResStatus) DeepCopyInto(out *DevResStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevResStatus.
func (in *DevResStatus) DeepCopy() *DevResStatus {
	if in == nil {
		return nil
	}
	out := new(DevResStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PiccoloVolume) DeepCopyInto(out *PiccoloVolume) {
	*out = *in
	in.Volume.DeepCopyInto(&out.Volume)
	in.VolumeMount.DeepCopyInto(&out.VolumeMount)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PiccoloVolume.
func (in *PiccoloVolume) DeepCopy() *PiccoloVolume {
	if in == nil {
		return nil
	}
	out := new(PiccoloVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VehContainerConfig) DeepCopyInto(out *VehContainerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VehContainerConfig.
func (in *VehContainerConfig) DeepCopy() *VehContainerConfig {
	if in == nil {
		return nil
	}
	out := new(VehContainerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VehContainerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VehContainerConfigList) DeepCopyInto(out *VehContainerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VehContainerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VehContainerConfigList.
func (in *VehContainerConfigList) DeepCopy() *VehContainerConfigList {
	if in == nil {
		return nil
	}
	out := new(VehContainerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VehContainerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VehContainerConfigSpec) DeepCopyInto(out *VehContainerConfigSpec) {
	*out = *in
	if in.PiccoloVolumes != nil {
		in, out := &in.PiccoloVolumes, &out.PiccoloVolumes
		*out = make([]PiccoloVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PiccoloEnv != nil {
		in, out := &in.PiccoloEnv, &out.PiccoloEnv
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VehContainerConfigSpec.
func (in *VehContainerConfigSpec) DeepCopy() *VehContainerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(VehContainerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VehContainerConfigStatus) DeepCopyInto(out *VehContainerConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VehContainerConfigStatus.
func (in *VehContainerConfigStatus) DeepCopy() *VehContainerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(VehContainerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VehPod) DeepCopyInto(out *VehPod) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VehPod.
func (in *VehPod) DeepCopy() *VehPod {
	if in == nil {
		return nil
	}
	out := new(VehPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VehPod) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VehPodList) DeepCopyInto(out *VehPodList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VehPod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VehPodList.
func (in *VehPodList) DeepCopy() *VehPodList {
	if in == nil {
		return nil
	}
	out := new(VehPodList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VehPodList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VehPodSpec) DeepCopyInto(out *VehPodSpec) {
	*out = *in
	if in.DevResInfo != nil {
		in, out := &in.DevResInfo, &out.DevResInfo
		*out = make([]DevResInfo, len(*in))
		copy(*out, *in)
	}
	if in.VehContainerConfigInfo != nil {
		in, out := &in.VehContainerConfigInfo, &out.VehContainerConfigInfo
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.PodSpec.DeepCopyInto(&out.PodSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VehPodSpec.
func (in *VehPodSpec) DeepCopy() *VehPodSpec {
	if in == nil {
		return nil
	}
	out := new(VehPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VehPodStatus) DeepCopyInto(out *VehPodStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VehPodStatus.
func (in *VehPodStatus) DeepCopy() *VehPodStatus {
	if in == nil {
		return nil
	}
	out := new(VehPodStatus)
	in.DeepCopyInto(out)
	return out
}
