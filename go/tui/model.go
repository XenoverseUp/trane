package tui

import (
	"context"
	"os/exec"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
)

type commandState int

const (
	running commandState = iota
	err
	success
)

type Tab struct {
	Title      string
	Command    string
	Args       []string
	Cwd        string
	output     string
	state      commandState

	cancelFunc context.CancelFunc
	pid        int
	cmd        *exec.Cmd
}

type model struct {
	tabs      []*Tab
	activeTab int
	width     int
	height    int
	spinner   spinner.Model
	viewport  viewport.Model
}

type outputMsg struct {
	index int
	line  string
}

type doneMsg struct {
	index int
	err   error
}
