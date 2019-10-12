package main

import (
	"github.com/mrz1836/postmark"
)

func main() {
	client := postmark.NewClient("[SERVER-TOKEN]", "[ACCOUNT-TOKEN]")

	email := postmark.Email{
		From:       "no-reply@example.com",
		To:         "tito@example.com",
		Subject:    "Reset your password",
		HtmlBody:   "...",
		TextBody:   "...",
		Tag:        "pw-reset",
		TrackOpens: true,
	}

	_, err := client.SendEmail(email)
	if err != nil {
		panic(err)
	}
}
