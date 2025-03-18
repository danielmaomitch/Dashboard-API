package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetPlaylists(accessTok string) Playlists {

	url := "https://api.spotify.com/v1/me/playlists?limit=20"

	bearer := "Bearer " + accessTok
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error formatting API Request")
	}
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error getting Spotify API Response ", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading API Response ", err)
	}

	playlists := Playlists{}
	if err := json.Unmarshal(respBody, &playlists); err != nil {
		log.Println("Error unmarhsalling playlists JSON ", err)
	}

	return playlists
}

func UpdatePlaylist(accessTok string, uri string, id string) {

	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks?uris=%s", id, uri)

	bearer := "Bearer " + accessTok
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println("Error formatting API Request ", err)
	}
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	client.Do(req)
}

func UpdatePlaylistWithCurrent(accessTok string, id string) {

	curTrack := GetCurrentTrack(accessTok)
	uri := curTrack.Item.URI
	UpdatePlaylist(accessTok, uri, id)
}
