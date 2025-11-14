package coinglassclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"
)

const (
	CustomHeaderNameKey = "CG-API-KEY"
	DefaultAPIUrl       = "https://open-api-v4.coinglass.com/api"
)

type Client struct {
	ApiUrl string
	secret string
	c      *http.Client
}

func NewClient(secret, url string, timeout time.Duration) *Client {
	return &Client{
		secret: secret,
		ApiUrl: url,
		c: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) Url(endpoint string) string {
	return fmt.Sprintf("%s%s", c.ApiUrl, endpoint)
}

type SuccessResponse struct {
	Code json.Number     `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

func (c *Client) doCall(ctx context.Context, req *Request, response any) (*http.Response, error) {
	if reflect.TypeOf(response).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("Response struct is not a pointer")
	}
	httpRequest, err := req.NewHttpRequest(ctx, c.secret)
	if err != nil {
		return nil, fmt.Errorf("api call %v() on %v: %v", req.Method, req.Endpoint, err.Error())
	}
	httpResponse, err := c.c.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("api call %v() on %v: %v", req.Method, httpRequest.URL.String(), err.Error())
	}

	bodyBytes, err := io.ReadAll(httpResponse.Body)
	// parsing error
	if err != nil {
		return nil, fmt.Errorf(
			"call %v() on %v status code: %v. could not decode body to response: %v",
			req.Method,
			httpRequest.URL.String(),
			httpResponse.StatusCode,
			err.Error())
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf(
			"call %v() on %v status code: %v.raw body %v",
			req.Method,
			httpRequest.URL.String(),
			httpResponse.StatusCode,
			string(bodyBytes))
	}

	var base200Response SuccessResponse

	err = json.Unmarshal(bodyBytes, &base200Response)

	if err != nil {
		return nil, fmt.Errorf(
			"call %v() on %v status code: %v. could not decode body to base 200 response model: %v",
			req.Method,
			httpRequest.URL.String(),
			httpResponse.StatusCode,
			err.Error())
	}
	code, err := base200Response.Code.Int64()
	if err != nil {
		return nil, fmt.Errorf("call %v() on %v status, in message code: %s, code Parsing error: %s",
			req.Method,
			httpRequest.URL.String(),
			base200Response.Code.String(),
			base200Response.Msg,
		)
	}
	if code >= 400 {
		return nil, fmt.Errorf("call %v() on %v status, in message code: %s, Error: %s",
			req.Method,
			httpRequest.URL.String(),
			base200Response.Code.String(),
			base200Response.Msg,
		)
	}
	err = json.Unmarshal(base200Response.Data, response)

	if err != nil {
		return nil, fmt.Errorf(
			"call %v() on %v status code: %v. could not decode body to response model: %v",
			req.Method,
			httpRequest.URL.String(),
			httpResponse.StatusCode,
			err.Error())
	}
	if response == nil {
		// if we have some http error, return it
		return nil, fmt.Errorf("call %v() on %v status code: %v. response missing",
			req.Method,
			httpRequest.URL.String(),
			httpResponse.StatusCode)
	}

	return httpResponse, nil
}

type Request struct {
	Body     any
	Endpoint string
	Method   string
}

func NewRequest(endpoint string, body any, methods ...string) *Request {
	method := http.MethodGet
	if len(methods) != 0 {
		method = methods[0]
	}
	return &Request{
		Body:     body,
		Endpoint: endpoint,
		Method:   method,
	}
}

func (r *Request) NewHttpRequest(ctx context.Context, secret string) (*http.Request, error) {
	var bodyReader io.Reader
	if r.Method != http.MethodGet {
		byteBody, err := json.Marshal(r.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(byteBody)
	}
	request, err := http.NewRequest(r.Method, r.Endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	//Custom headers
	request.Header.Set(CustomHeaderNameKey, secret)
	request = request.WithContext(ctx)

	return request, nil
}
