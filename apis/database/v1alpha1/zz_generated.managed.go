/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	runtimev1alpha1 "github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

// GetBindingPhase of this CockroachCluster.
func (mg *CockroachCluster) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this CockroachCluster.
func (mg *CockroachCluster) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetNonPortableClassReference of this CockroachCluster.
func (mg *CockroachCluster) GetNonPortableClassReference() *corev1.ObjectReference {
	return mg.Spec.NonPortableClassReference
}

// GetReclaimPolicy of this CockroachCluster.
func (mg *CockroachCluster) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this CockroachCluster.
func (mg *CockroachCluster) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this CockroachCluster.
func (mg *CockroachCluster) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this CockroachCluster.
func (mg *CockroachCluster) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetConditions of this CockroachCluster.
func (mg *CockroachCluster) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetNonPortableClassReference of this CockroachCluster.
func (mg *CockroachCluster) SetNonPortableClassReference(r *corev1.ObjectReference) {
	mg.Spec.NonPortableClassReference = r
}

// SetReclaimPolicy of this CockroachCluster.
func (mg *CockroachCluster) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this CockroachCluster.
func (mg *CockroachCluster) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this YugabyteCluster.
func (mg *YugabyteCluster) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this YugabyteCluster.
func (mg *YugabyteCluster) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetNonPortableClassReference of this YugabyteCluster.
func (mg *YugabyteCluster) GetNonPortableClassReference() *corev1.ObjectReference {
	return mg.Spec.NonPortableClassReference
}

// GetReclaimPolicy of this YugabyteCluster.
func (mg *YugabyteCluster) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this YugabyteCluster.
func (mg *YugabyteCluster) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this YugabyteCluster.
func (mg *YugabyteCluster) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this YugabyteCluster.
func (mg *YugabyteCluster) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetConditions of this YugabyteCluster.
func (mg *YugabyteCluster) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetNonPortableClassReference of this YugabyteCluster.
func (mg *YugabyteCluster) SetNonPortableClassReference(r *corev1.ObjectReference) {
	mg.Spec.NonPortableClassReference = r
}

// SetReclaimPolicy of this YugabyteCluster.
func (mg *YugabyteCluster) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this YugabyteCluster.
func (mg *YugabyteCluster) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
