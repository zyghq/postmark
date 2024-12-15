package main

import (
	"context"
	"github.com/zyghq/postmark"
	"log"
)

func main() {
	client := postmark.NewClient("[SERVER-TOKEN]", "[ACCOUNT-TOKEN]")

	req := postmark.CreateDomainRequest{
		Name: "example.com",
	}
	domain, err := client.CreateDomain(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	log.Println(domain.ID)
	log.Println(domain.Name)
	log.Println(domain.SPFVerified)
	log.Println(domain.DKIMVerified)
	log.Println(domain.WeakDKIM)
	log.Println(domain.ReturnPathDomainVerified)

	log.Println("SPFHost:", domain.SPFHost)
	log.Println("SPFTextValue:", domain.SPFTextValue)
	log.Println("DKIMHost:", domain.DKIMHost)
	log.Println("DKIMTextValue:", domain.DKIMTextValue)
	log.Println("DKIMPendingHost:", domain.DKIMPendingHost)
	log.Println("DKIMPendingTextValue:", domain.DKIMPendingTextValue)
	log.Println("DKIMRevokedHost:", domain.DKIMRevokedHost)
	log.Println("DKIMRevokedTextValue:", domain.DKIMRevokedTextValue)
	log.Println("SafeToRemoveRevokedKey:", domain.SafeToRemoveRevokedKeyFromDNS)
	log.Println("DKIMUpdateStatus:", domain.DKIMUpdateStatus)
	log.Println("ReturnPathDomain:", domain.ReturnPathDomain)
	log.Println("ReturnPathDomainCNAME:", domain.ReturnPathDomainCNAMEValue)
}
