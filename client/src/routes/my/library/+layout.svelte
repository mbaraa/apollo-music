<script lang="ts">
	import { page } from "$app/stores";
	import RouteLink from "$lib/components/Navbar/RouteLink.svelte";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";

	$: path = $page.url.pathname;
	const libraryEntries = [
		{
			name: translate(TranslationKeys.LIBRARY_NAV_ALL_SONGS),
			icon: null,
			url: "/my/library/music/all-songs"
		},
		{
			name: translate(TranslationKeys.LIBRARY_NAV_ALBUMS),
			icon: null,
			url: "/my/library/music/albums"
		},
		{
			name: translate(TranslationKeys.LIBRARY_NAV_ARTISTS),
			icon: null,
			url: "/my/library/music/artists"
		},
		{
			name: translate(TranslationKeys.LIBRARY_NAV_FAVORITES),
			icon: null,
			url: "/my/library/favorites"
		},
		{
			name: translate(TranslationKeys.LIBRARY_NAV_PLAYLISTS),
			icon: null,
			url: "/my/library/playlists"
		},
		{
			name: translate(TranslationKeys.LIBRARY_NAV_GENRES),
			icon: null,
			url: "/my/library/music/genres"
		},
		{
			name: translate(TranslationKeys.LIBRARY_NAV_YEARS),
			icon: null,
			url: "/my/library/music/years"
		},
		//{ name: "Playing Queue", icon: null, url: "/my/library/queue" },
		//{ name: "Recently Added", icon: null, url: "/my/library/recently-added" },
		{ name: translate(TranslationKeys.LIBRARY_NAV_UPLOAD), icon: null, url: "/my/library/upload" }
	];
</script>

<div class="bg-dark-primary">
	<h1 class="font-Comfortaa text-dark-accent text-[30px] p-[25px]">
		{translate(TranslationKeys.LIBRARY_HEAD)}
	</h1>

	<nav class="border-b-[1px] border-b-dark-neutral">
		{#key path}
			<ul
				class="m-0 list-none flex justify-between items-center gap-[20px] pb-[20px] px-[25px] w-full overflow-x-scroll"
			>
				{#each libraryEntries as entry}
					<li>
						<RouteLink
							_class="w-[85px] rounded-[10px] border-[2px] p-[5px] border-dark-accent [&.active]:bg-dark-accent [&.active]:text-dark-primary [&.active]:font-[500]"
							href={entry.url}
							title={entry.name}
						/>
					</li>
				{/each}
			</ul>
		{/key}
	</nav>

	<slot />
</div>
