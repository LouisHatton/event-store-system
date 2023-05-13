<script lang="ts">
	import { classnames } from '$lib/util';

	type BoxColour = 'green' | 'yellow' | 'orange' | 'blue' | 'red' | 'neutral';
	export let colour: BoxColour = 'green';
	export let stat = '24';
	export let source = 'Event Sources';
	export let message = 'Up 5 from last week';
	export let href = '';

	$: boxColour = getBoxColour(colour);

	function getBoxColour(c: BoxColour) {
		switch (c) {
			case 'green':
				return 'bg-green-100 hover:bg-green-200 dark:bg-green-600/50 dark:hover:bg-green-700/50';
			case 'yellow':
				return 'bg-yellow-100 hover:bg-yellow-200 dark:bg-yellow-600/50 dark:hover:bg-yellow-700/50';
			case 'orange':
				return 'bg-orange-100 hover:bg-orange-200 dark:bg-orange-500/50 dark:hover:bg-orange-600/50';
			case 'blue':
				return 'bg-blue-100 hover:bg-blue-200 dark:bg-blue-600/50 dark:hover:bg-blue-700/50';
			case 'red':
				return 'bg-red-100 hover:bg-red-200 dark:bg-red-600/50 dark:hover:bg-red-700/50';
			case 'neutral':
				return 'bg-neutral-100 hover:bg-neutral-200 dark:bg-neutral-500/50 dark:hover:bg-neutral-600/50';
		}
	}
</script>

<div class="relative">
	<div class="absolute right-3 top-3 rotate-45">
		<svg
			clip-rule="evenodd"
			fill-rule="evenodd"
			stroke-linejoin="round"
			stroke-miterlimit="2"
			viewBox="0 0 24 24"
			class=" h-4 fill-current"
			xmlns="http://www.w3.org/2000/svg"
			><path
				d="m18.787 9.473s-4.505-4.502-6.259-6.255c-.147-.146-.339-.22-.53-.22-.192 0-.384.074-.531.22-1.753 1.753-6.256 6.252-6.256 6.252-.147.147-.219.339-.217.532.001.19.075.38.221.525.292.293.766.295 1.056.004l4.977-4.976v14.692c0 .414.336.75.75.75.413 0 .75-.336.75-.75v-14.692l4.978 4.978c.289.29.762.287 1.055-.006.145-.145.219-.335.221-.525.002-.192-.07-.384-.215-.529z"
				fill-rule="nonzero"
			/></svg
		>
	</div>
	<a
		{href}
		class={classnames(
			'flex flex-col justify-center items-center px-4 xl:px-16 py-6 transition-all ease-out duration-500 rounded-lg',
			boxColour
		)}
	>
		<p class="text-6xl pb-3">{stat}</p>
		<p class="text-base pb-1">{source}</p>
		<p class="text-sm">{message}</p>
	</a>
</div>
