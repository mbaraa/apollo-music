<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Album, Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import Player from "$lib/components/Player/index.svelte";
	import { page } from "$app/stores";
	import { currentAlbumCover, playNow, playingQueue, songToPlay } from "$lib/store";
	import MusicTile from "$lib/components/Music/MusicTile.svelte";

	const albumPublicId = $page.params.publicId;

	let album: Album;

	async function playAlbum() {
		album = await Requests.makeAuthRequest("GET", `library/album/${albumPublicId}`, null)
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
	<title>{`${album?.title} - ${translate(TranslationKeys.TITLE_LIBRARY_MUSIC_ALBUMS)}`}</title>
</svelte:head>

<main>
	{#if album && album.songs}
		<div class="h-full text-dark-secondary">
			{#each album.songs as song}
				<button
					class="block w-full"
					on:click={() => {
						playNow.set(true);
						songToPlay.set(song);
						playingQueue.set(album.songs);
						currentAlbumCover.set(`data:image/*;base64,${album.coverB64}`);
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
