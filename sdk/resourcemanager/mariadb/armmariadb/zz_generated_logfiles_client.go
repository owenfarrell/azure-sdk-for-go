//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmariadb

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

// LogFilesClient contains the methods for the LogFiles group.
// Don't use this type directly, use NewLogFilesClient() instead.
type LogFilesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLogFilesClient creates a new instance of LogFilesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLogFilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LogFilesClient, error) {
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
	client := &LogFilesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByServerPager - List all the log files in a given server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// options - LogFilesClientListByServerOptions contains the optional parameters for the LogFilesClient.ListByServer method.
func (client *LogFilesClient) NewListByServerPager(resourceGroupName string, serverName string, options *LogFilesClientListByServerOptions) *runtime.Pager[LogFilesClientListByServerResponse] {
	return runtime.NewPager(runtime.PageProcessor[LogFilesClientListByServerResponse]{
		More: func(page LogFilesClientListByServerResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *LogFilesClientListByServerResponse) (LogFilesClientListByServerResponse, error) {
			req, err := client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			if err != nil {
				return LogFilesClientListByServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LogFilesClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LogFilesClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *LogFilesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *LogFilesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/logFiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *LogFilesClient) listByServerHandleResponse(resp *http.Response) (LogFilesClientListByServerResponse, error) {
	result := LogFilesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogFileListResult); err != nil {
		return LogFilesClientListByServerResponse{}, err
	}
	return result, nil
}
