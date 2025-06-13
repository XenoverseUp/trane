import {useEffect, useState} from 'react';
import {spawn} from 'child_process';
import {Command, Proc, ProcStatus} from '../lib/types.js';

export default function useProcesses(commands: Command[]): Proc[] {
	const [procs, setProcs] = useState<(Proc & {buffer?: string})[]>([]);

	useEffect(() => {
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
				buffer: '', // temporary line buffer
			};
		});

		setProcs(running);

		running.forEach(proc => {
			const {process: child, label} = proc;

			const handleChunk = (data: Buffer, isError = false) => {
				setProcs(prev =>
					prev.map(p => {
						if (p.label !== label) return p;

						const raw = (p.buffer ?? '') + data.toString();
						const parts = raw.split('\n');

						const completedLines = parts
							.slice(0, -1)
							.map(line => (isError ? `\u001b[31m${line}` : line));

						return {
							...p,
							output: [...p.output, ...completedLines],
							buffer: parts?.[-1] ?? '',
						};
					}),
				);
			};

			child.stdout?.on('data', data => handleChunk(data, false));
			child.stderr?.on('data', data => handleChunk(data, true));

			child.on('exit', code => {
				setProcs(prev =>
					prev.map(p => {
						if (p.label !== label) return p;

						const flushedLine = p.buffer ? [p.buffer] : [];
						const exitMsg = `[${label}] exited with code ${code}`;

						return {
							...p,
							status: code === 0 ? 'success' : 'error',
							output: [...p.output, ...flushedLine, exitMsg],
							buffer: '',
						};
					}),
				);
			});
		});

		return () => {
			running.forEach(p => p.process.kill());
		};
	}, [commands]);

	// Strip internal buffer before returning
	return procs.map(({buffer, ...rest}) => rest);
}
