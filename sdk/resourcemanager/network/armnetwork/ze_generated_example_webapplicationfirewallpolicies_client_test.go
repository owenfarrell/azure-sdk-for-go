//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafListPolicies.json
func ExampleWebApplicationFirewallPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("<resource-group-name>",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafListAllPolicies.json
func ExampleWebApplicationFirewallPoliciesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListAllPager(nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafPolicyGet.json
func ExampleWebApplicationFirewallPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<policy-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafPolicyCreateOrUpdate.json
func ExampleWebApplicationFirewallPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<policy-name>",
		armnetwork.WebApplicationFirewallPolicy{
			Location: to.Ptr("<location>"),
			Properties: &armnetwork.WebApplicationFirewallPolicyPropertiesFormat{
				CustomRules: []*armnetwork.WebApplicationFirewallCustomRule{
					{
						Name:   to.Ptr("<name>"),
						Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
						MatchConditions: []*armnetwork.MatchCondition{
							{
								MatchValues: []*string{
									to.Ptr("192.168.1.0/24"),
									to.Ptr("10.0.0.0/24")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
									}},
								Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
							}},
						Priority: to.Ptr[int32](1),
						RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeMatchRule),
					},
					{
						Name:   to.Ptr("<name>"),
						Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
						MatchConditions: []*armnetwork.MatchCondition{
							{
								MatchValues: []*string{
									to.Ptr("192.168.1.0/24")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
									}},
								Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
							},
							{
								MatchValues: []*string{
									to.Ptr("Windows")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										Selector:     to.Ptr("<selector>"),
										VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRequestHeaders),
									}},
								Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorContains),
							}},
						Priority: to.Ptr[int32](2),
						RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeMatchRule),
					}},
				ManagedRules: &armnetwork.ManagedRulesDefinition{
					Exclusions: []*armnetwork.OwaspCrsExclusionEntry{
						{
							ExclusionManagedRuleSets: []*armnetwork.ExclusionManagedRuleSet{
								{
									RuleGroups: []*armnetwork.ExclusionManagedRuleGroup{
										{
											RuleGroupName: to.Ptr("<rule-group-name>"),
											Rules: []*armnetwork.ExclusionManagedRule{
												{
													RuleID: to.Ptr("<rule-id>"),
												}},
										},
										{
											RuleGroupName: to.Ptr("<rule-group-name>"),
										}},
									RuleSetType:    to.Ptr("<rule-set-type>"),
									RuleSetVersion: to.Ptr("<rule-set-version>"),
								}},
							MatchVariable:         to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgNames),
							Selector:              to.Ptr("<selector>"),
							SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith),
						},
						{
							ExclusionManagedRuleSets: []*armnetwork.ExclusionManagedRuleSet{
								{
									RuleGroups:     []*armnetwork.ExclusionManagedRuleGroup{},
									RuleSetType:    to.Ptr("<rule-set-type>"),
									RuleSetVersion: to.Ptr("<rule-set-version>"),
								}},
							MatchVariable:         to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgNames),
							Selector:              to.Ptr("<selector>"),
							SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorEndsWith),
						},
						{
							MatchVariable:         to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgNames),
							Selector:              to.Ptr("<selector>"),
							SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith),
						},
						{
							MatchVariable:         to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgValues),
							Selector:              to.Ptr("<selector>"),
							SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith),
						}},
					ManagedRuleSets: []*armnetwork.ManagedRuleSet{
						{
							RuleSetType:    to.Ptr("<rule-set-type>"),
							RuleSetVersion: to.Ptr("<rule-set-version>"),
						}},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafPolicyDelete.json
func ExampleWebApplicationFirewallPoliciesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<policy-name>",
		&armnetwork.WebApplicationFirewallPoliciesClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
