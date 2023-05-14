import type Music from "./Music";

export default interface Album {
	publicId: string;
	title: string;
	artistName: string;
	year: string;
	genre: string;
	songs: Music[];
}
