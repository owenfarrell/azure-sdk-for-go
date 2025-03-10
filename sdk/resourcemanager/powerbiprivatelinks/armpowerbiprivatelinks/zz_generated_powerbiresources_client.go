//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

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

// PowerBIResourcesClient contains the methods for the PowerBIResources group.
// Don't use this type directly, use NewPowerBIResourcesClient() instead.
type PowerBIResourcesClient struct {
	host              string
	subscriptionID    string
	resourceGroupName string
	azureResourceName string
	pl                runtime.Pipeline
}

// NewPowerBIResourcesClient creates a new instance of PowerBIResourcesClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// resourceGroupName - The name of the resource group.
// azureResourceName - The name of the Azure resource.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPowerBIResourcesClient(subscriptionID string, resourceGroupName string, azureResourceName string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PowerBIResourcesClient, error) {
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
	client := &PowerBIResourcesClient{
		subscriptionID:    subscriptionID,
		resourceGroupName: resourceGroupName,
		azureResourceName: azureResourceName,
		host:              ep,
		pl:                pl,
	}
	return client, nil
}

// Create - Creates or updates a Private Link Service Resource for Power BI.
// If the operation fails it returns an *azcore.ResponseError type.
// body - Tenant resource to be created or updated.
// options - PowerBIResourcesClientCreateOptions contains the optional parameters for the PowerBIResourcesClient.Create method.
func (client *PowerBIResourcesClient) Create(ctx context.Context, body TenantResource, options *PowerBIResourcesClientCreateOptions) (PowerBIResourcesClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, body, options)
	if err != nil {
		return PowerBIResourcesClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PowerBIResourcesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PowerBIResourcesClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *PowerBIResourcesClient) createCreateRequest(ctx context.Context, body TenantResource, options *PowerBIResourcesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{azureResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.azureResourceName == "" {
		return nil, errors.New("parameter client.azureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureResourceName}", url.PathEscape(client.azureResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientTenantID != nil {
		req.Raw().Header.Set("x-ms-client-tenant-id", *options.ClientTenantID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createHandleResponse handles the Create response.
func (client *PowerBIResourcesClient) createHandleResponse(resp *http.Response) (PowerBIResourcesClientCreateResponse, error) {
	result := PowerBIResourcesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantResource); err != nil {
		return PowerBIResourcesClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Private Link Service Resource for Power BI.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PowerBIResourcesClientDeleteOptions contains the optional parameters for the PowerBIResourcesClient.Delete method.
func (client *PowerBIResourcesClient) Delete(ctx context.Context, options *PowerBIResourcesClientDeleteOptions) (PowerBIResourcesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, options)
	if err != nil {
		return PowerBIResourcesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PowerBIResourcesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PowerBIResourcesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return PowerBIResourcesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PowerBIResourcesClient) deleteCreateRequest(ctx context.Context, options *PowerBIResourcesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{azureResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.azureResourceName == "" {
		return nil, errors.New("parameter client.azureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureResourceName}", url.PathEscape(client.azureResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceName - Gets all the private link resources for the given Azure resource.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PowerBIResourcesClientListByResourceNameOptions contains the optional parameters for the PowerBIResourcesClient.ListByResourceName
// method.
func (client *PowerBIResourcesClient) ListByResourceName(ctx context.Context, options *PowerBIResourcesClientListByResourceNameOptions) (PowerBIResourcesClientListByResourceNameResponse, error) {
	req, err := client.listByResourceNameCreateRequest(ctx, options)
	if err != nil {
		return PowerBIResourcesClientListByResourceNameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PowerBIResourcesClientListByResourceNameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PowerBIResourcesClientListByResourceNameResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByResourceNameHandleResponse(resp)
}

// listByResourceNameCreateRequest creates the ListByResourceName request.
func (client *PowerBIResourcesClient) listByResourceNameCreateRequest(ctx context.Context, options *PowerBIResourcesClientListByResourceNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{azureResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.azureResourceName == "" {
		return nil, errors.New("parameter client.azureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureResourceName}", url.PathEscape(client.azureResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceNameHandleResponse handles the ListByResourceName response.
func (client *PowerBIResourcesClient) listByResourceNameHandleResponse(resp *http.Response) (PowerBIResourcesClientListByResourceNameResponse, error) {
	result := PowerBIResourcesClientListByResourceNameResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantResourceArray); err != nil {
		return PowerBIResourcesClientListByResourceNameResponse{}, err
	}
	return result, nil
}

// Update - Creates or updates a Private Link Service Resource for Power BI.
// If the operation fails it returns an *azcore.ResponseError type.
// body - Tenant resource to be created or updated.
// options - PowerBIResourcesClientUpdateOptions contains the optional parameters for the PowerBIResourcesClient.Update method.
func (client *PowerBIResourcesClient) Update(ctx context.Context, body TenantResource, options *PowerBIResourcesClientUpdateOptions) (PowerBIResourcesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, body, options)
	if err != nil {
		return PowerBIResourcesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PowerBIResourcesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PowerBIResourcesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PowerBIResourcesClient) updateCreateRequest(ctx context.Context, body TenantResource, options *PowerBIResourcesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{azureResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.azureResourceName == "" {
		return nil, errors.New("parameter client.azureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureResourceName}", url.PathEscape(client.azureResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientTenantID != nil {
		req.Raw().Header.Set("x-ms-client-tenant-id", *options.ClientTenantID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *PowerBIResourcesClient) updateHandleResponse(resp *http.Response) (PowerBIResourcesClientUpdateResponse, error) {
	result := PowerBIResourcesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantResource); err != nil {
		return PowerBIResourcesClientUpdateResponse{}, err
	}
	return result, nil
}
