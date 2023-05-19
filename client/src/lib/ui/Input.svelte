<script lang="ts">
	export let _class = "";
	export let type: "text" | "password" | "email" | "number" = "text";
	export let placeholder: string;
	export let required = false;
	export let value: any;

	const classes =
		"block font-IBMPlexSans bg-dark-primary text-dark-secondary border-[2px] border-dark-accent rounded-[20px] p-[16px] h-[56px] w-[330px]";
</script>

{#if type === "text"}
	<input type="text" bind:value {placeholder} class={`${classes} ${_class}`} {required} />
{:else if type === "email"}
	<input type="email" bind:value {placeholder} class={`${classes} ${_class}`} {required} />
{:else if type === "password"}
	<input type="password" bind:value {placeholder} class={`${classes} ${_class}`} {required} />
{/if}
