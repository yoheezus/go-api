package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	tokensEndpoint = "https://accounts.spotify.com/api/token"
	apiEndpoint    = "https://api.spotify.com/v1/"
)

// Auth contains Authentication details to be used in requests
type Auth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// RequestAuthenticationToken calls off to tokensEndpoint to retrieve a Bearer token
func RequestAuthenticationToken() (Auth, error) {
	spotifyCreds := os.Getenv("SPOTIFY_BASE64")

	body := url.Values{}
	body.Set("grant_type", "client_credentials")

	spotifyClient := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, tokensEndpoint, strings.NewReader(body.Encode()))
	if err != nil {
		return Auth{}, err
	}
	req.Header.Add("Authorization", "Basic "+spotifyCreds)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := spotifyClient.Do(req)
	if err != nil {
		return Auth{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Auth{}, err
	}

	sauth := Auth{}
	err = json.Unmarshal(respBody, &sauth)
	if err != nil {
		return Auth{}, err
	}

	fmt.Println(sauth.AccessToken)
	fmt.Println(sauth.ExpiresIn)
	fmt.Println(sauth.TokenType)

	return sauth, nil

}
