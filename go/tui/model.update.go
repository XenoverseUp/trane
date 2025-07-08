package tui

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			for i := range m.tabs {
				if m.tabs[i].cancelFunc != nil {
					m.tabs[i].cancelFunc()
				}
			}
			return m, tea.Quit
		case "1", "2", "3", "4":
			idx := int(msg.String()[0] - '1')
			if idx >= 0 && idx < len(m.tabs) {
				m.activeTab = idx
			}
		case "right":
			m.activeTab = (m.activeTab + 1) % len(m.tabs)
		case "left":
			m.activeTab = ((m.activeTab - 1) + len(m.tabs)) % len(m.tabs)
		case "j":
			tab := m.tabs[m.activeTab]
			tab.output += "Hello World! " + strconv.Itoa(rand.Int()) + "\n"
		}

	case tea.MouseMsg:
		if msg.Action == tea.MouseAction(tea.MouseButtonWheelUp) {
			m.viewport.ScrollUp(1)
			return m, nil
		}
		if msg.Action == tea.MouseAction(tea.MouseButtonWheelDown) {
			m.viewport.ScrollDown(1)
			return m, nil
		}
		if msg.Action == tea.MouseActionRelease && msg.Button == tea.MouseButtonLeft {
			for i := range m.tabs {
				if zone.Get(fmt.Sprintf("tab:%d", i)).InBounds(msg) {
					m.activeTab = i
				}
			}
			return m, nil
		}

	case outputMsg:
		tab := m.tabs[msg.index]
		tab.output += msg.line + "\n"
		if msg.index == m.activeTab {
			m.updateViewportContent()
		}

	case doneMsg:
		tab := m.tabs[msg.index]
		if msg.err != nil {
			tab.state = err
			tab.output += fmt.Sprintf("\nError: %v\n", msg.err)
		} else {
			tab.state = success
		}
		if msg.index == m.activeTab {
			m.updateViewportContent()
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.renderViewport()
	}

	m.updateViewportContent()

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
