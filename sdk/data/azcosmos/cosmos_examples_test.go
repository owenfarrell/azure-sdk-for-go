// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"context"
	"encoding/json"
	"log"
	"os"
)

// This example shows you how to get started using the Azure Cosmos DB SDK for Go.
// NewCosmosClient creates a new instance of CosmosClient with the specified values. It uses the default pipeline configuration.
func Example() {

	endpoint, _ := os.LookupEnv("SOME_ENDPOINT")
	key, _ := os.LookupEnv("SOME_KEY")

	// Create new Cosmos DB client.
	cred, _ := NewSharedKeyCredential(key)
	client, err := NewCosmosClient(endpoint, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	// All operations for Go operate on a context.Context, allowing you to control cancellation/timeout.
	ctx := context.Background()

	databaseName := CosmosDatabaseProperties{Id: "databaseName"}
	database, err := client.CreateDatabase(ctx, databaseName, nil, nil)
	log.Panic(err)

	// This example showcases several common operations to help you get started, such as:

	// ===== 1. Creating a container =====

	properties := CosmosContainerProperties{
		Id: "aContainer",
		PartitionKeyDefinition: PartitionKeyDefinition{
			Paths: []string{"/id"},
		},
		IndexingPolicy: &IndexingPolicy{
			IncludedPaths: []IncludedPath{
				{Path: "/*"},
			},
			ExcludedPaths: []ExcludedPath{
				{Path: "/\"_etag\"/?"},
			},
			Automatic:    true,
			IndexingMode: IndexingModeConsistent,
		},
	}

	throughput := NewManualThroughputProperties(400)

	resp, _ := database.DatabaseProperties.Database.CreateContainer(ctx, properties, throughput, nil)
	container := resp.ContainerProperties.Container

	// ===== 2. Update container properties =====

	resp, _ = container.Read(ctx, nil)

	updatedProperties := CosmosContainerProperties{
		Id: "aContainer",
		PartitionKeyDefinition: PartitionKeyDefinition{
			Paths: []string{"/id"},
		},
		ETag: "someEtag",
		IndexingPolicy: &IndexingPolicy{
			IncludedPaths: []IncludedPath{},
			ExcludedPaths: []ExcludedPath{},
			Automatic:     false,
			IndexingMode:  IndexingModeNone,
		},
	}

	// Replace container properties
	resp, _ = container.Replace(ctx, updatedProperties, nil)

	// Read the manual throughput property
	throughputResponse, _ := container.ReadThroughput(ctx, nil)
	_, _ = throughputResponse.ThroughputProperties.ManualThroughput()

	// Replace manual throughput property
	newScale := NewManualThroughputProperties(500)
	_, _ = container.ReplaceThroughput(ctx, *newScale, nil)

	// Container autoscale throughput
	_, _ = throughputResponse.ThroughputProperties.AutoscaleMaxThroughput()

	newScale = NewAutoscaleThroughputProperties(10000)
	_, _ = container.ReplaceThroughput(ctx, *newScale, nil)

	// Delete container
	resp, _ = container.Delete(ctx, nil)

	// ===== 3. Item CRUD =====

	// Items in an Azure Cosmos container are uniquely identified by their id and partition key value.
	pk, _ := NewPartitionKey("1")

	item := map[string]string{
		"id":    "1",
		"value": "2",
	}

	// Create item.
	itemResponse, _ := container.CreateItem(ctx, pk, item, nil)

	// Read item.
	itemResponse, _ = container.ReadItem(ctx, pk, "1", nil)

	// Item response body
	var itemResponseBody map[string]interface{}
	err = json.Unmarshal(itemResponse.Value, &itemResponseBody)
	log.Panic(err)

	// ===== 4. Optimistic Concurrency / PreConditionFail =====

	// Azure Cosmos DB supports optimistic concurrency control to prevent lost updates or deletes and detection of conflicting operations.
	itemResponse, _ = container.UpsertItem(ctx, pk, item, &CosmosItemRequestOptions{EnableContentResponseOnWrite: true})

	itemSessionToken := itemResponse.SessionToken()

	itemResponse, _ = container.ReadItem(ctx, pk, "1", &CosmosItemRequestOptions{SessionToken: itemSessionToken})

	// Delete item.
	itemResponse, _ = container.DeleteItem(ctx, pk, "1", nil)
}
