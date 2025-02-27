//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

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

// PermissionsClient contains the methods for the Permissions group.
// Don't use this type directly, use NewPermissionsClient() instead.
type PermissionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPermissionsClient creates a new instance of PermissionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPermissionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PermissionsClient, error) {
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
	client := &PermissionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListForResourcePager - Gets all permissions the caller has for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceProviderNamespace - The namespace of the resource provider.
// parentResourcePath - The parent resource identity.
// resourceType - The resource type of the resource.
// resourceName - The name of the resource to get the permissions for.
// options - PermissionsClientListForResourceOptions contains the optional parameters for the PermissionsClient.ListForResource
// method.
func (client *PermissionsClient) NewListForResourcePager(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *PermissionsClientListForResourceOptions) *runtime.Pager[PermissionsClientListForResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PermissionsClientListForResourceResponse]{
		More: func(page PermissionsClientListForResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PermissionsClientListForResourceResponse) (PermissionsClientListForResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForResourceCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PermissionsClientListForResourceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PermissionsClientListForResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PermissionsClientListForResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForResourceHandleResponse(resp)
		},
	})
}

// listForResourceCreateRequest creates the ListForResource request.
func (client *PermissionsClient) listForResourceCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *PermissionsClientListForResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/permissions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", resourceProviderNamespace)
	urlPath = strings.ReplaceAll(urlPath, "{parentResourcePath}", parentResourcePath)
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", resourceType)
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForResourceHandleResponse handles the ListForResource response.
func (client *PermissionsClient) listForResourceHandleResponse(resp *http.Response) (PermissionsClientListForResourceResponse, error) {
	result := PermissionsClientListForResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PermissionGetResult); err != nil {
		return PermissionsClientListForResourceResponse{}, err
	}
	return result, nil
}

// NewListForResourceGroupPager - Gets all permissions the caller has for a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - PermissionsClientListForResourceGroupOptions contains the optional parameters for the PermissionsClient.ListForResourceGroup
// method.
func (client *PermissionsClient) NewListForResourceGroupPager(resourceGroupName string, options *PermissionsClientListForResourceGroupOptions) *runtime.Pager[PermissionsClientListForResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PermissionsClientListForResourceGroupResponse]{
		More: func(page PermissionsClientListForResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PermissionsClientListForResourceGroupResponse) (PermissionsClientListForResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PermissionsClientListForResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PermissionsClientListForResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PermissionsClientListForResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForResourceGroupHandleResponse(resp)
		},
	})
}

// listForResourceGroupCreateRequest creates the ListForResourceGroup request.
func (client *PermissionsClient) listForResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PermissionsClientListForResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Authorization/permissions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForResourceGroupHandleResponse handles the ListForResourceGroup response.
func (client *PermissionsClient) listForResourceGroupHandleResponse(resp *http.Response) (PermissionsClientListForResourceGroupResponse, error) {
	result := PermissionsClientListForResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PermissionGetResult); err != nil {
		return PermissionsClientListForResourceGroupResponse{}, err
	}
	return result, nil
}
