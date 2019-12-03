# Postmark

[![Build Status](https://travis-ci.com/mrz1836/postmark.svg?branch=master)](https://travis-ci.com/mrz1836/postmark)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrz1836/postmark)](https://goreportcard.com/report/github.com/mrz1836/postmark)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/0f79f6c5a19c41ce85fe0dc5a5c288b2)](https://www.codacy.com/app/mrz1818/postmark?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/postmark&amp;utm_campaign=Badge_Grade)
[![Release](https://img.shields.io/github/release-pre/mrz1836/postmark.svg?style=flat&v=1)](https://github.com/mrz1836/postmark/releases)
[![GoDoc](https://godoc.org/github.com/mrz1836/postmark?status.svg)](https://godoc.org/github.com/mrz1836/postmark)

A Golang package for the using Postmark API.

### Installation
```bash
go get -u github.com/mrz1836/postmark
```

### Basic Usage

Grab your [`Server Token`](https://account.postmarkapp.com/servers/XXXX/credentials), and your [`Account Token`](https://account.postmarkapp.com/account/edit).

```go
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
```
Swap out HTTPClient for use on Google App Engine:

```go
import (
    "github.com/mrz1836/postmark"
    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
)

// ....

client := postmark.NewClient("[SERVER-TOKEN]", "[ACCOUNT-TOKEN]")

ctx := appengine.NewContext(req)
client.HTTPClient = urlfetch.Client(ctx)

// ...
```

### API Coverage

* [x] Emails
    * [x] `POST /email`
    * [x] `POST /email/batch`
    * [x] `POST /email/withTemplate`
    * [x] `POST /email/batchWithTemplates`
* [x] Bounces
    * [x] `GET /deliverystats`
    * [x] `GET /bounces`
    * [x] `GET /bounces/:id`
    * [x] `GET /bounces/:id/dump`
    * [x] `PUT /bounces/:id/activate`
    * [x] `GET /bounces/tags`
* [ ] Templates
    * [x] `GET /templates`
    * [x] `POST /templates`
    * [x] `GET /templates/:id`
    * [x] `PUT /templates/:id`
    * [x] `DELETE /templates/:id`
    * [x] `POST /templates/validate`
* [x] Servers
    * [x] `GET /servers/:id`
    * [x] `PUT /servers/:id`
* [x] Outbound Messages
    * [x] `GET /messages/outbound`
    * [x] `GET /messages/outbound/:id/details`
    * [x] `GET /messages/outbound/:id/dump`
    * [x] `GET /messages/outbound/opens`
    * [x] `GET /messages/outbound/opens/:id`
* [x] Inbound Messages
    * [x] `GET /messages/inbound`
    * [x] `GET /messages/inbound/:id/details`
    * [x] `PUT /messages/inbound/:id/bypass`
    * [x] `PUT /messages/inbound/:id/retry`
* [ ] Sender signatures
    * [x] `GET /senders`
    * [ ] Get a sender signature’s details
    * [ ] Create a signature
    * [ ] Edit a signature
    * [ ] Delete a signature
    * [ ] Resend a confirmation
    * [ ] Verify an SPF record
    * [ ] Request a new DKIM
* [ ] Stats
    * [x] `GET /stats/outbound`
    * [x] `GET /stats/outbound/sends`
    * [x] `GET /stats/outbound/bounces`
    * [x] `GET /stats/outbound/spam`
    * [x] `GET /stats/outbound/tracked`
    * [x] `GET /stats/outbound/opens`
    * [x] `GET /stats/outbound/platform`
    * [ ] Get email client usage
    * [ ] Get email read times
* [ ] Triggers
    * [ ] Tags triggers
        * [ ] Create a trigger for a tag
        * [ ] Get a single trigger
        * [ ] Edit a single trigger
        * [ ] Delete a single trigger
        * [ ] Search triggers
    * [ ] Inbound rules triggers
        * [ ] Create a trigger for inbound rule
        * [ ] Delete a single trigger
        * [ ] List triggers
