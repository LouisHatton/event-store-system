<script lang="ts">
	import { page } from '$app/stores';
	import { classnames } from '$lib/util';

	export let href: string;

	$: active = $page.url.pathname === href;

	$: activeClass = active ? 'bg-amber-500/10' : '';
	let baseClass =
		'text-left cursor-pointer flex flex-row items-center  hover:bg-amber-500/10 transition-all ease-in-out px-4 py-3 rounded-2xl text-zinc-800 font-semibold text-lg dark:text-zinc-100';
</script>

<a {href} class={classnames(baseClass, activeClass)}>
	<div class="h-6 w-6 mr-3 fill-current">
		<slot name="svg" />
	</div>
	<slot />
</a>
