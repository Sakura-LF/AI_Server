package file

import (
	"AI_Server/configs"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
)

func SetYaml(c *configs.Config) error {
	byteData, _ := yaml.Marshal(c)

	err := os.WriteFile("configs/config.yaml", byteData, 0600)
	if err != nil {
		log.Error().Err(err).Msg("写入配置文件失败")
		return err
	}
	log.Info().Msg("写入配置文件成功")
	return nil
}
