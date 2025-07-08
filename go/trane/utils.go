package trane

import "github.com/charmbracelet/bubbles/spinner"

func getStateIcon(state commandState, spinner *spinner.Model) string {
  switch state {
    case running: return spinner.View()
    case success: return "✔"
    case err:     return "✖"
  }

  return ""
}
