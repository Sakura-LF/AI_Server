package rand

import (
	rand2 "golang.org/x/exp/rand"
	"strconv"
	"strings"
)

//var letterBytes = []byte("0123456789")

func GetRandomCode(n int) string {
	//rand.New(rand2.NewSource(uint64(time.Now().UnixNano())))
	// 初始化随机数生成器
	//rand.New()
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(strconv.Itoa(rand2.Intn(10)))
	}
	return builder.String()
}
