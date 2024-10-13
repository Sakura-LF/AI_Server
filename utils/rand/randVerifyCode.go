package rand

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

//var letterStrings = "0123456789"

func GetRandomCode(n int) string {
	// 初始化随机数生成器
	var builder strings.Builder
	for i := 0; i < n; i++ {

		builder.WriteString(strconv.Itoa(rand.IntN(10)))
	}
	return builder.String()
}
