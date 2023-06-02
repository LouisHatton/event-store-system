<script lang="ts">
	import Tab from '$lib/components/Tabs/Tab.svelte';
	import TabsRow from '$lib/components/Tabs/TabsRow.svelte';
	import Arrow from '$lib/icons/Arrow.svelte';
	import { page } from '$app/stores';
	import type { Connection } from '$lib/services/ConnectionService';
	import ConnectionService from '$lib/services/ConnectionService';
	import { activeProject } from '$lib/stores/projectStore';
	import { Skeleton } from 'flowbite-svelte';
	import { classnames } from '$lib/util';
	import { goto } from '$app/navigation';
	import type { ApiError } from '$lib/services/NetworkService';
	import { connection } from './store';

	$: getConnection($page.params.id, $activeProject?.id);

	async function getConnection(id: string, projectId = '') {
		if ($connection?.id === id && $connection.projectId === projectId) return;
		if (!projectId) return;
		$connection = undefined;
		try {
			$connection = await ConnectionService.getConnection(projectId, id);
		} catch (err: unknown) {
			let error = err as ApiError;
			if (error.status === 404) {
				await goto('/connections');
				return;
			}
			throw err;
		}
	}

	let titleSpacing = 'mt-8 mb-4';
</script>

<div class="py-4 px-6">
	<a
		href="/connections"
		class="text-orange-400 hover:text-orange-500 dark:hover:text-orange-300 flex flex-row items-center gap-x-2"
		><Arrow direction="left" class="fill-current w-5" />Back to Connections
	</a>
	{#if $connection}
		<h2 class={classnames('text-3xl font-semibold', titleSpacing)}>{$connection.name}</h2>
	{:else}
		<div
			class={classnames(
				'w-80 h-9 bg-zinc-200 dark:bg-zinc-600 animate-pulse rounded-md',
				titleSpacing
			)}
		/>
	{/if}
	<div class="mt-8">
		<TabsRow>
			<Tab title="Overview" href="/connections/{$page.params.id}" />
			<Tab title="Events Log" href="/connections/{$page.params.id}/events" />
		</TabsRow>
	</div>
</div>
<slot />
