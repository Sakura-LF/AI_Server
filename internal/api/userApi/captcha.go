package userApi

import (
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v2"
	"github.com/mojocn/base64Captcha"
)

func (userApi *UserApi) Captcha(c *fiber.Ctx) error {
	driverString := base64Captcha.NewDriverString(
		80,
		240,
		5,
		5,
		5,
		"0123456789",
		nil,
		nil,
		nil,
	)
	dirver := base64Captcha.NewCaptcha(driverString, base64Captcha.DefaultMemStore)
	id, s, _, err := dirver.Generate()
	if err != nil {
		return res.FailWithMsgAndReason(c, err.Error(), "生成验证码失败")
	}

	return res.OkWithData(c, map[string]any{
		"captchaId": id,
		"captcha":   s,
	})
}
