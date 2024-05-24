package config

import (
	"fmt"
	"os"

	"github.com/ojo-network/ethereum-api/pool"
	"github.com/ojo-network/indexer/server"
	"gopkg.in/yaml.v3"
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

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, config.Validate()
}

// Validate returns an error if the Config object is invalid.
func (c Config) Validate() (err error) {
	for _, exchange := range c.Exchanges{
		if _, ok := pool.SupportedExchanges[exchange.Name]; !ok {
			return fmt.Errorf("unsupported exchange: %s", exchange.Name)
		}
	}

	return nil
}

func (c *Config) AssetPairs() []server.AssetPair {
	var assetPairs []server.AssetPair
	for _, exchange := range c.Exchanges {
		for _, pool := range exchange.Pools {
			assetPairs = append(assetPairs, server.AssetPair{
				Base:  pool.Base,
				Quote: pool.Quote,
			})
		}
	}

	return assetPairs
}
