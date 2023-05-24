<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import { page } from "$app/stores";
	import { playNow, playingQueue, songToPlay } from "../../../../../store";
	import MusicTile from "$lib/components/Music/MusicTile.svelte";

	const albumPublicId = $page.params.publicId;

	$: songs = new Array<Music>();

	async function playAlbum() {
		songs = await Requests.makeAuthRequest("GET", `library/album/${albumPublicId}`, null)
			.then((resp) => resp.json())
			.then((resp) => resp["data"]["songs"])
			.catch((err) => {
				console.error(err);
			});
	}
	onMount(() => {
		playAlbum();
	});
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_LIBRARY)}</title>
</svelte:head>

<main>
	{#if songs}
		<div class="h-[85vh] overflow-y-scroll">
			{#each songs as song}
				<button
					class="block w-full"
					on:click={() => {
						playNow.set(true);
						songToPlay.set(song);
						playingQueue.set(songs);
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
