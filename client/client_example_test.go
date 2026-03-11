package client_test

import (
	"fmt"
	"net/url"

	"github.com/munhdalai/go_helper/client"
)

func ExampleMakeHTTPRequest() {
	// Энгийн GET хүсэлт
	type Response struct {
		Origin string `json:"origin"`
	}

	result, reqErr := client.MakeHTTPRequest[Response](&client.RequestConfig{
		Url: "https://httpbin.org/get",
	})
	if reqErr != nil {
		fmt.Println("Алдаа:", reqErr.Error())
		return
	}
	fmt.Println(result.Origin)
}

func ExampleMakeHTTPRequest_post() {
	// POST хүсэлт body-тэй
	type RequestBody struct {
		Name string `json:"name"`
	}
	type Response struct {
		JSON RequestBody `json:"json"`
	}

	result, reqErr := client.MakeHTTPRequest[Response](&client.RequestConfig{
		Url:    "https://httpbin.org/post",
		Method: "POST",
		Body:   RequestBody{Name: "Бат"},
	})
	if reqErr != nil {
		fmt.Println("Алдаа:", reqErr.Error())
		return
	}
	fmt.Println(result.JSON.Name)
}

func ExampleMakeHTTPRequest_withHeaders() {
	// Custom header-тэй хүсэлт
	type Response struct {
		Headers map[string]string `json:"headers"`
	}

	headers := map[string]string{
		"Authorization": "Bearer mytoken123",
		"X-Custom":      "test",
	}

	_, reqErr := client.MakeHTTPRequest[Response](&client.RequestConfig{
		Url:     "https://httpbin.org/get",
		Headers: &headers,
		Timeout: 10,
	})
	if reqErr != nil {
		fmt.Println("Алдаа:", reqErr.Error())
		return
	}
}

func ExampleMakeHTTPRequest_withParameters() {
	// Query parameter-тэй хүсэлт
	type Response struct {
		Args map[string]string `json:"args"`
	}

	params := url.Values{}
	params.Set("search", "golang")
	params.Set("page", "1")

	result, reqErr := client.MakeHTTPRequest[Response](&client.RequestConfig{
		Url:        "https://httpbin.org/get",
		Parameters: &params,
	})
	if reqErr != nil {
		fmt.Println("Алдаа:", reqErr.Error())
		return
	}
	fmt.Println(result.Args["search"])
}

func ExampleRequestError_Error() {
	err := &client.RequestError{
		StatusCode: 404,
		Err:        fmt.Errorf("not found"),
	}
	fmt.Println(err.Error())
	// Output: status 404: err not found
}
