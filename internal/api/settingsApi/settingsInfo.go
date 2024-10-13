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
	switch name {
	case "site_info":
		res.OkWithData(c, conf.GlobalConfig.SiteInfo)
		return nil
	default:
		res.FailWithMsg(c, "参数错误")
	}
	return nil
}
