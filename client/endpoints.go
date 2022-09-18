package client

import (
	"io"
	"net/url"
)

type ApiEndpoint struct {
	ApiUrl string
}

func NewApiEndpoint(template, host string) *ApiEndpoint {
	u, err := url.Parse(template)
	if err != nil {
		panic(err)
	}
	u.Host = host
	return &ApiEndpoint{
		ApiUrl: u.String(),
	}
}

func (e *ApiEndpoint) Put(contentType string, body io.Reader) []byte {
	resp, _ := Put(e.ApiUrl, contentType, body)
	data, _ := ReadResponse(resp)
	return data
}

func (e *ApiEndpoint) Get() []byte {
	resp, _ := DefaultClient.Get(e.ApiUrl)
	data, _ := ReadResponse(resp)
	return data
}
