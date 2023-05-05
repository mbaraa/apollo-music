export enum TranslationKeys {
	SYMBOL = "SYMBOL",
	NAME = "NAME",
	UI_DIRECTION = "UI_DIRECTION",

	TITLE_ABOUT = "TITLE_ABOUT",
	TITLE_RECENTS = "TITLE_RECENTS",
	TITLE_LIBRARY = "TITLE_LIBRARY",
	TITLE_EQUALIZER = "TITLE_EQUALIZER",
	TITLE_SETTINGS = "TITLE_SETTINGS",
	TITLE_PLAYING_SUFFIX = "TITLE_PLAYING_SUFFIX"
}

export type TranslationValues = { [key in TranslationKeys]: string };
