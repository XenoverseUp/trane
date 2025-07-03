package trane

import "github.com/charmbracelet/lipgloss"

var (
  purple    = lipgloss.Color("99")
  gray      = lipgloss.Color("245")
  lightGray = lipgloss.Color("241")
  lightPurple = lipgloss.Color("147")


	activeTabStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#00FFAA")).
			Padding(0,1 )

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#888888")).
				Padding(0, 1)
)
