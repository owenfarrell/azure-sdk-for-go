//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/BudgetsList.json
func ExampleBudgetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconsumption.NewBudgetsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("<scope>",
		nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/Budget.json
func ExampleBudgetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconsumption.NewBudgetsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<scope>",
		"<budget-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/CreateOrUpdateBudget.json
func ExampleBudgetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconsumption.NewBudgetsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		"<budget-name>",
		armconsumption.Budget{
			ETag: to.Ptr("<etag>"),
			Properties: &armconsumption.BudgetProperties{
				Amount:   to.Ptr[float64](100.65),
				Category: to.Ptr(armconsumption.CategoryTypeCost),
				Filter: &armconsumption.BudgetFilter{
					And: []*armconsumption.BudgetFilterProperties{
						{
							Dimensions: &armconsumption.BudgetComparisonExpression{
								Name:     to.Ptr("<name>"),
								Operator: to.Ptr(armconsumption.BudgetOperatorTypeIn),
								Values: []*string{
									to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2"),
									to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1")},
							},
						},
						{
							Tags: &armconsumption.BudgetComparisonExpression{
								Name:     to.Ptr("<name>"),
								Operator: to.Ptr(armconsumption.BudgetOperatorTypeIn),
								Values: []*string{
									to.Ptr("Dev"),
									to.Ptr("Prod")},
							},
						},
						{
							Tags: &armconsumption.BudgetComparisonExpression{
								Name:     to.Ptr("<name>"),
								Operator: to.Ptr(armconsumption.BudgetOperatorTypeIn),
								Values: []*string{
									to.Ptr("engineering"),
									to.Ptr("sales")},
							},
						}},
				},
				Notifications: map[string]*armconsumption.Notification{
					"Actual_GreaterThan_80_Percent": {
						ContactEmails: []*string{
							to.Ptr("johndoe@contoso.com"),
							to.Ptr("janesmith@contoso.com")},
						ContactGroups: []*string{
							to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup")},
						ContactRoles: []*string{
							to.Ptr("Contributor"),
							to.Ptr("Reader")},
						Enabled:       to.Ptr(true),
						Locale:        to.Ptr(armconsumption.CultureCodeEnUs),
						Operator:      to.Ptr(armconsumption.OperatorTypeGreaterThan),
						Threshold:     to.Ptr[float64](80),
						ThresholdType: to.Ptr(armconsumption.ThresholdTypeActual),
					},
				},
				TimeGrain: to.Ptr(armconsumption.TimeGrainTypeMonthly),
				TimePeriod: &armconsumption.BudgetTimePeriod{
					EndDate:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T00:00:00Z"); return t }()),
					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-01T00:00:00Z"); return t }()),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/DeleteBudget.json
func ExampleBudgetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconsumption.NewBudgetsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"<scope>",
		"<budget-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
