#!/usr/bin/env node

import { platform, arch } from 'node:os'
import { spawnSync } from 'node:child_process'
import { join, dirname } from 'node:path'
import { fileURLToPath } from 'node:url'
import { existsSync } from 'node:fs'
import { globSync } from 'glob'

const __dirname = dirname(fileURLToPath(import.meta.url))

const osMap = {
  win32: 'windows',
  darwin: 'darwin',
  linux: 'linux',
}

const archMap = {
  x64: 'amd64',
  arm64: 'arm64',
}

const os = osMap[platform()]
const cpu = archMap[arch()]

if (!os || !cpu) {
  console.error(`Unsupported platform or architecture: ${platform()} ${arch()}`)
  process.exit(1)
}

const binaryName = os === 'windows' ? 'trane.exe' : 'trane'

const globPattern = join(__dirname, `bin/trane_${os}_${cpu}_v*/`)

const matches = globSync(globPattern)

if (!matches.length) {
  console.error(`No binary found for ${os}-${cpu} in npm/bin/`)
  process.exit(1)
}

const binaryDir = matches[0]
const binaryPath = join(binaryDir, binaryName)

if (!existsSync(binaryPath)) {
  console.error(`Binary not found at ${binaryPath}`)
  process.exit(1)
}

const result = spawnSync(binaryPath, process.argv.slice(2), {
  stdio: 'inherit',
})

process.exit(result.status ?? 1)
