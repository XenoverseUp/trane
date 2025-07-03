package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	filePath = flag.String("file", "trane.json", "Path to command JSON file")
	vFlag    = flag.Bool("version", false, "Show CLI version")
	hFlag    = flag.Bool("help", false, "Show help")

	fShort = flag.String("f", "", "Shorthand for --file")
	vShort = flag.Bool("v", false, "Shorthand for --version")
	hShort = flag.Bool("h", false, "Shorthand for --help")
)

func printHelp() {
	fmt.Println(`Usage:
  trane [options] [alias]

Options:
  --file, -f     Path to command JSON file (default: ./trane.json)
  --version, -v  Show CLI version
  --help, -h     Show this help

Examples:
  trane --file=./my-commands.json
  trane build`)
}

func main() {
	flag.Parse()

	if *hFlag || *hShort {
		printHelp()
		return
	}
	if *vFlag || *vShort {
		fmt.Println(version)
		return
	}

	actualFile := *filePath
	if *fShort != "" {
		actualFile = *fShort
	}
	jsonPath, err := filepath.Abs(actualFile)
	if err != nil {
		log.Fatalf("Error resolving file path: %v", err)
	}

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}


	args := flag.Args()

	/* Alias Mode */

	if len(args) > 0 {
  	alias := args[0]
  	task, ok := config.Tasks[alias]
  	if !ok {
  		log.Fatalf("Alias %q not found in %s", alias, actualFile)
  	}
  	cwd := task.CWD
  	if cwd == "" {
  		cwd = "."
  	}
  	fmt.Printf("Parsed task [%s]: %s %v (cwd: %s)\n", alias, task.Command, task.Args, cwd)
	}

	/* TUI Mode */

	var tabs = []trane.Tab{}

	for _, task := range config.Tasks {

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
}
