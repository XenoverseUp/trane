import {useEffect, useState} from 'react';
import {spawn} from 'child_process';
import {Command, Proc, ProcStatus} from '../lib/types.js';

export default function useProcesses(commands: Command[]): Proc[] {
	const [procs, setProcs] = useState<Proc[]>([]);

	useEffect(() => {
		// Spawn all processes
		const running = commands.map(cmd => {
			const child = spawn(cmd.command, cmd.args, {
				cwd: cmd.cwd,
				shell: true,
			});

			return {
				...cmd,
				process: child,
				output: [] as string[],
				status: 'running' as ProcStatus,
			};
		});

		setProcs(running);

		// Attach listeners to each child process
		running.forEach(p => {
			const {process: child, label} = p;

			const appendLine = (line: string) => {
				setProcs(prev =>
					prev.map(pr =>
						pr.label === label ? {...pr, output: [...pr.output, line]} : pr,
					),
				);
			};

			child.stdout?.on('data', data => {
				appendLine(data.toString());
			});

			child.stderr?.on('data', data => {
				appendLine(`\u001b[31m${data.toString()}\u001b[0m`);
			});

			child.on('exit', code => {
				appendLine(`\n[${label}] exited with code ${code}`);
				setProcs(prev =>
					prev.map(pr =>
						pr.label === label
							? {...pr, status: code === 0 ? 'success' : 'error'}
							: pr,
					),
				);
			});
		});

		// Cleanup on unmount or commands change
		return () => {
			running.forEach(p => p.process.kill());
		};
	}, [commands]);

	return procs;
}
