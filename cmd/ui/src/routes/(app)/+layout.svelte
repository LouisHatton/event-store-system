<script>
	import Sidebar from './Sidebar.svelte';
</script>

<section
	class="min-h-screen max-w-full flex flex-row w-full bg-zinc-50 text-zinc-800 dark:bg-gray-700 dark:text-zinc-50"
>
	<Sidebar />

	<div class="w-full relative">
		<slot />
	</div>
</section>
