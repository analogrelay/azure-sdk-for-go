//go:build go1.21
// +build go1.21

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/shared"
	"github.com/stretchr/testify/require"
)

func TestAssistants(t *testing.T) {
	if recording.GetRecordMode() != recording.LiveMode {
		t.Skip("https://github.com/Azure/azure-sdk-for-go/issues/22869")
	}

	testFn := func(t *testing.T, useAPIKey bool) {
		assistantClient := newStainlessTestClientWithOptions(t, azureOpenAI.Assistants.Endpoint, &stainlessTestClientOptions{
			UseAPIKey: useAPIKey,
		}).Beta.Assistants

		assistant, err := assistantClient.New(context.Background(), openai.BetaAssistantNewParams{
			Model:        shared.ChatModel(azureOpenAI.Assistants.Model),
			Instructions: openai.String("Answer questions in any manner possible"),
		})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, err := assistantClient.Delete(context.Background(), assistant.ID)
			require.NoError(t, err)
		})

		junkAssistant, err := assistantClient.New(context.Background(), openai.BetaAssistantNewParams{
			Model:        shared.ChatModel(azureOpenAI.Assistants.Model),
			Instructions: openai.String("Answer questions in any manner possible"),
		})
		require.NoError(t, err)

		// ensure there's at least one assistant "after" what we request
		t.Cleanup(func() {
			_, err := assistantClient.Delete(context.Background(), junkAssistant.ID)
			require.NoError(t, err)
		})

		const desc = "This is a newly updated description"

		// update the assistant's description
		{
			updatedAssistant, err := assistantClient.Update(context.Background(), assistant.ID, openai.BetaAssistantUpdateParams{
				Description: openai.String(desc),
			})
			require.NoError(t, err)
			require.Equal(t, desc, updatedAssistant.Description)
			require.Equal(t, assistant.ID, updatedAssistant.ID)
		}

		// get the same assistant back again
		{
			assistant2, err := assistantClient.Get(context.Background(), assistant.ID)
			require.NoError(t, err)
			require.Equal(t, assistant.ID, assistant2.ID)
			require.Equal(t, desc, assistant2.Description)
		}

		// listing assistants
		{
			pager, err := assistantClient.List(context.Background(), openai.BetaAssistantListParams{
				After: openai.String(assistant.ID),
				Limit: openai.Int(1),
			})
			require.NoError(t, err)

			var pages []openai.Assistant = pager.Data
			require.NotEmpty(t, pages)
		}
	}

	t.Run("APIKey", func(t *testing.T) {
		testFn(t, true)
	})

	t.Run("TokenCredential", func(t *testing.T) {
		testFn(t, false)
	})
}

func TestAssistantsThreads(t *testing.T) {
	if recording.GetRecordMode() != recording.LiveMode {
		t.Skip("https://github.com/Azure/azure-sdk-for-go/issues/22869")
	}

	testFn := func(t *testing.T, useAPIKey bool) {
		// NOTE: if you want to hit the OpenAI service instead...
		// assistantClient := openai.NewClient(
		// 	option.WithHeader("OpenAI-Beta", "assistants=v2"),
		// )
		beta := newStainlessTestClientWithOptions(t, azureOpenAI.Assistants.Endpoint, &stainlessTestClientOptions{
			UseAPIKey: useAPIKey,
		}).Beta
		assistantClient := beta.Assistants
		threadClient := beta.Threads

		assistant, err := assistantClient.New(context.Background(), openai.BetaAssistantNewParams{
			Model:        shared.ChatModel(azureOpenAI.Assistants.Model),
			Instructions: openai.String("Answer questions in any manner possible"),
		})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, err := assistantClient.Delete(context.Background(), assistant.ID)
			require.NoError(t, err)
		})

		thread, err := threadClient.New(context.Background(), openai.BetaThreadNewParams{})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, err := threadClient.Delete(context.Background(), thread.ID)
			require.NoError(t, err)
		})

		metadata := shared.MetadataParam{"hello": "world"}

		// update the thread
		{
			updatedThread, err := threadClient.Update(context.Background(), thread.ID, openai.BetaThreadUpdateParams{
				Metadata: metadata,
			})
			require.NoError(t, err)
			require.Equal(t, thread.ID, updatedThread.ID)
			require.Equal(t, shared.Metadata(metadata), updatedThread.Metadata)
		}

		// get the thread back
		{
			gotThread, err := threadClient.Get(context.Background(), thread.ID)
			require.NoError(t, err)
			require.Equal(t, thread.ID, gotThread.ID)
			require.Equal(t, shared.Metadata(metadata), gotThread.Metadata)
		}
	}

	t.Run("APIKey", func(t *testing.T) {
		testFn(t, true)
	})

	t.Run("TokenCredential", func(t *testing.T) {
		testFn(t, false)
	})
}

func TestAssistantRun(t *testing.T) {
	if recording.GetRecordMode() != recording.LiveMode {
		t.Skip("https://github.com/Azure/azure-sdk-for-go/issues/22869")
	}

	testFn := func(t *testing.T, useAPIKey bool) {
		// NOTE: if you want to hit the OpenAI service instead...
		// assistantClient := openai.NewClient(
		// 	option.WithHeader("OpenAI-Beta", "assistants=v2"),
		// )

		client := newStainlessTestClientWithOptions(t, azureOpenAI.Assistants.Endpoint, &stainlessTestClientOptions{
			UseAPIKey: useAPIKey,
		})

		// (this is the test, verbatim, from openai-go: https://github.com/openai/openai-go/blob/main/examples/assistant-streaming/main.go)

		ctx := context.Background()

		// Create an assistant
		println("Create an assistant")
		assistant, err := client.Beta.Assistants.New(ctx, openai.BetaAssistantNewParams{
			Name:         openai.String("Math Tutor"),
			Instructions: openai.String("You are a personal math tutor. Write and run code to answer math questions."),
			Tools: []openai.AssistantToolUnionParam{{
				OfCodeInterpreter: &openai.CodeInterpreterToolParam{},
			}},
			Model: shared.ChatModel(azureOpenAI.Assistants.Model),
		})

		if err != nil {
			panic(err)
		}

		// Create a thread
		println("Create an thread")
		thread, err := client.Beta.Threads.New(ctx, openai.BetaThreadNewParams{})
		if err != nil {
			panic(err)
		}

		// Create a message in the thread
		println("Create a message")
		_, err = client.Beta.Threads.Messages.New(ctx, thread.ID, openai.BetaThreadMessageNewParams{
			Role: openai.BetaThreadMessageNewParamsRoleAssistant,
			Content: openai.BetaThreadMessageNewParamsContentUnion{
				OfArrayOfContentParts: []openai.MessageContentPartParamUnion{{
					OfText: &openai.TextContentBlockParam{
						Text: "I need to solve the equation `3x + 11 = 14`. Can you help me?",
					},
				}},
			},
		})
		if err != nil {
			panic(err)
		}

		// Create a run
		println("Create a run")
		stream := client.Beta.Threads.Runs.NewStreaming(ctx, thread.ID, openai.BetaThreadRunNewParams{
			AssistantID:  assistant.ID,
			Instructions: openai.String("Please address the user as Jane Doe. The user has a premium account."),
		})

		for stream.Next() {
			evt := stream.Current()
			println(fmt.Sprintf("%T", evt.Data))
		}
	}
	t.Run("APIKey", func(t *testing.T) {
		testFn(t, true)
	})

	t.Run("TokenCredential", func(t *testing.T) {
		testFn(t, false)
	})
}
