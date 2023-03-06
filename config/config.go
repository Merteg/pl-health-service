package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Environment string
	Mongo       map[string]string
}

func GetConfig() Configuration {
	conf := Configuration{}
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = v.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}
