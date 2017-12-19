# dialogflow
dialogflow client library. This library supports api.ai webhook.

# parse request from dialogflow

use `json.Unmarshal()`

```
func handler(w http.ResponseWriter, r *http.Request) {
        buf := new(bytes.Buffer)
        buf.ReadFrom(r.Body)
        var request dialogflow.Request
        err := json.Unmarshal(buf.Bytes(), &request)
        ...
```

# build response for dialogflow

```
// w is http.ResponseWriter
resp = dialogflow.NewSSMLResponse(
    "speak response",
    "displayText response",
    true,
)    
data, _ := resp.ToJSON()
w.Write(data)
```                
