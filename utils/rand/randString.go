package rand

import (
	"github.com/google/uuid"
	"strings"
)

//var letterStrings = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetRandomUserName() (string, error) {
	var builder strings.Builder
	v7, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	uuid.NewString()
	//split := strings.Split(v7.String(), "-")
	builder.WriteString(v7.String())
	return builder.String(), err
}

func GetRandomNickName(str string) string {
	var builder strings.Builder
	strs := strings.Split(str, "-")
	builder.WriteString("用户")
	builder.WriteString(strs[4])
	return builder.String()
}
