<script lang="ts">
	import Card from '$lib/components/Card.svelte';
	import type { User } from 'firebase/auth';

	import PageWrapper from '$lib/components/PageWrapper.svelte';
	import { authStore } from '$lib/stores/authStore';
	import { Button, Input, Label } from 'flowbite-svelte';
	import UserService from '$lib/services/UserService';
	import LoadingButton from '$lib/components/LoadingButton.svelte';

	let emailAddress = '';
	let name = '';
	let saving = false;
	let loadingData = true;

	$: $authStore.user ? updateVars($authStore.user) : (loadingData = true);

	function updateVars(user: User) {
		loadingData = false;
		if (user.email) {
			emailAddress = user.email;
		}
		if (user.displayName) {
			name = user.displayName;
		}
	}

	async function hanldeClick() {
		saving = true;
		await UserService.updateUser(emailAddress, name);
		saving = false;
	}
</script>

<PageWrapper>
	<h2 class="text-4xl font-semibold">Settings</h2>
	<div class="grid grid-cols-3 gap-6 mt-14">
		<Card class="col-span-2">
			<h3 class="text-2xl font-semibold">Your Account</h3>
			<div class="mt-6">
				<div>
					<Label class="mb-2">Your Name:</Label>
					<Input disabled={loadingData} class="mb-4" bind:value={name} />
				</div>
				<div>
					<Label class="mb-2">Email Address:</Label>
					<Input disabled={loadingData} class="mb-4" bind:value={emailAddress} />
				</div>
				<div class="flex flex-row justify-between">
					<LoadingButton color="yellow" loading={saving} on:click={hanldeClick}>Save</LoadingButton>
					<Button color="light">Change Password</Button>
				</div>
			</div>
		</Card>
		<div>
			<Card>
				<h4 class="text-lg font-semibold mb-4">Danger Zone</h4>
				<Button color="red">Delete Account</Button>
			</Card>
		</div>
	</div>
</PageWrapper>
