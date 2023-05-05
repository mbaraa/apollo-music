<script lang="ts">
	export let tracks: string[] = [];

	let player: HTMLAudioElement;
	export let currentTrack = "";

	function handleLoad(event: Event) {
		if ((event.target as HTMLAudioElement).readyState >= 2) {
			player.load();
			player.play();
		}
	}

	function handledEnded() {
		next();
	}

	function play() {
		player.play();
	}
	function pause() {
		player.pause();
	}
	function stop() {}
	function next() {}
	function previous() {}
</script>

<audio
	bind:this={player}
	hidden
	controls
	src={currentTrack}
	on:loadeddata={handleLoad}
	on:ended={handledEnded}
/>
{#if player}
	<input type="range" bind:value={player.currentTime} />
	<button class="p-[10px] border-[1px] border-black rounded-[10px]" on:click={play}>Play</button>
	<button class="p-[10px] border-[1px] border-black rounded-[10px]" on:click={pause}>Pause</button>
{/if}
