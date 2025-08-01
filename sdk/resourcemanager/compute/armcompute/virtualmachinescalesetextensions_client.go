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

// VirtualMachineScaleSetExtensionsClient contains the methods for the VirtualMachineScaleSetExtensions group.
// Don't use this type directly, use NewVirtualMachineScaleSetExtensionsClient() instead.
type VirtualMachineScaleSetExtensionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualMachineScaleSetExtensionsClient creates a new instance of VirtualMachineScaleSetExtensionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualMachineScaleSetExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineScaleSetExtensionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualMachineScaleSetExtensionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update an extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmScaleSetName - The name of the VM scale set.
//   - vmssExtensionName - The name of the VM scale set extension.
//   - extensionParameters - Parameters supplied to the Create VM scale set Extension operation.
//   - options - VirtualMachineScaleSetExtensionsClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachineScaleSetExtensionsClient.BeginCreateOrUpdate
//     method.
func (client *VirtualMachineScaleSetExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachineScaleSetExtensionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineScaleSetExtensionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineScaleSetExtensionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - The operation to create or update an extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *VirtualMachineScaleSetExtensionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineScaleSetExtensionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualMachineScaleSetExtensionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, _ *VirtualMachineScaleSetExtensionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, extensionParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmScaleSetName - The name of the VM scale set.
//   - vmssExtensionName - The name of the VM scale set extension.
//   - options - VirtualMachineScaleSetExtensionsClientBeginDeleteOptions contains the optional parameters for the VirtualMachineScaleSetExtensionsClient.BeginDelete
//     method.
func (client *VirtualMachineScaleSetExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsClientBeginDeleteOptions) (*runtime.Poller[VirtualMachineScaleSetExtensionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineScaleSetExtensionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineScaleSetExtensionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The operation to delete the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *VirtualMachineScaleSetExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineScaleSetExtensionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualMachineScaleSetExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, _ *VirtualMachineScaleSetExtensionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
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

// Get - The operation to get the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmScaleSetName - The name of the VM scale set.
//   - vmssExtensionName - The name of the VM scale set extension.
//   - options - VirtualMachineScaleSetExtensionsClientGetOptions contains the optional parameters for the VirtualMachineScaleSetExtensionsClient.Get
//     method.
func (client *VirtualMachineScaleSetExtensionsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsClientGetOptions) (VirtualMachineScaleSetExtensionsClientGetResponse, error) {
	var err error
	const operationName = "VirtualMachineScaleSetExtensionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
	if err != nil {
		return VirtualMachineScaleSetExtensionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineScaleSetExtensionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineScaleSetExtensionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineScaleSetExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetExtensionsClient) getHandleResponse(resp *http.Response) (VirtualMachineScaleSetExtensionsClientGetResponse, error) {
	result := VirtualMachineScaleSetExtensionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineScaleSetExtension); err != nil {
		return VirtualMachineScaleSetExtensionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of all extensions in a VM scale set.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmScaleSetName - The name of the VM scale set.
//   - options - VirtualMachineScaleSetExtensionsClientListOptions contains the optional parameters for the VirtualMachineScaleSetExtensionsClient.NewListPager
//     method.
func (client *VirtualMachineScaleSetExtensionsClient) NewListPager(resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetExtensionsClientListOptions) *runtime.Pager[VirtualMachineScaleSetExtensionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachineScaleSetExtensionsClientListResponse]{
		More: func(page VirtualMachineScaleSetExtensionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineScaleSetExtensionsClientListResponse) (VirtualMachineScaleSetExtensionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VirtualMachineScaleSetExtensionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
			}, nil)
			if err != nil {
				return VirtualMachineScaleSetExtensionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *VirtualMachineScaleSetExtensionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, _ *VirtualMachineScaleSetExtensionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
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

// listHandleResponse handles the List response.
func (client *VirtualMachineScaleSetExtensionsClient) listHandleResponse(resp *http.Response) (VirtualMachineScaleSetExtensionsClientListResponse, error) {
	result := VirtualMachineScaleSetExtensionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineScaleSetExtensionListResult); err != nil {
		return VirtualMachineScaleSetExtensionsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update an extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmScaleSetName - The name of the VM scale set.
//   - vmssExtensionName - The name of the VM scale set extension.
//   - extensionParameters - Parameters supplied to the Update VM scale set Extension operation.
//   - options - VirtualMachineScaleSetExtensionsClientBeginUpdateOptions contains the optional parameters for the VirtualMachineScaleSetExtensionsClient.BeginUpdate
//     method.
func (client *VirtualMachineScaleSetExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsClientBeginUpdateOptions) (*runtime.Poller[VirtualMachineScaleSetExtensionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineScaleSetExtensionsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineScaleSetExtensionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - The operation to update an extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *VirtualMachineScaleSetExtensionsClient) update(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineScaleSetExtensionsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, _ *VirtualMachineScaleSetExtensionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, extensionParameters); err != nil {
		return nil, err
	}
	return req, nil
}
