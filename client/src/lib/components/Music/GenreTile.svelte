<script lang="ts">
	import type { Genre } from "$lib/entities";
	import Marquee from "$lib/ui/Marquee.svelte";

	export let genre: Genre;
</script>

{#if genre}
	<div class="border-[1px]">
		<Marquee title={genre.name} _class="text-[20px] w-[220px] sm:w-[300px]" />
	</div>
{/if}
