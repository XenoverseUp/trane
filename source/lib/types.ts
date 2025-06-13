import {ChildProcess} from 'child_process';

export type ProcStatus = 'running' | 'success' | 'error';

export interface Command {
	label: string;
	command: string;
	args?: string[];
	cwd?: string;
}

export interface Proc extends Command {
	process: ChildProcess;
	output: string[];
	status: ProcStatus;
}
