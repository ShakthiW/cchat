/*
Copyright © 2025 SHAKTHI WARNAKULASURIYA <shakthiraveen2002@gmail.com>
*/
package providers

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

var PromptTemplate = `
	You are a helpful assistant that lives in the terminal. So keep the response short and concise.
	Don't use styling through markdown.

	User Question: %s
`

func GetGeminiResponse(message string, apiKey string) (string, error) {
	ctx := context.Background()

	clientConfig := &genai.ClientConfig{
		APIKey: apiKey,
	}
	client, err := genai.NewClient(ctx, clientConfig)
	if err != nil {
		return "", fmt.Errorf("failed to create genai client: %w", err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(fmt.Sprintf(PromptTemplate, message)),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	return result.Text(), nil
}
