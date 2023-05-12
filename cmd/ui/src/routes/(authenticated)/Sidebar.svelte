<script>
	import BarChart from '$lib/icons/BarChart.svelte';
	import DashboardIcon from '$lib/icons/DashboardIcon.svelte';
	import LifeBuoy from '$lib/icons/LifeBuoy.svelte';
	import Megaphone from '$lib/icons/Megaphone.svelte';
	import PulseIcon from '$lib/icons/PulseIcon.svelte';
	import Settings from '$lib/icons/Settings.svelte';
	import Speedometer from '$lib/icons/Speedometer.svelte';
	import { darkMode } from '$lib/stores/darkMode';
	import SidebarButton from './SidebarButton.svelte';

	function toggleDarkMode() {
		darkMode.toggle();
	}
</script>

<aside class="w-72 border-r border-r-zinc-200 bg-white dark:bg-zinc-900 dark:border-r-zinc-800">
	<div class="flex flex-col px-6">
		<div class="flex flex-row justify-center items-center gap-x-2 my-8">
			<img
				src="/static/insight-wave-800x800.png"
				alt="logo"
				class="w-10"
				on:keydown={toggleDarkMode}
				on:click={toggleDarkMode}
			/>
			<h1 class="text-2xl text-amber-500 font-semibold">InsightWave</h1>
		</div>
		<div class="my-6">
			<div>
				<span class="tracking-wide font-medium text-zinc-400/60 mx-2">MAIN MENU</span>
			</div>
			<div class="flex flex-col gap-y-4 my-4">
				<SidebarButton href="/"><DashboardIcon filled slot="svg" />Dashboard</SidebarButton>
				<SidebarButton href="/events"><Speedometer slot="svg" />Events</SidebarButton>
				<SidebarButton href="/view-stats"><BarChart slot="svg" />Statistics</SidebarButton>
				<SidebarButton href="/settings"><Settings slot="svg" />Settings</SidebarButton>
			</div>
		</div>
		<div class="border-b border-b-zinc-100 dark:border-b-zinc-800" />
		<div class="my-6">
			<div>
				<span class="tracking-wide font-medium text-zinc-400/60 mx-2">SUPPORT</span>
			</div>
			<div class="flex flex-col gap-y-4 my-4">
				<SidebarButton href="/help"><LifeBuoy filled slot="svg" />Get Help</SidebarButton>
				<SidebarButton href="/feedback"><Megaphone filled slot="svg" />Feedback</SidebarButton>
			</div>
		</div>
	</div>
</aside>
