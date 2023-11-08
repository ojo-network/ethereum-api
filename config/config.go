package config

import (
	"os"

	"github.com/ojo-network/ethereum-api/client"
	"github.com/ojo-network/indexer/server"
	"gopkg.in/yaml.v3"
)

type Config struct {
	NodeUrl string              `yaml:"node_url"`
	Server  server.ServerConfig `yaml:"server"`
	Pools   []client.Pool       `yaml:"pools"`
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
	return &config, nil
}

func (c *Config) AssetPairs() []server.AssetPair {
	var assetPairs []server.AssetPair
	for _, pool := range c.Pools {
		assetPairs = append(assetPairs, server.AssetPair{
			Base:  pool.Base,
			Quote: pool.Quote,
		})
	}
	return assetPairs
}
