<script lang="ts">
	import { page } from "$app/stores";

	export let title: string;
	export let href = "#";
	export let _class: string = "";

	let isActive = $page.url.pathname.includes(href);
	const classes =
		"hover:opacity-[0.9] grid grid-cols-1 place-items-center font-IBMPlexSans font-[400] text-dark-accent ";
</script>

<a {href} class:active={isActive} class={classes + _class}>
	<slot />
	{title}
</a>
