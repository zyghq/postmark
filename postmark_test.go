package postmark

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"goji.io"
)

var (
	tMux    = goji.NewMux()
	tServer *httptest.Server
	client  *Client
)

func init() { //nolint:gochecknoinits // need to refactor this in the future
	tServer = httptest.NewServer(tMux)

	transport := &http.Transport{
		Proxy: func(_ *http.Request) (*url.URL, error) {
			// Reroute...
			return url.Parse(tServer.URL)
		},
	}

	client = NewClient("", "")
	client.HTTPClient = &http.Client{Transport: transport}
	client.BaseURL = tServer.URL
}
