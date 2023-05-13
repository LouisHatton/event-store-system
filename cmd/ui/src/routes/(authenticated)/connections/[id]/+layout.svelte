<script>
	import Tab from '$lib/components/Tabs/Tab.svelte';
	import TabsRow from '$lib/components/Tabs/TabsRow.svelte';
	import Arrow from '$lib/icons/Arrow.svelte';
	import { page } from '$app/stores';
</script>

<div class="py-4 px-6">
	<a
		href="/connections"
		class="text-orange-400 hover:text-orange-500 dark:hover:text-orange-300 flex flex-row items-center gap-x-2"
		><Arrow direction="left" class="fill-current w-5" />Back to Connections
	</a>
	<h2 class="text-3xl font-semibold mt-8 mb-4">Stripe | New Customer</h2>
	<div class="mt-8">
		<TabsRow>
			<Tab title="Overview" href="/connections/{$page.params.id}" />
			<Tab title="Events Log" href="/connections/{$page.params.id}/events" />
		</TabsRow>
	</div>
</div>
<slot />
