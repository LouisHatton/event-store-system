<script>
	import Button from '$lib/components/Button.svelte';
	import TwitterIcon from '$lib/components/icons/TwitterIcon.svelte';
	import '../app.postcss';
</script>

<header
	class="flex flex-row justify-between items-center mx-4 py-6 border-b border-b-neutral-200/80"
>
	<div class="flex flex-row items-center gap-x-3">
		<img src="/static/insight-wave-800x800.png" alt="logo" class="w-10" />
		<h1 class="text-2xl text-amber-500 font-semibold">InsightWave</h1>
	</div>
	<Button light shadow>Sign In</Button>
</header>
<slot />
<footer class="bg-white mt-32 mb-12 max-w-5xl mx-auto px-4">
	<section class="grid grid-cols-5 md:grid-cols-6">
		<div class="col-span-2 flex flex-col gap-y-2">
			<img src="/static/insight-wave-800x800.png" alt="logo" class="w-10" />
			<h5 class="text-2xl text-zinc-800 font-semibold">InsightWave</h5>
			<p class="text-md text-neutral-600">
				Realtime metrics for your internal tools, viweing realtime stats from any source which
				provides a webhook
			</p>
		</div>
		<div class="md:col-span-2" />
		<div class="col-span-2 flex flex-col gap-y-8 mt-auto items-end">
			<div class="flex flex-col text-right gap-y-4">
				<a href="/#" class="hover:underline">Privacy Policy</a>
				<a href="/#" class="hover:underline">Terms of Service</a>
			</div>
			<div>
				<a href="https://twitter.com/thedev_louis" target="_blank">
					<TwitterIcon class="w-6 fill-current text-neutral-600 hover:text-neutral-700" />
				</a>
			</div>
		</div>
	</section>
</footer>
