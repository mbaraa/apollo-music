<script lang="ts">
	import RouteLink from "./RouteLink.svelte";
	import { page } from "$app/stores";
	import LibraryIcon from "./LibraryIcon.svelte";
	import SearchIcon from "./SearchIcon.svelte";
	import SettingsIcon from "./SettingsIcon.svelte";

	$: path = $page.url.pathname;
</script>

<nav class="bg-dark-primary h-[60px]">
	{#key path}
		<ul class="m-0 list-none flex justify-between items-center gap-[20px] pt-[8px] px-[30px]">
			<li><RouteLink href="/my/library" title="Library"><LibraryIcon /></RouteLink></li>
			<li><RouteLink href="/my/search" title="Search"><SearchIcon /></RouteLink></li>
			<li><RouteLink href="/my/settings" title="Settings"><SettingsIcon /></RouteLink></li>
		</ul>
	{/key}
</nav>
