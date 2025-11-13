package catalogue

import "github.com/NOLLYWOOD-COM/go-sdk/internal/httpclient"

type WorkSvc struct {
	httpClient httpclient.Client
}

type Work struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	Title      string `json:"title"`
}
