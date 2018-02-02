package insights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// ComponentFeatureCapabilitiesClient is the composite Swagger for Application Insights Management Client
type ComponentFeatureCapabilitiesClient struct {
	ManagementClient
}

// NewComponentFeatureCapabilitiesClient creates an instance of the ComponentFeatureCapabilitiesClient client.
func NewComponentFeatureCapabilitiesClient(p pipeline.Pipeline) ComponentFeatureCapabilitiesClient {
	return ComponentFeatureCapabilitiesClient{NewManagementClient(p)}
}

// Get returns feature capabilites of the application insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights component
// resource.
func (client ComponentFeatureCapabilitiesClient) Get(ctx context.Context, resourceGroupName string, resourceName string) (*ApplicationInsightsComponentFeatureCapabilities, error) {
	req, err := client.getPreparer(resourceGroupName, resourceName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ApplicationInsightsComponentFeatureCapabilities), err
}

// getPreparer prepares the Get request.
func (client ComponentFeatureCapabilitiesClient) getPreparer(resourceGroupName string, resourceName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/featurecapabilities"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client ComponentFeatureCapabilitiesClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ApplicationInsightsComponentFeatureCapabilities{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
