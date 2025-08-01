// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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

// ResourceSKUsClient contains the methods for the ResourceSKUs group.
// Don't use this type directly, use NewResourceSKUsClient() instead.
type ResourceSKUsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResourceSKUsClient creates a new instance of ResourceSKUsClient with the specified values.
//   - subscriptionID - Subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceSKUsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceSKUsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListSKUsPager - The skus action returns the list of SKUs that DMS (classic) supports.
//
// Generated from API version 2025-03-15-preview
//   - options - ResourceSKUsClientListSKUsOptions contains the optional parameters for the ResourceSKUsClient.NewListSKUsPager
//     method.
func (client *ResourceSKUsClient) NewListSKUsPager(options *ResourceSKUsClientListSKUsOptions) *runtime.Pager[ResourceSKUsClientListSKUsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceSKUsClientListSKUsResponse]{
		More: func(page ResourceSKUsClientListSKUsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceSKUsClientListSKUsResponse) (ResourceSKUsClientListSKUsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ResourceSKUsClient.NewListSKUsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSKUsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ResourceSKUsClientListSKUsResponse{}, err
			}
			return client.listSKUsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *ResourceSKUsClient) listSKUsCreateRequest(ctx context.Context, _ *ResourceSKUsClientListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *ResourceSKUsClient) listSKUsHandleResponse(resp *http.Response) (ResourceSKUsClientListSKUsResponse, error) {
	result := ResourceSKUsClientListSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSKUsResult); err != nil {
		return ResourceSKUsClientListSKUsResponse{}, err
	}
	return result, nil
}
