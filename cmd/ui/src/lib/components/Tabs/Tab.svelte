<script lang="ts">
	import { page } from '$app/stores';

	export let title: string;
	export let href = '#';
	export let disabled = false;

	let barWidth = 0;

	$: active = $page.url.pathname === href;

	$: active ? startAnimation() : (barWidth = 0);

	function startAnimation() {
		setTimeout(() => {
			barWidth = 0;
			barWidth += 80;
		}, 100);
	}
</script>

<a href={disabled ? '#' : href} class="group transition-all ease-in-out">
	<span class="py-2 px-4 hover:bg-neutral-100 dark:hover:bg-slate-600 rounded-md">
		{#if disabled}
			<span class="opacity-50">{title}</span>
		{:else}
			<a {href}>{title} </a>
		{/if}
	</span>

	<div
		class="h-[2px] mt-3 mx-auto transition-all ease-in-out duration-200 bg-orange-400"
		style="width: {barWidth}%;"
	/>
</a>
