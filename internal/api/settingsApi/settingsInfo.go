package settingsApi

import (
	"AI_Server/init/conf"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v3"
)

type SettingsApi struct {
}

func (settings *SettingsApi) SettingInfoView(c fiber.Ctx) error {
	name := c.Params("name")
	// todo 加入权限校验
	switch name {
	case "site_info":
		return res.OkWithData(c, conf.GlobalConfig.SiteInfo)
	default:
		return res.FailWithMsg(c, "参数错误")
	}
}
