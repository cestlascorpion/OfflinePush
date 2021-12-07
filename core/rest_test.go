package core

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGET(t *testing.T) {
	client, err := NewRestyClient(http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	resp, err := client.GET("https://httpbin.org/get", "", nil)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(string(resp))
}

func TestGET2(t *testing.T) {
	client, err := NewRestyClient(http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	resp, err := client.GET("https://www.baidu.com", "", nil)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(string(resp))
}
