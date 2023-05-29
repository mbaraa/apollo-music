<script lang="ts">
	import Navbar from "$lib/components/Navbar/index.svelte";
	import Player from "$lib/components/Player/index.svelte";

	let innerHeight = 0;
	let othersHeight = 0;
	$: slotHeight = innerHeight - othersHeight /*height of the player and navbar*/;
</script>

<div bind:offsetHeight={innerHeight} class="absolute z-[-1] h-[100vh]" />
<div style="height: {slotHeight}px;" class="overflow-y-scroll bg-dark-primary">
	<slot />
</div>
<div bind:offsetHeight={othersHeight}>
	<Player />
	<Navbar />
</div>
