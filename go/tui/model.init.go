package tui

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	cmds := []tea.Cmd{m.spinner.Tick}
	return tea.Batch(cmds...)
}
