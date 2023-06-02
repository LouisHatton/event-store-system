<script lang="ts">
	import { Button, Spinner } from 'flowbite-svelte';
	import type { ButtonColorType } from 'flowbite-svelte/dist/types';

	export let loading = false;
	export let color: ButtonColorType;
</script>

<Button disabled={loading} {color} on:click>
	{#if loading}
		<Spinner class="mr-3" size="4" color="white" />
	{/if}
	<slot />
</Button>
