//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package pollers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
)

// FinalStateVia is the enumerated type for the possible final-state-via values.
type FinalStateVia string

const (
	// FinalStateViaAzureAsyncOp indicates the final payload comes from the Azure-AsyncOperation URL.
	FinalStateViaAzureAsyncOp FinalStateVia = "azure-async-operation"

	// FinalStateViaLocation indicates the final payload comes from the Location URL.
	FinalStateViaLocation FinalStateVia = "location"

	// FinalStateViaOriginalURI indicates the final payload comes from the original URL.
	FinalStateViaOriginalURI FinalStateVia = "original-uri"

	// FinalStateViaOpLocation indicates the final payload comes from the Operation-Location URL.
	FinalStateViaOpLocation FinalStateVia = "operation-location"
)

// KindFromToken extracts the poller kind from the provided token.
// If the pollerID doesn't match what's in the token an error is returned.
func KindFromToken(pollerID, token string) (string, error) {
	// unmarshal into JSON object to determine the poller type
	obj := map[string]interface{}{}
	err := json.Unmarshal([]byte(token), &obj)
	if err != nil {
		return "", err
	}
	t, ok := obj["type"]
	if !ok {
		return "", errors.New("missing type field")
	}
	tt, ok := t.(string)
	if !ok {
		return "", fmt.Errorf("invalid type format %T", t)
	}
	ttID, ttKind, err := DecodeID(tt)
	if err != nil {
		return "", err
	}
	// ensure poller types match
	if ttID != pollerID {
		return "", fmt.Errorf("cannot resume from this poller token.  expected %s, received %s", pollerID, ttID)
	}
	return ttKind, nil
}

// PollerType returns the concrete type of the poller (FOR TESTING PURPOSES).
func PollerType(p *Poller) reflect.Type {
	return reflect.TypeOf(p.lro)
}

// NewPoller creates a Poller from the specified input.
func NewPoller(lro Operation, resp *http.Response, pl exported.Pipeline) *Poller {
	return &Poller{lro: lro, pl: pl, resp: resp}
}

// Poller encapsulates state and logic for polling on long-running operations.
type Poller struct {
	lro  Operation
	pl   exported.Pipeline
	resp *http.Response
	err  error
}

// Done returns true if the LRO has reached a terminal state.
func (l *Poller) Done() bool {
	if l.err != nil {
		return true
	}
	return l.lro.State() != OperationStateInProgress
}

// Poll sends a polling request to the polling endpoint and returns the response or error.
func (l *Poller) Poll(ctx context.Context) (*http.Response, error) {
	if l.Done() {
		// the LRO has reached a terminal state, don't poll again
		if l.resp != nil {
			return l.resp, nil
		}
		return nil, l.err
	}
	req, err := exported.NewRequest(ctx, http.MethodGet, l.lro.URL())
	if err != nil {
		return nil, err
	}
	resp, err := l.pl.Do(req)
	if err != nil {
		// don't update the poller for failed requests
		return nil, err
	}
	defer resp.Body.Close()
	if !StatusCodeValid(resp) {
		// the LRO failed.  unmarshall the error and update state
		l.err = exported.NewResponseError(resp)
		l.resp = nil
		return nil, l.err
	}
	if err = l.lro.Update(resp); err != nil {
		return nil, err
	}
	l.resp = resp
	log.Writef(log.EventLRO, "State %s", l.lro.State())
	if l.lro.State() == OperationStateFailed {
		l.err = exported.NewResponseError(resp)
		l.resp = nil
		return nil, l.err
	}
	return l.resp, nil
}

// ResumeToken returns a token string that can be used to resume a poller that has not yet reached a terminal state.
func (l *Poller) ResumeToken() (string, error) {
	if l.Done() {
		return "", errors.New("cannot create a ResumeToken from a poller in a terminal state")
	}
	b, err := json.Marshal(l.lro)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FinalResponse will perform a final GET request and return the final HTTP response for the polling
// operation and unmarshall the content of the payload into the respType interface that is provided.
func (l *Poller) FinalResponse(ctx context.Context, respType interface{}) (*http.Response, error) {
	if !l.Done() {
		return nil, errors.New("cannot return a final response from a poller in a non-terminal state")
	}
	// update l.resp with the content from final GET if applicable
	if u := l.lro.FinalGetURL(); u != "" {
		log.Write(log.EventLRO, "Performing final GET.")
		req, err := exported.NewRequest(ctx, http.MethodGet, u)
		if err != nil {
			return nil, err
		}
		resp, err := l.pl.Do(req)
		if err != nil {
			return nil, err
		}
		if !StatusCodeValid(resp) {
			return nil, exported.NewResponseError(resp)
		}
		l.resp = resp
	}
	// if there's nothing to unmarshall into or no response body just return the final response
	if respType == nil {
		return l.resp, nil
	} else if l.resp.StatusCode == http.StatusNoContent || l.resp.ContentLength == 0 {
		log.Write(log.EventLRO, "final response specifies a response type but no payload was received")
		return l.resp, nil
	}
	body, err := exported.Payload(l.resp)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, respType); err != nil {
		return nil, err
	}
	return l.resp, nil
}

// PollUntilDone will handle the entire span of the polling operation until a terminal state is reached,
// then return the final HTTP response for the polling operation and unmarshal the content of the payload
// into the respType interface that is provided.
// freq - the time to wait between intervals in absence of a Retry-After header.  Minimum is one second.
func (l *Poller) PollUntilDone(ctx context.Context, freq time.Duration, respType interface{}) (*http.Response, error) {
	if freq < time.Second {
		return nil, errors.New("polling frequency minimum is one second")
	}
	start := time.Now()
	logPollUntilDoneExit := func(v interface{}) {
		log.Writef(log.EventLRO, "END PollUntilDone() for %T: %v, total time: %s", l.lro, v, time.Since(start))
	}
	log.Writef(log.EventLRO, "BEGIN PollUntilDone() for %T", l.lro)
	if l.resp != nil {
		// initial check for a retry-after header existing on the initial response
		if retryAfter := shared.RetryAfter(l.resp); retryAfter > 0 {
			log.Writef(log.EventLRO, "initial Retry-After delay for %s", retryAfter.String())
			if err := shared.Delay(ctx, retryAfter); err != nil {
				logPollUntilDoneExit(err)
				return nil, err
			}
		}
	}
	// begin polling the endpoint until a terminal state is reached
	for {
		resp, err := l.Poll(ctx)
		if err != nil {
			logPollUntilDoneExit(err)
			return nil, err
		}
		if l.Done() {
			logPollUntilDoneExit(l.lro.State())
			return l.FinalResponse(ctx, respType)
		}
		d := freq
		if retryAfter := shared.RetryAfter(resp); retryAfter > 0 {
			log.Writef(log.EventLRO, "Retry-After delay for %s", retryAfter.String())
			d = retryAfter
		} else {
			log.Writef(log.EventLRO, "delay for %s", d.String())
		}
		if err = shared.Delay(ctx, d); err != nil {
			logPollUntilDoneExit(err)
			return nil, err
		}
	}
}
