import type Music from "./Music";

export default interface Artist {
	publicId: string;
	name: string;
	songs: Music[];
}
