//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// FlowLogsClient contains the methods for the FlowLogs group.
// Don't use this type directly, use NewFlowLogsClient() instead.
type FlowLogsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFlowLogsClient creates a new instance of FlowLogsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFlowLogsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FlowLogsClient, error) {
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
	client := &FlowLogsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a flow log for the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkWatcherName - The name of the network watcher.
// flowLogName - The name of the flow log.
// parameters - Parameters that define the create or update flow log resource.
// options - FlowLogsClientBeginCreateOrUpdateOptions contains the optional parameters for the FlowLogsClient.BeginCreateOrUpdate
// method.
func (client *FlowLogsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog, options *FlowLogsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[FlowLogsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkWatcherName, flowLogName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[FlowLogsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[FlowLogsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a flow log for the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FlowLogsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog, options *FlowLogsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkWatcherName, flowLogName, parameters, options)
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
func (client *FlowLogsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog, options *FlowLogsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if flowLogName == "" {
		return nil, errors.New("parameter flowLogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified flow log resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkWatcherName - The name of the network watcher.
// flowLogName - The name of the flow log resource.
// options - FlowLogsClientBeginDeleteOptions contains the optional parameters for the FlowLogsClient.BeginDelete method.
func (client *FlowLogsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsClientBeginDeleteOptions) (*armruntime.Poller[FlowLogsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkWatcherName, flowLogName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[FlowLogsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[FlowLogsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified flow log resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FlowLogsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkWatcherName, flowLogName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FlowLogsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if flowLogName == "" {
		return nil, errors.New("parameter flowLogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a flow log resource by name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkWatcherName - The name of the network watcher.
// flowLogName - The name of the flow log resource.
// options - FlowLogsClientGetOptions contains the optional parameters for the FlowLogsClient.Get method.
func (client *FlowLogsClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsClientGetOptions) (FlowLogsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkWatcherName, flowLogName, options)
	if err != nil {
		return FlowLogsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FlowLogsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FlowLogsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FlowLogsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if flowLogName == "" {
		return nil, errors.New("parameter flowLogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FlowLogsClient) getHandleResponse(resp *http.Response) (FlowLogsClientGetResponse, error) {
	result := FlowLogsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FlowLog); err != nil {
		return FlowLogsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all flow log resources for the specified Network Watcher.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing Network Watcher.
// networkWatcherName - The name of the Network Watcher resource.
// options - FlowLogsClientListOptions contains the optional parameters for the FlowLogsClient.List method.
func (client *FlowLogsClient) NewListPager(resourceGroupName string, networkWatcherName string, options *FlowLogsClientListOptions) *runtime.Pager[FlowLogsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[FlowLogsClientListResponse]{
		More: func(page FlowLogsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FlowLogsClientListResponse) (FlowLogsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkWatcherName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FlowLogsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FlowLogsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FlowLogsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FlowLogsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, options *FlowLogsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FlowLogsClient) listHandleResponse(resp *http.Response) (FlowLogsClientListResponse, error) {
	result := FlowLogsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FlowLogListResult); err != nil {
		return FlowLogsClientListResponse{}, err
	}
	return result, nil
}

// UpdateTags - Update tags of the specified flow log.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkWatcherName - The name of the network watcher.
// flowLogName - The name of the flow log.
// parameters - Parameters supplied to update flow log tags.
// options - FlowLogsClientUpdateTagsOptions contains the optional parameters for the FlowLogsClient.UpdateTags method.
func (client *FlowLogsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters TagsObject, options *FlowLogsClientUpdateTagsOptions) (FlowLogsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkWatcherName, flowLogName, parameters, options)
	if err != nil {
		return FlowLogsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FlowLogsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FlowLogsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *FlowLogsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters TagsObject, options *FlowLogsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if flowLogName == "" {
		return nil, errors.New("parameter flowLogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *FlowLogsClient) updateTagsHandleResponse(resp *http.Response) (FlowLogsClientUpdateTagsResponse, error) {
	result := FlowLogsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FlowLog); err != nil {
		return FlowLogsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
