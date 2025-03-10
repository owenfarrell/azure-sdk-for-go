//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

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

// ApplicationPackageClient contains the methods for the ApplicationPackage group.
// Don't use this type directly, use NewApplicationPackageClient() instead.
type ApplicationPackageClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationPackageClient creates a new instance of ApplicationPackageClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApplicationPackageClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationPackageClient, error) {
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
	client := &ApplicationPackageClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Activate - Activates the specified application package. This should be done after the ApplicationPackage was created and
// uploaded. This needs to be done before an ApplicationPackage can be used on Pools or
// Tasks.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// applicationName - The name of the application. This must be unique within the account.
// versionName - The version of the application.
// parameters - The parameters for the request.
// options - ApplicationPackageClientActivateOptions contains the optional parameters for the ApplicationPackageClient.Activate
// method.
func (client *ApplicationPackageClient) Activate(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, parameters ActivateApplicationPackageParameters, options *ApplicationPackageClientActivateOptions) (ApplicationPackageClientActivateResponse, error) {
	req, err := client.activateCreateRequest(ctx, resourceGroupName, accountName, applicationName, versionName, parameters, options)
	if err != nil {
		return ApplicationPackageClientActivateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationPackageClientActivateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationPackageClientActivateResponse{}, runtime.NewResponseError(resp)
	}
	return client.activateHandleResponse(resp)
}

// activateCreateRequest creates the Activate request.
func (client *ApplicationPackageClient) activateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, parameters ActivateApplicationPackageParameters, options *ApplicationPackageClientActivateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}/activate"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// activateHandleResponse handles the Activate response.
func (client *ApplicationPackageClient) activateHandleResponse(resp *http.Response) (ApplicationPackageClientActivateResponse, error) {
	result := ApplicationPackageClientActivateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationPackage); err != nil {
		return ApplicationPackageClientActivateResponse{}, err
	}
	return result, nil
}

// Create - Creates an application package record. The record contains a storageUrl where the package should be uploaded to.
// Once it is uploaded the ApplicationPackage needs to be activated using
// ApplicationPackageActive before it can be used. If the auto storage account was configured to use storage keys, the URL
// returned will contain a SAS.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// applicationName - The name of the application. This must be unique within the account.
// versionName - The version of the application.
// options - ApplicationPackageClientCreateOptions contains the optional parameters for the ApplicationPackageClient.Create
// method.
func (client *ApplicationPackageClient) Create(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, options *ApplicationPackageClientCreateOptions) (ApplicationPackageClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, applicationName, versionName, options)
	if err != nil {
		return ApplicationPackageClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationPackageClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationPackageClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ApplicationPackageClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, options *ApplicationPackageClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ApplicationPackageClient) createHandleResponse(resp *http.Response) (ApplicationPackageClientCreateResponse, error) {
	result := ApplicationPackageClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationPackage); err != nil {
		return ApplicationPackageClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an application package record and its associated binary file.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// applicationName - The name of the application. This must be unique within the account.
// versionName - The version of the application.
// options - ApplicationPackageClientDeleteOptions contains the optional parameters for the ApplicationPackageClient.Delete
// method.
func (client *ApplicationPackageClient) Delete(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, options *ApplicationPackageClientDeleteOptions) (ApplicationPackageClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, applicationName, versionName, options)
	if err != nil {
		return ApplicationPackageClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationPackageClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ApplicationPackageClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ApplicationPackageClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationPackageClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, options *ApplicationPackageClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets information about the specified application package.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// applicationName - The name of the application. This must be unique within the account.
// versionName - The version of the application.
// options - ApplicationPackageClientGetOptions contains the optional parameters for the ApplicationPackageClient.Get method.
func (client *ApplicationPackageClient) Get(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, options *ApplicationPackageClientGetOptions) (ApplicationPackageClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, applicationName, versionName, options)
	if err != nil {
		return ApplicationPackageClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationPackageClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationPackageClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationPackageClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, options *ApplicationPackageClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationPackageClient) getHandleResponse(resp *http.Response) (ApplicationPackageClientGetResponse, error) {
	result := ApplicationPackageClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationPackage); err != nil {
		return ApplicationPackageClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all of the application packages in the specified application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// applicationName - The name of the application. This must be unique within the account.
// options - ApplicationPackageClientListOptions contains the optional parameters for the ApplicationPackageClient.List method.
func (client *ApplicationPackageClient) NewListPager(resourceGroupName string, accountName string, applicationName string, options *ApplicationPackageClientListOptions) *runtime.Pager[ApplicationPackageClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ApplicationPackageClientListResponse]{
		More: func(page ApplicationPackageClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationPackageClientListResponse) (ApplicationPackageClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, applicationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationPackageClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationPackageClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationPackageClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ApplicationPackageClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, applicationName string, options *ApplicationPackageClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationPackageClient) listHandleResponse(resp *http.Response) (ApplicationPackageClientListResponse, error) {
	result := ApplicationPackageClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListApplicationPackagesResult); err != nil {
		return ApplicationPackageClientListResponse{}, err
	}
	return result, nil
}
