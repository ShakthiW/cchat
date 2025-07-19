package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const openaiEndpoint = "https://api.openai.com/v1/chat/completions"
const openaiModel = "gpt-4o-mini"

type OpenAIRequest struct {
	Model    string          `json:"model"`
	Messages []OpenAIMessage `json:"messages"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message OpenAIMessage `json:"message"`
	} `json:"choices"`
}

func GetOpenAIResponse(message string, apiKey string) (string, error) {
	reqBody := OpenAIRequest{
		Model: openaiModel,
		Messages: []OpenAIMessage{
			{
				Role:    "user",
				Content: message,
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", openaiEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("OpenAI API error: %s", string(errorBody))
	}

	var openAIResponse OpenAIResponse
	if err := json.NewDecoder(res.Body).Decode(&openAIResponse); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(openAIResponse.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return openAIResponse.Choices[0].Message.Content, nil
}
