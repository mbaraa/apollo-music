<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Year, Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import Player from "$lib/components/Player/index.svelte";
	import { page } from "$app/stores";
	import { playNow, playingQueue, songToPlay } from "../../../../../../store";
	import MusicTile from "$lib/components/Music/MusicTile.svelte";

	const yearPublicId = $page.params.publicId;

	let year: Year;

	async function playYear() {
		year = await Requests.makeAuthRequest("GET", `library/year/${yearPublicId}`, null)
			.then((resp) => resp.json())
			.then((resp) => resp["data"])
			.catch((err) => {
				console.error(err);
			});
	}
	onMount(() => {
		playYear();
	});
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_LIBRARY)}</title>
</svelte:head>

<main>
	{#if year && year.songs}
		<div class="h-[85vh] overflow-y-scroll">
			{#each year.songs as song}
				<button
					class="block w-full"
					on:click={() => {
						playNow.set(true);
						songToPlay.set(song);
						playingQueue.set(year.songs);
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
