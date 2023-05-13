<script lang="ts">
	import { classnames } from '$lib/util';

	export let sm = false;
</script>

{#if sm}
	<div class={classnames('px-6 py-4 w-full', $$props.class)}><slot /></div>
{:else}
	<div class="px-6 py-10 w-full"><slot /></div>
{/if}
