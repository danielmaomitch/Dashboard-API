package spotify

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Play(accessTok string) {

	url := "https://api.spotify.com/v1/me/player/play"

	bearer := "Bearer " + accessTok
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Println("Error formatting API Request")
	}
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	client.Do(req)
}

func Pause(accessTok string) {
	url := "https://api.spotify.com/v1/me/player/pause"

	bearer := "Bearer " + accessTok
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Println("Error formatting API Request")
	}
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	client.Do(req)
}

func GetCurrentTrack(accessTok string) CurrentlyPlaying {

	url := "https://api.spotify.com/v1/me/player/currently-playing"

	bearer := "Bearer " + accessTok
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error formatting API Request ", err)
		return CurrentlyPlaying{}
	}
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error getting Spotify API response ", err)
		return CurrentlyPlaying{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading Spotify API response ", err)
		return CurrentlyPlaying{}
	}

	track := CurrentlyPlaying{}
	if err := json.Unmarshal(body, &track); err != nil {
		log.Println("Error unmarshalling currently-playing JSON; May be nil ", err)
		return CurrentlyPlaying{}
	}

	return track
}
