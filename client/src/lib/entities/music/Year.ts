import type Music from "./Music";

export default interface Year {
	publicId: string;
	name: string;
	songs: Music[];
}
