// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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

// NetworkSecurityPerimeterConfigurationsClient contains the methods for the NetworkSecurityPerimeterConfigurations group.
// Don't use this type directly, use NewNetworkSecurityPerimeterConfigurationsClient() instead.
type NetworkSecurityPerimeterConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkSecurityPerimeterConfigurationsClient creates a new instance of NetworkSecurityPerimeterConfigurationsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkSecurityPerimeterConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkSecurityPerimeterConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkSecurityPerimeterConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a specific network security perimeter configuration with a topic or domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - resourceType - The type of the resource. This can be either \'topics\', or \'domains\'.
//   - resourceName - The name of the resource (namely, either, the topic name or domain name).
//   - perimeterGUID - Unique identifier for perimeter
//   - associationName - Association name to association network security perimeter resource to profile
//   - options - NetworkSecurityPerimeterConfigurationsClientGetOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.Get
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) Get(ctx context.Context, resourceGroupName string, resourceType NetworkSecurityPerimeterResourceType, resourceName string, perimeterGUID string, associationName string, options *NetworkSecurityPerimeterConfigurationsClientGetOptions) (NetworkSecurityPerimeterConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "NetworkSecurityPerimeterConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceType, resourceName, perimeterGUID, associationName, options)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkSecurityPerimeterConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceType NetworkSecurityPerimeterResourceType, resourceName string, perimeterGUID string, associationName string, _ *NetworkSecurityPerimeterConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{resourceType}/{resourceName}/networkSecurityPerimeterConfigurations/{perimeterGuid}.{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(string(resourceType)))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if perimeterGUID == "" {
		return nil, errors.New("parameter perimeterGUID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{perimeterGuid}", url.PathEscape(perimeterGUID))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkSecurityPerimeterConfigurationsClient) getHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationsClientGetResponse, error) {
	result := NetworkSecurityPerimeterConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfiguration); err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all network security perimeter configurations associated with a topic or domain.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - resourceType - The type of the resource. This can be either \'topics\' or \'domains\'.
//   - resourceName - The name of the resource (namely, either, the topic name or domain name).
//   - options - NetworkSecurityPerimeterConfigurationsClientListOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.NewListPager
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) NewListPager(resourceGroupName string, resourceType NetworkSecurityPerimeterResourceType, resourceName string, options *NetworkSecurityPerimeterConfigurationsClientListOptions) *runtime.Pager[NetworkSecurityPerimeterConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkSecurityPerimeterConfigurationsClientListResponse]{
		More: func(page NetworkSecurityPerimeterConfigurationsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *NetworkSecurityPerimeterConfigurationsClientListResponse) (NetworkSecurityPerimeterConfigurationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkSecurityPerimeterConfigurationsClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, resourceType, resourceName, options)
			if err != nil {
				return NetworkSecurityPerimeterConfigurationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NetworkSecurityPerimeterConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkSecurityPerimeterConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *NetworkSecurityPerimeterConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceType NetworkSecurityPerimeterResourceType, resourceName string, _ *NetworkSecurityPerimeterConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{resourceType}/{resourceName}/networkSecurityPerimeterConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(string(resourceType)))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NetworkSecurityPerimeterConfigurationsClient) listHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationsClientListResponse, error) {
	result := NetworkSecurityPerimeterConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfigurationList); err != nil {
		return NetworkSecurityPerimeterConfigurationsClientListResponse{}, err
	}
	return result, nil
}

// BeginReconcile - Reconcile a specific network security perimeter configuration for a given network security perimeter association
// with a topic or domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - resourceType - The type of the resource. This can be either \'topics\' or \'domains\'.
//   - resourceName - The name of the resource (namely, either, the topic name or domain name).
//   - perimeterGUID - Unique identifier for perimeter
//   - associationName - Association name to association network security perimeter resource to profile
//   - options - NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.BeginReconcile
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) BeginReconcile(ctx context.Context, resourceGroupName string, resourceType NetworkSecurityPerimeterResourceType, resourceName string, perimeterGUID string, associationName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*runtime.Poller[NetworkSecurityPerimeterConfigurationsClientReconcileResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reconcile(ctx, resourceGroupName, resourceType, resourceName, perimeterGUID, associationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkSecurityPerimeterConfigurationsClientReconcileResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkSecurityPerimeterConfigurationsClientReconcileResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Reconcile - Reconcile a specific network security perimeter configuration for a given network security perimeter association
// with a topic or domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
func (client *NetworkSecurityPerimeterConfigurationsClient) reconcile(ctx context.Context, resourceGroupName string, resourceType NetworkSecurityPerimeterResourceType, resourceName string, perimeterGUID string, associationName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkSecurityPerimeterConfigurationsClient.BeginReconcile"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.reconcileCreateRequest(ctx, resourceGroupName, resourceType, resourceName, perimeterGUID, associationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// reconcileCreateRequest creates the Reconcile request.
func (client *NetworkSecurityPerimeterConfigurationsClient) reconcileCreateRequest(ctx context.Context, resourceGroupName string, resourceType NetworkSecurityPerimeterResourceType, resourceName string, perimeterGUID string, associationName string, _ *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{resourceType}/{resourceName}/networkSecurityPerimeterConfigurations/{perimeterGuid}.{associationName}/reconcile"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(string(resourceType)))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if perimeterGUID == "" {
		return nil, errors.New("parameter perimeterGUID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{perimeterGuid}", url.PathEscape(perimeterGUID))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
