//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdesktopvirtualization

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

// StartMenuItemsClient contains the methods for the StartMenuItems group.
// Don't use this type directly, use NewStartMenuItemsClient() instead.
type StartMenuItemsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStartMenuItemsClient creates a new instance of StartMenuItemsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStartMenuItemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StartMenuItemsClient, error) {
	cl, err := arm.NewClient(moduleName+".StartMenuItemsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StartMenuItemsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - List start menu items in the given application group.
//
// Generated from API version 2022-09-09
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationGroupName - The name of the application group
//   - options - StartMenuItemsClientListOptions contains the optional parameters for the StartMenuItemsClient.NewListPager method.
func (client *StartMenuItemsClient) NewListPager(resourceGroupName string, applicationGroupName string, options *StartMenuItemsClientListOptions) *runtime.Pager[StartMenuItemsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[StartMenuItemsClientListResponse]{
		More: func(page StartMenuItemsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StartMenuItemsClientListResponse) (StartMenuItemsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StartMenuItemsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return StartMenuItemsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StartMenuItemsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *StartMenuItemsClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *StartMenuItemsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}/startMenuItems"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-09")
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.IsDescending != nil {
		reqQP.Set("isDescending", strconv.FormatBool(*options.IsDescending))
	}
	if options != nil && options.InitialSkip != nil {
		reqQP.Set("initialSkip", strconv.FormatInt(int64(*options.InitialSkip), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *StartMenuItemsClient) listHandleResponse(resp *http.Response) (StartMenuItemsClientListResponse, error) {
	result := StartMenuItemsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StartMenuItemList); err != nil {
		return StartMenuItemsClientListResponse{}, err
	}
	return result, nil
}