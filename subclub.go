package subclub

import (
	"net/http"
	"time"
)

const (
	apiURL = "https://api.sub.club/public"

	timeoutSec = 20
)

// Client holds an http.Client and ClientConfig for interacting with the API.
type Client struct {
	Client *http.Client
	Config *ClientConfig

	APIKey string
}

// NewClient creates a new Client with the supplied sub.club API key.
func NewClient(key string) *Client {
	return &Client{
		Client: &http.Client{Timeout: timeoutSec * time.Second},
		Config: NewClientConfig(apiURL, "go-subclub"),
		APIKey: key,
	}
}
