<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Album, Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import Player from "$lib/components/Player/index.svelte";
	import { page } from "$app/stores";
	import { playNow, playingQueue, songToPlay } from "../../../../../../store";
	import MusicTile from "$lib/components/Music/MusicTile.svelte";
	import AlbumTile from "$lib/components/Music/AlbumTile.svelte";

	const artistPublicId = $page.params.publicId;
	let albums = new Array<Album>();

	async function playAlbum() {
		albums = await Requests.makeAuthRequest("GET", `library/artist/${artistPublicId}`, null)
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
	<title>{translate(TranslationKeys.TITLE_LIBRARY_MUSIC_ARTISTS)}</title>
</svelte:head>

<main>
	{#if albums}
		<div class="h-[85vh] overflow-y-scroll">
			{#each albums as album}
				<a href={`/my/library/music/albums/${album.publicId}`}>
					<AlbumTile {album} />
				</a>
			{/each}
		</div>
	{:else}
		<Loading />
	{/if}
</main>
