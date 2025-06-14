import {Badge} from '@inkjs/ui';
import {Box, Text} from 'ink';
import {Proc} from '../lib/types.js';
import {statusColor, statusSymbol} from '../lib/utils.js';

type Props = {
	activeIndex: number;
	procs: Proc[];
};

export default function Tabbar({activeIndex, procs}: Props) {
	return (
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
			<Box
				borderRight
				borderLeft={false}
				borderTop={false}
				borderBottom={false}
				borderStyle="classic"
				alignItems="center"
				paddingRight={1}
				marginRight={1}
			>
				<Badge color="blueBright">⛬ Trane</Badge>
			</Box>
			<Box gap={2}>
				{procs.map((proc, i) => (
					<Text
						key={proc.label}
						color={statusColor(proc.status)}
						inverse={i === activeIndex}
					>
						{' '}
						({i + 1}) {proc.label} {statusSymbol(proc.status)}{' '}
					</Text>
				))}
			</Box>
		</Box>
	);
}
