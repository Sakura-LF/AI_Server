package validate

import "regexp"

var regEmail = regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
var regTel = regexp.MustCompile("^1[345789]{1}\\d{9}$") // 假设手机号为11位数字

// ValidateEmail 校验邮箱格式
func ValidateEmail(email string) bool {
	return regEmail.MatchString(email)
}

// ValidatePhone 校验电话号码格式
func ValidatePhone(phone string) bool {
	return regTel.MatchString(phone)
}
