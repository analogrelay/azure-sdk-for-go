// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DedicatedHostGroupsClient contains the methods for the DedicatedHostGroups group.
// Don't use this type directly, use NewDedicatedHostGroupsClient() instead.
type DedicatedHostGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDedicatedHostGroupsClient creates a new instance of DedicatedHostGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDedicatedHostGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DedicatedHostGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DedicatedHostGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a dedicated host group. For details of Dedicated Host and Dedicated Host Groups please
// see Dedicated Host Documentation [https://go.microsoft.com/fwlink/?linkid=2082596]
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostGroupName - The name of the dedicated host group.
//   - parameters - Parameters supplied to the Create Dedicated Host Group.
//   - options - DedicatedHostGroupsClientCreateOrUpdateOptions contains the optional parameters for the DedicatedHostGroupsClient.CreateOrUpdate
//     method.
func (client *DedicatedHostGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroup, options *DedicatedHostGroupsClientCreateOrUpdateOptions) (DedicatedHostGroupsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DedicatedHostGroupsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostGroupName, parameters, options)
	if err != nil {
		return DedicatedHostGroupsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DedicatedHostGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DedicatedHostGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DedicatedHostGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroup, _ *DedicatedHostGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DedicatedHostGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (DedicatedHostGroupsClientCreateOrUpdateResponse, error) {
	result := DedicatedHostGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroup); err != nil {
		return DedicatedHostGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a dedicated host group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostGroupName - The name of the dedicated host group.
//   - options - DedicatedHostGroupsClientDeleteOptions contains the optional parameters for the DedicatedHostGroupsClient.Delete
//     method.
func (client *DedicatedHostGroupsClient) Delete(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsClientDeleteOptions) (DedicatedHostGroupsClientDeleteResponse, error) {
	var err error
	const operationName = "DedicatedHostGroupsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostGroupName, options)
	if err != nil {
		return DedicatedHostGroupsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DedicatedHostGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DedicatedHostGroupsClientDeleteResponse{}, err
	}
	return DedicatedHostGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DedicatedHostGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, _ *DedicatedHostGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about a dedicated host group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostGroupName - The name of the dedicated host group.
//   - options - DedicatedHostGroupsClientGetOptions contains the optional parameters for the DedicatedHostGroupsClient.Get method.
func (client *DedicatedHostGroupsClient) Get(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsClientGetOptions) (DedicatedHostGroupsClientGetResponse, error) {
	var err error
	const operationName = "DedicatedHostGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostGroupName, options)
	if err != nil {
		return DedicatedHostGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DedicatedHostGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DedicatedHostGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DedicatedHostGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedHostGroupsClient) getHandleResponse(resp *http.Response) (DedicatedHostGroupsClientGetResponse, error) {
	result := DedicatedHostGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroup); err != nil {
		return DedicatedHostGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all of the dedicated host groups in the specified resource group. Use the nextLink
// property in the response to get the next page of dedicated host groups.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DedicatedHostGroupsClientListByResourceGroupOptions contains the optional parameters for the DedicatedHostGroupsClient.NewListByResourceGroupPager
//     method.
func (client *DedicatedHostGroupsClient) NewListByResourceGroupPager(resourceGroupName string, options *DedicatedHostGroupsClientListByResourceGroupOptions) *runtime.Pager[DedicatedHostGroupsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DedicatedHostGroupsClientListByResourceGroupResponse]{
		More: func(page DedicatedHostGroupsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHostGroupsClientListByResourceGroupResponse) (DedicatedHostGroupsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHostGroupsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return DedicatedHostGroupsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DedicatedHostGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *DedicatedHostGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DedicatedHostGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (DedicatedHostGroupsClientListByResourceGroupResponse, error) {
	result := DedicatedHostGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroupListResult); err != nil {
		return DedicatedHostGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all of the dedicated host groups in the subscription. Use the nextLink property in the
// response to get the next page of dedicated host groups.
//
// Generated from API version 2024-11-01
//   - options - DedicatedHostGroupsClientListBySubscriptionOptions contains the optional parameters for the DedicatedHostGroupsClient.NewListBySubscriptionPager
//     method.
func (client *DedicatedHostGroupsClient) NewListBySubscriptionPager(options *DedicatedHostGroupsClientListBySubscriptionOptions) *runtime.Pager[DedicatedHostGroupsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DedicatedHostGroupsClientListBySubscriptionResponse]{
		More: func(page DedicatedHostGroupsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHostGroupsClientListBySubscriptionResponse) (DedicatedHostGroupsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHostGroupsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DedicatedHostGroupsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DedicatedHostGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, _ *DedicatedHostGroupsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/hostGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DedicatedHostGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (DedicatedHostGroupsClientListBySubscriptionResponse, error) {
	result := DedicatedHostGroupsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroupListResult); err != nil {
		return DedicatedHostGroupsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an dedicated host group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostGroupName - The name of the dedicated host group.
//   - parameters - Parameters supplied to the Update Dedicated Host Group operation.
//   - options - DedicatedHostGroupsClientUpdateOptions contains the optional parameters for the DedicatedHostGroupsClient.Update
//     method.
func (client *DedicatedHostGroupsClient) Update(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroupUpdate, options *DedicatedHostGroupsClientUpdateOptions) (DedicatedHostGroupsClientUpdateResponse, error) {
	var err error
	const operationName = "DedicatedHostGroupsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostGroupName, parameters, options)
	if err != nil {
		return DedicatedHostGroupsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DedicatedHostGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DedicatedHostGroupsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *DedicatedHostGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroupUpdate, _ *DedicatedHostGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DedicatedHostGroupsClient) updateHandleResponse(resp *http.Response) (DedicatedHostGroupsClientUpdateResponse, error) {
	result := DedicatedHostGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroup); err != nil {
		return DedicatedHostGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
