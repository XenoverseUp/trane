package trane

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)


func (m model) renderHeader() (string, int) {
	traneStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#A6ABF1")).
		Padding(0, 1).
		MarginLeft(1).
		Border(lipgloss.MarkdownBorder(), false, true, false, false).
		BorderForeground(lipgloss.Color("#555555"))

	traneText := traneStyle.Render("⛬ TRANE")

	var tabLabels []string
	for i, t := range m.tabs {
		label := fmt.Sprintf("%d) %s ", i+1, t.Title)
		if i == m.activeTab {
			label += "✔ "
			tabLabels = append(tabLabels, activeTabStyle.Render(label))
		} else {
			tabLabels = append(tabLabels, zone.Mark(fmt.Sprintf("tab:%d", i), inactiveTabStyle.Render(label)))
		}
	}
	tabsContent := lipgloss.JoinHorizontal(lipgloss.Left, tabLabels...)

	remainingWidth := max(m.width-lipgloss.Width(traneText)-lipgloss.Width(tabsContent), 0)

	headerRow := lipgloss.JoinHorizontal(
		lipgloss.Left,
		traneText,
		lipgloss.NewStyle().Width(remainingWidth).Render(""),
		tabsContent,
	)

	headerWithBorder := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color("#A6ABF1")).
		Width(m.width).
		Render(headerRow)

	return headerWithBorder, lipgloss.Height(headerWithBorder)
}

func (m model) renderInfoBar() (string, int) {
	infoBar := "←/→ or 1-9 to switch tabs, 'q' to quit."
	infoBarStyled := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Padding(0, 1).
		Render(infoBar)

	return infoBarStyled, lipgloss.Height(infoBarStyled)
}
