<script lang="ts">
	import LaunchListSignup from '$lib/components/LaunchListSignup.svelte';
	import { Badge, Indicator } from 'flowbite-svelte';
</script>

<svelte:head>
	<title>InsightWave | Real Time Metrics for Internal Tools</title>
</svelte:head>

<main class="bg-gradient-to-b from-white via-slate-100 to-slate-100 px-4">
	<seciton class="flex flex-col items-center justify-between container max-w-5xl mx-auto">
		<div class="text-left mx-auto my-20 md:my-32">
			<Badge color="dark" rounded class="px-2.5 py-0.5" large>
				<Indicator color="yellow" size="md" class="mr-1" />Currently in Alpha!
			</Badge>
			<h1 class="text-6xl font-semibold text-neutral-900 py-6 max-w-3xl">
				Realtime metrics for your <span
					class="text-transparent bg-clip-text bg-gradient-to-r from-amber-500 to-orange-500"
					>internal tools.</span
				>
			</h1>
			<p class="text-2xl text-neutral-600">
				Monitor events and actions fron your internal tools. Embed customizable dashboards to
				visualize your metrics in real-time and optimize your business for success.
			</p>
			<div class="max-w-xl mr-auto mt-16">
				<LaunchListSignup />
			</div>
		</div>
	</seciton>
	<section class="my-20 md:my-32 max-w-5xl mx-auto">
		<div class="grid md:grid-cols-2 gap-x-12 gap-y-8 items-center">
			<div class="flex flex-col gap-y-8">
				<h2 class="text-5xl font-semibold text-neutral-800">Receive events from any source.</h2>
				<p class="text-2xl text-neutral-600">
					Connect any platorm which provides a webhook and start visualizing the events in
					real-time. Export this data into a beautiful chart and embed anywhere!
				</p>
			</div>
			<div>
				<img
					alt="Data from stripe, vercel and mail chimp flowing into insight wave"
					src="/static/data-in-insight-wave.png"
				/>
			</div>
		</div>
	</section>
	<section class="max-w-5xl mx-auto py-12">
		<div>
			<h2 class="text-5xl font-semibold text-neutral-800">Visibility Everywhere.</h2>
			<p class="text-2xl text-neutral-600 mt-4">
				Keep in the loop on what is happening accross your business
			</p>
		</div>
		<div class="grid md:grid-cols-2 grid-rows-2 gap-4 md:gap-8 mt-10 md:mt-12">
			<div class="bg-white rounded-md px-6 py-10">
				<p class="text-4xl mb-4 -ml-2">⏱️</p>
				<h4 class="text-3xl font-semibold text-neutral-800">Realtime updates</h4>
				<p class="text-xl text-neutral-600 mt-4">
					Watch events as they come in or allow us to notify you if something is unexpected.
				</p>
			</div>
			<div class="bg-white rounded-md px-6 py-10">
				<p class="text-4xl mb-4 -ml-2">📈</p>
				<h4 class="text-3xl font-semibold text-neutral-800">Overview Dashboards</h4>
				<p class="text-xl text-neutral-600 mt-4">
					View what's happening across the board and customize to suit you.
				</p>
			</div>
			<div class="bg-white rounded-md px-6 py-10">
				<p class="text-4xl mb-4 -ml-2">🖼️</p>
				<h4 class="text-3xl font-semibold text-neutral-800">Embed Charts</h4>
				<p class="text-xl text-neutral-600 mt-4">
					Embed your Dashboards and charts into Confluence, Notion etc with realtime updates.
				</p>
			</div>
			<div class="bg-white rounded-md px-6 py-10">
				<p class="text-4xl mb-4 -ml-2">♾️</p>
				<h4 class="text-3xl font-semibold text-neutral-800">Unlimited Sources</h4>
				<p class="text-xl text-neutral-600 mt-4">
					We won't limit how many services you can track, just pay for what you use.
				</p>
			</div>
		</div>
	</section>
	<section class="max-w-5xl mx-auto py-32">
		<div class="mb-6">
			<h2 class="text-5xl font-semibold text-neutral-800">Join the waitlist</h2>
			<p class="text-2xl text-neutral-600 mt-4">
				Join the waitlist below and get beta access to InsightWave when its ready:
			</p>
		</div>
		<div class="max-w-2xl">
			<LaunchListSignup noText />
		</div>
	</section>
</main>
