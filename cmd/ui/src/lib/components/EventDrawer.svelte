<script lang="ts">
	import { Button, CloseButton, Drawer } from 'flowbite-svelte';
	import { sineIn } from 'svelte/easing';

	export let hidden = false;
	export let data: object = {};
	let transitionParamsRight = {
		x: 600,
		duration: 300,
		easing: sineIn
	};
</script>

<Drawer
	placement="right"
	transitionType="fly"
	transitionParams={transitionParamsRight}
	bind:hidden
	id="sidebar"
	class="w-4/5 md:w-3/5 xl:w-2/5"
>
	<div class="flex items-center">
		<h5
			id="drawer-label"
			class="inline-flex items-center mb-4 text-xl font-semibold text-zinc-900 dark:text-white"
		>
			<svg
				class="w-6 h-6 mr-2"
				aria-hidden="true"
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path
					fill-rule="evenodd"
					d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
					clip-rule="evenodd"
				/></svg
			>Event at 2023-05-10 13:15
		</h5>
		<CloseButton on:click={() => (hidden = true)} class="mb-4 dark:text-white" />
	</div>
	<div>
		<pre>
			{JSON.stringify(data, undefined, 2)}
		</pre>
	</div>
</Drawer>
