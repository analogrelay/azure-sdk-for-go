// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"bytes"
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/uuid"
)

// ContainerClient lets you perform read, update, change throughput, and delete container operations.
// It also lets you perform read, update, change throughput, and delete item operations.
type ContainerClient struct {
	// The Id of the Cosmos container
	id string
	// The database that contains the container
	database *DatabaseClient
	// The resource link
	link string
}

func newContainer(id string, database *DatabaseClient) (*ContainerClient, error) {
	return &ContainerClient{
		id:       id,
		database: database,
		link:     createLink(database.link, pathSegmentCollection, id)}, nil
}

// ID returns the identifier of the Cosmos container.
func (c *ContainerClient) ID() string {
	return c.id
}

// Read obtains the information for a Cosmos container.
// ctx - The context for the request.
// o - Options for the operation.
func (c *ContainerClient) Read(
	ctx context.Context,
	o *ReadContainerOptions) (ContainerResponse, error) {
	var err error
	spanName, err := c.getSpanForContainer(operationTypeRead, resourceTypeCollection, c.id)
	if err != nil {
		return ContainerResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if o == nil {
		o = &ReadContainerOptions{}
	}

	operationContext := pipelineRequestOptions{
		resourceType:    resourceTypeCollection,
		resourceAddress: c.link,
	}

	path, err := generatePathForNameBased(resourceTypeCollection, c.link, false)
	if err != nil {
		return ContainerResponse{}, err
	}

	azResponse, err := c.database.client.sendGetRequest(
		path,
		ctx,
		operationContext,
		o,
		nil)
	if err != nil {
		return ContainerResponse{}, err
	}

	response, err := newContainerResponse(azResponse)
	return response, err
}

// Replace a Cosmos container.
// ctx - The context for the request.
// o - Options for the operation.
func (c *ContainerClient) Replace(
	ctx context.Context,
	containerProperties ContainerProperties,
	o *ReplaceContainerOptions) (ContainerResponse, error) {
	var err error
	spanName, err := c.getSpanForContainer(operationTypeReplace, resourceTypeCollection, c.id)
	if err != nil {
		return ContainerResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if o == nil {
		o = &ReplaceContainerOptions{}
	}

	operationContext := pipelineRequestOptions{
		resourceType:     resourceTypeCollection,
		resourceAddress:  c.link,
		isWriteOperation: true,
	}

	path, err := generatePathForNameBased(resourceTypeCollection, c.link, false)
	if err != nil {
		return ContainerResponse{}, err
	}

	azResponse, err := c.database.client.sendPutRequest(
		path,
		ctx,
		containerProperties,
		operationContext,
		o,
		nil)
	if err != nil {
		return ContainerResponse{}, err
	}

	response, err := newContainerResponse(azResponse)
	return response, err
}

// Delete a Cosmos container.
// ctx - The context for the request.
// o - Options for the operation.
func (c *ContainerClient) Delete(
	ctx context.Context,
	o *DeleteContainerOptions) (ContainerResponse, error) {
	var err error
	spanName, err := c.getSpanForContainer(operationTypeDelete, resourceTypeCollection, c.id)
	if err != nil {
		return ContainerResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if o == nil {
		o = &DeleteContainerOptions{}
	}

	operationContext := pipelineRequestOptions{
		resourceType:     resourceTypeCollection,
		resourceAddress:  c.link,
		isWriteOperation: true,
	}

	path, err := generatePathForNameBased(resourceTypeCollection, c.link, false)
	if err != nil {
		return ContainerResponse{}, err
	}

	azResponse, err := c.database.client.sendDeleteRequest(
		path,
		ctx,
		operationContext,
		o,
		nil)
	if err != nil {
		return ContainerResponse{}, err
	}

	response, err := newContainerResponse(azResponse)
	return response, err
}

// ReadThroughput obtains the provisioned throughput information for the container.
// ctx - The context for the request.
// o - Options for the operation.
func (c *ContainerClient) ReadThroughput(
	ctx context.Context,
	o *ThroughputOptions) (ThroughputResponse, error) {
	var err error
	spanName, err := c.getSpanForContainer(operationTypeRead, resourceTypeOffer, c.id)
	if err != nil {
		return ThroughputResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if o == nil {
		o = &ThroughputOptions{}
	}

	rid, err := c.getRID(ctx)
	if err != nil {
		return ThroughputResponse{}, err
	}

	offers := &cosmosOffers{client: c.database.client}
	response, err := offers.ReadThroughputIfExists(ctx, rid, o)
	return response, err
}

// ReplaceThroughput updates the provisioned throughput for the container.
// ctx - The context for the request.
// throughputProperties - The throughput configuration of the container.
// o - Options for the operation.
func (c *ContainerClient) ReplaceThroughput(
	ctx context.Context,
	throughputProperties ThroughputProperties,
	o *ThroughputOptions) (ThroughputResponse, error) {
	var err error
	spanName, err := c.getSpanForContainer(operationTypeReplace, resourceTypeOffer, c.id)
	if err != nil {
		return ThroughputResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if o == nil {
		o = &ThroughputOptions{}
	}

	rid, err := c.getRID(ctx)
	if err != nil {
		return ThroughputResponse{}, err
	}

	offers := &cosmosOffers{client: c.database.client}
	response, err := offers.ReplaceThroughputIfExists(ctx, throughputProperties, rid, o)
	return response, err
}

// CreateItem creates an item in a Cosmos container.
// ctx - The context for the request.
// partitionKey - The partition key for the item.
// item - The item to create.
// o - Options for the operation.
func (c *ContainerClient) CreateItem(
	ctx context.Context,
	partitionKey PartitionKey,
	item []byte,
	o *ItemOptions) (ItemResponse, error) {
	var err error
	spanName, err := c.getSpanForItems(operationTypeCreate)
	if err != nil {
		return ItemResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	h := headerOptionsOverride{
		partitionKey: &partitionKey,
	}

	if o == nil {
		o = &ItemOptions{}
	} else {
		h.enableContentResponseOnWrite = &o.EnableContentResponseOnWrite
	}

	operationContext := pipelineRequestOptions{
		resourceType:          resourceTypeDocument,
		resourceAddress:       c.link,
		isWriteOperation:      true,
		headerOptionsOverride: &h}

	path, err := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, true)
	if err != nil {
		return ItemResponse{}, err
	}

	azResponse, err := c.database.client.sendPostRequest(
		path,
		ctx,
		item,
		operationContext,
		o,
		nil)
	if err != nil {
		return ItemResponse{}, err
	}

	response, err := newItemResponse(azResponse)
	return response, err
}

// UpsertItem creates or replaces an item in a Cosmos container.
// ctx - The context for the request.
// partitionKey - The partition key for the item.
// item - The item to upsert.
// o - Options for the operation.
func (c *ContainerClient) UpsertItem(
	ctx context.Context,
	partitionKey PartitionKey,
	item []byte,
	o *ItemOptions) (ItemResponse, error) {
	var err error
	spanName, err := c.getSpanForItems(operationTypeUpsert)
	if err != nil {
		return ItemResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	h := headerOptionsOverride{
		partitionKey: &partitionKey,
	}

	addHeader := func(r *policy.Request) {
		r.Raw().Header.Add(cosmosHeaderIsUpsert, "true")
	}

	if o == nil {
		o = &ItemOptions{}
	} else {
		h.enableContentResponseOnWrite = &o.EnableContentResponseOnWrite
	}

	operationContext := pipelineRequestOptions{
		resourceType:          resourceTypeDocument,
		resourceAddress:       c.link,
		isWriteOperation:      true,
		headerOptionsOverride: &h}

	path, err := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, true)
	if err != nil {
		return ItemResponse{}, err
	}

	azResponse, err := c.database.client.sendPostRequest(
		path,
		ctx,
		item,
		operationContext,
		o,
		addHeader)
	if err != nil {
		return ItemResponse{}, err
	}

	response, err := newItemResponse(azResponse)
	return response, err
}

// ReplaceItem replaces an item in a Cosmos container.
// ctx - The context for the request.
// partitionKey - The partition key of the item to replace.
// itemId - The id of the item to replace.
// item - The content to be used to replace.
// o - Options for the operation.
func (c *ContainerClient) ReplaceItem(
	ctx context.Context,
	partitionKey PartitionKey,
	itemId string,
	item []byte,
	o *ItemOptions) (ItemResponse, error) {
	var err error
	spanName, err := c.getSpanForItems(operationTypeReplace)
	if err != nil {
		return ItemResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	h := headerOptionsOverride{
		partitionKey: &partitionKey,
	}

	if o == nil {
		o = &ItemOptions{}
	} else {
		h.enableContentResponseOnWrite = &o.EnableContentResponseOnWrite
	}

	operationContext := pipelineRequestOptions{
		resourceType:          resourceTypeDocument,
		resourceAddress:       createLink(c.link, pathSegmentDocument, itemId),
		isWriteOperation:      true,
		headerOptionsOverride: &h}

	path, err := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, false)
	if err != nil {
		return ItemResponse{}, err
	}

	azResponse, err := c.database.client.sendPutRequest(
		path,
		ctx,
		item,
		operationContext,
		o,
		nil)
	if err != nil {
		return ItemResponse{}, err
	}

	response, err := newItemResponse(azResponse)
	return response, err
}

// ReadItem reads an item in a Cosmos container.
// ctx - The context for the request.
// partitionKey - The partition key for the item.
// itemId - The id of the item to read.
// o - Options for the operation.
func (c *ContainerClient) ReadItem(
	ctx context.Context,
	partitionKey PartitionKey,
	itemId string,
	o *ItemOptions) (ItemResponse, error) {
	var err error
	spanName, err := c.getSpanForItems(operationTypeRead)
	if err != nil {
		return ItemResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	h := headerOptionsOverride{
		partitionKey: &partitionKey,
	}

	if o == nil {
		o = &ItemOptions{}
	}

	operationContext := pipelineRequestOptions{
		resourceType:          resourceTypeDocument,
		resourceAddress:       createLink(c.link, pathSegmentDocument, itemId),
		headerOptionsOverride: &h}

	path, err := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, false)
	if err != nil {
		return ItemResponse{}, err
	}

	azResponse, err := c.database.client.sendGetRequest(
		path,
		ctx,
		operationContext,
		o,
		nil)
	if err != nil {
		return ItemResponse{}, err
	}

	response, err := newItemResponse(azResponse)
	return response, err
}

// DeleteItem deletes an item in a Cosmos container.
// ctx - The context for the request.
// partitionKey - The partition key for the item.
// itemId - The id of the item to delete.
// o - Options for the operation.
func (c *ContainerClient) DeleteItem(
	ctx context.Context,
	partitionKey PartitionKey,
	itemId string,
	o *ItemOptions) (ItemResponse, error) {
	var err error
	spanName, err := c.getSpanForItems(operationTypeDelete)
	if err != nil {
		return ItemResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	h := headerOptionsOverride{
		partitionKey: &partitionKey,
	}

	if o == nil {
		o = &ItemOptions{}
	} else {
		h.enableContentResponseOnWrite = &o.EnableContentResponseOnWrite
	}

	operationContext := pipelineRequestOptions{
		resourceType:          resourceTypeDocument,
		resourceAddress:       createLink(c.link, pathSegmentDocument, itemId),
		isWriteOperation:      true,
		headerOptionsOverride: &h}

	path, err := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, false)
	if err != nil {
		return ItemResponse{}, err
	}

	azResponse, err := c.database.client.sendDeleteRequest(
		path,
		ctx,
		operationContext,
		o,
		nil)
	if err != nil {
		return ItemResponse{}, err
	}

	response, err := newItemResponse(azResponse)
	return response, err
}

// NewQueryItemsPager executes a single partition query in a Cosmos container.
// query - The SQL query to execute.
// partitionKey - The partition key to scope the query on. See below for more information on cross partition queries.
// o - Options for the operation.
//
// You can specify an empty list of partition keys by passing `NewPartitionKey()` to the `partitionKey` parameter, to indicate that the query WHERE clauses will specify which partitions to query.
//
// Limited cross partition queries ARE possible with the Go SDK.
// If you specify partition keys in the `partitionKey` parameter, you must specify ALL partition keys that the container has (in the case of hierarchical partitioning).
//
// If the query itself contains WHERE clauses that filter down to a single partition, the query will be executed on that partition.
// If the query does not filter down to a single partition (i.e. it does not filter on partition key at all, or filters on only some of the partition keys a container defines), the query will be executed as a cross partition query.
// The Azure Cosmos DB Gateway API, used by the Go SDK, can only perform a LIMITED set of cross-partition queries.
// Specifically, the gateway can only perform simple projections and filtering on cross partition queries.
// See https://learn.microsoft.com/rest/api/cosmos-db/querying-cosmosdb-resources-using-the-rest-api#queries-that-cannot-be-served-by-gateway for more details.
//
// When performing a cross-partition query, the Gateway may return pages of inconsistent size, or even empty pages (while still having a non-nil continuation token).
// Ensure you fully iterate the pager, even if you receive empty pages, to ensure you get all results.
//
// If you provide a query that the gateway cannot execute, it will return a BadRequest error.
func (c *ContainerClient) NewQueryItemsPager(query string, partitionKey PartitionKey, o *QueryOptions) *runtime.Pager[QueryItemsResponse] {
	correlatedActivityId, _ := uuid.New()
	h := headerOptionsOverride{
		partitionKey:         &partitionKey,
		correlatedActivityId: &correlatedActivityId,
	}

	queryOptions := &QueryOptions{}
	if o != nil {
		originalOptions := *o
		queryOptions = &originalOptions
	}

	operationContext := pipelineRequestOptions{
		resourceType:          resourceTypeDocument,
		resourceAddress:       c.link,
		headerOptionsOverride: &h,
	}

	if c.database.client.queryEngine != nil {
		return c.executeQueryWithEngine(c.database.client.queryEngine, query, partitionKey, o, operationContext)
	}

	path, _ := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, true)

	return runtime.NewPager(runtime.PagingHandler[QueryItemsResponse]{
		More: func(page QueryItemsResponse) bool {
			return page.ContinuationToken != nil
		},
		Fetcher: func(ctx context.Context, page *QueryItemsResponse) (QueryItemsResponse, error) {
			var err error
			spanName, err := c.getSpanForItems(operationTypeQuery)
			if err != nil {
				return QueryItemsResponse{}, err
			}
			ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
			defer func() { endSpan(err) }()
			if page != nil {
				if page.ContinuationToken != nil {
					// Use the previous page continuation if available
					queryOptions.ContinuationToken = page.ContinuationToken
				}
			}

			azResponse, err := c.database.client.sendQueryRequest(
				path,
				ctx,
				query,
				queryOptions.QueryParameters,
				operationContext,
				queryOptions,
				nil)

			if err != nil {
				return QueryItemsResponse{}, err
			}

			return newQueryResponse(azResponse)
		},
	})
}

// PatchItem patches an item in a Cosmos container.
// ctx - The context for the request.
// partitionKey - The partition key for the item.
// itemId - The id of the item to patch.
// ops - Operations to perform on the patch
// o - Options for the operation.
func (c *ContainerClient) PatchItem(
	ctx context.Context,
	partitionKey PartitionKey,
	itemId string,
	ops PatchOperations,
	o *ItemOptions) (ItemResponse, error) {
	var err error
	spanName, err := c.getSpanForItems(operationTypePatch)
	if err != nil {
		return ItemResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	h := headerOptionsOverride{
		partitionKey: &partitionKey,
	}

	if o == nil {
		o = &ItemOptions{}
	} else {
		h.enableContentResponseOnWrite = &o.EnableContentResponseOnWrite
	}

	operationContext := pipelineRequestOptions{
		resourceType:          resourceTypeDocument,
		resourceAddress:       createLink(c.link, pathSegmentDocument, itemId),
		isWriteOperation:      true,
		headerOptionsOverride: &h}

	path, err := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, false)
	if err != nil {
		return ItemResponse{}, err
	}

	azResponse, err := c.database.client.sendPatchRequest(
		path,
		ctx,
		ops,
		operationContext,
		o,
		nil)
	if err != nil {
		return ItemResponse{}, err
	}

	response, err := newItemResponse(azResponse)
	return response, err
}

// NewTransactionalBatch creates a batch of operations to be committed as a single unit.
// See https://docs.microsoft.com/azure/cosmos-db/sql/transactional-batch
func (c *ContainerClient) NewTransactionalBatch(partitionKey PartitionKey) TransactionalBatch {
	return TransactionalBatch{partitionKey: partitionKey}
}

// ExecuteTransactionalBatch executes a transactional batch.
// Once executed, verify the Success property of the response to determine if the batch was committed
func (c *ContainerClient) ExecuteTransactionalBatch(ctx context.Context, b TransactionalBatch, o *TransactionalBatchOptions) (TransactionalBatchResponse, error) {
	var err error
	spanName, err := c.getSpanForContainer(operationTypeBatch, resourceTypeCollection, c.id)
	if err != nil {
		return TransactionalBatchResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.database.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if len(b.operations) == 0 {
		return TransactionalBatchResponse{}, errors.New("no operations in batch")
	}

	h := headerOptionsOverride{
		partitionKey: &b.partitionKey,
	}

	if o == nil {
		o = &TransactionalBatchOptions{}
	} else {
		h.enableContentResponseOnWrite = &o.EnableContentResponseOnWrite
	}

	// If contentResponseOnWrite is not enabled at the client level the
	// service will not even send a batch response payload
	// Instead we should automatically enforce contentResponseOnWrite for all
	// batch requests whenever at least one of the item operations requires a content response (read operation)
	enableContentResponseOnWriteForReadOperations := true
	for _, op := range b.operations {
		if op.getOperationType() == operationTypeRead {
			h.enableContentResponseOnWrite = &enableContentResponseOnWriteForReadOperations
			break
		}
	}

	operationContext := pipelineRequestOptions{
		resourceType:          resourceTypeDocument,
		resourceAddress:       c.link,
		isWriteOperation:      true,
		headerOptionsOverride: &h}

	path, err := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, true)
	if err != nil {
		return TransactionalBatchResponse{}, err
	}

	azResponse, err := c.database.client.sendBatchRequest(
		ctx,
		path,
		b.operations,
		operationContext,
		o,
		nil)
	if err != nil {
		return TransactionalBatchResponse{}, err
	}

	response, err := newTransactionalBatchResponse(azResponse)
	return response, err
}

func (c *ContainerClient) getRID(ctx context.Context) (string, error) {
	containerResponse, err := c.Read(ctx, nil)
	if err != nil {
		return "", err
	}

	return containerResponse.ContainerProperties.ResourceID, nil
}

func (c *ContainerClient) getSpanForContainer(operationType operationType, resourceType resourceType, id string) (span, error) {
	return getSpanNameForContainers(c.database.client.accountEndpointUrl(), operationType, resourceType, c.database.id, id)
}

func (c *ContainerClient) getSpanForItems(operationType operationType) (span, error) {
	return getSpanNameForItems(c.database.client.accountEndpointUrl(), operationType, c.database.id, c.id)
}

func (c *ContainerClient) getQueryPlanFromGateway(ctx context.Context, query string, queryOptions *QueryOptions, operationContext pipelineRequestOptions) ([]byte, error) {
	path, _ := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, true)
	azResponse, err := c.database.client.sendQueryRequest(
		path,
		ctx,
		query,
		queryOptions.QueryParameters,
		operationContext,
		queryOptions,
		func(req *policy.Request) {
			req.Raw().Header.Set(cosmosHeaderIsQueryPlanRequest, "True")
		})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(azResponse.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *ContainerClient) getPartitionKeyRanges(ctx context.Context, queryOptions *QueryOptions, operationContext pipelineRequestOptions) ([]byte, error) {
	path, _ := generatePathForNameBased(resourceTypePartitionKeyRange, operationContext.resourceAddress, true)
	azResponse, err := c.database.client.sendGetRequest(
		path,
		ctx,
		pipelineRequestOptions{
			resourceType:          resourceTypePartitionKeyRange,
			resourceAddress:       operationContext.resourceAddress,
			headerOptionsOverride: operationContext.headerOptionsOverride,
		},
		nil,
		nil)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(azResponse.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Executes a query using the provided query engine.
func (c *ContainerClient) executeQueryWithEngine(ctx context.Context, queryEngine QueryEngine, query string, partitionKey PartitionKey, queryOptions *QueryOptions, operationContext pipelineRequestOptions) *runtime.Pager[QueryItemsResponse] {
	var queryPipeline QueryPipeline
	path, _ := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, true)
	return runtime.NewPager(runtime.PagingHandler[QueryItemsResponse]{
		More: func(page QueryItemsResponse) bool {
			return queryPipeline == nil || !queryPipeline.IsComplete()
		},
		Fetcher: func(ctx context.Context, page *QueryItemsResponse) (QueryItemsResponse, error) {
			if queryPipeline == nil {
				// First page, we need to fetch the query plan and PK ranges
				// TODO: We could proactively try to run this against the gateway and then fall back to the engine.
				plan, err := c.getQueryPlanFromGateway(ctx, query, queryOptions, operationContext)
				if err != nil {
					return QueryItemsResponse{}, err
				}
				pkranges, err := c.getPartitionKeyRanges()
				if err != nil {
					return QueryItemsResponse{}, err
				}

				// Create a query pipeline
				queryPipeline, err = queryEngine.CreateQueryPipeline(plan, pkranges)
				if err != nil {
					return QueryItemsResponse{}, err
				}
				query = queryPipeline.GetRewrittenQuery()
			}

			for {
				if queryPipeline.IsComplete() {
					// Nil out the query pipeline, which will cause More to stop returning true.
					queryPipeline = nil
					return QueryItemsResponse{
						Response: Response{}, // TODO: Synthesize something here?
						Items:    nil,
					}, nil
				}
				// Fetch more data from the pipeline
				items, err := queryPipeline.NextBatch(queryOptions.PageSizeHint)
				if err != nil {
					return QueryItemsResponse{}, nil
				}

				if len(items) > 0 {
					return QueryItemsResponse{
						Response: Response{}, // TODO: Synthesize something here?
						Items:    items,
					}, nil
				}

				// We didn't get any items, so we need to put more data in.
				requests, err := queryPipeline.NextRequests()
				for _, request := range requests {
					azResponse, err := c.database.client.sendQueryRequest(
						path,
						ctx,
						query,
						queryOptions.QueryParameters,
						operationContext,
						&request,
						nil)
					buf := new(bytes.Buffer)
					_, err = buf.ReadFrom(azResponse.Body)
					if err != nil {
						return QueryItemsResponse{}, err
					}
					data := buf.Bytes()
					continuation := azResponse.Header.Get(cosmosHeaderContinuationToken)

					if err = queryPipeline.ProvideData(request.PartitionKeyRangeID, data, continuation); err != nil {
						return QueryItemsResponse{}, err
					}
				}

				// We've provided more data to the pipeline, so let's try to run the pipeline forward again.
			}
		},
	})
}
