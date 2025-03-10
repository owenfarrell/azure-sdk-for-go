//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

import "time"

// AvailableContacts - Customer retrieves list of Available Contacts for a spacecraft resource. Later, one of the available
// contact can be selected to create a contact.
type AvailableContacts struct {
	// Properties of Contact resource.
	Properties *ContactInstanceProperties `json:"properties,omitempty"`

	// The reference to the spacecraft resource.
	Spacecraft *ResourceReference `json:"spacecraft,omitempty"`

	// READ-ONLY; Name of Azure Ground Station.
	GroundStationName *string `json:"groundStationName,omitempty" azure:"ro"`
}

// AvailableContactsListResult - Response for the ListAvailableContacts API service call.
type AvailableContactsListResult struct {
	// A list of available contacts
	Value []*AvailableContacts `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// AvailableGroundStation - GroundStations available to schedule Contacts
type AvailableGroundStation struct {
	// Azure region
	Location *string `json:"location,omitempty"`

	// The properties bag for this resource
	Properties *AvailableGroundStationProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of groundStation
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the ground station.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AvailableGroundStationListResult - Response for the AvailableGroundStations API service call.
type AvailableGroundStationListResult struct {
	// A list of ground station resources.
	Value []*AvailableGroundStation `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// AvailableGroundStationProperties - Properties object for Available groundstation.
type AvailableGroundStationProperties struct {
	// Altitude of the ground station
	AltitudeMeters *float32 `json:"altitudeMeters,omitempty"`

	// City of ground station.
	City *string `json:"city,omitempty"`

	// Latitude of the ground station in decimal degrees.
	LatitudeDegrees *float32 `json:"latitudeDegrees,omitempty"`

	// Longitude of the ground station in decimal degrees.
	LongitudeDegrees *float32 `json:"longitudeDegrees,omitempty"`

	// Ground station provider name.
	ProviderName *string `json:"providerName,omitempty"`
}

// AvailableGroundStationsClientGetOptions contains the optional parameters for the AvailableGroundStationsClient.Get method.
type AvailableGroundStationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AvailableGroundStationsClientListByCapabilityOptions contains the optional parameters for the AvailableGroundStationsClient.ListByCapability
// method.
type AvailableGroundStationsClientListByCapabilityOptions struct {
	// placeholder for future optional parameters
}

// CloudError - An error response from the service.
type CloudError struct {
	// An error response from the service.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody - An error response from the service.
type CloudErrorBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`

	// A list of additional details about the error.
	Details []*CloudErrorBody `json:"details,omitempty"`

	// A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`

	// The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
}

// Contact - Customer creates a contact resource for a spacecraft resource.
type Contact struct {
	// Properties of the Contact Resource.
	Properties *ContactsProperties `json:"properties,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ContactInstanceProperties - Contact Instance Properties
type ContactInstanceProperties struct {
	// READ-ONLY; Azimuth of the antenna at the end of the contact in decimal degrees.
	EndAzimuthDegrees *float32 `json:"endAzimuthDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Spacecraft elevation above the horizon at contact end.
	EndElevationDegrees *float32 `json:"endElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Maximum elevation of the antenna during the contact in decimal degrees.
	MaximumElevationDegrees *float32 `json:"maximumElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Time to lost receiving a signal.
	RxEndTime *time.Time `json:"rxEndTime,omitempty" azure:"ro"`

	// READ-ONLY; Earliest time to receive a signal.
	RxStartTime *time.Time `json:"rxStartTime,omitempty" azure:"ro"`

	// READ-ONLY; Azimuth of the antenna at the start of the contact in decimal degrees.
	StartAzimuthDegrees *float32 `json:"startAzimuthDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Spacecraft elevation above the horizon at contact start.
	StartElevationDegrees *float32 `json:"startElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Time at which antenna transmit will be disabled.
	TxEndTime *time.Time `json:"txEndTime,omitempty" azure:"ro"`

	// READ-ONLY; Time at which antenna transmit will be enabled.
	TxStartTime *time.Time `json:"txStartTime,omitempty" azure:"ro"`
}

// ContactListResult - Response for the ListContacts API service call.
type ContactListResult struct {
	// A list of contact resources in a resource group.
	Value []*Contact `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ContactParameters - Parameters that define the contact resource.
type ContactParameters struct {
	// REQUIRED; The reference to the contact profile resource.
	ContactProfile *ResourceReference `json:"contactProfile,omitempty"`

	// REQUIRED; End time of a contact.
	EndTime *time.Time `json:"endTime,omitempty"`

	// REQUIRED; Name of Azure Ground Station.
	GroundStationName *string `json:"groundStationName,omitempty"`

	// REQUIRED; Start time of a contact.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// ContactProfile - Customer creates a Contact Profile Resource, which will contain all of the configurations required for
// scheduling a contact.
type ContactProfile struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Properties of the spacecraft resource.
	Properties *ContactProfilesProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ContactProfileLink - Contact Profile link
type ContactProfileLink struct {
	// REQUIRED; Contact Profile Link Channel
	Channels []*ContactProfileLinkChannel `json:"channels,omitempty"`

	// REQUIRED; Direction (uplink or downlink)
	Direction *Direction `json:"direction,omitempty"`

	// REQUIRED; polarization. eg (RHCP, LHCP)
	Polarization *Polarization `json:"polarization,omitempty"`

	// Effective Isotropic Radiated Power (EIRP) in dBW.
	EirpdBW *float32 `json:"eirpdBW,omitempty"`

	// Gain To Noise Temperature in db/K.
	GainOverTemperature *float32 `json:"gainOverTemperature,omitempty"`
}

// ContactProfileLinkChannel - Contact Profile Link Channel
type ContactProfileLinkChannel struct {
	// REQUIRED; Bandwidth in MHz
	BandwidthMHz *float32 `json:"bandwidthMHz,omitempty"`

	// REQUIRED; Center Frequency in MHz
	CenterFrequencyMHz *float32 `json:"centerFrequencyMHz,omitempty"`

	// REQUIRED; Customer End point to store/retrieve data during a contact.
	EndPoint *EndPoint `json:"endPoint,omitempty"`

	// Configuration for decoding
	DecodingConfiguration *string `json:"decodingConfiguration,omitempty"`

	// Configuration for demodulation
	DemodulationConfiguration *string `json:"demodulationConfiguration,omitempty"`

	// Configuration for encoding
	EncodingConfiguration *string `json:"encodingConfiguration,omitempty"`

	// Configuration for modulation
	ModulationConfiguration *string `json:"modulationConfiguration,omitempty"`
}

// ContactProfileListResult - Response for the ListContactProfiles API service call.
type ContactProfileListResult struct {
	// A list of contact profile resources in a resource group.
	Value []*ContactProfile `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ContactProfilesClientBeginCreateOrUpdateOptions contains the optional parameters for the ContactProfilesClient.BeginCreateOrUpdate
// method.
type ContactProfilesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ContactProfilesClientBeginDeleteOptions contains the optional parameters for the ContactProfilesClient.BeginDelete method.
type ContactProfilesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ContactProfilesClientGetOptions contains the optional parameters for the ContactProfilesClient.Get method.
type ContactProfilesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ContactProfilesClientListBySubscriptionOptions contains the optional parameters for the ContactProfilesClient.ListBySubscription
// method.
type ContactProfilesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ContactProfilesClientListOptions contains the optional parameters for the ContactProfilesClient.List method.
type ContactProfilesClientListOptions struct {
	// placeholder for future optional parameters
}

// ContactProfilesClientUpdateTagsOptions contains the optional parameters for the ContactProfilesClient.UpdateTags method.
type ContactProfilesClientUpdateTagsOptions struct {
	// placeholder for future optional parameters
}

// ContactProfilesProperties - List of Contact Profile Resource Properties.
type ContactProfilesProperties struct {
	// REQUIRED; Links of the Contact Profile
	Links []*ContactProfileLink `json:"links,omitempty"`

	// Auto track configuration.
	AutoTrackingConfiguration *AutoTrackingConfiguration `json:"autoTrackingConfiguration,omitempty"`

	// The URI of the Event Hub used for telemetry
	EventHubURI *string `json:"eventHubUri,omitempty"`

	// Minimum viable elevation for the contact in decimal degrees.
	MinimumElevationDegrees *float32 `json:"minimumElevationDegrees,omitempty"`

	// Minimum viable contact duration in ISO 8601 format.
	MinimumViableContactDuration *string `json:"minimumViableContactDuration,omitempty"`
}

// ContactsClientBeginCreateOptions contains the optional parameters for the ContactsClient.BeginCreate method.
type ContactsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ContactsClientBeginDeleteOptions contains the optional parameters for the ContactsClient.BeginDelete method.
type ContactsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ContactsClientGetOptions contains the optional parameters for the ContactsClient.Get method.
type ContactsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ContactsClientListOptions contains the optional parameters for the ContactsClient.List method.
type ContactsClientListOptions struct {
	// placeholder for future optional parameters
}

// ContactsProperties - Properties of the Contact Resource.
type ContactsProperties struct {
	// REQUIRED; The reference to the contact profile resource.
	ContactProfile *ResourceReference `json:"contactProfile,omitempty"`

	// REQUIRED; Azure Ground Station name.
	GroundStationName *string `json:"groundStationName,omitempty"`

	// REQUIRED; Reservation end time of a contact.
	ReservationEndTime *time.Time `json:"reservationEndTime,omitempty"`

	// REQUIRED; Reservation start time of a contact.
	ReservationStartTime *time.Time `json:"reservationStartTime,omitempty"`

	// READ-ONLY; Azimuth of the antenna at the end of the contact in decimal degrees.
	EndAzimuthDegrees *float32 `json:"endAzimuthDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Spacecraft elevation above the horizon at contact end.
	EndElevationDegrees *float32 `json:"endElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Any error message while scheduling a contact.
	ErrorMessage *string `json:"errorMessage,omitempty" azure:"ro"`

	// READ-ONLY; Maximum elevation of the antenna during the contact in decimal degrees.
	MaximumElevationDegrees *float32 `json:"maximumElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Receive end time of a contact.
	RxEndTime *time.Time `json:"rxEndTime,omitempty" azure:"ro"`

	// READ-ONLY; Receive start time of a contact.
	RxStartTime *time.Time `json:"rxStartTime,omitempty" azure:"ro"`

	// READ-ONLY; Azimuth of the antenna at the start of the contact in decimal degrees.
	StartAzimuthDegrees *float32 `json:"startAzimuthDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Spacecraft elevation above the horizon at contact start.
	StartElevationDegrees *float32 `json:"startElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Status of a contact.
	Status *Status `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; Transmit end time of a contact.
	TxEndTime *time.Time `json:"txEndTime,omitempty" azure:"ro"`

	// READ-ONLY; Transmit start time of a contact.
	TxStartTime *time.Time `json:"txStartTime,omitempty" azure:"ro"`
}

// EndPoint - Customer End point to store/retrieve data during a contact.
type EndPoint struct {
	// REQUIRED; Name of an end point.
	EndPointName *string `json:"endPointName,omitempty"`

	// REQUIRED; IP Address.
	IPAddress *string `json:"ipAddress,omitempty"`

	// REQUIRED; TCP port to listen on to receive data.
	Port *string `json:"port,omitempty"`

	// REQUIRED; Protocol either UDP or TCP.
	Protocol *Protocol `json:"protocol,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceIDListResult - Response for an API service call that lists the resource IDs of resources associated with another
// resource.
type ResourceIDListResult struct {
	// A list of Azure Resource IDs.
	Value []*ResourceIDListResultValueItem `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

type ResourceIDListResultValueItem struct {
	// The Azure Resource ID
	ID *string `json:"id,omitempty"`
}

// ResourceReference - Resource Reference
type ResourceReference struct {
	// Resource ID.
	ID *string `json:"id,omitempty"`
}

// Spacecraft - Customer creates a spacecraft resource to schedule a contact.
type Spacecraft struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Spacecraft Properties
	Properties *SpacecraftsProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SpacecraftLink - Spacecraft Link
type SpacecraftLink struct {
	// REQUIRED; Bandwidth in MHz
	BandwidthMHz *float32 `json:"bandwidthMHz,omitempty"`

	// REQUIRED; Center Frequency in MHz
	CenterFrequencyMHz *float32 `json:"centerFrequencyMHz,omitempty"`

	// REQUIRED; Direction (uplink or downlink)
	Direction *Direction `json:"direction,omitempty"`

	// REQUIRED; polarization. eg (RHCP, LHCP)
	Polarization *Polarization `json:"polarization,omitempty"`
}

// SpacecraftListResult - Response for the ListSpacecrafts API service call.
type SpacecraftListResult struct {
	// A list of spacecraft resources in a resource group.
	Value []*Spacecraft `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// SpacecraftsClientBeginCreateOrUpdateOptions contains the optional parameters for the SpacecraftsClient.BeginCreateOrUpdate
// method.
type SpacecraftsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SpacecraftsClientBeginDeleteOptions contains the optional parameters for the SpacecraftsClient.BeginDelete method.
type SpacecraftsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SpacecraftsClientBeginListAvailableContactsOptions contains the optional parameters for the SpacecraftsClient.BeginListAvailableContacts
// method.
type SpacecraftsClientBeginListAvailableContactsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SpacecraftsClientGetOptions contains the optional parameters for the SpacecraftsClient.Get method.
type SpacecraftsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SpacecraftsClientListBySubscriptionOptions contains the optional parameters for the SpacecraftsClient.ListBySubscription
// method.
type SpacecraftsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// SpacecraftsClientListOptions contains the optional parameters for the SpacecraftsClient.List method.
type SpacecraftsClientListOptions struct {
	// placeholder for future optional parameters
}

// SpacecraftsClientUpdateTagsOptions contains the optional parameters for the SpacecraftsClient.UpdateTags method.
type SpacecraftsClientUpdateTagsOptions struct {
	// placeholder for future optional parameters
}

// SpacecraftsProperties - List of Spacecraft Resource Properties.
type SpacecraftsProperties struct {
	// REQUIRED; NORAD ID of the spacecraft.
	NoradID *string `json:"noradId,omitempty"`

	// Links of the Spacecraft
	Links []*SpacecraftLink `json:"links,omitempty"`

	// Title line of Two Line Element (TLE).
	TitleLine *string `json:"titleLine,omitempty"`

	// Line 1 of Two Line Element (TLE).
	TleLine1 *string `json:"tleLine1,omitempty"`

	// Line 2 of Two Line Element (TLE).
	TleLine2 *string `json:"tleLine2,omitempty"`

	// READ-ONLY; Authorization status of spacecraft.
	AuthorizationStatus *AuthorizationStatus `json:"authorizationStatus,omitempty" azure:"ro"`

	// READ-ONLY; Details of the authorization status.
	AuthorizationStatusExtended *string `json:"authorizationStatusExtended,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TagsObject - Tags object for patch operations.
type TagsObject struct {
	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}
