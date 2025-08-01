# Release History

## 2.0.0 (2025-07-25)
### Breaking Changes

- Type of `ErrorResponse.Error` has been changed from `*ErrorResponseError` to `*ErrorDetail`
- Type of `ExtendedLocationOptions.SupportedPolicy` has been changed from `*string` to `*ResourceTypeExtendedLocationPolicy`
- Type of `ExtendedLocationOptions.Type` has been changed from `*string` to `*ExtendedLocationType`
- Type of `OperationsDefinition.ActionType` has been changed from `*OperationsDefinitionActionType` to `*OperationActionType`
- Type of `OperationsDefinition.Origin` has been changed from `*OperationsDefinitionOrigin` to `*OperationOrigins`
- Type of `ResourceProviderManagement.ResourceAccessPolicy` has been changed from `*ResourceProviderManagementResourceAccessPolicy` to `*ResourceAccessPolicy`
- Type of `ResourceProviderManagement.ResourceAccessRoles` has been changed from `[]any` to `[]*ResourceAccessRole`
- Type of `ResourceProviderManifestManagement.ResourceAccessPolicy` has been changed from `*ResourceProviderManagementResourceAccessPolicy` to `*ResourceAccessPolicy`
- Type of `ResourceProviderManifestManagement.ResourceAccessRoles` has been changed from `[]any` to `[]*ResourceAccessRole`
- Type of `ResourceProviderManifestPropertiesManagement.ResourceAccessPolicy` has been changed from `*ResourceProviderManagementResourceAccessPolicy` to `*ResourceAccessPolicy`
- Type of `ResourceProviderManifestPropertiesManagement.ResourceAccessRoles` has been changed from `[]any` to `[]*ResourceAccessRole`
- Type of `ResourceType.MarketplaceType` has been changed from `*ResourceTypeMarketplaceType` to `*MarketplaceType`
- Type of `ResourceTypeRegistrationProperties.MarketplaceType` has been changed from `*ResourceTypeRegistrationPropertiesMarketplaceType` to `*MarketplaceType`
- Type of `SKULocationInfo.Type` has been changed from `*SKULocationInfoType` to `*ExtendedLocationType`
- Enum `OperationsDefinitionActionType` has been removed
- Enum `OperationsDefinitionOrigin` has been removed
- Enum `ResourceProviderManagementResourceAccessPolicy` has been removed
- Enum `ResourceTypeMarketplaceType` has been removed
- Enum `ResourceTypeRegistrationPropertiesMarketplaceType` has been removed
- Enum `SKULocationInfoType` has been removed
- Operation `*CustomRolloutsClient.CreateOrUpdate` has been changed to LRO, use `*CustomRolloutsClient.BeginCreateOrUpdate` instead.
- Operation `*ResourceTypeRegistrationsClient.Delete` has been changed to LRO, use `*ResourceTypeRegistrationsClient.BeginDelete` instead.
- Struct `CustomRolloutPropertiesAutoGenerated` has been removed
- Struct `DefaultRolloutPropertiesAutoGenerated` has been removed
- Struct `Error` has been removed
- Struct `ErrorInnerError` has been removed
- Struct `ErrorResponseError` has been removed
- Struct `InnerError` has been removed
- Struct `NotificationRegistrationPropertiesAutoGenerated` has been removed
- Struct `ProviderRegistrationPropertiesAutoGenerated` has been removed
- Struct `ResourceTypeRegistrationPropertiesAutoGenerated` has been removed
- Field `OperationsContent` of struct `OperationsClientCreateOrUpdateResponse` has been removed
- Field `Contents` of struct `OperationsPutContent` has been removed

### Features Added

- New value `ExtensionCategoryBestMatchOperationBegin`, `ExtensionCategorySubscriptionLifecycleNotificationDeletion` added to enum type `ExtensionCategory`
- New value `OptInHeaderTypeClientPrincipalNameEncoded`, `OptInHeaderTypeMSIResourceIDEncoded`, `OptInHeaderTypeManagementGroupAncestorsEncoded`, `OptInHeaderTypePrivateLinkID`, `OptInHeaderTypePrivateLinkResourceID`, `OptInHeaderTypePrivateLinkVnetTrafficTag`, `OptInHeaderTypeResourceGroupLocation` added to enum type `OptInHeaderType`
- New value `RoutingTypeBypassEndpointSelectionOptimization`, `RoutingTypeCascadeAuthorizedExtension`, `RoutingTypeChildFanout`, `RoutingTypeLocationMapping`, `RoutingTypeServiceFanout` added to enum type `RoutingType`
- New enum type `AdditionalOptions` with values `AdditionalOptionsProtectedAsyncOperationPolling`, `AdditionalOptionsProtectedAsyncOperationPollingAuditOnly`
- New enum type `AdditionalOptionsAsyncOperation` with values `AdditionalOptionsAsyncOperationProtectedAsyncOperationPolling`, `AdditionalOptionsAsyncOperationProtectedAsyncOperationPollingAuditOnly`
- New enum type `AdditionalOptionsResourceTypeRegistration` with values `AdditionalOptionsResourceTypeRegistrationProtectedAsyncOperationPolling`, `AdditionalOptionsResourceTypeRegistrationProtectedAsyncOperationPollingAuditOnly`
- New enum type `AuthenticationScheme` with values `AuthenticationSchemeBearer`, `AuthenticationSchemePoP`
- New enum type `AvailabilityZonePolicy` with values `AvailabilityZonePolicyMultiZoned`, `AvailabilityZonePolicyNotSpecified`, `AvailabilityZonePolicySingleZoned`
- New enum type `AvailableCheckInManifestEnvironment` with values `AvailableCheckInManifestEnvironmentAll`, `AvailableCheckInManifestEnvironmentCanary`, `AvailableCheckInManifestEnvironmentFairfax`, `AvailableCheckInManifestEnvironmentMooncake`, `AvailableCheckInManifestEnvironmentNotSpecified`, `AvailableCheckInManifestEnvironmentProd`
- New enum type `BlockActionVerb` with values `BlockActionVerbAction`, `BlockActionVerbDelete`, `BlockActionVerbNotSpecified`, `BlockActionVerbRead`, `BlockActionVerbUnrecognized`, `BlockActionVerbWrite`
- New enum type `CapacityPolicy` with values `CapacityPolicyDefault`, `CapacityPolicyRestricted`
- New enum type `CommonAPIVersionsMergeMode` with values `CommonAPIVersionsMergeModeMerge`, `CommonAPIVersionsMergeModeOverwrite`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `CrossTenantTokenValidation` with values `CrossTenantTokenValidationEnsureSecureValidation`, `CrossTenantTokenValidationPassthroughInsecureToken`
- New enum type `DataBoundary` with values `DataBoundaryEU`, `DataBoundaryGlobal`, `DataBoundaryNotDefined`, `DataBoundaryUS`
- New enum type `EndpointType` with values `EndpointTypeCanary`, `EndpointTypeNotSpecified`, `EndpointTypeProduction`, `EndpointTypeTestInProduction`
- New enum type `EndpointTypeResourceType` with values `EndpointTypeResourceTypeCanary`, `EndpointTypeResourceTypeNotSpecified`, `EndpointTypeResourceTypeProduction`, `EndpointTypeResourceTypeTestInProduction`
- New enum type `ExpeditedRolloutIntent` with values `ExpeditedRolloutIntentHotfix`, `ExpeditedRolloutIntentNotSpecified`
- New enum type `ExtendedLocationType` with values `ExtendedLocationTypeArcZone`, `ExtendedLocationTypeCustomLocation`, `ExtendedLocationTypeEdgeZone`, `ExtendedLocationTypeNotSpecified`
- New enum type `FilterOption` with values `FilterOptionEnableSubscriptionFilterOnTenant`, `FilterOptionNotSpecified`
- New enum type `FrontdoorRequestMode` with values `FrontdoorRequestModeNotSpecified`, `FrontdoorRequestModeUseManifest`
- New enum type `Intent` with values `IntentDEFERREDACCESSCHECK`, `IntentLOWPRIVILEGE`, `IntentNOTSPECIFIED`, `IntentRPCONTRACT`
- New enum type `LegacyOperation` with values `LegacyOperationAction`, `LegacyOperationAzureAsyncOperationWaiting`, `LegacyOperationCreate`, `LegacyOperationDelete`, `LegacyOperationDeploymentCleanup`, `LegacyOperationEvaluateDeploymentOutput`, `LegacyOperationNotSpecified`, `LegacyOperationRead`, `LegacyOperationResourceCacheWaiting`, `LegacyOperationWaiting`
- New enum type `MarketplaceType` with values `MarketplaceTypeAddOn`, `MarketplaceTypeBypass`, `MarketplaceTypeNotSpecified`, `MarketplaceTypeStore`
- New enum type `NotificationEndpointType` with values `NotificationEndpointTypeEventhub`, `NotificationEndpointTypeWebhook`
- New enum type `NotificationOptions` with values `NotificationOptionsEmitSpendingLimit`, `NotificationOptionsNone`, `NotificationOptionsNotSpecified`
- New enum type `NotificationType` with values `NotificationTypeSubscriptionNotification`, `NotificationTypeUnspecified`
- New enum type `OperationActionType` with values `OperationActionTypeInternal`, `OperationActionTypeNotSpecified`
- New enum type `OperationOrigins` with values `OperationOriginsNotSpecified`, `OperationOriginsSystem`, `OperationOriginsUser`
- New enum type `OptOutHeaderType` with values `OptOutHeaderTypeNotSpecified`, `OptOutHeaderTypeSystemDataCreatedByLastModifiedBy`
- New enum type `Policy` with values `PolicyNotSpecified`, `PolicySynchronizeBeginExtension`
- New enum type `PolicyExecutionType` with values `PolicyExecutionTypeBypassPolicies`, `PolicyExecutionTypeExecutePolicies`, `PolicyExecutionTypeExpectPartialPutRequests`, `PolicyExecutionTypeNotSpecified`
- New enum type `ProviderRegistrationKind` with values `ProviderRegistrationKindDirect`, `ProviderRegistrationKindHybrid`, `ProviderRegistrationKindManaged`
- New enum type `QuotaPolicy` with values `QuotaPolicyDefault`, `QuotaPolicyNone`, `QuotaPolicyRestricted`
- New enum type `Readiness` with values `ReadinessClosingDown`, `ReadinessDeprecated`, `ReadinessGA`, `ReadinessInDevelopment`, `ReadinessInternalOnly`, `ReadinessPrivatePreview`, `ReadinessPublicPreview`, `ReadinessRemovedFromARM`, `ReadinessRetired`
- New enum type `ResourceAccessPolicy` with values `ResourceAccessPolicyAcisActionAllowed`, `ResourceAccessPolicyAcisReadAllowed`, `ResourceAccessPolicyNotSpecified`
- New enum type `ResourceSubType` with values `ResourceSubTypeAsyncOperation`, `ResourceSubTypeNotSpecified`
- New enum type `ResourceTypeCategory` with values `ResourceTypeCategoryFreeForm`, `ResourceTypeCategoryInternal`, `ResourceTypeCategoryNone`, `ResourceTypeCategoryPureProxy`
- New enum type `ResourceTypeEndpointKind` with values `ResourceTypeEndpointKindDirect`, `ResourceTypeEndpointKindManaged`
- New enum type `ResourceTypeExtendedLocationPolicy` with values `ResourceTypeExtendedLocationPolicyAll`, `ResourceTypeExtendedLocationPolicyNotSpecified`
- New enum type `ResourceTypeRegistrationKind` with values `ResourceTypeRegistrationKindDirect`, `ResourceTypeRegistrationKindHybrid`, `ResourceTypeRegistrationKindManaged`
- New enum type `Role` with values `RoleLimitedOwner`, `RoleServiceOwner`
- New enum type `ServerFailureResponseMessageType` with values `ServerFailureResponseMessageTypeNotSpecified`, `ServerFailureResponseMessageTypeOutageReporting`
- New enum type `ServiceClientOptionsType` with values `ServiceClientOptionsTypeDisableAutomaticDecompression`, `ServiceClientOptionsTypeNotSpecified`
- New enum type `ServiceFeatureFlagAction` with values `ServiceFeatureFlagActionCreate`, `ServiceFeatureFlagActionDoNotCreate`
- New enum type `ServiceStatus` with values `ServiceStatusActive`, `ServiceStatusInactive`
- New enum type `SignedRequestScope` with values `SignedRequestScopeEndpoint`, `SignedRequestScopeResourceURI`
- New enum type `SkipNotifications` with values `SkipNotificationsDisabled`, `SkipNotificationsEnabled`, `SkipNotificationsUnspecified`
- New enum type `SupportedOperations` with values `SupportedOperationsDelete`, `SupportedOperationsGet`, `SupportedOperationsNotSpecified`
- New enum type `TemplateDeploymentPreflightNotifications` with values `TemplateDeploymentPreflightNotificationsNone`, `TemplateDeploymentPreflightNotificationsUnregisteredSubscriptions`
- New function `NewAuthorizedApplicationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AuthorizedApplicationsClient, error)`
- New function `*AuthorizedApplicationsClient.BeginCreateOrUpdate(context.Context, string, string, AuthorizedApplication, *AuthorizedApplicationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AuthorizedApplicationsClientCreateOrUpdateResponse], error)`
- New function `*AuthorizedApplicationsClient.Delete(context.Context, string, string, *AuthorizedApplicationsClientDeleteOptions) (AuthorizedApplicationsClientDeleteResponse, error)`
- New function `*AuthorizedApplicationsClient.Get(context.Context, string, string, *AuthorizedApplicationsClientGetOptions) (AuthorizedApplicationsClientGetResponse, error)`
- New function `*AuthorizedApplicationsClient.NewListPager(string, *AuthorizedApplicationsClientListOptions) *runtime.Pager[AuthorizedApplicationsClientListResponse]`
- New function `*ClientFactory.NewAuthorizedApplicationsClient() *AuthorizedApplicationsClient`
- New function `*ClientFactory.NewNewRegionFrontloadReleaseClient() *NewRegionFrontloadReleaseClient`
- New function `*ClientFactory.NewProviderMonitorSettingsClient() *ProviderMonitorSettingsClient`
- New function `*ClientFactory.NewResourceActionsClient() *ResourceActionsClient`
- New function `*CustomRolloutsClient.Delete(context.Context, string, string, *CustomRolloutsClientDeleteOptions) (CustomRolloutsClientDeleteResponse, error)`
- New function `*CustomRolloutsClient.Stop(context.Context, string, string, *CustomRolloutsClientStopOptions) (CustomRolloutsClientStopResponse, error)`
- New function `PossibleAdditionalOptionsValues() []AdditionalOptions`
- New function `PossibleEndpointTypeValues() []EndpointType`
- New function `PossiblePolicyValues() []Policy`
- New function `NewProviderMonitorSettingsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ProviderMonitorSettingsClient, error)`
- New function `*ProviderMonitorSettingsClient.BeginCreate(context.Context, string, string, ProviderMonitorSetting, *ProviderMonitorSettingsClientBeginCreateOptions) (*runtime.Poller[ProviderMonitorSettingsClientCreateResponse], error)`
- New function `*ProviderMonitorSettingsClient.Delete(context.Context, string, string, *ProviderMonitorSettingsClientDeleteOptions) (ProviderMonitorSettingsClientDeleteResponse, error)`
- New function `*ProviderMonitorSettingsClient.Get(context.Context, string, string, *ProviderMonitorSettingsClientGetOptions) (ProviderMonitorSettingsClientGetResponse, error)`
- New function `*ProviderMonitorSettingsClient.NewListByResourceGroupPager(string, *ProviderMonitorSettingsClientListByResourceGroupOptions) *runtime.Pager[ProviderMonitorSettingsClientListByResourceGroupResponse]`
- New function `*ProviderMonitorSettingsClient.NewListBySubscriptionPager(*ProviderMonitorSettingsClientListBySubscriptionOptions) *runtime.Pager[ProviderMonitorSettingsClientListBySubscriptionResponse]`
- New function `*ProviderMonitorSettingsClient.Update(context.Context, string, string, *ProviderMonitorSettingsClientUpdateOptions) (ProviderMonitorSettingsClientUpdateResponse, error)`
- New function `NewNewRegionFrontloadReleaseClient(string, azcore.TokenCredential, *arm.ClientOptions) (*NewRegionFrontloadReleaseClient, error)`
- New function `*NewRegionFrontloadReleaseClient.CreateOrUpdate(context.Context, string, string, FrontloadPayload, *NewRegionFrontloadReleaseClientCreateOrUpdateOptions) (NewRegionFrontloadReleaseClientCreateOrUpdateResponse, error)`
- New function `*NewRegionFrontloadReleaseClient.GenerateManifest(context.Context, string, FrontloadPayload, *NewRegionFrontloadReleaseClientGenerateManifestOptions) (NewRegionFrontloadReleaseClientGenerateManifestResponse, error)`
- New function `*NewRegionFrontloadReleaseClient.Get(context.Context, string, string, *NewRegionFrontloadReleaseClientGetOptions) (NewRegionFrontloadReleaseClientGetResponse, error)`
- New function `*NewRegionFrontloadReleaseClient.Stop(context.Context, string, string, *NewRegionFrontloadReleaseClientStopOptions) (NewRegionFrontloadReleaseClientStopResponse, error)`
- New function `NewResourceActionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ResourceActionsClient, error)`
- New function `*ResourceActionsClient.BeginDeleteResources(context.Context, string, string, ResourceManagementAction, *ResourceActionsClientBeginDeleteResourcesOptions) (*runtime.Poller[ResourceActionsClientDeleteResourcesResponse], error)`
- New struct `APIProfile`
- New struct `AdditionalAuthorization`
- New struct `AllowedResourceName`
- New struct `AllowedUnauthorizedActionsExtension`
- New struct `ApplicationDataAuthorization`
- New struct `ApplicationProviderAuthorization`
- New struct `AsyncOperationPollingRules`
- New struct `AsyncTimeoutRule`
- New struct `AuthorizedApplication`
- New struct `AuthorizedApplicationArrayResponseWithContinuation`
- New struct `AuthorizedApplicationProperties`
- New struct `CustomRolloutSpecificationAutoProvisionConfig`
- New struct `CustomRolloutStatusManifestCheckinStatus`
- New struct `DefaultRolloutSpecificationAutoProvisionConfig`
- New struct `DefaultRolloutSpecificationExpeditedRollout`
- New struct `DefaultRolloutStatusManifestCheckinStatus`
- New struct `DeleteDependency`
- New struct `DstsConfiguration`
- New struct `EndpointInformation`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ExpeditedRolloutDefinition`
- New struct `FanoutLinkedNotificationRule`
- New struct `FanoutLinkedNotificationRuleDstsConfiguration`
- New struct `FilterRule`
- New struct `FrontloadPayload`
- New struct `FrontloadPayloadProperties`
- New struct `FrontloadPayloadPropertiesOverrideEndpointLevelFields`
- New struct `FrontloadPayloadPropertiesOverrideManifestLevelFields`
- New struct `GroupConnectivityInformation`
- New struct `LegacyDisallowedCondition`
- New struct `LinkedNotificationRule`
- New struct `LocalizedOperationDefinition`
- New struct `LocalizedOperationDefinitionDisplay`
- New struct `LocalizedOperationDisplayDefinition`
- New struct `LocalizedOperationDisplayDefinitionCs`
- New struct `LocalizedOperationDisplayDefinitionDe`
- New struct `LocalizedOperationDisplayDefinitionDefault`
- New struct `LocalizedOperationDisplayDefinitionEn`
- New struct `LocalizedOperationDisplayDefinitionEs`
- New struct `LocalizedOperationDisplayDefinitionFr`
- New struct `LocalizedOperationDisplayDefinitionHu`
- New struct `LocalizedOperationDisplayDefinitionIt`
- New struct `LocalizedOperationDisplayDefinitionJa`
- New struct `LocalizedOperationDisplayDefinitionKo`
- New struct `LocalizedOperationDisplayDefinitionNl`
- New struct `LocalizedOperationDisplayDefinitionPl`
- New struct `LocalizedOperationDisplayDefinitionPt`
- New struct `LocalizedOperationDisplayDefinitionPtBR`
- New struct `LocalizedOperationDisplayDefinitionRu`
- New struct `LocalizedOperationDisplayDefinitionSv`
- New struct `LocalizedOperationDisplayDefinitionZhHans`
- New struct `LocalizedOperationDisplayDefinitionZhHant`
- New struct `LocationQuotaRule`
- New struct `ManifestLevelPropertyBag`
- New struct `Notification`
- New struct `OpenAPIConfiguration`
- New struct `OpenAPIValidation`
- New struct `OperationsContentProperties`
- New struct `OperationsPutContentProperties`
- New struct `PrivateResourceProviderConfiguration`
- New struct `ProviderMonitorSetting`
- New struct `ProviderMonitorSettingArrayResponseWithContinuation`
- New struct `ProviderMonitorSettingProperties`
- New struct `ProviderRegistrationPropertiesPrivateResourceProviderConfiguration`
- New struct `QuotaRule`
- New struct `ResourceAccessRole`
- New struct `ResourceConcurrencyControlOption`
- New struct `ResourceGraphConfiguration`
- New struct `ResourceHydrationAccount`
- New struct `ResourceManagementAction`
- New struct `ResourceManagementEntity`
- New struct `ResourceProviderAuthorizationManagedByAuthorization`
- New struct `ResourceProviderAuthorizationRules`
- New struct `ResourceProviderManagementErrorResponseMessageOptions`
- New struct `ResourceProviderManagementExpeditedRolloutMetadata`
- New struct `ResourceProviderManifestPropertiesDstsConfiguration`
- New struct `ResourceProviderManifestPropertiesNotificationSettings`
- New struct `ResourceProviderManifestPropertiesResourceGroupLockOptionDuringMove`
- New struct `ResourceProviderManifestPropertiesResponseOptions`
- New struct `ResourceProviderService`
- New struct `ResourceTypeEndpointBase`
- New struct `ResourceTypeEndpointBaseDstsConfiguration`
- New struct `ResourceTypeEndpointBaseFeaturesRule`
- New struct `ResourceTypeEndpointDstsConfiguration`
- New struct `ResourceTypeOnBehalfOfToken`
- New struct `ResourceTypeRegistrationPropertiesAvailabilityZoneRule`
- New struct `ResourceTypeRegistrationPropertiesCapacityRule`
- New struct `ResourceTypeRegistrationPropertiesDstsConfiguration`
- New struct `ResourceTypeRegistrationPropertiesLegacyPolicy`
- New struct `ResourceTypeRegistrationPropertiesManagement`
- New struct `ResourceTypeRegistrationPropertiesMarketplaceOptions`
- New struct `ResourceTypeRegistrationPropertiesResourceCache`
- New struct `ResourceTypeRegistrationPropertiesResourceGraphConfiguration`
- New struct `ResourceTypeRegistrationPropertiesResourceManagementOptions`
- New struct `ResourceTypeRegistrationPropertiesResourceManagementOptionsBatchProvisioningSupport`
- New struct `ResourceTypeRegistrationPropertiesResourceManagementOptionsNestedProvisioningSupport`
- New struct `ResourceTypeRegistrationPropertiesResourceQueryManagement`
- New struct `ResourceTypeRegistrationPropertiesResourceTypeCommonAttributeManagement`
- New struct `ResourceTypeRegistrationPropertiesRoutingRule`
- New struct `ResourceTypeRegistrationPropertiesTemplateDeploymentPolicy`
- New struct `SubscriberSetting`
- New struct `SystemData`
- New struct `ThirdPartyExtension`
- New struct `TokenAuthConfiguration`
- New struct `TrackedResource`
- New field `SystemData` in struct `CustomRollout`
- New field `AutoProvisionConfig`, `RefreshSubscriptionRegistration`, `ReleaseScopes`, `SkipReleaseScopeValidation` in struct `CustomRolloutPropertiesSpecification`
- New field `ManifestCheckinStatus` in struct `CustomRolloutPropertiesStatus`
- New field `AutoProvisionConfig`, `RefreshSubscriptionRegistration`, `ReleaseScopes`, `SkipReleaseScopeValidation` in struct `CustomRolloutSpecification`
- New field `Kind`, `SystemData` in struct `CustomRolloutSpecificationProviderRegistration`
- New field `ManifestCheckinStatus` in struct `CustomRolloutStatus`
- New field `SystemData` in struct `DefaultRollout`
- New field `AutoProvisionConfig`, `ExpeditedRollout` in struct `DefaultRolloutPropertiesSpecification`
- New field `ManifestCheckinStatus` in struct `DefaultRolloutPropertiesStatus`
- New field `AutoProvisionConfig`, `ExpeditedRollout` in struct `DefaultRolloutSpecification`
- New field `Kind`, `SystemData` in struct `DefaultRolloutSpecificationProviderRegistration`
- New field `ManifestCheckinStatus` in struct `DefaultRolloutStatus`
- New field `ApplicationIDs`, `DelegationAppIDs` in struct `IdentityManagementProperties`
- New field `DependsOnTypes` in struct `LinkedOperationRule`
- New field `DirectRpRoleDefinitionID`, `GlobalAsyncOperationResourceTypeName`, `RegionalAsyncOperationResourceTypeName` in struct `Metadata`
- New field `SystemData` in struct `NotificationRegistration`
- New anonymous field `OperationsPutContent` in struct `OperationsClientCreateOrUpdateResponse`
- New field `SystemData` in struct `OperationsContent`
- New field `ID`, `Name`, `Properties`, `SystemData`, `Type` in struct `OperationsPutContent`
- New field `Kind`, `SystemData` in struct `ProviderRegistration`
- New field `CrossTenantTokenValidation`, `CustomManifestVersion`, `DstsConfiguration`, `EnableTenantLinkedNotification`, `GlobalNotificationEndpoints`, `LegacyNamespace`, `LegacyRegistrations`, `LinkedNotificationRules`, `ManagementGroupGlobalNotificationEndpoints`, `NotificationOptions`, `NotificationSettings`, `Notifications`, `OptionalFeatures`, `PrivateResourceProviderConfiguration`, `ResourceGroupLockOptionDuringMove`, `ResourceHydrationAccounts`, `ResourceProviderAuthorizationRules`, `ResponseOptions`, `ServiceName`, `Services`, `TokenAuthConfiguration` in struct `ProviderRegistrationProperties`
- New field `DirectRpRoleDefinitionID`, `GlobalAsyncOperationResourceTypeName`, `RegionalAsyncOperationResourceTypeName` in struct `ProviderRegistrationPropertiesProviderHubMetadata`
- New field `SystemData` in struct `ProxyResource`
- New field `OptOutHeaders` in struct `RequestHeaderOptions`
- New field `SystemData` in struct `Resource`
- New field `AllowedThirdPartyExtensions`, `GroupingTag`, `ManagedByAuthorization` in struct `ResourceProviderAuthorization`
- New field `EndpointType`, `SKULink` in struct `ResourceProviderEndpoint`
- New field `AuthorizationOwners`, `CanaryManifestOwners`, `ErrorResponseMessageOptions`, `ExpeditedRolloutMetadata`, `ExpeditedRolloutSubmitters`, `PcCode`, `ProfitCenterProgramID` in struct `ResourceProviderManagement`
- New field `CrossTenantTokenValidation`, `EnableTenantLinkedNotification`, `LinkedNotificationRules`, `Notifications`, `ResourceProviderAuthorizationRules`, `ServiceName`, `Services` in struct `ResourceProviderManifest`
- New field `AuthorizationOwners`, `CanaryManifestOwners`, `ErrorResponseMessageOptions`, `ExpeditedRolloutMetadata`, `ExpeditedRolloutSubmitters`, `PcCode`, `ProfitCenterProgramID` in struct `ResourceProviderManifestManagement`
- New field `CrossTenantTokenValidation`, `CustomManifestVersion`, `DstsConfiguration`, `EnableTenantLinkedNotification`, `GlobalNotificationEndpoints`, `LegacyNamespace`, `LegacyRegistrations`, `LinkedNotificationRules`, `ManagementGroupGlobalNotificationEndpoints`, `NotificationOptions`, `NotificationSettings`, `Notifications`, `OptionalFeatures`, `ResourceGroupLockOptionDuringMove`, `ResourceHydrationAccounts`, `ResourceProviderAuthorizationRules`, `ResponseOptions`, `ServiceName`, `Services` in struct `ResourceProviderManifestProperties`
- New field `AuthorizationOwners`, `CanaryManifestOwners`, `ErrorResponseMessageOptions`, `ExpeditedRolloutMetadata`, `ExpeditedRolloutSubmitters`, `PcCode`, `ProfitCenterProgramID` in struct `ResourceProviderManifestPropertiesManagement`
- New field `OptOutHeaders` in struct `ResourceProviderManifestPropertiesRequestHeaderOptions`
- New field `OptOutHeaders` in struct `ResourceProviderManifestRequestHeaderOptions`
- New field `AdditionalOptions`, `AllowedUnauthorizedActionsExtensions`, `CrossTenantTokenValidation`, `LinkedNotificationRules`, `Notifications`, `QuotaRule`, `ResourceProviderAuthorizationRules` in struct `ResourceType`
- New field `APIVersion`, `DataBoundary`, `DstsConfiguration`, `EndpointType`, `EndpointURI`, `Kind`, `SKULink`, `TokenAuthConfiguration`, `Zones` in struct `ResourceTypeEndpoint`
- New field `Kind`, `SystemData` in struct `ResourceTypeRegistration`
- New field `APIProfiles`, `AddResourceListTargetLocations`, `AdditionalOptions`, `AllowEmptyRoleAssignments`, `AllowedResourceNames`, `AllowedTemplateDeploymentReferenceActions`, `AllowedUnauthorizedActionsExtensions`, `AsyncOperationResourceTypeName`, `AsyncTimeoutRules`, `AvailabilityZoneRule`, `CapacityRule`, `Category`, `CommonAPIVersions`, `CrossTenantTokenValidation`, `DisallowedEndUserOperations`, `DstsConfiguration`, `FrontdoorRequestMode`, `GroupingTag`, `LegacyName`, `LegacyNames`, `LegacyPolicy`, `LinkedNotificationRules`, `LinkedOperationRules`, `Management`, `ManifestLink`, `MarketplaceOptions`, `Metadata`, `Notifications`, `OnBehalfOfTokens`, `OpenAPIConfiguration`, `PolicyExecutionType`, `QuotaRule`, `ResourceCache`, `ResourceConcurrencyControlOptions`, `ResourceGraphConfiguration`, `ResourceManagementOptions`, `ResourceProviderAuthorizationRules`, `ResourceQueryManagement`, `ResourceSubType`, `ResourceTypeCommonAttributeManagement`, `ResourceValidation`, `RoutingRule`, `SKULink`, `SupportsTags`, `TemplateDeploymentPolicy`, `TokenAuthConfiguration` in struct `ResourceTypeRegistrationProperties`
- New field `ApplicationIDs`, `DelegationAppIDs` in struct `ResourceTypeRegistrationPropertiesIdentityManagement`
- New field `OptOutHeaders` in struct `ResourceTypeRegistrationPropertiesRequestHeaderOptions`
- New field `OptOutHeaders` in struct `ResourceTypeRequestHeaderOptions`
- New field `PreflightNotifications` in struct `ResourceTypeTemplateDeploymentPolicy`
- New field `SystemData` in struct `SKUResource`
- New field `Readiness` in struct `ServiceTreeInfo`
- New field `PreflightNotifications` in struct `TemplateDeploymentPolicy`
- New field `ApplicationID` in struct `ThrottlingRule`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).