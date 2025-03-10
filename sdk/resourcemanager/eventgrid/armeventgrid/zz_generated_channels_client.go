//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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

// ChannelsClient contains the methods for the Channels group.
// Don't use this type directly, use NewChannelsClient() instead.
type ChannelsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewChannelsClient creates a new instance of ChannelsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewChannelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ChannelsClient, error) {
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
	client := &ChannelsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Synchronously creates or updates a new channel with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the partners subscription.
// partnerNamespaceName - Name of the partner namespace.
// channelName - Name of the channel.
// channelInfo - Channel information.
// options - ChannelsClientCreateOrUpdateOptions contains the optional parameters for the ChannelsClient.CreateOrUpdate method.
func (client *ChannelsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, channelInfo Channel, options *ChannelsClientCreateOrUpdateOptions) (ChannelsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, partnerNamespaceName, channelName, channelInfo, options)
	if err != nil {
		return ChannelsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ChannelsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ChannelsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ChannelsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, channelInfo Channel, options *ChannelsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/channels/{channelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(channelName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, channelInfo)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ChannelsClient) createOrUpdateHandleResponse(resp *http.Response) (ChannelsClientCreateOrUpdateResponse, error) {
	result := ChannelsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Channel); err != nil {
		return ChannelsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete an existing channel.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the partners subscription.
// partnerNamespaceName - Name of the partner namespace.
// channelName - Name of the channel.
// options - ChannelsClientBeginDeleteOptions contains the optional parameters for the ChannelsClient.BeginDelete method.
func (client *ChannelsClient) BeginDelete(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, options *ChannelsClientBeginDeleteOptions) (*armruntime.Poller[ChannelsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, partnerNamespaceName, channelName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ChannelsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ChannelsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete an existing channel.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ChannelsClient) deleteOperation(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, options *ChannelsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, partnerNamespaceName, channelName, options)
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
func (client *ChannelsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, options *ChannelsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/channels/{channelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(channelName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get properties of a channel.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the partners subscription.
// partnerNamespaceName - Name of the partner namespace.
// channelName - Name of the channel.
// options - ChannelsClientGetOptions contains the optional parameters for the ChannelsClient.Get method.
func (client *ChannelsClient) Get(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, options *ChannelsClientGetOptions) (ChannelsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, partnerNamespaceName, channelName, options)
	if err != nil {
		return ChannelsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ChannelsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ChannelsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ChannelsClient) getCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, options *ChannelsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/channels/{channelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(channelName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ChannelsClient) getHandleResponse(resp *http.Response) (ChannelsClientGetResponse, error) {
	result := ChannelsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Channel); err != nil {
		return ChannelsClientGetResponse{}, err
	}
	return result, nil
}

// GetFullURL - Get the full endpoint URL of a partner destination channel.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the partners subscription.
// partnerNamespaceName - Name of the partner namespace.
// channelName - Name of the Channel.
// options - ChannelsClientGetFullURLOptions contains the optional parameters for the ChannelsClient.GetFullURL method.
func (client *ChannelsClient) GetFullURL(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, options *ChannelsClientGetFullURLOptions) (ChannelsClientGetFullURLResponse, error) {
	req, err := client.getFullURLCreateRequest(ctx, resourceGroupName, partnerNamespaceName, channelName, options)
	if err != nil {
		return ChannelsClientGetFullURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ChannelsClientGetFullURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ChannelsClientGetFullURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.getFullURLHandleResponse(resp)
}

// getFullURLCreateRequest creates the GetFullURL request.
func (client *ChannelsClient) getFullURLCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, options *ChannelsClientGetFullURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/channels/{channelName}/getFullUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(channelName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getFullURLHandleResponse handles the GetFullURL response.
func (client *ChannelsClient) getFullURLHandleResponse(resp *http.Response) (ChannelsClientGetFullURLResponse, error) {
	result := ChannelsClientGetFullURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionFullURL); err != nil {
		return ChannelsClientGetFullURLResponse{}, err
	}
	return result, nil
}

// NewListByPartnerNamespacePager - List all the channels in a partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the partners subscription.
// partnerNamespaceName - Name of the partner namespace.
// options - ChannelsClientListByPartnerNamespaceOptions contains the optional parameters for the ChannelsClient.ListByPartnerNamespace
// method.
func (client *ChannelsClient) NewListByPartnerNamespacePager(resourceGroupName string, partnerNamespaceName string, options *ChannelsClientListByPartnerNamespaceOptions) *runtime.Pager[ChannelsClientListByPartnerNamespaceResponse] {
	return runtime.NewPager(runtime.PageProcessor[ChannelsClientListByPartnerNamespaceResponse]{
		More: func(page ChannelsClientListByPartnerNamespaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ChannelsClientListByPartnerNamespaceResponse) (ChannelsClientListByPartnerNamespaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPartnerNamespaceCreateRequest(ctx, resourceGroupName, partnerNamespaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ChannelsClientListByPartnerNamespaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ChannelsClientListByPartnerNamespaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ChannelsClientListByPartnerNamespaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPartnerNamespaceHandleResponse(resp)
		},
	})
}

// listByPartnerNamespaceCreateRequest creates the ListByPartnerNamespace request.
func (client *ChannelsClient) listByPartnerNamespaceCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *ChannelsClientListByPartnerNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/channels"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByPartnerNamespaceHandleResponse handles the ListByPartnerNamespace response.
func (client *ChannelsClient) listByPartnerNamespaceHandleResponse(resp *http.Response) (ChannelsClientListByPartnerNamespaceResponse, error) {
	result := ChannelsClientListByPartnerNamespaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChannelsListResult); err != nil {
		return ChannelsClientListByPartnerNamespaceResponse{}, err
	}
	return result, nil
}

// Update - Synchronously updates a channel with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the partners subscription.
// partnerNamespaceName - Name of the partner namespace.
// channelName - Name of the channel.
// channelUpdateParameters - Channel update information.
// options - ChannelsClientUpdateOptions contains the optional parameters for the ChannelsClient.Update method.
func (client *ChannelsClient) Update(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, channelUpdateParameters ChannelUpdateParameters, options *ChannelsClientUpdateOptions) (ChannelsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, partnerNamespaceName, channelName, channelUpdateParameters, options)
	if err != nil {
		return ChannelsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ChannelsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ChannelsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return ChannelsClientUpdateResponse{}, nil
}

// updateCreateRequest creates the Update request.
func (client *ChannelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, channelName string, channelUpdateParameters ChannelUpdateParameters, options *ChannelsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/channels/{channelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(channelName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, channelUpdateParameters)
}
