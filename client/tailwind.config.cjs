const defaultTheme = require("tailwindcss/defaultTheme");
const plugin = require("tailwindcss/plugin");

/** @type {import('tailwindcss').Config} */
export default {
	content: ["./src/**/*.{html,js,svelte,ts}"],
	theme: {
		container: {
			center: true
		},
		colors: {
			dark: {
				primary: "#051220",
				secondary: "#fff",
				accent: "#add8fb",
				accent2: "#add8fb99",
				neutral: "#273d50"
			},
			light: {},
			...defaultTheme.colors
		},
		extend: {
			fontFamily: {
				Comfortaa: ["Comfortaa", "sans"],
				IBMPlexSans: ["IBM Plex Sans", "sans"]
			}
		}
	},
	plugins: [
		plugin(function ({ addVariant }) {
			addVariant("current", "&.active");
		})
	]
};
