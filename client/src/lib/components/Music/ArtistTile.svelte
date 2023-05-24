<script lang="ts">
	import type { Artist } from "$lib/entities";

	export let artist: Artist;
</script>

{#if artist}
	<div class="p-[5px] border-[1px]">
		<span>{artist.name}</span>
	</div>
{/if}
