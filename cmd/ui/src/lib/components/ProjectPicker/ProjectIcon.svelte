<script lang="ts">
	import Dots from '$lib/icons/Dots.svelte';
	import Project from '$lib/icons/Project.svelte';
	import { classnames } from '$lib/util';
	import { Avatar, Chevron, ChevronUp, DropdownItem, MenuButton } from 'flowbite-svelte';

	export let chevron = false;
	export let colour = 'bg-green-400';

	$: syleClass = chevron ? 'py-3 px-4' : 'py-2 px-4';
</script>

<button
	class={classnames(
		'w-full font-medium text-sm hover:bg-gray-100 dark:hover:bg-gray-600 text-left flex items-center text-zinc-700 dark:text-zinc-100 rounded-md',
		syleClass
	)}
>
	<div
		class={classnames(
			'w-10 h-10 p-2 fill-current text-neutral-100 rounded-md flex-shrink-0',
			colour
		)}
	>
		<Project />
	</div>
	<span class="ml-4 font-semibold line-clamp-2"> My New Project</span>
	<span class="ml-auto mr-2 fill-current w-5">
		{#if chevron}
			<Dots horizontal />
		{/if}
	</span>
</button>
