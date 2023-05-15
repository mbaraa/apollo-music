<script lang="ts">
	import config from "$lib/config";
	import Requests from "$lib/utils/requests/Requests";

	let file: HTMLInputElement;
	let files: HTMLInputElement;

	function handleChange(e: Event) {
		const inEv = e.target as HTMLInputElement;
		if (inEv.files) {
			console.log(inEv.files[0].name);
		}
	}
	async function upload() {
		const formData = new FormData();
		if (file && file.files) {
			formData.append("audioFile", file.files[0] as File);
		} else {
			console.error("meow");
		}
		await Requests.makeAuthRequest(
			"POST",
			"upload/file/music",
			formData,
			{},
			{},
			(b: FormData): FormData => b
		)
			.then((resp) => {
				console.log("ok: ", resp.ok);
			})
			.catch((err) => {
				console.error(err);
			});
	}
	async function uploadDirectory() {
		const formData = new FormData();
		if (files && files.files) {
			for (let file of Array.from(files.files)) {
				formData.append("audioFile", file);
				console.log(file.name);
			}
		} else {
			console.error("meow");
		}
		await Requests.makeAuthRequest(
			"POST",
			"upload/directory/music",
			formData,
			{},
			{},
			(b: FormData): FormData => b
		)
			.then((resp) => {
				console.log("ok: ", resp.ok);
			})
			.catch((err) => {
				console.error(err);
			});
	}
</script>

<input type="file" bind:this={file} on:change={handleChange} />
<button on:click={upload}>Upload File</button>
<br />
<input type="file" bind:this={files} webkitdirectory multiple />
<button on:click={uploadDirectory}>Upload Directory</button>
