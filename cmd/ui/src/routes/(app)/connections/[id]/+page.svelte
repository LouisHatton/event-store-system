<script lang="ts">
	import Card from '$lib/components/Card.svelte';
	import CopyText from '$lib/components/CopyText.svelte';
	import PageLoader from '$lib/components/PageLoader.svelte';
	import PageWrapper from '$lib/components/PageWrapper.svelte';
	import { Button, Input, Label } from 'flowbite-svelte';
	import { connection } from './store';
	import LoadingButton from '$lib/components/LoadingButton.svelte';
	import ConnectionService, { type NewConnection } from '$lib/services/ConnectionService';
	import { goto } from '$app/navigation';

	let connectionName = '';
	let updatingConnection = false;

	$: initConnectionName($connection?.name ?? '');
	const initConnectionName = (name: string) => {
		connectionName = name;
	};

	async function updateConnection() {
		if (!$connection) return;
		updatingConnection = true;
		let updatedConnection: NewConnection = {
			name: connectionName,
			tags: $connection.tags
		};
		try {
			$connection = await ConnectionService.set(
				$connection.projectId,
				$connection.id,
				updatedConnection
			);
		} catch {
			alert('failed to update connection, please try again later');
		} finally {
			updatingConnection = false;
		}
	}

	async function handleDeleteConnection() {
		if (!$connection) return;
		try {
			await ConnectionService.delete($connection.projectId, $connection.id);
			goto('/connections');
		} catch (err: unknown) {
			console.log(err);
			alert('error deleting connection, please try again later');
		}
	}
</script>

<PageWrapper sm>
	<PageLoader loading={!$connection}>
		{#if $connection}
			<div class="grid grid-cols-3 gap-6">
				<div class="col-span-2 flex flex-col gap-6">
					<Card class="flex flex-col gap-y-2">
						<h3 class="text-xl font-semibold mb-2">Event URL</h3>
						<CopyText text="https://event.api.insightwave.co/{$connection.urlId}" />
					</Card>
					<Card class="flex flex-col gap-y-2">
						<h3 class="text-xl font-semibold mb-2">Schema</h3>
						<p class="opacity-90">When we receive your first event we will show the schema here!</p>
					</Card>
				</div>
				<div>
					<Card>
						<h4 class="text-lg font-semibold mb-4">Settings</h4>
						<div>
							<Label class="mb-2">Connection Name:</Label>
							<Input class="mb-4" bind:value={connectionName} />
							<div>
								<LoadingButton
									loading={updatingConnection}
									color="yellow"
									on:click={updateConnection}>Update</LoadingButton
								>
							</div>
						</div>
						<div class="mt-12">
							<p class="font-semibold mb-4">Danger Zone</p>
							<Button color="red" on:click={handleDeleteConnection}>Delete Connection</Button>
						</div>
					</Card>
				</div>
			</div>
		{/if}
	</PageLoader>
</PageWrapper>
