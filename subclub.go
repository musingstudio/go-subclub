package subclub

import (
	"net/http"
	"time"
)

const (
	apiURL = "https://api.sub.club/public"

	timeoutSec = 20
)

type Client struct {
	Client *http.Client
	Config *ClientConfig

	APIKey string
}

func NewClient(key string) *Client {
	return &Client{
		Client: &http.Client{Timeout: timeoutSec * time.Second},
		Config: NewClientConfig(apiURL, "go-subclub"),
		APIKey: key,
	}
}
