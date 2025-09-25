package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
		App      AppConfig
	}

	ServerConfig struct {
		Port string
		Host string
	}

	DatabaseConfig struct {
		URL      string
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		SSLMode  string
	}

	AppConfig struct {
		Environment string
		Debug       bool
	}
)

func Load() *Config {
	// Set config file settings
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Set default values
	setDefaults()

	// Enable automatic env variable reading
	viper.AutomaticEnv()

	// Read config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	// Mapping config
	config := &Config{
		Server: ServerConfig{
			Port: viper.GetString("server.port"),
			Host: viper.GetString("server.host"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			Name:     viper.GetString("database.name"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			SSLMode:  viper.GetString("database.sslmode"),
		},
		App: AppConfig{
			Environment: viper.GetString("app.environment"),
			Debug:       viper.GetBool("app.debug"),
		},
	}

	// Build database URL
	config.Database.URL = buildDatabaseURL(&config.Database)

	fmt.Println(*config)
	return config
}

func setDefaults() {
	// Server defaults
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.host", "0.0.0.0")

	// Database defaults
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.name", "checkout_db")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.sslmode", "disable")

	// App defaults
	viper.SetDefault("app.environment", "development")
	viper.SetDefault("app.debug", true)
}

func buildDatabaseURL(db *DatabaseConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		db.User, db.Password, db.Host, db.Port, db.Name, db.SSLMode)
}

// Helper methods
func (c *Config) IsProduction() bool {
	return c.App.Environment == "production"
}

func (c *Config) IsDevelopment() bool {
	return c.App.Environment == "development"
}

func (c *Config) GetServerAddress() string {
	return c.Server.Host + ":" + c.Server.Port
}
