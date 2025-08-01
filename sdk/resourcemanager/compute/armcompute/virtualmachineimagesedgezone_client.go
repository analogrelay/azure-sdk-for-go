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
	"strconv"
	"strings"
)

// VirtualMachineImagesEdgeZoneClient contains the methods for the VirtualMachineImagesEdgeZone group.
// Don't use this type directly, use NewVirtualMachineImagesEdgeZoneClient() instead.
type VirtualMachineImagesEdgeZoneClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualMachineImagesEdgeZoneClient creates a new instance of VirtualMachineImagesEdgeZoneClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualMachineImagesEdgeZoneClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineImagesEdgeZoneClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualMachineImagesEdgeZoneClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a virtual machine image in an edge zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - location - The name of Azure region.
//   - edgeZone - The name of the edge zone.
//   - publisherName - A valid image publisher.
//   - offer - A valid image publisher offer.
//   - skus - A valid image SKU.
//   - version - A valid image SKU version.
//   - options - VirtualMachineImagesEdgeZoneClientGetOptions contains the optional parameters for the VirtualMachineImagesEdgeZoneClient.Get
//     method.
func (client *VirtualMachineImagesEdgeZoneClient) Get(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesEdgeZoneClientGetOptions) (VirtualMachineImagesEdgeZoneClientGetResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesEdgeZoneClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, edgeZone, publisherName, offer, skus, version, options)
	if err != nil {
		return VirtualMachineImagesEdgeZoneClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesEdgeZoneClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesEdgeZoneClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineImagesEdgeZoneClient) getCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, version string, _ *VirtualMachineImagesEdgeZoneClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if skus == "" {
		return nil, errors.New("parameter skus cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
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

// getHandleResponse handles the Get response.
func (client *VirtualMachineImagesEdgeZoneClient) getHandleResponse(resp *http.Response) (VirtualMachineImagesEdgeZoneClientGetResponse, error) {
	result := VirtualMachineImagesEdgeZoneClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImage); err != nil {
		return VirtualMachineImagesEdgeZoneClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of all virtual machine image versions for the specified location, edge zone, publisher, offer, and SKU.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - location - The name of Azure region.
//   - edgeZone - The name of the edge zone.
//   - publisherName - A valid image publisher.
//   - offer - A valid image publisher offer.
//   - skus - A valid image SKU.
//   - options - VirtualMachineImagesEdgeZoneClientListOptions contains the optional parameters for the VirtualMachineImagesEdgeZoneClient.List
//     method.
func (client *VirtualMachineImagesEdgeZoneClient) List(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, options *VirtualMachineImagesEdgeZoneClientListOptions) (VirtualMachineImagesEdgeZoneClientListResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesEdgeZoneClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, location, edgeZone, publisherName, offer, skus, options)
	if err != nil {
		return VirtualMachineImagesEdgeZoneClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesEdgeZoneClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesEdgeZoneClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *VirtualMachineImagesEdgeZoneClient) listCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, options *VirtualMachineImagesEdgeZoneClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if skus == "" {
		return nil, errors.New("parameter skus cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineImagesEdgeZoneClient) listHandleResponse(resp *http.Response) (VirtualMachineImagesEdgeZoneClientListResponse, error) {
	result := VirtualMachineImagesEdgeZoneClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesEdgeZoneClientListResponse{}, err
	}
	return result, nil
}

// ListOffers - Gets a list of virtual machine image offers for the specified location, edge zone and publisher.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - location - The name of Azure region.
//   - edgeZone - The name of the edge zone.
//   - publisherName - A valid image publisher.
//   - options - VirtualMachineImagesEdgeZoneClientListOffersOptions contains the optional parameters for the VirtualMachineImagesEdgeZoneClient.ListOffers
//     method.
func (client *VirtualMachineImagesEdgeZoneClient) ListOffers(ctx context.Context, location string, edgeZone string, publisherName string, options *VirtualMachineImagesEdgeZoneClientListOffersOptions) (VirtualMachineImagesEdgeZoneClientListOffersResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesEdgeZoneClient.ListOffers"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listOffersCreateRequest(ctx, location, edgeZone, publisherName, options)
	if err != nil {
		return VirtualMachineImagesEdgeZoneClientListOffersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesEdgeZoneClientListOffersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesEdgeZoneClientListOffersResponse{}, err
	}
	resp, err := client.listOffersHandleResponse(httpResp)
	return resp, err
}

// listOffersCreateRequest creates the ListOffers request.
func (client *VirtualMachineImagesEdgeZoneClient) listOffersCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, _ *VirtualMachineImagesEdgeZoneClientListOffersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
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

// listOffersHandleResponse handles the ListOffers response.
func (client *VirtualMachineImagesEdgeZoneClient) listOffersHandleResponse(resp *http.Response) (VirtualMachineImagesEdgeZoneClientListOffersResponse, error) {
	result := VirtualMachineImagesEdgeZoneClientListOffersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesEdgeZoneClientListOffersResponse{}, err
	}
	return result, nil
}

// ListPublishers - Gets a list of virtual machine image publishers for the specified Azure location and edge zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - location - The name of Azure region.
//   - edgeZone - The name of the edge zone.
//   - options - VirtualMachineImagesEdgeZoneClientListPublishersOptions contains the optional parameters for the VirtualMachineImagesEdgeZoneClient.ListPublishers
//     method.
func (client *VirtualMachineImagesEdgeZoneClient) ListPublishers(ctx context.Context, location string, edgeZone string, options *VirtualMachineImagesEdgeZoneClientListPublishersOptions) (VirtualMachineImagesEdgeZoneClientListPublishersResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesEdgeZoneClient.ListPublishers"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listPublishersCreateRequest(ctx, location, edgeZone, options)
	if err != nil {
		return VirtualMachineImagesEdgeZoneClientListPublishersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesEdgeZoneClientListPublishersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesEdgeZoneClientListPublishersResponse{}, err
	}
	resp, err := client.listPublishersHandleResponse(httpResp)
	return resp, err
}

// listPublishersCreateRequest creates the ListPublishers request.
func (client *VirtualMachineImagesEdgeZoneClient) listPublishersCreateRequest(ctx context.Context, location string, edgeZone string, _ *VirtualMachineImagesEdgeZoneClientListPublishersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
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

// listPublishersHandleResponse handles the ListPublishers response.
func (client *VirtualMachineImagesEdgeZoneClient) listPublishersHandleResponse(resp *http.Response) (VirtualMachineImagesEdgeZoneClientListPublishersResponse, error) {
	result := VirtualMachineImagesEdgeZoneClientListPublishersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesEdgeZoneClientListPublishersResponse{}, err
	}
	return result, nil
}

// ListSKUs - Gets a list of virtual machine image SKUs for the specified location, edge zone, publisher, and offer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - location - The name of Azure region.
//   - edgeZone - The name of the edge zone.
//   - publisherName - A valid image publisher.
//   - offer - A valid image publisher offer.
//   - options - VirtualMachineImagesEdgeZoneClientListSKUsOptions contains the optional parameters for the VirtualMachineImagesEdgeZoneClient.ListSKUs
//     method.
func (client *VirtualMachineImagesEdgeZoneClient) ListSKUs(ctx context.Context, location string, edgeZone string, publisherName string, offer string, options *VirtualMachineImagesEdgeZoneClientListSKUsOptions) (VirtualMachineImagesEdgeZoneClientListSKUsResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesEdgeZoneClient.ListSKUs"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listSKUsCreateRequest(ctx, location, edgeZone, publisherName, offer, options)
	if err != nil {
		return VirtualMachineImagesEdgeZoneClientListSKUsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesEdgeZoneClientListSKUsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesEdgeZoneClientListSKUsResponse{}, err
	}
	resp, err := client.listSKUsHandleResponse(httpResp)
	return resp, err
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *VirtualMachineImagesEdgeZoneClient) listSKUsCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, offer string, _ *VirtualMachineImagesEdgeZoneClientListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
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

// listSKUsHandleResponse handles the ListSKUs response.
func (client *VirtualMachineImagesEdgeZoneClient) listSKUsHandleResponse(resp *http.Response) (VirtualMachineImagesEdgeZoneClientListSKUsResponse, error) {
	result := VirtualMachineImagesEdgeZoneClientListSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesEdgeZoneClientListSKUsResponse{}, err
	}
	return result, nil
}
