<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import type { ButtonColor } from './flowbiteUtils';

	export let light = false;
	export let shadow = false;

	$: tailwindColours = calculateColours(light, shadow);

	function calculateColours(light: boolean, shadow: boolean) {
		let className = '';
		if (light) {
			className += 'bg-white text-zinc-900 border-gray-200 border hover:bg-slate-50 ';
			if (shadow) className += 'shadow-md shadow-amber-100 ';
		}
		className += 'bg-amber-500 hover:bg-amber-600 text-white ';

		return className;
	}
</script>

<button
	class={`px-6 py-2 transition-all ease-in-out duration-300 font-semibold rounded-md flex-shrink-0 cursor-pointer ` +
		tailwindColours}
	{...$$props}><slot /></button
>
