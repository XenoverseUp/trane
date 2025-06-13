import React, {useEffect, useRef, useState} from 'react';
import {Box, Text, useInput, measureElement} from 'ink';

type Props = {
	autoScroll: boolean;
	output: string[];
};

const SCROLL_SPEED = 2;

export default function TabContent({autoScroll, output}: Props) {
	const [offset, setOffset] = useState(0);
	const [maxLines, setMaxLines] = useState(-1);
	const containerRef = useRef(null);

	useEffect(() => {
		if (containerRef.current) {
			const {height} = measureElement(containerRef.current);
			setMaxLines(Math.max(1, height));
		}
	}, [containerRef.current]);

	useEffect(() => {
		if (autoScroll) setOffset(Math.max(0, output.length - maxLines));
	}, [output.length, autoScroll, maxLines]);

	useInput((input, key) => {
		if (autoScroll) return;

		if (key.upArrow || input === 'k')
			setOffset(prev => Math.max(0, prev - SCROLL_SPEED));

		if (key.downArrow || input === 'j')
			setOffset(prev =>
				Math.min(Math.max(0, output.length - maxLines), prev + SCROLL_SPEED),
			);
	});

	if (maxLines === -1)
		return <Box ref={containerRef} flexGrow={1} padding={1} />;

	const visible = output.slice(offset, offset + maxLines);

	return (
		<Box ref={containerRef} flexGrow={1} padding={1} flexDirection="column">
			{maxLines !== -1 && visible.map((line, i) => <Text key={i}>{line}</Text>)}
		</Box>
	);
}
