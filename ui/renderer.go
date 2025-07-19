/*
Copyright © 2025 SHAKTHI WARNAKULASURIYA <shakthiraveen2002@gmail.com>

*/
package ui

import "github.com/charmbracelet/lipgloss"

var (
	// For successful AI responses
	ResponseStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			MarginTop(1).
			MarginBottom(1).
			Width(80)

	// For error messages
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("1")).
			Bold(true)

	// For subtle status/info
	InfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("7")).
			Italic(true)
)
