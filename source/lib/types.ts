import {ChildProcess} from 'child_process';
import {z} from 'zod';
import {CommandSchema} from './schema.js';

export type ProcStatus = 'running' | 'success' | 'error';

export type Command = z.infer<typeof CommandSchema>;

export interface Proc extends Command {
	process: ChildProcess;
	output: string[];
	status: ProcStatus;
}
