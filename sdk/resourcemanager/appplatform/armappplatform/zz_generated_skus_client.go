//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// SKUsClient contains the methods for the SKUs group.
// Don't use this type directly, use NewSKUsClient() instead.
type SKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSKUsClient creates a new instance of SKUsClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SKUsClient, error) {
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
	client := &SKUsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Lists all of the available skus of the Microsoft.AppPlatform provider.
// If the operation fails it returns an *azcore.ResponseError type.
// options - SKUsClientListOptions contains the optional parameters for the SKUsClient.List method.
func (client *SKUsClient) NewListPager(options *SKUsClientListOptions) *runtime.Pager[SKUsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[SKUsClientListResponse]{
		More: func(page SKUsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SKUsClientListResponse) (SKUsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SKUsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SKUsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SKUsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SKUsClient) listCreateRequest(ctx context.Context, options *SKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppPlatform/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SKUsClient) listHandleResponse(resp *http.Response) (SKUsClientListResponse, error) {
	result := SKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSKUCollection); err != nil {
		return SKUsClientListResponse{}, err
	}
	return result, nil
}
