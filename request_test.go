package apiai

import (
	"encoding/json"
	"testing"
)

func Test_1(t *testing.T) {
	var r Request
	err := json.Unmarshal(input1, &r)
	if err != nil {
		t.Errorf("Failed to unmarshal : %s", err)
		return
	}
	// Assertion
	// OriginalRequest
	originalRequest := r.OriginalRequest
	if originalRequest.Data.User.ID != "1502731363742" {
		t.Errorf("Wrong User ID : %s", originalRequest.Data.User.ID)
		return
	}
	if originalRequest.Data.User.Locale != "en-US" {
		t.Errorf("Wrong User Locale : %s", originalRequest.Data.User.Locale)
		return
	}

	if r.ID != "ff0f8352-772a-4af9-b1d2-db7d2c1ca8e0" {
		t.Errorf("Wrong ID : %s", r.ID)
		return
	}
	if r.SessionID != "1502731363742" {
		t.Errorf("Wrong SessionID : %s", r.SessionID)
		return
	}

	result := r.Result
	if result.Source != "agent" {
		t.Errorf("Wrong Source : %s", result.Source)
		return
	}
	fulfillment := result.Fulfillment
	if fulfillment.Speech != "Hi!" {
		t.Errorf("Wrong Speech : %s", fulfillment.Speech)
		return
	}
	if len(fulfillment.Messages) != 1 {
		t.Errorf("Wrong Message count : %d", len(fulfillment.Messages))
		return
	}
	msg := fulfillment.Messages[0]
	if msg.Type != 0 {
		t.Errorf("Wrong Message type : %d", msg.Type)
		return
	}
	if msg.Speech != "Hi!" {
		t.Errorf("Wrong Speech : %s", msg.Speech)
		return
	}
	t.Errorf("OK")
}
