package dialogflow

type Request struct {
	OriginalRequest OriginalRequest `json:originalRequest`
	ID              string          `json:id`
	Result          Result          `json:result`
	SessionID       string          `json:sessionId`
	Status          Status          `json:status`
}

type OriginalRequest struct {
	Source string `json:source`
	Data   struct {
		Inputs []Input `json:inputs`
		User   User    `json:user`
	} `json:data`
}

type Input struct {
	Intent    string     `json:intent`
	RaqInputs []RawInput `json:"raw_inputs"`
}

type RawInput struct {
	Query string `json:query`
	Type  int    `json:"input_type"`
}

type User struct {
	ID     string `json:"user_id"`
	Locale string `json:locale`
}

type Result struct {
	Source        string                 `json:source`
	ResolvedQuery string                 `json:resolvedQuery`
	Fulfillment   Fulfillment            `json:fulfillment`
	MetaData      MetaData               `json:metadata`
	Params        map[string]interface{} `json:"parameters"`
}

type MetaData struct {
	IntentID   string `json:intentId`
	IntentName string `json:intentName`
}

type Fulfillment struct {
	Speech   string    `json:speech`
	Messages []Message `json:messages`
}

type Message struct {
	Type   int    `json:type`
	Speech string `json:speech`
}

type Status struct {
	Code      int    `json:code`
	ErrorType string `json:errorType`
}
