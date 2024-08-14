package pool

type ExchangeName string
type PoolContract string

const ExchangeUniswap ExchangeName = "uniswap"
const ExchangeCamelot ExchangeName = "camelot"
const ExchangeBalancer ExchangeName = "balancer"
const ExchangePancake ExchangeName = "pancake"
const PoolUniswap PoolContract = "uniswappool"
const PoolAlgebra PoolContract = "algebrapool"
const PoolBalancer PoolContract = "balancerpool"
const PoolPancake PoolContract = "pancakepool"

// maps exchange to pool contract of that exchange
var SupportedExchanges = map[ExchangeName]PoolContract{
	ExchangeUniswap:  PoolUniswap,
	ExchangeCamelot:  PoolAlgebra,
	ExchangeBalancer: PoolBalancer,
	ExchangePancake:  PoolPancake,
}
