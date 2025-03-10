//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoritesList.json
func ExampleFavoritesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx,
		"<resource-group-name>",
		"<resource-name>",
		&armapplicationinsights.FavoritesClientListOptions{FavoriteType: nil,
			SourceType:      nil,
			CanFetchContent: nil,
			Tags:            []string{},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteGet.json
func ExampleFavoritesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<favorite-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteAdd.json
func ExampleFavoritesClient_Add() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Add(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<favorite-id>",
		armapplicationinsights.ComponentFavorite{
			Config:                  to.Ptr("<config>"),
			FavoriteID:              to.Ptr("<favorite-id>"),
			FavoriteType:            to.Ptr(armapplicationinsights.FavoriteTypeShared),
			IsGeneratedFromTemplate: to.Ptr(false),
			Name:                    to.Ptr("<name>"),
			Tags: []*string{
				to.Ptr("TagSample01"),
				to.Ptr("TagSample02")},
			Version: to.Ptr("<version>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteUpdate.json
func ExampleFavoritesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<favorite-id>",
		armapplicationinsights.ComponentFavorite{
			Config:                  to.Ptr("<config>"),
			FavoriteID:              to.Ptr("<favorite-id>"),
			FavoriteType:            to.Ptr(armapplicationinsights.FavoriteTypeShared),
			IsGeneratedFromTemplate: to.Ptr(false),
			Name:                    to.Ptr("<name>"),
			Tags: []*string{
				to.Ptr("TagSample01"),
				to.Ptr("TagSample02"),
				to.Ptr("TagSample03")},
			TimeModified: to.Ptr("<time-modified>"),
			Version:      to.Ptr("<version>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteDelete.json
func ExampleFavoritesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<favorite-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
