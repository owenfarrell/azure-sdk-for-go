//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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

// InvoiceSectionsClient contains the methods for the InvoiceSections group.
// Don't use this type directly, use NewInvoiceSectionsClient() instead.
type InvoiceSectionsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewInvoiceSectionsClient creates a new instance of InvoiceSectionsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInvoiceSectionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*InvoiceSectionsClient, error) {
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
	client := &InvoiceSectionsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an invoice section. The operation is supported only for billing accounts with
// agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// invoiceSectionName - The ID that uniquely identifies an invoice section.
// parameters - The new or updated invoice section.
// options - InvoiceSectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the InvoiceSectionsClient.BeginCreateOrUpdate
// method.
func (client *InvoiceSectionsClient) BeginCreateOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InvoiceSection, options *InvoiceSectionsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[InvoiceSectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, billingAccountName, billingProfileName, invoiceSectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[InvoiceSectionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[InvoiceSectionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an invoice section. The operation is supported only for billing accounts with agreement
// type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InvoiceSectionsClient) createOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InvoiceSection, options *InvoiceSectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InvoiceSectionsClient) createOrUpdateCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InvoiceSection, options *InvoiceSectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Gets an invoice section by its ID. The operation is supported only for billing accounts with agreement type Microsoft
// Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// invoiceSectionName - The ID that uniquely identifies an invoice section.
// options - InvoiceSectionsClientGetOptions contains the optional parameters for the InvoiceSectionsClient.Get method.
func (client *InvoiceSectionsClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *InvoiceSectionsClientGetOptions) (InvoiceSectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
	if err != nil {
		return InvoiceSectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InvoiceSectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InvoiceSectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InvoiceSectionsClient) getCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *InvoiceSectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InvoiceSectionsClient) getHandleResponse(resp *http.Response) (InvoiceSectionsClientGetResponse, error) {
	result := InvoiceSectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvoiceSection); err != nil {
		return InvoiceSectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - Lists the invoice sections that a user has access to. The operation is supported only for
// billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// options - InvoiceSectionsClientListByBillingProfileOptions contains the optional parameters for the InvoiceSectionsClient.ListByBillingProfile
// method.
func (client *InvoiceSectionsClient) NewListByBillingProfilePager(billingAccountName string, billingProfileName string, options *InvoiceSectionsClientListByBillingProfileOptions) *runtime.Pager[InvoiceSectionsClientListByBillingProfileResponse] {
	return runtime.NewPager(runtime.PageProcessor[InvoiceSectionsClientListByBillingProfileResponse]{
		More: func(page InvoiceSectionsClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InvoiceSectionsClientListByBillingProfileResponse) (InvoiceSectionsClientListByBillingProfileResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InvoiceSectionsClientListByBillingProfileResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return InvoiceSectionsClientListByBillingProfileResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InvoiceSectionsClientListByBillingProfileResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *InvoiceSectionsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *InvoiceSectionsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *InvoiceSectionsClient) listByBillingProfileHandleResponse(resp *http.Response) (InvoiceSectionsClientListByBillingProfileResponse, error) {
	result := InvoiceSectionsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvoiceSectionListResult); err != nil {
		return InvoiceSectionsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}
