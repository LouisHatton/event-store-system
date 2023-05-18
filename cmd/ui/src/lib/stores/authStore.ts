import type { User } from 'firebase/auth';
import { writable } from 'svelte/store';

export const authStore = writable<{
	isLoggedIn: boolean;
	user?: User | null;
	initialised: boolean;
}>({
	isLoggedIn: false,
	initialised: false
});
