package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	SetDefaultPostgresConfig()
	SetDefaultServerConfig()
}

func InitConfig() {
	viper.SetConfigName("main")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/vk_edu_db_project/configs")
	viper.AddConfigPath("/forum/configs")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}

func SetDefaultPostgresConfig() {
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.username", "forum")
	viper.SetDefault("db.name", "forum")
	viper.SetDefault("db.password", "password")
}

func SetDefaultServerConfig() {
	viper.SetDefault("host", "forum")
	viper.SetDefault("port", "5000")
	viper.SetDefault("metrics_host", "forum")
	viper.SetDefault("metrics_port", "8080")
}
