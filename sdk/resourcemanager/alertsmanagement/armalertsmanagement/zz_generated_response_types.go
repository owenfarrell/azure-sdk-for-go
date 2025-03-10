//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

// AlertProcessingRulesClientCreateOrUpdateResponse contains the response from method AlertProcessingRulesClient.CreateOrUpdate.
type AlertProcessingRulesClientCreateOrUpdateResponse struct {
	AlertProcessingRule
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertProcessingRulesClientDeleteResponse contains the response from method AlertProcessingRulesClient.Delete.
type AlertProcessingRulesClientDeleteResponse struct {
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertProcessingRulesClientGetByNameResponse contains the response from method AlertProcessingRulesClient.GetByName.
type AlertProcessingRulesClientGetByNameResponse struct {
	AlertProcessingRule
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertProcessingRulesClientListByResourceGroupResponse contains the response from method AlertProcessingRulesClient.ListByResourceGroup.
type AlertProcessingRulesClientListByResourceGroupResponse struct {
	AlertProcessingRulesList
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertProcessingRulesClientListBySubscriptionResponse contains the response from method AlertProcessingRulesClient.ListBySubscription.
type AlertProcessingRulesClientListBySubscriptionResponse struct {
	AlertProcessingRulesList
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertProcessingRulesClientUpdateResponse contains the response from method AlertProcessingRulesClient.Update.
type AlertProcessingRulesClientUpdateResponse struct {
	AlertProcessingRule
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertsClientChangeStateResponse contains the response from method AlertsClient.ChangeState.
type AlertsClientChangeStateResponse struct {
	Alert
}

// AlertsClientGetAllResponse contains the response from method AlertsClient.GetAll.
type AlertsClientGetAllResponse struct {
	AlertsList
}

// AlertsClientGetByIDResponse contains the response from method AlertsClient.GetByID.
type AlertsClientGetByIDResponse struct {
	Alert
}

// AlertsClientGetHistoryResponse contains the response from method AlertsClient.GetHistory.
type AlertsClientGetHistoryResponse struct {
	AlertModification
}

// AlertsClientGetSummaryResponse contains the response from method AlertsClient.GetSummary.
type AlertsClientGetSummaryResponse struct {
	AlertsSummary
}

// AlertsClientMetaDataResponse contains the response from method AlertsClient.MetaData.
type AlertsClientMetaDataResponse struct {
	AlertsMetaData
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsList
}

// SmartGroupsClientChangeStateResponse contains the response from method SmartGroupsClient.ChangeState.
type SmartGroupsClientChangeStateResponse struct {
	SmartGroup
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// SmartGroupsClientGetAllResponse contains the response from method SmartGroupsClient.GetAll.
type SmartGroupsClientGetAllResponse struct {
	SmartGroupsList
}

// SmartGroupsClientGetByIDResponse contains the response from method SmartGroupsClient.GetByID.
type SmartGroupsClientGetByIDResponse struct {
	SmartGroup
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// SmartGroupsClientGetHistoryResponse contains the response from method SmartGroupsClient.GetHistory.
type SmartGroupsClientGetHistoryResponse struct {
	SmartGroupModification
}
