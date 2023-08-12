// Package postmark ...
package postmark

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var postmarkURL = `https://api.postmarkapp.com`

// Client provides a connection to the Postmark API
type Client struct {
	// HTTPClient is &http.Client{} by default
	HTTPClient *http.Client
	// Server Token: Used for requests that require server level privileges. This token can be found on the Credentials tab under your Postmark server.
	ServerToken string
	// AccountToken: Used for requests that require account level privileges. This token is only accessible by the account owner, and can be found on the Account tab of your Postmark account.
	AccountToken string
	// BaseURL is the root API endpoint
	BaseURL string
}

const (
	accountToken = "account"
	serverToken  = "server"
)

// Options is an object to hold variable parameters to perform request.
type parameters struct {
	// Method is HTTP method type.
	Method string
	// Path is postfix for URI.
	Path string
	// Payload for the request.
	Payload interface{}
	// TokenType defines which token to use
	TokenType string
}

// NewClient builds a new Client pointer using the provided tokens, a default HTTPClient, and a default API base URL
// Accepts `Server Token`, and `Account Token` as arguments
// http://developer.postmarkapp.com/developer-api-overview.html#authentication
func NewClient(serverToken string, accountToken string) *Client {
	return &Client{
		HTTPClient:   &http.Client{},
		ServerToken:  serverToken,
		AccountToken: accountToken,
		BaseURL:      postmarkURL,
	}
}

func (client *Client) doRequest(ctx context.Context, opts parameters, dst interface{}) (err error) {
	url := fmt.Sprintf("%s/%s", client.BaseURL, opts.Path)

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, opts.Method, url, nil)
	if err != nil {
		return
	}

	if opts.Payload != nil {
		var payloadData []byte
		payloadData, err = json.Marshal(opts.Payload)
		if err != nil {
			return
		}
		req.Body = io.NopCloser(bytes.NewBuffer(payloadData))
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	switch opts.TokenType {
	case accountToken:
		req.Header.Add("X-Postmark-Account-Token", client.AccountToken)

	default:
		req.Header.Add("X-Postmark-Server-Token", client.ServerToken)
	}

	var res *http.Response
	res, err = client.HTTPClient.Do(req)
	if err != nil {
		return
	}

	defer func() {
		_ = res.Body.Close()
	}()
	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode >= 400 {
		// If the status code is not a success, attempt to unmarshall the body into the APIError struct.
		var apiErr APIError
		err = json.Unmarshal(body, &apiErr)
		if err != nil {
			return
		}
		return apiErr
	}

	err = json.Unmarshal(body, dst)
	return
}

// APIError represents errors returned by Postmark
type APIError struct {
	// ErrorCode: see error codes here (https://postmarkapp.com/developer/api/overview#error-codes)
	ErrorCode int64 `json:"ErrorCode"`
	// Message contains error details
	Message string `json:"Message"`
}

// Error returns the error message details
func (res APIError) Error() string {
	return res.Message
}
