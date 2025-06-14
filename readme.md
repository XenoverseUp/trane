# Trane &middot; [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/xenoverseup/trane/blob/main/LICENSE) [![npm version](https://img.shields.io/npm/v/react.svg?style=flat)](https://www.npmjs.com/package/@xenoverseup/trane)

Trane is a modern take on parallel task running. The NPM package [Concurrently](https://www.npmjs.com/package/concurrently) is pretty popular in Node environment. However, it doesn't provide a user-friendly DX and UX. Trane aims to provide a similar functionality with a config based approach and a nicer TUI on top.

<img src="./docs/showcase.gif" alt="The usage of Trane." />

## Install

```bash
$ npm install --global trane
```

## CLI

```bash
$ trane --help

	Usage
    $ trane [options]

  Options
    --file, -f     Path to command JSON file (default: ./trane.json)
    --version, -v  Show CLI version
    --help, -h     Show this help

  Example
    $ trane --file=./my-commands.json
```

## Usage

## Technicality & Implementation

### Technologies:

## Todo:

- Runtime safety for `trane.json` files.
- Define config in the `trane.json`.
- Better handling of scroll area.
- Help screen for commands.
- Support native buffer format.
- Modes for interaction and commands.
- Export output logs.
- Clear output buffer.
- Timestamps.
- Provide both config based and cli based command running.

Muhammed Can Durmus | 2025
