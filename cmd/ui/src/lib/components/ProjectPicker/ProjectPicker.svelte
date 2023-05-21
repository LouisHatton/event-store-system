<script lang="ts">
	import { Dropdown } from 'flowbite-svelte';
	import ProjectIcon from './ProjectIcon.svelte';
	import ProjectService from '$lib/services/ProjectService';
	import { activeProject, projectStore } from '$lib/stores/projectStore';

	export let placement = 'top';
	let open = false;

	$: if (!$projectStore) ProjectService.listProjects().then((p) => projectStore.set(p));
	$: if ($projectStore && !$activeProject) activeProject.set($projectStore[0]);

	function updateActiveProject(id: string) {
		open = false;
		if (!$projectStore) return;
		activeProject.set($projectStore.find((p) => p.id === id));
	}
</script>

<button color="light" class="w-full">
	{#if $activeProject}
		<ProjectIcon colour={$activeProject.config.colour} name={$activeProject.name} chevron />
	{:else}
		<div class="flex px-4 py-3 gap-x-4 items-center">
			<div class="w-10 flex-shrink-0 h-10 bg-zinc-300 dark:bg-zinc-600 animate-pulse rounded-md" />
			<div class="w-full h-4 bg-zinc-300 dark:bg-zinc-600 animate-pulse rounded-md" />
		</div>
	{/if}
</button>
{#if $projectStore}
	<Dropdown
		{placement}
		bind:open
		class="w-64 overflow-y-auto py-1 h-48 gap-y-2 flex flex-col rounded-md border-t border-l border-r border-gray-200 dark:border-gray-600"
	>
		{#each $projectStore as project}
			<ProjectIcon
				name={project.name}
				colour={project.config.colour}
				on:click={() => updateActiveProject(project.id)}
			/>
		{/each}
		<a
			slot="footer"
			href="/"
			class="flex items-center px-3 py-4 -mb-1 text-sm font-medium text-amber-600 hover:bg-gray-100 dark:bg-gray-700 dark:hover:bg-gray-600 dark:text-amber-500"
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 24 24"
				stroke-width="1.5"
				stroke="currentColor"
				class="h-5 w-5 mr-1"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M19 7.5v3m0 0v3m0-3h3m-3 0h-3m-2.25-4.125a3.375 3.375 0 11-6.75 0 3.375 3.375 0 016.75 0zM4 19.235v-.11a6.375 6.375 0 0112.75 0v.109A12.318 12.318 0 0110.374 21c-2.331 0-4.512-.645-6.374-1.766z"
				/></svg
			>Create New Project
		</a>
	</Dropdown>
{/if}
