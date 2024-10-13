package rand

import (
	"fmt"
	"testing"
)

func TestGerRandomCode(t *testing.T) {
	for i := 0; i < 10; i++ {
		code := GetRandomCode(6)
		fmt.Println(code)
	}
}
