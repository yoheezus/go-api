package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetTrackInfoFromID takes a shared track ID and returns artistName and albumName
func GetTrackInfoFromID(token, fullTrackID string) (string, string, error) {
	const endpoint = "tracks/"

	trackID := strings.Trim(fullTrackID, "spotify:track:")
	requestURL := apiEndpoint + endpoint + trackID

	spotifyClient := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return "", "", err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := spotifyClient.Do(req)
	if err != nil {
		return "", "", err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	trackInfo := FullTrackObject{}
	err = json.Unmarshal(respBody, &trackInfo)

	trackArtistName := trackInfo.Artists[0].Name
	trackAlbumName := trackInfo.Album.Name

	return trackArtistName, trackAlbumName, nil
}
