import {Badge} from '@inkjs/ui';
import {Box, Spacer, Text, useApp, useInput} from 'ink';

type Props = {
	error: string;
};

export default function Error({error}: Props) {
	const {exit} = useApp();

	useInput((input, _) => input === 'q' && exit());

	return (
		<Box flexDirection="column" width="100%" height="100%">
			<Box
				borderStyle="round"
				borderLeft={false}
				borderTop={false}
				borderRight={false}
				paddingX={1}
				borderBottomColor="blueBright"
				flexDirection="row"
				justifyContent="space-between"
				flexShrink={0}
			>
				<Badge color="blueBright">⛬ Trane</Badge>
			</Box>
			<Box flexDirection="column" width="100%" paddingX={1} flexGrow={1}>
				<Text color="redBright">{error}</Text>
				<Spacer />
				<Text color="gray">Press key 'q' to exit...</Text>
			</Box>
		</Box>
	);
}
