package core

import (
	"fmt"
	"testing"
	"time"
)

func TestGET(t *testing.T) {
	resp, err := GET("https://httpbin.org/get", "", nil, time.Second*3)
	if err != nil {
		t.Failed()
	}
	fmt.Println(string(resp))
}

func TestGET2(t *testing.T) {
	resp, err := GET("https://www.baidu.com", "", nil, time.Second*3)
	if err != nil {
		t.Failed()
	}
	fmt.Println(string(resp))
}
