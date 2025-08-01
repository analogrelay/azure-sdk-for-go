//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_ListBySubscription.json
func ExampleDiskEncryptionSetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiskEncryptionSetsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DiskEncryptionSetList = armcompute.DiskEncryptionSetList{
		// 	Value: []*armcompute.DiskEncryptionSet{
		// 		{
		// 			Name: to.Ptr("myDiskEncryptionSet"),
		// 			Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("Encryption"),
		// 			},
		// 			Identity: &armcompute.EncryptionSetIdentity{
		// 				Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		// 			},
		// 			Properties: &armcompute.EncryptionSetProperties{
		// 				ActiveKey: &armcompute.KeyForDiskEncryptionSet{
		// 					KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
		// 					SourceVault: &armcompute.SourceVault{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 					},
		// 				},
		// 				EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
		// 				PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myDiskEncryptionSet2"),
		// 			Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/mySecondResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("Encryption"),
		// 			},
		// 			Identity: &armcompute.EncryptionSetIdentity{
		// 				Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		// 			},
		// 			Properties: &armcompute.EncryptionSetProperties{
		// 				ActiveKey: &armcompute.KeyForDiskEncryptionSet{
		// 					KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
		// 					SourceVault: &armcompute.SourceVault{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/mySecondResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault2"),
		// 					},
		// 				},
		// 				EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
		// 				PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_ListByResourceGroup.json
func ExampleDiskEncryptionSetsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiskEncryptionSetsClient().NewListByResourceGroupPager("myResourceGroup", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DiskEncryptionSetList = armcompute.DiskEncryptionSetList{
		// 	Value: []*armcompute.DiskEncryptionSet{
		// 		{
		// 			Name: to.Ptr("myDiskEncryptionSet"),
		// 			Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("Encryption"),
		// 			},
		// 			Identity: &armcompute.EncryptionSetIdentity{
		// 				Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		// 			},
		// 			Properties: &armcompute.EncryptionSetProperties{
		// 				ActiveKey: &armcompute.KeyForDiskEncryptionSet{
		// 					KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
		// 					SourceVault: &armcompute.SourceVault{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 					},
		// 				},
		// 				EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
		// 				PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myDiskEncryptionSet2"),
		// 			Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("Encryption"),
		// 			},
		// 			Identity: &armcompute.EncryptionSetIdentity{
		// 				Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		// 			},
		// 			Properties: &armcompute.EncryptionSetProperties{
		// 				ActiveKey: &armcompute.KeyForDiskEncryptionSet{
		// 					KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
		// 					SourceVault: &armcompute.SourceVault{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault2"),
		// 					},
		// 				},
		// 				EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
		// 				PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Get_WithAutoKeyRotationError.json
func ExampleDiskEncryptionSetsClient_Get_getInformationAboutADiskEncryptionSetWhenAutoKeyRotationFailed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiskEncryptionSetsClient().Get(ctx, "myResourceGroup", "myDiskEncryptionSet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("Development"),
	// 		"project": to.Ptr("Encryption"),
	// 	},
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
	// 			SourceVault: &armcompute.SourceVault{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
	// 			},
	// 		},
	// 		AutoKeyRotationError: &armcompute.APIError{
	// 			Code: to.Ptr("ManagedServiceIdentityNotFound"),
	// 			Message: to.Ptr("Auto-key rotation was disabled as managed service identity associated with DiskEncryptionSet 'myDiskEncryptionSet' was not found. Please update the resource with correct identity to re-enable auto-key rotation."),
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RotationToLatestKeyVersionEnabled: to.Ptr(true),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Get.json
func ExampleDiskEncryptionSetsClient_Get_getInformationAboutADiskEncryptionSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiskEncryptionSetsClient().Get(ctx, "myResourceGroup", "myDiskEncryptionSet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("Development"),
	// 		"project": to.Ptr("Encryption"),
	// 	},
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
	// 			SourceVault: &armcompute.SourceVault{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
	// 			},
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentSubscription.json
func ExampleDiskEncryptionSetsClient_BeginCreateOrUpdate_createADiskEncryptionSetWithKeyVaultFromADifferentSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskEncryptionSetsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDiskEncryptionSet", armcompute.DiskEncryptionSet{
		Location: to.Ptr("West US"),
		Identity: &armcompute.EncryptionSetIdentity{
			Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		},
		Properties: &armcompute.EncryptionSetProperties{
			ActiveKey: &armcompute.KeyForDiskEncryptionSet{
				KeyURL: to.Ptr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/{key}"),
			},
			EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/{key}"),
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentTenant.json
func ExampleDiskEncryptionSetsClient_BeginCreateOrUpdate_createADiskEncryptionSetWithKeyVaultFromADifferentTenant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskEncryptionSetsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDiskEncryptionSet", armcompute.DiskEncryptionSet{
		Location: to.Ptr("West US"),
		Identity: &armcompute.EncryptionSetIdentity{
			Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
				"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}": {},
			},
		},
		Properties: &armcompute.EncryptionSetProperties{
			ActiveKey: &armcompute.KeyForDiskEncryptionSet{
				KeyURL: to.Ptr("https://myvaultdifferenttenant.vault-int.azure-int.net/keys/{key}"),
			},
			EncryptionType:    to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
			FederatedClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
	// 			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}": &armcompute.UserAssignedIdentitiesValue{
	// 			},
	// 		},
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvaultdifferenttenant.vault-int.azure-int.net/keys/{key}"),
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		FederatedClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Create.json
func ExampleDiskEncryptionSetsClient_BeginCreateOrUpdate_createADiskEncryptionSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskEncryptionSetsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDiskEncryptionSet", armcompute.DiskEncryptionSet{
		Location: to.Ptr("West US"),
		Identity: &armcompute.EncryptionSetIdentity{
			Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		},
		Properties: &armcompute.EncryptionSetProperties{
			ActiveKey: &armcompute.KeyForDiskEncryptionSet{
				KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
				SourceVault: &armcompute.SourceVault{
					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
				},
			},
			EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
	// 			SourceVault: &armcompute.SourceVault{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
	// 			},
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Update_WithRotationToLatestKeyVersionEnabled.json
func ExampleDiskEncryptionSetsClient_BeginUpdate_updateADiskEncryptionSetWithRotationToLatestKeyVersionEnabledSetToTrueSucceeded() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskEncryptionSetsClient().BeginUpdate(ctx, "myResourceGroup", "myDiskEncryptionSet", armcompute.DiskEncryptionSetUpdate{
		Identity: &armcompute.EncryptionSetIdentity{
			Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		},
		Properties: &armcompute.DiskEncryptionSetUpdateProperties{
			ActiveKey: &armcompute.KeyForDiskEncryptionSet{
				KeyURL: to.Ptr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/keyVersion1"),
			},
			EncryptionType:                    to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
			RotationToLatestKeyVersionEnabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/KeyVersion2"),
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T04:41:35.079Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RotationToLatestKeyVersionEnabled: to.Ptr(true),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Update_WithRotationToLatestKeyVersionEnabledInProgress.json
func ExampleDiskEncryptionSetsClient_BeginUpdate_updateADiskEncryptionSetWithRotationToLatestKeyVersionEnabledSetToTrueUpdating() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskEncryptionSetsClient().BeginUpdate(ctx, "myResourceGroup", "myDiskEncryptionSet", armcompute.DiskEncryptionSetUpdate{
		Identity: &armcompute.EncryptionSetIdentity{
			Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		},
		Properties: &armcompute.DiskEncryptionSetUpdateProperties{
			ActiveKey: &armcompute.KeyForDiskEncryptionSet{
				KeyURL: to.Ptr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/keyVersion1"),
			},
			EncryptionType:                    to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
			RotationToLatestKeyVersionEnabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/keyVersion2"),
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T04:41:35.079Z"); return t}()),
	// 		PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 			{
	// 				KeyURL: to.Ptr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/keyVersion1"),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RotationToLatestKeyVersionEnabled: to.Ptr(true),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Update.json
func ExampleDiskEncryptionSetsClient_BeginUpdate_updateADiskEncryptionSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskEncryptionSetsClient().BeginUpdate(ctx, "myResourceGroup", "myDiskEncryptionSet", armcompute.DiskEncryptionSetUpdate{
		Properties: &armcompute.DiskEncryptionSetUpdateProperties{
			ActiveKey: &armcompute.KeyForDiskEncryptionSet{
				KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/keyName/keyVersion"),
				SourceVault: &armcompute.SourceVault{
					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
				},
			},
			EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
		},
		Tags: map[string]*string{
			"department": to.Ptr("Development"),
			"project":    to.Ptr("Encryption"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("Development"),
	// 		"project": to.Ptr("Encryption"),
	// 	},
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/keyName/keyVersion"),
	// 			SourceVault: &armcompute.SourceVault{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
	// 			},
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T04:41:35.079Z"); return t}()),
	// 		PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Delete.json
func ExampleDiskEncryptionSetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskEncryptionSetsClient().BeginDelete(ctx, "myResourceGroup", "myDiskEncryptionSet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_ListAssociatedResources.json
func ExampleDiskEncryptionSetsClient_NewListAssociatedResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiskEncryptionSetsClient().NewListAssociatedResourcesPager("myResourceGroup", "myDiskEncryptionSet", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ResourceURIList = armcompute.ResourceURIList{
		// 	Value: []*string{
		// 		to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
		// 		to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot")},
		// 	}
	}
}
