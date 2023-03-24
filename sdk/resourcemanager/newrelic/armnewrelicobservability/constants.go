//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnewrelicobservability

const (
	moduleName    = "armnewrelicobservability"
	moduleVersion = "v0.1.0"
)

// AccountCreationSource - Source of Account creation
type AccountCreationSource string

const (
	// AccountCreationSourceLIFTR - Account is created from LIFTR
	AccountCreationSourceLIFTR AccountCreationSource = "LIFTR"
	// AccountCreationSourceNEWRELIC - Account is created from NEWRELIC
	AccountCreationSourceNEWRELIC AccountCreationSource = "NEWRELIC"
)

// PossibleAccountCreationSourceValues returns the possible values for the AccountCreationSource const type.
func PossibleAccountCreationSourceValues() []AccountCreationSource {
	return []AccountCreationSource{
		AccountCreationSourceLIFTR,
		AccountCreationSourceNEWRELIC,
	}
}

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// BillingCycle - Different usage type like YEARLY/MONTHLY
type BillingCycle string

const (
	// BillingCycleMONTHLY - Billing cycle is MONTHLY
	BillingCycleMONTHLY BillingCycle = "MONTHLY"
	// BillingCycleWEEKLY - Billing cycle is WEEKLY
	BillingCycleWEEKLY BillingCycle = "WEEKLY"
	// BillingCycleYEARLY - Billing cycle is YEARLY
	BillingCycleYEARLY BillingCycle = "YEARLY"
)

// PossibleBillingCycleValues returns the possible values for the BillingCycle const type.
func PossibleBillingCycleValues() []BillingCycle {
	return []BillingCycle{
		BillingCycleMONTHLY,
		BillingCycleWEEKLY,
		BillingCycleYEARLY,
	}
}

// BillingSource - Billing source
type BillingSource string

const (
	// BillingSourceAZURE - Billing source is Azure
	BillingSourceAZURE    BillingSource = "AZURE"
	BillingSourceNEWRELIC BillingSource = "NEWRELIC"
)

// PossibleBillingSourceValues returns the possible values for the BillingSource const type.
func PossibleBillingSourceValues() []BillingSource {
	return []BillingSource{
		BillingSourceAZURE,
		BillingSourceNEWRELIC,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// LiftrResourceCategories - Liftr Resource category.
type LiftrResourceCategories string

const (
	LiftrResourceCategoriesMonitorLogs LiftrResourceCategories = "MonitorLogs"
	LiftrResourceCategoriesUnknown     LiftrResourceCategories = "Unknown"
)

// PossibleLiftrResourceCategoriesValues returns the possible values for the LiftrResourceCategories const type.
func PossibleLiftrResourceCategoriesValues() []LiftrResourceCategories {
	return []LiftrResourceCategories{
		LiftrResourceCategoriesMonitorLogs,
		LiftrResourceCategoriesUnknown,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// MarketplaceSubscriptionStatus - Flag specifying the Marketplace Subscription Status of the resource. If payment is not
// made in time, the resource will go in Suspended state.
type MarketplaceSubscriptionStatus string

const (
	// MarketplaceSubscriptionStatusActive - monitoring is enabled
	MarketplaceSubscriptionStatusActive MarketplaceSubscriptionStatus = "Active"
	// MarketplaceSubscriptionStatusSuspended - monitoring is disabled
	MarketplaceSubscriptionStatusSuspended MarketplaceSubscriptionStatus = "Suspended"
)

// PossibleMarketplaceSubscriptionStatusValues returns the possible values for the MarketplaceSubscriptionStatus const type.
func PossibleMarketplaceSubscriptionStatusValues() []MarketplaceSubscriptionStatus {
	return []MarketplaceSubscriptionStatus{
		MarketplaceSubscriptionStatusActive,
		MarketplaceSubscriptionStatusSuspended,
	}
}

// MonitoringStatus - Flag specifying if the resource monitoring is enabled or disabled.
type MonitoringStatus string

const (
	// MonitoringStatusDisabled - monitoring is disabled
	MonitoringStatusDisabled MonitoringStatus = "Disabled"
	// MonitoringStatusEnabled - monitoring is enabled
	MonitoringStatusEnabled MonitoringStatus = "Enabled"
)

// PossibleMonitoringStatusValues returns the possible values for the MonitoringStatus const type.
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return []MonitoringStatus{
		MonitoringStatusDisabled,
		MonitoringStatusEnabled,
	}
}

// OrgCreationSource - Source of Org creation
type OrgCreationSource string

const (
	// OrgCreationSourceLIFTR - Org is created from LIFTR
	OrgCreationSourceLIFTR OrgCreationSource = "LIFTR"
	// OrgCreationSourceNEWRELIC - Org is created from NEWRELIC
	OrgCreationSourceNEWRELIC OrgCreationSource = "NEWRELIC"
)

// PossibleOrgCreationSourceValues returns the possible values for the OrgCreationSource const type.
func PossibleOrgCreationSourceValues() []OrgCreationSource {
	return []OrgCreationSource{
		OrgCreationSourceLIFTR,
		OrgCreationSourceNEWRELIC,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - Provisioning State of the Monitor resource
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Monitor resource creation request accepted
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Monitor resource creation canceled
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating - Monitor resource creation started
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleted - Monitor resource is deleted
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting - Monitor resource deletion started
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Monitor resource creation failed
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateNotSpecified - Monitor resource state is unknown
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	// ProvisioningStateSucceeded - Monitor resource creation successful
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Monitor resource is being updated
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateNotSpecified,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// SendAADLogsStatus - Indicates whether AAD logs are being sent.
type SendAADLogsStatus string

const (
	SendAADLogsStatusDisabled SendAADLogsStatus = "Disabled"
	SendAADLogsStatusEnabled  SendAADLogsStatus = "Enabled"
)

// PossibleSendAADLogsStatusValues returns the possible values for the SendAADLogsStatus const type.
func PossibleSendAADLogsStatusValues() []SendAADLogsStatus {
	return []SendAADLogsStatus{
		SendAADLogsStatusDisabled,
		SendAADLogsStatusEnabled,
	}
}

// SendActivityLogsStatus - Indicates whether activity logs are being sent.
type SendActivityLogsStatus string

const (
	SendActivityLogsStatusDisabled SendActivityLogsStatus = "Disabled"
	SendActivityLogsStatusEnabled  SendActivityLogsStatus = "Enabled"
)

// PossibleSendActivityLogsStatusValues returns the possible values for the SendActivityLogsStatus const type.
func PossibleSendActivityLogsStatusValues() []SendActivityLogsStatus {
	return []SendActivityLogsStatus{
		SendActivityLogsStatusDisabled,
		SendActivityLogsStatusEnabled,
	}
}

// SendMetricsStatus - Indicates whether metrics are being sent.
type SendMetricsStatus string

const (
	SendMetricsStatusDisabled SendMetricsStatus = "Disabled"
	SendMetricsStatusEnabled  SendMetricsStatus = "Enabled"
)

// PossibleSendMetricsStatusValues returns the possible values for the SendMetricsStatus const type.
func PossibleSendMetricsStatusValues() []SendMetricsStatus {
	return []SendMetricsStatus{
		SendMetricsStatusDisabled,
		SendMetricsStatusEnabled,
	}
}

// SendSubscriptionLogsStatus - Indicates whether subscription logs are being sent.
type SendSubscriptionLogsStatus string

const (
	SendSubscriptionLogsStatusDisabled SendSubscriptionLogsStatus = "Disabled"
	SendSubscriptionLogsStatusEnabled  SendSubscriptionLogsStatus = "Enabled"
)

// PossibleSendSubscriptionLogsStatusValues returns the possible values for the SendSubscriptionLogsStatus const type.
func PossibleSendSubscriptionLogsStatusValues() []SendSubscriptionLogsStatus {
	return []SendSubscriptionLogsStatus{
		SendSubscriptionLogsStatusDisabled,
		SendSubscriptionLogsStatusEnabled,
	}
}

// SendingLogsStatus - Indicates whether logs are being sent.
type SendingLogsStatus string

const (
	SendingLogsStatusDisabled SendingLogsStatus = "Disabled"
	SendingLogsStatusEnabled  SendingLogsStatus = "Enabled"
)

// PossibleSendingLogsStatusValues returns the possible values for the SendingLogsStatus const type.
func PossibleSendingLogsStatusValues() []SendingLogsStatus {
	return []SendingLogsStatus{
		SendingLogsStatusDisabled,
		SendingLogsStatusEnabled,
	}
}

// SendingMetricsStatus - Indicates whether metrics are being sent.
type SendingMetricsStatus string

const (
	SendingMetricsStatusDisabled SendingMetricsStatus = "Disabled"
	SendingMetricsStatusEnabled  SendingMetricsStatus = "Enabled"
)

// PossibleSendingMetricsStatusValues returns the possible values for the SendingMetricsStatus const type.
func PossibleSendingMetricsStatusValues() []SendingMetricsStatus {
	return []SendingMetricsStatus{
		SendingMetricsStatusDisabled,
		SendingMetricsStatusEnabled,
	}
}

// SingleSignOnStates - Various states of the SSO resource
type SingleSignOnStates string

const (
	SingleSignOnStatesDisable  SingleSignOnStates = "Disable"
	SingleSignOnStatesEnable   SingleSignOnStates = "Enable"
	SingleSignOnStatesExisting SingleSignOnStates = "Existing"
	SingleSignOnStatesInitial  SingleSignOnStates = "Initial"
)

// PossibleSingleSignOnStatesValues returns the possible values for the SingleSignOnStates const type.
func PossibleSingleSignOnStatesValues() []SingleSignOnStates {
	return []SingleSignOnStates{
		SingleSignOnStatesDisable,
		SingleSignOnStatesEnable,
		SingleSignOnStatesExisting,
		SingleSignOnStatesInitial,
	}
}

// TagAction - Valid actions for a filtering tag. Exclusion takes priority over inclusion.
type TagAction string

const (
	TagActionExclude TagAction = "Exclude"
	TagActionInclude TagAction = "Include"
)

// PossibleTagActionValues returns the possible values for the TagAction const type.
func PossibleTagActionValues() []TagAction {
	return []TagAction{
		TagActionExclude,
		TagActionInclude,
	}
}

// UsageType - Different usage type like PAYG/COMMITTED
type UsageType string

const (
	// UsageTypeCOMMITTED - Usage type is COMMITTED
	UsageTypeCOMMITTED UsageType = "COMMITTED"
	// UsageTypePAYG - Usage type is PAYG
	UsageTypePAYG UsageType = "PAYG"
)

// PossibleUsageTypeValues returns the possible values for the UsageType const type.
func PossibleUsageTypeValues() []UsageType {
	return []UsageType{
		UsageTypeCOMMITTED,
		UsageTypePAYG,
	}
}
