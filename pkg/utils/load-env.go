package utils

import (
	"git.sxxfuture.net/filfi/letsfil/fil-data/pkg/log"
	"github.com/joho/godotenv"
	"os"
	"path"
)

func LoadEnv() error {
	// 环境变量存在 FILBASE_MANAGER_CONFIG，读对应文件
	// 失败则改读 path.Join(lo.must(os.Getwd()), ".env")
	// 失败则改读 path.Join(lo.must(os.UserHomeDir()), ".filbase-manager", "config.env")
	// 再失败直接 return nil

	confFile := os.Getenv("FILBASE_MANAGER_CONFIG")
	if confFile == "" {
		confFile = ".env"
	}
	// https://github.com/joho/godotenv
	err := godotenv.Load(confFile)
	if err != nil {
		log.Error("Error loading .env file")
		pwd, err := os.Getwd()
		if err != nil {
			log.Error("Error loading .env file")
			return err
		}
		err = godotenv.Load(path.Join(pwd, ".env"))
		if err != nil {
			log.Error("Error loading .env file")
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			err = godotenv.Load(path.Join(home, ".filbase-manager", "config.env"))
			if err != nil {
				log.Error("Error loading .env file")
				return err
			}
		}
	}
	return nil

}
