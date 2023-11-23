//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// BenefitRecommendationsClient contains the methods for the BenefitRecommendations group.
// Don't use this type directly, use NewBenefitRecommendationsClient() instead.
type BenefitRecommendationsClient struct {
	internal *arm.Client
}

// NewBenefitRecommendationsClient creates a new instance of BenefitRecommendationsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBenefitRecommendationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*BenefitRecommendationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BenefitRecommendationsClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - List of recommendations for purchasing savings plan.
//
// Generated from API version 2022-10-01
//   - billingScope - The scope associated with benefit recommendation operations. This includes '/subscriptions/{subscriptionId}/'
//     for subscription scope,
//     '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resource group scope, /providers/Microsoft.Billing/billingAccounts/{billingAccountId}'
//     for enterprise agreement scope, and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billing profile
//     scope
//   - options - BenefitRecommendationsClientListOptions contains the optional parameters for the BenefitRecommendationsClient.NewListPager
//     method.
func (client *BenefitRecommendationsClient) NewListPager(billingScope string, options *BenefitRecommendationsClientListOptions) *runtime.Pager[BenefitRecommendationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BenefitRecommendationsClientListResponse]{
		More: func(page BenefitRecommendationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BenefitRecommendationsClientListResponse) (BenefitRecommendationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BenefitRecommendationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, billingScope, options)
			}, nil)
			if err != nil {
				return BenefitRecommendationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *BenefitRecommendationsClient) listCreateRequest(ctx context.Context, billingScope string, options *BenefitRecommendationsClientListOptions) (*policy.Request, error) {
	urlPath := "/{billingScope}/providers/Microsoft.CostManagement/benefitRecommendations"
	urlPath = strings.ReplaceAll(urlPath, "{billingScope}", billingScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BenefitRecommendationsClient) listHandleResponse(resp *http.Response) (BenefitRecommendationsClientListResponse, error) {
	result := BenefitRecommendationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitRecommendationsListResult); err != nil {
		return BenefitRecommendationsClientListResponse{}, err
	}
	return result, nil
}
