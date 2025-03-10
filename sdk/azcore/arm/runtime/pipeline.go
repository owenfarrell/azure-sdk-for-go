//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"errors"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	azpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// NewPipeline creates a pipeline from connection options.
// The telemetry policy, when enabled, will use the specified module and version info.
func NewPipeline(module, version string, cred azcore.TokenCredential, plOpts azruntime.PipelineOptions, options *arm.ClientOptions) (azruntime.Pipeline, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	conf, err := getConfiguration(&options.ClientOptions)
	if err != nil {
		return azruntime.Pipeline{}, err
	}
	authPolicy := NewBearerTokenPolicy(cred, &armpolicy.BearerTokenOptions{Scopes: []string{conf.Audience + "/.default"}})
	perRetry := make([]azpolicy.Policy, 0, len(plOpts.PerRetry)+1)
	copy(perRetry, plOpts.PerRetry)
	plOpts.PerRetry = append(perRetry, authPolicy)
	if !options.DisableRPRegistration {
		regRPOpts := armpolicy.RegistrationOptions{ClientOptions: options.ClientOptions}
		regPolicy, err := NewRPRegistrationPolicy(cred, &regRPOpts)
		if err != nil {
			return azruntime.Pipeline{}, err
		}
		perCall := make([]azpolicy.Policy, 0, len(plOpts.PerCall)+1)
		copy(perCall, plOpts.PerCall)
		plOpts.PerCall = append(perCall, regPolicy)
	}
	return azruntime.NewPipeline(module, version, plOpts, &options.ClientOptions), nil
}

func getConfiguration(o *azpolicy.ClientOptions) (cloud.ServiceConfiguration, error) {
	c := cloud.AzurePublic
	if !reflect.ValueOf(o.Cloud).IsZero() {
		c = o.Cloud
	}
	if conf, ok := c.Services[cloud.ResourceManager]; ok && conf.Endpoint != "" && conf.Audience != "" {
		return conf, nil
	} else {
		return conf, errors.New("provided Cloud field is missing Azure Resource Manager configuration")
	}
}
