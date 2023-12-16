<script lang="ts">
	import config from "$lib/config";
	import type { Music } from "$lib/entities";
	import { translate } from "$lib/locale";
	import { TranslationKeys } from "$lib/strings/keys";
	import Next from "./Next.svelte";
	import Pause from "./Pause.svelte";
	import Play from "./Play.svelte";
	import Previous from "./Previous.svelte";
	import { playingQueue, songToPlay, currentAlbumCover, playNow, shuffleSongs } from "$lib/store";
	import Marquee from "$lib/ui/Marquee.svelte";
	import ArrowDown from "./ArrowDown.svelte";
	import ArrowUp from "./ArrowUp.svelte";
	import Menu from "./Menu.svelte";
	import Shuffle from "./Shuffle.svelte";
	import Repeat from "./Repeat.svelte";
	import { onMount } from "svelte";

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
	$: canPlay = playlist !== null && playlist.length > 0 && currentTime !== undefined;

	let player: HTMLAudioElement;
	let currentTime = 0;
	let duration = 0;
	let currentAudio: Music = { title: "Play Queue is Empty" } as Music;
	let paused = true;

	let pageTitle = translate(TranslationKeys.TITLE_LIBRARY);
	let expand = false;
	let height = "75px";

	function toggleExpand() {
		if (!canPlay) return;
		expand = !expand;
		if (expand) {
			playNow.set(false);
			height = "85vh";
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
		paused = false;

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
		paused = false;
		fetchMusic(playlist[prevSongIndex]);
	}

	function togglePlayPause() {
		paused = !paused;
		if (player.paused) player.play();
		else player.pause();
	}

	function handleFinishedSong(event: Event) {
		if (playlist.findIndex((m) => m.publicId === currentAudio.publicId) === playlist.length - 1) {
			console.log("finished");
			return;
		}
		next();
		paused = false;
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
		return (m >= 10 ? "" : "0") + m.toString();
	}

	function formatTime(time: number): string {
		const ss = Math.floor(time % 60);
		const mm = Math.floor((time / 60) % 60);
		const hh = Math.floor((time / 60 / 60) % 60);

		return `${hh > 0 ? `${formatNumber(hh)}:` : ""}${formatNumber(mm)}:${formatNumber(ss)}`;
	}

	function setMediaSession() {
		if ("mediaSession" in navigator) {
			navigator.mediaSession.metadata = new MediaMetadata({
				title: currentAudio.title,
				artist: currentAudio.artistName,
				album: currentAudio.albumTitle,
				artwork: [
					{
						src: cover,
						sizes: "96x96",
						type: "image/png"
					}
				]
			});

			navigator.mediaSession.setActionHandler("play", () => {
				player?.play();
			});
			navigator.mediaSession.setActionHandler("pause", () => {
				player?.pause();
			});
			navigator.mediaSession.setActionHandler("stop", () => {
				player?.pause();
				player.currentTime = 0;
			});
			navigator.mediaSession.setActionHandler("seekbackward", () => {
				let seekTo = -10;
				if (currentTime + seekTo < 0) {
					seekTo = 0;
				}
				player.currentTime += seekTo;
			});
			navigator.mediaSession.setActionHandler("seekforward", () => {
				let seekTo = +10;
				if (currentTime + seekTo > duration) {
					seekTo = 0;
				}
				player.currentTime += seekTo;
			});
			navigator.mediaSession.setActionHandler("seekto", (a: MediaSessionActionDetails) => {
				const seekTime = Number(a.seekOffset);
				player.currentTime = seekTime;
			});
			navigator.mediaSession.setActionHandler("previoustrack", () => {
				previous();
			});
			navigator.mediaSession.setActionHandler("nexttrack", () => {
				next();
			});
		}
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
		paused = false;
		player.load();
		player.play();
		currentAudio = music;
		pageTitle = currentAudio.title + translate(TranslationKeys.TITLE_PLAYING_SUFFIX);
		setMediaSession();
	}

	onMount(() => {
		window.addEventListener("popstate", (e) => {
			if (expand && canPlay) {
				e.stopImmediatePropagation();
				toggleExpand();
				return;
			}
			history.go(-1);
		});
	});
</script>

<svelte:head>
	<title>{pageTitle}</title>
</svelte:head>

<div class="hidden">
	{#if canPlay && playOnAdd}
		{toggleExpand()}
		{fetchMusic(selectedSong)}
	{/if}
</div>

<div
	class="text-dark-secondary w-[100vw] bg-dark-neutral"
	style="height: {height}; position: {expand ? 'absolute' : 'inherit'}; bottom: {expand
		? '60px'
		: '0'}; display: {canPlay ? 'block' : 'none'}"
>
	{#if expand && canPlay}
		<div class="h-full w-full">
			<div class="flex justify-between items-center px-10 p-5">
				<button class="text-dark-accent p-1.5" on:click={toggleExpand}><ArrowDown /></button>
				<h1 class="font-IBMPlexSans text-xl font-medium">
					{translate(TranslationKeys.PLAYER_NOW_PLAYING)}
				</h1>
				<button class="text-dark-accent p-1.5" on:click={toggleExpand}>
					<Menu />
				</button>
			</div>
			<div class="flex w-full items-center justify-center flex-col mt-[20px]">
				<div class="w-[80%] h-[80%]">
					<img src={cover} class="w-full h-full rounded-[10px]" alt="Album Cover" />
				</div>
			</div>
			<div class="px-[10%] pt-[16px]">
				<div>
					<div class="flex justify-between items-center font-IBMPlexSans">
						<div>
							<h2 class="text-[25px] text-dark-secondary">{currentAudio.title}</h2>
							<h3 class="text-[16px] text-dark-secondary opacity-[0.7] pt-[1px]">
								{currentAudio.artistName}
							</h3>
						</div>
						<div>like</div>
					</div>
					<input
						type="range"
						on:change={handleSeeking}
						max={duration}
						value={currentTime}
						class="pt-[35px] w-full h-[5px]"
					/>
					<div class="flex justify-between">
						<span class="text-dark-secondary text-[15px] opacity-[0.7]"
							>{formatTime(currentTime)}</span
						>
						<span class="text-dark-secondary text-[15px] opacity-[0.7]">{formatTime(duration)}</span
						>
					</div>
				</div>
				<div class="pt-[20px] flex justify-between items-center">
					<button><Shuffle /></button>
					<div class="text-dark-secondary flex justify-between space-x-[20px]">
						<button on:click={previous}>
							<Previous />
						</button>
						<button
							class="w-[64px] h-[64px] rounded-[100%] bg-dark-primary m-auto shadow-inner grid grid-cols-1 place-items-center"
							on:click={togglePlayPause}
						>
							{#if paused}<Play /> {:else} <Pause /> {/if}
						</button>
						<button on:click={next}>
							<Next />
						</button>
					</div>
					<button><Repeat /></button>
				</div>
				<div>queue</div>
			</div>
		</div>
	{:else}
		<div class="" on:click={toggleExpand} on:keydown={() => {}}>
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
						on:click={(e) => {
							e.stopPropagation();
							previous();
						}}
					>
						<Previous />
					</button>
					<button
						class="p-[5px]"
						on:click={(e) => {
							e.stopPropagation();
							togglePlayPause();
						}}
					>
						{#if paused}<Play /> {:else} <Pause /> {/if}
					</button>
					<button
						class="p-[5px]"
						on:click={(e) => {
							e.stopPropagation();
							next();
						}}><Next /></button
					>
				</div>
			</div>

			<progress max={duration} value={currentTime} class="w-full h-[5px] absolute bottom-[0] p-0" />
		</div>
	{/if}
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

<style lang="postcss" scoped>
	.container {
		@apply bg-[#ff9922];
	}

	progress::-moz-progress-bar {
		background: #9cc7ea;
	}

	progress::-webkit-progress-value {
		background: #9cc7ea;
	}
</style>
