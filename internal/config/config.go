package config

import (
    "github.com/spf13/viper"
	"github.com/akramov1ch/bookstore/internal/models"
)


func LoadConfig() (*models.Config, error) {
    viper.SetConfigFile(".env")
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }
    config := &models.Config{
        DBHost:     viper.GetString("DB_HOST"),
        DBPort:     viper.GetString("DB_PORT"),
        DBUser:     viper.GetString("DB_USER"),
        DBPassword: viper.GetString("DB_PASSWORD"),
        DBName:     viper.GetString("DB_NAME"),
    }
    return config, nil
}
