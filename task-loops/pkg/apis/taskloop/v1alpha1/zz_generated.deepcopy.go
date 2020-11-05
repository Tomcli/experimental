// +build !ignore_autogenerated

/*
Copyright 2020 The Tekton Authors

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
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskLoop) DeepCopyInto(out *TaskLoop) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskLoop.
func (in *TaskLoop) DeepCopy() *TaskLoop {
	if in == nil {
		return nil
	}
	out := new(TaskLoop)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskLoop) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskLoopList) DeepCopyInto(out *TaskLoopList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TaskLoop, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskLoopList.
func (in *TaskLoopList) DeepCopy() *TaskLoopList {
	if in == nil {
		return nil
	}
	out := new(TaskLoopList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskLoopList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskLoopRunStatus) DeepCopyInto(out *TaskLoopRunStatus) {
	*out = *in
	if in.TaskLoopSpec != nil {
		in, out := &in.TaskLoopSpec, &out.TaskLoopSpec
		*out = new(TaskLoopSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TaskRuns != nil {
		in, out := &in.TaskRuns, &out.TaskRuns
		*out = make(map[string]*TaskLoopTaskRunStatus, len(*in))
		for key, val := range *in {
			var outVal *TaskLoopTaskRunStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(TaskLoopTaskRunStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskLoopRunStatus.
func (in *TaskLoopRunStatus) DeepCopy() *TaskLoopRunStatus {
	if in == nil {
		return nil
	}
	out := new(TaskLoopRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskLoopSpec) DeepCopyInto(out *TaskLoopSpec) {
	*out = *in
	if in.TaskRef != nil {
		in, out := &in.TaskRef, &out.TaskRef
		*out = new(v1beta1.TaskRef)
		**out = **in
	}
	if in.TaskSpec != nil {
		in, out := &in.TaskSpec, &out.TaskSpec
		*out = new(v1beta1.TaskSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskLoopSpec.
func (in *TaskLoopSpec) DeepCopy() *TaskLoopSpec {
	if in == nil {
		return nil
	}
	out := new(TaskLoopSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskLoopTaskRunStatus) DeepCopyInto(out *TaskLoopTaskRunStatus) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(v1beta1.TaskRunStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskLoopTaskRunStatus.
func (in *TaskLoopTaskRunStatus) DeepCopy() *TaskLoopTaskRunStatus {
	if in == nil {
		return nil
	}
	out := new(TaskLoopTaskRunStatus)
	in.DeepCopyInto(out)
	return out
}