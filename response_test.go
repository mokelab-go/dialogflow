package dialogflow

import (
	"encoding/json"
	"testing"
)

func getString(json map[string]interface{}, key string) string {
	v, ok := json[key]
	if !ok {
		return "not found"
	}
	if v2, ok := v.(string); ok {
		return v2
	} else {
		return "not string"
	}
}

func getBool(json map[string]interface{}, key string) bool {
	v, ok := json[key]
	if !ok {
		return false
	}
	if v2, ok := v.(bool); ok {
		return v2
	} else {
		return false
	}
}

func getObject(json map[string]interface{}, key string) map[string]interface{} {
	v, ok := json[key]
	if !ok {
		return map[string]interface{}{}
	}
	if v2, ok := v.(map[string]interface{}); ok {
		return v2
	} else {
		return map[string]interface{}{}
	}
}

func TestReponse_1(t *testing.T) {
	r := NewSSMLResponse("Hi", "Hi!", true)
	b, err := r.ToJSON()
	if err != nil {
		t.Errorf("Failed to build json : %s", err)
		return
	}
	// Assertion
	var j map[string]interface{}
	err = json.Unmarshal(b, &j)
	if err != nil {
		t.Errorf("Failed to parse json : %s", err)
		return
	}
	if len(j) != 3 {
		t.Errorf("Unexpected root count : %d", len(j))
		return
	}
	v := getString(j, "speech")
	if v != "<speak>Hi</speak>" {
		t.Errorf("Unexpected speech value : %s", v)
		return
	}
	v = getString(j, "displayText")
	if v != "Hi!" {
		t.Errorf("Unexpected displayText value : %s", v)
		return
	}
	o := getObject(j, "data")
	if len(o) != 1 {
		t.Errorf("Unexpected data object : %s", o)
		return
	}
	o = getObject(o, "google")
	if len(o) != 2 {
		t.Errorf("Unexpected data object : %s", o)
		return
	}
	b2 := getBool(o, "expect_user_response")
	if !b2 {
		t.Errorf("Unexpected expect_user_response : %s", o)
		return
	}
	b2 = getBool(o, "is_ssml")
	if !b2 {
		t.Errorf("Unexpected is_ssml : %s", o)
		return
	}
}
