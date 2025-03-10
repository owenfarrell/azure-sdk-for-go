//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceprovisioningservices

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CertificateListDescription.
func (c CertificateListDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CertificateProperties.
func (c CertificateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateByteArray(objectMap, "certificate", c.Certificate, runtime.Base64StdFormat)
	populateTimeRFC1123(objectMap, "created", c.Created)
	populateTimeRFC1123(objectMap, "expiry", c.Expiry)
	populate(objectMap, "isVerified", c.IsVerified)
	populate(objectMap, "subject", c.Subject)
	populate(objectMap, "thumbprint", c.Thumbprint)
	populateTimeRFC1123(objectMap, "updated", c.Updated)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CertificateProperties.
func (c *CertificateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "certificate":
			err = runtime.DecodeByteArray(string(val), &c.Certificate, runtime.Base64StdFormat)
			delete(rawMsg, key)
		case "created":
			err = unpopulateTimeRFC1123(val, &c.Created)
			delete(rawMsg, key)
		case "expiry":
			err = unpopulateTimeRFC1123(val, &c.Expiry)
			delete(rawMsg, key)
		case "isVerified":
			err = unpopulate(val, &c.IsVerified)
			delete(rawMsg, key)
		case "subject":
			err = unpopulate(val, &c.Subject)
			delete(rawMsg, key)
		case "thumbprint":
			err = unpopulate(val, &c.Thumbprint)
			delete(rawMsg, key)
		case "updated":
			err = unpopulateTimeRFC1123(val, &c.Updated)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GroupIDInformationProperties.
func (g GroupIDInformationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", g.GroupID)
	populate(objectMap, "requiredMembers", g.RequiredMembers)
	populate(objectMap, "requiredZoneNames", g.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IotDpsPropertiesDescription.
func (i IotDpsPropertiesDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allocationPolicy", i.AllocationPolicy)
	populate(objectMap, "authorizationPolicies", i.AuthorizationPolicies)
	populate(objectMap, "deviceProvisioningHostName", i.DeviceProvisioningHostName)
	populate(objectMap, "enableDataResidency", i.EnableDataResidency)
	populate(objectMap, "idScope", i.IDScope)
	populate(objectMap, "ipFilterRules", i.IPFilterRules)
	populate(objectMap, "iotHubs", i.IotHubs)
	populate(objectMap, "privateEndpointConnections", i.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", i.PublicNetworkAccess)
	populate(objectMap, "serviceOperationsHostName", i.ServiceOperationsHostName)
	populate(objectMap, "state", i.State)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IotDpsSKUDefinitionListResult.
func (i IotDpsSKUDefinitionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", i.NextLink)
	populate(objectMap, "value", i.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResources.
func (p PrivateLinkResources) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProvisioningServiceDescription.
func (p ProvisioningServiceDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", p.Etag)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "sku", p.SKU)
	populate(objectMap, "systemData", p.SystemData)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProvisioningServiceDescriptionListResult.
func (p ProvisioningServiceDescriptionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SharedAccessSignatureAuthorizationRuleListResult.
func (s SharedAccessSignatureAuthorizationRuleListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TagsResource.
func (t TagsResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VerificationCodeResponseProperties.
func (v VerificationCodeResponseProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateByteArray(objectMap, "certificate", v.Certificate, runtime.Base64StdFormat)
	populate(objectMap, "created", v.Created)
	populate(objectMap, "expiry", v.Expiry)
	populate(objectMap, "isVerified", v.IsVerified)
	populate(objectMap, "subject", v.Subject)
	populate(objectMap, "thumbprint", v.Thumbprint)
	populate(objectMap, "updated", v.Updated)
	populate(objectMap, "verificationCode", v.VerificationCode)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VerificationCodeResponseProperties.
func (v *VerificationCodeResponseProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "certificate":
			err = runtime.DecodeByteArray(string(val), &v.Certificate, runtime.Base64StdFormat)
			delete(rawMsg, key)
		case "created":
			err = unpopulate(val, &v.Created)
			delete(rawMsg, key)
		case "expiry":
			err = unpopulate(val, &v.Expiry)
			delete(rawMsg, key)
		case "isVerified":
			err = unpopulate(val, &v.IsVerified)
			delete(rawMsg, key)
		case "subject":
			err = unpopulate(val, &v.Subject)
			delete(rawMsg, key)
		case "thumbprint":
			err = unpopulate(val, &v.Thumbprint)
			delete(rawMsg, key)
		case "updated":
			err = unpopulate(val, &v.Updated)
			delete(rawMsg, key)
		case "verificationCode":
			err = unpopulate(val, &v.VerificationCode)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateByteArray(m map[string]interface{}, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
