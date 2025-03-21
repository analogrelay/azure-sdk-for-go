//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherConnectionMonitorCreate.json
func ExampleConnectionMonitorsClient_BeginCreateOrUpdate_createConnectionMonitorV1() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionMonitorsClient().BeginCreateOrUpdate(ctx, "rg1", "nw1", "cm1", armnetwork.ConnectionMonitor{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.ConnectionMonitorParameters{
			Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
				{
					Name:       to.Ptr("source"),
					ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
				},
				{
					Name:    to.Ptr("destination"),
					Address: to.Ptr("bing.com"),
				}},
			TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
				{
					Name: to.Ptr("tcp"),
					TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
						Port: to.Ptr[int32](80),
					},
					TestFrequencySec: to.Ptr[int32](60),
					Protocol:         to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
				}},
			TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
				{
					Name: to.Ptr("tg"),
					Destinations: []*string{
						to.Ptr("destination")},
					Sources: []*string{
						to.Ptr("source")},
					TestConfigurations: []*string{
						to.Ptr("tcp")},
				}},
		},
	}, &armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions{Migrate: nil})
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
	// res.ConnectionMonitorResult = armnetwork.ConnectionMonitorResult{
	// 	Name: to.Ptr("cm1"),
	// 	Type: to.Ptr("Microsoft.Network/networkWatchers/connectionMonitors"),
	// 	Etag: to.Ptr("W/\"e7497f26-5f09-4559-900b-fe98f3dedb6f\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1/connectionMonitors/cm1"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.ConnectionMonitorResultProperties{
	// 		Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
	// 			{
	// 				Name: to.Ptr("source"),
	// 				ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("destination"),
	// 				Address: to.Ptr("bing.com"),
	// 		}},
	// 		TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
	// 			{
	// 				Name: to.Ptr("tcp"),
	// 				TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
	// 					Port: to.Ptr[int32](80),
	// 				},
	// 				TestFrequencySec: to.Ptr[int32](60),
	// 				Protocol: to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
	// 		}},
	// 		TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
	// 			{
	// 				Name: to.Ptr("tg"),
	// 				Destinations: []*string{
	// 					to.Ptr("destination")},
	// 					Sources: []*string{
	// 						to.Ptr("source")},
	// 						TestConfigurations: []*string{
	// 							to.Ptr("tcp")},
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherConnectionMonitorV2Create.json
func ExampleConnectionMonitorsClient_BeginCreateOrUpdate_createConnectionMonitorV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionMonitorsClient().BeginCreateOrUpdate(ctx, "rg1", "nw1", "cm1", armnetwork.ConnectionMonitor{
		Properties: &armnetwork.ConnectionMonitorParameters{
			Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
				{
					Name:       to.Ptr("vm1"),
					ResourceID: to.Ptr("/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/NwRgIrinaCentralUSEUAP/providers/Microsoft.Compute/virtualMachines/vm1"),
				},
				{
					Name: to.Ptr("CanaryWorkspaceVamshi"),
					Filter: &armnetwork.ConnectionMonitorEndpointFilter{
						Type: to.Ptr(armnetwork.ConnectionMonitorEndpointFilterTypeInclude),
						Items: []*armnetwork.ConnectionMonitorEndpointFilterItem{
							{
								Type:    to.Ptr(armnetwork.ConnectionMonitorEndpointFilterItemTypeAgentAddress),
								Address: to.Ptr("npmuser"),
							}},
					},
					ResourceID: to.Ptr("/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/vasamudrRG/providers/Microsoft.OperationalInsights/workspaces/vasamudrWorkspace"),
				},
				{
					Name:    to.Ptr("bing"),
					Address: to.Ptr("bing.com"),
				},
				{
					Name:    to.Ptr("google"),
					Address: to.Ptr("google.com"),
				}},
			Outputs: []*armnetwork.ConnectionMonitorOutput{},
			TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
				{
					Name: to.Ptr("testConfig1"),
					TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
						DisableTraceRoute: to.Ptr(false),
						Port:              to.Ptr[int32](80),
					},
					TestFrequencySec: to.Ptr[int32](60),
					Protocol:         to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
				}},
			TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
				{
					Name: to.Ptr("test1"),
					Destinations: []*string{
						to.Ptr("bing"),
						to.Ptr("google")},
					Disable: to.Ptr(false),
					Sources: []*string{
						to.Ptr("vm1"),
						to.Ptr("CanaryWorkspaceVamshi")},
					TestConfigurations: []*string{
						to.Ptr("testConfig1")},
				}},
		},
	}, &armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions{Migrate: nil})
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
	// res.ConnectionMonitorResult = armnetwork.ConnectionMonitorResult{
	// 	Name: to.Ptr("cm1"),
	// 	Type: to.Ptr("Microsoft.Network/networkWatchers/connectionMonitors"),
	// 	Etag: to.Ptr("W/\"e7497f26-5f09-4559-900b-fe98f3dedb6f\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1/connectionMonitors/cm1"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: &armnetwork.ConnectionMonitorResultProperties{
	// 		Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
	// 			{
	// 				Name: to.Ptr("vm1"),
	// 				ResourceID: to.Ptr("/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/NwRgIrinaCentralUSEUAP/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CanaryWorkspaceVamshi"),
	// 				Filter: &armnetwork.ConnectionMonitorEndpointFilter{
	// 					Type: to.Ptr(armnetwork.ConnectionMonitorEndpointFilterTypeInclude),
	// 					Items: []*armnetwork.ConnectionMonitorEndpointFilterItem{
	// 						{
	// 							Type: to.Ptr(armnetwork.ConnectionMonitorEndpointFilterItemTypeAgentAddress),
	// 							Address: to.Ptr("npmuser"),
	// 					}},
	// 				},
	// 				ResourceID: to.Ptr("/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/vasamudrRG/providers/Microsoft.OperationalInsights/workspaces/vasamudrWorkspace"),
	// 			},
	// 			{
	// 				Name: to.Ptr("bing"),
	// 				Address: to.Ptr("bing.com"),
	// 			},
	// 			{
	// 				Name: to.Ptr("google"),
	// 				Address: to.Ptr("google.com"),
	// 		}},
	// 		Outputs: []*armnetwork.ConnectionMonitorOutput{
	// 		},
	// 		TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
	// 			{
	// 				Name: to.Ptr("testConfig1"),
	// 				TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
	// 					DisableTraceRoute: to.Ptr(false),
	// 					Port: to.Ptr[int32](80),
	// 				},
	// 				TestFrequencySec: to.Ptr[int32](60),
	// 				Protocol: to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
	// 		}},
	// 		TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
	// 			{
	// 				Name: to.Ptr("test1"),
	// 				Destinations: []*string{
	// 					to.Ptr("bing"),
	// 					to.Ptr("google")},
	// 					Disable: to.Ptr(false),
	// 					Sources: []*string{
	// 						to.Ptr("vm1"),
	// 						to.Ptr("CanaryWorkspaceVamshi")},
	// 						TestConfigurations: []*string{
	// 							to.Ptr("testConfig1")},
	// 					}},
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherConnectionMonitorCreateWithArcNetwork.json
func ExampleConnectionMonitorsClient_BeginCreateOrUpdate_createConnectionMonitorWithArcNetwork() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionMonitorsClient().BeginCreateOrUpdate(ctx, "rg1", "nw1", "cm1", armnetwork.ConnectionMonitor{
		Properties: &armnetwork.ConnectionMonitorParameters{
			Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
				{
					Name:       to.Ptr("vm1"),
					Type:       to.Ptr(armnetwork.EndpointTypeAzureVM),
					ResourceID: to.Ptr("/subscriptions/9cece3e3-0f7d-47ca-af0e-9772773f90b7/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/TESTVM"),
				},
				{
					Name:    to.Ptr("bing"),
					Type:    to.Ptr(armnetwork.EndpointTypeExternalAddress),
					Address: to.Ptr("bing.com"),
				},
				{
					Name:    to.Ptr("google"),
					Type:    to.Ptr(armnetwork.EndpointTypeExternalAddress),
					Address: to.Ptr("google.com"),
				},
				{
					Name: to.Ptr("ArcBasedNetwork"),
					Type: to.Ptr(armnetwork.EndpointTypeAzureArcNetwork),
					LocationDetails: &armnetwork.ConnectionMonitorEndpointLocationDetails{
						Region: to.Ptr("eastus"),
					},
					Scope: &armnetwork.ConnectionMonitorEndpointScope{
						Include: []*armnetwork.ConnectionMonitorEndpointScopeItem{
							{
								Address: to.Ptr("172.21.128.0/20"),
							}},
					},
					SubscriptionID: to.Ptr("9cece3e3-0f7d-47ca-af0e-9772773f90b7"),
				}},
			Outputs: []*armnetwork.ConnectionMonitorOutput{},
			TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
				{
					Name: to.Ptr("testConfig1"),
					TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
						DisableTraceRoute: to.Ptr(false),
						Port:              to.Ptr[int32](80),
					},
					TestFrequencySec: to.Ptr[int32](60),
					Protocol:         to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
				}},
			TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
				{
					Name: to.Ptr("test1"),
					Destinations: []*string{
						to.Ptr("bing"),
						to.Ptr("google")},
					Disable: to.Ptr(false),
					Sources: []*string{
						to.Ptr("vm1"),
						to.Ptr("ArcBasedNetwork")},
					TestConfigurations: []*string{
						to.Ptr("testConfig1")},
				}},
		},
	}, &armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions{Migrate: nil})
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
	// res.ConnectionMonitorResult = armnetwork.ConnectionMonitorResult{
	// 	Name: to.Ptr("cm1"),
	// 	Type: to.Ptr("Microsoft.Network/networkWatchers/connectionMonitors"),
	// 	Etag: to.Ptr("W/\"e7497f26-5f09-4559-900b-fe98f3dedb6f\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1/connectionMonitors/cm1"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: &armnetwork.ConnectionMonitorResultProperties{
	// 		Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
	// 			{
	// 				Name: to.Ptr("vm1"),
	// 				ResourceID: to.Ptr("/subscriptions/9cece3e3-0f7d-47ca-af0e-9772773f90b7/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/TESTVM"),
	// 			},
	// 			{
	// 				Name: to.Ptr("ArcBasedNetwork"),
	// 				Type: to.Ptr(armnetwork.EndpointTypeAzureArcNetwork),
	// 				LocationDetails: &armnetwork.ConnectionMonitorEndpointLocationDetails{
	// 					Region: to.Ptr("eastus"),
	// 				},
	// 				Scope: &armnetwork.ConnectionMonitorEndpointScope{
	// 					Include: []*armnetwork.ConnectionMonitorEndpointScopeItem{
	// 						{
	// 							Address: to.Ptr("172.21.128.0/20"),
	// 					}},
	// 				},
	// 				SubscriptionID: to.Ptr("9cece3e3-0f7d-47ca-af0e-9772773f90b7"),
	// 			},
	// 			{
	// 				Name: to.Ptr("bing"),
	// 				Address: to.Ptr("bing.com"),
	// 			},
	// 			{
	// 				Name: to.Ptr("google"),
	// 				Address: to.Ptr("google.com"),
	// 		}},
	// 		Outputs: []*armnetwork.ConnectionMonitorOutput{
	// 		},
	// 		TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
	// 			{
	// 				Name: to.Ptr("testConfig1"),
	// 				TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
	// 					DisableTraceRoute: to.Ptr(false),
	// 					Port: to.Ptr[int32](80),
	// 				},
	// 				TestFrequencySec: to.Ptr[int32](60),
	// 				Protocol: to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
	// 		}},
	// 		TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
	// 			{
	// 				Name: to.Ptr("test1"),
	// 				Destinations: []*string{
	// 					to.Ptr("bing"),
	// 					to.Ptr("google")},
	// 					Disable: to.Ptr(false),
	// 					Sources: []*string{
	// 						to.Ptr("vm1"),
	// 						to.Ptr("ArcBasedNetwork")},
	// 						TestConfigurations: []*string{
	// 							to.Ptr("testConfig1")},
	// 					}},
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherConnectionMonitorGet.json
func ExampleConnectionMonitorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionMonitorsClient().Get(ctx, "rg1", "nw1", "cm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionMonitorResult = armnetwork.ConnectionMonitorResult{
	// 	Name: to.Ptr("cm1"),
	// 	Type: to.Ptr("Microsoft.Network/networkWatchers/connectionMonitors"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1/connectionMonitors/cm1"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.ConnectionMonitorResultProperties{
	// 		Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
	// 			{
	// 				Name: to.Ptr("source"),
	// 				ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("destination"),
	// 				Address: to.Ptr("bing.com"),
	// 		}},
	// 		TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
	// 			{
	// 				Name: to.Ptr("tcp"),
	// 				TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
	// 					Port: to.Ptr[int32](80),
	// 				},
	// 				TestFrequencySec: to.Ptr[int32](60),
	// 				Protocol: to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
	// 		}},
	// 		TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
	// 			{
	// 				Name: to.Ptr("tg"),
	// 				Destinations: []*string{
	// 					to.Ptr("destination")},
	// 					Sources: []*string{
	// 						to.Ptr("source")},
	// 						TestConfigurations: []*string{
	// 							to.Ptr("tcp")},
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherConnectionMonitorDelete.json
func ExampleConnectionMonitorsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionMonitorsClient().BeginDelete(ctx, "rg1", "nw1", "cm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherConnectionMonitorUpdateTags.json
func ExampleConnectionMonitorsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionMonitorsClient().UpdateTags(ctx, "rg1", "nw1", "cm1", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionMonitorResult = armnetwork.ConnectionMonitorResult{
	// 	Name: to.Ptr("cm1"),
	// 	Type: to.Ptr("Microsoft.Network/networkWatchers/connectionMonitors"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	ID: to.Ptr("/subscriptions/subid/`/rg1/providers/Microsoft.Network/networkWatchers/nw1/connectionMonitors/cm1"),
	// 	Location: to.Ptr("westcentralus"),
	// 	Properties: &armnetwork.ConnectionMonitorResultProperties{
	// 		AutoStart: to.Ptr(true),
	// 		Destination: &armnetwork.ConnectionMonitorDestination{
	// 			Address: to.Ptr("bing.com"),
	// 			Port: to.Ptr[int32](80),
	// 		},
	// 		MonitoringIntervalInSeconds: to.Ptr[int32](60),
	// 		Source: &armnetwork.ConnectionMonitorSource{
	// 			Port: to.Ptr[int32](0),
	// 			ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 		},
	// 		MonitoringStatus: to.Ptr("Running"),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T02:48:10.679Z"); return t}()),
	// 	},
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherConnectionMonitorStop.json
func ExampleConnectionMonitorsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionMonitorsClient().BeginStop(ctx, "rg1", "nw1", "cm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherConnectionMonitorStart.json
func ExampleConnectionMonitorsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionMonitorsClient().BeginStart(ctx, "rg1", "nw1", "cm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherConnectionMonitorQuery.json
func ExampleConnectionMonitorsClient_BeginQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionMonitorsClient().BeginQuery(ctx, "rg1", "nw1", "cm1", nil)
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
	// res.ConnectionMonitorQueryResult = armnetwork.ConnectionMonitorQueryResult{
	// 	SourceStatus: to.Ptr(armnetwork.ConnectionMonitorSourceStatusActive),
	// 	States: []*armnetwork.ConnectionStateSnapshot{
	// 		{
	// 			ConnectionState: to.Ptr(armnetwork.ConnectionStateReachable),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-08T05:12:41.526Z"); return t}()),
	// 			EvaluationState: to.Ptr(armnetwork.EvaluationStateCompleted),
	// 			Hops: []*armnetwork.ConnectivityHop{
	// 				{
	// 					Type: to.Ptr("Source"),
	// 					Address: to.Ptr("10.1.1.4"),
	// 					ID: to.Ptr("7dbbe7aa-60ba-4650-831e-63d775d38e9e"),
	// 					Issues: []*armnetwork.ConnectivityIssue{
	// 					},
	// 					NextHopIDs: []*string{
	// 						to.Ptr("75c8d819-b208-4584-a311-1aa45ce753f9")},
	// 						ResourceID: to.Ptr("subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic0/ipConfigurations/ipconfig1"),
	// 					},
	// 					{
	// 						Type: to.Ptr("VirtualNetwork"),
	// 						Address: to.Ptr("192.168.100.4"),
	// 						ID: to.Ptr("75c8d819-b208-4584-a311-1aa45ce753f9"),
	// 						Issues: []*armnetwork.ConnectivityIssue{
	// 						},
	// 						NextHopIDs: []*string{
	// 						},
	// 						ResourceID: to.Ptr("subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/ipconfig1"),
	// 				}},
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-08T03:42:33.338Z"); return t}()),
	// 		}},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherConnectionMonitorList.json
func ExampleConnectionMonitorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionMonitorsClient().NewListPager("rg1", "nw1", nil)
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
		// page.ConnectionMonitorListResult = armnetwork.ConnectionMonitorListResult{
		// 	Value: []*armnetwork.ConnectionMonitorResult{
		// 		{
		// 			Name: to.Ptr("cm1"),
		// 			Type: to.Ptr("Microsoft.Network/networkWatchers/connectionMonitors"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1/connectionMonitors/cm1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetwork.ConnectionMonitorResultProperties{
		// 				Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
		// 					{
		// 						Name: to.Ptr("source"),
		// 						ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
		// 					},
		// 					{
		// 						Name: to.Ptr("destination"),
		// 						Address: to.Ptr("bing.com"),
		// 				}},
		// 				TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
		// 					{
		// 						Name: to.Ptr("tcp"),
		// 						TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
		// 							Port: to.Ptr[int32](80),
		// 						},
		// 						TestFrequencySec: to.Ptr[int32](60),
		// 						Protocol: to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
		// 				}},
		// 				TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
		// 					{
		// 						Name: to.Ptr("tg"),
		// 						Destinations: []*string{
		// 							to.Ptr("destination")},
		// 							Sources: []*string{
		// 								to.Ptr("source")},
		// 								TestConfigurations: []*string{
		// 									to.Ptr("tcp")},
		// 							}},
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("cm2"),
		// 						Type: to.Ptr("Microsoft.Network/networkWatchers/connectionMonitors"),
		// 						Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1/connectionMonitors/cm2"),
		// 						Location: to.Ptr("eastus"),
		// 						Properties: &armnetwork.ConnectionMonitorResultProperties{
		// 							Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
		// 								{
		// 									Name: to.Ptr("source"),
		// 									ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct2"),
		// 								},
		// 								{
		// 									Name: to.Ptr("destination"),
		// 									Address: to.Ptr("google.com"),
		// 							}},
		// 							TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
		// 								{
		// 									Name: to.Ptr("tcp"),
		// 									TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
		// 										Port: to.Ptr[int32](80),
		// 									},
		// 									TestFrequencySec: to.Ptr[int32](60),
		// 									Protocol: to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
		// 							}},
		// 							TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
		// 								{
		// 									Name: to.Ptr("tg"),
		// 									Destinations: []*string{
		// 										to.Ptr("destination")},
		// 										Sources: []*string{
		// 											to.Ptr("source")},
		// 											TestConfigurations: []*string{
		// 												to.Ptr("tcp")},
		// 										}},
		// 										ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 									},
		// 							}},
		// 						}
	}
}
