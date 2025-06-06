// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515/storage"
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
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default
type SqlDatabaseContainerThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseContainerThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            SqlDatabaseContainerThroughputSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *SqlDatabaseContainerThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerThroughputSetting{}

// ConvertFrom populates our SqlDatabaseContainerThroughputSetting from the provided hub SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source storage.SqlDatabaseContainerThroughputSetting

	err := source.ConvertFrom(hub)
	if err != nil {
		return eris.Wrap(err, "converting from hub to source")
	}

	err = setting.AssignProperties_From_SqlDatabaseContainerThroughputSetting(&source)
	if err != nil {
		return eris.Wrap(err, "converting from source to setting")
	}

	return nil
}

// ConvertTo populates the provided hub SqlDatabaseContainerThroughputSetting from our SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.SqlDatabaseContainerThroughputSetting
	err := setting.AssignProperties_To_SqlDatabaseContainerThroughputSetting(&destination)
	if err != nil {
		return eris.Wrap(err, "converting to destination from setting")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return eris.Wrap(err, "converting from destination to hub")
	}

	return nil
}

var _ configmaps.Exporter = &SqlDatabaseContainerThroughputSetting{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (setting *SqlDatabaseContainerThroughputSetting) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &SqlDatabaseContainerThroughputSetting{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (setting *SqlDatabaseContainerThroughputSetting) SecretDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *SqlDatabaseContainerThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting SqlDatabaseContainerThroughputSetting) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceScope returns the scope of the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *SqlDatabaseContainerThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *SqlDatabaseContainerThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
func (setting *SqlDatabaseContainerThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *SqlDatabaseContainerThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlDatabaseContainerThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *SqlDatabaseContainerThroughputSetting) Owner() *genruntime.ResourceReference {
	if setting.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *SqlDatabaseContainerThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlDatabaseContainerThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlDatabaseContainerThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// AssignProperties_From_SqlDatabaseContainerThroughputSetting populates our SqlDatabaseContainerThroughputSetting from the provided source SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) AssignProperties_From_SqlDatabaseContainerThroughputSetting(source *storage.SqlDatabaseContainerThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec SqlDatabaseContainerThroughputSetting_Spec
	err := spec.AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status SqlDatabaseContainerThroughputSetting_STATUS
	err = status.AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseContainerThroughputSetting populates the provided destination SqlDatabaseContainerThroughputSetting from our SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) AssignProperties_To_SqlDatabaseContainerThroughputSetting(destination *storage.SqlDatabaseContainerThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.SqlDatabaseContainerThroughputSetting_Spec
	err := setting.Spec.AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.SqlDatabaseContainerThroughputSetting_STATUS
	err = setting.Status.AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *SqlDatabaseContainerThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseContainerThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default
type SqlDatabaseContainerThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerThroughputSetting `json:"items"`
}

type SqlDatabaseContainerThroughputSetting_Spec struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *SqlDatabaseContainerThroughputSettingOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`

	// +kubebuilder:validation:Required
	// Resource: The standard JSON format of a resource throughput
	Resource *ThroughputSettingsResource `json:"resource,omitempty"`
	Tags     map[string]string           `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &SqlDatabaseContainerThroughputSetting_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (setting *SqlDatabaseContainerThroughputSetting_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if setting == nil {
		return nil, nil
	}
	result := &arm.SqlDatabaseContainerThroughputSetting_Spec{}

	// Set property "Location":
	if setting.Location != nil {
		location := *setting.Location
		result.Location = &location
	}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if setting.Resource != nil {
		result.Properties = &arm.ThroughputSettingsUpdateProperties{}
	}
	if setting.Resource != nil {
		resource_ARM, err := (*setting.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := *resource_ARM.(*arm.ThroughputSettingsResource)
		result.Properties.Resource = &resource
	}

	// Set property "Tags":
	if setting.Tags != nil {
		result.Tags = make(map[string]string, len(setting.Tags))
		for key, value := range setting.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *SqlDatabaseContainerThroughputSetting_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.SqlDatabaseContainerThroughputSetting_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *SqlDatabaseContainerThroughputSetting_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.SqlDatabaseContainerThroughputSetting_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.SqlDatabaseContainerThroughputSetting_Spec, got %T", armInput)
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		setting.Location = &location
	}

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	setting.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Resource":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			setting.Resource = &resource
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		setting.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			setting.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &SqlDatabaseContainerThroughputSetting_Spec{}

// ConvertSpecFrom populates our SqlDatabaseContainerThroughputSetting_Spec from the provided source
func (setting *SqlDatabaseContainerThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.SqlDatabaseContainerThroughputSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlDatabaseContainerThroughputSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our SqlDatabaseContainerThroughputSetting_Spec
func (setting *SqlDatabaseContainerThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.SqlDatabaseContainerThroughputSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlDatabaseContainerThroughputSetting_Spec{}
	err := setting.AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec(dst)
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

// AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec populates our SqlDatabaseContainerThroughputSetting_Spec from the provided source SqlDatabaseContainerThroughputSetting_Spec
func (setting *SqlDatabaseContainerThroughputSetting_Spec) AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec(source *storage.SqlDatabaseContainerThroughputSetting_Spec) error {

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec SqlDatabaseContainerThroughputSettingOperatorSpec
		err := operatorSpec.AssignProperties_From_SqlDatabaseContainerThroughputSettingOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_SqlDatabaseContainerThroughputSettingOperatorSpec() to populate field OperatorSpec")
		}
		setting.OperatorSpec = &operatorSpec
	} else {
		setting.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		setting.Owner = &owner
	} else {
		setting.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignProperties_From_ThroughputSettingsResource(source.Resource)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_ThroughputSettingsResource() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec populates the provided destination SqlDatabaseContainerThroughputSetting_Spec from our SqlDatabaseContainerThroughputSetting_Spec
func (setting *SqlDatabaseContainerThroughputSetting_Spec) AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec(destination *storage.SqlDatabaseContainerThroughputSetting_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// OperatorSpec
	if setting.OperatorSpec != nil {
		var operatorSpec storage.SqlDatabaseContainerThroughputSettingOperatorSpec
		err := setting.OperatorSpec.AssignProperties_To_SqlDatabaseContainerThroughputSettingOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SqlDatabaseContainerThroughputSettingOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = setting.OriginalVersion()

	// Owner
	if setting.Owner != nil {
		owner := setting.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if setting.Resource != nil {
		var resource storage.ThroughputSettingsResource
		err := setting.Resource.AssignProperties_To_ThroughputSettingsResource(&resource)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_ThroughputSettingsResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (setting *SqlDatabaseContainerThroughputSetting_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type SqlDatabaseContainerThroughputSetting_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: The name of the ARM resource.
	Name     *string                                          `json:"name,omitempty"`
	Resource *ThroughputSettingsGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags     map[string]string                                `json:"tags,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlDatabaseContainerThroughputSetting_STATUS{}

// ConvertStatusFrom populates our SqlDatabaseContainerThroughputSetting_STATUS from the provided source
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.SqlDatabaseContainerThroughputSetting_STATUS)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlDatabaseContainerThroughputSetting_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlDatabaseContainerThroughputSetting_STATUS
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.SqlDatabaseContainerThroughputSetting_STATUS)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlDatabaseContainerThroughputSetting_STATUS{}
	err := setting.AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS(dst)
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

var _ genruntime.FromARMConverter = &SqlDatabaseContainerThroughputSetting_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.SqlDatabaseContainerThroughputSetting_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.SqlDatabaseContainerThroughputSetting_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.SqlDatabaseContainerThroughputSetting_STATUS, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		setting.Id = &id
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		setting.Location = &location
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		setting.Name = &name
	}

	// Set property "Resource":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsGetProperties_Resource_STATUS
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			setting.Resource = &resource
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		setting.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			setting.Tags[key] = value
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		setting.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS populates our SqlDatabaseContainerThroughputSetting_STATUS from the provided source SqlDatabaseContainerThroughputSetting_STATUS
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS(source *storage.SqlDatabaseContainerThroughputSetting_STATUS) error {

	// Conditions
	setting.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	setting.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	setting.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsGetProperties_Resource_STATUS
		err := resource.AssignProperties_From_ThroughputSettingsGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	setting.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS populates the provided destination SqlDatabaseContainerThroughputSetting_STATUS from our SqlDatabaseContainerThroughputSetting_STATUS
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS(destination *storage.SqlDatabaseContainerThroughputSetting_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(setting.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(setting.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(setting.Name)

	// Resource
	if setting.Resource != nil {
		var resource storage.ThroughputSettingsGetProperties_Resource_STATUS
		err := setting.Resource.AssignProperties_To_ThroughputSettingsGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(setting.Type)

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
type SqlDatabaseContainerThroughputSettingOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_SqlDatabaseContainerThroughputSettingOperatorSpec populates our SqlDatabaseContainerThroughputSettingOperatorSpec from the provided source SqlDatabaseContainerThroughputSettingOperatorSpec
func (operator *SqlDatabaseContainerThroughputSettingOperatorSpec) AssignProperties_From_SqlDatabaseContainerThroughputSettingOperatorSpec(source *storage.SqlDatabaseContainerThroughputSettingOperatorSpec) error {

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

// AssignProperties_To_SqlDatabaseContainerThroughputSettingOperatorSpec populates the provided destination SqlDatabaseContainerThroughputSettingOperatorSpec from our SqlDatabaseContainerThroughputSettingOperatorSpec
func (operator *SqlDatabaseContainerThroughputSettingOperatorSpec) AssignProperties_To_SqlDatabaseContainerThroughputSettingOperatorSpec(destination *storage.SqlDatabaseContainerThroughputSettingOperatorSpec) error {
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

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerThroughputSetting{}, &SqlDatabaseContainerThroughputSettingList{})
}
