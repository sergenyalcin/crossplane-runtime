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

package resource

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource/unstructured/claim"
)

// A Conditioned may have conditions set or retrieved. Conditions are typically
// indicate the status of both a resource and its reconciliation process.
type Conditioned interface {
	SetConditions(c ...xpv1.Condition)
	GetCondition(ct xpv1.ConditionType) xpv1.Condition
}

// A ClaimReferencer may reference a resource claim.
type ClaimReferencer interface {
	SetClaimReference(r *claim.Reference)
	GetClaimReference() *claim.Reference
}

// A ManagedResourceReferencer may reference a concrete managed resource.
type ManagedResourceReferencer interface {
	SetResourceReference(r *corev1.ObjectReference)
	GetResourceReference() *corev1.ObjectReference
}

// A LocalConnectionSecretWriterTo may write a connection secret to its own
// namespace.
type LocalConnectionSecretWriterTo interface {
	SetWriteConnectionSecretToReference(r *xpv1.LocalSecretReference)
	GetWriteConnectionSecretToReference() *xpv1.LocalSecretReference
}

// A ConnectionSecretWriterTo may write a connection secret to an arbitrary
// namespace.
type ConnectionSecretWriterTo interface {
	SetWriteConnectionSecretToReference(r *xpv1.SecretReference)
	GetWriteConnectionSecretToReference() *xpv1.SecretReference
}

// A ConnectionDetailsPublisherTo may write a connection details secret to a
// secret store.
type ConnectionDetailsPublisherTo interface {
	SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo)
	GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo
}

// A Manageable resource may specify a ManagementPolicies.
type Manageable interface {
	SetManagementPolicies(p xpv1.ManagementPolicies)
	GetManagementPolicies() xpv1.ManagementPolicies
}

// An Orphanable resource may specify a DeletionPolicy.
type Orphanable interface {
	SetDeletionPolicy(p xpv1.DeletionPolicy)
	GetDeletionPolicy() xpv1.DeletionPolicy
}

type CustomReconciliation interface {
	SetReconciliationPolicy(p *xpv1.ReconciliationPolicy)
	GetReconciliationPolicy() *xpv1.ReconciliationPolicy
}

// A ProviderConfigReferencer may reference a provider config resource.
type ProviderConfigReferencer interface {
	GetProviderConfigReference() *xpv1.Reference
	SetProviderConfigReference(p *xpv1.Reference)
}

// A RequiredProviderConfigReferencer may reference a provider config resource.
// Unlike ProviderConfigReferencer, the reference is required (i.e. not nil).
type RequiredProviderConfigReferencer interface {
	GetProviderConfigReference() xpv1.Reference
	SetProviderConfigReference(p xpv1.Reference)
}

// A RequiredTypedResourceReferencer can reference a resource.
type RequiredTypedResourceReferencer interface {
	SetResourceReference(r xpv1.TypedReference)
	GetResourceReference() xpv1.TypedReference
}

// A Finalizer manages the finalizers on the resource.
type Finalizer interface {
	AddFinalizer(ctx context.Context, obj Object) error
	RemoveFinalizer(ctx context.Context, obj Object) error
}

// A CompositionSelector may select a composition of resources.
type CompositionSelector interface {
	SetCompositionSelector(s *metav1.LabelSelector)
	GetCompositionSelector() *metav1.LabelSelector
}

// A CompositionReferencer may reference a composition of resources.
type CompositionReferencer interface {
	SetCompositionReference(ref *corev1.ObjectReference)
	GetCompositionReference() *corev1.ObjectReference
}

// A CompositionRevisionReferencer may reference a specific revision of a
// composition of resources.
type CompositionRevisionReferencer interface {
	SetCompositionRevisionReference(ref *corev1.ObjectReference)
	GetCompositionRevisionReference() *corev1.ObjectReference
}

// A CompositionRevisionSelector may reference a set of
// composition revisions.
type CompositionRevisionSelector interface {
	SetCompositionRevisionSelector(selector *metav1.LabelSelector)
	GetCompositionRevisionSelector() *metav1.LabelSelector
}

// A CompositionUpdater uses a composition, and may update which revision of
// that composition it uses.
type CompositionUpdater interface {
	SetCompositionUpdatePolicy(p *xpv1.UpdatePolicy)
	GetCompositionUpdatePolicy() *xpv1.UpdatePolicy
}

// A CompositeResourceDeleter creates a composite, and controls the policy
// used to delete the composite.
type CompositeResourceDeleter interface {
	SetCompositeDeletePolicy(policy *xpv1.CompositeDeletePolicy)
	GetCompositeDeletePolicy() *xpv1.CompositeDeletePolicy
}

// A ComposedResourcesReferencer may reference the resources it composes.
type ComposedResourcesReferencer interface {
	SetResourceReferences(refs []corev1.ObjectReference)
	GetResourceReferences() []corev1.ObjectReference
}

// A CompositeResourceReferencer can reference a composite resource.
type CompositeResourceReferencer interface {
	SetResourceReference(r *corev1.ObjectReference)
	GetResourceReference() *corev1.ObjectReference
}

// An EnvironmentConfigReferencer references a list of EnvironmentConfigs.
type EnvironmentConfigReferencer interface {
	SetEnvironmentConfigReferences(refs []corev1.ObjectReference)
	GetEnvironmentConfigReferences() []corev1.ObjectReference
}

// A UserCounter can count how many users it has.
type UserCounter interface {
	SetUsers(i int64)
	GetUsers() int64
}

// A ConnectionDetailsPublishedTimer can record the last time its connection
// details were published.
type ConnectionDetailsPublishedTimer interface {
	SetConnectionDetailsLastPublishedTime(t *metav1.Time)
	GetConnectionDetailsLastPublishedTime() *metav1.Time
}

// ReconciliationObserver can track data observed by resource reconciler.
type ReconciliationObserver interface {
	SetObservedGeneration(generation int64)
	GetObservedGeneration() int64
}

// An Object is a Kubernetes object.
type Object interface {
	metav1.Object
	runtime.Object
}

// A Managed is a Kubernetes object representing a concrete managed
// resource (e.g. a CloudSQL instance).
type Managed interface { //nolint:interfacebloat // This interface has to be big.
	Object

	ProviderConfigReferencer
	ConnectionSecretWriterTo
	ConnectionDetailsPublisherTo
	Manageable
	Orphanable
	CustomReconciliation

	Conditioned
}

// A ManagedList is a list of managed resources.
type ManagedList interface {
	client.ObjectList

	// GetItems returns the list of managed resources.
	GetItems() []Managed
}

// A ProviderConfig configures a Crossplane provider.
type ProviderConfig interface {
	Object

	UserCounter
	Conditioned
}

// A ProviderConfigUsage indicates a usage of a Crossplane provider config.
type ProviderConfigUsage interface {
	Object

	RequiredProviderConfigReferencer
	RequiredTypedResourceReferencer
}

// A ProviderConfigUsageList is a list of provider config usages.
type ProviderConfigUsageList interface {
	client.ObjectList

	// GetItems returns the list of provider config usages.
	GetItems() []ProviderConfigUsage
}

// A Composite resource composes one or more Composed resources.
type Composite interface { //nolint:interfacebloat // This interface has to be big.
	Object

	CompositionSelector
	CompositionReferencer
	CompositionUpdater
	CompositionRevisionReferencer
	CompositionRevisionSelector
	ComposedResourcesReferencer
	EnvironmentConfigReferencer
	ClaimReferencer
	ConnectionSecretWriterTo
	ConnectionDetailsPublisherTo

	Conditioned
	ConnectionDetailsPublishedTimer
	ReconciliationObserver
}

// Composed resources can be a composed into a Composite resource.
type Composed interface {
	Object

	Conditioned
	ConnectionSecretWriterTo
	ConnectionDetailsPublisherTo
	ReconciliationObserver
}

// A CompositeClaim for a Composite resource.
type CompositeClaim interface { //nolint:interfacebloat // This interface has to be big.
	Object

	CompositionSelector
	CompositionReferencer
	CompositionUpdater
	CompositionRevisionReferencer
	CompositionRevisionSelector
	CompositeResourceDeleter
	CompositeResourceReferencer
	LocalConnectionSecretWriterTo
	ConnectionDetailsPublisherTo

	Conditioned
	ConnectionDetailsPublishedTimer
	ReconciliationObserver
}
