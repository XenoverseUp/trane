// TODO: Sanitize Full Screen
// TODO: Stable ordering
// TODO: Handle quitting
// TODO: Graceful errors
// TODO: ASCII art help

package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"

	zone "github.com/lrstanley/bubblezone"
)

var _ tea.Model = model{}

func CreateTrane(tabs []Tab) {
	tabsPtrs := make([]*Tab, len(tabs))
	for i := range tabs {
		tabsPtrs[i] = &tabs[i]
	}


  	s := spinner.New()
	s.Spinner = spinner.MiniDot

	var m = model{
		tabs:      tabsPtrs,
		activeTab: 0,
		spinner:   s,
	}

  	zone.NewGlobal()

	program := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())


	run(m.tabs, program)

	_, err := program.Run()

	if err != nil {
		fmt.Println("Error:", err)
	}
}
