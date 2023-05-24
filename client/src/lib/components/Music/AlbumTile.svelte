<script lang="ts">
	import type { Album } from "$lib/entities";
	import Marquee from "$lib/ui/Marquee.svelte";

	export let album: Album;
</script>

{#if album}
	<div class="border-[1px]">
		<img
			src={`data:image/*;base64,${album.coverB64}`}
			alt="cover"
			class="inline-block w-[75px] h-[75px] rounded-[0px] mr-[5px]"
		/>
		<Marquee title={album.title} _class="text-[20px] w-[220px] sm:w-[300px]" />
	</div>
{/if}
