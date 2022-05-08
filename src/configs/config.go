package configs

import "github.com/spf13/viper"

type ServerConfig struct {
	ConnectionString string `mapstructure:"DB_DSN"`
	JWTsecret        string `mapstructure:"JWT_SECRET"`
	// ServerPort       string `mapstructure:"SERVER_PORT"`
	// ServerHost       string `mapstructure:"SERVER_HOST"`
}

func LoadServerConfig(path string) (ServerConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	config := ServerConfig{}
	err := viper.ReadInConfig()
	viper.Unmarshal(&config)
	return config, err
}
