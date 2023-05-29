<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Album, Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import AlbumTile from "$lib/components/Music/AlbumTile.svelte";

	let albums: Album[];

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
	<title>{translate(TranslationKeys.TITLE_LIBRARY_MUSIC_ALBUMS)}</title>
</svelte:head>

<main>
	{#if albums}
		<div class="px-[25px] font-IBMPlexSans text-dark-secondary">
			<h1 class="py-[15px] text-[24px]">
				{albums.length}
				{translate(TranslationKeys.LIBRARY_NAV_ALBUMS)}
			</h1>
			<div class="h-full grid grid-cols-2 gap-[15px]">
				{#each albums as album}
					<a href={`/my/library/music/albums/${album.publicId}`}>
						<AlbumTile {album} />
					</a>
				{/each}
			</div>
		</div>
	{:else}
		<Loading />
	{/if}
</main>
