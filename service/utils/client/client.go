package client

import (
	"io"
	"net/http"
)

// TODO 封装三方Response
func Do(url string) (string, []byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "500", []byte{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "500", []byte{}, err
	}
	return resp.Status, body, nil
}
