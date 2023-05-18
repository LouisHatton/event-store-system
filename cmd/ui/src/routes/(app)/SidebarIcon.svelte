<script lang="ts">
	import { page } from '$app/stores';
	import { classnames } from '$lib/util';

	export let href: string | undefined = undefined;

	$: active = $page.url.pathname === href;

	$: activeClass = active ? 'bg-amber-500/10' : '';
	let baseClass =
		'text-left cursor-pointer flex flex-row items-center  hover:bg-amber-500/10 transition-all ease-in-out p-4 rounded-full text-zinc-800 font-semibold text-lg dark:text-zinc-100';
</script>

{#if href}
	<a {href} class={classnames(baseClass, activeClass)} on:click>
		<div class="h-6 w-6 fill-current">
			<slot />
		</div>
	</a>
{:else}
	<button class={classnames(baseClass, activeClass)} on:click>
		<div class="h-6 w-6 fill-current">
			<slot />
		</div>
	</button>
{/if}
