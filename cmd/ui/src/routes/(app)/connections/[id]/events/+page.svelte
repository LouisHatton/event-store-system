<script lang="ts">
	import { goto } from '$app/navigation';
	import Card from '$lib/components/Card.svelte';
	import EventDrawer from '$lib/components/EventDrawer.svelte';
	import PageWrapper from '$lib/components/PageWrapper.svelte';
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

	let loading = false;
	let hideEventDrawer = true;
	let selectedEvent: object;

	const events = [
		{
			id: 1,
			name: 'Stripe | New Customer',
			currentMonthTotal: 30
		},
		{ id: 2, name: 'Stripe | New Purchase', currentMonthTotal: 48 },
		{ id: 3, name: 'Vercel | Deployment Failed', currentMonthTotal: 12 }
	];

	function loadMore() {
		loading = true;
		setTimeout(() => {
			loading = false;
		}, 2000);
	}
</script>

<EventDrawer bind:hidden={hideEventDrawer} data={selectedEvent} />
<PageWrapper sm class="w-full">
	<Card class="max-w-full overflow-hidden">
		<Table>
			<TableHead>
				<TableHeadCell>Date</TableHeadCell>
				<TableHeadCell>Time Since Event</TableHeadCell>
				<TableHeadCell />
			</TableHead>
			<TableBody>
				{#each events as event}
					<TableBodyRow
						class="hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer"
						on:click={() => {
							hideEventDrawer = false;
							selectedEvent = event;
						}}
					>
						<TableBodyCell>{event.name}</TableBodyCell>
						<TableBodyCell>Active</TableBodyCell>
						<TableBodyCell><Button color="light">View</Button></TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
		</Table>
		<div class="flex flex-row justify-center mt-6">
			<Button color="light" on:click={loadMore}>
				{#if loading}
					<Spinner class="mr-3" size="4" color="white" />
					Loading
				{:else}
					Load More
				{/if}
			</Button>
		</div>
	</Card>
</PageWrapper>
