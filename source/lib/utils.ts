import {ProcStatus} from './types.js';

export const statusSymbol = (status: ProcStatus) => {
	switch (status) {
		case 'running':
			return '●';
		case 'success':
			return '✔';
		case 'error':
			return '✖';
	}
};

export const statusColor = (status: ProcStatus) => {
	switch (status) {
		case 'running':
			return 'yellowBright';
		case 'success':
			return 'greenBright';
		case 'error':
			return 'redBright';
	}
};
