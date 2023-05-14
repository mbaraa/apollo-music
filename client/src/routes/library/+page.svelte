<script lang="ts">
	import { goto } from "$app/navigation";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";

	onMount(() => {
		goto("/library/music/albums");
	});
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_LIBRARY)}</title>
</svelte:head>
