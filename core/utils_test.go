package core

import (
	"fmt"
	"testing"
)

const (
	TestAppKey       = "bgiopsdfghoenrbgof84cn"
	TestMasterSecret = "bnofbp0jq3pg9uy5906afc"
)

func TestSignature(t *testing.T) {
	sign, tp := Signature(TestAppKey, TestMasterSecret)
	fmt.Println(sign, tp)
}
