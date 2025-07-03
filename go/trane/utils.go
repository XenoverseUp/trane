package trane

func getStateIcon(state commandState) string {
  switch state {
    case running: return "●"
    case success: return "✔"
    case err:     return "✖"
  }

  return ""
}

func getStateColor(state commandState) string {
  switch state {
    case running: return "●"
    case success: return "✔"
    case err:     return "✖"
  }

  return ""
}
