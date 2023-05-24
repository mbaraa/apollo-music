<script lang="ts">
	import type { Year } from "$lib/entities";
	import Marquee from "$lib/ui/Marquee.svelte";

	export let year: Year;
</script>

{#if year}
	<div class="border-[1px]">
		<Marquee
			title={year.name === "0" ? "Unknown Year" : year.name}
			_class="text-[20px] w-[220px] sm:w-[300px]"
		/>
	</div>
{/if}
