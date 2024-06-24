package pool

type ExchangeName string
type PoolContract string

const ExchangeUniswap ExchangeName = "uniswap"
const ExchangeCamelot ExchangeName = "camelot"
const ExchangeBalancer ExchangeName = "balancer"
const PoolUniswap PoolContract = "uniswappool"
const PoolAlgebra PoolContract = "algebrapool"
const PoolBalancer PoolContract = "balancerpool"

// maps exchange to pool contract of that exchange
var SupportedExchanges = map[ExchangeName]PoolContract{
	ExchangeUniswap:  PoolUniswap,
	ExchangeCamelot:  PoolAlgebra,
	ExchangeBalancer: PoolBalancer,
}
