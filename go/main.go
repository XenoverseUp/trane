package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

  switch msg := msg.(type) {
    case tea.KeyMsg:
      switch msg.String() {
        case "ctrl+c":
          return m, tea.Quit
      }
  }

  return m, nil
}

func (m model) View() string {
  return "Hello World!"
}

func main() {
  p:= tea.NewProgram(model{}, tea.WithAltScreen())

  if _, err := p.Run(); err != nil {
    fmt.Print("Error!")
  }
}
