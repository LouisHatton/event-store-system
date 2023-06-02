<script lang="ts">
	import { goto } from '$app/navigation';
	import LoadingButton from '$lib/components/LoadingButton.svelte';
	import type { NewConnection } from '$lib/services/ConnectionService';
	import ConnectionService from '$lib/services/ConnectionService';
	import type { ApiError } from '$lib/services/NetworkService';
	import { Drawer, Button, CloseButton, Input, Label } from 'flowbite-svelte';
	import { createEventDispatcher } from 'svelte';
	import { sineIn } from 'svelte/easing';

	let hidden = true;
	let transitionParamsRight = {
		x: 320,
		duration: 200,
		easing: sineIn
	};

	export let disabled = false;
	export let projectId: string;

	let name = '';
	let loading = false;

	async function createConnection() {
		loading = true;
		let newConnection: NewConnection = {
			name,
			tags: []
		};
		try {
			let connection = await ConnectionService.create(projectId, newConnection);
			goto(`/connections/${connection.id}`);
		} catch (e: unknown) {
			alert('There was an error processing your request');
		} finally {
			loading = false;
		}
	}
</script>

<Button on:click={() => (hidden = false)} color="yellow" {disabled}>Add New</Button>
<Drawer
	placement="right"
	transitionType="fly"
	width="w-[500px]"
	transitionParams={transitionParamsRight}
	bind:hidden
	id="sidebar"
>
	<div class="py-2 px-4">
		<div class="flex items-center">
			<h5
				id="drawer-label"
				class="inline-flex items-center mb-4 text-2xl font-semibold text-zinc-800 dark:text-gray-200"
			>
				New Connection
			</h5>
			<CloseButton on:click={() => (hidden = true)} class="mb-4 dark:text-white" />
		</div>
		<div class="my-6 flex flex-col">
			<Label class="mb-2">Connection Name:</Label>
			<Input class="mb-4" bind:value={name} />

			<LoadingButton {loading} color="yellow" on:click={createConnection}>Create</LoadingButton>
		</div>
	</div>
</Drawer>
