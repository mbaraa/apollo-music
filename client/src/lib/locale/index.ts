import config from "$lib/config";
import en from "$lib/strings/en";
import { TranslationKeys, type TranslationValues } from "$lib/strings/keys";

const locales: TranslationValues[] = [en];

const fallbackLocale = locales[0];

let activeLocale =
	locales.find((locale) => locale[TranslationKeys.SYMBOL] === config.defaultLocale) ??
	fallbackLocale;

export function translate(key: TranslationKeys): string {
	let text = activeLocale[key];
	if (!text) {
		return "";
	}
	return text;
}

export function setLocale(localeSymbol: string) {
	activeLocale =
		locales.find((locale) => locale[TranslationKeys.SYMBOL] === localeSymbol) ?? fallbackLocale;
}

export function getLocales(): unknown {
	return locales.map((locale) => {
		return { symbol: locale[TranslationKeys.SYMBOL], name: locale[TranslationKeys.NAME] };
	});
}
