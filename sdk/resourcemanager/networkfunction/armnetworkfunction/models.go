//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkfunction

import "time"

// AzureTrafficCollector - Azure Traffic Collector resource.
type AzureTrafficCollector struct {
	// Resource location.
	Location *string `json:"location,omitempty"`

	// Properties of the Azure Traffic Collector.
	Properties *AzureTrafficCollectorPropertiesFormat `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *TrackedResourceSystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AzureTrafficCollectorListResult - Response for the ListTrafficCollectors API service call.
type AzureTrafficCollectorListResult struct {
	// A list of Traffic Collector resources.
	Value []*AzureTrafficCollector `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// AzureTrafficCollectorPropertiesFormat - Azure Traffic Collector resource properties.
type AzureTrafficCollectorPropertiesFormat struct {
	// Collector Policies for Azure Traffic Collector.
	CollectorPolicies []*CollectorPolicy `json:"collectorPolicies,omitempty"`

	// The virtualHub to which the Azure Traffic Collector belongs.
	VirtualHub *ResourceReference `json:"virtualHub,omitempty"`

	// READ-ONLY; The provisioning state of the application rule collection resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// AzureTrafficCollectorsByResourceGroupClientListOptions contains the optional parameters for the AzureTrafficCollectorsByResourceGroupClient.List
// method.
type AzureTrafficCollectorsByResourceGroupClientListOptions struct {
	// placeholder for future optional parameters
}

// AzureTrafficCollectorsBySubscriptionClientListOptions contains the optional parameters for the AzureTrafficCollectorsBySubscriptionClient.List
// method.
type AzureTrafficCollectorsBySubscriptionClientListOptions struct {
	// placeholder for future optional parameters
}

// AzureTrafficCollectorsClientBeginCreateOrUpdateOptions contains the optional parameters for the AzureTrafficCollectorsClient.BeginCreateOrUpdate
// method.
type AzureTrafficCollectorsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureTrafficCollectorsClientBeginDeleteOptions contains the optional parameters for the AzureTrafficCollectorsClient.BeginDelete
// method.
type AzureTrafficCollectorsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureTrafficCollectorsClientGetOptions contains the optional parameters for the AzureTrafficCollectorsClient.Get method.
type AzureTrafficCollectorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureTrafficCollectorsClientUpdateTagsOptions contains the optional parameters for the AzureTrafficCollectorsClient.UpdateTags
// method.
type AzureTrafficCollectorsClientUpdateTagsOptions struct {
	// placeholder for future optional parameters
}

// ClientListOperationsOptions contains the optional parameters for the Client.ListOperations method.
type ClientListOperationsOptions struct {
	// placeholder for future optional parameters
}

// CloudError - An error response from the service.
type CloudError struct {
	// An error response from the service.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody - An error response from the service.
type CloudErrorBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`

	// A list of additional details about the error.
	Details []*CloudErrorBody `json:"details,omitempty"`

	// A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`

	// The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
}

// CollectorPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the CollectorPoliciesClient.BeginCreateOrUpdate
// method.
type CollectorPoliciesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CollectorPoliciesClientBeginDeleteOptions contains the optional parameters for the CollectorPoliciesClient.BeginDelete
// method.
type CollectorPoliciesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CollectorPoliciesClientGetOptions contains the optional parameters for the CollectorPoliciesClient.Get method.
type CollectorPoliciesClientGetOptions struct {
	// placeholder for future optional parameters
}

// CollectorPoliciesClientListOptions contains the optional parameters for the CollectorPoliciesClient.List method.
type CollectorPoliciesClientListOptions struct {
	// placeholder for future optional parameters
}

// CollectorPolicy - Collector policy resource.
type CollectorPolicy struct {
	// Properties of the Collector Policy.
	Properties *CollectorPolicyPropertiesFormat `json:"properties,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *CollectorPolicySystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CollectorPolicyListResult - Response for the ListCollectorPolicies API service call.
type CollectorPolicyListResult struct {
	// A list of collection policies.
	Value []*CollectorPolicy `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// CollectorPolicyPropertiesFormat - Collection policy properties.
type CollectorPolicyPropertiesFormat struct {
	// Emission policies.
	EmissionPolicies []*EmissionPoliciesPropertiesFormat `json:"emissionPolicies,omitempty"`

	// Ingestion policies.
	IngestionPolicy *IngestionPolicyPropertiesFormat `json:"ingestionPolicy,omitempty"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// CollectorPolicySystemData - Metadata pertaining to creation and last modification of the resource.
type CollectorPolicySystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// EmissionPoliciesPropertiesFormat - Emission policy properties.
type EmissionPoliciesPropertiesFormat struct {
	// Emission policy destinations.
	EmissionDestinations []*EmissionPolicyDestination `json:"emissionDestinations,omitempty"`

	// Emission format type.
	EmissionType *EmissionType `json:"emissionType,omitempty"`
}

// EmissionPolicyDestination - Emission policy destination properties.
type EmissionPolicyDestination struct {
	// Emission destination type.
	DestinationType *DestinationType `json:"destinationType,omitempty"`
}

// IngestionPolicyPropertiesFormat - Ingestion Policy properties.
type IngestionPolicyPropertiesFormat struct {
	// Ingestion Sources.
	IngestionSources []*IngestionSourcesPropertiesFormat `json:"ingestionSources,omitempty"`

	// The ingestion type.
	IngestionType *IngestionType `json:"ingestionType,omitempty"`
}

// IngestionSourcesPropertiesFormat - Ingestion policy properties.
type IngestionSourcesPropertiesFormat struct {
	// Resource ID.
	ResourceID *string `json:"resourceId,omitempty"`

	// Ingestion source type.
	SourceType *SourceType `json:"sourceType,omitempty"`
}

// Operation - Azure Traffic Collector REST API operation definition.
type Operation struct {
	// Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`

	// Origin of the operation
	Origin *string `json:"origin,omitempty"`
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// Description of the operation.
	Description *string `json:"description,omitempty"`

	// Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft NetworkFunction.
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of the request to list Azure Traffic Collector operations. It contains a list of operations
// and a URL link to get the next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of operations supported by the Azure Traffic Collector resource provider.
	Value []*Operation `json:"value,omitempty"`
}

// ProxyResource - An azure resource object
type ProxyResource struct {
	// READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceReference - Resource reference properties.
type ResourceReference struct {
	// READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TagsObject - Tags object for patch operations.
type TagsObject struct {
	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// TrackedResource - Common resource representation.
type TrackedResource struct {
	// Resource location.
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *TrackedResourceSystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TrackedResourceSystemData - Metadata pertaining to creation and last modification of the resource.
type TrackedResourceSystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}