<script lang="ts">
	import type { Album } from "$lib/entities";

	export let album: Album;
</script>

{#if album}
	<div class="w-full h-2/4 space-y-[8px] text-dark-secondary">
		<img src={`data:image/*;base64,${album.coverB64}`} alt="cover" class="rounded-[6px] w-full" />
		<div class="space-y-[2px]">
			<span class="text-[18px] block">
				{album.title}
			</span>
			<span class="text-[14px] block font-[300]">{album.artistName}</span>
		</div>
	</div>
{/if}
