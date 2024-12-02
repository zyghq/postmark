package postmark

import (
	"context"
	"net/http"
	"testing"

	"goji.io/pat"
)

func TestGetSuppressions(t *testing.T) {
	responseJSON := `{
		"Suppressions":[
		  {
			"EmailAddress":"address@wildbit.com",
			"SuppressionReason":"ManualSuppression",
			"Origin": "Recipient",
			"CreatedAt":"2019-12-10T08:58:33-05:00"
		  },
		  {
			"EmailAddress":"bounce.address@wildbit.com",
			"SuppressionReason":"HardBounce",
			"Origin": "Recipient",
			"CreatedAt":"2019-12-11T08:58:33-05:00"
		  },
		  {
			"EmailAddress":"spam.complaint.address@wildbit.com",
			"SuppressionReason":"SpamComplaint",
			"Origin": "Recipient",
			"CreatedAt":"2019-12-12T08:58:33-05:00"
		  }
		]
	  }`

	tMux.HandleFunc(pat.Get("/message-streams/:StreamID/suppressions/dump"), func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(responseJSON))
	})

	res, err := client.GetSuppressions(context.Background(), "outbound", nil)
	if err != nil {
		t.Fatalf("GetSuppressions: %s", err.Error())
	}

	if len(res) != 3 {
		t.Fatalf("GetSuppressions: wrong number of suppression (%d)", len(res))
	}

	if res[0].EmailAddress != "address@wildbit.com" {
		t.Fatalf("GetSuppressions: wrong suppression email address: %s", res[0].EmailAddress)
	}

	responseJSON = `{
		"Suppressions":[
		  {
			"EmailAddress":"address@wildbit.com",
			"SuppressionReason":"ManualSuppression",
			"Origin": "Recipient",
			"CreatedAt":"2019-12-10T08:58:33-05:00"
		  }
		]
	  }`

	tMux.HandleFunc(pat.Get("/message-streams/:StreamID/suppressions/dump"), func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(responseJSON))
	})

	res, err = client.GetSuppressions(context.Background(), "outbound", map[string]interface{}{
		"emailaddress":      "address@wildbit.com",
		"fromdate":          "2019-12-10",
		"todate":            "2019-12-11",
		"suppressionreason": HardBounceReason,
		"origin":            RecipientOrigin,
	})
	if err != nil {
		t.Fatalf("%v", err.Error())
	}

	if len(res) != 1 {
		t.Fatalf("GetSuppressions: wrong number of suppression (%d)", len(res))
	}

	if res[0].EmailAddress != "address@wildbit.com" {
		t.Fatalf("GetSuppressions: wrong suppression email address: %s", res[0].EmailAddress)
	}
}

func TestCreateSuppressions(t *testing.T) {
	responseJSON := `{
		"Suppressions":[
		  {
			"EmailAddress":"good.address@wildbit.com",
			"Status":"Suppressed",
			"Message": null
		  },
		  {
			"EmailAddress":"spammy.address@wildbit.com",
			"Status":"Failed",
			"Message": "You do not have the required authority to change this suppression."
		  },
		  {
			"EmailAddress":"invalid-email-address",
			"Status":"Failed",
			"Message": "An invalid email address was provided."
		  }
		]
	  }`

	tMux.HandleFunc(pat.Post("/message-streams/:StreamID/suppressions"), func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(responseJSON))
	})

	res, err := client.CreateSuppressions(context.Background(), "outbound", []Suppression{})
	if err != nil {
		t.Fatalf("%v", err.Error())
	}

	if len(res) != 3 {
		t.Fatalf("CreateSuppressions: wrong number of suppression (%d)", len(res))
	}

	if res[0].EmailAddress != "good.address@wildbit.com" {
		t.Fatalf("CreateSuppressions: wrong suppression email address: %s", res[0].EmailAddress)
	}
}

func TestDeleteSuppressions(t *testing.T) {
	responseJSON := `{
		"Suppressions":[
		  {
			"EmailAddress":"good.address@wildbit.com"
		  },
		  {
			"EmailAddress":"not.suppressed@wildbit.com"
		  },
		  {
			"EmailAddress":"spammy.address@wildbit.com"
		  }
		]
	  }`

	tMux.HandleFunc(pat.Post("/message-streams/:StreamID/suppressions/delete"), func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(responseJSON))
	})

	res, err := client.DeleteSuppressions(context.Background(), "outbound", []Suppression{})
	if err != nil {
		t.Fatalf("%v", err.Error())
	}

	if len(res) != 3 {
		t.Fatalf("CreateSuppressions: wrong number of suppression (%d)", len(res))
	}

	if res[0].EmailAddress != "good.address@wildbit.com" {
		t.Fatalf("CreateSuppressions: wrong suppression email address: %s", res[0].EmailAddress)
	}
}
