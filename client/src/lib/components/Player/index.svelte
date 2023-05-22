<script lang="ts">
	import config from "$lib/config";
	import type { Music } from "$lib/entities";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import Requests from "$lib/utils/requests/Requests";
	import { createEventDispatcher } from "svelte";
	import Next from "./Next.svelte";
	import Pause from "./Pause.svelte";
	import Play from "./Play.svelte";
	import Previous from "./Previous.svelte";

	export let playlist: Music[];

	let player: HTMLAudioElement;
	let currentTime = 0;
	let duration = 0;
	let currentAudio: Music = { title: "Play Queue is Empty" } as Music;

	let pageTitle = translate(TranslationKeys.TITLE_LIBRARY);
	let expand = false;
	let height = "15vh";

	function toggleExpand() {
		expand = !expand;
		if (expand) {
			height = "100vh";
		} else {
			height = "15vh";
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

{#if playlist && playlist.length > 0}
	<div
		class="text-dark-secondary absolute bottom-0 w-[100vw] bg-dark-primary"
		on:click={toggleExpand}
		on:keydown={() => {}}
		style="height: {height};"
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
		<div class="p-[10px]">
			<div class="float-left text-dark-secondary font-IBMPlexSans pb-[10px]" on:keydown={() => {}}>
				<img src="/favicon.ico" class="w-[52px] h-[52px] inline" alt="Album Cover" />
				<div class="pl-[10px] text-[18px] font-bold move-on-overflow inline-block">
					<p>
						{currentAudio.title}
					</p>
				</div>
				<!-- <span class="text-[20px]">{formatTime(currentTime)}/{formatTime(duration)}</span> -->
			</div>
			<div class="float-right">
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
			<input
				type="range"
				class="w-full"
				value={currentTime}
				on:change={handleSeeking}
				min={0}
				max={duration}
			/>
		</div>
	</div>
{/if}

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
	.move-on-overflow {
		width: 150px;
		overflow-x: scroll;
		white-space: nowrap;
	}

	.move-on-overflow p {
		animation-name: scroll;
		animation-duration: 5s;
		animation-iteration-count: infinite;
	}

	@keyframes scroll {
		0% {
			transform: translateX(25%);
		}
		100% {
			transform: translateX(-75%);
		}
	}
</style>
