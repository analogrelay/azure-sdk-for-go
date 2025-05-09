// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcarbonoptimization

// CarbonEmissionDataClassification provides polymorphic access to related types.
// Call the interface's GetCarbonEmissionData() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CarbonEmissionData, *CarbonEmissionItemDetailData, *CarbonEmissionMonthlySummaryData, *CarbonEmissionOverallSummaryData,
// - *CarbonEmissionTopItemMonthlySummaryData, *CarbonEmissionTopItemsSummaryData, *ResourceCarbonEmissionItemDetailData,
// - *ResourceCarbonEmissionTopItemMonthlySummaryData, *ResourceCarbonEmissionTopItemsSummaryData, *ResourceGroupCarbonEmissionItemDetailData,
// - *ResourceGroupCarbonEmissionTopItemMonthlySummaryData, *ResourceGroupCarbonEmissionTopItemsSummaryData
type CarbonEmissionDataClassification interface {
	// GetCarbonEmissionData returns the CarbonEmissionData content of the underlying type.
	GetCarbonEmissionData() *CarbonEmissionData
}

// QueryFilterClassification provides polymorphic access to related types.
// Call the interface's GetQueryFilter() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ItemDetailsQueryFilter, *MonthlySummaryReportQueryFilter, *OverallSummaryReportQueryFilter, *QueryFilter, *TopItemsMonthlySummaryReportQueryFilter,
// - *TopItemsSummaryReportQueryFilter
type QueryFilterClassification interface {
	// GetQueryFilter returns the QueryFilter content of the underlying type.
	GetQueryFilter() *QueryFilter
}
