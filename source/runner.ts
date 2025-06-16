import execa from 'execa';
import readline from 'readline';
import {Command} from './lib/types.js';

function promptYesNo(question: string): Promise<boolean> {
	return new Promise(resolve => {
		const rl = readline.createInterface({
			input: process.stdin,
			output: process.stdout,
		});
		rl.question(question + ' [y/N] ', answer => {
			rl.close();
			resolve(/^y(es)?$/i.test(answer));
		});
	});
}

export async function runCommand(cmd: Command) {
	console.log(`💡 Running: ${cmd.command} ${cmd.args?.join(' ')}`);
	console.log(`📁 cwd: ${cmd.cwd}`);

	const proceed = await promptYesNo('Do you want to execute this command?');
	if (!proceed) {
		console.log('❌ Aborted.');
		process.exit(0);
	}

	const subprocess = execa(cmd.command, cmd.args, {
		cwd: cmd.cwd,
		stdio: 'inherit',
	});
	await subprocess;
}
