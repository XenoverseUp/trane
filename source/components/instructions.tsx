import {Box, Text} from 'ink';
import React from 'react';

type Props = {
	autoScroll: boolean;
};

function Instructions({autoScroll}: Props) {
	return (
		<Box justifyContent="space-between" paddingX={2} flexShrink={0}>
			<Text color="gray">←/→ or 1-9 to switch tabs, 'q' to quit</Text>
			<Text color="gray">
				Auto-scroll: {autoScroll ? 'ON' : 'OFF'} (toggle with 'a')
			</Text>
		</Box>
	);
}

export default React.memo(
	Instructions,
	(prev, next) => prev.autoScroll === next.autoScroll,
);
