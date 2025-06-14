#!/usr/bin/env node

import {render, Box, Text} from 'ink';
import fs from 'fs';
import path from 'path';
import meow from 'meow';
import App from './app.js';
import {Command} from './lib/types.js';

const cli = meow(
	`
  Usage
    $ trane [options]

  Options
    --file, -f     Path to command JSON file (default: ./trane.json)
    --version, -v  Show CLI version
    --help, -h     Show this help

  Example
    $ trane --file=./my-commands.json
`,
	{
		importMeta: import.meta,
		flags: {
			file: {
				type: 'string',
				alias: 'f',
				default: './trane.json',
			},
			help: {
				type: 'boolean',
				alias: 'h',
			},
			version: {
				type: 'boolean',
				alias: 'v',
			},
		},
		autoHelp: false,
		autoVersion: false,
		version: '0.0.1',
	},
);

if (cli.flags.help) {
	cli.showHelp();
	process.exit(0);
}

if (cli.flags.version) {
	cli.showVersion();
	process.exit(0);
}

function tryLoadCommands(filePath: string): {
	commands: Command[];
	error?: string;
} {
	const resolvedPath = path.resolve(process.cwd(), filePath);

	if (!fs.existsSync(resolvedPath)) {
		return {
			commands: [],
			error: `❌ File not found: ${resolvedPath}`,
		};
	}

	try {
		const raw = fs.readFileSync(resolvedPath, 'utf8');
		const parsed = JSON.parse(raw);

		if (!Array.isArray(parsed)) {
			return {
				commands: [],
				error: `❌ JSON must be an array of command objects.`,
			};
		}

		const commands = parsed.map((cmd, i) => ({
			label: cmd.label ?? `Command ${i + 1}`,
			command: String(cmd.command),
			args: Array.isArray(cmd.args) ? cmd.args.map(String) : [],
			cwd: cmd.cwd ?? process.cwd(),
		}));

		return {commands};
	} catch (err) {
		return {
			commands: [],
			error: `❌ Failed to parse JSON: ${(err as Error).message}`,
		};
	}
}

const {commands, error} = tryLoadCommands(cli.flags.file);

process.stdout.write('\x1b[?1049h');

process.on('exit', () => {
	process.stdout.write('\x1b[?1049l');
});

render(
	<Box width={process.stdout.columns} height={process.stdout.rows}>
		{error ? (
			<Box flexDirection="column" padding={1}>
				<Text color="redBright">{error}</Text>
				<Text color="gray">Press any key to exit...</Text>
			</Box>
		) : (
			<App commands={commands} />
		)}
	</Box>,
);

if (error) {
	process.stdin.setRawMode?.(true);
	process.stdin.resume();
	process.stdin.once('data', () => process.exit(1));
}
