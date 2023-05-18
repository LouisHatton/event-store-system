<script>
	import GoogleLogo from '$lib/icons/GoogleLogo.svelte';
	import AuthenticationService from '$lib/services/AuthenticationService';

	async function handleClick() {
		await AuthenticationService.signInWithGoogle();
	}
</script>

<button
	on:click={handleClick}
	class="border border-zinc-200 dark:border-gray-600 dark:hover:bg-gray-600 transition-all ease-out rounded-md py-3 hover:bg-zinc-100 flex flex-row justify-center gap-x-3 items-center"
>
	<GoogleLogo class="h-6 w-6" /> Sign In with Google
</button>
