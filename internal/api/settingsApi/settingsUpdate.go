package settingsApi

import (
	"AI_Server/configs"
	"AI_Server/init/conf"
	"AI_Server/utils/file"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v3"
)

func (settings *SettingsApi) SettingUpdateView(c fiber.Ctx) error {
	name := c.Params("name")
	switch name {
	case "site_info":
		siteInfo := new(configs.SiteInfo)
		if err := c.Bind().Body(siteInfo); err != nil {
			res.FailWithMsg(c, err.Error())
		}
		conf.GlobalConfig.SiteInfo = *siteInfo
		res.FailWithMsg(c, "更新成功")
		file.SetYaml(conf.GlobalConfig)
	}
	return nil
}
