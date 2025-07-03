package main

import (
	"context"
	"os/exec"

	"github.com/charmbracelet/bubbles/spinner"
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
	spinner   spinner.Model
	width     int
}

type outputMsg struct {
	index int
	line  string
}

type doneMsg struct {
	index int
	err   error
}
