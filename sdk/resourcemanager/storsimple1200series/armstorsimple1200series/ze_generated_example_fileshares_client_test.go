//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesListByFileServer.json
func ExampleFileSharesClient_NewListByFileServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewFileSharesClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByFileServerPager("HSDK-ARCSX4MVKZ",
		"HSDK-ARCSX4MVKZ",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesGet.json
func ExampleFileSharesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewFileSharesClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"HSDK-4XY4FI2IVG",
		"HSDK-4XY4FI2IVG",
		"Auto-TestFileShare1",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesCreateOrUpdate.json
func ExampleFileSharesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewFileSharesClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"HSDK-4XY4FI2IVG",
		"HSDK-4XY4FI2IVG",
		"Auto-TestFileShare1",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		armstorsimple1200series.FileShare{
			Name: to.Ptr("Auto-TestFileShare1"),
			Properties: &armstorsimple1200series.FileShareProperties{
				Description:                to.Ptr("Demo FileShare for SDK Test Tiered"),
				AdminUser:                  to.Ptr("fareast\\idcdlslb"),
				DataPolicy:                 to.Ptr(armstorsimple1200series.DataPolicyTiered),
				MonitoringStatus:           to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
				ShareStatus:                to.Ptr(armstorsimple1200series.ShareStatusOnline),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesDelete.json
func ExampleFileSharesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewFileSharesClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"HSDK-DMNJB2PET0",
		"HSDK-DMNJB2PET0",
		"Auto-TestFileShare2",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesListMetrics.json
func ExampleFileSharesClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewFileSharesClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListMetricsPager("HSDK-DMNJB2PET0",
		"HSDK-DMNJB2PET0",
		"Auto-TestFileShare2",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		&armstorsimple1200series.FileSharesClientListMetricsOptions{Filter: to.Ptr("startTime%20ge%20'2018-08-10T18:30:00Z'%20and%20endTime%20le%20'2018-08-11T18:30:00Z'")})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesListMetricDefinition.json
func ExampleFileSharesClient_NewListMetricDefinitionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewFileSharesClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListMetricDefinitionPager("HSDK-DMNJB2PET0",
		"HSDK-DMNJB2PET0",
		"Auto-TestFileShare2",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesListByDevice.json
func ExampleFileSharesClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewFileSharesClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByDevicePager("HSDK-4XY4FI2IVG",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}