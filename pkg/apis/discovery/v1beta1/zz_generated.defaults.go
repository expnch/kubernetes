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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1beta1

import (
	discoveryv1beta1 "k8s.io/api/discovery/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&discoveryv1beta1.EndpointSlice{}, func(obj interface{}) { SetObjectDefaults_EndpointSlice(obj.(*discoveryv1beta1.EndpointSlice)) })
	scheme.AddTypeDefaultingFunc(&discoveryv1beta1.EndpointSliceList{}, func(obj interface{}) { SetObjectDefaults_EndpointSliceList(obj.(*discoveryv1beta1.EndpointSliceList)) })
	return nil
}

func SetObjectDefaults_EndpointSlice(in *discoveryv1beta1.EndpointSlice) {
	for i := range in.Ports {
		a := &in.Ports[i]
		SetDefaults_EndpointPort(a)
	}
}

func SetObjectDefaults_EndpointSliceList(in *discoveryv1beta1.EndpointSliceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_EndpointSlice(a)
	}
}
