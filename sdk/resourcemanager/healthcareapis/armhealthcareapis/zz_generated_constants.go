//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

const (
	moduleName    = "armhealthcareapis"
	moduleVersion = "v0.4.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// FhirResourceVersionPolicy - Controls how resources are versioned on the FHIR service
type FhirResourceVersionPolicy string

const (
	FhirResourceVersionPolicyNoVersion       FhirResourceVersionPolicy = "no-version"
	FhirResourceVersionPolicyVersioned       FhirResourceVersionPolicy = "versioned"
	FhirResourceVersionPolicyVersionedUpdate FhirResourceVersionPolicy = "versioned-update"
)

// PossibleFhirResourceVersionPolicyValues returns the possible values for the FhirResourceVersionPolicy const type.
func PossibleFhirResourceVersionPolicyValues() []FhirResourceVersionPolicy {
	return []FhirResourceVersionPolicy{
		FhirResourceVersionPolicyNoVersion,
		FhirResourceVersionPolicyVersioned,
		FhirResourceVersionPolicyVersionedUpdate,
	}
}

// FhirServiceKind - The kind of the service.
type FhirServiceKind string

const (
	FhirServiceKindFhirR4   FhirServiceKind = "fhir-R4"
	FhirServiceKindFhirStu3 FhirServiceKind = "fhir-Stu3"
)

// PossibleFhirServiceKindValues returns the possible values for the FhirServiceKind const type.
func PossibleFhirServiceKindValues() []FhirServiceKind {
	return []FhirServiceKind{
		FhirServiceKindFhirR4,
		FhirServiceKindFhirStu3,
	}
}

// IotIdentityResolutionType - The type of IoT identity resolution to use with the destination.
type IotIdentityResolutionType string

const (
	IotIdentityResolutionTypeCreate IotIdentityResolutionType = "Create"
	IotIdentityResolutionTypeLookup IotIdentityResolutionType = "Lookup"
)

// PossibleIotIdentityResolutionTypeValues returns the possible values for the IotIdentityResolutionType const type.
func PossibleIotIdentityResolutionTypeValues() []IotIdentityResolutionType {
	return []IotIdentityResolutionType{
		IotIdentityResolutionTypeCreate,
		IotIdentityResolutionTypeLookup,
	}
}

// Kind - The kind of the service.
type Kind string

const (
	KindFhir     Kind = "fhir"
	KindFhirStu3 Kind = "fhir-Stu3"
	KindFhirR4   Kind = "fhir-R4"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindFhir,
		KindFhirStu3,
		KindFhirR4,
	}
}

// ManagedServiceIdentityType - Type of identity being specified, currently SystemAssigned and None are allowed.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone           ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned ManagedServiceIdentityType = "SystemAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
	}
}

// OperationResultStatus - The status of the operation being performed.
type OperationResultStatus string

const (
	OperationResultStatusCanceled  OperationResultStatus = "Canceled"
	OperationResultStatusFailed    OperationResultStatus = "Failed"
	OperationResultStatusRequested OperationResultStatus = "Requested"
	OperationResultStatusRunning   OperationResultStatus = "Running"
	OperationResultStatusSucceeded OperationResultStatus = "Succeeded"
)

// PossibleOperationResultStatusValues returns the possible values for the OperationResultStatus const type.
func PossibleOperationResultStatusValues() []OperationResultStatus {
	return []OperationResultStatus{
		OperationResultStatusCanceled,
		OperationResultStatusFailed,
		OperationResultStatusRequested,
		OperationResultStatusRunning,
		OperationResultStatusSucceeded,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ProvisioningState - The provisioning state.
type ProvisioningState string

const (
	ProvisioningStateAccepted          ProvisioningState = "Accepted"
	ProvisioningStateCanceled          ProvisioningState = "Canceled"
	ProvisioningStateCreating          ProvisioningState = "Creating"
	ProvisioningStateDeleting          ProvisioningState = "Deleting"
	ProvisioningStateDeprovisioned     ProvisioningState = "Deprovisioned"
	ProvisioningStateFailed            ProvisioningState = "Failed"
	ProvisioningStateMoving            ProvisioningState = "Moving"
	ProvisioningStateSucceeded         ProvisioningState = "Succeeded"
	ProvisioningStateSuspended         ProvisioningState = "Suspended"
	ProvisioningStateSystemMaintenance ProvisioningState = "SystemMaintenance"
	ProvisioningStateUpdating          ProvisioningState = "Updating"
	ProvisioningStateVerifying         ProvisioningState = "Verifying"
	ProvisioningStateWarned            ProvisioningState = "Warned"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateDeprovisioned,
		ProvisioningStateFailed,
		ProvisioningStateMoving,
		ProvisioningStateSucceeded,
		ProvisioningStateSuspended,
		ProvisioningStateSystemMaintenance,
		ProvisioningStateUpdating,
		ProvisioningStateVerifying,
		ProvisioningStateWarned,
	}
}

// PublicNetworkAccess - Control permission for data plane traffic coming from public networks while private endpoint is enabled.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ServiceEventState - Indicates the current status of event support for the resource.
type ServiceEventState string

const (
	ServiceEventStateDisabled ServiceEventState = "Disabled"
	ServiceEventStateEnabled  ServiceEventState = "Enabled"
	ServiceEventStateUpdating ServiceEventState = "Updating"
)

// PossibleServiceEventStateValues returns the possible values for the ServiceEventState const type.
func PossibleServiceEventStateValues() []ServiceEventState {
	return []ServiceEventState{
		ServiceEventStateDisabled,
		ServiceEventStateEnabled,
		ServiceEventStateUpdating,
	}
}

// ServiceManagedIdentityType - Type of identity being specified, currently SystemAssigned and None are allowed.
type ServiceManagedIdentityType string

const (
	ServiceManagedIdentityTypeNone                       ServiceManagedIdentityType = "None"
	ServiceManagedIdentityTypeSystemAssigned             ServiceManagedIdentityType = "SystemAssigned"
	ServiceManagedIdentityTypeSystemAssignedUserAssigned ServiceManagedIdentityType = "SystemAssigned,UserAssigned"
	ServiceManagedIdentityTypeUserAssigned               ServiceManagedIdentityType = "UserAssigned"
)

// PossibleServiceManagedIdentityTypeValues returns the possible values for the ServiceManagedIdentityType const type.
func PossibleServiceManagedIdentityTypeValues() []ServiceManagedIdentityType {
	return []ServiceManagedIdentityType{
		ServiceManagedIdentityTypeNone,
		ServiceManagedIdentityTypeSystemAssigned,
		ServiceManagedIdentityTypeSystemAssignedUserAssigned,
		ServiceManagedIdentityTypeUserAssigned,
	}
}

// ServiceNameUnavailabilityReason - The reason for unavailability.
type ServiceNameUnavailabilityReason string

const (
	ServiceNameUnavailabilityReasonInvalid       ServiceNameUnavailabilityReason = "Invalid"
	ServiceNameUnavailabilityReasonAlreadyExists ServiceNameUnavailabilityReason = "AlreadyExists"
)

// PossibleServiceNameUnavailabilityReasonValues returns the possible values for the ServiceNameUnavailabilityReason const type.
func PossibleServiceNameUnavailabilityReasonValues() []ServiceNameUnavailabilityReason {
	return []ServiceNameUnavailabilityReason{
		ServiceNameUnavailabilityReasonInvalid,
		ServiceNameUnavailabilityReasonAlreadyExists,
	}
}
