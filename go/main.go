package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xenoverseup/trane"
)

const version = "0.0.13"

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
	file       string
	showVersion bool
	cfg        Config
	jsonPath   string
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
		var tabs []trane.Tab
		for _, task := range cfg.Tasks {
			cwd := task.CWD
			if cwd == "" {
				cwd = "."
			}
			tabs = append(tabs, trane.Tab{
				Title:   task.Label,
				Command: task.Command,
				Args:    task.Args,
				Cwd:     cwd,
			})
		}

		trane.CreateTrane(tabs)
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
			log.Fatalf("Alias %q not found in %s", alias, jsonPath)
		}

		cwd := task.CWD
		if cwd == "" {
			cwd = "."
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
		log.Fatalf("Error resolving file path: %v", err)
	}

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("Error reading file %s: %v", jsonPath, err)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
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
