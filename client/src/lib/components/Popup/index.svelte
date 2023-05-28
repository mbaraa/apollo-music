<script lang="ts">
	import { popupMessage, popupType, showPopup } from "$lib/store";
	import HappyCloud from "./HappyCloud.svelte";
	import SadCloud from "./SadCloud.svelte";

	let message: string;
	let type: "info" | "error";
	let show = false;

	popupMessage.subscribe((m) => {
		message = m;
	});
	popupType.subscribe((t) => {
		type = t;
	});
	showPopup.subscribe((s) => {
		show = s;
	});
</script>

<div
	class="absolute bottom-[30px] rounded-[10px] min-h-[60px] h-auto w-[90vw] mx-[5%] bg-dark-accent2 text-dark-primary p-[10px] fade-in font-IBMPlexSans"
	style="display: {show ? 'block' : 'none'}; border-left: {type === 'error'
		? '#ff0000'
		: '#388e3c'} solid 15px;"
>
	{#if type === "error"}
		<SadCloud />
	{:else}
		<HappyCloud />
	{/if}
	<span class="px-[10px]">
		{message}
	</span>
</div>

<style>
	.fade-in {
		animation: fadeIn 1s;
	}

	@keyframes fadeIn {
		0% {
			opacity: 0;
		}
		100% {
			opacity: 1;
		}
	}
</style>
