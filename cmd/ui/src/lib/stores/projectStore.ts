import type { Project } from '$lib/services/ProjectService';
import { writable } from 'svelte/store';

export const projectStore = writable<Project[] | undefined>(undefined);

export const activeProject = writable<Project | undefined>(undefined);
