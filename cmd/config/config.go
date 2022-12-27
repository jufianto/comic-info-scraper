package config

import "github.com/spf13/viper"

func GetConfig() (*viper.Viper, error) {
	config := viper.New()
	viper.SetConfigType("yaml")
	config.SetConfigFile("./config/scraper.yaml")

	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}

	return config, nil
}
