<script lang="ts">
	import type { Album } from "$lib/entities";

	export let album: Album;
</script>

{#if album}
	<div class="border-[1px]">
		<span>{album.title}</span>
		<br />
		<span>{album.artistName}</span>
		<br />
		<span>{album.year}</span>
	</div>
{/if}
