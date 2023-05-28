import { popupMessage, popupType, showPopup as showpopup } from "$lib/store";

export async function sleep(time: number): Promise<void> {
	await new Promise((resolve) => setTimeout(resolve, time));
}

export function showPopup(message: string, type: "info" | "error") {
	showpopup.set(true);
	popupMessage.set(message);
	popupType.set(type);
	sleep(5000).then(() => {
		showpopup.set(false);
		popupMessage.set("");
	});
}
