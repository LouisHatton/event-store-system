<script>
	import Sidebar from './Sidebar.svelte';
</script>

<section
	class="min-h-screen flex flex-row w-full bg-zinc-100 text-zinc-800 dark:bg-zinc-950 dark:text-zinc-50"
>
	<Sidebar />
	<div>
		<slot />
	</div>
</section>
