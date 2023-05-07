package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HTTPClient interface {
	Do(request *http.Request) (*http.Response, error)
}

type Client struct {
	baseURL string
	headers map[string]string
	client  HTTPClient
}

func NewClient(baseURL string, headers map[string]string, httpClient HTTPClient) *Client {
	return &Client{
		baseURL: baseURL,
		headers: headers,
		client:  httpClient,
	}
}

func (api *Client) Get(ctx context.Context, endpoint string, responseStruct interface{}) error {
	responseData, err := api.sendRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		return err
	}
	if responseData == nil {
		return nil
	}
	return json.Unmarshal(responseData, responseStruct)
}

func (api *Client) Post(ctx context.Context, endpoint string, requestData, responseStruct interface{}) error {
	data, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	responseData, err := api.sendRequest(ctx, "POST", endpoint, data)
	if err != nil {
		return err
	}
	if responseData == nil {
		return nil
	}

	if responseStruct == nil {
		return nil
	}

	return json.Unmarshal(responseData, responseStruct)
}

func (api *Client) Put(ctx context.Context, endpoint string, requestData, responseStruct interface{}) error {
	data, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	responseData, err := api.sendRequest(ctx, "PUT", endpoint, data)
	if err != nil {
		return err
	}
	if responseData == nil {
		return nil
	}
	return json.Unmarshal(responseData, responseStruct)
}

func (api *Client) Delete(ctx context.Context, endpoint string, responseStruct interface{}) error {
	responseData, err := api.sendRequest(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	if responseData == nil {
		return nil
	}
	return json.Unmarshal(responseData, responseStruct)
}

func (api *Client) sendRequest(ctx context.Context, method, endpoint string, data []byte) ([]byte, error) {
	u, err := url.Parse(api.baseURL)
	if err != nil {
		panic(err)
	}
	u.Path += endpoint

	if data == nil {
		data = []byte{}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	for k, v := range api.headers {
		req.Header.Set(k, v)
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New("non-2xx status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return nil, nil
	}

	return body, nil
}
