/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
		fmt.Println("chat called")
		if len(args) == 0 {
			fmt.Println("❌ Please provide a message to send. Example: chat 'Tell me a joke'")
			os.Exit(1)
		}

		message := strings.Join(args, " ")

		// TODO: Implement the chat logic
		fmt.Println(message)
	},
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
