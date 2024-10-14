package userApi

import (
	"AI_Server/common/message"
	"AI_Server/init/data"
	"AI_Server/internal/data/mysql/user"
	"AI_Server/internal/modeles"
	"AI_Server/utils/res"
	"AI_Server/utils/validate"
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/mojocn/base64Captcha"
)

type RegisterRequest struct {
	Value        string                 `json:"value"`
	RegisterType modeles.RegisterSource `json:"registerType"` // 注册方式 0 邮箱 1 手机 2 微信
	Captcha      string                 `json:"captcha"`      // 图形验证码
	CaptchaID    string                 `json:"captchaID"`    // 图形验证码ID
	Code         string                 `json:"code"`         // 手机/邮箱验证码
	Step         int8                   `json:"step"`         // 步骤 1 第一步 2 第二步
}

func (userApi *UserApi) Register(c fiber.Ctx) error {
	// 校验请求参数
	var req RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return res.FailWithMsgAndReason(c, err.Error(), "请求参数错误")
	}
	// 根据 RegisterType 校验 邮箱/电话
	switch req.RegisterType {
	case modeles.EmailRegister:
		if !validate.ValidateEmail(req.Value) {
			return res.FailWithMsg(c, "邮箱格式错误")
		}
	case modeles.TelRegister:
		if !validate.ValidatePhone(req.Value) {
			return res.FailWithMsg(c, "邮箱格式错误")
		}
	default:
		return res.FailWithMsg(c, "未知注册方式")
	}

	// 校验步骤
	switch req.Step {
	case 1:
		if req.Captcha == "" || req.CaptchaID == "" {
			return res.FailWithMsg(c, "请输入图像验证码")
		}
		// todo 校验错误次数
		// 校验图形验证码
		ok := base64Captcha.DefaultMemStore.Verify(req.CaptchaID, req.Captcha, false)
		if !ok {
			return res.FailWithMsg(c, "图形验证码错误")
		}
		// 根据注册类型发送验证码
		go func() {
			err := message.SendCode(modeles.EmailRegister, req.Value)
			if err != nil {
				return
			}
		}()
		return res.OkWithMsg(c, "发送验证码成功,请注意查收")
	case 2:
		if req.Code == "" {
			return res.FailWithMsg(c, "验证码不能为空")
		}
		// 获取验证码
		code, err := data.RDB.Get(context.Background(), req.Value).Result()
		if err != nil {
			return res.FailWithMsg(c, "验证码已过期")
		}
		if code != req.Code {
			return res.FailWithMsg(c, "验证码错误")
		}
		// 删除验证码,直接同意判断 del的状态
		if intCmd := data.RDB.Del(context.Background(), req.Value); intCmd.Val() == 0 {
			return res.FailWithMsg(c, "验证码已过期,请重新获取验证码")
		}

		// 创建用户
		if err := user.CreateUser(modeles.EmailRegister, req.Value); err != nil {
			return res.FailWithMsgAndReason(c, "注册失败", err.Error())
		}
		return res.OkWithMsg(c, "注册成功")
	}

	return res.FailWithMsg(c, "未知错误")
}
