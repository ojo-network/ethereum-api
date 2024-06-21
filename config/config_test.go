package config

import (
	"os"
	"testing"

	"github.com/ojo-network/ethereum-api/pool"
	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(`
exchanges:
  - name: uniswap
    node_urls:
      - http://node1.com
    pools:
      - base: "WBTC"
        quote: "WETH"
        address: "testAddress1"
  - name: camelot
    node_urls:
      - http://node2.com
    pools:
      - base: "WETH"
        quote: "USDC"
        address: "testAddress2"
server:
  listen_addr: "http://localhost:8080"
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
	assert.Equal(t, pool.ExchangeUniswap, config.Exchanges[0].Name)
	assert.Equal(t, pool.ExchangeCamelot, config.Exchanges[1].Name)

	assert.Equal(t, "http://node1.com", config.Exchanges[0].NodeUrls[0])
	assert.Equal(t, "http://node2.com", config.Exchanges[1].NodeUrls[0])

	assert.Equal(t, 1, len(config.Exchanges[0].Pools))
	assert.Equal(t, 1, len(config.Exchanges[1].Pools))

	assert.Equal(t, "WBTC/WETH", config.Exchanges[0].Pools[0].ExchangePair())
	assert.Equal(t, "testAddress1", config.Exchanges[0].Pools[0].Address)

	assert.Equal(t, "WETH/USDC", config.Exchanges[1].Pools[0].ExchangePair())
	assert.Equal(t, "testAddress2", config.Exchanges[1].Pools[0].Address)

	assert.Equal(t, "http://localhost:8080", config.Server.ListenAddr)
}

func TestInvalidExchanges(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(`
exchanges:
  - name: invalid exchange
    node_urls:
      - http://node1.com
    pools:
      - base: "WBTC"
        quote: "WETH"
        address: "testAddress1"
  - name: camelot
    node_urls:
      - http://node2.com
    pools:
      - base: "WETH"
        quote: "USDC"
        address: "testAddress2"
server:
  listen_addr: "http://localhost:8080"
`))
	if err != nil {
		t.Fatal(err)
	}

	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	_, err = ParseConfig(tmpfile.Name())
	assert.ErrorContains(t, err, "unsupported exchange: invalid exchange")
}
