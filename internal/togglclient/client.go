package togglclient

import (
	"encoding/base64"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func New() *Client {
	return &Client{
		HTTPClient: &http.Client{},
		BaseURL:    "https://api.track.toggl.com/api/v9",
	}
}

func encodeToken(token string) string {
	creds := token + ":api_token"
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(creds))
}
