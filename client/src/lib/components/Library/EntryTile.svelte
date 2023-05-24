<script lang="ts">
	export let name: string;
	export let icon = "";
	export let url: string;
</script>

<a href={url}>
	<div class="p-[20px] font-Comfortaa">
		<img src="/favicon.ico" alt="icon" class="inline-block" />
		<h1 class="inline-block pl-[10px] text-[20px]">{name}</h1>
	</div>
</a>
