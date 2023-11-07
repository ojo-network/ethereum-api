package client

import (
	"math/big"

	sdkmath "cosmossdk.io/math"
	"github.com/ojo-network/ethereum-api/abi"
	"github.com/ojo-network/indexer/indexer"
	"github.com/ojo-network/indexer/utils"
)

type Pool struct {
	Address      string
	ExchangePair string
	BaseDecimal  uint64
	QuoteDecimal uint64
	InvertPrice  bool
}

func (p *Pool) convertEventToSpotPrice(event *abi.PoolSwap) indexer.SpotPrice {
	return indexer.SpotPrice{
		BlockNum:     indexer.BlockNum(event.Raw.BlockNumber),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair,
		Price:        p.sqrtPriceX96ToDec(event.SqrtPriceX96),
	}
}

func (p *Pool) convertEventToSwap(event *abi.PoolSwap) indexer.Swap {
	return indexer.Swap{
		BlockNum:     indexer.BlockNum(event.Raw.BlockNumber),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair,
		Price:        p.sqrtPriceX96ToDec(event.SqrtPriceX96),
		Volume:       p.swapVolume(event),
	}
}

func (p *Pool) sqrtPriceX96ToDec(sqrtPriceX96 *big.Int) sdkmath.LegacyDec {
	sdkValue := sdkmath.LegacyNewDecFromBigInt(sqrtPriceX96)

	// Convert from Q notation to the “actual value” by dividing by 2^k where
	// k is the value after the X. Convert sqrtPriceX96 to sqrtPrice by dividing by 2^96
	sqrtPrice := sdkValue.Quo(sdkmath.LegacyNewDec(2).Power(uint64(96)))
	price := sqrtPrice.Power(2)

	// erc20 tokens have built in decimal values. For example, 1 WETH actually
	// represents WETH in the contract whereas USDC is 10^6
	// Therefore, USDC has 6 decimals and WETH has 18.
	// To get the adjusted price, divide the result by (10^18 / 10^6)
	adjPower := p.BaseDecimal - p.QuoteDecimal
	adjPrice := price.Quo(sdkmath.LegacyNewDec(10).Power(uint64(adjPower)))

	// Divide 1 by the price if the desired price is inverted from the pool
	if p.InvertPrice {
		return sdkmath.LegacyNewDec(1).Quo(adjPrice)
	} else {
		return adjPrice
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
