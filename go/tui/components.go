package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/muesli/reflow/wordwrap"
)

func (m model) renderHeader() (string, int) {
	traneStyle := lipgloss.NewStyle().
		Foreground(Black).
		Background(Accent).
		Padding(0, 1).
		Border(lipgloss.MarkdownBorder(), false, true, false, false).
		BorderForeground(GrayLight)

	traneText := traneStyle.Render("⛬ TRANE")

	var tabLabels []string
	for i, t := range m.tabs {
		label := fmt.Sprintf("(%d) %s %s", i+1, t.Title, getStateIcon(t.state, &m.spinner))
		if i == m.activeTab {
			tabLabels = append(tabLabels, activeTabStyle(t.state).Render(label))
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
		BorderForeground(Accent).
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
	var content string

	if tab.output != "" {
		content = wordwrap.String(tab.output, m.viewport.Width)
	}

	m.viewport.SetContent(content)
}

func (m model) renderInfoBar() (string, int) {
	infoBar := "←/→ or 1-9 to switch tabs, 'q' to quit."
	infoBarStyled := lipgloss.NewStyle().
		Foreground(GrayLight).
		Padding(0, 1).
		Border(lipgloss.MarkdownBorder(), true, false, false, false).
		BorderForeground(GrayDark).
		Width(m.width).
		Render(infoBar)

	return infoBarStyled, lipgloss.Height(infoBarStyled)
}
