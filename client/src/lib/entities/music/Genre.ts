import type Music from "./Music";

export default interface Genre {
	publicId: string;
	name: string;
	songs: Music[];
}
