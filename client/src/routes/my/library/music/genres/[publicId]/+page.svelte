<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Album, Genre, Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import Player from "$lib/components/Player/index.svelte";
	import { page } from "$app/stores";
	import { currentAlbumCover, playNow, playingQueue, songToPlay } from "../../../../../../store";
	import MusicTile from "$lib/components/Music/MusicTile.svelte";

	const genrePublicId = $page.params.publicId;

	let genre: Genre;

	async function playAlbum() {
		genre = await Requests.makeAuthRequest("GET", `library/genre/${genrePublicId}`, null)
			.then((resp) => resp.json())
			.then((resp) => resp["data"])
			.catch((err) => {
				console.error(err);
			});
	}
	onMount(() => {
		playAlbum();
	});
</script>

<svelte:head>
	<title>{`${genre?.name} - ${translate(TranslationKeys.TITLE_LIBRARY_MUSIC_GENRERS)}`}</title>
</svelte:head>

<main>
	{#if genre && genre.songs}
		<div class="h-[85vh] overflow-y-scroll">
			{#each genre.songs as song}
				<button
					class="block w-full"
					on:click={() => {
						playNow.set(true);
						songToPlay.set(song);
						playingQueue.set(genre.songs);
					}}
				>
					<MusicTile music={song} />
				</button>
			{/each}
		</div>
	{:else}
		<Loading />
	{/if}
</main>
