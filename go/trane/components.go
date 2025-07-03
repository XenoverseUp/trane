package trane

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)


func (m model) renderHeader() (string, int) {
	traneStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(lightPurple).
		Padding(0, 1).
		MarginLeft(1).
		Border(lipgloss.MarkdownBorder(), false, true, false, false).
		BorderForeground(lipgloss.Color("#555555"))

	traneText := traneStyle.Render("⛬ TRANE")

	var tabLabels []string
	for i, t := range m.tabs {
		label := fmt.Sprintf("(%d) %s %s", i+1, t.Title, getStateIcon(t.state))
		if i == m.activeTab {
			tabLabels = append(tabLabels, activeTabStyle.Render(label))
		} else {
			tabLabels = append(tabLabels, zone.Mark(fmt.Sprintf("tab:%d", i), inactiveTabStyle.Render(label)))
		}
	}

	tabsContent := lipgloss.JoinHorizontal(lipgloss.Left, tabLabels[0])
  for _, label := range tabLabels[1:] {
      tabsContent += " " + label
  }

	remainingWidth := max(m.width-lipgloss.Width(traneText)-lipgloss.Width(tabsContent), 0)

	headerRow := lipgloss.JoinHorizontal(
		lipgloss.Left,
		traneText,
		lipgloss.NewStyle().Width(remainingWidth).Render(""),
		tabsContent,
	)

	headerWithBorder := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lightPurple).
		Width(m.width).
		Render(headerRow)

	return headerWithBorder, lipgloss.Height(headerWithBorder)
}

func (m *model) renderViewport() {
	_, headerHeight := m.renderHeader()
	_, infoBarHeight := m.renderInfoBar()

	viewportHeight := max(m.height-headerHeight-infoBarHeight, 1)
	m.viewport = viewport.New(m.width, viewportHeight)
	m.updateViewportContent()
}

func (m *model) updateViewportContent() {
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
	m.viewport.SetContent(content.String())
}


func (m model) renderInfoBar() (string, int) {
	infoBar := "←/→ or 1-9 to switch tabs, 'q' to quit."
	infoBarStyled := lipgloss.NewStyle().
		Foreground(lightGray).
		Padding(0, 1).
		Render(infoBar)

	return infoBarStyled, lipgloss.Height(infoBarStyled)
}
