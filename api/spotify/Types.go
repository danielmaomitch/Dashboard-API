package spotify

type LogIn struct {
	Name        string `json:"name"`
	AccessToken string `json:"access-token"`
	ExpiryDate  string `json:"expiry-date"`
}

type SpotifyReq struct {
	AccessToken string `json:"access-token"`
	ID          string `json:"id"`
	URI         string `json:"uri"`
}

type CurrentlyPlaying struct {
	Item struct {
		Album struct {
			Href   string `json:"href"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"artists"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"item"`
	IsPlaying bool `json:"is_playing"`
}

type Playlists struct {
	Href  string `json:"href"`
	Limit int    `json:"limit"`
	Next  string `json:"next"`
	Total int    `json:"total"`
	Items []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		ID     string `json:"id"`
		Images []struct {
			Height int    `json:"height"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name  string `json:"name"`
		Owner struct {
			DisplayName string `json:"display_name"`
		} `json:"owner"`
		URI string `json:"uri"`
	} `json:"items"`
}
