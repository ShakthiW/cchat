/*
Copyright © 2025 SHAKTHI WARNAKULASURIYA <shakthiraveen2002@gmail.com>
*/
package cmd

import (
	"chatcli/config"
	"chatcli/providers"
	"chatcli/ui"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with an LLM using OpenAI, Gemini, or Claude",
	Long: `Send a prompt to an LLM from the terminal.

Examples:
  chat "Tell me a joke"
  chat --gemini "What's the weather?"
  chat --claude "Summarize this paragraph..."`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ Please provide a message to send. Example: chat 'Tell me a joke'")
			os.Exit(1)
		}

		message := strings.Join(args, " ")

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("❌ Could not load config. Have you run 'chat init'?")
			os.Exit(1)
		}

		var provider string
		if cmd.Flag("openai").Changed {
			provider = "openai"
		} else if cmd.Flag("gemini").Changed {
			provider = "gemini"
		} else if cmd.Flag("claude").Changed {
			provider = "claude"
		} else {
			provider = cfg.DefaultProvider
		}

		var response string

		switch provider {
		case "openai":
			response, err = getOpenAIResponse(message, cfg)
			if err != nil {
				fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("❌ Failed to get OpenAI response: %v", err)))
				os.Exit(1)
			}
		case "gemini":
			response, err = getGeminiResponse(message, cfg)
			if err != nil {
				fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("❌ Failed to get Gemini response: %v", err)))
				os.Exit(1)
			}
		case "claude":
			fmt.Println(ui.InfoStyle.Render("❌ Claude is not supported yet"))
			os.Exit(1)
		default:
			fmt.Println(ui.ErrorStyle.Render("❌ No provider selected. Use --openai, --gemini, or --claude"))
			os.Exit(1)
		}

		fmt.Println(ui.ResponseStyle.Render(response))
	},
}

func getOpenAIResponse(message string, cfg *config.Config) (string, error) {
	if cfg.OpenAIKey == "" {
		fmt.Println(ui.ErrorStyle.Render("❌ OpenAI key not set. Use `chat init --openai-key=...`"))
		os.Exit(1)
	}

	response, err := providers.GetOpenAIResponse(message, cfg.OpenAIKey)
	if err != nil {
		return "", fmt.Errorf("failed to get OpenAI response: %w", err)
	}

	return response, nil
}

func getGeminiResponse(message string, cfg *config.Config) (string, error) {
	if cfg.GeminiKey == "" {
		fmt.Println(ui.ErrorStyle.Render("❌ Gemini key not set. Use `chat init --gemini-key=...`"))
		os.Exit(1)
	}

	response, err := providers.GetGeminiResponse(message, cfg.GeminiKey)
	if err != nil {
		return "", fmt.Errorf("failed to get Gemini response: %w", err)
	}

	return response, nil
}

func init() {
	rootCmd.AddCommand(chatCmd)

	chatCmd.Flags().BoolP("openai", "o", false, "Chat with OpenAI")
	chatCmd.Flags().BoolP("gemini", "g", false, "Chat with Gemini")
	chatCmd.Flags().BoolP("claude", "c", false, "Chat with Claude")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
