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
			<li>
				<RouteLink href="/my/library" _class="[&.active]:font-[600]" title="Library"
					><LibraryIcon /></RouteLink
				>
			</li>
			<li>
				<RouteLink href="/my/search" _class="[&.active]:font-[600]" title="Search"
					><SearchIcon /></RouteLink
				>
			</li>
			<li>
				<RouteLink href="/my/settings" _class="[&.active]:font-[600]" title="Settings"
					><SettingsIcon /></RouteLink
				>
			</li>
		</ul>
	{/key}
</nav>
