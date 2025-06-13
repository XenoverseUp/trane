import React, {useState} from 'react';
import {Box, useApp, useInput} from 'ink';

import Tabbar from './components/tab-bar.js';
import {Proc} from './lib/types.js';
import useProcesses from './hooks/useProcess.js';
import TabContent from './components/tab-content.js';
import Instructions from './components/instructions.js';

const commands = [
	{label: 'Server', command: 'ls', args: ['-la'], cwd: '.'},
	{label: 'Dashboard', command: 'sleep', args: ['15'], cwd: '.'},
	{label: 'Landing', command: 'brew', args: ['install'], cwd: '.'},
];

export default function App() {
	const {exit} = useApp();
	const [tabIndex, setTabIndex] = useState(0);
	const [autoScroll, setAutoScroll] = useState(true);

	const procs: Proc[] = useProcesses(commands);

	useInput((input, key) => {
		if (input === 'q') exit();

		if (key.leftArrow || input === 'h')
			setTabIndex(i => (i - 1 + procs.length) % procs.length);
		if (key.rightArrow || input === 'l')
			setTabIndex(i => (i + 1) % procs.length);

		if (input === 'a') setAutoScroll(prev => !prev);
		if (key.upArrow || input === 'k') setAutoScroll(false);

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

			<TabContent
				key={activeProc?.label ?? 'none'}
				autoScroll={autoScroll}
				output={outputLines}
			/>

			{/* Instructions */}
			<Instructions {...{autoScroll}} />
		</Box>
	);
}
