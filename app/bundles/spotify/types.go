package spotify

// SimpleTrackObject contains basic information about a Track.
type SimpleTrackObject struct {
	Artists          []SimpleArtistObject `json:"artists"`
	AvailableMarkets []string             `json:"available_markets"`
	DiscNumber       int                  `json:"disc_number"`
	Duration         int                  `json:"duration_ms"`
	Explicit         bool                 `json:"explicit"`
	ExternalURLs     map[string]string    `json:"external_urls"`
	Endpoint         string               `json:"href"`
	ID               string               `json:"id"`
	Name             string               `json:"name"`
	PreviewURL       string               `json:"preview_url"`
	TrackNumber      int                  `json:"track_number"`
	URI              string               `json:"uri"`
}

// FullTrackObject contains additional Track Information ontop of SimpleTrackObject
type FullTrackObject struct {
	SimpleTrackObject
	Album       SimpleAlbumObject `json:"album"`
	ExternalIDs map[string]string `json:"external_ids"`
	Popularity  int               `json:"popularity"`
}

// SimpleArtistObject contains basic information about an Artist
type SimpleArtistObject struct {
	ExternalURLs map[string]string `json:"external_urls"`
	Followers    map[string]int    `json:"followers"`
	Genres       []string          `json:"genres"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Images       []ImageObject     `json:"images"`
	Name         string            `json:"name"`
	Popularity   int               `json:"popularity"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

// SimpleAlbumObject contains basic information about an Album
type SimpleAlbumObject struct {
	AlbumGroup           string               `json:"album_group,omitempty"`
	AlbumType            string               `json:"album_type"`
	Artists              []SimpleArtistObject `json:"artists"`
	AvailableMarkets     []string             `json:"available_markets"`
	ExternalURLs         map[string]string    `json:"external_urls"`
	Href                 string               `json:"href"`
	ID                   string               `json:"id"`
	Images               []ImageObject        `json:"images"`
	Name                 string               `json:"name"`
	ReleaseDate          string               `json:"release_date"`
	ReleaseDatePrecision string               `json:"release_date_precision"`
	Type                 string               `json:"type"`
	URI                  string               `json:"uri"`
}

// ImageObject contains details about an image
type ImageObject struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
