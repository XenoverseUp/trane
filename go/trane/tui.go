package trane

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"

	zone "github.com/lrstanley/bubblezone"
)

var _ tea.Model = model{}


/*** Init ***/

func (m model) Init() tea.Cmd {
	cmds := []tea.Cmd{m.spinner.Tick}

	return tea.Batch(cmds...)
}

/*** Update ***/

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

	case tea.MouseMsg:
		if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft {
			return m, nil
		}

		for i := range m.tabs {
  		if zone.Get(fmt.Sprintf("tab:%d", i)).InBounds(msg) {
  			m.activeTab = i
  		}
		}

		return m, nil


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
		m.renderViewport()
	}

	tab := m.tabs[m.activeTab]

	var content strings.Builder
	switch tab.state {
 	  case running:
  		content.WriteString(fmt.Sprintf("%s CMD: `%s`\n", m.spinner.View(), tab.Command))
   	case success:
  		content.WriteString(fmt.Sprintf("Finished: %s\n", tab.Command))
   	case err:
  		content.WriteString(fmt.Sprintf("Error: %s\n", tab.Command))
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

/*** View ***/

func (m model) View() string {
  var b strings.Builder

	header, _ := m.renderHeader()
	infoBar, _ := m.renderInfoBar()

	b.WriteString(header)
	b.WriteString("\n")

	b.WriteString(m.viewport.View())
	b.WriteString("\n")

	b.WriteString(infoBar)

	return zone.Scan(b.String())
}


func CreateTrane(tabs []Tab) {
  s := spinner.New()
	s.Spinner = spinner.Meter

	var m = model{
		tabs:      tabs,
		activeTab: 0,
		spinner:   s,
	}

  zone.NewGlobal()

	program := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	_, err := program.Run()

	if err != nil {
		fmt.Println("Error:", err)
	}
}
