package tests

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty"
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
