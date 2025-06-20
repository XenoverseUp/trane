import {useMemo, useState} from 'react';
import {Box, useApp, useInput} from 'ink';

import Tabbar from './components/tab-bar.js';
import {Command, Proc} from './lib/types.js';
import useProcesses from './hooks/useProcess.js';
import TabContent from './components/tab-content.js';
import Instructions from './components/instructions.js';

type Props = {
	commands: Command[];
};

export default function App({commands}: Props) {
	const {exit} = useApp();
	const [tabIndex, setTabIndex] = useState(0);
	const [autoScroll, setAutoScroll] = useState(true);

	const {procs, killProc, restartProc} = useProcesses(commands);

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

		if (input === 'x') {
			const active = procs[tabIndex];
			if (active) killProc(active.label);
		}
		if (input === 'r') {
			const active = procs[tabIndex];
			if (active) restartProc(active.label);
		}
	});

	const activeProc: Proc | undefined = useMemo(
		() => procs[tabIndex],
		[procs, tabIndex],
	);

	return (
		<Box flexDirection="column" width="100%" height="100%">
			{/* Tab bar */}
			<Tabbar procs={procs} activeIndex={tabIndex} />

			{/* Output */}

			<TabContent
				key={activeProc?.label ?? 'none'}
				autoScroll={autoScroll}
				output={activeProc?.output ?? []}
			/>

			{/* Instructions */}
			<Instructions {...{autoScroll}} />
		</Box>
	);
}
