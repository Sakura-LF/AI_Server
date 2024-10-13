package userApi

import (
	"AI_Server/common/email"
	"AI_Server/utils/rand"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v3"
	"github.com/mojocn/base64Captcha"
	"regexp"
)

type RegisterRequest struct {
	Value        string `json:"value"`
	RegisterType int8   `json:"registerType"` // 注册方式 1 手机号 2 邮箱
	Captcha      string `json:"captcha"`      // 图形验证码
	CaptchaID    string `json:"captchaID"`    // 图形验证码ID
	Code         string `json:"code"`         // 手机验证码
	Step         int8   `json:"step"`         // 步骤 1 第一步 2 第二步
}

func (userApi *UserApi) Register(c fiber.Ctx) error {
	// 1. 校验图形验证码
	// 2. 给邮箱/手机发送验证码
	// 3. 校验验证码
	var req RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return res.FailWithMsgAndReason(c, err.Error(), "请求参数错误")
	}
	switch req.Step {
	case 1:
		// 校验图形验证码
		ok := base64Captcha.DefaultMemStore.Verify(req.CaptchaID, req.Captcha, false)
		if !ok {
			return res.FailWithMsg(c, "图形验证码错误")
		}
		if req.RegisterType == 1 {
			// 校验是不是邮箱
			result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, req.Value)
			if !result {
				return res.FailWithMsg(c, "邮箱格式错误")
			}
			// 发邮件
			code := rand.GetRandomCode(6)
			if err := email.SendVerifyCode(req.Value, code); err != nil {
				return res.FailWithMsgAndReason(c, "发送验证码失败", err.Error())
			}
			// 将验证码存到redis中
		}
	case 2:
		return res.FailWithMsg(c, "注册功能暂未开放")
	}

	return nil
}
