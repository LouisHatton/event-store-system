<script lang="ts">
	import DataBox from '$lib/components/DataBox.svelte';
	import PageWrapper from '$lib/components/PageWrapper.svelte';
	import ToastHandler from './ToastHandler.svelte';
</script>

<PageWrapper>
	<h2 class="text-4xl font-semibold">Dashboard</h2>
	<div class="mt-14">
		<h3 class="text-2xl font-semibold mb-4">Events Overview</h3>
		<div class="grid lg:grid-cols-4 md:grid-cols-2 gap-4">
			<DataBox />
			<DataBox colour="blue" />
			<DataBox colour="orange" />
			<DataBox colour="red" />
		</div>
	</div>
</PageWrapper>
