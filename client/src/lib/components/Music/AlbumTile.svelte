<script lang="ts">
	import type { Album } from "$lib/entities";

	export let album: Album;
</script>

{#if album}
	<div class="">
		<img
			src={`data:image/*;base64,${album.coverB64}`}
			alt="cover"
			class="w-[150px] h-[150px] rounded-[0px] mr-[5px] rounded-[10px]"
		/>
		<span class="text-[18px] block">
			{album.title}
		</span>
		<span class="text-[14px] block font-[300]">{album.artistName}</span>
	</div>
{/if}
