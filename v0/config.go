package v0

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config stores configuration read from config.yml
type Config struct {
	RPCEndpoint       string   `mapstructure:"rpcEndpoint"`
	ContractAddresses []string `mapstructure:"contractAddresses"`
	LogLevel          string   `mapstructure:"logLevel"`
}

// ConfigLoader defines an interface for loading and checking configuration
type ConfigLoader interface {
	LoadConfig() (*Config, error)
	CheckConfig(*Config) error
}

// ViperConfigLoader implements ConfigLoader using Viper
type ViperConfigLoader struct{}

func (v *ViperConfigLoader) LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (v *ViperConfigLoader) CheckConfig(config *Config) error {
	for _, address := range config.ContractAddresses {
		if !common.IsHexAddress(address) {
			return errors.New("invalid Ethereum address found in config: " + address)
		}
	}
	return nil
}

// InitializeLogger initializes a Zap logger based on the log level in the config
func InitializeLogger(logLevel string) (*zap.SugaredLogger, error) {
	var level zapcore.Level
	if err := level.Set(logLevel); err != nil {
		return nil, err
	}

	config := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
