<script>
	import Card from '$lib/components/Card.svelte';

	import PageWrapper from '$lib/components/PageWrapper.svelte';
	import { Button, Input, Label } from 'flowbite-svelte';
</script>

<PageWrapper>
	<h2 class="text-4xl font-semibold">Project Settings</h2>
	<div class="grid grid-cols-3 gap-6 mt-14">
		<Card class="col-span-2" />
		<div>
			<Card>
				<h4 class="text-lg font-semibold mb-4">Danger Zone</h4>
				<Button color="red">Delete Project</Button>
			</Card>
		</div>
	</div>
</PageWrapper>
