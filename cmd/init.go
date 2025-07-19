/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"chatcli/config"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Configure API keys and default provider for chatcli",
	Long: `Use this command to set or update your API keys and default LLM provider.
You can update one or more fields at a time using flags.

Examples:
  chat init --openai-key=sk-xxx
  chat init --default-provider=openai
  chat init --claude-key=abc123 --default-provider=claude`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			cfg = &config.Config{}
		}

		openaiKey, _ := cmd.Flags().GetString("openai-key")
		geminiKey, _ := cmd.Flags().GetString("gemini-key")
		claudeKey, _ := cmd.Flags().GetString("claude-key")
		defaultProvider, _ := cmd.Flags().GetString("default-provider")

		// Set only if provided
		if openaiKey != "" {
			cfg.OpenAIKey = openaiKey
		}
		if geminiKey != "" {
			cfg.GeminiKey = geminiKey
		}
		if claudeKey != "" {
			cfg.ClaudeKey = claudeKey
		}
		if defaultProvider != "" {
			defaultProvider = strings.ToLower(defaultProvider)
			if defaultProvider != "openai" && defaultProvider != "gemini" && defaultProvider != "claude" {
				fmt.Println("❌ Invalid default provider. Choose from: openai, gemini, claude")
				return
			}
			cfg.DefaultProvider = defaultProvider
		}

		// save updated config
		err = config.SaveConfig(cfg.OpenAIKey, cfg.GeminiKey, cfg.ClaudeKey, cfg.DefaultProvider)
		if err != nil {
			fmt.Println("❌ Error saving config:", err)
			return
		}

		fmt.Println("✅ Config updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String("openai-key", "", "Set OpenAI API key")
	initCmd.Flags().String("gemini-key", "", "Set Gemini API key")
	initCmd.Flags().String("claude-key", "", "Set Claude API key")
	initCmd.Flags().String("default-provider", "", "Set default provider")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
