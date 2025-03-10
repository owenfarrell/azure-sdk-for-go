//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

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

// TargetTypesClient contains the methods for the TargetTypes group.
// Don't use this type directly, use NewTargetTypesClient() instead.
type TargetTypesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTargetTypesClient creates a new instance of TargetTypesClient with the specified values.
// subscriptionID - GUID that represents an Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTargetTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TargetTypesClient, error) {
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
	client := &TargetTypesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get a Target Type resources for given location.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - String that represents a Location resource name.
// targetTypeName - String that represents a Target Type resource name.
// options - TargetTypesClientGetOptions contains the optional parameters for the TargetTypesClient.Get method.
func (client *TargetTypesClient) Get(ctx context.Context, locationName string, targetTypeName string, options *TargetTypesClientGetOptions) (TargetTypesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, targetTypeName, options)
	if err != nil {
		return TargetTypesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TargetTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TargetTypesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TargetTypesClient) getCreateRequest(ctx context.Context, locationName string, targetTypeName string, options *TargetTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/locations/{locationName}/targetTypes/{targetTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if targetTypeName == "" {
		return nil, errors.New("parameter targetTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetTypeName}", url.PathEscape(targetTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TargetTypesClient) getHandleResponse(resp *http.Response) (TargetTypesClientGetResponse, error) {
	result := TargetTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TargetType); err != nil {
		return TargetTypesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of Target Type resources for given location.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - String that represents a Location resource name.
// options - TargetTypesClientListOptions contains the optional parameters for the TargetTypesClient.List method.
func (client *TargetTypesClient) NewListPager(locationName string, options *TargetTypesClientListOptions) *runtime.Pager[TargetTypesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[TargetTypesClientListResponse]{
		More: func(page TargetTypesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TargetTypesClientListResponse) (TargetTypesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, locationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TargetTypesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TargetTypesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TargetTypesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TargetTypesClient) listCreateRequest(ctx context.Context, locationName string, options *TargetTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/locations/{locationName}/targetTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TargetTypesClient) listHandleResponse(resp *http.Response) (TargetTypesClientListResponse, error) {
	result := TargetTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TargetTypeListResult); err != nil {
		return TargetTypesClientListResponse{}, err
	}
	return result, nil
}
