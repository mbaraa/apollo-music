<script lang="ts">
	import { goto } from "$app/navigation";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import EntryTile from "$lib/components/Library/EntryTile.svelte";

	const libraryEntries = [
		{ name: "Favorites", icon: null, url: "/my/library/favorites" },
		{ name: "Playlists", icon: null, url: "/my/library/playlists" },
		{ name: "All Songs", icon: null, url: "/my/library/music/all-songs" },
		{ name: "Albums", icon: null, url: "/my/library/music/albums" },
		{ name: "Artists", icon: null, url: "/my/library/music/artists" },
		{ name: "Genres", icon: null, url: "/my/library/music/genres" },
		{ name: "Years", icon: null, url: "/my/library/music/years" },
		//{ name: "Playing Queue", icon: null, url: "/my/library/queue" },
		{ name: "Recently Added", icon: null, url: "/my/library/recently-added" },
		{ name: "Upload", icon: null, url: "/my/library/upload" }
	];
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_LIBRARY)}</title>
</svelte:head>

<main>
	<div class="h-[85vh] overflow-y-scroll">
		{#each libraryEntries as entry}
			<EntryTile name={entry.name} url={entry.url} />
		{/each}
	</div>
</main>
