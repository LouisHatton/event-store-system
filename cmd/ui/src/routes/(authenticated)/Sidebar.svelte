<script>
	import ProjectPicker from '$lib/components/ProjectPicker/ProjectPicker.svelte';
	import BarChart from '$lib/icons/BarChart.svelte';
	import DashboardIcon from '$lib/icons/DashboardIcon.svelte';
	import LifeBuoy from '$lib/icons/LifeBuoy.svelte';
	import LogOut from '$lib/icons/LogOut.svelte';
	import Megaphone from '$lib/icons/Megaphone.svelte';
	import Moon from '$lib/icons/Moon.svelte';
	import Settings from '$lib/icons/Settings.svelte';
	import Speedometer from '$lib/icons/Speedometer.svelte';
	import Sun from '$lib/icons/Sun.svelte';
	import { darkMode } from '$lib/stores/darkMode';
	import SidebarButton from './SidebarButton.svelte';
	import SidebarIcon from './SidebarIcon.svelte';
	import { Button, Dropdown, DropdownItem, Chevron, Avatar } from 'flowbite-svelte';

	function toggleDarkMode() {
		darkMode.toggle();
	}
</script>

<div
	class="w-72 border-r border-r-zinc-200 bg-white dark:bg-gray-800 dark:border-r-zinc-600 flex-shrink-0"
>
	<div class="flex flex-col px-6 h-full">
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
				<span class="tracking-wide font-medium text-zinc-400/60 mx-2">MENU</span>
			</div>
			<div class="flex flex-col gap-y-4 my-4">
				<SidebarButton href="/"><DashboardIcon filled slot="svg" />Dashboard</SidebarButton>
				<SidebarButton href="/connections"><Speedometer slot="svg" />Connections</SidebarButton>
				<SidebarButton disabled href="/charts"><BarChart slot="svg" />View Data</SidebarButton>
				<SidebarButton href="/settings"><Settings slot="svg" />Settings</SidebarButton>
			</div>
		</div>
		<div class="border-b border-b-zinc-100 dark:border-b-zinc-800" />
		<div class="my-6">
			<div>
				<span class="tracking-wide font-medium text-zinc-400/60 mx-2">SUPPORT</span>
			</div>
			<div class="flex flex-col gap-y-4 my-4">
				<SidebarButton href="mailto:help@insightwave.co">
					<LifeBuoy filled slot="svg" />
					Get Help
				</SidebarButton>
				<SidebarButton href="mailto:feedback@insightwave.co">
					<Megaphone filled slot="svg" />
					Feedback
				</SidebarButton>
			</div>
		</div>
		<div class="mt-auto">
			<ProjectPicker />
			<div
				class="flex flex-row justify-center gap-x-10 items-center my-4 pt-4 border-t border-gray-200 dark:border-gray-600"
			>
				<SidebarIcon on:click={darkMode.toggle}>
					{#if $darkMode}
						<Sun />
					{:else}
						<Moon />
					{/if}
				</SidebarIcon>
				<SidebarIcon>
					<LogOut />
				</SidebarIcon>
			</div>
		</div>
	</div>
</div>
