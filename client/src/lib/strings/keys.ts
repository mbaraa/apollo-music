export enum TranslationKeys {
	SYMBOL = "SYMBOL",
	NAME = "NAME",
	UI_DIRECTION = "UI_DIRECTION",

	TITLE_MAIN = "TITLE_MAIN",
	TITLE_ABOUT = "TITLE_ABOUT",
	TITLE_RECENTS = "TITLE_RECENTS",
	TITLE_LIBRARY = "TITLE_LIBRARY",
	TITLE_LIBRARY_MUSIC = "TITLE_LIBRARY_MUSIC",
	TITLE_LIBRARY_MUSIC_SONGS = "TITLE_LIBRARY_MUSIC_SONGS",
	TITLE_LIBRARY_MUSIC_ALBUMS = "TITLE_LIBRARY_MUSIC_ALBUMS",
	TITLE_LIBRARY_MUSIC_ARTISTS = "TITLE_LIBRARY_MUSIC_ARTISTS",
	TITLE_LIBRARY_MUSIC_YEARS = "TITLE_LIBRARY_MUSIC_YEARS",
	TITLE_LIBRARY_MUSIC_GENRERS = "TITLE_LIBRARY_MUSIC_GENRES",

	TITLE_EQUALIZER = "TITLE_EQUALIZER",
	TITLE_SETTINGS = "TITLE_SETTINGS",
	TITLE_PLAYING_SUFFIX = "TITLE_PLAYING_SUFFIX",
	TITLE_SIGN_IN = "TITLE_SIGN_IN",
	TITLE_RESET_PASSWORD = "TITLE_RESET_PASSWORD",
	TITLE_FORGOT_PASSWORD = "TITLE_FORGOT_PASSWORD",

	SIGN_IN_HEADER = "SIGN_IN_HEADER",
	SIGN_IN_EMAIL = "SIGN_IN_EMAIL",
	SIGN_IN_PASSWORD = "SIGN_IN_PASSWORD",
	SIGN_IN_BUTTON = "SIGN_IN_BUTTON",
	SIGN_IN_FORGOT_PASSWORD = "SIGN_IN_FORGOT_PASSWORD",
	SIGN_IN_SIGN_UP = "SIGN_IN_SIGN_UP",
	SIGN_IN_SIGN_UP_HERE = "SIGN_IN_SIGN_UP_HERE",

	FORGOT_PASSWORD_EMAIL = "FORGOT_PASSWORD_EMAIL",
	FORGOT_PASSWORD_BUTTON = "FORGOT_PASSWORD_BUTTON",
	PASSWORD_RESET_NEW_PASSWORD = "PASSWORD_RESET_NEW_PASSWORD",
	PASSWORD_RESET_UPDATE_BUTTON = "PASSWORD_RESET_UPDATE_BUTTON",

	LIBRARY_HEAD = "LIBRARY_HEAD",
	LIBRARY_NAV_ALBUMS = "LIBRARY_NAV_ALBUMS",
	LIBRARY_NAV_ARTISTS = "LIBRARY_NAV_ARTISTS",
	LIBRARY_NAV_ALL_SONGS = "LIBRARY_NAV_ALL_SONGS",
	LIBRARY_NAV_FAVORITES = "LIBRARY_NAV_FAVORITES",
	LIBRARY_NAV_PLAYLISTS = "LIBRARY_NAV_PLAYLISTS",
	LIBRARY_NAV_YEARS = "LIBRARY_NAV_YEARS",
	LIBRARY_NAV_GENRES = "LIBRARY_NAV_GENRES",
	LIBRARY_NAV_UPLOAD = "LIBRARY_NAV_UPLOAD",

	PLAYER_NOW_PLAYING = "PLAYER_NOW_PLAYING"
}

export type TranslationValues = {
	[key in TranslationKeys]: string;
};
