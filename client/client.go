package client

import (
	"io"
	"net/http"
)

var DefaultClient = &http.Client{}

func Put(url, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ReadResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
