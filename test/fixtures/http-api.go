package fixtures

import (
	"bytes"
	"encoding/json"
)

func ComposeJSONBody(body any) *bytes.Buffer {
	b, _ := json.Marshal(body)
	return bytes.NewBuffer(b)
}
