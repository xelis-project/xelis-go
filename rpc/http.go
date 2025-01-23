package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

// not using a JSONRPC lib because we can only pass array or object in params based on https://www.jsonrpc.org/specification - sometime we want to pass a value (string, int) :S

type Http struct {
	RequestTimeout time.Duration
	Endpoint       *url.URL
	Header         http.Header
	client         *http.Client
}

func NewHttp(endpoint string, header http.Header) (*Http, error) {
	e, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	h := &Http{
		Endpoint:       e,
		RequestTimeout: 3 * time.Second,
		Header:         header,
		client:         &http.Client{},
	}

	return h, nil
}

func (h *Http) Request(method string, params interface{}, result interface{}) (res *http.Response, err error) {
	h.client.Timeout = h.RequestTimeout

	rpcRequest := RPCRequest{ID: 0, JSONRPC: "2.0", Method: method, Params: params}
	jsonParams, err := json.Marshal(rpcRequest)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", h.Endpoint.String(), bytes.NewBuffer(jsonParams))
	if err != nil {
		return
	}

	if h.Header != nil {
		req.Header = h.Header
	}

	res, err = h.client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	var rpcResponse RPCResponse
	err = json.Unmarshal(body, &rpcResponse)
	if err != nil {
		return
	}

	if rpcResponse.Error != nil {
		err = errors.New(rpcResponse.Error.Message)
		return
	}

	err = json.Unmarshal(rpcResponse.Result, result)
	return
}
