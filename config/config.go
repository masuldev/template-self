package config

import "github.com/spf13/viper"

var environment *Env

type Env struct {
	Application Application
}

type Application struct {
	Version string
	Port    string
}

func Init() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	env := &Env{
		Application: Application{
			Port:    viper.GetString("PORT"),
			Version: viper.GetString("VERSION"),
		},
	}

	environment = env
	return nil
}

func GetAll() *Env {
	return environment
}
