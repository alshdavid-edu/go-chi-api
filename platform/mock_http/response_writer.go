package mock_http

import (
	"encoding/json"
	"net/http"
)

type ResponseWriter struct {
	Written    []byte
	StatusCode int
}

func (m *ResponseWriter) Write(w []byte) (int, error) {
	m.Written = w
	return 0, nil
}

func (m *ResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m *ResponseWriter) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}

func (m *ResponseWriter) GetBodyJSON() map[string]interface{} {
	var v map[string]interface{}
	json.Unmarshal(m.Written, &v)
	return v
}

func (m *ResponseWriter) GetBodyJSONArray() []map[string]interface{} {
	var v []map[string]interface{}
	json.Unmarshal(m.Written, &v)
	return v
}

func (m *ResponseWriter) GetBodyString() string {
	return string(m.Written)
}
