/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha3

import "fmt"

const (
	// ClusterctlLabelName is applied to all components managed by clusterctl.
	ClusterctlLabelName = "clusterctl.cluster.x-k8s.io"

	// ClusterctlCoreLabelName is applied to all the core objects managed by clusterctl.
	ClusterctlCoreLabelName = "clusterctl.cluster.x-k8s.io/core"

	// ClusterctlResourceLifecyleLabelName describes the lifecyle for a specific resource.
	//
	// Example: resources shared between instances of the same provider:  CRDs,
	// ValidatingWebhookConfiguration, MutatingWebhookConfiguration, and so on.
	ClusterctlResourceLifecyleLabelName = "clusterctl.cluster.x-k8s.io/lifecycle"
)

// ResourceLifecycle configures the lifecycle of a resource
type ResourceLifecycle string

const (
	// ResourceLifecycleShared is used to indicate that a resource is shared between
	// multiple instances of a provider.
	ResourceLifecycleShared = ResourceLifecycle("shared")
)

// ManifestLabel returns the cluster.x-k8s.io/provider label value for a provider/type.
//
// Note: the label uniquely describes the provider type and its kind (e.g. bootstrap-kubeadm);
// it's not meant to be used to describe each instance of a particular provider.
func ManifestLabel(name string, providerType ProviderType) string {
	switch providerType {
	case CoreProviderType:
		return name
	case BootstrapProviderType:
		return fmt.Sprintf("bootstrap-%s", name)
	case ControlPlaneProviderType:
		return fmt.Sprintf("control-plane-%s", name)
	case InfrastructureProviderType:
		return fmt.Sprintf("infrastructure-%s", name)
	default:
		return fmt.Sprintf("unknown-type-%s", name)
	}
}
