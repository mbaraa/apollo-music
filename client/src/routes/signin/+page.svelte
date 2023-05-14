<script lang="ts">
	import { goto } from "$app/navigation";
	import type { User } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	let user: User = { email: "", password: "" };

	async function signin() {
		await Requests.makeRequest("POST", "auth/signin/email", user)
			.then((resp) => resp.json())
			.then((resp) => {
				localStorage.setItem("token", resp["data"]["token"]);
			})
			.catch((err) => console.error(err));
	}
</script>

<form
	on:submit={(e) => {
		e.preventDefault();
		signin();
		goto("/library");
	}}
>
	<label for="login.email">Email: </label>
	<input id="login.email" type="email" bind:value={user.email} />
	<br />
	<label for="login.password">Password: </label>
	<input id="login.password" type="password" bind:value={user.password} />
	<br />
	<input type="submit" />
</form>
