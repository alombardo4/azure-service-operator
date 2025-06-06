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

// +kubebuilder:rbac:groups=apimanagement.azure.com,resources=apis,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apimanagement.azure.com,resources={apis/status,apis/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220801.Api
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimapis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}
type Api struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Api_Spec   `json:"spec,omitempty"`
	Status            Api_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Api{}

// GetConditions returns the conditions of the resource
func (api *Api) GetConditions() conditions.Conditions {
	return api.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (api *Api) SetConditions(conditions conditions.Conditions) { api.Status.Conditions = conditions }

var _ configmaps.Exporter = &Api{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (api *Api) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if api.Spec.OperatorSpec == nil {
		return nil
	}
	return api.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Api{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (api *Api) SecretDestinationExpressions() []*core.DestinationExpression {
	if api.Spec.OperatorSpec == nil {
		return nil
	}
	return api.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Api{}

// AzureName returns the Azure name of the resource
func (api *Api) AzureName() string {
	return api.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (api Api) GetAPIVersion() string {
	return "2022-08-01"
}

// GetResourceScope returns the scope of the resource
func (api *Api) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (api *Api) GetSpec() genruntime.ConvertibleSpec {
	return &api.Spec
}

// GetStatus returns the status of this resource
func (api *Api) GetStatus() genruntime.ConvertibleStatus {
	return &api.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (api *Api) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/apis"
func (api *Api) GetType() string {
	return "Microsoft.ApiManagement/service/apis"
}

// NewEmptyStatus returns a new empty (blank) status
func (api *Api) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Api_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (api *Api) Owner() *genruntime.ResourceReference {
	if api.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(api.Spec)
	return api.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (api *Api) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Api_STATUS); ok {
		api.Status = *st
		return nil
	}

	// Convert status to required version
	var st Api_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	api.Status = st
	return nil
}

// Hub marks that this Api is the hub type for conversion
func (api *Api) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (api *Api) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: api.Spec.OriginalVersion,
		Kind:    "Api",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220801.Api
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimapis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}
type ApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Api `json:"items"`
}

// Storage version of v1api20220801.Api_Spec
type Api_Spec struct {
	APIVersion             *string                       `json:"apiVersion,omitempty"`
	ApiRevision            *string                       `json:"apiRevision,omitempty"`
	ApiRevisionDescription *string                       `json:"apiRevisionDescription,omitempty"`
	ApiType                *string                       `json:"apiType,omitempty"`
	ApiVersionDescription  *string                       `json:"apiVersionDescription,omitempty"`
	ApiVersionSet          *ApiVersionSetContractDetails `json:"apiVersionSet,omitempty"`

	// ApiVersionSetReference: A resource identifier for the related ApiVersionSet.
	ApiVersionSetReference *genruntime.ResourceReference   `armReference:"ApiVersionSetId" json:"apiVersionSetReference,omitempty"`
	AuthenticationSettings *AuthenticationSettingsContract `json:"authenticationSettings,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                 `json:"azureName,omitempty"`
	Contact         *ApiContactInformation `json:"contact,omitempty"`
	Description     *string                `json:"description,omitempty"`
	DisplayName     *string                `json:"displayName,omitempty"`
	Format          *string                `json:"format,omitempty"`
	IsCurrent       *bool                  `json:"isCurrent,omitempty"`
	License         *ApiLicenseInformation `json:"license,omitempty"`
	OperatorSpec    *ApiOperatorSpec       `json:"operatorSpec,omitempty"`
	OriginalVersion string                 `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner       *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	Path        *string                            `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Protocols   []string                           `json:"protocols,omitempty"`
	ServiceUrl  *string                            `json:"serviceUrl,omitempty"`

	// SourceApiReference: API identifier of the source API.
	SourceApiReference               *genruntime.ResourceReference             `armReference:"SourceApiId" json:"sourceApiReference,omitempty"`
	SubscriptionKeyParameterNames    *SubscriptionKeyParameterNamesContract    `json:"subscriptionKeyParameterNames,omitempty"`
	SubscriptionRequired             *bool                                     `json:"subscriptionRequired,omitempty"`
	TermsOfServiceUrl                *string                                   `json:"termsOfServiceUrl,omitempty"`
	TranslateRequiredQueryParameters *string                                   `json:"translateRequiredQueryParameters,omitempty"`
	Type                             *string                                   `json:"type,omitempty"`
	Value                            *string                                   `json:"value,omitempty"`
	WsdlSelector                     *ApiCreateOrUpdateProperties_WsdlSelector `json:"wsdlSelector,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Api_Spec{}

// ConvertSpecFrom populates our Api_Spec from the provided source
func (api *Api_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == api {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(api)
}

// ConvertSpecTo populates the provided destination from our Api_Spec
func (api *Api_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == api {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(api)
}

// Storage version of v1api20220801.Api_STATUS
type Api_STATUS struct {
	APIVersion                    *string                                       `json:"apiVersion,omitempty"`
	ApiRevision                   *string                                       `json:"apiRevision,omitempty"`
	ApiRevisionDescription        *string                                       `json:"apiRevisionDescription,omitempty"`
	ApiVersionDescription         *string                                       `json:"apiVersionDescription,omitempty"`
	ApiVersionSet                 *ApiVersionSetContractDetails_STATUS          `json:"apiVersionSet,omitempty"`
	ApiVersionSetId               *string                                       `json:"apiVersionSetId,omitempty"`
	AuthenticationSettings        *AuthenticationSettingsContract_STATUS        `json:"authenticationSettings,omitempty"`
	Conditions                    []conditions.Condition                        `json:"conditions,omitempty"`
	Contact                       *ApiContactInformation_STATUS                 `json:"contact,omitempty"`
	Description                   *string                                       `json:"description,omitempty"`
	DisplayName                   *string                                       `json:"displayName,omitempty"`
	Id                            *string                                       `json:"id,omitempty"`
	IsCurrent                     *bool                                         `json:"isCurrent,omitempty"`
	IsOnline                      *bool                                         `json:"isOnline,omitempty"`
	License                       *ApiLicenseInformation_STATUS                 `json:"license,omitempty"`
	Name                          *string                                       `json:"name,omitempty"`
	Path                          *string                                       `json:"path,omitempty"`
	PropertiesType                *string                                       `json:"properties_type,omitempty"`
	PropertyBag                   genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	Protocols                     []string                                      `json:"protocols,omitempty"`
	ServiceUrl                    *string                                       `json:"serviceUrl,omitempty"`
	SourceApiId                   *string                                       `json:"sourceApiId,omitempty"`
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContract_STATUS `json:"subscriptionKeyParameterNames,omitempty"`
	SubscriptionRequired          *bool                                         `json:"subscriptionRequired,omitempty"`
	TermsOfServiceUrl             *string                                       `json:"termsOfServiceUrl,omitempty"`
	Type                          *string                                       `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Api_STATUS{}

// ConvertStatusFrom populates our Api_STATUS from the provided source
func (api *Api_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == api {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(api)
}

// ConvertStatusTo populates the provided destination from our Api_STATUS
func (api *Api_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == api {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(api)
}

// Storage version of v1api20220801.APIVersion
// +kubebuilder:validation:Enum={"2022-08-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-08-01")

// Storage version of v1api20220801.ApiContactInformation
// API contact information
type ApiContactInformation struct {
	Email       *string                `json:"email,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Url         *string                `json:"url,omitempty"`
}

// Storage version of v1api20220801.ApiContactInformation_STATUS
// API contact information
type ApiContactInformation_STATUS struct {
	Email       *string                `json:"email,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Url         *string                `json:"url,omitempty"`
}

// Storage version of v1api20220801.ApiCreateOrUpdateProperties_WsdlSelector
type ApiCreateOrUpdateProperties_WsdlSelector struct {
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	WsdlEndpointName *string                `json:"wsdlEndpointName,omitempty"`
	WsdlServiceName  *string                `json:"wsdlServiceName,omitempty"`
}

// Storage version of v1api20220801.ApiLicenseInformation
// API license information
type ApiLicenseInformation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Url         *string                `json:"url,omitempty"`
}

// Storage version of v1api20220801.ApiLicenseInformation_STATUS
// API license information
type ApiLicenseInformation_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Url         *string                `json:"url,omitempty"`
}

// Storage version of v1api20220801.ApiOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ApiOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20220801.ApiVersionSetContractDetails
// An API Version Set contains the common configuration for a set of API Versions relating
type ApiVersionSetContractDetails struct {
	Description *string                `json:"description,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Identifier for existing API Version Set. Omit this value to create a new Version Set.
	Reference         *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
	VersionHeaderName *string                       `json:"versionHeaderName,omitempty"`
	VersionQueryName  *string                       `json:"versionQueryName,omitempty"`
	VersioningScheme  *string                       `json:"versioningScheme,omitempty"`
}

// Storage version of v1api20220801.ApiVersionSetContractDetails_STATUS
// An API Version Set contains the common configuration for a set of API Versions relating
type ApiVersionSetContractDetails_STATUS struct {
	Description       *string                `json:"description,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VersionHeaderName *string                `json:"versionHeaderName,omitempty"`
	VersionQueryName  *string                `json:"versionQueryName,omitempty"`
	VersioningScheme  *string                `json:"versioningScheme,omitempty"`
}

// Storage version of v1api20220801.AuthenticationSettingsContract
// API Authentication Settings.
type AuthenticationSettingsContract struct {
	OAuth2                       *OAuth2AuthenticationSettingsContract  `json:"oAuth2,omitempty"`
	OAuth2AuthenticationSettings []OAuth2AuthenticationSettingsContract `json:"oAuth2AuthenticationSettings,omitempty"`
	Openid                       *OpenIdAuthenticationSettingsContract  `json:"openid,omitempty"`
	OpenidAuthenticationSettings []OpenIdAuthenticationSettingsContract `json:"openidAuthenticationSettings,omitempty"`
	PropertyBag                  genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220801.AuthenticationSettingsContract_STATUS
// API Authentication Settings.
type AuthenticationSettingsContract_STATUS struct {
	OAuth2                       *OAuth2AuthenticationSettingsContract_STATUS  `json:"oAuth2,omitempty"`
	OAuth2AuthenticationSettings []OAuth2AuthenticationSettingsContract_STATUS `json:"oAuth2AuthenticationSettings,omitempty"`
	Openid                       *OpenIdAuthenticationSettingsContract_STATUS  `json:"openid,omitempty"`
	OpenidAuthenticationSettings []OpenIdAuthenticationSettingsContract_STATUS `json:"openidAuthenticationSettings,omitempty"`
	PropertyBag                  genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220801.SubscriptionKeyParameterNamesContract
// Subscription key parameter names details.
type SubscriptionKeyParameterNamesContract struct {
	Header      *string                `json:"header,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Query       *string                `json:"query,omitempty"`
}

// Storage version of v1api20220801.SubscriptionKeyParameterNamesContract_STATUS
// Subscription key parameter names details.
type SubscriptionKeyParameterNamesContract_STATUS struct {
	Header      *string                `json:"header,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Query       *string                `json:"query,omitempty"`
}

// Storage version of v1api20220801.OAuth2AuthenticationSettingsContract
// API OAuth2 Authentication settings details.
type OAuth2AuthenticationSettingsContract struct {
	AuthorizationServerId *string                `json:"authorizationServerId,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Scope                 *string                `json:"scope,omitempty"`
}

// Storage version of v1api20220801.OAuth2AuthenticationSettingsContract_STATUS
// API OAuth2 Authentication settings details.
type OAuth2AuthenticationSettingsContract_STATUS struct {
	AuthorizationServerId *string                `json:"authorizationServerId,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Scope                 *string                `json:"scope,omitempty"`
}

// Storage version of v1api20220801.OpenIdAuthenticationSettingsContract
// API OAuth2 Authentication settings details.
type OpenIdAuthenticationSettingsContract struct {
	BearerTokenSendingMethods []string               `json:"bearerTokenSendingMethods,omitempty"`
	OpenidProviderId          *string                `json:"openidProviderId,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220801.OpenIdAuthenticationSettingsContract_STATUS
// API OAuth2 Authentication settings details.
type OpenIdAuthenticationSettingsContract_STATUS struct {
	BearerTokenSendingMethods []string               `json:"bearerTokenSendingMethods,omitempty"`
	OpenidProviderId          *string                `json:"openidProviderId,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Api{}, &ApiList{})
}
