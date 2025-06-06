// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversdatabasesbackuplongtermretentionpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversdatabasesbackuplongtermretentionpolicies/status,serversdatabasesbackuplongtermretentionpolicies/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.ServersDatabasesBackupLongTermRetentionPolicy
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/LongTermRetentionPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupLongTermRetentionPolicies/default
type ServersDatabasesBackupLongTermRetentionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServersDatabasesBackupLongTermRetentionPolicy_Spec   `json:"spec,omitempty"`
	Status            ServersDatabasesBackupLongTermRetentionPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabasesBackupLongTermRetentionPolicy{}

// GetConditions returns the conditions of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ configmaps.Exporter = &ServersDatabasesBackupLongTermRetentionPolicy{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ServersDatabasesBackupLongTermRetentionPolicy{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) SecretDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &ServersDatabasesBackupLongTermRetentionPolicy{}

// AzureName returns the Azure name of the resource (always "default")
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersDatabasesBackupLongTermRetentionPolicy) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceScope returns the scope of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/backupLongTermRetentionPolicies"
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetType() string {
	return "Microsoft.Sql/servers/databases/backupLongTermRetentionPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ServersDatabasesBackupLongTermRetentionPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) Owner() *genruntime.ResourceReference {
	if policy.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ServersDatabasesBackupLongTermRetentionPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st ServersDatabasesBackupLongTermRetentionPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// Hub marks that this ServersDatabasesBackupLongTermRetentionPolicy is the hub type for conversion
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion,
		Kind:    "ServersDatabasesBackupLongTermRetentionPolicy",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.ServersDatabasesBackupLongTermRetentionPolicy
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/LongTermRetentionPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupLongTermRetentionPolicies/default
type ServersDatabasesBackupLongTermRetentionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabasesBackupLongTermRetentionPolicy `json:"items"`
}

// Storage version of v1api20211101.ServersDatabasesBackupLongTermRetentionPolicy_Spec
type ServersDatabasesBackupLongTermRetentionPolicy_Spec struct {
	MonthlyRetention *string                                                    `json:"monthlyRetention,omitempty"`
	OperatorSpec     *ServersDatabasesBackupLongTermRetentionPolicyOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion  string                                                     `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner           *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`
	PropertyBag     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	WeekOfYear      *int                               `json:"weekOfYear,omitempty"`
	WeeklyRetention *string                            `json:"weeklyRetention,omitempty"`
	YearlyRetention *string                            `json:"yearlyRetention,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ServersDatabasesBackupLongTermRetentionPolicy_Spec{}

// ConvertSpecFrom populates our ServersDatabasesBackupLongTermRetentionPolicy_Spec from the provided source
func (policy *ServersDatabasesBackupLongTermRetentionPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(policy)
}

// ConvertSpecTo populates the provided destination from our ServersDatabasesBackupLongTermRetentionPolicy_Spec
func (policy *ServersDatabasesBackupLongTermRetentionPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(policy)
}

// Storage version of v1api20211101.ServersDatabasesBackupLongTermRetentionPolicy_STATUS
type ServersDatabasesBackupLongTermRetentionPolicy_STATUS struct {
	Conditions       []conditions.Condition `json:"conditions,omitempty"`
	Id               *string                `json:"id,omitempty"`
	MonthlyRetention *string                `json:"monthlyRetention,omitempty"`
	Name             *string                `json:"name,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type             *string                `json:"type,omitempty"`
	WeekOfYear       *int                   `json:"weekOfYear,omitempty"`
	WeeklyRetention  *string                `json:"weeklyRetention,omitempty"`
	YearlyRetention  *string                `json:"yearlyRetention,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ServersDatabasesBackupLongTermRetentionPolicy_STATUS{}

// ConvertStatusFrom populates our ServersDatabasesBackupLongTermRetentionPolicy_STATUS from the provided source
func (policy *ServersDatabasesBackupLongTermRetentionPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(policy)
}

// ConvertStatusTo populates the provided destination from our ServersDatabasesBackupLongTermRetentionPolicy_STATUS
func (policy *ServersDatabasesBackupLongTermRetentionPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(policy)
}

// Storage version of v1api20211101.ServersDatabasesBackupLongTermRetentionPolicyOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ServersDatabasesBackupLongTermRetentionPolicyOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ServersDatabasesBackupLongTermRetentionPolicy{}, &ServersDatabasesBackupLongTermRetentionPolicyList{})
}
