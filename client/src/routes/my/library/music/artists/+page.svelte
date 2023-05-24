<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Artist } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import ArtistTile from "$lib/components/Music/ArtistTile.svelte";

	let artists: Artist[];

	onMount(async () => {
		artists = await Requests.makeAuthRequest("GET", "library/artists", null)
			.then((resp) => resp.json())
			.then((resp) => resp["data"])
			.catch((err) => {
				console.error(err);
			});
	});
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_LIBRARY)}</title>
</svelte:head>

<main>
	{#if artists}
		<div class="h-[85vh] overflow-y-scroll">
			{#each artists as artist}
				<a href={`/my/library/music/artists/${artist.publicId}`}>
					<ArtistTile {artist} />
				</a>
			{/each}
		</div>
	{:else}
		<Loading />
	{/if}
</main>
