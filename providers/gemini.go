package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const geminiEndpoint = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key="
const geminiModel = "gemini-2.0-flash"

type GeminiRequest struct {
	Model    string          `json:"model"`
	Messages []GeminiMessage `json:"messages"`
}

type GeminiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GeminiResponse struct {
	Choices []struct {
		Message GeminiMessage `json:"message"`
	} `json:"choices"`
}

func GetGeminiResponse(message string, apiKey string) (string, error) {
	reqBody := GeminiRequest{
		Model: geminiModel,
		Messages: []GeminiMessage{
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

	req, err := http.NewRequest("POST", geminiEndpoint, bytes.NewBuffer(jsonData))
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
		return "", fmt.Errorf("gemini API error: %s", string(errorBody))
	}

	var geminiResponse GeminiResponse
	if err := json.NewDecoder(res.Body).Decode(&geminiResponse); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(geminiResponse.Choices) == 0 {
		return "", fmt.Errorf("no response from Gemini")
	}

	return geminiResponse.Choices[0].Message.Content, nil
}
