/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_TrafficInfluence

import (
	"regexp"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIClient manages communication with the 3gpp-traffic-influence API v1.1.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	IndividualTrafficInfluenceSubscriptionApi *IndividualTrafficInfluenceSubscriptionApiService

	TrafficInfluenceSubscriptionApi *TrafficInfluenceSubscriptionApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.IndividualTrafficInfluenceSubscriptionApi = (*IndividualTrafficInfluenceSubscriptionApiService)(&c.common)
	c.TrafficInfluenceSubscriptionApi = (*TrafficInfluenceSubscriptionApiService)(&c.common)

	return c
}