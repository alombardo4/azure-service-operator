// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}
type DomainsTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainsTopic_Spec   `json:"spec,omitempty"`
	Status            DomainsTopic_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DomainsTopic{}

// GetConditions returns the conditions of the resource
func (topic *DomainsTopic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *DomainsTopic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ conversion.Convertible = &DomainsTopic{}

// ConvertFrom populates our DomainsTopic from the provided hub DomainsTopic
func (topic *DomainsTopic) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.DomainsTopic)
	if !ok {
		return fmt.Errorf("expected eventgrid/v1api20200601/storage/DomainsTopic but received %T instead", hub)
	}

	return topic.AssignProperties_From_DomainsTopic(source)
}

// ConvertTo populates the provided hub DomainsTopic from our DomainsTopic
func (topic *DomainsTopic) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.DomainsTopic)
	if !ok {
		return fmt.Errorf("expected eventgrid/v1api20200601/storage/DomainsTopic but received %T instead", hub)
	}

	return topic.AssignProperties_To_DomainsTopic(destination)
}

var _ configmaps.Exporter = &DomainsTopic{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (topic *DomainsTopic) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if topic.Spec.OperatorSpec == nil {
		return nil
	}
	return topic.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &DomainsTopic{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (topic *DomainsTopic) SecretDestinationExpressions() []*core.DestinationExpression {
	if topic.Spec.OperatorSpec == nil {
		return nil
	}
	return topic.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &DomainsTopic{}

// InitializeSpec initializes the spec for this resource from the given status
func (topic *DomainsTopic) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*DomainsTopic_STATUS); ok {
		return topic.Spec.Initialize_From_DomainsTopic_STATUS(s)
	}

	return fmt.Errorf("expected Status of type DomainsTopic_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &DomainsTopic{}

// AzureName returns the Azure name of the resource
func (topic *DomainsTopic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic DomainsTopic) GetAPIVersion() string {
	return "2020-06-01"
}

// GetResourceScope returns the scope of the resource
func (topic *DomainsTopic) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (topic *DomainsTopic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *DomainsTopic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (topic *DomainsTopic) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (topic *DomainsTopic) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *DomainsTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DomainsTopic_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (topic *DomainsTopic) Owner() *genruntime.ResourceReference {
	if topic.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return topic.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (topic *DomainsTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DomainsTopic_STATUS); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st DomainsTopic_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// AssignProperties_From_DomainsTopic populates our DomainsTopic from the provided source DomainsTopic
func (topic *DomainsTopic) AssignProperties_From_DomainsTopic(source *storage.DomainsTopic) error {

	// ObjectMeta
	topic.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DomainsTopic_Spec
	err := spec.AssignProperties_From_DomainsTopic_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_DomainsTopic_Spec() to populate field Spec")
	}
	topic.Spec = spec

	// Status
	var status DomainsTopic_STATUS
	err = status.AssignProperties_From_DomainsTopic_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_DomainsTopic_STATUS() to populate field Status")
	}
	topic.Status = status

	// No error
	return nil
}

// AssignProperties_To_DomainsTopic populates the provided destination DomainsTopic from our DomainsTopic
func (topic *DomainsTopic) AssignProperties_To_DomainsTopic(destination *storage.DomainsTopic) error {

	// ObjectMeta
	destination.ObjectMeta = *topic.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.DomainsTopic_Spec
	err := topic.Spec.AssignProperties_To_DomainsTopic_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_DomainsTopic_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.DomainsTopic_STATUS
	err = topic.Status.AssignProperties_To_DomainsTopic_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_DomainsTopic_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *DomainsTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion(),
		Kind:    "DomainsTopic",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}
type DomainsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainsTopic `json:"items"`
}

type DomainsTopic_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *DomainsTopicOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventgrid.azure.com/Domain resource
	Owner *genruntime.KnownResourceReference `group:"eventgrid.azure.com" json:"owner,omitempty" kind:"Domain"`
}

var _ genruntime.ARMTransformer = &DomainsTopic_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (topic *DomainsTopic_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if topic == nil {
		return nil, nil
	}
	result := &arm.DomainsTopic_Spec{}

	// Set property "Name":
	result.Name = resolved.Name
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (topic *DomainsTopic_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.DomainsTopic_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (topic *DomainsTopic_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.DomainsTopic_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.DomainsTopic_Spec, got %T", armInput)
	}

	// Set property "AzureName":
	topic.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	topic.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DomainsTopic_Spec{}

// ConvertSpecFrom populates our DomainsTopic_Spec from the provided source
func (topic *DomainsTopic_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.DomainsTopic_Spec)
	if ok {
		// Populate our instance from source
		return topic.AssignProperties_From_DomainsTopic_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.DomainsTopic_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = topic.AssignProperties_From_DomainsTopic_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DomainsTopic_Spec
func (topic *DomainsTopic_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.DomainsTopic_Spec)
	if ok {
		// Populate destination from our instance
		return topic.AssignProperties_To_DomainsTopic_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.DomainsTopic_Spec{}
	err := topic.AssignProperties_To_DomainsTopic_Spec(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_DomainsTopic_Spec populates our DomainsTopic_Spec from the provided source DomainsTopic_Spec
func (topic *DomainsTopic_Spec) AssignProperties_From_DomainsTopic_Spec(source *storage.DomainsTopic_Spec) error {

	// AzureName
	topic.AzureName = source.AzureName

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec DomainsTopicOperatorSpec
		err := operatorSpec.AssignProperties_From_DomainsTopicOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_DomainsTopicOperatorSpec() to populate field OperatorSpec")
		}
		topic.OperatorSpec = &operatorSpec
	} else {
		topic.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		topic.Owner = &owner
	} else {
		topic.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_DomainsTopic_Spec populates the provided destination DomainsTopic_Spec from our DomainsTopic_Spec
func (topic *DomainsTopic_Spec) AssignProperties_To_DomainsTopic_Spec(destination *storage.DomainsTopic_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = topic.AzureName

	// OperatorSpec
	if topic.OperatorSpec != nil {
		var operatorSpec storage.DomainsTopicOperatorSpec
		err := topic.OperatorSpec.AssignProperties_To_DomainsTopicOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_DomainsTopicOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = topic.OriginalVersion()

	// Owner
	if topic.Owner != nil {
		owner := topic.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_DomainsTopic_STATUS populates our DomainsTopic_Spec from the provided source DomainsTopic_STATUS
func (topic *DomainsTopic_Spec) Initialize_From_DomainsTopic_STATUS(source *DomainsTopic_STATUS) error {

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (topic *DomainsTopic_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (topic *DomainsTopic_Spec) SetAzureName(azureName string) { topic.AzureName = azureName }

type DomainsTopic_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	// Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	// ProvisioningState: Provisioning state of the domain topic.
	ProvisioningState *DomainTopicProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// SystemData: The system metadata relating to Domain Topic resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DomainsTopic_STATUS{}

// ConvertStatusFrom populates our DomainsTopic_STATUS from the provided source
func (topic *DomainsTopic_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.DomainsTopic_STATUS)
	if ok {
		// Populate our instance from source
		return topic.AssignProperties_From_DomainsTopic_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.DomainsTopic_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = topic.AssignProperties_From_DomainsTopic_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DomainsTopic_STATUS
func (topic *DomainsTopic_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.DomainsTopic_STATUS)
	if ok {
		// Populate destination from our instance
		return topic.AssignProperties_To_DomainsTopic_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.DomainsTopic_STATUS{}
	err := topic.AssignProperties_To_DomainsTopic_STATUS(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &DomainsTopic_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (topic *DomainsTopic_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.DomainsTopic_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (topic *DomainsTopic_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.DomainsTopic_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.DomainsTopic_STATUS, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		topic.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		topic.Name = &name
	}

	// Set property "ProvisioningState":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			var temp string
			temp = string(*typedInput.Properties.ProvisioningState)
			provisioningState := DomainTopicProperties_ProvisioningState_STATUS(temp)
			topic.ProvisioningState = &provisioningState
		}
	}

	// Set property "SystemData":
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		topic.SystemData = &systemData
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		topic.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_DomainsTopic_STATUS populates our DomainsTopic_STATUS from the provided source DomainsTopic_STATUS
func (topic *DomainsTopic_STATUS) AssignProperties_From_DomainsTopic_STATUS(source *storage.DomainsTopic_STATUS) error {

	// Conditions
	topic.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	topic.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	topic.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := *source.ProvisioningState
		provisioningStateTemp := genruntime.ToEnum(provisioningState, domainTopicProperties_ProvisioningState_STATUS_Values)
		topic.ProvisioningState = &provisioningStateTemp
	} else {
		topic.ProvisioningState = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		topic.SystemData = &systemDatum
	} else {
		topic.SystemData = nil
	}

	// Type
	topic.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_DomainsTopic_STATUS populates the provided destination DomainsTopic_STATUS from our DomainsTopic_STATUS
func (topic *DomainsTopic_STATUS) AssignProperties_To_DomainsTopic_STATUS(destination *storage.DomainsTopic_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(topic.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(topic.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(topic.Name)

	// ProvisioningState
	if topic.ProvisioningState != nil {
		provisioningState := string(*topic.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// SystemData
	if topic.SystemData != nil {
		var systemDatum storage.SystemData_STATUS
		err := topic.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(topic.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type DomainsTopicOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_DomainsTopicOperatorSpec populates our DomainsTopicOperatorSpec from the provided source DomainsTopicOperatorSpec
func (operator *DomainsTopicOperatorSpec) AssignProperties_From_DomainsTopicOperatorSpec(source *storage.DomainsTopicOperatorSpec) error {

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// No error
	return nil
}

// AssignProperties_To_DomainsTopicOperatorSpec populates the provided destination DomainsTopicOperatorSpec from our DomainsTopicOperatorSpec
func (operator *DomainsTopicOperatorSpec) AssignProperties_To_DomainsTopicOperatorSpec(destination *storage.DomainsTopicOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type DomainTopicProperties_ProvisioningState_STATUS string

const (
	DomainTopicProperties_ProvisioningState_STATUS_Canceled  = DomainTopicProperties_ProvisioningState_STATUS("Canceled")
	DomainTopicProperties_ProvisioningState_STATUS_Creating  = DomainTopicProperties_ProvisioningState_STATUS("Creating")
	DomainTopicProperties_ProvisioningState_STATUS_Deleting  = DomainTopicProperties_ProvisioningState_STATUS("Deleting")
	DomainTopicProperties_ProvisioningState_STATUS_Failed    = DomainTopicProperties_ProvisioningState_STATUS("Failed")
	DomainTopicProperties_ProvisioningState_STATUS_Succeeded = DomainTopicProperties_ProvisioningState_STATUS("Succeeded")
	DomainTopicProperties_ProvisioningState_STATUS_Updating  = DomainTopicProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to DomainTopicProperties_ProvisioningState_STATUS
var domainTopicProperties_ProvisioningState_STATUS_Values = map[string]DomainTopicProperties_ProvisioningState_STATUS{
	"canceled":  DomainTopicProperties_ProvisioningState_STATUS_Canceled,
	"creating":  DomainTopicProperties_ProvisioningState_STATUS_Creating,
	"deleting":  DomainTopicProperties_ProvisioningState_STATUS_Deleting,
	"failed":    DomainTopicProperties_ProvisioningState_STATUS_Failed,
	"succeeded": DomainTopicProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  DomainTopicProperties_ProvisioningState_STATUS_Updating,
}

func init() {
	SchemeBuilder.Register(&DomainsTopic{}, &DomainsTopicList{})
}
