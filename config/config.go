package config

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/ojo-network/ethereum-api/pool"
	"github.com/ojo-network/indexer/server"
	"github.com/spf13/viper"
)

const (
	configFileType    = "yaml"
	envVariablePrefix = "eth_api"
	structTagName     = "yaml"
)

type Exchange struct {
	Name     pool.ExchangeName `yaml:"name"`
	NodeUrls []string          `yaml:"node_urls"`
	Pools    []pool.Pool       `yaml:"pools"`
}

type Config struct {
	Exchanges []Exchange          `yaml:"exchanges"`
	Server    server.ServerConfig `yaml:"server"`
}

func ParseConfig(filePath string) (*Config, error) {
	config := Config{}

	vpr := viper.New()
	vpr.SetConfigFile(filePath)
	vpr.SetConfigType(configFileType)
	vpr.SetEnvPrefix(envVariablePrefix)
	vpr.AutomaticEnv()
	vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := vpr.ReadInConfig(); err != nil {
		return &config, fmt.Errorf("couldn't load config: %s", err)
	}

	tagConfig := func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = structTagName
	}

	if err := vpr.Unmarshal(&config, tagConfig); err != nil {
		return &config, fmt.Errorf("couldn't read config: %s", err)
	}

	// Split NodeUrls from environment variables
	for i := range config.Exchanges {
		envVarName := fmt.Sprintf("exchanges.%d.node_urls", i)
		if v := vpr.GetString(envVarName); v != "" {
			config.Exchanges[i].NodeUrls = strings.Split(v, ",")
		}
	}

	return &config, config.Validate()
}

// Validate returns an error if the Config object is invalid.
func (c Config) Validate() (err error) {
	for _, exchange := range c.Exchanges {
		if _, ok := pool.SupportedExchanges[exchange.Name]; !ok {
			return fmt.Errorf("unsupported exchange: %s", exchange.Name)
		}
	}

	return nil
}

func (c *Config) AssetPairs(exchange Exchange) []server.AssetPair {
	var assetPairs []server.AssetPair
	for _, pool := range exchange.Pools {
		assetPairs = append(assetPairs, server.AssetPair{
			Base:  pool.Base,
			Quote: pool.Quote,
		})
	}

	return assetPairs
}
