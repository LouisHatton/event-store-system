<script lang="ts">
	import Dots from '$lib/icons/Dots.svelte';
	import Project from '$lib/icons/Project.svelte';
	import { classnames } from '$lib/util';

	export let chevron = false;
	export let colour: string;
	export let name: string;

	$: syleClass = chevron ? 'py-3 px-4' : 'py-2 px-4';
	$: classColour = getClassColour(colour);

	function getClassColour(c: string) {
		switch (c) {
			case 'green':
				return 'bg-green-400';
			case 'red':
			default:
				return 'bg-red-400';
		}
	}
</script>

<button
	on:click
	class={classnames(
		'w-full font-medium text-sm hover:bg-gray-100 dark:hover:bg-gray-600 text-left flex items-center text-zinc-700 dark:text-zinc-100 rounded-md',
		syleClass
	)}
>
	<div
		class={classnames(
			'w-10 h-10 p-2 fill-current text-neutral-100 rounded-md flex-shrink-0',
			classColour
		)}
	>
		<Project />
	</div>
	<span class="ml-4 font-semibold line-clamp-2">{name}</span>
	<span class="ml-auto mr-2 fill-current w-5">
		{#if chevron}
			<Dots horizontal />
		{/if}
	</span>
</button>
