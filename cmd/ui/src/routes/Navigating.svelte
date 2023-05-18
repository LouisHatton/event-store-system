<script lang="ts">
	import { navigating } from '$app/stores';
	import { classnames } from '$lib/util';
	import { tick } from 'svelte';

	export let externalLoading = false;

	$: if ($navigating || externalLoading) startLoad();
	else stopLoad();

	let loaderStyle = 'fixed top-0 left-0 right-0 h-1 bg-amber-500 dark:bg-amber-600 z-50';

	let loading = true;
	let loadingComplete = false;
	let loadingProgress = 0;
	let opacity = 1;

	$: if (loadingProgress > 100) {
		loadingProgress = 100;
	}

	let loadingInterval: NodeJS.Timer;

	/**
	 * This loading bar lasts around 24 seconds before reaching the end.
	 */
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
				if (increment > 0.2) increment = increment / multiplier;
				if (multiplier > 0.1) multiplier -= 0.0005;
			}
			if (loadingProgress >= 100) {
				clearInterval(loadingInterval);
			}
		}, 100);
	}

	async function stopLoad() {
		opacity = 1;
		clearInterval(loadingInterval);
		loadingComplete = true;
		loadingProgress = 0;

		loading = true;
		loadingInterval = setInterval(() => {
			if (loading) loadingProgress += 25;
			if (loadingProgress >= 100) {
				clearInterval(loadingInterval);
				opacity = 0.5;
				setTimeout(() => {
					opacity = 0;
					loadingComplete = false;
					loading = false;
				}, 150);
			}
		}, 50);
	}
</script>

{#if loadingComplete}
	<div
		class={classnames(loaderStyle, `transition-all ease-linear duration-75`)}
		style={`width: ${loadingProgress.toString()}%;  opacity: ${opacity.toString()};`}
	/>
{:else}
	<div
		class={classnames(
			loaderStyle,
			`${loading ? 'transition-all ease-linear duration-700' : 'hidden'}`,
			`${loading && loadingProgress > 100 ? 'animate-pulse' : ''}`
		)}
		style={`width: ${loadingProgress.toString()}%;`}
	/>
{/if}
