package postmark

import (
	"context"
	"net/http"
)

type Domain struct {
	ID                       int64
	Name                     string
	SPFVerified              bool
	DKIMVerified             bool
	WeakDKIM                 bool
	ReturnPathDomainVerified bool
}

type DomainDetail struct {
	Domain
	SPFHost                       string
	SPFTextValue                  string
	DKIMHost                      string
	DKIMTextValue                 string
	DKIMPendingHost               string
	DKIMPendingTextValue          string
	DKIMRevokedHost               string
	DKIMRevokedTextValue          string
	SafeToRemoveRevokedKeyFromDNS bool
	DKIMUpdateStatus              string
	ReturnPathDomain              string
	ReturnPathDomainCNAMEValue    string
}

type CreateDomainRequest struct {
	// Name represents the domain name
	Name string `json:"Name"`
	// ReturnPathDomain is optional field but must be a subdomain of your From email domain.
	ReturnPathDomain *string `json:"ReturnPathDomain,omitempty"`
}

// CreateDomain creates a domain
func (client *Client) CreateDomain(ctx context.Context, req CreateDomainRequest) (DomainDetail, error) {
	res := DomainDetail{}
	err := client.doRequest(ctx, parameters{
		Method:    http.MethodPost,
		Path:      "domains",
		TokenType: accountToken,
		Payload:   req,
	}, &res)
	return res, err
}
