package facest

import (
	"net/http"
	"time"
)

const (
	BaseURLV1 = "http://api.facest.io/v1"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURLV1,
		apiKey:  appKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type seccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
