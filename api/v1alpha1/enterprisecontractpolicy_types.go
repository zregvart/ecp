/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "make" to regenerate code after modifying this file
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type EnterpriseContractVersionName string

// EnterpriseContractPolicySpec represents the desired state of EnterpriseContractPolicy
type EnterpriseContractPolicySpec struct {
	// Description text describing the the policy or it's intended use
	// +optional
	Description *string `json:"description"`
	// Revisions of the policy, there needs to be at least one
	// +kubebuilder:validation:MinProperties:=1
	Revisions map[EnterpriseContractVersionName]EnterpriseContractPolicyRevision `json:"revisions"`
}

// EnterpriseContractPolicyRevision represents version of an enterprise contract
type EnterpriseContractPolicyRevision struct {
	// EffectiveFrom represents the time when the policy comes into effect, if empty immediately
	// +kubebuilder:validation:Format:=date-time
	// +optional
	EffectiveFrom *metav1.Time `json:"effectiveFrom"`
	// Additive is true if this revision is additive to up to the last non-additive or first revision, false if it's non-additive and older revisions are not taken into consideration. This allows for changes to be incremental and removes duplication. `false` by default
	// +kubebuilder:default:=false
	Additive bool `json:"additive,omitempty"`
	// Sources is list of policy sources
	// +kubebuilder:validation:MinItems:=1
	Sources []PolicySource `json:"sources"`
	// Exceptions configures exceptions under which the policy is evaluated as successful even if the listed policy checks have reported failure
	Exceptions EnterpriseContractPolicyExceptions `json:"exceptions,omitempty"`
}

// PolicySource represents the configuration of the source for the policy
type PolicySource struct {
	// GitRepository configures fetching of the policies from a Git repository
	// +optional
	GitRepository *GitPolicySource `json:"git,omitempty"`
}

type GitPolicySource struct {
	// Repository URL
	Repository string `json:"repository"`
	// Revision matching the branch, commit id or similar to fetch. Defaults to `main`
	// +kubebuilder:default:=main
	// +optional
	Revision *string `json:"revision"`
}

// EnterpriseContractPolicyExceptions configuration of exceptions for the policy evaluation
type EnterpriseContractPolicyExceptions struct {
	// +optional
	NonBlocking []string `json:"nonBlocking,omitempty"`
}

// EnterpriseContractPolicyStatus defines the observed state of EnterpriseContractPolicy
type EnterpriseContractPolicyStatus struct {
	// EffectiveVersion the currently effective version
	EffectiveVersion EnterpriseContractVersionName `json:"effectiveVersion"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Effective version",type=string,JSONPath=`.spec.effectiveVersion`
// EnterpriseContractPolicy is the Schema for the enterprisecontractpolicies API
type EnterpriseContractPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnterpriseContractPolicySpec   `json:"spec,omitempty"`
	Status EnterpriseContractPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnterpriseContractPolicyList contains a list of EnterpriseContractPolicy
type EnterpriseContractPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnterpriseContractPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EnterpriseContractPolicy{}, &EnterpriseContractPolicyList{})
}
