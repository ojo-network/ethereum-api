package config

import (
	"math/rand"
	"os"

	"github.com/ojo-network/ethereum-api/pool"
	"github.com/ojo-network/indexer/server"
	"gopkg.in/yaml.v3"
)

type Config struct {
	NodeUrls []string            `yaml:"node_urls"`
	Server   server.ServerConfig `yaml:"server"`
	Pools    []pool.Pool         `yaml:"pools"`
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

func (c *Config) RandomNodeUrl() string {
	index := rand.Intn(len(c.NodeUrls))
	return c.NodeUrls[index]
}
