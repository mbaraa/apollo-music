<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Album, Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import AlbumTile from "$lib/components/Music/AlbumTile.svelte";

	let albums: Promise<Album[]>;

	async function fetchAlbums(fetchCovers: boolean) {
		return (await Requests.makeAuthRequest("GET", "library/albums", null, {
			"fetch-covers": fetchCovers
		})
			.then((resp) => resp.json())
			.then((resp) => resp["data"] as Album[])
			.catch((err) => {
				console.error(err);
				return [];
			})) as Album[];
	}

	onMount(async () => {
		albums = fetchAlbums(false);
		let c = 0;
		let localCoversAreGood = (await albums).every((a) => {
			const cover = localStorage.getItem("cover_" + a.publicId);
			console.log(cover);
			if (cover) {
				c++;
				return true;
			}
			return false;
		});
		console.log("good", localCoversAreGood, c);

		if (!localCoversAreGood) {
			albums = fetchAlbums(true);
			(await albums).flat().forEach((a) => {
				localStorage.setItem("cover_" + a.publicId, a.coverB64);
			});
		}
	});
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_LIBRARY_MUSIC_ALBUMS)}</title>
</svelte:head>

<main>
	{#await albums}
		<Loading />
	{:then albums}
		{#if albums}
			<div class="px-[25px] font-IBMPlexSans text-dark-secondary">
				<h1 class="py-[15px] text-[24px]">
					{albums.length}
					{translate(TranslationKeys.LIBRARY_NAV_ALBUMS)}
				</h1>
				<div class="h-full grid grid-cols-2 gap-x-[25px] gap-y-[32px]">
					{#each albums as album}
						<a href={`/my/library/music/albums/${album.publicId}`}>
							<AlbumTile {album} />
						</a>
					{/each}
				</div>
			</div>
		{/if}
	{:catch}
		oops
	{/await}
</main>
