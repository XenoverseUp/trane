import {Command, ProcStatus} from './types.js';

import fs from 'fs';
import path from 'path';
import {ConfigSchema} from './schema.js';

export const statusSymbol = (status: ProcStatus) => {
	switch (status) {
		case 'running':
			return '●';
		case 'success':
			return '✔';
		case 'error':
			return '✖';
	}
};

export const statusColor = (status: ProcStatus) => {
	switch (status) {
		case 'running':
			return 'yellowBright';
		case 'success':
			return 'greenBright';
		case 'error':
			return 'redBright';
	}
};

export function tryLoadCommands(filePath: string): {
	commands: Command[];
	error?: string;
} {
	const resolvedPath = path.resolve(process.cwd(), filePath);

	if (!fs.existsSync(resolvedPath)) {
		return {
			commands: [],
			error: `✖ File not found: ${resolvedPath}`,
		};
	}

	try {
		const raw = fs.readFileSync(resolvedPath, 'utf8');
		const parsed = JSON.parse(raw);

		const validated = ConfigSchema.safeParse(parsed);

		if (!validated.success) {
			const message = validated.error.errors
				.map(e => `- ${e.path.join('. ')}: ${e.message}.`)
				.join('\n');
			return {
				commands: [],
				error: `✖ Invalid config format:\n${message}`,
			};
		}

		const commands = validated.data.map((cmd, i) => ({
			label: cmd.label ?? cmd.alias ?? `Command ${i + 1}`,
			alias: cmd.alias,
			command: cmd.command,
			args: cmd.args ?? [],
			cwd: cmd.cwd ?? process.cwd(),
		}));

		return {commands};
	} catch (err) {
		return {
			commands: [],
			error: `✖ Failed to parse JSON: ${(err as Error).message}.`,
		};
	}
}
