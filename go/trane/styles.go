package trane

import "github.com/charmbracelet/lipgloss"


var (
	Black           = lipgloss.Color("#000000") // Default text
	Accent          = lipgloss.Color("147") // Headers, highlights, active
	SuccessGreen    = lipgloss.Color("#a6e3a1") // Success state
	ErrorRed        = lipgloss.Color("#f38ba8") // Error/failure
	WarningYellow   = lipgloss.Color("#f9e2af") // Warnings, in-progress
	InfoBlue        = lipgloss.Color("#89b4fa") // Info, links
	HighlightPink   = lipgloss.Color("#f5c2e7") // Secondary highlights
	TagCyan         = lipgloss.Color("#94e2d5") // Metadata, tags

	GrayLight       = lipgloss.Color("#6c7086") // Muted/inactive text
	GrayDark        = lipgloss.Color("#45475a") // Borders, dark UI accents
)


var inactiveTabStyle = lipgloss.NewStyle().
				Foreground(GrayLight).
				Padding(0, 1)


func activeTabStyle(state commandState) lipgloss.Style {
  style := lipgloss.NewStyle().
			Foreground(Black).
			Padding(0, 1)

  switch state {
    case running: return style.Background(WarningYellow)
    case success: return style.Background(SuccessGreen)
    case err:     return style.Background(ErrorRed)
  }

  return style
}
