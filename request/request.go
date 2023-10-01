package request

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

type RequestType int

const (
	Get RequestType = 1 << iota
	Post
)

func (e RequestType) String() string {
	switch e {
	case Get:
		return "GET"
	case Post:
		return "POST"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

func RequestWithContextAndData(ctx context.Context, typeRequest RequestType, url string, data []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, typeRequest.String(), url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func RequestWithContext(ctx context.Context, typeRequest RequestType, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, typeRequest.String(), url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Request(typeRequest RequestType, url string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(typeRequest.String(), url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
