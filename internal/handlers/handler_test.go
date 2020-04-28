package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
)

func encodePayload(payload interface{}) io.Reader {
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(payload)

	return b
}
