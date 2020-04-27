/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"fmt"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

type FeatureResource resource

type Feature struct {
	Links       interface{}   `json:"_links,omitempty"`
	Description string        `json:"description,omitempty"`
	Id          string        `json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Stage       *FeatureStage `json:"stage,omitempty"`
	Status      string        `json:"status,omitempty"`
	Type        string        `json:"type,omitempty"`
}

func (m *FeatureResource) GetFeature(featureId string) (*Feature, *Response, error) {
	url := fmt.Sprintf("/api/v1/features/%v", featureId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var feature *Feature

	resp, err := m.client.requestExecutor.Do(req, &feature)
	if err != nil {
		return nil, resp, err
	}

	return feature, resp, nil
}

func (m *FeatureResource) ListFeatures() ([]*Feature, *Response, error) {
	url := fmt.Sprintf("/api/v1/features")

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var feature []*Feature

	resp, err := m.client.requestExecutor.Do(req, &feature)
	if err != nil {
		return nil, resp, err
	}

	return feature, resp, nil
}

func (m *FeatureResource) ListFeatureDependencies(featureId string) ([]*Feature, *Response, error) {
	url := fmt.Sprintf("/api/v1/features/%v/dependencies", featureId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var feature []*Feature

	resp, err := m.client.requestExecutor.Do(req, &feature)
	if err != nil {
		return nil, resp, err
	}

	return feature, resp, nil
}

func (m *FeatureResource) ListFeatureDependents(featureId string) ([]*Feature, *Response, error) {
	url := fmt.Sprintf("/api/v1/features/%v/dependents", featureId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var feature []*Feature

	resp, err := m.client.requestExecutor.Do(req, &feature)
	if err != nil {
		return nil, resp, err
	}

	return feature, resp, nil
}

func (m *FeatureResource) UpdateFeatureLifecycle(featureId string, lifecycle string, qp *query.Params) (*Feature, *Response, error) {
	url := fmt.Sprintf("/api/v1/features/%v/%v", featureId, lifecycle)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var feature *Feature

	resp, err := m.client.requestExecutor.Do(req, &feature)
	if err != nil {
		return nil, resp, err
	}

	return feature, resp, nil
}