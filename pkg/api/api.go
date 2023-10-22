package api

type Album struct {
	URL string
	Title string
	Artist string
	Year int
	Image string
}

func GetAlbum(guid string) Album {
	result := Album {
		URL: "https://open.spotify.com/album/7uQZYsvK048nZWgB10cMbe?si=K7W2DAVlQ2W82w5DOIhUzA",
		Title: "Album Name",
		Year: 2023,
		Artist: "Artist Name",
		Image: "https://i.scdn.co/image/ab67616d0000b27373601c5b76c54845c4941b32",
	}

	return result
}
