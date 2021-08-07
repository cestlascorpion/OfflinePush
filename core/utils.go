package core

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func Signature(appKey string, masterSecret string) (signature string, timestamp string) {
	timestamp = strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	original := appKey + timestamp + masterSecret

	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum), timestamp
}
