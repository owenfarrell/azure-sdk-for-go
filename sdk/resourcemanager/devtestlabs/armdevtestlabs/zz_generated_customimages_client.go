//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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

// CustomImagesClient contains the methods for the CustomImages group.
// Don't use this type directly, use NewCustomImagesClient() instead.
type CustomImagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCustomImagesClient creates a new instance of CustomImagesClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCustomImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CustomImagesClient, error) {
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
	client := &CustomImagesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or replace an existing custom image. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the custom image.
// customImage - A custom image.
// options - CustomImagesClientBeginCreateOrUpdateOptions contains the optional parameters for the CustomImagesClient.BeginCreateOrUpdate
// method.
func (client *CustomImagesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, customImage CustomImage, options *CustomImagesClientBeginCreateOrUpdateOptions) (*armruntime.Poller[CustomImagesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, labName, name, customImage, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[CustomImagesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[CustomImagesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or replace an existing custom image. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CustomImagesClient) createOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, customImage CustomImage, options *CustomImagesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, name, customImage, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CustomImagesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, customImage CustomImage, options *CustomImagesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, customImage)
}

// BeginDelete - Delete custom image. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the custom image.
// options - CustomImagesClientBeginDeleteOptions contains the optional parameters for the CustomImagesClient.BeginDelete
// method.
func (client *CustomImagesClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, name string, options *CustomImagesClientBeginDeleteOptions) (*armruntime.Poller[CustomImagesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, labName, name, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[CustomImagesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[CustomImagesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete custom image. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CustomImagesClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, name string, options *CustomImagesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomImagesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *CustomImagesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get custom image.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the custom image.
// options - CustomImagesClientGetOptions contains the optional parameters for the CustomImagesClient.Get method.
func (client *CustomImagesClient) Get(ctx context.Context, resourceGroupName string, labName string, name string, options *CustomImagesClientGetOptions) (CustomImagesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return CustomImagesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomImagesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomImagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *CustomImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomImagesClient) getHandleResponse(resp *http.Response) (CustomImagesClientGetResponse, error) {
	result := CustomImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomImage); err != nil {
		return CustomImagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List custom images in a given lab.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// options - CustomImagesClientListOptions contains the optional parameters for the CustomImagesClient.List method.
func (client *CustomImagesClient) NewListPager(resourceGroupName string, labName string, options *CustomImagesClientListOptions) *runtime.Pager[CustomImagesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[CustomImagesClientListResponse]{
		More: func(page CustomImagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomImagesClientListResponse) (CustomImagesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, labName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CustomImagesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CustomImagesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CustomImagesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CustomImagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *CustomImagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CustomImagesClient) listHandleResponse(resp *http.Response) (CustomImagesClientListResponse, error) {
	result := CustomImagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomImageList); err != nil {
		return CustomImagesClientListResponse{}, err
	}
	return result, nil
}

// Update - Allows modifying tags of custom images. All other properties will be ignored.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the custom image.
// customImage - A custom image.
// options - CustomImagesClientUpdateOptions contains the optional parameters for the CustomImagesClient.Update method.
func (client *CustomImagesClient) Update(ctx context.Context, resourceGroupName string, labName string, name string, customImage CustomImageFragment, options *CustomImagesClientUpdateOptions) (CustomImagesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, name, customImage, options)
	if err != nil {
		return CustomImagesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomImagesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomImagesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CustomImagesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, customImage CustomImageFragment, options *CustomImagesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, customImage)
}

// updateHandleResponse handles the Update response.
func (client *CustomImagesClient) updateHandleResponse(resp *http.Response) (CustomImagesClientUpdateResponse, error) {
	result := CustomImagesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomImage); err != nil {
		return CustomImagesClientUpdateResponse{}, err
	}
	return result, nil
}
