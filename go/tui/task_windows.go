//go:build windows
// +build windows

package tui

import (
	"bufio"
	"context"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"

	tea "github.com/charmbracelet/bubbletea"
)

func run(tabs []*Tab, program *tea.Program) {
	for i, tab := range tabs {
		i := i
		tab := tab

		go func() {
			ctx, cancel := context.WithCancel(context.Background())

			cmd := exec.CommandContext(ctx, tab.Command, tab.Args...)
			cmd.Dir = tab.Cwd

			cmd.SysProcAttr = &syscall.SysProcAttr{
				HideWindow:    true,
				CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
			}

			stdoutPipe, err := cmd.StdoutPipe()
			if err != nil {
				program.Send(doneMsg{index: i, err: err})
				cancel()
				return
			}
			stderrPipe, err := cmd.StderrPipe()
			if err != nil {
				program.Send(doneMsg{index: i, err: err})
				cancel()
				return
			}

			if err := cmd.Start(); err != nil {
				program.Send(doneMsg{index: i, err: err})
				cancel()
				return
			}

			tab.cmd = cmd

			tab.cancelFunc = func() {
				cancel()
				windows.GenerateConsoleCtrlEvent(windows.CTRL_BREAK_EVENT, uint32(cmd.Process.Pid))
				cmd.Process.Kill()
			}

			go func () {
				scanner := bufio.NewScanner(stdoutPipe)
				for scanner.Scan() {
					program.Send(outputMsg{index: i, line: scanner.Text()})
				}
			}()
			go func () {
				scanner := bufio.NewScanner(stderrPipe)
				for scanner.Scan() {
					program.Send(outputMsg{index: i, line: scanner.Text()})
				}
			}()

			err = cmd.Wait()
			program.Send(doneMsg{index: i, err: err})
		}()
	}
}
