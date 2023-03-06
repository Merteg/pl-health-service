package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Environment string
	Mongo       MongoConfiguration
}

type MongoConfiguration struct {
	mongoURI        string
	port            string
	dbName          string
	targetsCollName string
	healthCollName  string
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

	fmt.Println(conf)
	return conf

}
