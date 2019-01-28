package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Random(args []int, length int) []int {
	if len(args) <= 0 {
		return args
	}

	if length <= 0 || len(args) <= length {
		return args
	}

	for i := len(args) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		args[i], args[num] = args[num], args[i]
	}
	return args
}

func CreateCaptcha() string {
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
}