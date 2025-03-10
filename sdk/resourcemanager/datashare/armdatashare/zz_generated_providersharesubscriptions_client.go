//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare

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

// ProviderShareSubscriptionsClient contains the methods for the ProviderShareSubscriptions group.
// Don't use this type directly, use NewProviderShareSubscriptionsClient() instead.
type ProviderShareSubscriptionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProviderShareSubscriptionsClient creates a new instance of ProviderShareSubscriptionsClient with the specified values.
// subscriptionID - The subscription identifier
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProviderShareSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProviderShareSubscriptionsClient, error) {
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
	client := &ProviderShareSubscriptionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Adjust - Adjust a share subscription's expiration date in a provider share
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - The name of the share account.
// shareName - The name of the share.
// providerShareSubscriptionID - To locate shareSubscription
// providerShareSubscription - The provider share subscription
// options - ProviderShareSubscriptionsClientAdjustOptions contains the optional parameters for the ProviderShareSubscriptionsClient.Adjust
// method.
func (client *ProviderShareSubscriptionsClient) Adjust(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, providerShareSubscription ProviderShareSubscription, options *ProviderShareSubscriptionsClientAdjustOptions) (ProviderShareSubscriptionsClientAdjustResponse, error) {
	req, err := client.adjustCreateRequest(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID, providerShareSubscription, options)
	if err != nil {
		return ProviderShareSubscriptionsClientAdjustResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderShareSubscriptionsClientAdjustResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderShareSubscriptionsClientAdjustResponse{}, runtime.NewResponseError(resp)
	}
	return client.adjustHandleResponse(resp)
}

// adjustCreateRequest creates the Adjust request.
func (client *ProviderShareSubscriptionsClient) adjustCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, providerShareSubscription ProviderShareSubscription, options *ProviderShareSubscriptionsClientAdjustOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}/adjust"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if providerShareSubscriptionID == "" {
		return nil, errors.New("parameter providerShareSubscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerShareSubscriptionId}", url.PathEscape(providerShareSubscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, providerShareSubscription)
}

// adjustHandleResponse handles the Adjust response.
func (client *ProviderShareSubscriptionsClient) adjustHandleResponse(resp *http.Response) (ProviderShareSubscriptionsClientAdjustResponse, error) {
	result := ProviderShareSubscriptionsClientAdjustResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderShareSubscription); err != nil {
		return ProviderShareSubscriptionsClientAdjustResponse{}, err
	}
	return result, nil
}

// GetByShare - Get share subscription in a provider share
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - The name of the share account.
// shareName - The name of the share.
// providerShareSubscriptionID - To locate shareSubscription
// options - ProviderShareSubscriptionsClientGetByShareOptions contains the optional parameters for the ProviderShareSubscriptionsClient.GetByShare
// method.
func (client *ProviderShareSubscriptionsClient) GetByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, options *ProviderShareSubscriptionsClientGetByShareOptions) (ProviderShareSubscriptionsClientGetByShareResponse, error) {
	req, err := client.getByShareCreateRequest(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID, options)
	if err != nil {
		return ProviderShareSubscriptionsClientGetByShareResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderShareSubscriptionsClientGetByShareResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderShareSubscriptionsClientGetByShareResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByShareHandleResponse(resp)
}

// getByShareCreateRequest creates the GetByShare request.
func (client *ProviderShareSubscriptionsClient) getByShareCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, options *ProviderShareSubscriptionsClientGetByShareOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if providerShareSubscriptionID == "" {
		return nil, errors.New("parameter providerShareSubscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerShareSubscriptionId}", url.PathEscape(providerShareSubscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByShareHandleResponse handles the GetByShare response.
func (client *ProviderShareSubscriptionsClient) getByShareHandleResponse(resp *http.Response) (ProviderShareSubscriptionsClientGetByShareResponse, error) {
	result := ProviderShareSubscriptionsClientGetByShareResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderShareSubscription); err != nil {
		return ProviderShareSubscriptionsClientGetByShareResponse{}, err
	}
	return result, nil
}

// NewListBySharePager - List share subscriptions in a provider share
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - The name of the share account.
// shareName - The name of the share.
// options - ProviderShareSubscriptionsClientListByShareOptions contains the optional parameters for the ProviderShareSubscriptionsClient.ListByShare
// method.
func (client *ProviderShareSubscriptionsClient) NewListBySharePager(resourceGroupName string, accountName string, shareName string, options *ProviderShareSubscriptionsClientListByShareOptions) *runtime.Pager[ProviderShareSubscriptionsClientListByShareResponse] {
	return runtime.NewPager(runtime.PageProcessor[ProviderShareSubscriptionsClientListByShareResponse]{
		More: func(page ProviderShareSubscriptionsClientListByShareResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProviderShareSubscriptionsClientListByShareResponse) (ProviderShareSubscriptionsClientListByShareResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByShareCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProviderShareSubscriptionsClientListByShareResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProviderShareSubscriptionsClientListByShareResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProviderShareSubscriptionsClientListByShareResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByShareHandleResponse(resp)
		},
	})
}

// listByShareCreateRequest creates the ListByShare request.
func (client *ProviderShareSubscriptionsClient) listByShareCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *ProviderShareSubscriptionsClientListByShareOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByShareHandleResponse handles the ListByShare response.
func (client *ProviderShareSubscriptionsClient) listByShareHandleResponse(resp *http.Response) (ProviderShareSubscriptionsClientListByShareResponse, error) {
	result := ProviderShareSubscriptionsClientListByShareResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderShareSubscriptionList); err != nil {
		return ProviderShareSubscriptionsClientListByShareResponse{}, err
	}
	return result, nil
}

// Reinstate - Reinstate share subscription in a provider share
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - The name of the share account.
// shareName - The name of the share.
// providerShareSubscriptionID - To locate shareSubscription
// providerShareSubscription - The provider share subscription
// options - ProviderShareSubscriptionsClientReinstateOptions contains the optional parameters for the ProviderShareSubscriptionsClient.Reinstate
// method.
func (client *ProviderShareSubscriptionsClient) Reinstate(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, providerShareSubscription ProviderShareSubscription, options *ProviderShareSubscriptionsClientReinstateOptions) (ProviderShareSubscriptionsClientReinstateResponse, error) {
	req, err := client.reinstateCreateRequest(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID, providerShareSubscription, options)
	if err != nil {
		return ProviderShareSubscriptionsClientReinstateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderShareSubscriptionsClientReinstateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderShareSubscriptionsClientReinstateResponse{}, runtime.NewResponseError(resp)
	}
	return client.reinstateHandleResponse(resp)
}

// reinstateCreateRequest creates the Reinstate request.
func (client *ProviderShareSubscriptionsClient) reinstateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, providerShareSubscription ProviderShareSubscription, options *ProviderShareSubscriptionsClientReinstateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}/reinstate"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if providerShareSubscriptionID == "" {
		return nil, errors.New("parameter providerShareSubscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerShareSubscriptionId}", url.PathEscape(providerShareSubscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, providerShareSubscription)
}

// reinstateHandleResponse handles the Reinstate response.
func (client *ProviderShareSubscriptionsClient) reinstateHandleResponse(resp *http.Response) (ProviderShareSubscriptionsClientReinstateResponse, error) {
	result := ProviderShareSubscriptionsClientReinstateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderShareSubscription); err != nil {
		return ProviderShareSubscriptionsClientReinstateResponse{}, err
	}
	return result, nil
}

// BeginRevoke - Revoke share subscription in a provider share
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - The name of the share account.
// shareName - The name of the share.
// providerShareSubscriptionID - To locate shareSubscription
// options - ProviderShareSubscriptionsClientBeginRevokeOptions contains the optional parameters for the ProviderShareSubscriptionsClient.BeginRevoke
// method.
func (client *ProviderShareSubscriptionsClient) BeginRevoke(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, options *ProviderShareSubscriptionsClientBeginRevokeOptions) (*armruntime.Poller[ProviderShareSubscriptionsClientRevokeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.revoke(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ProviderShareSubscriptionsClientRevokeResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ProviderShareSubscriptionsClientRevokeResponse](options.ResumeToken, client.pl, nil)
	}
}

// Revoke - Revoke share subscription in a provider share
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ProviderShareSubscriptionsClient) revoke(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, options *ProviderShareSubscriptionsClientBeginRevokeOptions) (*http.Response, error) {
	req, err := client.revokeCreateRequest(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// revokeCreateRequest creates the Revoke request.
func (client *ProviderShareSubscriptionsClient) revokeCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, options *ProviderShareSubscriptionsClientBeginRevokeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}/revoke"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if providerShareSubscriptionID == "" {
		return nil, errors.New("parameter providerShareSubscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerShareSubscriptionId}", url.PathEscape(providerShareSubscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
