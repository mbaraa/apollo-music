<script lang="ts">
	export let title: string;
	export let _class = "";
</script>

<p class={`marquee inline-block ${_class}`}>
	<span>{title}</span>
</p>

<style>
	@media (max-width: 668px) {
		.marquee {
			margin: 0 auto;
			white-space: nowrap;
			overflow: hidden;
		}

		.marquee span {
			display: inline-block;
			padding-left: 100%;
			animation: marquee 15s linear infinite;
		}
	}

	@keyframes marquee {
		0% {
			transform: translate(0, 0);
		}
		100% {
			transform: translate(-100%, 0);
		}
	}
</style>
