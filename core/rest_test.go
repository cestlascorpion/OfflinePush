package core

import (
	"fmt"
	"testing"
	"time"
)

func TestGET(t *testing.T) {
	client, err := NewRestyClient(time.Second * 3)
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
	client, err := NewRestyClient(time.Second * 3)
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
