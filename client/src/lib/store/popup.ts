import { writable } from "svelte/store";

export const popupMessage = writable("");
export const popupType = writable<"info" | "error">("info");
export const showPopup = writable(false);
