// +build !ignore_autogenerated

/*
Copyright 2021 Red Hat, Inc.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/gateway-api/apis/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (in *API) DeepCopy() *API {
	if in == nil {
		return nil
	}
	out := new(API)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *API) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyAuth) DeepCopyInto(out *APIKeyAuth) {
	*out = *in
	in.CredentialSource.DeepCopyInto(&out.CredentialSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyAuth.
func (in *APIKeyAuth) DeepCopy() *APIKeyAuth {
	if in == nil {
		return nil
	}
	out := new(APIKeyAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyAuthCredentials) DeepCopyInto(out *APIKeyAuthCredentials) {
	*out = *in
	if in.LabelSelectors != nil {
		in, out := &in.LabelSelectors, &out.LabelSelectors
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyAuthCredentials.
func (in *APIKeyAuthCredentials) DeepCopy() *APIKeyAuthCredentials {
	if in == nil {
		return nil
	}
	out := new(APIKeyAuthCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIList) DeepCopyInto(out *APIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]API, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIList.
func (in *APIList) DeepCopy() *APIList {
	if in == nil {
		return nil
	}
	out := new(APIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIMappings) DeepCopyInto(out *APIMappings) {
	*out = *in
	if in.OAS != nil {
		in, out := &in.OAS, &out.OAS
		*out = new(string)
		**out = **in
	}
	if in.HTTPPathMatch != nil {
		in, out := &in.HTTPPathMatch, &out.HTTPPathMatch
		*out = new(v1alpha1.HTTPPathMatch)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIMappings.
func (in *APIMappings) DeepCopy() *APIMappings {
	if in == nil {
		return nil
	}
	out := new(APIMappings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIProduct) DeepCopyInto(out *APIProduct) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIProduct.
func (in *APIProduct) DeepCopy() *APIProduct {
	if in == nil {
		return nil
	}
	out := new(APIProduct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIProduct) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIProductList) DeepCopyInto(out *APIProductList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIProduct, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIProductList.
func (in *APIProductList) DeepCopy() *APIProductList {
	if in == nil {
		return nil
	}
	out := new(APIProductList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIProductList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIProductSpec) DeepCopyInto(out *APIProductSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityScheme != nil {
		in, out := &in.SecurityScheme, &out.SecurityScheme
		*out = make([]SecurityScheme, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.APIs != nil {
		in, out := &in.APIs, &out.APIs
		*out = make([]APIReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = new(RateLimitSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIProductSpec.
func (in *APIProductSpec) DeepCopy() *APIProductSpec {
	if in == nil {
		return nil
	}
	out := new(APIProductSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIProductStatus) DeepCopyInto(out *APIProductStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIProductStatus.
func (in *APIProductStatus) DeepCopy() *APIProductStatus {
	if in == nil {
		return nil
	}
	out := new(APIProductStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIReference) DeepCopyInto(out *APIReference) {
	*out = *in
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIReference.
func (in *APIReference) DeepCopy() *APIReference {
	if in == nil {
		return nil
	}
	out := new(APIReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISpec) DeepCopyInto(out *APISpec) {
	*out = *in
	in.Destination.DeepCopyInto(&out.Destination)
	in.Mappings.DeepCopyInto(&out.Mappings)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISpec.
func (in *APISpec) DeepCopy() *APISpec {
	if in == nil {
		return nil
	}
	out := new(APISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIStatus) DeepCopyInto(out *APIStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIStatus.
func (in *APIStatus) DeepCopy() *APIStatus {
	if in == nil {
		return nil
	}
	out := new(APIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	in.ServiceReference.DeepCopyInto(&out.ServiceReference)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenIDConnectAuth) DeepCopyInto(out *OpenIDConnectAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenIDConnectAuth.
func (in *OpenIDConnectAuth) DeepCopy() *OpenIDConnectAuth {
	if in == nil {
		return nil
	}
	out := new(OpenIDConnectAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitDefinitionSpec) DeepCopyInto(out *RateLimitDefinitionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitDefinitionSpec.
func (in *RateLimitDefinitionSpec) DeepCopy() *RateLimitDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitSpec) DeepCopyInto(out *RateLimitSpec) {
	*out = *in
	if in.GlobalRateLimit != nil {
		in, out := &in.GlobalRateLimit, &out.GlobalRateLimit
		*out = new(RateLimitDefinitionSpec)
		**out = **in
	}
	if in.PerRemoteIPRateLimit != nil {
		in, out := &in.PerRemoteIPRateLimit, &out.PerRemoteIPRateLimit
		*out = new(RateLimitDefinitionSpec)
		**out = **in
	}
	if in.AuthRateLimit != nil {
		in, out := &in.AuthRateLimit, &out.AuthRateLimit
		*out = new(RateLimitDefinitionSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitSpec.
func (in *RateLimitSpec) DeepCopy() *RateLimitSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityScheme) DeepCopyInto(out *SecurityScheme) {
	*out = *in
	if in.APIKeyAuth != nil {
		in, out := &in.APIKeyAuth, &out.APIKeyAuth
		*out = new(APIKeyAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.OpenIDConnectAuth != nil {
		in, out := &in.OpenIDConnectAuth, &out.OpenIDConnectAuth
		*out = new(OpenIDConnectAuth)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityScheme.
func (in *SecurityScheme) DeepCopy() *SecurityScheme {
	if in == nil {
		return nil
	}
	out := new(SecurityScheme)
	in.DeepCopyInto(out)
	return out
}