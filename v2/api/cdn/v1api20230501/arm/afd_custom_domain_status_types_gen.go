// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type AfdCustomDomain_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties of the domain to create.
	Properties *AFDDomainProperties_STATUS `json:"properties,omitempty"`

	// SystemData: Read only system data
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The JSON object that contains the properties of the domain to create.
type AFDDomainProperties_STATUS struct {
	// AzureDnsZone: Resource reference to the Azure DNS zone
	AzureDnsZone     *ResourceReference_STATUS                    `json:"azureDnsZone,omitempty"`
	DeploymentStatus *AFDDomainProperties_DeploymentStatus_STATUS `json:"deploymentStatus,omitempty"`

	// DomainValidationState: Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step.
	// DCV stands for DomainControlValidation.
	DomainValidationState *AFDDomainProperties_DomainValidationState_STATUS `json:"domainValidationState,omitempty"`

	// ExtendedProperties: Key-Value pair representing migration properties for domains.
	ExtendedProperties map[string]string `json:"extendedProperties,omitempty"`

	// HostName: The host name of the domain. Must be a domain name.
	HostName *string `json:"hostName,omitempty"`

	// PreValidatedCustomDomainResourceId: Resource reference to the Azure resource where custom domain ownership was
	// prevalidated
	PreValidatedCustomDomainResourceId *ResourceReference_STATUS `json:"preValidatedCustomDomainResourceId,omitempty"`

	// ProfileName: The name of the profile which holds the domain.
	ProfileName *string `json:"profileName,omitempty"`

	// ProvisioningState: Provisioning status
	ProvisioningState *AFDDomainProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// TlsSettings: The configuration specifying how to enable HTTPS for the domain - using AzureFrontDoor managed certificate
	// or user's own certificate. If not specified, enabling ssl uses AzureFrontDoor managed certificate by default.
	TlsSettings *AFDDomainHttpsParameters_STATUS `json:"tlsSettings,omitempty"`

	// ValidationProperties: Values the customer needs to validate domain ownership
	ValidationProperties *DomainValidationProperties_STATUS `json:"validationProperties,omitempty"`
}

// Read only system data
type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC)
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: An identifier for the identity that created the resource
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource
	CreatedByType *IdentityType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: An identifier for the identity that last modified the resource
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource
	LastModifiedByType *IdentityType_STATUS `json:"lastModifiedByType,omitempty"`
}

// The JSON object that contains the properties to secure a domain.
type AFDDomainHttpsParameters_STATUS struct {
	// CertificateType: Defines the source of the SSL certificate.
	CertificateType *AFDDomainHttpsParameters_CertificateType_STATUS `json:"certificateType,omitempty"`

	// MinimumTlsVersion: TLS protocol version that will be used for Https
	MinimumTlsVersion *AFDDomainHttpsParameters_MinimumTlsVersion_STATUS `json:"minimumTlsVersion,omitempty"`

	// Secret: Resource reference to the secret. ie. subs/rg/profile/secret
	Secret *ResourceReference_STATUS `json:"secret,omitempty"`
}

type AFDDomainProperties_DeploymentStatus_STATUS string

const (
	AFDDomainProperties_DeploymentStatus_STATUS_Failed     = AFDDomainProperties_DeploymentStatus_STATUS("Failed")
	AFDDomainProperties_DeploymentStatus_STATUS_InProgress = AFDDomainProperties_DeploymentStatus_STATUS("InProgress")
	AFDDomainProperties_DeploymentStatus_STATUS_NotStarted = AFDDomainProperties_DeploymentStatus_STATUS("NotStarted")
	AFDDomainProperties_DeploymentStatus_STATUS_Succeeded  = AFDDomainProperties_DeploymentStatus_STATUS("Succeeded")
)

// Mapping from string to AFDDomainProperties_DeploymentStatus_STATUS
var aFDDomainProperties_DeploymentStatus_STATUS_Values = map[string]AFDDomainProperties_DeploymentStatus_STATUS{
	"failed":     AFDDomainProperties_DeploymentStatus_STATUS_Failed,
	"inprogress": AFDDomainProperties_DeploymentStatus_STATUS_InProgress,
	"notstarted": AFDDomainProperties_DeploymentStatus_STATUS_NotStarted,
	"succeeded":  AFDDomainProperties_DeploymentStatus_STATUS_Succeeded,
}

type AFDDomainProperties_DomainValidationState_STATUS string

const (
	AFDDomainProperties_DomainValidationState_STATUS_Approved                  = AFDDomainProperties_DomainValidationState_STATUS("Approved")
	AFDDomainProperties_DomainValidationState_STATUS_InternalError             = AFDDomainProperties_DomainValidationState_STATUS("InternalError")
	AFDDomainProperties_DomainValidationState_STATUS_Pending                   = AFDDomainProperties_DomainValidationState_STATUS("Pending")
	AFDDomainProperties_DomainValidationState_STATUS_PendingRevalidation       = AFDDomainProperties_DomainValidationState_STATUS("PendingRevalidation")
	AFDDomainProperties_DomainValidationState_STATUS_RefreshingValidationToken = AFDDomainProperties_DomainValidationState_STATUS("RefreshingValidationToken")
	AFDDomainProperties_DomainValidationState_STATUS_Rejected                  = AFDDomainProperties_DomainValidationState_STATUS("Rejected")
	AFDDomainProperties_DomainValidationState_STATUS_Submitting                = AFDDomainProperties_DomainValidationState_STATUS("Submitting")
	AFDDomainProperties_DomainValidationState_STATUS_TimedOut                  = AFDDomainProperties_DomainValidationState_STATUS("TimedOut")
	AFDDomainProperties_DomainValidationState_STATUS_Unknown                   = AFDDomainProperties_DomainValidationState_STATUS("Unknown")
)

// Mapping from string to AFDDomainProperties_DomainValidationState_STATUS
var aFDDomainProperties_DomainValidationState_STATUS_Values = map[string]AFDDomainProperties_DomainValidationState_STATUS{
	"approved":                  AFDDomainProperties_DomainValidationState_STATUS_Approved,
	"internalerror":             AFDDomainProperties_DomainValidationState_STATUS_InternalError,
	"pending":                   AFDDomainProperties_DomainValidationState_STATUS_Pending,
	"pendingrevalidation":       AFDDomainProperties_DomainValidationState_STATUS_PendingRevalidation,
	"refreshingvalidationtoken": AFDDomainProperties_DomainValidationState_STATUS_RefreshingValidationToken,
	"rejected":                  AFDDomainProperties_DomainValidationState_STATUS_Rejected,
	"submitting":                AFDDomainProperties_DomainValidationState_STATUS_Submitting,
	"timedout":                  AFDDomainProperties_DomainValidationState_STATUS_TimedOut,
	"unknown":                   AFDDomainProperties_DomainValidationState_STATUS_Unknown,
}

type AFDDomainProperties_ProvisioningState_STATUS string

const (
	AFDDomainProperties_ProvisioningState_STATUS_Creating  = AFDDomainProperties_ProvisioningState_STATUS("Creating")
	AFDDomainProperties_ProvisioningState_STATUS_Deleting  = AFDDomainProperties_ProvisioningState_STATUS("Deleting")
	AFDDomainProperties_ProvisioningState_STATUS_Failed    = AFDDomainProperties_ProvisioningState_STATUS("Failed")
	AFDDomainProperties_ProvisioningState_STATUS_Succeeded = AFDDomainProperties_ProvisioningState_STATUS("Succeeded")
	AFDDomainProperties_ProvisioningState_STATUS_Updating  = AFDDomainProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to AFDDomainProperties_ProvisioningState_STATUS
var aFDDomainProperties_ProvisioningState_STATUS_Values = map[string]AFDDomainProperties_ProvisioningState_STATUS{
	"creating":  AFDDomainProperties_ProvisioningState_STATUS_Creating,
	"deleting":  AFDDomainProperties_ProvisioningState_STATUS_Deleting,
	"failed":    AFDDomainProperties_ProvisioningState_STATUS_Failed,
	"succeeded": AFDDomainProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  AFDDomainProperties_ProvisioningState_STATUS_Updating,
}

// The JSON object that contains the properties to validate a domain.
type DomainValidationProperties_STATUS struct {
	// ExpirationDate: The date time that the token expires
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// ValidationToken: Challenge used for DNS TXT record or file based validation
	ValidationToken *string `json:"validationToken,omitempty"`
}

// The type of identity that creates/modifies resources
type IdentityType_STATUS string

const (
	IdentityType_STATUS_Application     = IdentityType_STATUS("application")
	IdentityType_STATUS_Key             = IdentityType_STATUS("key")
	IdentityType_STATUS_ManagedIdentity = IdentityType_STATUS("managedIdentity")
	IdentityType_STATUS_User            = IdentityType_STATUS("user")
)

// Mapping from string to IdentityType_STATUS
var identityType_STATUS_Values = map[string]IdentityType_STATUS{
	"application":     IdentityType_STATUS_Application,
	"key":             IdentityType_STATUS_Key,
	"managedidentity": IdentityType_STATUS_ManagedIdentity,
	"user":            IdentityType_STATUS_User,
}

// Reference to another resource.
type ResourceReference_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type AFDDomainHttpsParameters_CertificateType_STATUS string

const (
	AFDDomainHttpsParameters_CertificateType_STATUS_AzureFirstPartyManagedCertificate = AFDDomainHttpsParameters_CertificateType_STATUS("AzureFirstPartyManagedCertificate")
	AFDDomainHttpsParameters_CertificateType_STATUS_CustomerCertificate               = AFDDomainHttpsParameters_CertificateType_STATUS("CustomerCertificate")
	AFDDomainHttpsParameters_CertificateType_STATUS_ManagedCertificate                = AFDDomainHttpsParameters_CertificateType_STATUS("ManagedCertificate")
)

// Mapping from string to AFDDomainHttpsParameters_CertificateType_STATUS
var aFDDomainHttpsParameters_CertificateType_STATUS_Values = map[string]AFDDomainHttpsParameters_CertificateType_STATUS{
	"azurefirstpartymanagedcertificate": AFDDomainHttpsParameters_CertificateType_STATUS_AzureFirstPartyManagedCertificate,
	"customercertificate":               AFDDomainHttpsParameters_CertificateType_STATUS_CustomerCertificate,
	"managedcertificate":                AFDDomainHttpsParameters_CertificateType_STATUS_ManagedCertificate,
}

type AFDDomainHttpsParameters_MinimumTlsVersion_STATUS string

const (
	AFDDomainHttpsParameters_MinimumTlsVersion_STATUS_TLS10 = AFDDomainHttpsParameters_MinimumTlsVersion_STATUS("TLS10")
	AFDDomainHttpsParameters_MinimumTlsVersion_STATUS_TLS12 = AFDDomainHttpsParameters_MinimumTlsVersion_STATUS("TLS12")
)

// Mapping from string to AFDDomainHttpsParameters_MinimumTlsVersion_STATUS
var aFDDomainHttpsParameters_MinimumTlsVersion_STATUS_Values = map[string]AFDDomainHttpsParameters_MinimumTlsVersion_STATUS{
	"tls10": AFDDomainHttpsParameters_MinimumTlsVersion_STATUS_TLS10,
	"tls12": AFDDomainHttpsParameters_MinimumTlsVersion_STATUS_TLS12,
}
