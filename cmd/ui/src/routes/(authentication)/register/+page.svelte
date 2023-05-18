<script lang="ts">
	import Alert from '$lib/components/Alert.svelte';
	import AuthenticationService from '$lib/services/AuthenticationService';
	import { validateEmail, validatePassword } from '$lib/util';
	import { Button, Checkbox, Label, Input, Spinner } from 'flowbite-svelte';
	import SignInWithGoogle from '../SignInWithGoogle.svelte';

	let email: string;
	let password: string;
	let passwordConfirm: string;
	let loading = false;

	const passwordRequirementsString =
		'Your password must consist of at least 8 characters, with at least a symbol, upper and lower case letters and a number.';
	let errorMessage = '';

	async function handleClick(e: MouseEvent) {
		let loading = true;
		e.preventDefault();
		errorMessage = '';
		if (!validateEmail(email)) {
			errorMessage = 'Invalid email provided';
			loading = false;
			return;
		}
		if (!validatePassword(password)) {
			errorMessage = passwordRequirementsString;
			loading = false;
			return;
		}
		if (password !== passwordConfirm) {
			errorMessage = 'Passwords do not match';
			loading = false;
			return;
		}
		await AuthenticationService.registerWithPassword(email, password);
	}
</script>

<div class="min-h-screen flex justify-center items-center">
	<div class="p-6 w-2/5 space-y-4 md:space-y-6 sm:p-8">
		<div class="flex flex-row justify-center items-center gap-x-2 my-8">
			<img src="/static/insight-wave-800x800.png" alt="logo" class="w-10" />
			<h1 class="text-2xl text-amber-500 font-semibold">InsightWave</h1>
		</div>
		<form class="flex flex-col space-y-6" action="/">
			<h3 class="text-xl font-medium text-gray-900 dark:text-white p-0">Register</h3>
			{#if errorMessage != ''}
				<Alert>{errorMessage}</Alert>
			{/if}
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
			<Label class="space-y-2">
				<span>Re-enter your password</span>
				<Input
					bind:value={passwordConfirm}
					type="password"
					name="password"
					placeholder="•••••"
					required
				/>
			</Label>
			<Button disabled={loading} on:click={handleClick} type="submit" class="w-full" color="yellow"
				>{#if loading}<Spinner class="mr-3" size="4" color="white" />{/if}Create Account</Button
			>
			<SignInWithGoogle />
			<p class="text-sm font-light text-gray-500 dark:text-gray-400">
				Already have an account? <a
					href="/login"
					class="font-medium text-primary-600 hover:underline dark:text-primary-500">Log In</a
				>
			</p>
		</form>
	</div>
</div>
