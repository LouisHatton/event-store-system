<script lang="ts">
	import AuthenticationService from '$lib/services/AuthenticationService';
	import { Button, Label, Input, Spinner } from 'flowbite-svelte';
	import SignInWithGoogle from '../SignInWithGoogle.svelte';

	let email: string;
	let password: string;
	let loading = false;

	async function handleLogin(e: MouseEvent) {
		e.preventDefault();
		loading = true;
		await AuthenticationService.signInWithPassword(email, password);
		console.log(AuthenticationService.currentUser);
		loading = false;
	}
</script>

<div class="min-h-screen flex justify-center items-center">
	<div class="p-6 w-4/5 lg:w-2/5 space-y-4 md:space-y-6 sm:p-8">
		<div class="flex flex-row justify-center items-center gap-x-2 my-8">
			<img src="/static/insight-wave-800x800.png" alt="logo" class="w-10" />
			<h1 class="text-2xl text-amber-500 font-semibold">InsightWave</h1>
		</div>
		<form class="flex flex-col space-y-6" action="/">
			<h3 class="text-xl font-medium text-gray-900 dark:text-white p-0">Login</h3>
			<Label class="space-y-2">
				<span>Your email</span>
				<Input
					bind:value={email}
					type="email"
					name="email"
					placeholder="name@company.com"
					required
				/>
			</Label>
			<Label class="space-y-2">
				<span>Your password</span>
				<Input bind:value={password} type="password" name="password" placeholder="•••••" required />
			</Label>
			<div class="flex items-start">
				<a
					href="/register"
					class="ml-auto text-sm text-orange-500 hover:underline dark:text-orange-400"
					>Forgot password?</a
				>
			</div>
			<Button disabled={loading} on:click={handleLogin} type="submit" class="w-full" color="yellow"
				>{#if loading}<Spinner class="mr-3" size="4" color="white" />{/if}Sign in</Button
			>
			<SignInWithGoogle />
			<p class="text-sm font-light text-gray-500 dark:text-gray-400">
				Don’t have an account yet? <a
					href="/register"
					class="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign up</a
				>
			</p>
		</form>
	</div>
</div>
