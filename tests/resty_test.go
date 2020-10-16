package tests

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestResty(t *testing.T) {

	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(resp.Body()))
}
