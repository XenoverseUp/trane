import {useEffect, useState, useRef, useCallback} from 'react';
import {spawn, ChildProcessWithoutNullStreams} from 'child_process';
import {Command, Proc} from '../lib/types.js';

type ProcWithBuffer = Proc & {
	buffer?: string;
	process: ChildProcessWithoutNullStreams;
};

export default function useProcesses(commands: Command[]) {
	const [procs, setProcs] = useState<ProcWithBuffer[]>([]);
	const commandsRef = useRef(commands);

	useEffect(() => {
		commandsRef.current = commands;
	}, [commands]);

	const spawnProcess = useCallback((cmd: Command): ProcWithBuffer => {
		const child = spawn(cmd.command, cmd.args, {
			cwd: cmd.cwd,
			shell: true,
		});

		const proc: ProcWithBuffer = {
			...cmd,
			process: child,
			output: [] as string[],
			status: 'running',
			buffer: '',
		};

		const handleChunk = (data: Buffer, isError = false) => {
			setProcs(prev =>
				prev.map(p => {
					if (p.label !== cmd.label) return p;

					const raw = (p.buffer ?? '') + data.toString();
					const parts = raw.split('\n');

					const completedLines = parts
						.slice(0, -1)
						.map(line => (isError ? `\u001b[31m${line}` : line));

					const lastLine = parts[parts.length - 1] ?? '';

					return {
						...p,
						output: [...p.output, ...completedLines],
						buffer: lastLine,
					};
				}),
			);
		};

		child.stdout.on('data', data => handleChunk(data, false));
		child.stderr.on('data', data => handleChunk(data, true));

		child.on('error', err => {
			setProcs(prev =>
				prev.map(p => {
					if (p.label !== cmd.label) return p;

					const errorMessage = `✖ Failed to start process: ${err.message}`;
					return {
						...p,
						status: 'error',
						output: [...p.output, errorMessage],
					};
				}),
			);
		});

		child.on('exit', code => {
			setProcs(prev =>
				prev.map(p => {
					if (p.label !== cmd.label) return p;

					const flushedLine = p.buffer ? [p.buffer] : [];
					const exitMsg = `[${cmd.label}] exited with code ${code}`;

					return {
						...p,
						status: code === 0 ? 'success' : 'error',
						output: [...p.output, ...flushedLine, exitMsg],
						buffer: '',
					};
				}),
			);
		});

		return proc;
	}, []);

	// On mount or commands change: spawn all processes
	useEffect(() => {
		const running = commandsRef.current.map(cmd => spawnProcess(cmd));
		setProcs(running);

		// Cleanup all processes on unmount
		return () => {
			running.forEach(p => {
				if (!p.process.killed) p.process.kill();
			});
		};
	}, [spawnProcess]);

	// Kill a process by label
	const killProc = useCallback((label: string) => {
		setProcs(prev => {
			return prev.map(p => {
				if (p.label === label && !p.process.killed) {
					p.process.kill();
					return {
						...p,
						status: 'error',
						output: [...p.output, `[${label}] killed`],
					};
				}
				return p;
			});
		});
	}, []);

	// Restart a process by label
	const restartProc = useCallback(
		(label: string) => {
			setProcs(prev => {
				// Kill old process if running
				const oldProc = prev.find(p => p.label === label);
				if (oldProc && !oldProc.process.killed) {
					oldProc.process.kill();
				}

				// Find command config for that label
				const cmd = commandsRef.current.find(c => c.label === label);
				if (!cmd) return prev;

				// Spawn new process for the label
				const newProc = spawnProcess(cmd);

				// Replace old proc with new proc
				return prev.map(p => (p.label === label ? newProc : p));
			});
		},
		[spawnProcess],
	);

	// Return processes *without* internal buffer, plus control functions
	const cleanProcs: Proc[] = procs.map(({buffer, ...rest}) => rest);

	return {procs: cleanProcs, killProc, restartProc};
}
