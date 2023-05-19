<script lang="ts">
	import { goto } from "$app/navigation";
	import type { User } from "$lib/entities";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import Input from "$lib/ui/Input.svelte";
	import Link from "$lib/ui/Link.svelte";
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

<main class="bg-dark-primary text-dark-secondary w-[100vw] h-[100vh]">
	<h1 class="font-[Comfortaa] pt-[32px] px-[35px] w-full font-[500] text-[30px]">
		{translate(TranslationKeys.SIGN_IN_HEADER)}
	</h1>
	<form
		class="pt-[66px] pl-[35px]"
		on:submit={async (e) => {
			e.preventDefault();
			const ok = await signin();
			if (ok) goto("/library");
		}}
	>
		<Input
			type="email"
			bind:value={user.email}
			placeholder={translate(TranslationKeys.SIGN_IN_EMAIL)}
			required
		/>
		<Input
			type="password"
			bind:value={user.password}
			placeholder={translate(TranslationKeys.SIGN_IN_PASSWORD)}
			required
			_class="mt-[24px]"
		/>
		<Link
			href="/forgot-password"
			title={translate(TranslationKeys.SIGN_IN_FORGOT_PASSWORD)}
			_class="float-right mt-[12px] mr-[35px] text-[16px]"
		/>

		<input
			type="submit"
			class="bg-dark-accent text-dark-neutral w-[330px] h-[48px] rounded-[20px] mt-[44px] text-[24px] font-IBMPlexSans cursor-pointer"
			value={translate(TranslationKeys.SIGN_IN_BUTTON)}
		/>
	</form>
	<div>
		<h3 class="font-IBMPlexSans text-dark-secondary text-[18px] w-full text-center mt-[30px]">
			{translate(TranslationKeys.SIGN_IN_SIGN_UP)}
			<Link
				href="https://checkout.apollo-music.app/sign-up"
				title={translate(TranslationKeys.SIGN_IN_SIGN_UP_HERE)}
			/>
		</h3>
	</div>
</main>
