<script lang="ts">
	import { goto } from '$app/navigation';
	import Card from '$lib/components/Card.svelte';
	import PageLoader from '$lib/components/PageLoader.svelte';
	import PageWrapper from '$lib/components/PageWrapper.svelte';
	import ConnectionService, { type Connection } from '$lib/services/ConnectionService';
	import { activeProject } from '$lib/stores/projectStore';
	import {
		Button,
		Spinner,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import NewConnection from './NewConnection.svelte';

	let connections: Connection[] | undefined = undefined;

	$: $activeProject && getConnections($activeProject.id);

	async function getConnections(projectId: string) {
		connections = undefined;
		connections = await ConnectionService.listConnections(projectId);
	}
</script>

<PageWrapper>
	<h2 class="text-4xl font-semibold">Connections</h2>
	<Card class="mt-14">
		<div class="flex flex-row justify-between items-center mb-6">
			<h3 class="text-2xl font-semibold">Your Connections</h3>
			{#if $activeProject}
				<NewConnection projectId={$activeProject.id} disabled={!connections} />
			{/if}
		</div>
		<PageLoader loading={!connections}>
			{#if connections}
				<Table>
					<TableHead>
						<TableHeadCell>Name</TableHeadCell>
						<TableHeadCell>Status</TableHeadCell>
						<TableHeadCell>Events this Month</TableHeadCell>
						<TableHeadCell />
					</TableHead>
					<TableBody>
						{#each connections as connection}
							<TableBodyRow
								class="hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer"
								on:click={() => {
									goto(`/connections/${connection.id}`);
								}}
							>
								<TableBodyCell>{connection.name}</TableBodyCell>
								<TableBodyCell>{connection.status}</TableBodyCell>
								<TableBodyCell>{'-'}</TableBodyCell>
								<TableBodyCell><Button color="light">Edit</Button></TableBodyCell>
							</TableBodyRow>
						{/each}
					</TableBody>
				</Table>
			{/if}
		</PageLoader>
	</Card>
</PageWrapper>
