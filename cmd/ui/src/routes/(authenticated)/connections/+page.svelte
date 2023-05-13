<script lang="ts">
	import { goto } from '$app/navigation';
	import Card from '$lib/components/Card.svelte';
	import PageWrapper from '$lib/components/PageWrapper.svelte';
	import {
		Button,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';

	const eventSources = [
		{
			id: 1,
			name: 'Stripe | New Customer',
			currentMonthTotal: 30
		},
		{ id: 2, name: 'Stripe | New Purchase', currentMonthTotal: 48 },
		{ id: 3, name: 'Vercel | Deployment Failed', currentMonthTotal: 12 }
	];
</script>

<PageWrapper>
	<h2 class="text-4xl font-semibold">Connections</h2>
	<Card class="mt-14">
		<div class="flex flex-row justify-between items-center mb-6">
			<h3 class="text-2xl font-semibold">Your Connections</h3>
			<Button color="yellow">Add New</Button>
		</div>
		<Table>
			<TableHead>
				<TableHeadCell>Name</TableHeadCell>
				<TableHeadCell>Status</TableHeadCell>
				<TableHeadCell>Events this Month</TableHeadCell>
				<TableHeadCell />
			</TableHead>
			<TableBody>
				{#each eventSources as event}
					<TableBodyRow
						class="hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer"
						on:click={() => {
							goto(`/connections/${event.id}`);
						}}
					>
						<TableBodyCell>{event.name}</TableBodyCell>
						<TableBodyCell>Active</TableBodyCell>
						<TableBodyCell>{event.currentMonthTotal}</TableBodyCell>
						<TableBodyCell><Button color="light">Edit</Button></TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
		</Table>
	</Card>
</PageWrapper>
