<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Album, Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import AlbumTile from "$lib/components/music/AlbumTile.svelte";
	import Player from "$lib/components/Player/index.svelte";

	let albums: Album[];
	$: songs = new Array<Music>();

	async function playAlbum(album: Album) {
		songs = await Requests.makeAuthRequest("GET", `library/album/${album.publicId}`, null)
			.then((resp) => resp.json())
			.then((resp) => resp["data"]["songs"])
			.catch((err) => {
				console.error(err);
			});
	}

	onMount(async () => {
		albums = await Requests.makeAuthRequest("GET", "library/albums", null)
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
	{#if albums}
		<div class="h-[85vh] overflow-y-scroll">
			{#each albums as album}
				<div on:click={() => playAlbum(album)}>
					<AlbumTile {album} />
				</div>
			{/each}
		</div>
		<Player playlist={songs} on:playlistchange={() => {}} />
	{:else}
		<Loading />
	{/if}
</main>
