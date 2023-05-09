<script lang="ts">
	import Button from './Button.svelte';

	let email: string;
	export let noText = false;

	$: valid = validateEmail(email);

	const validateEmail = (email: string) => {
		return String(email)
			.toLowerCase()
			.match(
				/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
			);
	};
</script>

<svelte:head>
	<script src="https://getlaunchlist.com/js/widget-diy.js" defer></script>
</svelte:head>

{#if !noText}
	<div class="text-left mb-4">
		<p class="text-xl text-neutral-700">Sign up for exclusive early access to InsightWave!</p>
	</div>
{/if}
<form
	class="launchlist-form flex flex-row gap-x-2"
	action="https://getlaunchlist.com/s/irFrVk"
	method="POST"
>
	<input
		bind:value={email}
		name="email"
		id="email"
		placeholder="Your Email"
		class="block w-full p-3 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-md focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
	/>
	<Button disabled={!valid}>Notify Me!</Button>
</form>
