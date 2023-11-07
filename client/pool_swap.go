package client

import (
	"math/big"

	sdkmath "cosmossdk.io/math"
	"github.com/ojo-network/ethereum-api/abi"
	"github.com/ojo-network/indexer/indexer"
	"github.com/ojo-network/indexer/utils"
)

func convertEventToSwap(event *abi.PoolSwap) indexer.Swap {
	amount1 := sdkmath.LegacyNewDecFromBigInt(event.Amount1)
	amount1Abs := amount1.Abs()
	// TODO: Retrieve the denom decimals from the contract automatically

	return indexer.Swap{
		BlockNum:     indexer.BlockNum(event.Raw.BlockNumber),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: "WETH/USDC",
		Price:        sqrtPriceX96ToDec(event.SqrtPriceX96),
		Volume:       amount1Abs.Quo(sdkmath.LegacyNewDec(10).Power(uint64(18))),
	}
}

// Convert from Q notation to the “actual value” by dividing by 2^k where
// k is the value after the X. For example, you can convert sqrtPriceX96
// to sqrtPrice by dividing by 2^96
func sqrtPriceX96ToDec(sqrtPriceX96 *big.Int) sdkmath.LegacyDec {
	sdkValue := sdkmath.LegacyNewDecFromBigInt(sqrtPriceX96)
	sqrtPrice := sdkValue.Quo(sdkmath.LegacyNewDec(2).Power(uint64(96)))
	price := sqrtPrice.Power(2)

	// 	erc20 tokens have built in decimal values. For example, 1 WETH actually
	// represents WETH in the contract whereas USDC is 10^6
	// Therefore, USDC has 6 decimals and WETH has 18.
	// to get the adjusted price divide the result by (10^18 / 10^6)
	adjPrice := price.Quo(sdkmath.LegacyNewDec(10).Power(uint64(12)))

	// Divide 1 by the price to get the price of WETH/USDC
	return sdkmath.LegacyNewDec(1).Quo(adjPrice)
}
