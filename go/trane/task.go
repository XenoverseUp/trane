package trane

import (
	"bufio"
	"context"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func run(tabs []*Tab, program *tea.Program) {
	for i, tab := range tabs {
		i := i
		tab := tab

		go func() {
			ctx, cancel := context.WithCancel(context.Background())
			tab.cancelFunc = cancel

			cmd := exec.CommandContext(ctx, tab.Command, tab.Args...)
			cmd.Dir = tab.Cwd

			stdoutPipe, err := cmd.StdoutPipe()
			if err != nil {
				program.Send(doneMsg{index: i, err: err})
				return
			}

			stderrPipe, err := cmd.StderrPipe()
			if err != nil {
				program.Send(doneMsg{index: i, err: err})
				return
			}

			if err := cmd.Start(); err != nil {
				program.Send(doneMsg{index: i, err: err})
				return
			}

			go func() {
				scanner := bufio.NewScanner(stdoutPipe)
				for scanner.Scan() {
					program.Send(outputMsg{index: i, line: scanner.Text()})
				}
			}()

			go func() {
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
