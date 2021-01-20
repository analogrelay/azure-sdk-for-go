// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// QueueClient contains the methods for the Queue group.
// Don't use this type directly, use NewQueueClient() instead.
type QueueClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewQueueClient creates a new instance of QueueClient with the specified values.
func NewQueueClient(con *armcore.Connection, subscriptionID string) *QueueClient {
	return &QueueClient{con: con, subscriptionID: subscriptionID}
}

// Create - Creates a new queue with the specified queue name, under the specified account.
func (client *QueueClient) Create(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue StorageQueue, options *QueueCreateOptions) (StorageQueueResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, queueName, queue, options)
	if err != nil {
		return StorageQueueResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return StorageQueueResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StorageQueueResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *QueueClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue StorageQueue, options *QueueCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(queue)
}

// createHandleResponse handles the Create response.
func (client *QueueClient) createHandleResponse(resp *azcore.Response) (StorageQueueResponse, error) {
	var val *StorageQueue
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StorageQueueResponse{}, err
	}
	return StorageQueueResponse{RawResponse: resp.Response, StorageQueue: val}, nil
}

// createHandleError handles the Create error response.
func (client *QueueClient) createHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete - Deletes the queue with the specified queue name, under the specified account if it exists.
func (client *QueueClient) Delete(ctx context.Context, resourceGroupName string, accountName string, queueName string, options *QueueDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, queueName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *QueueClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string, options *QueueDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *QueueClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the queue with the specified queue name, under the specified account if it exists.
func (client *QueueClient) Get(ctx context.Context, resourceGroupName string, accountName string, queueName string, options *QueueGetOptions) (StorageQueueResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, queueName, options)
	if err != nil {
		return StorageQueueResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return StorageQueueResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StorageQueueResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *QueueClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string, options *QueueGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QueueClient) getHandleResponse(resp *azcore.Response) (StorageQueueResponse, error) {
	var val *StorageQueue
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StorageQueueResponse{}, err
	}
	return StorageQueueResponse{RawResponse: resp.Response, StorageQueue: val}, nil
}

// getHandleError handles the Get error response.
func (client *QueueClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets a list of all the queues under the specified storage account
func (client *QueueClient) List(resourceGroupName string, accountName string, options *QueueListOptions) ListQueueResourcePager {
	return &listQueueResourcePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ListQueueResourceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListQueueResource.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *QueueClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *QueueListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	if options != nil && options.Maxpagesize != nil {
		query.Set("$maxpagesize", *options.Maxpagesize)
	}
	if options != nil && options.Filter != nil {
		query.Set("$filter", *options.Filter)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *QueueClient) listHandleResponse(resp *azcore.Response) (ListQueueResourceResponse, error) {
	var val *ListQueueResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListQueueResourceResponse{}, err
	}
	return ListQueueResourceResponse{RawResponse: resp.Response, ListQueueResource: val}, nil
}

// listHandleError handles the List error response.
func (client *QueueClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Update - Creates a new queue with the specified queue name, under the specified account.
func (client *QueueClient) Update(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue StorageQueue, options *QueueUpdateOptions) (StorageQueueResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, queueName, queue, options)
	if err != nil {
		return StorageQueueResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return StorageQueueResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StorageQueueResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *QueueClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue StorageQueue, options *QueueUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(queue)
}

// updateHandleResponse handles the Update response.
func (client *QueueClient) updateHandleResponse(resp *azcore.Response) (StorageQueueResponse, error) {
	var val *StorageQueue
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StorageQueueResponse{}, err
	}
	return StorageQueueResponse{RawResponse: resp.Response, StorageQueue: val}, nil
}

// updateHandleError handles the Update error response.
func (client *QueueClient) updateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
