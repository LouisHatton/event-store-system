<script lang="ts">
	import { classnames } from '$lib/util';
	import { Spinner } from 'flowbite-svelte';

	export let loading = false;
	export let size: 'sm' | 'md' | 'lg' | 'xl' | '2xl' = 'md';

	$: height = calculateHeight(size);

	function calculateHeight(size: 'sm' | 'md' | 'lg' | 'xl' | '2xl') {
		switch (size) {
			case 'sm':
				return 'h-6';
			case 'md':
				return 'h-28';
			case 'lg':
				return 'h-36';
			case 'xl':
				return 'h-52';
			case '2xl':
				return 'h-64';
		}
	}
</script>

{#if loading}
	<div class={classnames('flex flex-row justify-center items-center', height)}>
		<Spinner color="yellow" />
	</div>
{:else}
	<slot />
{/if}
