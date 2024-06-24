// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package camelot

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AlgebraPoolMetaData contains all meta data concerning the AlgebraPool contract.
var AlgebraPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"communityFee0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"communityFee1New\",\"type\":\"uint8\"}],\"name\":\"CommunityFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeZto\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeOtz\",\"type\":\"uint16\"}],\"name\":\"Fee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"virtualPoolAddress\",\"type\":\"address\"}],\"name\":\"Incentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"liquidityCooldown\",\"type\":\"uint32\"}],\"name\":\"LiquidityCooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"newTickSpacing\",\"type\":\"int24\"}],\"name\":\"TickSpacing\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activeIncentive\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataStorageOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"}],\"name\":\"getInnerCumulatives\",\"outputs\":[{\"internalType\":\"int56\",\"name\":\"innerTickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"innerSecondsSpentPerLiquidity\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"innerSecondsSpent\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"getTimepoints\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"tickCumulatives\",\"type\":\"int56[]\"},{\"internalType\":\"uint160[]\",\"name\":\"secondsPerLiquidityCumulatives\",\"type\":\"uint160[]\"},{\"internalType\":\"uint112[]\",\"name\":\"volatilityCumulatives\",\"type\":\"uint112[]\"},{\"internalType\":\"uint256[]\",\"name\":\"volumePerAvgLiquiditys\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalState\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"feeZto\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeOtz\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timepointIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"communityFeeToken0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"communityFeeToken1\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"initialPrice\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityCooldown\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidityPerTick\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidityDesired\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityActual\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastLiquidityAddTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"innerFeeGrowth0Token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"innerFeeGrowth1Token\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"fees0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fees1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"communityFee0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"communityFee1\",\"type\":\"uint8\"}],\"name\":\"setCommunityFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"virtualPoolAddress\",\"type\":\"address\"}],\"name\":\"setIncentive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newLiquidityCooldown\",\"type\":\"uint32\"}],\"name\":\"setLiquidityCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newTickSpacing\",\"type\":\"int24\"}],\"name\":\"setTickSpacing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountRequired\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountRequired\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapSupportingFeeOnInputTokens\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"name\":\"tickTable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityTotal\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"outerFeeGrowth0Token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outerFeeGrowth1Token\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"outerTickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"outerSecondsPerLiquidity\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"outerSecondsSpent\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"timepoints\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulative\",\"type\":\"uint160\"},{\"internalType\":\"uint88\",\"name\":\"volatilityCumulative\",\"type\":\"uint88\"},{\"internalType\":\"int24\",\"name\":\"averageTick\",\"type\":\"int24\"},{\"internalType\":\"uint144\",\"name\":\"volumePerLiquidityCumulative\",\"type\":\"uint144\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeGrowth0Token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeGrowth1Token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AlgebraPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AlgebraPoolMetaData.ABI instead.
var AlgebraPoolABI = AlgebraPoolMetaData.ABI

// AlgebraPool is an auto generated Go binding around an Ethereum contract.
type AlgebraPool struct {
	AlgebraPoolCaller     // Read-only binding to the contract
	AlgebraPoolTransactor // Write-only binding to the contract
	AlgebraPoolFilterer   // Log filterer for contract events
}

// AlgebraPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlgebraPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlgebraPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlgebraPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlgebraPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlgebraPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlgebraPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlgebraPoolSession struct {
	Contract     *AlgebraPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlgebraPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlgebraPoolCallerSession struct {
	Contract *AlgebraPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AlgebraPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlgebraPoolTransactorSession struct {
	Contract     *AlgebraPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AlgebraPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlgebraPoolRaw struct {
	Contract *AlgebraPool // Generic contract binding to access the raw methods on
}

// AlgebraPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlgebraPoolCallerRaw struct {
	Contract *AlgebraPoolCaller // Generic read-only contract binding to access the raw methods on
}

// AlgebraPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlgebraPoolTransactorRaw struct {
	Contract *AlgebraPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlgebraPool creates a new instance of AlgebraPool, bound to a specific deployed contract.
func NewAlgebraPool(address common.Address, backend bind.ContractBackend) (*AlgebraPool, error) {
	contract, err := bindAlgebraPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AlgebraPool{AlgebraPoolCaller: AlgebraPoolCaller{contract: contract}, AlgebraPoolTransactor: AlgebraPoolTransactor{contract: contract}, AlgebraPoolFilterer: AlgebraPoolFilterer{contract: contract}}, nil
}

// NewAlgebraPoolCaller creates a new read-only instance of AlgebraPool, bound to a specific deployed contract.
func NewAlgebraPoolCaller(address common.Address, caller bind.ContractCaller) (*AlgebraPoolCaller, error) {
	contract, err := bindAlgebraPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolCaller{contract: contract}, nil
}

// NewAlgebraPoolTransactor creates a new write-only instance of AlgebraPool, bound to a specific deployed contract.
func NewAlgebraPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AlgebraPoolTransactor, error) {
	contract, err := bindAlgebraPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolTransactor{contract: contract}, nil
}

// NewAlgebraPoolFilterer creates a new log filterer instance of AlgebraPool, bound to a specific deployed contract.
func NewAlgebraPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AlgebraPoolFilterer, error) {
	contract, err := bindAlgebraPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolFilterer{contract: contract}, nil
}

// bindAlgebraPool binds a generic wrapper to an already deployed contract.
func bindAlgebraPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AlgebraPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlgebraPool *AlgebraPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlgebraPool.Contract.AlgebraPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlgebraPool *AlgebraPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlgebraPool.Contract.AlgebraPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlgebraPool *AlgebraPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlgebraPool.Contract.AlgebraPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlgebraPool *AlgebraPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlgebraPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlgebraPool *AlgebraPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlgebraPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlgebraPool *AlgebraPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlgebraPool.Contract.contract.Transact(opts, method, params...)
}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address)
func (_AlgebraPool *AlgebraPoolCaller) ActiveIncentive(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "activeIncentive")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address)
func (_AlgebraPool *AlgebraPoolSession) ActiveIncentive() (common.Address, error) {
	return _AlgebraPool.Contract.ActiveIncentive(&_AlgebraPool.CallOpts)
}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address)
func (_AlgebraPool *AlgebraPoolCallerSession) ActiveIncentive() (common.Address, error) {
	return _AlgebraPool.Contract.ActiveIncentive(&_AlgebraPool.CallOpts)
}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_AlgebraPool *AlgebraPoolCaller) DataStorageOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "dataStorageOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_AlgebraPool *AlgebraPoolSession) DataStorageOperator() (common.Address, error) {
	return _AlgebraPool.Contract.DataStorageOperator(&_AlgebraPool.CallOpts)
}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_AlgebraPool *AlgebraPoolCallerSession) DataStorageOperator() (common.Address, error) {
	return _AlgebraPool.Contract.DataStorageOperator(&_AlgebraPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AlgebraPool *AlgebraPoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AlgebraPool *AlgebraPoolSession) Factory() (common.Address, error) {
	return _AlgebraPool.Contract.Factory(&_AlgebraPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AlgebraPool *AlgebraPoolCallerSession) Factory() (common.Address, error) {
	return _AlgebraPool.Contract.Factory(&_AlgebraPool.CallOpts)
}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_AlgebraPool *AlgebraPoolCaller) GetInnerCumulatives(opts *bind.CallOpts, bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "getInnerCumulatives", bottomTick, topTick)

	outstruct := new(struct {
		InnerTickCumulative           *big.Int
		InnerSecondsSpentPerLiquidity *big.Int
		InnerSecondsSpent             uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InnerTickCumulative = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.InnerSecondsSpentPerLiquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InnerSecondsSpent = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_AlgebraPool *AlgebraPoolSession) GetInnerCumulatives(bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	return _AlgebraPool.Contract.GetInnerCumulatives(&_AlgebraPool.CallOpts, bottomTick, topTick)
}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_AlgebraPool *AlgebraPoolCallerSession) GetInnerCumulatives(bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	return _AlgebraPool.Contract.GetInnerCumulatives(&_AlgebraPool.CallOpts, bottomTick, topTick)
}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_AlgebraPool *AlgebraPoolCaller) GetTimepoints(opts *bind.CallOpts, secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "getTimepoints", secondsAgos)

	outstruct := new(struct {
		TickCumulatives                []*big.Int
		SecondsPerLiquidityCumulatives []*big.Int
		VolatilityCumulatives          []*big.Int
		VolumePerAvgLiquiditys         []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulatives = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerLiquidityCumulatives = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.VolatilityCumulatives = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.VolumePerAvgLiquiditys = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_AlgebraPool *AlgebraPoolSession) GetTimepoints(secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	return _AlgebraPool.Contract.GetTimepoints(&_AlgebraPool.CallOpts, secondsAgos)
}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_AlgebraPool *AlgebraPoolCallerSession) GetTimepoints(secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	return _AlgebraPool.Contract.GetTimepoints(&_AlgebraPool.CallOpts, secondsAgos)
}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 feeZto, uint16 feeOtz, uint16 timepointIndex, uint8 communityFeeToken0, uint8 communityFeeToken1, bool unlocked)
func (_AlgebraPool *AlgebraPoolCaller) GlobalState(opts *bind.CallOpts) (struct {
	Price              *big.Int
	Tick               *big.Int
	FeeZto             uint16
	FeeOtz             uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "globalState")

	outstruct := new(struct {
		Price              *big.Int
		Tick               *big.Int
		FeeZto             uint16
		FeeOtz             uint16
		TimepointIndex     uint16
		CommunityFeeToken0 uint8
		CommunityFeeToken1 uint8
		Unlocked           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeZto = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.FeeOtz = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.TimepointIndex = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.CommunityFeeToken0 = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.CommunityFeeToken1 = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 feeZto, uint16 feeOtz, uint16 timepointIndex, uint8 communityFeeToken0, uint8 communityFeeToken1, bool unlocked)
func (_AlgebraPool *AlgebraPoolSession) GlobalState() (struct {
	Price              *big.Int
	Tick               *big.Int
	FeeZto             uint16
	FeeOtz             uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}, error) {
	return _AlgebraPool.Contract.GlobalState(&_AlgebraPool.CallOpts)
}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 feeZto, uint16 feeOtz, uint16 timepointIndex, uint8 communityFeeToken0, uint8 communityFeeToken1, bool unlocked)
func (_AlgebraPool *AlgebraPoolCallerSession) GlobalState() (struct {
	Price              *big.Int
	Tick               *big.Int
	FeeZto             uint16
	FeeOtz             uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}, error) {
	return _AlgebraPool.Contract.GlobalState(&_AlgebraPool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_AlgebraPool *AlgebraPoolCaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_AlgebraPool *AlgebraPoolSession) Liquidity() (*big.Int, error) {
	return _AlgebraPool.Contract.Liquidity(&_AlgebraPool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_AlgebraPool *AlgebraPoolCallerSession) Liquidity() (*big.Int, error) {
	return _AlgebraPool.Contract.Liquidity(&_AlgebraPool.CallOpts)
}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32)
func (_AlgebraPool *AlgebraPoolCaller) LiquidityCooldown(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "liquidityCooldown")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32)
func (_AlgebraPool *AlgebraPoolSession) LiquidityCooldown() (uint32, error) {
	return _AlgebraPool.Contract.LiquidityCooldown(&_AlgebraPool.CallOpts)
}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32)
func (_AlgebraPool *AlgebraPoolCallerSession) LiquidityCooldown() (uint32, error) {
	return _AlgebraPool.Contract.LiquidityCooldown(&_AlgebraPool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() pure returns(uint128)
func (_AlgebraPool *AlgebraPoolCaller) MaxLiquidityPerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "maxLiquidityPerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() pure returns(uint128)
func (_AlgebraPool *AlgebraPoolSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _AlgebraPool.Contract.MaxLiquidityPerTick(&_AlgebraPool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() pure returns(uint128)
func (_AlgebraPool *AlgebraPoolCallerSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _AlgebraPool.Contract.MaxLiquidityPerTick(&_AlgebraPool.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_AlgebraPool *AlgebraPoolCaller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Liquidity                 *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		Liquidity                 *big.Int
		LastLiquidityAddTimestamp uint32
		InnerFeeGrowth0Token      *big.Int
		InnerFeeGrowth1Token      *big.Int
		Fees0                     *big.Int
		Fees1                     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastLiquidityAddTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.InnerFeeGrowth0Token = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InnerFeeGrowth1Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Fees0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fees1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_AlgebraPool *AlgebraPoolSession) Positions(arg0 [32]byte) (struct {
	Liquidity                 *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	return _AlgebraPool.Contract.Positions(&_AlgebraPool.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_AlgebraPool *AlgebraPoolCallerSession) Positions(arg0 [32]byte) (struct {
	Liquidity                 *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	return _AlgebraPool.Contract.Positions(&_AlgebraPool.CallOpts, arg0)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AlgebraPool *AlgebraPoolCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AlgebraPool *AlgebraPoolSession) TickSpacing() (*big.Int, error) {
	return _AlgebraPool.Contract.TickSpacing(&_AlgebraPool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AlgebraPool *AlgebraPoolCallerSession) TickSpacing() (*big.Int, error) {
	return _AlgebraPool.Contract.TickSpacing(&_AlgebraPool.CallOpts)
}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 ) view returns(uint256)
func (_AlgebraPool *AlgebraPoolCaller) TickTable(opts *bind.CallOpts, arg0 int16) (*big.Int, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "tickTable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 ) view returns(uint256)
func (_AlgebraPool *AlgebraPoolSession) TickTable(arg0 int16) (*big.Int, error) {
	return _AlgebraPool.Contract.TickTable(&_AlgebraPool.CallOpts, arg0)
}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 ) view returns(uint256)
func (_AlgebraPool *AlgebraPoolCallerSession) TickTable(arg0 int16) (*big.Int, error) {
	return _AlgebraPool.Contract.TickTable(&_AlgebraPool.CallOpts, arg0)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_AlgebraPool *AlgebraPoolCaller) Ticks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "ticks", arg0)

	outstruct := new(struct {
		LiquidityTotal           *big.Int
		LiquidityDelta           *big.Int
		OuterFeeGrowth0Token     *big.Int
		OuterFeeGrowth1Token     *big.Int
		OuterTickCumulative      *big.Int
		OuterSecondsPerLiquidity *big.Int
		OuterSecondsSpent        uint32
		Initialized              bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityTotal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityDelta = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OuterFeeGrowth0Token = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OuterFeeGrowth1Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OuterTickCumulative = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OuterSecondsPerLiquidity = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OuterSecondsSpent = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.Initialized = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_AlgebraPool *AlgebraPoolSession) Ticks(arg0 *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	return _AlgebraPool.Contract.Ticks(&_AlgebraPool.CallOpts, arg0)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_AlgebraPool *AlgebraPoolCallerSession) Ticks(arg0 *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	return _AlgebraPool.Contract.Ticks(&_AlgebraPool.CallOpts, arg0)
}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_AlgebraPool *AlgebraPoolCaller) Timepoints(opts *bind.CallOpts, index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "timepoints", index)

	outstruct := new(struct {
		Initialized                   bool
		BlockTimestamp                uint32
		TickCumulative                *big.Int
		SecondsPerLiquidityCumulative *big.Int
		VolatilityCumulative          *big.Int
		AverageTick                   *big.Int
		VolumePerLiquidityCumulative  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.TickCumulative = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityCumulative = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VolatilityCumulative = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AverageTick = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.VolumePerLiquidityCumulative = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_AlgebraPool *AlgebraPoolSession) Timepoints(index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	return _AlgebraPool.Contract.Timepoints(&_AlgebraPool.CallOpts, index)
}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_AlgebraPool *AlgebraPoolCallerSession) Timepoints(index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	return _AlgebraPool.Contract.Timepoints(&_AlgebraPool.CallOpts, index)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AlgebraPool *AlgebraPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AlgebraPool *AlgebraPoolSession) Token0() (common.Address, error) {
	return _AlgebraPool.Contract.Token0(&_AlgebraPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AlgebraPool *AlgebraPoolCallerSession) Token0() (common.Address, error) {
	return _AlgebraPool.Contract.Token0(&_AlgebraPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AlgebraPool *AlgebraPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AlgebraPool *AlgebraPoolSession) Token1() (common.Address, error) {
	return _AlgebraPool.Contract.Token1(&_AlgebraPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AlgebraPool *AlgebraPoolCallerSession) Token1() (common.Address, error) {
	return _AlgebraPool.Contract.Token1(&_AlgebraPool.CallOpts)
}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_AlgebraPool *AlgebraPoolCaller) TotalFeeGrowth0Token(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "totalFeeGrowth0Token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_AlgebraPool *AlgebraPoolSession) TotalFeeGrowth0Token() (*big.Int, error) {
	return _AlgebraPool.Contract.TotalFeeGrowth0Token(&_AlgebraPool.CallOpts)
}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_AlgebraPool *AlgebraPoolCallerSession) TotalFeeGrowth0Token() (*big.Int, error) {
	return _AlgebraPool.Contract.TotalFeeGrowth0Token(&_AlgebraPool.CallOpts)
}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_AlgebraPool *AlgebraPoolCaller) TotalFeeGrowth1Token(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlgebraPool.contract.Call(opts, &out, "totalFeeGrowth1Token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_AlgebraPool *AlgebraPoolSession) TotalFeeGrowth1Token() (*big.Int, error) {
	return _AlgebraPool.Contract.TotalFeeGrowth1Token(&_AlgebraPool.CallOpts)
}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_AlgebraPool *AlgebraPoolCallerSession) TotalFeeGrowth1Token() (*big.Int, error) {
	return _AlgebraPool.Contract.TotalFeeGrowth1Token(&_AlgebraPool.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_AlgebraPool *AlgebraPoolTransactor) Burn(opts *bind.TransactOpts, bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "burn", bottomTick, topTick, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_AlgebraPool *AlgebraPoolSession) Burn(bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Burn(&_AlgebraPool.TransactOpts, bottomTick, topTick, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_AlgebraPool *AlgebraPoolTransactorSession) Burn(bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Burn(&_AlgebraPool.TransactOpts, bottomTick, topTick, amount)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_AlgebraPool *AlgebraPoolTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "collect", recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_AlgebraPool *AlgebraPoolSession) Collect(recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Collect(&_AlgebraPool.TransactOpts, recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_AlgebraPool *AlgebraPoolTransactorSession) Collect(recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Collect(&_AlgebraPool.TransactOpts, recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_AlgebraPool *AlgebraPoolTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "flash", recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_AlgebraPool *AlgebraPoolSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Flash(&_AlgebraPool.TransactOpts, recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_AlgebraPool *AlgebraPoolTransactorSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Flash(&_AlgebraPool.TransactOpts, recipient, amount0, amount1, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 initialPrice) returns()
func (_AlgebraPool *AlgebraPoolTransactor) Initialize(opts *bind.TransactOpts, initialPrice *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "initialize", initialPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 initialPrice) returns()
func (_AlgebraPool *AlgebraPoolSession) Initialize(initialPrice *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Initialize(&_AlgebraPool.TransactOpts, initialPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 initialPrice) returns()
func (_AlgebraPool *AlgebraPoolTransactorSession) Initialize(initialPrice *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Initialize(&_AlgebraPool.TransactOpts, initialPrice)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 liquidityDesired, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_AlgebraPool *AlgebraPoolTransactor) Mint(opts *bind.TransactOpts, sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, liquidityDesired *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "mint", sender, recipient, bottomTick, topTick, liquidityDesired, data)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 liquidityDesired, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_AlgebraPool *AlgebraPoolSession) Mint(sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, liquidityDesired *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Mint(&_AlgebraPool.TransactOpts, sender, recipient, bottomTick, topTick, liquidityDesired, data)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 liquidityDesired, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_AlgebraPool *AlgebraPoolTransactorSession) Mint(sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, liquidityDesired *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Mint(&_AlgebraPool.TransactOpts, sender, recipient, bottomTick, topTick, liquidityDesired, data)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x7c0112b7.
//
// Solidity: function setCommunityFee(uint8 communityFee0, uint8 communityFee1) returns()
func (_AlgebraPool *AlgebraPoolTransactor) SetCommunityFee(opts *bind.TransactOpts, communityFee0 uint8, communityFee1 uint8) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "setCommunityFee", communityFee0, communityFee1)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x7c0112b7.
//
// Solidity: function setCommunityFee(uint8 communityFee0, uint8 communityFee1) returns()
func (_AlgebraPool *AlgebraPoolSession) SetCommunityFee(communityFee0 uint8, communityFee1 uint8) (*types.Transaction, error) {
	return _AlgebraPool.Contract.SetCommunityFee(&_AlgebraPool.TransactOpts, communityFee0, communityFee1)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x7c0112b7.
//
// Solidity: function setCommunityFee(uint8 communityFee0, uint8 communityFee1) returns()
func (_AlgebraPool *AlgebraPoolTransactorSession) SetCommunityFee(communityFee0 uint8, communityFee1 uint8) (*types.Transaction, error) {
	return _AlgebraPool.Contract.SetCommunityFee(&_AlgebraPool.TransactOpts, communityFee0, communityFee1)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_AlgebraPool *AlgebraPoolTransactor) SetIncentive(opts *bind.TransactOpts, virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "setIncentive", virtualPoolAddress)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_AlgebraPool *AlgebraPoolSession) SetIncentive(virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _AlgebraPool.Contract.SetIncentive(&_AlgebraPool.TransactOpts, virtualPoolAddress)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_AlgebraPool *AlgebraPoolTransactorSession) SetIncentive(virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _AlgebraPool.Contract.SetIncentive(&_AlgebraPool.TransactOpts, virtualPoolAddress)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_AlgebraPool *AlgebraPoolTransactor) SetLiquidityCooldown(opts *bind.TransactOpts, newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "setLiquidityCooldown", newLiquidityCooldown)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_AlgebraPool *AlgebraPoolSession) SetLiquidityCooldown(newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _AlgebraPool.Contract.SetLiquidityCooldown(&_AlgebraPool.TransactOpts, newLiquidityCooldown)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_AlgebraPool *AlgebraPoolTransactorSession) SetLiquidityCooldown(newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _AlgebraPool.Contract.SetLiquidityCooldown(&_AlgebraPool.TransactOpts, newLiquidityCooldown)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_AlgebraPool *AlgebraPoolTransactor) SetTickSpacing(opts *bind.TransactOpts, newTickSpacing *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "setTickSpacing", newTickSpacing)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_AlgebraPool *AlgebraPoolSession) SetTickSpacing(newTickSpacing *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.Contract.SetTickSpacing(&_AlgebraPool.TransactOpts, newTickSpacing)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_AlgebraPool *AlgebraPoolTransactorSession) SetTickSpacing(newTickSpacing *big.Int) (*types.Transaction, error) {
	return _AlgebraPool.Contract.SetTickSpacing(&_AlgebraPool.TransactOpts, newTickSpacing)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_AlgebraPool *AlgebraPoolTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "swap", recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_AlgebraPool *AlgebraPoolSession) Swap(recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Swap(&_AlgebraPool.TransactOpts, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_AlgebraPool *AlgebraPoolTransactorSession) Swap(recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.Contract.Swap(&_AlgebraPool.TransactOpts, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_AlgebraPool *AlgebraPoolTransactor) SwapSupportingFeeOnInputTokens(opts *bind.TransactOpts, sender common.Address, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.contract.Transact(opts, "swapSupportingFeeOnInputTokens", sender, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_AlgebraPool *AlgebraPoolSession) SwapSupportingFeeOnInputTokens(sender common.Address, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.Contract.SwapSupportingFeeOnInputTokens(&_AlgebraPool.TransactOpts, sender, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_AlgebraPool *AlgebraPoolTransactorSession) SwapSupportingFeeOnInputTokens(sender common.Address, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _AlgebraPool.Contract.SwapSupportingFeeOnInputTokens(&_AlgebraPool.TransactOpts, sender, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// AlgebraPoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the AlgebraPool contract.
type AlgebraPoolBurnIterator struct {
	Event *AlgebraPoolBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolBurn represents a Burn event raised by the AlgebraPool contract.
type AlgebraPoolBurn struct {
	Owner           common.Address
	BottomTick      *big.Int
	TopTick         *big.Int
	LiquidityAmount *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_AlgebraPool *AlgebraPoolFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*AlgebraPoolBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "Burn", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolBurnIterator{contract: _AlgebraPool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_AlgebraPool *AlgebraPoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *AlgebraPoolBurn, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "Burn", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolBurn)
				if err := _AlgebraPool.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_AlgebraPool *AlgebraPoolFilterer) ParseBurn(log types.Log) (*AlgebraPoolBurn, error) {
	event := new(AlgebraPoolBurn)
	if err := _AlgebraPool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebraPoolCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the AlgebraPool contract.
type AlgebraPoolCollectIterator struct {
	Event *AlgebraPoolCollect // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolCollect)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolCollect)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolCollect represents a Collect event raised by the AlgebraPool contract.
type AlgebraPoolCollect struct {
	Owner      common.Address
	Recipient  common.Address
	BottomTick *big.Int
	TopTick    *big.Int
	Amount0    *big.Int
	Amount1    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_AlgebraPool *AlgebraPoolFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*AlgebraPoolCollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "Collect", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolCollectIterator{contract: _AlgebraPool.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_AlgebraPool *AlgebraPoolFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *AlgebraPoolCollect, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "Collect", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolCollect)
				if err := _AlgebraPool.contract.UnpackLog(event, "Collect", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollect is a log parse operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_AlgebraPool *AlgebraPoolFilterer) ParseCollect(log types.Log) (*AlgebraPoolCollect, error) {
	event := new(AlgebraPoolCollect)
	if err := _AlgebraPool.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebraPoolCommunityFeeIterator is returned from FilterCommunityFee and is used to iterate over the raw logs and unpacked data for CommunityFee events raised by the AlgebraPool contract.
type AlgebraPoolCommunityFeeIterator struct {
	Event *AlgebraPoolCommunityFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolCommunityFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolCommunityFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolCommunityFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolCommunityFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolCommunityFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolCommunityFee represents a CommunityFee event raised by the AlgebraPool contract.
type AlgebraPoolCommunityFee struct {
	CommunityFee0New uint8
	CommunityFee1New uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCommunityFee is a free log retrieval operation binding the contract event 0x9e22b964b08e25c3aaa72102bb0071c089258fb82d51271a8ddf5c24921356ee.
//
// Solidity: event CommunityFee(uint8 communityFee0New, uint8 communityFee1New)
func (_AlgebraPool *AlgebraPoolFilterer) FilterCommunityFee(opts *bind.FilterOpts) (*AlgebraPoolCommunityFeeIterator, error) {

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "CommunityFee")
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolCommunityFeeIterator{contract: _AlgebraPool.contract, event: "CommunityFee", logs: logs, sub: sub}, nil
}

// WatchCommunityFee is a free log subscription operation binding the contract event 0x9e22b964b08e25c3aaa72102bb0071c089258fb82d51271a8ddf5c24921356ee.
//
// Solidity: event CommunityFee(uint8 communityFee0New, uint8 communityFee1New)
func (_AlgebraPool *AlgebraPoolFilterer) WatchCommunityFee(opts *bind.WatchOpts, sink chan<- *AlgebraPoolCommunityFee) (event.Subscription, error) {

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "CommunityFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolCommunityFee)
				if err := _AlgebraPool.contract.UnpackLog(event, "CommunityFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommunityFee is a log parse operation binding the contract event 0x9e22b964b08e25c3aaa72102bb0071c089258fb82d51271a8ddf5c24921356ee.
//
// Solidity: event CommunityFee(uint8 communityFee0New, uint8 communityFee1New)
func (_AlgebraPool *AlgebraPoolFilterer) ParseCommunityFee(log types.Log) (*AlgebraPoolCommunityFee, error) {
	event := new(AlgebraPoolCommunityFee)
	if err := _AlgebraPool.contract.UnpackLog(event, "CommunityFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebraPoolFeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the AlgebraPool contract.
type AlgebraPoolFeeIterator struct {
	Event *AlgebraPoolFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolFee represents a Fee event raised by the AlgebraPool contract.
type AlgebraPoolFee struct {
	FeeZto uint16
	FeeOtz uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0x8a89de70856bccec096661388f305b9a75f5f65cb0d8a0e1e803c39dabedb57f.
//
// Solidity: event Fee(uint16 feeZto, uint16 feeOtz)
func (_AlgebraPool *AlgebraPoolFilterer) FilterFee(opts *bind.FilterOpts) (*AlgebraPoolFeeIterator, error) {

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolFeeIterator{contract: _AlgebraPool.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0x8a89de70856bccec096661388f305b9a75f5f65cb0d8a0e1e803c39dabedb57f.
//
// Solidity: event Fee(uint16 feeZto, uint16 feeOtz)
func (_AlgebraPool *AlgebraPoolFilterer) WatchFee(opts *bind.WatchOpts, sink chan<- *AlgebraPoolFee) (event.Subscription, error) {

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolFee)
				if err := _AlgebraPool.contract.UnpackLog(event, "Fee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFee is a log parse operation binding the contract event 0x8a89de70856bccec096661388f305b9a75f5f65cb0d8a0e1e803c39dabedb57f.
//
// Solidity: event Fee(uint16 feeZto, uint16 feeOtz)
func (_AlgebraPool *AlgebraPoolFilterer) ParseFee(log types.Log) (*AlgebraPoolFee, error) {
	event := new(AlgebraPoolFee)
	if err := _AlgebraPool.contract.UnpackLog(event, "Fee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebraPoolFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the AlgebraPool contract.
type AlgebraPoolFlashIterator struct {
	Event *AlgebraPoolFlash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolFlash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolFlash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolFlash represents a Flash event raised by the AlgebraPool contract.
type AlgebraPoolFlash struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Paid0     *big.Int
	Paid1     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_AlgebraPool *AlgebraPoolFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*AlgebraPoolFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolFlashIterator{contract: _AlgebraPool.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_AlgebraPool *AlgebraPoolFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *AlgebraPoolFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolFlash)
				if err := _AlgebraPool.contract.UnpackLog(event, "Flash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_AlgebraPool *AlgebraPoolFilterer) ParseFlash(log types.Log) (*AlgebraPoolFlash, error) {
	event := new(AlgebraPoolFlash)
	if err := _AlgebraPool.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebraPoolIncentiveIterator is returned from FilterIncentive and is used to iterate over the raw logs and unpacked data for Incentive events raised by the AlgebraPool contract.
type AlgebraPoolIncentiveIterator struct {
	Event *AlgebraPoolIncentive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolIncentive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolIncentive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolIncentive represents a Incentive event raised by the AlgebraPool contract.
type AlgebraPoolIncentive struct {
	VirtualPoolAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterIncentive is a free log retrieval operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_AlgebraPool *AlgebraPoolFilterer) FilterIncentive(opts *bind.FilterOpts, virtualPoolAddress []common.Address) (*AlgebraPoolIncentiveIterator, error) {

	var virtualPoolAddressRule []interface{}
	for _, virtualPoolAddressItem := range virtualPoolAddress {
		virtualPoolAddressRule = append(virtualPoolAddressRule, virtualPoolAddressItem)
	}

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "Incentive", virtualPoolAddressRule)
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolIncentiveIterator{contract: _AlgebraPool.contract, event: "Incentive", logs: logs, sub: sub}, nil
}

// WatchIncentive is a free log subscription operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_AlgebraPool *AlgebraPoolFilterer) WatchIncentive(opts *bind.WatchOpts, sink chan<- *AlgebraPoolIncentive, virtualPoolAddress []common.Address) (event.Subscription, error) {

	var virtualPoolAddressRule []interface{}
	for _, virtualPoolAddressItem := range virtualPoolAddress {
		virtualPoolAddressRule = append(virtualPoolAddressRule, virtualPoolAddressItem)
	}

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "Incentive", virtualPoolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolIncentive)
				if err := _AlgebraPool.contract.UnpackLog(event, "Incentive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncentive is a log parse operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_AlgebraPool *AlgebraPoolFilterer) ParseIncentive(log types.Log) (*AlgebraPoolIncentive, error) {
	event := new(AlgebraPoolIncentive)
	if err := _AlgebraPool.contract.UnpackLog(event, "Incentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebraPoolInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the AlgebraPool contract.
type AlgebraPoolInitializeIterator struct {
	Event *AlgebraPoolInitialize // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolInitialize)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolInitialize)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolInitialize represents a Initialize event raised by the AlgebraPool contract.
type AlgebraPoolInitialize struct {
	Price *big.Int
	Tick  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_AlgebraPool *AlgebraPoolFilterer) FilterInitialize(opts *bind.FilterOpts) (*AlgebraPoolInitializeIterator, error) {

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolInitializeIterator{contract: _AlgebraPool.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_AlgebraPool *AlgebraPoolFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *AlgebraPoolInitialize) (event.Subscription, error) {

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolInitialize)
				if err := _AlgebraPool.contract.UnpackLog(event, "Initialize", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialize is a log parse operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_AlgebraPool *AlgebraPoolFilterer) ParseInitialize(log types.Log) (*AlgebraPoolInitialize, error) {
	event := new(AlgebraPoolInitialize)
	if err := _AlgebraPool.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebraPoolLiquidityCooldownIterator is returned from FilterLiquidityCooldown and is used to iterate over the raw logs and unpacked data for LiquidityCooldown events raised by the AlgebraPool contract.
type AlgebraPoolLiquidityCooldownIterator struct {
	Event *AlgebraPoolLiquidityCooldown // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolLiquidityCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolLiquidityCooldown)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolLiquidityCooldown)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolLiquidityCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolLiquidityCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolLiquidityCooldown represents a LiquidityCooldown event raised by the AlgebraPool contract.
type AlgebraPoolLiquidityCooldown struct {
	LiquidityCooldown uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidityCooldown is a free log retrieval operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_AlgebraPool *AlgebraPoolFilterer) FilterLiquidityCooldown(opts *bind.FilterOpts) (*AlgebraPoolLiquidityCooldownIterator, error) {

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "LiquidityCooldown")
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolLiquidityCooldownIterator{contract: _AlgebraPool.contract, event: "LiquidityCooldown", logs: logs, sub: sub}, nil
}

// WatchLiquidityCooldown is a free log subscription operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_AlgebraPool *AlgebraPoolFilterer) WatchLiquidityCooldown(opts *bind.WatchOpts, sink chan<- *AlgebraPoolLiquidityCooldown) (event.Subscription, error) {

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "LiquidityCooldown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolLiquidityCooldown)
				if err := _AlgebraPool.contract.UnpackLog(event, "LiquidityCooldown", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidityCooldown is a log parse operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_AlgebraPool *AlgebraPoolFilterer) ParseLiquidityCooldown(log types.Log) (*AlgebraPoolLiquidityCooldown, error) {
	event := new(AlgebraPoolLiquidityCooldown)
	if err := _AlgebraPool.contract.UnpackLog(event, "LiquidityCooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebraPoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the AlgebraPool contract.
type AlgebraPoolMintIterator struct {
	Event *AlgebraPoolMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolMint represents a Mint event raised by the AlgebraPool contract.
type AlgebraPoolMint struct {
	Sender          common.Address
	Owner           common.Address
	BottomTick      *big.Int
	TopTick         *big.Int
	LiquidityAmount *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_AlgebraPool *AlgebraPoolFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*AlgebraPoolMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "Mint", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolMintIterator{contract: _AlgebraPool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_AlgebraPool *AlgebraPoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AlgebraPoolMint, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "Mint", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolMint)
				if err := _AlgebraPool.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_AlgebraPool *AlgebraPoolFilterer) ParseMint(log types.Log) (*AlgebraPoolMint, error) {
	event := new(AlgebraPoolMint)
	if err := _AlgebraPool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebraPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the AlgebraPool contract.
type AlgebraPoolSwapIterator struct {
	Event *AlgebraPoolSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolSwap represents a Swap event raised by the AlgebraPool contract.
type AlgebraPoolSwap struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Price     *big.Int
	Liquidity *big.Int
	Tick      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_AlgebraPool *AlgebraPoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*AlgebraPoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolSwapIterator{contract: _AlgebraPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_AlgebraPool *AlgebraPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *AlgebraPoolSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolSwap)
				if err := _AlgebraPool.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_AlgebraPool *AlgebraPoolFilterer) ParseSwap(log types.Log) (*AlgebraPoolSwap, error) {
	event := new(AlgebraPoolSwap)
	if err := _AlgebraPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebraPoolTickSpacingIterator is returned from FilterTickSpacing and is used to iterate over the raw logs and unpacked data for TickSpacing events raised by the AlgebraPool contract.
type AlgebraPoolTickSpacingIterator struct {
	Event *AlgebraPoolTickSpacing // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlgebraPoolTickSpacingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebraPoolTickSpacing)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlgebraPoolTickSpacing)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlgebraPoolTickSpacingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebraPoolTickSpacingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebraPoolTickSpacing represents a TickSpacing event raised by the AlgebraPool contract.
type AlgebraPoolTickSpacing struct {
	NewTickSpacing *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTickSpacing is a free log retrieval operation binding the contract event 0x01413b1d5d4c359e9a0daa7909ecda165f6e8c51fe2ff529d74b22a5a7c02645.
//
// Solidity: event TickSpacing(int24 newTickSpacing)
func (_AlgebraPool *AlgebraPoolFilterer) FilterTickSpacing(opts *bind.FilterOpts) (*AlgebraPoolTickSpacingIterator, error) {

	logs, sub, err := _AlgebraPool.contract.FilterLogs(opts, "TickSpacing")
	if err != nil {
		return nil, err
	}
	return &AlgebraPoolTickSpacingIterator{contract: _AlgebraPool.contract, event: "TickSpacing", logs: logs, sub: sub}, nil
}

// WatchTickSpacing is a free log subscription operation binding the contract event 0x01413b1d5d4c359e9a0daa7909ecda165f6e8c51fe2ff529d74b22a5a7c02645.
//
// Solidity: event TickSpacing(int24 newTickSpacing)
func (_AlgebraPool *AlgebraPoolFilterer) WatchTickSpacing(opts *bind.WatchOpts, sink chan<- *AlgebraPoolTickSpacing) (event.Subscription, error) {

	logs, sub, err := _AlgebraPool.contract.WatchLogs(opts, "TickSpacing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebraPoolTickSpacing)
				if err := _AlgebraPool.contract.UnpackLog(event, "TickSpacing", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTickSpacing is a log parse operation binding the contract event 0x01413b1d5d4c359e9a0daa7909ecda165f6e8c51fe2ff529d74b22a5a7c02645.
//
// Solidity: event TickSpacing(int24 newTickSpacing)
func (_AlgebraPool *AlgebraPoolFilterer) ParseTickSpacing(log types.Log) (*AlgebraPoolTickSpacing, error) {
	event := new(AlgebraPoolTickSpacing)
	if err := _AlgebraPool.contract.UnpackLog(event, "TickSpacing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
