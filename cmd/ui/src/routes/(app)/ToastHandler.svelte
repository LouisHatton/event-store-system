<script lang="ts">
	import { Toast, Button } from 'flowbite-svelte';
	import { slide } from 'svelte/transition';

	let show = true;
	let counter = 6;

	function trigger() {
		show = true;
		counter = 6;
		timeout();
	}

	function timeout(): NodeJS.Timer | undefined {
		if (--counter > 0) return setTimeout(timeout, 1000);
		show = false;
	}
</script>

<Toast color="green" transition={slide} bind:open={show} position="top-right" class="z-50">
	<svelte:fragment slot="icon">
		<svg
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 24 24"
			stroke-width="1.5"
			stroke="currentColor"
			class="w-6 h-6"
			><path
				stroke-linecap="round"
				stroke-linejoin="round"
				d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
			/></svg
		>
	</svelte:fragment>
	Autohide in {counter}s.
</Toast>
