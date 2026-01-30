package SpotifyEntity

// #TODO Api Pública não está funcionando
type Playlist struct {
	Name   string
	Tracks []Tracks
}

type Tracks struct {
	Name   string
	Artist string
	Link   string
}
