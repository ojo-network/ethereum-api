// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool

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

// ComposableStablePoolNewPoolParams is an auto generated low-level Go binding around an user-defined struct.
type ComposableStablePoolNewPoolParams struct {
	Vault                           common.Address
	ProtocolFeeProvider             common.Address
	Name                            string
	Symbol                          string
	Tokens                          []common.Address
	RateProviders                   []common.Address
	TokenRateCacheDurations         []*big.Int
	ExemptFromYieldProtocolFeeFlags []bool
	AmplificationParameter          *big.Int
	SwapFeePercentage               *big.Int
	PauseWindowDuration             *big.Int
	BufferPeriodDuration            *big.Int
	Owner                           common.Address
	Version                         string
}

// IPoolSwapStructsSwapRequest is an auto generated low-level Go binding around an user-defined struct.
type IPoolSwapStructsSwapRequest struct {
	Kind            uint8
	TokenIn         common.Address
	TokenOut        common.Address
	Amount          *big.Int
	PoolId          [32]byte
	LastChangeBlock *big.Int
	From            common.Address
	To              common.Address
	UserData        []byte
}

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"contractIProtocolFeePercentagesProvider\",\"name\":\"protocolFeeProvider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIRateProvider[]\",\"name\":\"rateProviders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenRateCacheDurations\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"exemptFromYieldProtocolFeeFlags\",\"type\":\"bool[]\"},{\"internalType\":\"uint256\",\"name\":\"amplificationParameter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structComposableStablePool.NewPoolParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AmpUpdateStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentValue\",\"type\":\"uint256\"}],\"name\":\"AmpUpdateStopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"feeCache\",\"type\":\"bytes32\"}],\"name\":\"ProtocolFeePercentageCacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"RecoveryModeStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"TokenRateCacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIRateProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cacheDuration\",\"type\":\"uint256\"}],\"name\":\"TokenRateProviderSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActualSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAmplificationParameter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isUpdating\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"precision\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBptIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastJoinExitData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastJoinExitAmplification\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPostJoinExitInvariant\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumBpt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"}],\"name\":\"getProtocolFeePercentageCache\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeesCollector\",\"outputs\":[{\"internalType\":\"contractIProtocolFeesCollector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"}],\"name\":\"getProviderFeeId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateProviders\",\"outputs\":[{\"internalType\":\"contractIRateProvider[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getScalingFactors\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenRateCache\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inRecoveryMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenExemptFromYieldProtocolFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onExitPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onJoinPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIPoolSwapStructs.SwapRequest\",\"name\":\"swapRequest\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"indexIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexOut\",\"type\":\"uint256\"}],\"name\":\"onSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryJoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setSwapFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setTokenRateCacheDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rawEndValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"startAmplificationParameterUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopAmplificationParameterUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateProtocolFeePercentageCache\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"updateTokenRateCache\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pool *PoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pool *PoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Pool.Contract.DOMAINSEPARATOR(&_Pool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pool *PoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Pool.Contract.DOMAINSEPARATOR(&_Pool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pool *PoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pool *PoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Pool.Contract.Allowance(&_Pool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pool *PoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Pool.Contract.Allowance(&_Pool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pool *PoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pool *PoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Pool.Contract.BalanceOf(&_Pool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pool *PoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Pool.Contract.BalanceOf(&_Pool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pool *PoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pool *PoolSession) Decimals() (uint8, error) {
	return _Pool.Contract.Decimals(&_Pool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pool *PoolCallerSession) Decimals() (uint8, error) {
	return _Pool.Contract.Decimals(&_Pool.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Pool *PoolCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Pool *PoolSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _Pool.Contract.GetActionId(&_Pool.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Pool *PoolCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _Pool.Contract.GetActionId(&_Pool.CallOpts, selector)
}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_Pool *PoolCaller) GetActualSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getActualSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_Pool *PoolSession) GetActualSupply() (*big.Int, error) {
	return _Pool.Contract.GetActualSupply(&_Pool.CallOpts)
}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_Pool *PoolCallerSession) GetActualSupply() (*big.Int, error) {
	return _Pool.Contract.GetActualSupply(&_Pool.CallOpts)
}

// GetAmplificationParameter is a free data retrieval call binding the contract method 0x6daccffa.
//
// Solidity: function getAmplificationParameter() view returns(uint256 value, bool isUpdating, uint256 precision)
func (_Pool *PoolCaller) GetAmplificationParameter(opts *bind.CallOpts) (struct {
	Value      *big.Int
	IsUpdating bool
	Precision  *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAmplificationParameter")

	outstruct := new(struct {
		Value      *big.Int
		IsUpdating bool
		Precision  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsUpdating = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Precision = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAmplificationParameter is a free data retrieval call binding the contract method 0x6daccffa.
//
// Solidity: function getAmplificationParameter() view returns(uint256 value, bool isUpdating, uint256 precision)
func (_Pool *PoolSession) GetAmplificationParameter() (struct {
	Value      *big.Int
	IsUpdating bool
	Precision  *big.Int
}, error) {
	return _Pool.Contract.GetAmplificationParameter(&_Pool.CallOpts)
}

// GetAmplificationParameter is a free data retrieval call binding the contract method 0x6daccffa.
//
// Solidity: function getAmplificationParameter() view returns(uint256 value, bool isUpdating, uint256 precision)
func (_Pool *PoolCallerSession) GetAmplificationParameter() (struct {
	Value      *big.Int
	IsUpdating bool
	Precision  *big.Int
}, error) {
	return _Pool.Contract.GetAmplificationParameter(&_Pool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Pool *PoolCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Pool *PoolSession) GetAuthorizer() (common.Address, error) {
	return _Pool.Contract.GetAuthorizer(&_Pool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Pool *PoolCallerSession) GetAuthorizer() (common.Address, error) {
	return _Pool.Contract.GetAuthorizer(&_Pool.CallOpts)
}

// GetBptIndex is a free data retrieval call binding the contract method 0x82687a56.
//
// Solidity: function getBptIndex() view returns(uint256)
func (_Pool *PoolCaller) GetBptIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getBptIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBptIndex is a free data retrieval call binding the contract method 0x82687a56.
//
// Solidity: function getBptIndex() view returns(uint256)
func (_Pool *PoolSession) GetBptIndex() (*big.Int, error) {
	return _Pool.Contract.GetBptIndex(&_Pool.CallOpts)
}

// GetBptIndex is a free data retrieval call binding the contract method 0x82687a56.
//
// Solidity: function getBptIndex() view returns(uint256)
func (_Pool *PoolCallerSession) GetBptIndex() (*big.Int, error) {
	return _Pool.Contract.GetBptIndex(&_Pool.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Pool *PoolCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Pool *PoolSession) GetDomainSeparator() ([32]byte, error) {
	return _Pool.Contract.GetDomainSeparator(&_Pool.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Pool *PoolCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _Pool.Contract.GetDomainSeparator(&_Pool.CallOpts)
}

// GetLastJoinExitData is a free data retrieval call binding the contract method 0x3c975d51.
//
// Solidity: function getLastJoinExitData() view returns(uint256 lastJoinExitAmplification, uint256 lastPostJoinExitInvariant)
func (_Pool *PoolCaller) GetLastJoinExitData(opts *bind.CallOpts) (struct {
	LastJoinExitAmplification *big.Int
	LastPostJoinExitInvariant *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getLastJoinExitData")

	outstruct := new(struct {
		LastJoinExitAmplification *big.Int
		LastPostJoinExitInvariant *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastJoinExitAmplification = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastPostJoinExitInvariant = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLastJoinExitData is a free data retrieval call binding the contract method 0x3c975d51.
//
// Solidity: function getLastJoinExitData() view returns(uint256 lastJoinExitAmplification, uint256 lastPostJoinExitInvariant)
func (_Pool *PoolSession) GetLastJoinExitData() (struct {
	LastJoinExitAmplification *big.Int
	LastPostJoinExitInvariant *big.Int
}, error) {
	return _Pool.Contract.GetLastJoinExitData(&_Pool.CallOpts)
}

// GetLastJoinExitData is a free data retrieval call binding the contract method 0x3c975d51.
//
// Solidity: function getLastJoinExitData() view returns(uint256 lastJoinExitAmplification, uint256 lastPostJoinExitInvariant)
func (_Pool *PoolCallerSession) GetLastJoinExitData() (struct {
	LastJoinExitAmplification *big.Int
	LastPostJoinExitInvariant *big.Int
}, error) {
	return _Pool.Contract.GetLastJoinExitData(&_Pool.CallOpts)
}

// GetMinimumBpt is a free data retrieval call binding the contract method 0x04842d4c.
//
// Solidity: function getMinimumBpt() pure returns(uint256)
func (_Pool *PoolCaller) GetMinimumBpt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getMinimumBpt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumBpt is a free data retrieval call binding the contract method 0x04842d4c.
//
// Solidity: function getMinimumBpt() pure returns(uint256)
func (_Pool *PoolSession) GetMinimumBpt() (*big.Int, error) {
	return _Pool.Contract.GetMinimumBpt(&_Pool.CallOpts)
}

// GetMinimumBpt is a free data retrieval call binding the contract method 0x04842d4c.
//
// Solidity: function getMinimumBpt() pure returns(uint256)
func (_Pool *PoolCallerSession) GetMinimumBpt() (*big.Int, error) {
	return _Pool.Contract.GetMinimumBpt(&_Pool.CallOpts)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_Pool *PoolCaller) GetNextNonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getNextNonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_Pool *PoolSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _Pool.Contract.GetNextNonce(&_Pool.CallOpts, account)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_Pool *PoolCallerSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _Pool.Contract.GetNextNonce(&_Pool.CallOpts, account)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Pool *PoolCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Pool *PoolSession) GetOwner() (common.Address, error) {
	return _Pool.Contract.GetOwner(&_Pool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Pool *PoolCallerSession) GetOwner() (common.Address, error) {
	return _Pool.Contract.GetOwner(&_Pool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Pool *PoolCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getPausedState")

	outstruct := new(struct {
		Paused              bool
		PauseWindowEndTime  *big.Int
		BufferPeriodEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PauseWindowEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BufferPeriodEndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Pool *PoolSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _Pool.Contract.GetPausedState(&_Pool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Pool *PoolCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _Pool.Contract.GetPausedState(&_Pool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_Pool *PoolCaller) GetPoolId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getPoolId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_Pool *PoolSession) GetPoolId() ([32]byte, error) {
	return _Pool.Contract.GetPoolId(&_Pool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_Pool *PoolCallerSession) GetPoolId() ([32]byte, error) {
	return _Pool.Contract.GetPoolId(&_Pool.CallOpts)
}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_Pool *PoolCaller) GetProtocolFeePercentageCache(opts *bind.CallOpts, feeType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getProtocolFeePercentageCache", feeType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_Pool *PoolSession) GetProtocolFeePercentageCache(feeType *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetProtocolFeePercentageCache(&_Pool.CallOpts, feeType)
}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_Pool *PoolCallerSession) GetProtocolFeePercentageCache(feeType *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetProtocolFeePercentageCache(&_Pool.CallOpts, feeType)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Pool *PoolCaller) GetProtocolFeesCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getProtocolFeesCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Pool *PoolSession) GetProtocolFeesCollector() (common.Address, error) {
	return _Pool.Contract.GetProtocolFeesCollector(&_Pool.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Pool *PoolCallerSession) GetProtocolFeesCollector() (common.Address, error) {
	return _Pool.Contract.GetProtocolFeesCollector(&_Pool.CallOpts)
}

// GetProviderFeeId is a free data retrieval call binding the contract method 0x4df77ce0.
//
// Solidity: function getProviderFeeId(uint256 feeType) view returns(uint256)
func (_Pool *PoolCaller) GetProviderFeeId(opts *bind.CallOpts, feeType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getProviderFeeId", feeType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProviderFeeId is a free data retrieval call binding the contract method 0x4df77ce0.
//
// Solidity: function getProviderFeeId(uint256 feeType) view returns(uint256)
func (_Pool *PoolSession) GetProviderFeeId(feeType *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetProviderFeeId(&_Pool.CallOpts, feeType)
}

// GetProviderFeeId is a free data retrieval call binding the contract method 0x4df77ce0.
//
// Solidity: function getProviderFeeId(uint256 feeType) view returns(uint256)
func (_Pool *PoolCallerSession) GetProviderFeeId(feeType *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetProviderFeeId(&_Pool.CallOpts, feeType)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_Pool *PoolCaller) GetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_Pool *PoolSession) GetRate() (*big.Int, error) {
	return _Pool.Contract.GetRate(&_Pool.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_Pool *PoolCallerSession) GetRate() (*big.Int, error) {
	return _Pool.Contract.GetRate(&_Pool.CallOpts)
}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_Pool *PoolCaller) GetRateProviders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getRateProviders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_Pool *PoolSession) GetRateProviders() ([]common.Address, error) {
	return _Pool.Contract.GetRateProviders(&_Pool.CallOpts)
}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_Pool *PoolCallerSession) GetRateProviders() ([]common.Address, error) {
	return _Pool.Contract.GetRateProviders(&_Pool.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_Pool *PoolCaller) GetScalingFactors(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getScalingFactors")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_Pool *PoolSession) GetScalingFactors() ([]*big.Int, error) {
	return _Pool.Contract.GetScalingFactors(&_Pool.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_Pool *PoolCallerSession) GetScalingFactors() ([]*big.Int, error) {
	return _Pool.Contract.GetScalingFactors(&_Pool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_Pool *PoolCaller) GetSwapFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getSwapFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_Pool *PoolSession) GetSwapFeePercentage() (*big.Int, error) {
	return _Pool.Contract.GetSwapFeePercentage(&_Pool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_Pool *PoolCallerSession) GetSwapFeePercentage() (*big.Int, error) {
	return _Pool.Contract.GetSwapFeePercentage(&_Pool.CallOpts)
}

// GetTokenRate is a free data retrieval call binding the contract method 0x54dea00a.
//
// Solidity: function getTokenRate(address token) view returns(uint256)
func (_Pool *PoolCaller) GetTokenRate(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getTokenRate", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenRate is a free data retrieval call binding the contract method 0x54dea00a.
//
// Solidity: function getTokenRate(address token) view returns(uint256)
func (_Pool *PoolSession) GetTokenRate(token common.Address) (*big.Int, error) {
	return _Pool.Contract.GetTokenRate(&_Pool.CallOpts, token)
}

// GetTokenRate is a free data retrieval call binding the contract method 0x54dea00a.
//
// Solidity: function getTokenRate(address token) view returns(uint256)
func (_Pool *PoolCallerSession) GetTokenRate(token common.Address) (*big.Int, error) {
	return _Pool.Contract.GetTokenRate(&_Pool.CallOpts, token)
}

// GetTokenRateCache is a free data retrieval call binding the contract method 0x7f1260d1.
//
// Solidity: function getTokenRateCache(address token) view returns(uint256 rate, uint256 oldRate, uint256 duration, uint256 expires)
func (_Pool *PoolCaller) GetTokenRateCache(opts *bind.CallOpts, token common.Address) (struct {
	Rate     *big.Int
	OldRate  *big.Int
	Duration *big.Int
	Expires  *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getTokenRateCache", token)

	outstruct := new(struct {
		Rate     *big.Int
		OldRate  *big.Int
		Duration *big.Int
		Expires  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OldRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Expires = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenRateCache is a free data retrieval call binding the contract method 0x7f1260d1.
//
// Solidity: function getTokenRateCache(address token) view returns(uint256 rate, uint256 oldRate, uint256 duration, uint256 expires)
func (_Pool *PoolSession) GetTokenRateCache(token common.Address) (struct {
	Rate     *big.Int
	OldRate  *big.Int
	Duration *big.Int
	Expires  *big.Int
}, error) {
	return _Pool.Contract.GetTokenRateCache(&_Pool.CallOpts, token)
}

// GetTokenRateCache is a free data retrieval call binding the contract method 0x7f1260d1.
//
// Solidity: function getTokenRateCache(address token) view returns(uint256 rate, uint256 oldRate, uint256 duration, uint256 expires)
func (_Pool *PoolCallerSession) GetTokenRateCache(token common.Address) (struct {
	Rate     *big.Int
	OldRate  *big.Int
	Duration *big.Int
	Expires  *big.Int
}, error) {
	return _Pool.Contract.GetTokenRateCache(&_Pool.CallOpts, token)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_Pool *PoolCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_Pool *PoolSession) GetVault() (common.Address, error) {
	return _Pool.Contract.GetVault(&_Pool.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_Pool *PoolCallerSession) GetVault() (common.Address, error) {
	return _Pool.Contract.GetVault(&_Pool.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_Pool *PoolCaller) InRecoveryMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "inRecoveryMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_Pool *PoolSession) InRecoveryMode() (bool, error) {
	return _Pool.Contract.InRecoveryMode(&_Pool.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_Pool *PoolCallerSession) InRecoveryMode() (bool, error) {
	return _Pool.Contract.InRecoveryMode(&_Pool.CallOpts)
}

// IsTokenExemptFromYieldProtocolFee is a free data retrieval call binding the contract method 0xab7759f1.
//
// Solidity: function isTokenExemptFromYieldProtocolFee(address token) view returns(bool)
func (_Pool *PoolCaller) IsTokenExemptFromYieldProtocolFee(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "isTokenExemptFromYieldProtocolFee", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenExemptFromYieldProtocolFee is a free data retrieval call binding the contract method 0xab7759f1.
//
// Solidity: function isTokenExemptFromYieldProtocolFee(address token) view returns(bool)
func (_Pool *PoolSession) IsTokenExemptFromYieldProtocolFee(token common.Address) (bool, error) {
	return _Pool.Contract.IsTokenExemptFromYieldProtocolFee(&_Pool.CallOpts, token)
}

// IsTokenExemptFromYieldProtocolFee is a free data retrieval call binding the contract method 0xab7759f1.
//
// Solidity: function isTokenExemptFromYieldProtocolFee(address token) view returns(bool)
func (_Pool *PoolCallerSession) IsTokenExemptFromYieldProtocolFee(token common.Address) (bool, error) {
	return _Pool.Contract.IsTokenExemptFromYieldProtocolFee(&_Pool.CallOpts, token)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolSession) Name() (string, error) {
	return _Pool.Contract.Name(&_Pool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolCallerSession) Name() (string, error) {
	return _Pool.Contract.Name(&_Pool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Pool *PoolCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Pool *PoolSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Pool.Contract.Nonces(&_Pool.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Pool *PoolCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Pool.Contract.Nonces(&_Pool.CallOpts, owner)
}

// QueryExit is a free data retrieval call binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptIn, uint256[] amountsOut)
func (_Pool *PoolCaller) QueryExit(opts *bind.CallOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptIn      *big.Int
	AmountsOut []*big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "queryExit", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)

	outstruct := new(struct {
		BptIn      *big.Int
		AmountsOut []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BptIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountsOut = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// QueryExit is a free data retrieval call binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptIn, uint256[] amountsOut)
func (_Pool *PoolSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptIn      *big.Int
	AmountsOut []*big.Int
}, error) {
	return _Pool.Contract.QueryExit(&_Pool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a free data retrieval call binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptIn, uint256[] amountsOut)
func (_Pool *PoolCallerSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptIn      *big.Int
	AmountsOut []*big.Int
}, error) {
	return _Pool.Contract.QueryExit(&_Pool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a free data retrieval call binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptOut, uint256[] amountsIn)
func (_Pool *PoolCaller) QueryJoin(opts *bind.CallOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptOut    *big.Int
	AmountsIn []*big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "queryJoin", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)

	outstruct := new(struct {
		BptOut    *big.Int
		AmountsIn []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BptOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountsIn = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// QueryJoin is a free data retrieval call binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptOut, uint256[] amountsIn)
func (_Pool *PoolSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptOut    *big.Int
	AmountsIn []*big.Int
}, error) {
	return _Pool.Contract.QueryJoin(&_Pool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a free data retrieval call binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptOut, uint256[] amountsIn)
func (_Pool *PoolCallerSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptOut    *big.Int
	AmountsIn []*big.Int
}, error) {
	return _Pool.Contract.QueryJoin(&_Pool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolSession) Symbol() (string, error) {
	return _Pool.Contract.Symbol(&_Pool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolCallerSession) Symbol() (string, error) {
	return _Pool.Contract.Symbol(&_Pool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolSession) TotalSupply() (*big.Int, error) {
	return _Pool.Contract.TotalSupply(&_Pool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolCallerSession) TotalSupply() (*big.Int, error) {
	return _Pool.Contract.TotalSupply(&_Pool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Pool *PoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Pool *PoolSession) Version() (string, error) {
	return _Pool.Contract.Version(&_Pool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Pool *PoolCallerSession) Version() (string, error) {
	return _Pool.Contract.Version(&_Pool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pool *PoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Approve(&_Pool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Approve(&_Pool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Pool *PoolSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DecreaseAllowance(&_Pool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DecreaseAllowance(&_Pool.TransactOpts, spender, amount)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_Pool *PoolTransactor) DisableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "disableRecoveryMode")
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_Pool *PoolSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _Pool.Contract.DisableRecoveryMode(&_Pool.TransactOpts)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_Pool *PoolTransactorSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _Pool.Contract.DisableRecoveryMode(&_Pool.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_Pool *PoolTransactor) EnableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "enableRecoveryMode")
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_Pool *PoolSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _Pool.Contract.EnableRecoveryMode(&_Pool.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_Pool *PoolTransactorSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _Pool.Contract.EnableRecoveryMode(&_Pool.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pool *PoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pool *PoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.IncreaseAllowance(&_Pool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pool *PoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.IncreaseAllowance(&_Pool.TransactOpts, spender, addedValue)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_Pool *PoolTransactor) OnExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "onExitPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_Pool *PoolSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _Pool.Contract.OnExitPool(&_Pool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_Pool *PoolTransactorSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _Pool.Contract.OnExitPool(&_Pool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_Pool *PoolTransactor) OnJoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "onJoinPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_Pool *PoolSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _Pool.Contract.OnJoinPool(&_Pool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_Pool *PoolTransactorSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _Pool.Contract.OnJoinPool(&_Pool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnSwap is a paid mutator transaction binding the contract method 0x01ec954a.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) swapRequest, uint256[] balances, uint256 indexIn, uint256 indexOut) returns(uint256)
func (_Pool *PoolTransactor) OnSwap(opts *bind.TransactOpts, swapRequest IPoolSwapStructsSwapRequest, balances []*big.Int, indexIn *big.Int, indexOut *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "onSwap", swapRequest, balances, indexIn, indexOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x01ec954a.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) swapRequest, uint256[] balances, uint256 indexIn, uint256 indexOut) returns(uint256)
func (_Pool *PoolSession) OnSwap(swapRequest IPoolSwapStructsSwapRequest, balances []*big.Int, indexIn *big.Int, indexOut *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.OnSwap(&_Pool.TransactOpts, swapRequest, balances, indexIn, indexOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x01ec954a.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) swapRequest, uint256[] balances, uint256 indexIn, uint256 indexOut) returns(uint256)
func (_Pool *PoolTransactorSession) OnSwap(swapRequest IPoolSwapStructsSwapRequest, balances []*big.Int, indexIn *big.Int, indexOut *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.OnSwap(&_Pool.TransactOpts, swapRequest, balances, indexIn, indexOut)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolSession) Pause() (*types.Transaction, error) {
	return _Pool.Contract.Pause(&_Pool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolTransactorSession) Pause() (*types.Transaction, error) {
	return _Pool.Contract.Pause(&_Pool.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pool *PoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pool *PoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pool.Contract.Permit(&_Pool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pool *PoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pool.Contract.Permit(&_Pool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_Pool *PoolTransactor) SetSwapFeePercentage(opts *bind.TransactOpts, swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setSwapFeePercentage", swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_Pool *PoolSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetSwapFeePercentage(&_Pool.TransactOpts, swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_Pool *PoolTransactorSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetSwapFeePercentage(&_Pool.TransactOpts, swapFeePercentage)
}

// SetTokenRateCacheDuration is a paid mutator transaction binding the contract method 0xf4b7964d.
//
// Solidity: function setTokenRateCacheDuration(address token, uint256 duration) returns()
func (_Pool *PoolTransactor) SetTokenRateCacheDuration(opts *bind.TransactOpts, token common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setTokenRateCacheDuration", token, duration)
}

// SetTokenRateCacheDuration is a paid mutator transaction binding the contract method 0xf4b7964d.
//
// Solidity: function setTokenRateCacheDuration(address token, uint256 duration) returns()
func (_Pool *PoolSession) SetTokenRateCacheDuration(token common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetTokenRateCacheDuration(&_Pool.TransactOpts, token, duration)
}

// SetTokenRateCacheDuration is a paid mutator transaction binding the contract method 0xf4b7964d.
//
// Solidity: function setTokenRateCacheDuration(address token, uint256 duration) returns()
func (_Pool *PoolTransactorSession) SetTokenRateCacheDuration(token common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetTokenRateCacheDuration(&_Pool.TransactOpts, token, duration)
}

// StartAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0x2f1a0bc9.
//
// Solidity: function startAmplificationParameterUpdate(uint256 rawEndValue, uint256 endTime) returns()
func (_Pool *PoolTransactor) StartAmplificationParameterUpdate(opts *bind.TransactOpts, rawEndValue *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "startAmplificationParameterUpdate", rawEndValue, endTime)
}

// StartAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0x2f1a0bc9.
//
// Solidity: function startAmplificationParameterUpdate(uint256 rawEndValue, uint256 endTime) returns()
func (_Pool *PoolSession) StartAmplificationParameterUpdate(rawEndValue *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.StartAmplificationParameterUpdate(&_Pool.TransactOpts, rawEndValue, endTime)
}

// StartAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0x2f1a0bc9.
//
// Solidity: function startAmplificationParameterUpdate(uint256 rawEndValue, uint256 endTime) returns()
func (_Pool *PoolTransactorSession) StartAmplificationParameterUpdate(rawEndValue *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.StartAmplificationParameterUpdate(&_Pool.TransactOpts, rawEndValue, endTime)
}

// StopAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0xeb0f24d6.
//
// Solidity: function stopAmplificationParameterUpdate() returns()
func (_Pool *PoolTransactor) StopAmplificationParameterUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "stopAmplificationParameterUpdate")
}

// StopAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0xeb0f24d6.
//
// Solidity: function stopAmplificationParameterUpdate() returns()
func (_Pool *PoolSession) StopAmplificationParameterUpdate() (*types.Transaction, error) {
	return _Pool.Contract.StopAmplificationParameterUpdate(&_Pool.TransactOpts)
}

// StopAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0xeb0f24d6.
//
// Solidity: function stopAmplificationParameterUpdate() returns()
func (_Pool *PoolTransactorSession) StopAmplificationParameterUpdate() (*types.Transaction, error) {
	return _Pool.Contract.StopAmplificationParameterUpdate(&_Pool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Pool *PoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Transfer(&_Pool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Transfer(&_Pool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Pool *PoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferFrom(&_Pool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferFrom(&_Pool.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolSession) Unpause() (*types.Transaction, error) {
	return _Pool.Contract.Unpause(&_Pool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pool.Contract.Unpause(&_Pool.TransactOpts)
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_Pool *PoolTransactor) UpdateProtocolFeePercentageCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "updateProtocolFeePercentageCache")
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_Pool *PoolSession) UpdateProtocolFeePercentageCache() (*types.Transaction, error) {
	return _Pool.Contract.UpdateProtocolFeePercentageCache(&_Pool.TransactOpts)
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_Pool *PoolTransactorSession) UpdateProtocolFeePercentageCache() (*types.Transaction, error) {
	return _Pool.Contract.UpdateProtocolFeePercentageCache(&_Pool.TransactOpts)
}

// UpdateTokenRateCache is a paid mutator transaction binding the contract method 0x2df2c7c0.
//
// Solidity: function updateTokenRateCache(address token) returns()
func (_Pool *PoolTransactor) UpdateTokenRateCache(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "updateTokenRateCache", token)
}

// UpdateTokenRateCache is a paid mutator transaction binding the contract method 0x2df2c7c0.
//
// Solidity: function updateTokenRateCache(address token) returns()
func (_Pool *PoolSession) UpdateTokenRateCache(token common.Address) (*types.Transaction, error) {
	return _Pool.Contract.UpdateTokenRateCache(&_Pool.TransactOpts, token)
}

// UpdateTokenRateCache is a paid mutator transaction binding the contract method 0x2df2c7c0.
//
// Solidity: function updateTokenRateCache(address token) returns()
func (_Pool *PoolTransactorSession) UpdateTokenRateCache(token common.Address) (*types.Transaction, error) {
	return _Pool.Contract.UpdateTokenRateCache(&_Pool.TransactOpts, token)
}

// PoolAmpUpdateStartedIterator is returned from FilterAmpUpdateStarted and is used to iterate over the raw logs and unpacked data for AmpUpdateStarted events raised by the Pool contract.
type PoolAmpUpdateStartedIterator struct {
	Event *PoolAmpUpdateStarted // Event containing the contract specifics and raw log

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
func (it *PoolAmpUpdateStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAmpUpdateStarted)
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
		it.Event = new(PoolAmpUpdateStarted)
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
func (it *PoolAmpUpdateStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAmpUpdateStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAmpUpdateStarted represents a AmpUpdateStarted event raised by the Pool contract.
type PoolAmpUpdateStarted struct {
	StartValue *big.Int
	EndValue   *big.Int
	StartTime  *big.Int
	EndTime    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAmpUpdateStarted is a free log retrieval operation binding the contract event 0x1835882ee7a34ac194f717a35e09bb1d24c82a3b9d854ab6c9749525b714cdf2.
//
// Solidity: event AmpUpdateStarted(uint256 startValue, uint256 endValue, uint256 startTime, uint256 endTime)
func (_Pool *PoolFilterer) FilterAmpUpdateStarted(opts *bind.FilterOpts) (*PoolAmpUpdateStartedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AmpUpdateStarted")
	if err != nil {
		return nil, err
	}
	return &PoolAmpUpdateStartedIterator{contract: _Pool.contract, event: "AmpUpdateStarted", logs: logs, sub: sub}, nil
}

// WatchAmpUpdateStarted is a free log subscription operation binding the contract event 0x1835882ee7a34ac194f717a35e09bb1d24c82a3b9d854ab6c9749525b714cdf2.
//
// Solidity: event AmpUpdateStarted(uint256 startValue, uint256 endValue, uint256 startTime, uint256 endTime)
func (_Pool *PoolFilterer) WatchAmpUpdateStarted(opts *bind.WatchOpts, sink chan<- *PoolAmpUpdateStarted) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AmpUpdateStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAmpUpdateStarted)
				if err := _Pool.contract.UnpackLog(event, "AmpUpdateStarted", log); err != nil {
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

// ParseAmpUpdateStarted is a log parse operation binding the contract event 0x1835882ee7a34ac194f717a35e09bb1d24c82a3b9d854ab6c9749525b714cdf2.
//
// Solidity: event AmpUpdateStarted(uint256 startValue, uint256 endValue, uint256 startTime, uint256 endTime)
func (_Pool *PoolFilterer) ParseAmpUpdateStarted(log types.Log) (*PoolAmpUpdateStarted, error) {
	event := new(PoolAmpUpdateStarted)
	if err := _Pool.contract.UnpackLog(event, "AmpUpdateStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAmpUpdateStoppedIterator is returned from FilterAmpUpdateStopped and is used to iterate over the raw logs and unpacked data for AmpUpdateStopped events raised by the Pool contract.
type PoolAmpUpdateStoppedIterator struct {
	Event *PoolAmpUpdateStopped // Event containing the contract specifics and raw log

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
func (it *PoolAmpUpdateStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAmpUpdateStopped)
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
		it.Event = new(PoolAmpUpdateStopped)
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
func (it *PoolAmpUpdateStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAmpUpdateStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAmpUpdateStopped represents a AmpUpdateStopped event raised by the Pool contract.
type PoolAmpUpdateStopped struct {
	CurrentValue *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAmpUpdateStopped is a free log retrieval operation binding the contract event 0xa0d01593e47e69d07e0ccd87bece09411e07dd1ed40ca8f2e7af2976542a0233.
//
// Solidity: event AmpUpdateStopped(uint256 currentValue)
func (_Pool *PoolFilterer) FilterAmpUpdateStopped(opts *bind.FilterOpts) (*PoolAmpUpdateStoppedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AmpUpdateStopped")
	if err != nil {
		return nil, err
	}
	return &PoolAmpUpdateStoppedIterator{contract: _Pool.contract, event: "AmpUpdateStopped", logs: logs, sub: sub}, nil
}

// WatchAmpUpdateStopped is a free log subscription operation binding the contract event 0xa0d01593e47e69d07e0ccd87bece09411e07dd1ed40ca8f2e7af2976542a0233.
//
// Solidity: event AmpUpdateStopped(uint256 currentValue)
func (_Pool *PoolFilterer) WatchAmpUpdateStopped(opts *bind.WatchOpts, sink chan<- *PoolAmpUpdateStopped) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AmpUpdateStopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAmpUpdateStopped)
				if err := _Pool.contract.UnpackLog(event, "AmpUpdateStopped", log); err != nil {
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

// ParseAmpUpdateStopped is a log parse operation binding the contract event 0xa0d01593e47e69d07e0ccd87bece09411e07dd1ed40ca8f2e7af2976542a0233.
//
// Solidity: event AmpUpdateStopped(uint256 currentValue)
func (_Pool *PoolFilterer) ParseAmpUpdateStopped(log types.Log) (*PoolAmpUpdateStopped, error) {
	event := new(PoolAmpUpdateStopped)
	if err := _Pool.contract.UnpackLog(event, "AmpUpdateStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pool contract.
type PoolApprovalIterator struct {
	Event *PoolApproval // Event containing the contract specifics and raw log

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
func (it *PoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolApproval)
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
		it.Event = new(PoolApproval)
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
func (it *PoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolApproval represents a Approval event raised by the Pool contract.
type PoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PoolApprovalIterator{contract: _Pool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolApproval)
				if err := _Pool.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) ParseApproval(log types.Log) (*PoolApproval, error) {
	event := new(PoolApproval)
	if err := _Pool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the Pool contract.
type PoolPausedStateChangedIterator struct {
	Event *PoolPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *PoolPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPausedStateChanged)
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
		it.Event = new(PoolPausedStateChanged)
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
func (it *PoolPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPausedStateChanged represents a PausedStateChanged event raised by the Pool contract.
type PoolPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Pool *PoolFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*PoolPausedStateChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &PoolPausedStateChangedIterator{contract: _Pool.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Pool *PoolFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *PoolPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPausedStateChanged)
				if err := _Pool.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
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

// ParsePausedStateChanged is a log parse operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Pool *PoolFilterer) ParsePausedStateChanged(log types.Log) (*PoolPausedStateChanged, error) {
	event := new(PoolPausedStateChanged)
	if err := _Pool.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolProtocolFeePercentageCacheUpdatedIterator is returned from FilterProtocolFeePercentageCacheUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeePercentageCacheUpdated events raised by the Pool contract.
type PoolProtocolFeePercentageCacheUpdatedIterator struct {
	Event *PoolProtocolFeePercentageCacheUpdated // Event containing the contract specifics and raw log

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
func (it *PoolProtocolFeePercentageCacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolProtocolFeePercentageCacheUpdated)
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
		it.Event = new(PoolProtocolFeePercentageCacheUpdated)
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
func (it *PoolProtocolFeePercentageCacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolProtocolFeePercentageCacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolProtocolFeePercentageCacheUpdated represents a ProtocolFeePercentageCacheUpdated event raised by the Pool contract.
type PoolProtocolFeePercentageCacheUpdated struct {
	FeeCache [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeePercentageCacheUpdated is a free log retrieval operation binding the contract event 0xc3f0acc358200d8e08ac6ce20bc2f9f27893a344f813bf682b7859b3e521502e.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(bytes32 feeCache)
func (_Pool *PoolFilterer) FilterProtocolFeePercentageCacheUpdated(opts *bind.FilterOpts) (*PoolProtocolFeePercentageCacheUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ProtocolFeePercentageCacheUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolProtocolFeePercentageCacheUpdatedIterator{contract: _Pool.contract, event: "ProtocolFeePercentageCacheUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeePercentageCacheUpdated is a free log subscription operation binding the contract event 0xc3f0acc358200d8e08ac6ce20bc2f9f27893a344f813bf682b7859b3e521502e.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(bytes32 feeCache)
func (_Pool *PoolFilterer) WatchProtocolFeePercentageCacheUpdated(opts *bind.WatchOpts, sink chan<- *PoolProtocolFeePercentageCacheUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ProtocolFeePercentageCacheUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolProtocolFeePercentageCacheUpdated)
				if err := _Pool.contract.UnpackLog(event, "ProtocolFeePercentageCacheUpdated", log); err != nil {
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

// ParseProtocolFeePercentageCacheUpdated is a log parse operation binding the contract event 0xc3f0acc358200d8e08ac6ce20bc2f9f27893a344f813bf682b7859b3e521502e.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(bytes32 feeCache)
func (_Pool *PoolFilterer) ParseProtocolFeePercentageCacheUpdated(log types.Log) (*PoolProtocolFeePercentageCacheUpdated, error) {
	event := new(PoolProtocolFeePercentageCacheUpdated)
	if err := _Pool.contract.UnpackLog(event, "ProtocolFeePercentageCacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRecoveryModeStateChangedIterator is returned from FilterRecoveryModeStateChanged and is used to iterate over the raw logs and unpacked data for RecoveryModeStateChanged events raised by the Pool contract.
type PoolRecoveryModeStateChangedIterator struct {
	Event *PoolRecoveryModeStateChanged // Event containing the contract specifics and raw log

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
func (it *PoolRecoveryModeStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRecoveryModeStateChanged)
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
		it.Event = new(PoolRecoveryModeStateChanged)
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
func (it *PoolRecoveryModeStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRecoveryModeStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRecoveryModeStateChanged represents a RecoveryModeStateChanged event raised by the Pool contract.
type PoolRecoveryModeStateChanged struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecoveryModeStateChanged is a free log retrieval operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_Pool *PoolFilterer) FilterRecoveryModeStateChanged(opts *bind.FilterOpts) (*PoolRecoveryModeStateChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RecoveryModeStateChanged")
	if err != nil {
		return nil, err
	}
	return &PoolRecoveryModeStateChangedIterator{contract: _Pool.contract, event: "RecoveryModeStateChanged", logs: logs, sub: sub}, nil
}

// WatchRecoveryModeStateChanged is a free log subscription operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_Pool *PoolFilterer) WatchRecoveryModeStateChanged(opts *bind.WatchOpts, sink chan<- *PoolRecoveryModeStateChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RecoveryModeStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRecoveryModeStateChanged)
				if err := _Pool.contract.UnpackLog(event, "RecoveryModeStateChanged", log); err != nil {
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

// ParseRecoveryModeStateChanged is a log parse operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_Pool *PoolFilterer) ParseRecoveryModeStateChanged(log types.Log) (*PoolRecoveryModeStateChanged, error) {
	event := new(PoolRecoveryModeStateChanged)
	if err := _Pool.contract.UnpackLog(event, "RecoveryModeStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSwapFeePercentageChangedIterator is returned from FilterSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for SwapFeePercentageChanged events raised by the Pool contract.
type PoolSwapFeePercentageChangedIterator struct {
	Event *PoolSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *PoolSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSwapFeePercentageChanged)
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
		it.Event = new(PoolSwapFeePercentageChanged)
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
func (it *PoolSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSwapFeePercentageChanged represents a SwapFeePercentageChanged event raised by the Pool contract.
type PoolSwapFeePercentageChanged struct {
	SwapFeePercentage *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_Pool *PoolFilterer) FilterSwapFeePercentageChanged(opts *bind.FilterOpts) (*PoolSwapFeePercentageChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &PoolSwapFeePercentageChangedIterator{contract: _Pool.contract, event: "SwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeePercentageChanged is a free log subscription operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_Pool *PoolFilterer) WatchSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *PoolSwapFeePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSwapFeePercentageChanged)
				if err := _Pool.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
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

// ParseSwapFeePercentageChanged is a log parse operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_Pool *PoolFilterer) ParseSwapFeePercentageChanged(log types.Log) (*PoolSwapFeePercentageChanged, error) {
	event := new(PoolSwapFeePercentageChanged)
	if err := _Pool.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTokenRateCacheUpdatedIterator is returned from FilterTokenRateCacheUpdated and is used to iterate over the raw logs and unpacked data for TokenRateCacheUpdated events raised by the Pool contract.
type PoolTokenRateCacheUpdatedIterator struct {
	Event *PoolTokenRateCacheUpdated // Event containing the contract specifics and raw log

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
func (it *PoolTokenRateCacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTokenRateCacheUpdated)
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
		it.Event = new(PoolTokenRateCacheUpdated)
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
func (it *PoolTokenRateCacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTokenRateCacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTokenRateCacheUpdated represents a TokenRateCacheUpdated event raised by the Pool contract.
type PoolTokenRateCacheUpdated struct {
	TokenIndex *big.Int
	Rate       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenRateCacheUpdated is a free log retrieval operation binding the contract event 0xb77a83204ca282e08dc3a65b0a1ca32ea4e6875c38ef0bf5bf75e52a67354fac.
//
// Solidity: event TokenRateCacheUpdated(uint256 indexed tokenIndex, uint256 rate)
func (_Pool *PoolFilterer) FilterTokenRateCacheUpdated(opts *bind.FilterOpts, tokenIndex []*big.Int) (*PoolTokenRateCacheUpdatedIterator, error) {

	var tokenIndexRule []interface{}
	for _, tokenIndexItem := range tokenIndex {
		tokenIndexRule = append(tokenIndexRule, tokenIndexItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "TokenRateCacheUpdated", tokenIndexRule)
	if err != nil {
		return nil, err
	}
	return &PoolTokenRateCacheUpdatedIterator{contract: _Pool.contract, event: "TokenRateCacheUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenRateCacheUpdated is a free log subscription operation binding the contract event 0xb77a83204ca282e08dc3a65b0a1ca32ea4e6875c38ef0bf5bf75e52a67354fac.
//
// Solidity: event TokenRateCacheUpdated(uint256 indexed tokenIndex, uint256 rate)
func (_Pool *PoolFilterer) WatchTokenRateCacheUpdated(opts *bind.WatchOpts, sink chan<- *PoolTokenRateCacheUpdated, tokenIndex []*big.Int) (event.Subscription, error) {

	var tokenIndexRule []interface{}
	for _, tokenIndexItem := range tokenIndex {
		tokenIndexRule = append(tokenIndexRule, tokenIndexItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "TokenRateCacheUpdated", tokenIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTokenRateCacheUpdated)
				if err := _Pool.contract.UnpackLog(event, "TokenRateCacheUpdated", log); err != nil {
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

// ParseTokenRateCacheUpdated is a log parse operation binding the contract event 0xb77a83204ca282e08dc3a65b0a1ca32ea4e6875c38ef0bf5bf75e52a67354fac.
//
// Solidity: event TokenRateCacheUpdated(uint256 indexed tokenIndex, uint256 rate)
func (_Pool *PoolFilterer) ParseTokenRateCacheUpdated(log types.Log) (*PoolTokenRateCacheUpdated, error) {
	event := new(PoolTokenRateCacheUpdated)
	if err := _Pool.contract.UnpackLog(event, "TokenRateCacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTokenRateProviderSetIterator is returned from FilterTokenRateProviderSet and is used to iterate over the raw logs and unpacked data for TokenRateProviderSet events raised by the Pool contract.
type PoolTokenRateProviderSetIterator struct {
	Event *PoolTokenRateProviderSet // Event containing the contract specifics and raw log

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
func (it *PoolTokenRateProviderSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTokenRateProviderSet)
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
		it.Event = new(PoolTokenRateProviderSet)
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
func (it *PoolTokenRateProviderSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTokenRateProviderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTokenRateProviderSet represents a TokenRateProviderSet event raised by the Pool contract.
type PoolTokenRateProviderSet struct {
	TokenIndex    *big.Int
	Provider      common.Address
	CacheDuration *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenRateProviderSet is a free log retrieval operation binding the contract event 0xdd6d1c9badb346de6925b358a472c937b41698d2632696759e43fd6527feeec4.
//
// Solidity: event TokenRateProviderSet(uint256 indexed tokenIndex, address indexed provider, uint256 cacheDuration)
func (_Pool *PoolFilterer) FilterTokenRateProviderSet(opts *bind.FilterOpts, tokenIndex []*big.Int, provider []common.Address) (*PoolTokenRateProviderSetIterator, error) {

	var tokenIndexRule []interface{}
	for _, tokenIndexItem := range tokenIndex {
		tokenIndexRule = append(tokenIndexRule, tokenIndexItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "TokenRateProviderSet", tokenIndexRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &PoolTokenRateProviderSetIterator{contract: _Pool.contract, event: "TokenRateProviderSet", logs: logs, sub: sub}, nil
}

// WatchTokenRateProviderSet is a free log subscription operation binding the contract event 0xdd6d1c9badb346de6925b358a472c937b41698d2632696759e43fd6527feeec4.
//
// Solidity: event TokenRateProviderSet(uint256 indexed tokenIndex, address indexed provider, uint256 cacheDuration)
func (_Pool *PoolFilterer) WatchTokenRateProviderSet(opts *bind.WatchOpts, sink chan<- *PoolTokenRateProviderSet, tokenIndex []*big.Int, provider []common.Address) (event.Subscription, error) {

	var tokenIndexRule []interface{}
	for _, tokenIndexItem := range tokenIndex {
		tokenIndexRule = append(tokenIndexRule, tokenIndexItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "TokenRateProviderSet", tokenIndexRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTokenRateProviderSet)
				if err := _Pool.contract.UnpackLog(event, "TokenRateProviderSet", log); err != nil {
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

// ParseTokenRateProviderSet is a log parse operation binding the contract event 0xdd6d1c9badb346de6925b358a472c937b41698d2632696759e43fd6527feeec4.
//
// Solidity: event TokenRateProviderSet(uint256 indexed tokenIndex, address indexed provider, uint256 cacheDuration)
func (_Pool *PoolFilterer) ParseTokenRateProviderSet(log types.Log) (*PoolTokenRateProviderSet, error) {
	event := new(PoolTokenRateProviderSet)
	if err := _Pool.contract.UnpackLog(event, "TokenRateProviderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pool contract.
type PoolTransferIterator struct {
	Event *PoolTransfer // Event containing the contract specifics and raw log

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
func (it *PoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTransfer)
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
		it.Event = new(PoolTransfer)
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
func (it *PoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTransfer represents a Transfer event raised by the Pool contract.
type PoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolTransferIterator{contract: _Pool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTransfer)
				if err := _Pool.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) ParseTransfer(log types.Log) (*PoolTransfer, error) {
	event := new(PoolTransfer)
	if err := _Pool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
