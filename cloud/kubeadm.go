package cloud

import (
	"fmt"
	mrnd "math/rand"
	"time"
)

func GetKubeadmToken() string {
	return fmt.Sprintf("%s.%s", RandStringRunes(6), RandStringRunes(16))
}

func init() {
	mrnd.Seed(time.Now().UnixNano())
}

// Hexidecimal
var letterRunes = []rune("0123456789abcdef")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[mrnd.Intn(len(letterRunes))]
	}
	return string(b)
}
