// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// DatabaseClient lets you perform read, update, change throughput, and delete database operations.
type DatabaseClient struct {
	// The Id of the Cosmos database
	id string
	// The client associated with the Cosmos database
	client *Client
	// The resource link
	link string
}

func newDatabase(id string, client *Client) (*DatabaseClient, error) {
	return &DatabaseClient{
		id:     id,
		client: client,
		link:   createLink("", pathSegmentDatabase, id)}, nil
}

// ID returns the identifier of the Cosmos database.
func (db *DatabaseClient) ID() string {
	return db.id
}

// NewContainer returns a struct that represents the container and allows container level operations.
// id - The id of the container.
func (db *DatabaseClient) NewContainer(id string) (*ContainerClient, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	return newContainer(id, db)
}

// CreateContainer creates a container in the Cosmos database.
// ctx - The context for the request.
// containerProperties - The properties for the container.
// o - Options for the create container operation.
func (db *DatabaseClient) CreateContainer(
	ctx context.Context,
	containerProperties ContainerProperties,
	o *CreateContainerOptions) (ContainerResponse, error) {
	var err error
	spanName, err := db.getSpanForContainers(operationTypeCreate, resourceTypeCollection, containerProperties.ID)
	if err != nil {
		return ContainerResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, db.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if o == nil {
		o = &CreateContainerOptions{}
	}
	returnResponse := true
	h := &headerOptionsOverride{
		enableContentResponseOnWrite: &returnResponse,
	}

	operationContext := pipelineRequestOptions{
		resourceType:          resourceTypeCollection,
		resourceAddress:       db.link,
		isWriteOperation:      true,
		headerOptionsOverride: h,
	}

	path, err := generatePathForNameBased(resourceTypeCollection, db.link, true)
	if err != nil {
		return ContainerResponse{}, err
	}

	azResponse, err := db.client.sendPostRequest(
		path,
		ctx,
		containerProperties,
		operationContext,
		nil,
		o.ThroughputProperties.addHeadersToRequest)
	if err != nil {
		return ContainerResponse{}, err
	}

	response, err := newContainerResponse(azResponse)
	return response, err
}

// NewQueryContainersPager executes query for containers within a database.
// query - The SQL query to execute.
// o - Options for the operation.
func (c *DatabaseClient) NewQueryContainersPager(query string, o *QueryContainersOptions) *runtime.Pager[QueryContainersResponse] {
	queryOptions := &QueryContainersOptions{}
	if o != nil {
		originalOptions := *o
		queryOptions = &originalOptions
	}

	operationContext := pipelineRequestOptions{
		resourceType:    resourceTypeCollection,
		resourceAddress: c.link,
	}

	path, _ := generatePathForNameBased(resourceTypeCollection, operationContext.resourceAddress, true)

	return runtime.NewPager(runtime.PagingHandler[QueryContainersResponse]{
		More: func(page QueryContainersResponse) bool {
			return page.ContinuationToken != nil
		},
		Fetcher: func(ctx context.Context, page *QueryContainersResponse) (QueryContainersResponse, error) {
			var err error
			spanName, err := c.getSpanForDatabases(operationTypeQuery, resourceTypeCollection)
			if err != nil {
				return QueryContainersResponse{}, err
			}
			ctx, endSpan := runtime.StartSpan(ctx, spanName.name, c.client.internal.Tracer(), &spanName.options)
			defer func() { endSpan(err) }()
			if page != nil {
				if page.ContinuationToken != nil {
					// Use the previous page continuation if available
					queryOptions.ContinuationToken = page.ContinuationToken
				}
			}

			azResponse, err := c.client.sendQueryRequest(
				path,
				ctx,
				query,
				queryOptions.QueryParameters,
				operationContext,
				queryOptions,
				nil)

			if err != nil {
				return QueryContainersResponse{}, err
			}

			return newContainersQueryResponse(azResponse)
		},
	})
}

// Read obtains the information for a Cosmos database.
// ctx - The context for the request.
// o - Options for Read operation.
func (db *DatabaseClient) Read(
	ctx context.Context,
	o *ReadDatabaseOptions) (DatabaseResponse, error) {
	var err error
	spanName, err := db.getSpanForDatabases(operationTypeRead, resourceTypeDatabase)
	if err != nil {
		return DatabaseResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, db.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if o == nil {
		o = &ReadDatabaseOptions{}
	}

	operationContext := pipelineRequestOptions{
		resourceType:    resourceTypeDatabase,
		resourceAddress: db.link,
	}

	path, err := generatePathForNameBased(resourceTypeDatabase, db.link, false)
	if err != nil {
		return DatabaseResponse{}, err
	}

	azResponse, err := db.client.sendGetRequest(
		path,
		ctx,
		operationContext,
		o,
		nil)
	if err != nil {
		return DatabaseResponse{}, err
	}

	response, err := newDatabaseResponse(azResponse)
	return response, err
}

// ReadThroughput obtains the provisioned throughput information for the database.
// ctx - The context for the request.
// o - Options for the operation.
func (db *DatabaseClient) ReadThroughput(
	ctx context.Context,
	o *ThroughputOptions) (ThroughputResponse, error) {
	var err error
	spanName, err := db.getSpanForDatabases(operationTypeRead, resourceTypeOffer)
	if err != nil {
		return ThroughputResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, db.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if o == nil {
		o = &ThroughputOptions{}
	}

	rid, err := db.getRID(ctx)
	if err != nil {
		return ThroughputResponse{}, err
	}

	offers := &cosmosOffers{client: db.client}
	response, err := offers.ReadThroughputIfExists(ctx, rid, o)
	return response, err
}

// ReplaceThroughput updates the provisioned throughput for the database.
// ctx - The context for the request.
// throughputProperties - The throughput configuration of the database.
// o - Options for the operation.
func (db *DatabaseClient) ReplaceThroughput(
	ctx context.Context,
	throughputProperties ThroughputProperties,
	o *ThroughputOptions) (ThroughputResponse, error) {
	var err error
	spanName, err := db.getSpanForDatabases(operationTypeReplace, resourceTypeOffer)
	if err != nil {
		return ThroughputResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, db.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if o == nil {
		o = &ThroughputOptions{}
	}

	rid, err := db.getRID(ctx)
	if err != nil {
		return ThroughputResponse{}, err
	}

	offers := &cosmosOffers{client: db.client}
	response, err := offers.ReplaceThroughputIfExists(ctx, throughputProperties, rid, o)
	return response, err
}

// Delete a Cosmos database.
// ctx - The context for the request.
// o - Options for Read operation.
func (db *DatabaseClient) Delete(
	ctx context.Context,
	o *DeleteDatabaseOptions) (DatabaseResponse, error) {
	var err error
	spanName, err := db.getSpanForDatabases(operationTypeDelete, resourceTypeDatabase)
	if err != nil {
		return DatabaseResponse{}, err
	}
	ctx, endSpan := runtime.StartSpan(ctx, spanName.name, db.client.internal.Tracer(), &spanName.options)
	defer func() { endSpan(err) }()
	if o == nil {
		o = &DeleteDatabaseOptions{}
	}

	operationContext := pipelineRequestOptions{
		resourceType:     resourceTypeDatabase,
		resourceAddress:  db.link,
		isWriteOperation: true,
	}

	path, err := generatePathForNameBased(resourceTypeDatabase, db.link, false)
	if err != nil {
		return DatabaseResponse{}, err
	}

	azResponse, err := db.client.sendDeleteRequest(
		path,
		ctx,
		operationContext,
		o,
		nil)
	if err != nil {
		return DatabaseResponse{}, err
	}

	response, err := newDatabaseResponse(azResponse)
	return response, err
}

func (db *DatabaseClient) getRID(ctx context.Context) (string, error) {
	dbResponse, err := db.Read(ctx, nil)
	if err != nil {
		return "", err
	}

	return dbResponse.DatabaseProperties.ResourceID, nil
}

func (db *DatabaseClient) getSpanForDatabases(operationType operationType, resourceType resourceType) (span, error) {
	return getSpanNameForDatabases(db.client.accountEndpointUrl(), operationType, resourceType, db.id)
}

func (db *DatabaseClient) getSpanForContainers(operationType operationType, resourceType resourceType, id string) (span, error) {
	return getSpanNameForContainers(db.client.accountEndpointUrl(), operationType, resourceType, db.id, id)
}
