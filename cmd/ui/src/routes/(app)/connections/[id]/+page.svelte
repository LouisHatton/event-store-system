<script>
	import Card from '$lib/components/Card.svelte';
	import CopyText from '$lib/components/CopyText.svelte';
	import PageWrapper from '$lib/components/PageWrapper.svelte';
	import { Button, Input, Label } from 'flowbite-svelte';
</script>

<PageWrapper sm>
	<div class="grid grid-cols-3 gap-6">
		<div class="col-span-2 flex flex-col gap-6">
			<Card class="flex flex-col gap-y-2">
				<h3 class="text-xl font-semibold mb-2">Event URL</h3>
				<CopyText text="https://event.api.insightwave.co/abcd12-123456-bbcca-abdc12d" />
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
					<Input class="mb-4" />
					<div>
						<Button color="yellow">Update</Button>
					</div>
				</div>
				<div class="mt-12">
					<p class="font-semibold mb-4">Danger Zone</p>
					<Button color="red">Delete Connection</Button>
				</div>
			</Card>
		</div>
	</div>
</PageWrapper>
