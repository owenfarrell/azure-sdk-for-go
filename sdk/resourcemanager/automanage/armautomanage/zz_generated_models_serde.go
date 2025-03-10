//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AssignmentReportProperties.
func (a AssignmentReportProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "configurationProfile", a.ConfigurationProfile)
	populate(objectMap, "duration", a.Duration)
	populate(objectMap, "endTime", a.EndTime)
	populate(objectMap, "error", a.Error)
	populate(objectMap, "lastModifiedTime", a.LastModifiedTime)
	populate(objectMap, "reportFormatVersion", a.ReportFormatVersion)
	populate(objectMap, "resources", a.Resources)
	populate(objectMap, "startTime", a.StartTime)
	populate(objectMap, "status", a.Status)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BestPracticeList.
func (b BestPracticeList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", b.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationProfile.
func (c ConfigurationProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationProfileAssignmentList.
func (c ConfigurationProfileAssignmentList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationProfileAssignmentProperties.
func (c ConfigurationProfileAssignmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "configurationProfile", c.ConfigurationProfile)
	populate(objectMap, "profileOverrides", c.ProfileOverrides)
	populate(objectMap, "status", c.Status)
	populate(objectMap, "targetId", c.TargetID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationProfileList.
func (c ConfigurationProfileList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationProfileProperties.
func (c ConfigurationProfileProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "configuration", &c.Configuration)
	populate(objectMap, "overrides", c.Overrides)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationProfileUpdate.
func (c ConfigurationProfileUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ReportList.
func (r ReportList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServicePrincipalListResult.
func (s ServicePrincipalListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
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

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UpdateResource.
func (u UpdateResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", u.Tags)
	return json.Marshal(objectMap)
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

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
