<script lang="ts">
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import type { Genre } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";
	import Loading from "$lib/ui/Loading.svelte";
	import GenreTile from "$lib/components/Music/GenreTile.svelte";

	let genres: Genre[];

	onMount(async () => {
		genres = await Requests.makeAuthRequest("GET", "library/genres", null)
			.then((resp) => resp.json())
			.then((resp) => resp["data"])
			.catch((err) => {
				console.error(err);
			});
	});
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_LIBRARY_MUSIC_GENRERS)}</title>
</svelte:head>

<main>
	{#if genres}
		<div class="h-full">
			{#each genres as genre}
				<a href={`/my/library/music/genres/${genre.publicId}`}>
					<GenreTile {genre} />
				</a>
			{/each}
		</div>
	{:else}
		<Loading />
	{/if}
</main>
