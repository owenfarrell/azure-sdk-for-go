//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-list.json
func ExamplePipelineTopologiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("<resource-group-name>",
		"<account-name>",
		&armvideoanalyzer.PipelineTopologiesClientListOptions{Filter: nil,
			Top: to.Ptr[int32](2),
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-get-by-name.json
func ExamplePipelineTopologiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-topology-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-create.json
func ExamplePipelineTopologiesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-topology-name>",
		armvideoanalyzer.PipelineTopology{
			Kind: to.Ptr(armvideoanalyzer.KindLive),
			Properties: &armvideoanalyzer.PipelineTopologyProperties{
				Description: to.Ptr("<description>"),
				Parameters: []*armvideoanalyzer.ParameterDeclaration{
					{
						Name:        to.Ptr("<name>"),
						Type:        to.Ptr(armvideoanalyzer.ParameterTypeString),
						Description: to.Ptr("<description>"),
						Default:     to.Ptr("<default>"),
					},
					{
						Name:        to.Ptr("<name>"),
						Type:        to.Ptr(armvideoanalyzer.ParameterTypeSecretString),
						Description: to.Ptr("<description>"),
						Default:     to.Ptr("<default>"),
					}},
				Sinks: []armvideoanalyzer.SinkNodeBaseClassification{
					&armvideoanalyzer.VideoSink{
						Name: to.Ptr("<name>"),
						Type: to.Ptr("<type>"),
						Inputs: []*armvideoanalyzer.NodeInput{
							{
								NodeName: to.Ptr("<node-name>"),
							}},
						VideoCreationProperties: &armvideoanalyzer.VideoCreationProperties{
							Description:   to.Ptr("<description>"),
							SegmentLength: to.Ptr("<segment-length>"),
							Title:         to.Ptr("<title>"),
						},
						VideoName: to.Ptr("<video-name>"),
						VideoPublishingOptions: &armvideoanalyzer.VideoPublishingOptions{
							DisableArchive:        to.Ptr("<disable-archive>"),
							DisableRtspPublishing: to.Ptr("<disable-rtsp-publishing>"),
						},
					}},
				Sources: []armvideoanalyzer.SourceNodeBaseClassification{
					&armvideoanalyzer.RtspSource{
						Name: to.Ptr("<name>"),
						Type: to.Ptr("<type>"),
						Endpoint: &armvideoanalyzer.UnsecuredEndpoint{
							Type: to.Ptr("<type>"),
							Credentials: &armvideoanalyzer.UsernamePasswordCredentials{
								Type:     to.Ptr("<type>"),
								Password: to.Ptr("<password>"),
								Username: to.Ptr("<username>"),
							},
							URL: to.Ptr("<url>"),
						},
						Transport: to.Ptr(armvideoanalyzer.RtspTransportHTTP),
					}},
			},
			SKU: &armvideoanalyzer.SKU{
				Name: to.Ptr(armvideoanalyzer.SKUNameLiveS1),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-delete.json
func ExamplePipelineTopologiesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-topology-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-patch.json
func ExamplePipelineTopologiesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-topology-name>",
		armvideoanalyzer.PipelineTopologyUpdate{
			Properties: &armvideoanalyzer.PipelineTopologyPropertiesUpdate{
				Description: to.Ptr("<description>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
