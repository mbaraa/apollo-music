<script lang="ts">
	import config from "$lib/config";
	import type { Music } from "$lib/entities";
	import Requests from "$lib/utils/requests/Requests";

	export let playlist: Music[];

	let player: HTMLAudioElement;
	let currentTime = 0;
	let duration = 0;
	let currentAudio: Music = { title: "No current music" } as Music;

	let expand = false;
	let height = "100px";

	function toggleExpand() {
		expand = !expand;
		if (expand) {
			height = "100vh";
		} else {
			height = "100px";
		}
	}

	function handleSeeking(event: Event) {
		const s = event.target as HTMLInputElement;
		player.currentTime = Number(s.value);
		console.log(s.value);
		player = player;
	}

	function isPlayable(musicName: string): boolean {
		return (
			musicName.substring(musicName.length - 5) === ".flac" ||
			musicName.substring(musicName.length - 4) === ".mp3" ||
			musicName.substring(musicName.length - 4) === ".m4a" ||
			musicName.substring(musicName.length - 4) === ".oog" ||
			musicName.substring(musicName.length - 4) === ".acc" ||
			musicName.substring(musicName.length - 4) === ".wav"
		);
	}

	function next() {
		let nextSongIndex =
			(playlist.findIndex((m) => m.publicId === currentAudio.publicId) + 1) % playlist.length;
		//while (!isPlayable(playlist[nextSongIndex].)) {
		//	nextSongIndex = (nextSongIndex + 1) % playlist.length;
		//}
		fetchMusic(playlist[nextSongIndex]);
	}

	function previous() {
		let prevSongIndex =
			(playlist.findIndex((m) => m.publicId === currentAudio.publicId) - 1) % playlist.length;
		fetchMusic(playlist[prevSongIndex]);
	}

	function handleFinishedSong(event: Event) {
		next();
		console.log("finished");
	}

	function random() {
		fetchMusic(playlist[Math.floor(Math.random() * playlist.length)]);
	}

	function handleUpdateTime(event: Event) {
		const a = event.target as HTMLAudioElement;
		currentTime = Math.floor(a.currentTime);
	}

	function handleLoad(event: Event) {
		const a = event.target as HTMLAudioElement;
		duration = a.duration;
	}

	function formatNumber(m: number): string {
		return (m > 10 ? "" : "0") + m.toString();
	}

	function formatTime(time: number): string {
		const ss = Math.floor(time % 60);
		const mm = Math.floor((time / 60) % 60);
		const hh = Math.floor((time / 60 / 60) % 60);

		return `${hh > 0 ? `${formatNumber(hh)}:` : ""}${formatNumber(mm)}:${formatNumber(ss)}`;
	}

	async function fetchMusic(music: Music) {
		await Requests.makeAuthRequest("GET", `storage/${music.audio.publicPath}`, null)
			.then((resp) => resp.blob())
			.then((audio) => {
				if (player) {
					player.pause();
					player.currentTime = 0;
					currentTime = 0;
				}
				// player = new Audio(URL.createObjectURL(audio));
				player.src = URL.createObjectURL(audio);
				// player.preload = "auto";
				// player.controls = true;
				// player.oncanplaythrough = handleLoad;
				//player.onended = handleFinishedSong;
				player.load();
				player.play();
			});
		currentAudio = music;
	}
</script>

<div class="text-white absolute bottom-0 w-[100vw] bg-black" style="height: {height};">
	{#if expand}
		{#each playlist as music}
			<div
				class="border-[1px] h-[40px]"
				on:click={() => {
					fetchMusic(music);
				}}
			>
				{music.title}
			</div>
		{/each}
		<br />
	{/if}
	<div class="float-left">
		<br />
		<span class="text-[20px] font-bold">{currentAudio.title}</span>
		|
		<span class="text-[20px]">{formatTime(currentTime)}/{formatTime(duration)}</span>
		<br />
		<input
			type="range"
			class="w-full"
			value={currentTime}
			on:change={handleSeeking}
			min={0}
			max={duration}
		/>
		<br />
		<button class="p-[10px] border-[1px] border-white rounded-[10px]" on:click={previous}
			>Prevoius</button
		>
		<button
			class="p-[10px] border-[1px] border-white rounded-[10px]"
			on:click={() => {
				player.play();
			}}>Play</button
		>
		<button
			class="p-[10px] border-[1px] border-white rounded-[10px]"
			on:click={() => {
				player.pause();
			}}>Pause</button
		>
		<button class="p-[10px] border-[1px] border-white rounded-[10px]" on:click={next}>Next</button>
		<button class="p-[10px] border-[1px] border-white rounded-[10px]" on:click={random}
			>Random</button
		>
	</div>
	<button class="float-right" on:click={toggleExpand}>expand</button>
</div>

<audio
	class="hidden"
	bind:this={player}
	controls
	preload="auto"
	on:loadeddata={handleLoad}
	on:timeupdate={handleUpdateTime}
	on:ended={handleFinishedSong}
	on:progress={() => {
		console.log("downloading...");
	}}
>
	meow
</audio>
