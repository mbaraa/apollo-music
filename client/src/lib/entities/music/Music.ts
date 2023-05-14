export default interface Music {
	publicId: string;
	title: string;
	albumTitle: string;
	artistName: string;
	year: string;
	genre: string;
	audio: {
		fileName: string;
		fileSize: number;
		lastAccess: Date;
		accessTimes: number;
		publicPath: string;
		type: string;
	};
}
