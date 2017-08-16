# apiai
api.ai client library. This library supports api.ai webhook.

# parse request from api.ai

use `json.Unmarshal()`

```
func handler(w http.ResponseWriter, r *http.Request) {
        buf := new(bytes.Buffer)
        buf.ReadFrom(r.Body)
        var request apiai.Request
        err := json.Unmarshal(buf.Bytes(), &request)
        ...
```

# build response for api.ai

```
// w is http.ResponseWriter
resp = apiai.NewSSMLResponse(
    "speak response",
    "displayText response",
    true,
)    
data, _ := resp.ToJSON()
w.Write(data)
```                
