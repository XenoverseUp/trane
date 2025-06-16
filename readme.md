# Trane &middot; [![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/xenoverseup/trane/blob/main/LICENSE) [![npm version](https://img.shields.io/npm/v/@xenoverseup/trane)](https://www.npmjs.com/package/@xenoverseup/trane)

**Trane** is a modern CLI for running parallel tasks with a user-friendly TUI.
Think [Concurrently](https://www.npmjs.com/package/concurrently), but with a better interface, alias support, runtime validation, and developer-first ergonomics.

<p align="center">
  <img src="./docs/showcase.gif" alt="Trane showcase" width="700" />
</p>

## Install

```bash
npm install --global @xenoverseup/trane
```

## Usage

```bash
trane --help
```

### CLI Options

```
Usage
  $ trane [options] [alias]

Options
  --file, -f     Path to command config (default: ./trane.json)
  --list         List all available aliases
  --version, -v  Show CLI version
  --help, -h     Show CLI help

Examples
  $ trane                # Launch interactive TUI from trane.json
  $ trane --file=custom.json
  $ trane --list         # Show all available aliases
  $ trane build          # Run alias "build" defined in config
```

## Config File (`trane.json`)

Define tasks in a config file:

```json
[
	{
		"label": "Dev Server",
		"command": "npm",
		"args": ["run", "dev"],
		"alias": "dev"
	},
	{
		"label": "Type Check",
		"command": "tsc",
		"args": ["--watch"],
		"alias": "ts"
	}
]
```

### Types Safety for `trane.json`

If your editor supports it, enable auto-completion and type-checking:

```jsonc
// @ts-check
// @type {import('@xenoverseup/trane/trane-config').default}
[
	{
		"label": "Dev Server",
		"command": "npm",
		"args": ["run", "dev"],
		"alias": "dev"
	}
]
```

## Features

- **Config-first** setup using `trane.json`
- **Alias mode** for instant command execution
- **Runtime validation** of configs
- **Interactive TUI** with scrollable panes
- **`--list`** command to preview available tasks
- **Works out of the box** — no extra setup needed
- **Command logs** displayed side by side

## TODO Roadmap

- [ ] Export logs to file
- [ ] Timestamp each output
- [ ] Native buffer support
- [ ] Clear/reset output buffer
- [ ] Dual modes: interactive / headless
- [ ] Autocomplete alias names

---

[Muhammed Can Durmus](https://github.com/xenoverseup) · 2025
Licensed under the [MIT License](./LICENSE).
