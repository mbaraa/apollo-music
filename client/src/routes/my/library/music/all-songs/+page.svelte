<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import { page } from "$app/stores";
	import { playNow, playingQueue, songToPlay } from "$lib/store";
	import MusicTile from "$lib/components/Music/MusicTile.svelte";

	$: songs = new Array<Music>();

	async function playAlbum() {
		songs = await Requests.makeAuthRequest("GET", `library/music`, null)
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
	<title>{translate(TranslationKeys.TITLE_LIBRARY_MUSIC_SONGS)}</title>
</svelte:head>

<main>
	{#if songs}
		<div class="h-full">
			{#each songs as song}
				<button
					class="block w-full"
					on:click={() => {
						songToPlay.set(song);
						playingQueue.set(songs);
						playNow.set(true);
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
