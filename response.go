package dialogflow

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Speech             string
	DisplayText        string
	ExpectUserResponse bool
	IsSSML             bool
}

func NewSSMLResponse(ssml, display string, expectedUserResponse bool) Response {
	return Response{
		Speech:             fmt.Sprintf("<speak>%s</speak>", ssml),
		DisplayText:        display,
		ExpectUserResponse: expectedUserResponse,
		IsSSML:             true,
	}
}

func (r Response) ToJSON() ([]byte, error) {
	body := map[string]interface{}{
		"speech":      r.Speech,
		"displayText": r.DisplayText,
		"data": map[string]interface{}{
			"google": map[string]interface{}{
				"expect_user_response": r.ExpectUserResponse,
				"is_ssml":              r.IsSSML,
			},
		},
	}
	return json.Marshal(body)
}
