// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package main

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/v2/internal/eh/stress/tests"
)

func main() {
	tests := []struct {
		name string
		fn   func(ctx context.Context) error
	}{
		{name: "batch", fn: tests.BatchStressTester},
		{name: "balance", fn: tests.BalanceTester},
		{name: "multibalance", fn: tests.MultiBalanceTester},
		{name: "processor", fn: tests.ProcessorStressTester},
	}

	sort.Slice(tests, func(i, j int) bool {
		return tests[i].name < tests[j].name
	})

	if len(os.Args) < 2 {
		fmt.Printf("Usage: stress <scenario-name>\n")

		fmt.Printf("Scenarios:\n")

		for _, test := range tests {
			fmt.Printf("  %s\n", test.name)
		}

		os.Exit(1)
	}

	testName := os.Args[1]

	for _, test := range tests {
		if test.name == testName {
			if err := test.fn(context.Background()); err != nil {
				fmt.Printf("ERROR: %s\n", err)
				os.Exit(1)
			}

			os.Exit(0)
		}
	}

	fmt.Printf("No test with name %s", testName)
	os.Exit(1)
}
