<script>
	import Card from '$lib/components/Card.svelte';

	import PageWrapper from '$lib/components/PageWrapper.svelte';
	import { Button, Input, Label } from 'flowbite-svelte';
</script>

<PageWrapper>
	<h2 class="text-4xl font-semibold">Settings</h2>
	<div class="grid grid-cols-3 gap-6 mt-14">
		<Card class="col-span-2">
			<h3 class="text-2xl font-semibold">Your Account</h3>
			<div class="mt-6">
				<div>
					<Label class="mb-2">Your Name:</Label>
					<Input class="mb-4" />
				</div>
				<div>
					<Label class="mb-2">Email Address:</Label>
					<Input class="mb-4" />
				</div>
				<div class="flex flex-row justify-between">
					<Button color="yellow">Save</Button>
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
