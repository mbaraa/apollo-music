<script lang="ts">
	import Requests from "$lib/utils/requests/Requests";
	import { onMount } from "svelte";
	import config from "$lib/config";
	import Loading from "$lib/ui/Loading.svelte";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";

	let player: HTMLAudioElement;

	let musicList: string[] = [];
	let displayMusicList: string[] = [];
	let currentMusic = "";
	let currentTime = 0;
	let duration = 0;

	function search(event: Event) {
		const query = (event.target as HTMLInputElement).value;
		displayMusicList = musicList.filter((name) => {
			return name.toLowerCase().includes(query);
		});
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
		let nextSongIndex = (musicList.findIndex((m) => m === currentMusic) + 1) % musicList.length;
		while (!isPlayable(musicList[nextSongIndex])) {
			nextSongIndex = (nextSongIndex + 1) % musicList.length;
		}
		fetchMusic(musicList[nextSongIndex]);
	}

	function previous() {
		let prevSongIndex = (musicList.findIndex((m) => m === currentMusic) - 1) % musicList.length;
		while (!isPlayable(musicList[prevSongIndex])) {
			prevSongIndex = (prevSongIndex - 1) % musicList.length;
		}
		fetchMusic(musicList[prevSongIndex]);
	}

	function random() {
		fetchMusic(musicList[Math.floor(Math.random() * musicList.length)]);
	}

	function handleFinishedSong(event: Event) {
		next();
		console.log("finished");
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

	async function fetchMusic(music: string) {
		await fetch(`${config["backendAddress"]}/music-static/${music}`, {
			method: "GET",
			mode: "cors"
		})
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
		currentMusic = music;
	}

	onMount(async () => {
		musicList = await Requests.makeRequest("GET", "ls-music", null)
			.then((resp) => resp.json())
			.then((music) => music);

		displayMusicList = musicList;
	});
</script>

<svelte:head>
	<title>{translate(TranslationKeys.TITLE_LIBRARY)}</title>
</svelte:head>

{#if musicList.length > 0}
	<div class="bg-[#121212] h-[100vh] text-white">
		<div class="h-[400px] overflow-scroll">
			{#each displayMusicList as music}
				<button
					class="block w-[100%] p-[10px] border-[1px] border-white rounded-[10px]"
					on:click={() => {
						fetchMusic(music);
					}}>{music}</button
				>
			{/each}
		</div>
		<input
			class="p-[10px] w-full border-[1px] border-white bg-[#121212] rounded-[10px]"
			placeholder="find a song"
			on:keyup={search}
		/>
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
		/>
		{#if player}
			<br />
			<span class="text-[20px] font-bold">{currentMusic}</span>
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
			<button class="p-[10px] border-[1px] border-white rounded-[10px]" on:click={next}>Next</button
			>
			<button class="p-[10px] border-[1px] border-white rounded-[10px]" on:click={random}
				>Random</button
			>
		{/if}
	</div>
{:else}
	<Loading />
{/if}
