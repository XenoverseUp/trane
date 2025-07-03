package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var _ tea.Model = model{}

func initialModel() model {
	tabs := []tab{
		{title: "Echo", command: "echo Hello"},
		{title: "List", command: "ls -al"},
		{title: "Sleep", command: "sleep 2"},
		{title: "Brew", command: "brew help"},
	}

	s := spinner.New()
	s.Spinner = spinner.Dot

	m := model{
		tabs:      tabs,
		activeTab: 0,
		spinner:   s,
	}

	return m
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.spinner.Tick)

	// Start all commands
	// for i := range m.tabs {
	// 	cmds = append(cmds, runCommand(i, m.tabs[i].command))
	// }

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
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
		}

	case outputMsg:
		tab := &m.tabs[msg.index]
		tab.output += msg.line + "\n"

	case doneMsg:
		tab := &m.tabs[msg.index]

		if msg.err != nil {
			tab.state = err
		} else {
			tab.state = success
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)

	case tea.WindowSizeMsg:
		m.width = msg.Width
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	// --- Header Section ---

	traneStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFAA")).
		Background(lipgloss.Color("#000000")).
		Padding(0, 1).
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(lipgloss.Color("#555555"))

	traneText := traneStyle.Render("⛬ Trane")

	// Render tabs
	var tabLabels []string
	for i, t := range m.tabs {
		label := fmt.Sprintf("%d) %s ", i+1, t.title)
		if i == m.activeTab {

			label += "✔ "
			tabLabels = append(tabLabels, activeTabStyle.Render(label))
		} else {
			tabLabels = append(tabLabels, inactiveTabStyle.Render(label))
		}
	}
	tabsContent := lipgloss.JoinHorizontal(lipgloss.Left, tabLabels...)


	remainingWidth := max(m.width - lipgloss.Width(traneText) - lipgloss.Width(tabsContent), 0)

	headerRow := lipgloss.JoinHorizontal(
		lipgloss.Left,
		traneText,
		lipgloss.NewStyle().Width(remainingWidth).Render(""),
		tabsContent,
	)

	headerWithBorder := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color("#555555")).
		Width(m.width).
		Render(headerRow)

	b.WriteString(headerWithBorder)
	b.WriteString("\n")

	// --- Main Content Section ---

	tab := m.tabs[m.activeTab]
	if tab.state == running {
		b.WriteString(fmt.Sprintf("%s Running: %s\n", m.spinner.View(), tab.command))
	} else if tab.state == success {
		b.WriteString(fmt.Sprintf("Finished: %s\n", tab.command))
	} else if tab.state == err {
		b.WriteString(fmt.Sprintf("Error: %s\n", tab.command))
	}

	if tab.output != "" {
		b.WriteString("\n" + tab.output)
	}

	b.WriteString("\n\n←/→ or 1-9 to switch tabs, 'q' to quit.")
	return b.String()
}

var (
	activeTabStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#00FFAA")).
			Padding(0, 1).
			BorderForeground(lipgloss.Color("#00FFAA"))

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#888888")).
				Padding(0, 1)
)

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
