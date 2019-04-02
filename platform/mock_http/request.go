package mock_http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

func RequestBody(body interface{}) io.ReadCloser {
	b, _ := json.Marshal(body)
	return ioutil.NopCloser(bytes.NewReader(b))
}
