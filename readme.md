# Trane &middot; [![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/xenoverseup/trane/blob/main/LICENSE) [![npm version](https://img.shields.io/npm/v/@xenoverseup/trane?color=blue)](https://www.npmjs.com/package/@xenoverseup/trane) [![Downloads/week](https://img.shields.io/npm/dw/@xenoverseup/trane)](https://www.npmjs.com/package/@xenoverseup/trane) [![Stars](https://img.shields.io/github/stars/xenoverseup/trane?style=social)](https://github.com/xenoverseup/trane/stargazers)

**Trane** is a modern CLI for running parallel tasks with a user-friendly TUI.
Think [Concurrently](https://www.npmjs.com/package/concurrently), but with a better interface, alias support, runtime validation, and developer-first ergonomics. Create `trane.json` file in the root directory. Define commands and the working directories. Then simply run `trane`.

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
{
  "tasks": {
    "server": {
      "label": "Server",
      "command": "bun",
      "args": ["run", "dev"]
    },
    "dashboard": {
      "label": "Dashboard",
      "command": "bun",
      "args": ["dev"],
      "cwd": "./www/dashboard"
    },
    "landing": {
      "label": "Landing",
      "command": "npm",
      "args": ["run", "dev"],
      "cwd": "./www/landing"
    }
  }
}
```

### Types for `trane.json`

If your editor supports it, enable auto-completion and type-checking:

```jsonc
// @ts-check
// @type {import('@xenoverseup/trane/config').default}
{
  "tasks": {...}
}
```

## TODO Roadmap

- [x] Go migration
- [ ] Fix posinstall script
- [ ] Export logs to file
- [ ] Timestamp each output
- [ ] Native buffer support
- [ ] Clear/reset output buffer
- [ ] Dual modes: interactive / headless
- [ ] Autocomplete alias names

---

[Muhammed Can Durmus](https://github.com/xenoverseup) · 2025

Licensed under the [MIT License](./LICENSE).
