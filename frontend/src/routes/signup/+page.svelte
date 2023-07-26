<script lang="ts">
	import { page } from "$app/stores";

	let user: string;
	let email: string;
	let password: string;
	const signup = async () => {
		console.log("HITTING ENDPOINT");
		const response = await fetch(`${$page.url.origin}/auth/signup`, {
			method: "POST",
			credentials: "include",
			headers: {
				"Content-Type": "application/json",
			},

			body: JSON.stringify({ user:user, email: email, password: password }),
		});

		console.log(await response.json());
	};
</script>

<div class="flex flex-col w-full justify-center items-center">
	<div
		class=" flex flex-col space-y-4 p-6 bg-white rounded-xl bg-opacity-20 h-2/3 w-4/12"
	>
		<h1 class="text-3xl underline">Signup</h1>
		<form
			on:submit={(e) => {
				e.preventDefault();
				signup();
			}}
			class="flex flex-col space-y-6 items-center"
		>
			<div class="flex flex-col w-full cursor-none">
				<label for="user" class="text-xl">Username</label>
				<input
					name="user"
					type="text"
					bind:value={user}
					class="rounded-xl p-2"
				/>
			</div>
			<div class="flex flex-col w-full cursor-none">
				<label for="email" class="text-xl">Email</label>
				<input
					name="email"
					type="email"
					bind:value={email}
					class="rounded-xl p-2"
				/>
			</div>
			<div class="flex flex-col w-full cursor-none">
				<label for="password" class="text-xl">Password</label>
				<input
					name="password"
					type="password"
					bind:value={password}
					class="rounded-xl p-2"
				/>
			</div>
			<button class="">Login</button>
			<a href="/signup" class="">Don't have an account? Sign Up!</a>
		</form>
	</div>
</div>
