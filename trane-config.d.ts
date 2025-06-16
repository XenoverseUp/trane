export interface Command {
	label: string;
	command: string;
	args?: string[];
	alias?: string;
	cwd?: string;
}

declare const config: Command[];
export default config;
