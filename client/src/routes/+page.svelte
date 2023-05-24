<script lang="ts">
	import { goto } from "$app/navigation";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import Requests from "$lib/utils/requests/Requests";
	import { onMount } from "svelte";

	onMount(async () => {
		const validSession = await Requests.makeAuthRequest("GET", "auth/session/check", null)
			.then((resp) => {
				return resp.ok;
			})
			.catch(() => false);

		goto(validSession ? "/my/library" : "/sign-in");
	});
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_MAIN)}</title>
</svelte:head>
