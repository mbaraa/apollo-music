import { writable } from "svelte/store";
import type { Music } from "$lib/entities";

export const playingQueue = writable(new Array<Music>());
export const songToPlay = writable({} as Music);
export const currentAlbumCover = writable("/favicon.ico");
export const playNow = writable(false);
export const shuffleSongs = writable(false);
