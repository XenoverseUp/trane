#!/usr/bin/env node

// TODO: Update bin loading
import { fileURLToPath } from 'url'
import { dirname, join } from 'path'
import { spawn } from 'child_process'
import os from 'os'

const __dirname = dirname(fileURLToPath(import.meta.url))

let binary
switch (os.platform()) {
  case 'darwin':
    binary = join(__dirname, '../go/trane-darwin')
    break
  case 'linux':
    binary = join(__dirname, '../go/trane-linux')
    break
  case 'win32':
    binary = join(__dirname, '../go/trane-win.exe')
    break
  default:
    console.error(`Unsupported platform: ${os.platform()}`)
    process.exit(1)
}

const subprocess = spawn(binary, process.argv.slice(2), {
  stdio: 'inherit'
})

subprocess.on('exit', (code) => {
  process.exit(code)
})
