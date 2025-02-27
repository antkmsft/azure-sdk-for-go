//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BudgetsClient contains the methods for the Budgets group.
// Don't use this type directly, use NewBudgetsClient() instead.
type BudgetsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewBudgetsClient creates a new instance of BudgetsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBudgetsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*BudgetsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &BudgetsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// CreateOrUpdate - The operation to create or update a budget. You can optionally provide an eTag if desired as a form of
// concurrency control. To obtain the latest eTag for a given budget, perform a get operation prior
// to your put operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// scope - The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
// resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
// scope,
// '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
// for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope.
// budgetName - Budget Name.
// parameters - Parameters supplied to the Create Budget operation.
// options - BudgetsClientCreateOrUpdateOptions contains the optional parameters for the BudgetsClient.CreateOrUpdate method.
func (client *BudgetsClient) CreateOrUpdate(ctx context.Context, scope string, budgetName string, parameters Budget, options *BudgetsClientCreateOrUpdateOptions) (BudgetsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, scope, budgetName, parameters, options)
	if err != nil {
		return BudgetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BudgetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return BudgetsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BudgetsClient) createOrUpdateCreateRequest(ctx context.Context, scope string, budgetName string, parameters Budget, options *BudgetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if budgetName == "" {
		return nil, errors.New("parameter budgetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{budgetName}", url.PathEscape(budgetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BudgetsClient) createOrUpdateHandleResponse(resp *http.Response) (BudgetsClientCreateOrUpdateResponse, error) {
	result := BudgetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Budget); err != nil {
		return BudgetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - The operation to delete a budget.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// scope - The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
// resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
// scope,
// '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
// for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope.
// budgetName - Budget Name.
// options - BudgetsClientDeleteOptions contains the optional parameters for the BudgetsClient.Delete method.
func (client *BudgetsClient) Delete(ctx context.Context, scope string, budgetName string, options *BudgetsClientDeleteOptions) (BudgetsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, budgetName, options)
	if err != nil {
		return BudgetsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BudgetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BudgetsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return BudgetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BudgetsClient) deleteCreateRequest(ctx context.Context, scope string, budgetName string, options *BudgetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if budgetName == "" {
		return nil, errors.New("parameter budgetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{budgetName}", url.PathEscape(budgetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the budget for the scope by budget name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// scope - The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
// resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
// scope,
// '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
// for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope.
// budgetName - Budget Name.
// options - BudgetsClientGetOptions contains the optional parameters for the BudgetsClient.Get method.
func (client *BudgetsClient) Get(ctx context.Context, scope string, budgetName string, options *BudgetsClientGetOptions) (BudgetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, budgetName, options)
	if err != nil {
		return BudgetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BudgetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BudgetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BudgetsClient) getCreateRequest(ctx context.Context, scope string, budgetName string, options *BudgetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if budgetName == "" {
		return nil, errors.New("parameter budgetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{budgetName}", url.PathEscape(budgetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BudgetsClient) getHandleResponse(resp *http.Response) (BudgetsClientGetResponse, error) {
	result := BudgetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Budget); err != nil {
		return BudgetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all budgets for the defined scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// scope - The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
// resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
// scope,
// '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
// for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope.
// options - BudgetsClientListOptions contains the optional parameters for the BudgetsClient.List method.
func (client *BudgetsClient) NewListPager(scope string, options *BudgetsClientListOptions) *runtime.Pager[BudgetsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BudgetsClientListResponse]{
		More: func(page BudgetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BudgetsClientListResponse) (BudgetsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BudgetsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BudgetsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BudgetsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BudgetsClient) listCreateRequest(ctx context.Context, scope string, options *BudgetsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/budgets"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BudgetsClient) listHandleResponse(resp *http.Response) (BudgetsClientListResponse, error) {
	result := BudgetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BudgetsListResult); err != nil {
		return BudgetsClientListResponse{}, err
	}
	return result, nil
}
