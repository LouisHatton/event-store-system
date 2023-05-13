<script>
	import { Avatar, Button, Chevron, Dropdown, DropdownItem } from 'flowbite-svelte';
	import ProjectIcon from './ProjectIcon.svelte';

	let open = false;
	export let placement = 'top';
</script>

<button color="light" class="w-full"><ProjectIcon chevron /></button>
<Dropdown
	{placement}
	bind:open
	class="w-64 overflow-y-auto py-1 h-48 gap-y-2 flex flex-col rounded-md border-t border-l border-r border-gray-200 dark:border-gray-600"
>
	<ProjectIcon colour="bg-green-400" />
	<ProjectIcon colour="bg-red-400" />
	<ProjectIcon colour="bg-pink-400" />
	<ProjectIcon colour="bg-orange-400" />
	<a
		slot="footer"
		href="/"
		class="flex items-center px-3 py-4 -mb-1 text-sm font-medium text-amber-600 hover:bg-gray-100 dark:bg-gray-700 dark:hover:bg-gray-600 dark:text-amber-500"
	>
		<svg
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 24 24"
			stroke-width="1.5"
			stroke="currentColor"
			class="h-5 w-5 mr-1"
			><path
				stroke-linecap="round"
				stroke-linejoin="round"
				d="M19 7.5v3m0 0v3m0-3h3m-3 0h-3m-2.25-4.125a3.375 3.375 0 11-6.75 0 3.375 3.375 0 016.75 0zM4 19.235v-.11a6.375 6.375 0 0112.75 0v.109A12.318 12.318 0 0110.374 21c-2.331 0-4.512-.645-6.374-1.766z"
			/></svg
		>Create New Project
	</a>
</Dropdown>
