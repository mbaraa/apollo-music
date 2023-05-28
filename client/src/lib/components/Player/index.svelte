<script lang="ts">
	import config from "$lib/config";
	import type { Music } from "$lib/entities";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import { onMount } from "svelte";
	import Next from "./Next.svelte";
	import Pause from "./Pause.svelte";
	import Play from "./Play.svelte";
	import Previous from "./Previous.svelte";
	import { playingQueue, songToPlay, currentAlbumCover, playNow, shuffleSongs } from "$lib/store";
	import Marquee from "$lib/ui/Marquee.svelte";

	let playlist: Music[];
	let selectedSong: Music = {} as Music;
	let cover = "/favicon.ico";
	let playOnAdd = true;
	let shuffle = false;

	playingQueue.subscribe((queue) => {
		playlist = queue;
	});
	songToPlay.subscribe((song) => {
		selectedSong = song;
	});
	currentAlbumCover.subscribe((_cover) => {
		cover = _cover;
	});
	playNow.subscribe((play) => {
		playOnAdd = play;
	});
	shuffleSongs.subscribe((_shuffle) => {
		shuffle = _shuffle;
	});

	/************************************/
	$: canPlay = playlist !== null && playlist.length > 0;

	let player: HTMLAudioElement;
	let currentTime = 0;
	let duration = 0;
	let currentAudio: Music = { title: "Play Queue is Empty" } as Music;

	let pageTitle = translate(TranslationKeys.TITLE_LIBRARY);
	let expand = false;
	let height = "75px";

	function toggleExpand() {
		expand = !expand;
		if (expand) {
			height = "100vh";
		} else {
			height = "75px";
		}
	}

	function handleSeeking(event: Event) {
		const s = event.target as HTMLInputElement;
		const seekTime = Number(s.value);
		//const audio = document.getElementById("aud") as any;
		player.currentTime = seekTime;
		//player = audio;
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
		if (shuffle) {
			random();
			return;
		}
		let nextSongIndex =
			(playlist.findIndex((m) => m.publicId === currentAudio.publicId) + 1) % playlist.length;
		//while (!isPlayable(playlist[nextSongIndex].)) {
		//	nextSongIndex = (nextSongIndex + 1) % playlist.length;
		//}
		fetchMusic(playlist[nextSongIndex]);
	}

	function previous() {
		if (shuffle) {
			random();
			return;
		}
		let prevSongIndex =
			(playlist.findIndex((m) => m.publicId === currentAudio.publicId) - 1) % playlist.length;
		fetchMusic(playlist[prevSongIndex]);
	}

	function handleFinishedSong(event: Event) {
		if (playlist.findIndex((m) => m.publicId === currentAudio.publicId) === playlist.length - 1) {
			console.log("finished");
			return;
		}
		next();
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
		if (player) {
			player.pause();
			player.currentTime = 0;
			currentTime = 0;
		}
		player.src = `${config["backendAddress"]}/storage/${music.audio.publicPath}?token=${
			localStorage.getItem("token") ?? ""
		}`;
		player.load();
		player.play();
		currentAudio = music;
		pageTitle = currentAudio.title + translate(TranslationKeys.TITLE_PLAYING_SUFFIX);
	}
</script>

<svelte:head>
	<title>{pageTitle}</title>
</svelte:head>

<div class="hidden">
	{#if canPlay && playOnAdd}
		{fetchMusic(selectedSong)}
	{/if}
</div>

<div
	class="text-dark-secondary w-[100vw] bg-dark-neutral"
	on:click={toggleExpand}
	on:keydown={() => {}}
	style="height: {height}; position: {expand ? 'absolute' : 'inherit'}; bottom: 0;"
>
	{#if expand}
		<div class="h-[80vh] overflow-y-scroll">
			{#each playlist as music}
				<div
					class="border-[1px] h-[40px]"
					on:click={() => {
						toggleExpand();
						fetchMusic(music);
					}}
				>
					{music.title}
				</div>
			{/each}
		</div>
	{/if}
	<div class="">
		<div class="p-[10px] px-[25px] flex justify-between items-center">
			<div
				class="float-left text-dark-secondary font-IBMPlexSans pb-[10px] flex items-center"
				on:keydown={() => {}}
			>
				<img src={cover} class="rounded-[5px] w-[52px] h-[52px] inline" alt="Album Cover" />
				<div class="pl-[10px] text-[18px] font-bold w-[150px] inline-block">
					<Marquee title={currentAudio.title} _class="w-[100%] sm:w-[300px]" />
				</div>
			</div>
			<div class="text-dark-secondary mb-[10px]">
				<button
					class="p-[5px]"
					on:click={() => {
						toggleExpand();
						previous();
					}}
				>
					<Previous />
				</button>
				<button
					class="p-[5px]"
					on:click={() => {
						toggleExpand();
						if (player.paused) player.play();
						else player.pause();
					}}
					>{#if player?.paused}<Play /> {:else} <Pause /> {/if}</button
				>
				<button
					class="p-[5px]"
					on:click={() => {
						toggleExpand();
						next();
					}}><Next /></button
				>
			</div>
		</div>

		<progress max={duration} value={currentTime} class="w-full h-[5px] absolute bottom-[0] p-0" />
	</div>
</div>

<audio
	id="aud"
	class="hidden"
	bind:this={player}
	controls
	preload="none"
	on:loadeddata={handleLoad}
	on:timeupdate={handleUpdateTime}
	on:ended={handleFinishedSong}
	on:progress={() => {
		console.log("downloading...");
	}}
/>

<style>
	progress::-moz-progress-bar {
		background: #9cc7ea;
	}
	progress::-webkit-progress-value {
		background: #9cc7ea;
	}
</style>
