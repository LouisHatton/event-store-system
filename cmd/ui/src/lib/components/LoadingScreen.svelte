<script lang="ts">
	import { page } from '$app/stores';
	import type { Unsubscriber } from 'svelte/store';
	import { onDestroy, onMount } from 'svelte';
	import { Spinner } from 'flowbite-svelte';

	let showLoading = false;
	let unsubscribe: Unsubscriber;
	onMount(() => {
		unsubscribe = page.subscribe(() => {
			showLoading = false;
			setTimeout(() => {
				showLoading = true;
			}, 300);
		});
	});

	onDestroy(() => {
		unsubscribe();
	});
</script>

{#if showLoading}
	<div class="flex justify-center items-center min-h-screen dark:bg-gray-700">
		<Spinner color="yellow" size="12" />
	</div>
{/if}
