package strings

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixMilli())
	n := rand.Intn(1_000_000_000-987654) + 987654
	return fmt.Sprint(n)[:6]
}

func GeneratePublicId() string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	id := new(strings.Builder)
	for i := 0; i < 16; i++ {
		rand.Seed(time.Now().UnixMilli())
		index := rand.Intn(len(alphabet))
		id.WriteByte(alphabet[index])
		time.Sleep(time.Millisecond * 5)
	}
	return id.String()
}
