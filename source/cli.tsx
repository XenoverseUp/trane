#!/usr/bin/env node

import {render, Box} from 'ink';
import meow from 'meow';
import App from './app.js';
import {tryLoadCommands} from './lib/utils.js';
import Error from './error.js';

const cli = meow(
	`
  Usage
    $ trane [options] [alias]

  Options
    --file, -f     Path to command JSON file (default: ./trane.json)
    --version, -v  Show CLI version
    --help, -h     Show this help

  Example
    $ trane --file=./my-commands.json
    $ trane build
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
			list: {
				type: 'boolean',
				alias: 'l',
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

(function main() {
	if (cli.flags.help) {
		cli.showHelp();
		return;
	}

	if (cli.flags.version) {
		cli.showVersion();
		return;
	}

	const {commands, error} = tryLoadCommands(cli.flags.file);

	/* Error TUI */

	if (error) {
		process.stdout.write('\x1b[?1049h');
		process.on('exit', () => process.stdout.write('\x1b[?1049l'));

		render(
			<Box width={process.stdout.columns} height={process.stdout.rows}>
				<Error error={error} />
			</Box>,
		);

		return;
	}

	/* Alias Listing */

	if (cli.flags.list) {
		const listing = commands.filter(cmd => cmd.alias);

		if (listing.length === 0) {
			console.log('\nNo defined aliases command.');
			console.log('');
			return;
		}

		console.log('\nAvailable Commands:\n');
		listing.forEach(cmd => {
			const command = [cmd.command, ...(cmd?.args ?? [])].join(' ');
			return console.log(`- ${cmd.alias?.padEnd(12)} → ${command}`);
		});

		console.log('');
		return;
	}

	/* Alias Mode */

	if (cli.input.length > 0) {
		const alias = cli.input[0];
		const command = commands.find(c => c.alias === alias);
		if (!command) {
			console.error(`✖ No command found with alias "${alias}"`);
			process.exit(1);
		}
		import('./runner.js')
			.then(({runCommand}) => runCommand(command))
			.catch(err => {
				console.error('✖ Failed to run command:', err.message);
				process.exit(1);
			});

		return;
	}

	/* TUI Mode */
	process.stdout.write('\x1b[?1049h');
	process.on('exit', () => process.stdout.write('\x1b[?1049l'));

	render(
		<Box width={process.stdout.columns} height={process.stdout.rows}>
			<App commands={commands} />
		</Box>,
	);
})();
