package resourcehealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AvailabilityStatusesClient is the the Resource Health Client.
type AvailabilityStatusesClient struct {
	BaseClient
}

// NewAvailabilityStatusesClient creates an instance of the AvailabilityStatusesClient client.
func NewAvailabilityStatusesClient(subscriptionID string) AvailabilityStatusesClient {
	return NewAvailabilityStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAvailabilityStatusesClientWithBaseURI creates an instance of the AvailabilityStatusesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAvailabilityStatusesClientWithBaseURI(baseURI string, subscriptionID string) AvailabilityStatusesClient {
	return AvailabilityStatusesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByResource gets current availability status for a single resource
// Parameters:
// resourceURI - the fully qualified ID of the resource, including the resource name and resource type.
// Currently the API support not nested and one nesting level resource types :
// /subscriptions/{subscriptionId}/resourceGroups/{resource-group-name}/providers/{resource-provider-name}/{resource-type}/{resource-name}
// and
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resource-provider-name}/{parentResourceType}/{parentResourceName}/{resourceType}/{resourceName}
// filter - the filter to apply on the operation. For more information please see
// https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
// expand - setting $expand=recommendedactions in url query expands the recommendedactions in the response.
func (client AvailabilityStatusesClient) GetByResource(ctx context.Context, resourceURI string, filter string, expand string) (result AvailabilityStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailabilityStatusesClient.GetByResource")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByResourcePreparer(ctx, resourceURI, filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "GetByResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "GetByResource", resp, "Failure sending request")
		return
	}

	result, err = client.GetByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "GetByResource", resp, "Failure responding to request")
		return
	}

	return
}

// GetByResourcePreparer prepares the GetByResource request.
func (client AvailabilityStatusesClient) GetByResourcePreparer(ctx context.Context, resourceURI string, filter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": resourceURI,
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses/current", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByResourceSender sends the GetByResource request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilityStatusesClient) GetByResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByResourceResponder handles the response to the GetByResource request. The method always
// closes the http.Response Body.
func (client AvailabilityStatusesClient) GetByResourceResponder(resp *http.Response) (result AvailabilityStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all historical availability transitions and impacting events for a single resource.
// Parameters:
// resourceURI - the fully qualified ID of the resource, including the resource name and resource type.
// Currently the API support not nested and one nesting level resource types :
// /subscriptions/{subscriptionId}/resourceGroups/{resource-group-name}/providers/{resource-provider-name}/{resource-type}/{resource-name}
// and
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resource-provider-name}/{parentResourceType}/{parentResourceName}/{resourceType}/{resourceName}
// filter - the filter to apply on the operation. For more information please see
// https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
// expand - setting $expand=recommendedactions in url query expands the recommendedactions in the response.
func (client AvailabilityStatusesClient) List(ctx context.Context, resourceURI string, filter string, expand string) (result AvailabilityStatusListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailabilityStatusesClient.List")
		defer func() {
			sc := -1
			if result.aslr.Response.Response != nil {
				sc = result.aslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceURI, filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.aslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "List", resp, "Failure sending request")
		return
	}

	result.aslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.aslr.hasNextLink() && result.aslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AvailabilityStatusesClient) ListPreparer(ctx context.Context, resourceURI string, filter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": resourceURI,
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilityStatusesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AvailabilityStatusesClient) ListResponder(resp *http.Response) (result AvailabilityStatusListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AvailabilityStatusesClient) listNextResults(ctx context.Context, lastResults AvailabilityStatusListResult) (result AvailabilityStatusListResult, err error) {
	req, err := lastResults.availabilityStatusListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AvailabilityStatusesClient) ListComplete(ctx context.Context, resourceURI string, filter string, expand string) (result AvailabilityStatusListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailabilityStatusesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceURI, filter, expand)
	return
}

// ListByResourceGroup lists the current availability status for all the resources in the resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// filter - the filter to apply on the operation. For more information please see
// https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
// expand - setting $expand=recommendedactions in url query expands the recommendedactions in the response.
func (client AvailabilityStatusesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, expand string) (result AvailabilityStatusListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailabilityStatusesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.aslr.Response.Response != nil {
				sc = result.aslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.aslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.aslr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.aslr.hasNextLink() && result.aslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AvailabilityStatusesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceHealth/availabilityStatuses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilityStatusesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AvailabilityStatusesClient) ListByResourceGroupResponder(resp *http.Response) (result AvailabilityStatusListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client AvailabilityStatusesClient) listByResourceGroupNextResults(ctx context.Context, lastResults AvailabilityStatusListResult) (result AvailabilityStatusListResult, err error) {
	req, err := lastResults.availabilityStatusListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client AvailabilityStatusesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, expand string) (result AvailabilityStatusListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailabilityStatusesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, expand)
	return
}

// ListBySubscriptionID lists the current availability status for all the resources in the subscription.
// Parameters:
// filter - the filter to apply on the operation. For more information please see
// https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
// expand - setting $expand=recommendedactions in url query expands the recommendedactions in the response.
func (client AvailabilityStatusesClient) ListBySubscriptionID(ctx context.Context, filter string, expand string) (result AvailabilityStatusListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailabilityStatusesClient.ListBySubscriptionID")
		defer func() {
			sc := -1
			if result.aslr.Response.Response != nil {
				sc = result.aslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionIDNextResults
	req, err := client.ListBySubscriptionIDPreparer(ctx, filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListBySubscriptionID", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionIDSender(req)
	if err != nil {
		result.aslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListBySubscriptionID", resp, "Failure sending request")
		return
	}

	result.aslr, err = client.ListBySubscriptionIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListBySubscriptionID", resp, "Failure responding to request")
		return
	}
	if result.aslr.hasNextLink() && result.aslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionIDPreparer prepares the ListBySubscriptionID request.
func (client AvailabilityStatusesClient) ListBySubscriptionIDPreparer(ctx context.Context, filter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/availabilityStatuses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionIDSender sends the ListBySubscriptionID request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilityStatusesClient) ListBySubscriptionIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionIDResponder handles the response to the ListBySubscriptionID request. The method always
// closes the http.Response Body.
func (client AvailabilityStatusesClient) ListBySubscriptionIDResponder(resp *http.Response) (result AvailabilityStatusListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionIDNextResults retrieves the next set of results, if any.
func (client AvailabilityStatusesClient) listBySubscriptionIDNextResults(ctx context.Context, lastResults AvailabilityStatusListResult) (result AvailabilityStatusListResult, err error) {
	req, err := lastResults.availabilityStatusListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "listBySubscriptionIDNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "listBySubscriptionIDNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "listBySubscriptionIDNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionIDComplete enumerates all values, automatically crossing page boundaries as required.
func (client AvailabilityStatusesClient) ListBySubscriptionIDComplete(ctx context.Context, filter string, expand string) (result AvailabilityStatusListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailabilityStatusesClient.ListBySubscriptionID")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscriptionID(ctx, filter, expand)
	return
}
