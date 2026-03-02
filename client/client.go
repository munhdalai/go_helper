package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// RequestError нь HTTP хүсэлтийн алдааны мэдээллийг агуулна.
type RequestError struct {
	StatusCode   int
	ResponseData map[string]interface{}
	Err          error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

// RequestConfig нь HTTP хүсэлтийн тохиргоо.
//
// Url: хүсэлтийн URL (заавал)
// Method: HTTP method (default: GET)
// Headers: HTTP headers (default: Content-Type: application/json)
// Parameters: query parameters
// Body: хүсэлтийн бие (JSON болгон marshal хийгдэнэ)
// Timeout: хүсэлтийн timeout секундээр (default: 30)
type RequestConfig struct {
	Url        string
	Method     string
	Headers    *map[string]string
	Parameters *url.Values
	Body       interface{}
	Timeout    uint
}

// MakeHTTPRequest нь HTTP хүсэлт илгээж, хариуг generic type руу unmarshal хийнэ.
func MakeHTTPRequest[T any](config *RequestConfig) (*T, *RequestError) {
	if config.Method == "" {
		config.Method = "GET"
	}

	if config.Timeout == 0 {
		config.Timeout = 30
	}

	client := http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
	}

	u, err := url.Parse(config.Url)
	if err != nil {
		return nil, &RequestError{Err: err}
	}

	config.Method = strings.ToUpper(config.Method)

	q := u.Query()
	if config.Parameters != nil {
		for k, v := range *config.Parameters {
			q.Set(k, strings.Join(v, ","))
		}
	}
	u.RawQuery = q.Encode()

	var req *http.Request
	if config.Body != nil {
		jsonBytes, err := json.Marshal(config.Body)
		if err != nil {
			return nil, &RequestError{Err: err}
		}
		req, err = http.NewRequest(config.Method, u.String(), bytes.NewBuffer(jsonBytes))
		if err != nil {
			return nil, &RequestError{Err: err}
		}
	} else {
		req, err = http.NewRequest(config.Method, u.String(), nil)
		if err != nil {
			return nil, &RequestError{Err: err}
		}
	}

	req.Header.Add("Content-Type", "application/json")
	if config.Headers != nil {
		for k, v := range *config.Headers {
			req.Header.Set(k, v)
		}
	}

	log.Printf("%s %s\n", config.Method, req.URL.String())

	res, err := client.Do(req)
	if err != nil {
		return nil, &RequestError{Err: err}
	}

	if res == nil {
		return nil, &RequestError{Err: fmt.Errorf("calling %s returned empty response", u.String())}
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &RequestError{Err: err}
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		httpErr := &RequestError{
			StatusCode: res.StatusCode,
			Err:        errors.New("http request failed"),
		}
		json.Unmarshal(responseData, &httpErr.ResponseData)
		return nil, httpErr
	}

	responseObject := new(T)
	if err := json.Unmarshal(responseData, responseObject); err != nil {
		return nil, &RequestError{Err: fmt.Errorf("unmarshal response: %w", err)}
	}

	return responseObject, nil
}
