package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xenoverseup/trane/tui"
)

const version = "1.0.2"

type Task struct {
	Label   string   `json:"label"`
	Command string   `json:"command"`
	Args    []string `json:"args"`
	CWD     string   `json:"cwd,omitempty"`
}

type Config struct {
	Tasks map[string]Task `json:"tasks"`
}

var (
	file        string
	showVersion bool
	cfg         Config
	jsonPath    string
)

var rootCmd = &cobra.Command{
	Use:   "trane",
	Short: "Trane - Run aliases or launch the interactive TUI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println(version)
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		// TUI mode
		configDir := filepath.Dir(jsonPath)

		var tabs []tui.Tab
		for _, task := range cfg.Tasks {
			cwd := task.CWD
			if cwd == "" {
				cwd = configDir
			} else if !filepath.IsAbs(cwd) {
				cwd = filepath.Join(configDir, cwd)
			}

			tab := tui.Tab{
				Title:   task.Label,
				Command: task.Command,
				Args:    task.Args,
				Cwd:     cwd,
			}

			tabs = append(tabs, tab)
		}

		tui.CreateTrane(tabs)

	},
}

var runCmd = &cobra.Command{
	Use:   "run <alias>",
	Short: "Run a single alias from the config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		alias := args[0]
		task, ok := cfg.Tasks[alias]
		if !ok {
			fmt.Printf("Alias %q not found in %s", alias, jsonPath)
			os.Exit(0)
		}

		cwd := task.CWD
		if cwd == "" {
			cwd = filepath.Dir(jsonPath)
		} else if !filepath.IsAbs(cwd) {
			cwd = filepath.Join(filepath.Dir(jsonPath), cwd)
		}

		fmt.Printf("Parsed task [%s]: %s %v (cwd: %s)\n", alias, task.Command, task.Args, cwd)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func loadConfig() {
	var err error
	jsonPath, err = filepath.Abs(file)
	if err != nil {
		fmt.Printf("Error resolving the file path `%s`.\n", jsonPath)
		os.Exit(0)
	}

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Printf("Config find the config file `%s`.\n", jsonPath)
		os.Exit(0)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(0)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "trane.json", "Path to command JSON file")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show CLI version")

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
