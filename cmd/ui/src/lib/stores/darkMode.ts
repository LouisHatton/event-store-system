import { onMount } from 'svelte';
import { writable } from 'svelte/store';

function createDarkMode() {
	const DARK_MODE_LOCAL_STOREAGE_KEY = 'insight-wave-dark-mode';
	const { subscribe, set, update } = writable<boolean>(false);

	function updateDarkmode(enabled: boolean) {
		window.localStorage.setItem(DARK_MODE_LOCAL_STOREAGE_KEY, enabled.toString());
		set(enabled);
		return enabled;
	}

	let local = localStorage.getItem(DARK_MODE_LOCAL_STOREAGE_KEY);
	if (local === 'true') {
		updateDarkmode(true);
	} else {
		updateDarkmode(false);
	}

	return {
		subscribe,
		toggle: () => update((n) => updateDarkmode(!n))
	};
}

export const darkMode = createDarkMode();
