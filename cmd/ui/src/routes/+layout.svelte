<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import AuthenticationService from '$lib/services/AuthenticationService';
	import { authStore } from '$lib/stores/authStore';
	import { darkMode } from '$lib/stores/darkMode';
	import { Spinner } from 'flowbite-svelte';
	import '../app.css';
	import ToastHandler from './(app)/ToastHandler.svelte';
	import Navigating from './Navigating.svelte';
	import LoadingScreen from '$lib/components/LoadingScreen.svelte';
	$: $page.route && checkIfLoggedIn($authStore.initialised);
	let loading = true;

	async function checkIfLoggedIn(init: boolean) {
		loading = true;
		if (!init) return;
		try {
			const isInAuthRoute = $page.route.id?.includes('(authentication)');
			const token = await AuthenticationService.getToken();
			if (!token && !isInAuthRoute) {
				await goto('/login');
			} else if (token && isInAuthRoute) {
				await goto('/');
			}
		} finally {
			loading = false;
		}
	}
</script>

<main class={`${$darkMode ? 'dark' : ''}`}>
	<Navigating />
	<ToastHandler />
	{#if loading}
		<LoadingScreen />
	{:else}
		<slot />
	{/if}
</main>
