<script lang="ts">
	import { navigating } from '$app/stores';
	import { classnames } from '$lib/util';
	import { tick } from 'svelte';

	$: if ($navigating) startLoad();
	else stopLoad();

	let loaderStyle = 'fixed top-0 h-2 bg-amber-500 dark:bg-amber-600';

	let loading = true;
	let loadingComplete = false;
	let loadingProgress = 0;

	let loadingInterval: NodeJS.Timer;

	async function startLoad() {
		if (loading) return;
		clearInterval(loadingInterval);
		await tick();
		loadingComplete = false;
		loading = true;
		loadingProgress = 0;
		let increment = 10;
		let multiplier = 1.2;
		await tick();
		loadingInterval = setInterval(() => {
			if (loading) {
				loadingProgress += increment;
				increment = increment / multiplier;
				multiplier -= 0.0005;
			}
			if (loadingProgress > 100) {
				clearInterval(loadingInterval);
			}
		}, 100);
	}

	async function stopLoad() {
		clearInterval(loadingInterval);
		loadingComplete = true;
		loadingProgress = 0;

		loading = true;
		loadingInterval = setInterval(() => {
			if (loading) loadingProgress += 25;
			if (loadingProgress > 100) {
				clearInterval(loadingInterval);
				setTimeout(() => {
					loadingComplete = false;
					loading = false;
				}, 100);
			}
		}, 50);
	}
</script>

{#if loadingComplete}
	<div
		class={classnames(loaderStyle, `transition-all ease-linear duration-75`)}
		style={`width: ${loadingProgress.toString()}%;`}
	/>
{:else}
	<div
		class={classnames(
			loaderStyle,
			`${loading ? 'transition-all ease-linear duration-700' : 'hidden'}`
		)}
		style={`width: ${loadingProgress.toString()}%;`}
	/>
{/if}
