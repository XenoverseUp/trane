package main

import (
	"context"
	"os/exec"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
)

type CommandState int

const (
	running CommandState = iota
	err
	success
)

type tab struct {
	title      string
	command    string
	output     string
	state      CommandState
	cmd        *exec.Cmd
	cancelFunc context.CancelFunc
}

type model struct {
	tabs      []tab
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
