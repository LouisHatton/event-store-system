import type { Connection } from '$lib/services/ConnectionService';
import { writable } from 'svelte/store';

export const connection = writable<Connection | undefined>(undefined);
