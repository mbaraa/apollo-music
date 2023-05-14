<script lang="ts">
	import { goto } from "$app/navigation";
	import type { User } from "$lib/entities";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import Requests from "$lib/utils/requests/Requests";
	import { onMount } from "svelte";
	let user: User = { email: "", password: "" };

	async function signin(): Promise<boolean> {
		return await Requests.makeRequest("POST", "auth/signin/email", user)
			.then((resp) => resp.json())
			.then((resp) => {
				if (resp["data"]["token"]) {
					localStorage.setItem("token", resp["data"]["token"]);
					return true;
				}
				return false;
			})
			.catch(() => false);
	}

	onMount(async () => {
		const validSession = await Requests.makeAuthRequest("GET", "auth/session/check", null)
			.then((resp) => {
				return resp.ok;
			})
			.catch(() => false);

		goto(validSession ? "/library" : "/sign-in");
	});
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_SIGN_IN)}</title>
</svelte:head>

<form
	on:submit={async (e) => {
		e.preventDefault();
		const ok = await signin();
		if (ok) goto("/library");
	}}
>
	<label for="login.email">{translate(TranslationKeys.SIGN_IN_EMAIL)}: </label>
	<input id="login.email" type="email" bind:value={user.email} />
	<br />
	<label for="login.password">{translate(TranslationKeys.SIGN_IN_PASSWORD)}: </label>
	<input id="login.password" type="password" bind:value={user.password} />
	<br />
	<a href="/forgot-password">Forgot your password?</a>
	<br />
	<input type="submit" value={translate(TranslationKeys.SIGN_IN_BUTTON)} />
</form>
