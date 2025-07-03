package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"

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

func (m *model) InitViewport() {
	_, headerHeight := m.renderHeader()
	_, infoBarHeight := m.renderInfoBar()

	viewportHeight := max(m.height - headerHeight - infoBarHeight, 1)
	m.viewport = viewport.New(m.width, viewportHeight)
}


func (m model) Init() tea.Cmd {
	cmds := []tea.Cmd{m.spinner.Tick}

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
		m.height = msg.Height
		m.InitViewport()
	}

	tab := m.tabs[m.activeTab]

	var content strings.Builder
	switch tab.state {
 	  case running:
  		content.WriteString(fmt.Sprintf("%s Running: %s\n", m.spinner.View(), tab.command))
   	case success:
  		content.WriteString(fmt.Sprintf("Finished: %s\n", tab.command))
   	case err:
  		content.WriteString(fmt.Sprintf("Error: %s\n", tab.command))
	}

	if tab.output != "" {
		content.WriteString("\n" + tab.output)
	}

	var cmd tea.Cmd

	m.viewport.SetContent(content.String())
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
  var b strings.Builder

	header, _ := m.renderHeader()
	infoBar, _ := m.renderInfoBar()

	b.WriteString(header)
	b.WriteString("\n")

	b.WriteString(m.viewport.View())
	b.WriteString("\n")

	b.WriteString(infoBar)

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
