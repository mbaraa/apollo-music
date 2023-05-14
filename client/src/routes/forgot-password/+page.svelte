<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import Requests from "$lib/utils/requests/Requests";

	let email: string;

	async function requestPasswordUpdate() {
		await Requests.makeRequest("POST", "auth/password/reset", { userEmail: email })
			.then((resp) => resp.json())
			.then((resp) => console.log(resp))
			.catch((err) => console.error(err));
	}
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_FORGOT_PASSWORD)}</title>
</svelte:head>

<main>
	<form
		on:submit={async (e) => {
			e.preventDefault();
			await requestPasswordUpdate();
		}}
	>
		<label for="resetEmail">{translate(TranslationKeys.FORGOT_PASSWORD_EMAIL)}</label>
		<input id="resetEmail" type="email" bind:value={email} />
		<br />
		<input type="submit" value={translate(TranslationKeys.FORGOT_PASSWORD_BUTTON)} />
	</form>
</main>
