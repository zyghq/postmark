package main

import (
	"context"

	"github.com/zyghq/postmark"
)

func main() {
	client := postmark.NewClient("[SERVER-TOKEN]", "[ACCOUNT-TOKEN]")

	email := postmark.Email{
		From:       "no-reply@example.com",
		To:         "tito@example.com",
		Subject:    "Reset your password",
		HTMLBody:   "...",
		TextBody:   "...",
		Tag:        "pw-reset",
		TrackOpens: true,
	}

	_, err := client.SendEmail(context.Background(), email)
	if err != nil {
		panic(err)
	}
}
