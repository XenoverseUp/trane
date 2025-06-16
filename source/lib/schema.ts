import {z} from 'zod';

const shellSafeCommand = z
	.string()
	.min(1, 'Command cannot be empty')
	.regex(/^[^\s]+$/, 'Command must not contain spaces');

export const CommandSchema = z.object({
	label: z.string().max(12),
	alias: z
		.string()
		.regex(
			/^[a-zA-Z0-9-_]+$/,
			'Alias must be alphanumeric with dashes/underscores',
		)
		.optional(),
	command: shellSafeCommand,
	args: z.array(z.string()).optional(),
	cwd: z.string().optional(),
});

export const ConfigSchema = z.array(CommandSchema);
