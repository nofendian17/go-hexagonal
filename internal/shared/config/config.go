package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type (
	Config struct {
		App      app      `json:"app" validate:"required"`
		Database database `json:"database" validate:"required"`
		Log      log      `json:"log" validate:"required"`
	}

	app struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Version     string `json:"version" validate:"required"`
		Port        int    `json:"port" validate:"required"`
		Debug       bool   `json:"debug" json:"debug" validate:"required"`
		Key         string `json:"key" validate:"required"`
	}

	database struct {
		Pgsql struct {
			Host     string `json:"host" validate:"required"`
			Port     int    `json:"port" validate:"required"`
			Database string `json:"database" validate:"required"`
			Schema   string `json:"schema" validate:"required"`
			Username string `json:"username" validate:"required"`
			Password string `json:"password"  validate:"required"`
		} `json:"pgsql" validate:"required"`
		Redis struct {
			Host     string `json:"host" validate:"required"`
			Port     int    `json:"port" validate:"required"`
			Prefix   string `json:"prefix" validate:"required"`
			Lifetime int    `json:"lifetime" validate:"required"`
		} `json:"redis" validate:"required"`
	}

	log struct {
		FileLocation    string        `json:"fileLocation" validate:"required"`
		FileTdrLocation string        `json:"fileTdrLocation" validate:"required"`
		FileMaxAge      time.Duration `json:"fileMaxAge" validate:"required"`
		Stdout          bool          `json:"stdout"`
	}
)

func New() *Config {
	path := "configs.json"
	fmt.Println("Trying load configs:", path)
	viper.SetConfigFile(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")

	// for testing
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("../../configs")
	viper.AddConfigPath("../../../configs")
	viper.AddConfigPath("../../../../configs")
	viper.AddConfigPath("../../../../../configs")
	viper.AddConfigPath("../../../../../../configs")
	viper.WatchConfig()

	viper.SetDefault("Database.Pgsql.Host", "127.0.0.1")
	viper.SetDefault("Database.Pgsql.Port", 5432)
	viper.SetDefault("Database.Pgsql.Database", "postgres")
	viper.SetDefault("Database.Pgsql.Schema", "public")
	viper.SetDefault("Database.Pgsql.Username", "postgres")
	viper.SetDefault("Database.Redis.Host", "127.0.0.1")
	viper.SetDefault("Database.Redis.Port", 6379)
	viper.SetDefault("Database.Redis.Prefix", "app_")
	viper.SetDefault("Database.Redis.Lifetime", 600)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	config := Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Config loaded successfully.")
	return &config
}
