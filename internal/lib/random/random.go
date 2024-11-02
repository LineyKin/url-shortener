package random

import (
	"time"
)

func NewRandomString(x int) string {
	str := make([]rune, x)
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	for i := 0; i < x; i++ {
		str[i] = chars[rune((time.Now().UnixMicro()/int64(i+1))%63)]
	}

	return string(str)
}
