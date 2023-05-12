<script>
	import { darkMode } from '$lib/stores/darkMode';
	import '../app.css';
	import Navigating from './Navigating.svelte';
</script>

<main class={`${$darkMode ? 'dark' : ''}`}>
	<Navigating />
	<slot />
</main>
