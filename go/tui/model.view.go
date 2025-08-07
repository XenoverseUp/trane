package tui

import (
	"strings"

	zone "github.com/lrstanley/bubblezone"
)

func (m model) View() string {
	var b strings.Builder

	header, _ := m.renderHeader()
	infoBar, _ := m.renderInfoBar()

	zone.Scan(header)

	b.WriteString(header)
	b.WriteString("\n")

	b.WriteString(m.viewport.View())
	b.WriteString("\n")

	b.WriteString(infoBar)

	return b.String()
}
