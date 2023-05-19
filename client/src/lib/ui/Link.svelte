<script lang="ts">
	export let href: string;
	export let title: string;
	export let _class = "";

	const classes = "underline hover:opacity-[0.9] font-IBMPlexSans";
</script>

<a {href} class={`${classes} ${_class}`}>{title}</a>
