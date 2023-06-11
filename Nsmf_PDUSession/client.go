//go:build !debug
// +build !debug

/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

// APIClient manages communication with the Nsmf_PDUSession API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	IndividualPDUSessionHSMFApi        *IndividualPDUSessionHSMFApiService
	IndividualSMContextApi             *IndividualSMContextApiService
	IndividualSMContextNotificationApi *IndividualSMContextNotificationApiService
	PDUSessionsCollectionApi           *PDUSessionsCollectionApiService
	SMContextsCollectionApi            *SMContextsCollectionApiService
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
	c.IndividualPDUSessionHSMFApi = (*IndividualPDUSessionHSMFApiService)(&c.common)
	c.IndividualSMContextNotificationApi = (*IndividualSMContextNotificationApiService)(&c.common)
	c.IndividualSMContextApi = (*IndividualSMContextApiService)(&c.common)
	c.PDUSessionsCollectionApi = (*PDUSessionsCollectionApiService)(&c.common)
	c.SMContextsCollectionApi = (*SMContextsCollectionApiService)(&c.common)

	return c
}
