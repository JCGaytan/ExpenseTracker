package config


import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBPath    string
	JWTSecret string
}

var AppConfig Config

func LoadConfig() {
	viper.SetDefault("DB_PATH", "expense_tracker.db")
	viper.SetDefault("JWT_SECRET", "your_secret_key")
	viper.AutomaticEnv()

	AppConfig = Config{
		DBPath:    viper.GetString("DB_PATH"),
		JWTSecret: viper.GetString("JWT_SECRET"),
	}

	if AppConfig.JWTSecret == "your_secret_key" {
		log.Println("[WARN] Using default JWT secret. Set JWT_SECRET env variable in production.")
	}
}
