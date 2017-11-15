package redis

import (
	"fmt"
	"testing"
)

func TestIO(t *testing.T) {
	REDIS_HOST = "127.0.0.1:6379"
	REDIS_DB = 2
	fmt.Println(Store("abc", "mmm", "gg"))
	result, err := Get("abc", "mmm")
	fmt.Println(string(result.([]byte)), err)
}
