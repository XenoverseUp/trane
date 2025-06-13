import React, {useState} from 'react';
import {Box, Text, useApp, useInput} from 'ink';

import Tabbar from './components/tab-bar.js';
import {Proc} from './lib/types.js';
import useProcesses from './hooks/useProcess.js';

const commands = [
	{label: 'Server', command: 'ls', args: ['-la'], cwd: '.'},
	{label: 'Dashboard', command: 'sleep', args: ['15'], cwd: '.'},
	{label: 'Landing', command: 'lss', args: ['-ld'], cwd: '.'},
];

export default function App() {
	const {exit} = useApp();
	const [tabIndex, setTabIndex] = useState(0);
	const [autoScroll, setAutoScroll] = useState(true);

	const procs: Proc[] = useProcesses(commands);

	useInput((input, key) => {
		if (input === 'q') exit();
		if (key.leftArrow) setTabIndex(i => (i - 1 + procs.length) % procs.length);
		if (key.rightArrow) setTabIndex(i => (i + 1) % procs.length);
		if (input === 'a') setAutoScroll(prev => !prev);
		if (!isNaN(Number(input))) {
			const idx = parseInt(input, 10) - 1;
			if (idx >= 0 && idx < procs.length) setTabIndex(idx);
		}
	});

	const activeProc: Proc | undefined = procs[tabIndex];
	const outputLines = activeProc?.output ?? [];

	return (
		<Box flexDirection="column" width="100%" height="100%">
			{/* Tab bar */}
			<Tabbar procs={procs} activeIndex={tabIndex} />

			{/* Output */}

			<Box flexGrow={1} padding={1} flexDirection="column" overflow="hidden">
				{(autoScroll ? outputLines.slice(-30) : outputLines.slice(0, 30)).map(
					(line, i) => (
						<Text key={i}>{line}</Text>
					),
				)}
			</Box>

			{/* Instructions */}
			<Box justifyContent="space-between">
				<Text color="gray">←/→ or 1-9 to switch tabs, 'q' to quit</Text>
				<Text color="gray">
					Auto-scroll: {autoScroll ? 'ON' : 'OFF'} (toggle with 'a')
				</Text>
			</Box>
		</Box>
	);
}
