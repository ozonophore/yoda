package test

import (
	"bytes"
	json2 "encoding/json"
	"io"
	"net/http"
)

type MockClient struct {
	DoMatch func(method, path string, req *http.Request) any
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	obj := m.DoMatch(req.Method, req.URL.Path, req)
	json, _ := json2.Marshal(obj)

	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(json)),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
	}, nil
}

func NewMockClient(matcher func(method, path string, req *http.Request) any) *MockClient {
	return &MockClient{
		DoMatch: matcher,
	}

}
