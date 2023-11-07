package config

import (
	"math/big"
	"os"

	sdkmath "cosmossdk.io/math"
	"gopkg.in/yaml.v3"

	"github.com/ojo-network/ethereum-api/abi"
	"github.com/ojo-network/indexer/indexer"
	"github.com/ojo-network/indexer/utils"
)

type Config struct {
	NodeUrl string `yaml:"node_url"`
	Pools   []Pool `yaml:"pools"`
}

type Pool struct {
	Address      string `yaml:"address"`
	ExchangePair string `yaml:"exchange_pair"`
	BaseDecimal  uint64 `yaml:"base_decimal"`
	QuoteDecimal uint64 `yaml:"quote_decimal"`
	InvertPrice  bool   `yaml:"invert_price"`
}

func ParseConfig() (*Config, error) {
	config := Config{}

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (p *Pool) ConvertEventToSpotPrice(event *abi.PoolSwap) indexer.SpotPrice {
	return indexer.SpotPrice{
		BlockNum:     indexer.BlockNum(event.Raw.BlockNumber),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair,
		Price:        p.SqrtPriceX96ToDec(event.SqrtPriceX96),
	}
}

func (p *Pool) ConvertEventToSwap(event *abi.PoolSwap) indexer.Swap {
	return indexer.Swap{
		BlockNum:     indexer.BlockNum(event.Raw.BlockNumber),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair,
		Price:        p.SqrtPriceX96ToDec(event.SqrtPriceX96),
		Volume:       p.swapVolume(event),
	}
}

func (p *Pool) SqrtPriceX96ToDec(sqrtPriceX96 *big.Int) sdkmath.LegacyDec {
	sdkValue := sdkmath.LegacyNewDecFromBigInt(sqrtPriceX96)

	// Convert from Q notation to the “actual value” by dividing by 2^k where
	// k is the value after the X. Convert sqrtPriceX96 to sqrtPrice by dividing by 2^96
	sqrtPrice := sdkValue.Quo(sdkmath.LegacyNewDec(2).Power(uint64(96)))
	price := sqrtPrice.Power(2)

	// ERC-20 tokens have built in decimal values. For example, 1 WETH actually
	// represents WETH in the contract whereas USDC is 10^6
	// Therefore, USDC has 6 decimals and WETH has 18.
	// To get the adjusted price, divide the result by (10^18 / 10^6)
	basePower := sdkmath.LegacyNewDec(10).Power(p.BaseDecimal)
	quotePower := sdkmath.LegacyNewDec(10).Power(p.QuoteDecimal)
	adjPower := quotePower.Quo(basePower)
	price = price.Quo(adjPower)

	// Divide 1 by the price if the desired price is inverted from the pool
	if p.InvertPrice {
		return sdkmath.LegacyNewDec(1).Quo(price)
	} else {
		return price
	}
}

func (p *Pool) swapVolume(event *abi.PoolSwap) sdkmath.LegacyDec {
	if p.InvertPrice {
		volume := sdkmath.LegacyNewDecFromBigInt(event.Amount1).Abs()
		return volume.Quo(sdkmath.LegacyNewDec(10).Power(uint64(p.QuoteDecimal)))
	} else {
		volume := sdkmath.LegacyNewDecFromBigInt(event.Amount0).Abs()
		return volume.Quo(sdkmath.LegacyNewDec(10).Power(uint64(p.BaseDecimal)))
	}
}
