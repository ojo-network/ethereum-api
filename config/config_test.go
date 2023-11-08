package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(`
---
node_url: "http://localhost:8545"
server:
  listen_addr: "http://localhost:8080"
pools:
  - exchange_pair: "WBTC/WETH"
    address: "testAddress1"
  - exchange_pair: "WETH/USDC"
    address: "testAddress2"
`))
	if err != nil {
		t.Fatal(err)
	}

	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	config, err := ParseConfig(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the parsed config matches the expected values
	assert.Equal(t, "http://localhost:8545", config.NodeUrl)
	assert.Equal(t, "http://localhost:8080", config.Server.ListenAddr)
	assert.Equal(t, 2, len(config.Pools))
	assert.Equal(t, "WBTC/WETH", config.Pools[0].ExchangePair)
	assert.Equal(t, "testAddress1", config.Pools[0].Address)
	assert.Equal(t, "WETH/USDC", config.Pools[1].ExchangePair)
	assert.Equal(t, "testAddress2", config.Pools[1].Address)
}
