<script lang="ts">
	import { page } from "$app/stores";
	import RouteLink from "$lib/components/Navbar/RouteLink.svelte";

	$: path = $page.url.pathname;
	const libraryEntries = [
		{ name: "All Songs", icon: null, url: "/my/library/music/all-songs" },
		{ name: "Albums", icon: null, url: "/my/library/music/albums" },
		{ name: "Artists", icon: null, url: "/my/library/music/artists" },
		{ name: "Favorites", icon: null, url: "/my/library/favorites" },
		{ name: "Playlists", icon: null, url: "/my/library/playlists" },
		{ name: "Genres", icon: null, url: "/my/library/music/genres" },
		{ name: "Years", icon: null, url: "/my/library/music/years" },
		//{ name: "Playing Queue", icon: null, url: "/my/library/queue" },
		//{ name: "Recently Added", icon: null, url: "/my/library/recently-added" },
		{ name: "Upload", icon: null, url: "/my/library/upload" }
	];
</script>

<div class="bg-dark-primary">
	<h1 class="font-Comfortaa text-dark-accent text-[30px] p-[25px]">Apollo Library</h1>

	<nav>
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
