import React, {useEffect, useState} from 'react';
import {Box, Text} from 'ink';

type Props = {
	autoScroll: boolean;
	output: string[];
};

export default function TabContent({autoScroll, output}: Props) {
	const [visible, setVisible] = useState<string[]>([]);

	useEffect(() => {
		setVisible(autoScroll ? output.slice(-30) : output.slice(0, 30));
	}, [output, autoScroll]);

	return (
		<Box flexGrow={1} padding={1} flexDirection="column" overflow="hidden">
			{visible.map((line, i) => (
				<Text key={i}>{line}</Text>
			))}
		</Box>
	);
}
