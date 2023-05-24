<script lang="ts">
	import type { Music } from "$lib/entities";

	export let music: Music;
</script>

{#if music}
	<div class="block border-[1px] p-[10px] w-full">
		<span>{music.title}</span>
	</div>
{/if}
