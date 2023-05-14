<script lang="ts">
	import { page } from "$app/stores";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import Requests from "$lib/utils/requests/Requests";

	let newPassword: string;

	async function updatePassword() {
		await Requests.makeRequest(
			"POST",
			"auth/password/update",
			{ newPassword: newPassword },
			{},
			{
				Authorization: $page.params.passwordResetToken
			}
		)
			.then((resp) => resp.json())
			.then((resp) => console.log(resp))
			.catch((err) => console.error(err));
	}
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_RESET_PASSWORD)}</title>
</svelte:head>

<main>
	<form
		on:submit={async (e) => {
			e.preventDefault();
			await updatePassword();
		}}
	>
		<label for="newPassword">{translate(TranslationKeys.PASSWORD_RESET_NEW_PASSWORD)}</label>
		<input id="newPassword" type="password" bind:value={newPassword} />
		<br />
		<input type="submit" value={translate(TranslationKeys.PASSWORD_RESET_UPDATE_BUTTON)} />
	</form>
</main>
