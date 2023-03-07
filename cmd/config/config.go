package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// Config represents the application configuration.
type Config struct {
	DB  *DBConfig
	App *AppConfig
}

// DBConfig represents the database configuration.
type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

// AppConfig represents the application configuration.
type AppConfig struct {
	Port string
	Env  string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("./cmd/config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}

	return &cfg
}

func (cfg *Config) DSN() string {
	return cfg.DB.Username + ":" + cfg.DB.Password + "@tcp(" + cfg.DB.Host + ":" + cfg.DB.Port + ")/" + cfg.DB.Database
}

// WatchConfig watches the config file for changes and reloads it.
func WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file %s changed, reloading...\n", e.Name)
	})
}
