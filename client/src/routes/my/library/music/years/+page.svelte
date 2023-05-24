<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Year, Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import YearTile from "$lib/components/Music/YearTile.svelte";

	let years: Year[];

	onMount(async () => {
		years = await Requests.makeAuthRequest("GET", "library/years", null)
			.then((resp) => resp.json())
			.then((resp) => resp["data"])
			.catch((err) => {
				console.error(err);
			});
	});
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_LIBRARY_MUSIC_YEARS)}</title>
</svelte:head>

<main>
	{#if years}
		<div class="h-[85vh] overflow-y-scroll">
			{#each years as year}
				<a href={`/my/library/music/years/${year.publicId}`}>
					<YearTile {year} />
				</a>
			{/each}
		</div>
	{:else}
		<Loading />
	{/if}
</main>
