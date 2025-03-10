//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// KeysClient contains the methods for the Keys group.
// Don't use this type directly, use NewKeysClient() instead.
type KeysClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewKeysClient creates a new instance of KeysClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KeysClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &KeysClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// ListByAutomationAccount - Retrieve the automation keys for an account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - KeysClientListByAutomationAccountOptions contains the optional parameters for the KeysClient.ListByAutomationAccount
// method.
func (client *KeysClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, options *KeysClientListByAutomationAccountOptions) (KeysClientListByAutomationAccountResponse, error) {
	req, err := client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return KeysClientListByAutomationAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KeysClientListByAutomationAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KeysClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByAutomationAccountHandleResponse(resp)
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *KeysClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *KeysClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/listKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *KeysClient) listByAutomationAccountHandleResponse(resp *http.Response) (KeysClientListByAutomationAccountResponse, error) {
	result := KeysClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KeyListResult); err != nil {
		return KeysClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}
