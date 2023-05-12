import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	const delay = (ms) => new Promise((res) => setTimeout(res, ms));
	let ok = true;
	await delay(2000);

	return {
		ok
	};
};
