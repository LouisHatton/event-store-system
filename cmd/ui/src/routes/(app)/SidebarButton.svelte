<script lang="ts">
	import { page } from '$app/stores';
	import { classnames } from '$lib/util';

	export let href: string;
	export let disabled = false;

	$: active = href === '/' ? $page.url.pathname === href : $page.url.pathname.startsWith(href);

	$: console.log($page.url.pathname, activeClass);
	$: console.log(active);

	$: activeClass = active ? 'bg-amber-500/10' : '';
	let hoverClass = 'cursor-pointer hover:bg-amber-500/10 transition-all ease-in-out';
	let svgClass = 'h-6 w-6 mr-3 fill-current text-zinc-700 dark:text-zinc-200';
	let baseClass =
		'text-left flex flex-row items-center px-4 py-3 rounded-2xl text-zinc-800 font-medium text-lg dark:text-zinc-100';
</script>

{#if disabled}
	<div class={classnames(baseClass, 'cursor-default opacity-50')}>
		<div class={svgClass}>
			<slot name="svg" />
		</div>
		<slot />
	</div>
{:else}
	<a {href} class={classnames(baseClass, hoverClass, activeClass)}>
		<div class={svgClass}>
			<slot name="svg" />
		</div>
		<slot />
	</a>
{/if}
