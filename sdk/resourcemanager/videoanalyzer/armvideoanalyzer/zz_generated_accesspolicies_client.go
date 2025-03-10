//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer

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
	"strconv"
	"strings"
)

// AccessPoliciesClient contains the methods for the AccessPolicies group.
// Don't use this type directly, use NewAccessPoliciesClient() instead.
type AccessPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccessPoliciesClient creates a new instance of AccessPoliciesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccessPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessPoliciesClient, error) {
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
	client := &AccessPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new access policy resource or updates an existing one with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// accessPolicyName - The Access Policy name.
// parameters - The request parameters
// options - AccessPoliciesClientCreateOrUpdateOptions contains the optional parameters for the AccessPoliciesClient.CreateOrUpdate
// method.
func (client *AccessPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, accessPolicyName string, parameters AccessPolicyEntity, options *AccessPoliciesClientCreateOrUpdateOptions) (AccessPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, accessPolicyName, parameters, options)
	if err != nil {
		return AccessPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AccessPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AccessPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, accessPolicyName string, parameters AccessPolicyEntity, options *AccessPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/accessPolicies/{accessPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AccessPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (AccessPoliciesClientCreateOrUpdateResponse, error) {
	result := AccessPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessPolicyEntity); err != nil {
		return AccessPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing access policy resource with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// accessPolicyName - The Access Policy name.
// options - AccessPoliciesClientDeleteOptions contains the optional parameters for the AccessPoliciesClient.Delete method.
func (client *AccessPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, accessPolicyName string, options *AccessPoliciesClientDeleteOptions) (AccessPoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, accessPolicyName, options)
	if err != nil {
		return AccessPoliciesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AccessPoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AccessPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccessPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, accessPolicyName string, options *AccessPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/accessPolicies/{accessPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves an existing access policy resource with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// accessPolicyName - The Access Policy name.
// options - AccessPoliciesClientGetOptions contains the optional parameters for the AccessPoliciesClient.Get method.
func (client *AccessPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, accessPolicyName string, options *AccessPoliciesClientGetOptions) (AccessPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, accessPolicyName, options)
	if err != nil {
		return AccessPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccessPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, accessPolicyName string, options *AccessPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/accessPolicies/{accessPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccessPoliciesClient) getHandleResponse(resp *http.Response) (AccessPoliciesClientGetResponse, error) {
	result := AccessPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessPolicyEntity); err != nil {
		return AccessPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves all existing access policy resources, along with their JSON representations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// options - AccessPoliciesClientListOptions contains the optional parameters for the AccessPoliciesClient.List method.
func (client *AccessPoliciesClient) NewListPager(resourceGroupName string, accountName string, options *AccessPoliciesClientListOptions) *runtime.Pager[AccessPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AccessPoliciesClientListResponse]{
		More: func(page AccessPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessPoliciesClientListResponse) (AccessPoliciesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccessPoliciesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccessPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AccessPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccessPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/accessPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessPoliciesClient) listHandleResponse(resp *http.Response) (AccessPoliciesClientListResponse, error) {
	result := AccessPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessPolicyEntityCollection); err != nil {
		return AccessPoliciesClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates individual properties of an existing access policy resource with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// accessPolicyName - The Access Policy name.
// parameters - The request parameters
// options - AccessPoliciesClientUpdateOptions contains the optional parameters for the AccessPoliciesClient.Update method.
func (client *AccessPoliciesClient) Update(ctx context.Context, resourceGroupName string, accountName string, accessPolicyName string, parameters AccessPolicyEntity, options *AccessPoliciesClientUpdateOptions) (AccessPoliciesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, accessPolicyName, parameters, options)
	if err != nil {
		return AccessPoliciesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessPoliciesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessPoliciesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AccessPoliciesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, accessPolicyName string, parameters AccessPolicyEntity, options *AccessPoliciesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/accessPolicies/{accessPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *AccessPoliciesClient) updateHandleResponse(resp *http.Response) (AccessPoliciesClientUpdateResponse, error) {
	result := AccessPoliciesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessPolicyEntity); err != nil {
		return AccessPoliciesClientUpdateResponse{}, err
	}
	return result, nil
}
