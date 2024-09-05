// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curve

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

// TwocryptoOptimizedMetaData contains all meta data concerning the TwocryptoOptimized contract.
var TwocryptoOptimizedMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"packed_price_scale\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"packed_price_scale\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_index\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"approx_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"packed_price_scale\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampAgamma\",\"inputs\":[{\"name\":\"initial_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"current_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"current_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ClaimAdminFee\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true},{\"name\":\"tokens\",\"type\":\"uint256[2]\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[2]\"},{\"name\":\"_math\",\"type\":\"address\"},{\"name\":\"_salt\",\"type\":\"bytes32\"},{\"name\":\"packed_precisions\",\"type\":\"uint256\"},{\"name\":\"packed_gamma_A\",\"type\":\"uint256\"},{\"name\":\"packed_fee_params\",\"type\":\"uint256\"},{\"name\":\"packed_rebalancing_params\",\"type\":\"uint256\"},{\"name\":\"initial_price\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_received\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_received\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dx\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lp_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_scale\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_fee\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"xp\",\"type\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"precisions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_calc\",\"inputs\":[{\"name\":\"xp\",\"type\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A_gamma\",\"inputs\":[{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"future_gamma\",\"type\":\"uint256\"},{\"name\":\"future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A_gamma\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_parameters\",\"inputs\":[{\"name\":\"_new_mid_fee\",\"type\":\"uint256\"},{\"name\":\"_new_out_fee\",\"type\":\"uint256\"},{\"name\":\"_new_fee_gamma\",\"type\":\"uint256\"},{\"name\":\"_new_allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"_new_adjustment_step\",\"type\":\"uint256\"},{\"name\":\"_new_ma_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"MATH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit_a\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"packed_rebalancing_params\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"packed_fee_params\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ADMIN_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"salt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]}]",
}

// TwocryptoOptimizedABI is the input ABI used to generate the binding from.
// Deprecated: Use TwocryptoOptimizedMetaData.ABI instead.
var TwocryptoOptimizedABI = TwocryptoOptimizedMetaData.ABI

// TwocryptoOptimized is an auto generated Go binding around an Ethereum contract.
type TwocryptoOptimized struct {
	TwocryptoOptimizedCaller     // Read-only binding to the contract
	TwocryptoOptimizedTransactor // Write-only binding to the contract
	TwocryptoOptimizedFilterer   // Log filterer for contract events
}

// TwocryptoOptimizedCaller is an auto generated read-only Go binding around an Ethereum contract.
type TwocryptoOptimizedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwocryptoOptimizedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TwocryptoOptimizedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwocryptoOptimizedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TwocryptoOptimizedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwocryptoOptimizedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TwocryptoOptimizedSession struct {
	Contract     *TwocryptoOptimized // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TwocryptoOptimizedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TwocryptoOptimizedCallerSession struct {
	Contract *TwocryptoOptimizedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TwocryptoOptimizedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TwocryptoOptimizedTransactorSession struct {
	Contract     *TwocryptoOptimizedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TwocryptoOptimizedRaw is an auto generated low-level Go binding around an Ethereum contract.
type TwocryptoOptimizedRaw struct {
	Contract *TwocryptoOptimized // Generic contract binding to access the raw methods on
}

// TwocryptoOptimizedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TwocryptoOptimizedCallerRaw struct {
	Contract *TwocryptoOptimizedCaller // Generic read-only contract binding to access the raw methods on
}

// TwocryptoOptimizedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TwocryptoOptimizedTransactorRaw struct {
	Contract *TwocryptoOptimizedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTwocryptoOptimized creates a new instance of TwocryptoOptimized, bound to a specific deployed contract.
func NewTwocryptoOptimized(address common.Address, backend bind.ContractBackend) (*TwocryptoOptimized, error) {
	contract, err := bindTwocryptoOptimized(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimized{TwocryptoOptimizedCaller: TwocryptoOptimizedCaller{contract: contract}, TwocryptoOptimizedTransactor: TwocryptoOptimizedTransactor{contract: contract}, TwocryptoOptimizedFilterer: TwocryptoOptimizedFilterer{contract: contract}}, nil
}

// NewTwocryptoOptimizedCaller creates a new read-only instance of TwocryptoOptimized, bound to a specific deployed contract.
func NewTwocryptoOptimizedCaller(address common.Address, caller bind.ContractCaller) (*TwocryptoOptimizedCaller, error) {
	contract, err := bindTwocryptoOptimized(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedCaller{contract: contract}, nil
}

// NewTwocryptoOptimizedTransactor creates a new write-only instance of TwocryptoOptimized, bound to a specific deployed contract.
func NewTwocryptoOptimizedTransactor(address common.Address, transactor bind.ContractTransactor) (*TwocryptoOptimizedTransactor, error) {
	contract, err := bindTwocryptoOptimized(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedTransactor{contract: contract}, nil
}

// NewTwocryptoOptimizedFilterer creates a new log filterer instance of TwocryptoOptimized, bound to a specific deployed contract.
func NewTwocryptoOptimizedFilterer(address common.Address, filterer bind.ContractFilterer) (*TwocryptoOptimizedFilterer, error) {
	contract, err := bindTwocryptoOptimized(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedFilterer{contract: contract}, nil
}

// bindTwocryptoOptimized binds a generic wrapper to an already deployed contract.
func bindTwocryptoOptimized(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TwocryptoOptimizedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwocryptoOptimized *TwocryptoOptimizedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwocryptoOptimized.Contract.TwocryptoOptimizedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwocryptoOptimized *TwocryptoOptimizedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.TwocryptoOptimizedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwocryptoOptimized *TwocryptoOptimizedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.TwocryptoOptimizedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwocryptoOptimized *TwocryptoOptimizedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwocryptoOptimized.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) A() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.A(&_TwocryptoOptimized.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) A() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.A(&_TwocryptoOptimized.CallOpts)
}

// ADMINFEE is a free data retrieval call binding the contract method 0x4469ed14.
//
// Solidity: function ADMIN_FEE() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) ADMINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "ADMIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADMINFEE is a free data retrieval call binding the contract method 0x4469ed14.
//
// Solidity: function ADMIN_FEE() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) ADMINFEE() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.ADMINFEE(&_TwocryptoOptimized.CallOpts)
}

// ADMINFEE is a free data retrieval call binding the contract method 0x4469ed14.
//
// Solidity: function ADMIN_FEE() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) ADMINFEE() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.ADMINFEE(&_TwocryptoOptimized.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) D(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "D")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) D() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.D(&_TwocryptoOptimized.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) D() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.D(&_TwocryptoOptimized.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _TwocryptoOptimized.Contract.DOMAINSEPARATOR(&_TwocryptoOptimized.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _TwocryptoOptimized.Contract.DOMAINSEPARATOR(&_TwocryptoOptimized.CallOpts)
}

// MATH is a free data retrieval call binding the contract method 0xed6c1546.
//
// Solidity: function MATH() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) MATH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "MATH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MATH is a free data retrieval call binding the contract method 0xed6c1546.
//
// Solidity: function MATH() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) MATH() (common.Address, error) {
	return _TwocryptoOptimized.Contract.MATH(&_TwocryptoOptimized.CallOpts)
}

// MATH is a free data retrieval call binding the contract method 0xed6c1546.
//
// Solidity: function MATH() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) MATH() (common.Address, error) {
	return _TwocryptoOptimized.Contract.MATH(&_TwocryptoOptimized.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) AdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) AdjustmentStep() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.AdjustmentStep(&_TwocryptoOptimized.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) AdjustmentStep() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.AdjustmentStep(&_TwocryptoOptimized.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Admin() (common.Address, error) {
	return _TwocryptoOptimized.Contract.Admin(&_TwocryptoOptimized.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Admin() (common.Address, error) {
	return _TwocryptoOptimized.Contract.Admin(&_TwocryptoOptimized.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.Allowance(&_TwocryptoOptimized.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.Allowance(&_TwocryptoOptimized.CallOpts, arg0, arg1)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) AllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) AllowedExtraProfit() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.AllowedExtraProfit(&_TwocryptoOptimized.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) AllowedExtraProfit() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.AllowedExtraProfit(&_TwocryptoOptimized.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.BalanceOf(&_TwocryptoOptimized.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.BalanceOf(&_TwocryptoOptimized.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.Balances(&_TwocryptoOptimized.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.Balances(&_TwocryptoOptimized.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "calc_token_amount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) CalcTokenAmount(amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.CalcTokenAmount(&_TwocryptoOptimized.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) CalcTokenAmount(amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.CalcTokenAmount(&_TwocryptoOptimized.CallOpts, amounts, deposit)
}

// CalcTokenFee is a free data retrieval call binding the contract method 0xbcc8342e.
//
// Solidity: function calc_token_fee(uint256[2] amounts, uint256[2] xp) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) CalcTokenFee(opts *bind.CallOpts, amounts [2]*big.Int, xp [2]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "calc_token_fee", amounts, xp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenFee is a free data retrieval call binding the contract method 0xbcc8342e.
//
// Solidity: function calc_token_fee(uint256[2] amounts, uint256[2] xp) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) CalcTokenFee(amounts [2]*big.Int, xp [2]*big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.CalcTokenFee(&_TwocryptoOptimized.CallOpts, amounts, xp)
}

// CalcTokenFee is a free data retrieval call binding the contract method 0xbcc8342e.
//
// Solidity: function calc_token_fee(uint256[2] amounts, uint256[2] xp) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) CalcTokenFee(amounts [2]*big.Int, xp [2]*big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.CalcTokenFee(&_TwocryptoOptimized.CallOpts, amounts, xp)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "calc_withdraw_one_coin", token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.CalcWithdrawOneCoin(&_TwocryptoOptimized.CallOpts, token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.CalcWithdrawOneCoin(&_TwocryptoOptimized.CallOpts, token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _TwocryptoOptimized.Contract.Coins(&_TwocryptoOptimized.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _TwocryptoOptimized.Contract.Coins(&_TwocryptoOptimized.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Decimals() (uint8, error) {
	return _TwocryptoOptimized.Contract.Decimals(&_TwocryptoOptimized.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Decimals() (uint8, error) {
	return _TwocryptoOptimized.Contract.Decimals(&_TwocryptoOptimized.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Factory() (common.Address, error) {
	return _TwocryptoOptimized.Contract.Factory(&_TwocryptoOptimized.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Factory() (common.Address, error) {
	return _TwocryptoOptimized.Contract.Factory(&_TwocryptoOptimized.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Fee() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.Fee(&_TwocryptoOptimized.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Fee() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.Fee(&_TwocryptoOptimized.CallOpts)
}

// FeeCalc is a free data retrieval call binding the contract method 0x80823d9e.
//
// Solidity: function fee_calc(uint256[2] xp) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) FeeCalc(opts *bind.CallOpts, xp [2]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "fee_calc", xp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeCalc is a free data retrieval call binding the contract method 0x80823d9e.
//
// Solidity: function fee_calc(uint256[2] xp) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) FeeCalc(xp [2]*big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.FeeCalc(&_TwocryptoOptimized.CallOpts, xp)
}

// FeeCalc is a free data retrieval call binding the contract method 0x80823d9e.
//
// Solidity: function fee_calc(uint256[2] xp) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) FeeCalc(xp [2]*big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.FeeCalc(&_TwocryptoOptimized.CallOpts, xp)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) FeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) FeeGamma() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.FeeGamma(&_TwocryptoOptimized.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) FeeGamma() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.FeeGamma(&_TwocryptoOptimized.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) FeeReceiver() (common.Address, error) {
	return _TwocryptoOptimized.Contract.FeeReceiver(&_TwocryptoOptimized.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) FeeReceiver() (common.Address, error) {
	return _TwocryptoOptimized.Contract.FeeReceiver(&_TwocryptoOptimized.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) FutureAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "future_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) FutureAGamma() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.FutureAGamma(&_TwocryptoOptimized.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) FutureAGamma() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.FutureAGamma(&_TwocryptoOptimized.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) FutureAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "future_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) FutureAGammaTime() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.FutureAGammaTime(&_TwocryptoOptimized.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) FutureAGammaTime() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.FutureAGammaTime(&_TwocryptoOptimized.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Gamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Gamma() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.Gamma(&_TwocryptoOptimized.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Gamma() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.Gamma(&_TwocryptoOptimized.CallOpts)
}

// GetDx is a free data retrieval call binding the contract method 0x37ed3a7a.
//
// Solidity: function get_dx(uint256 i, uint256 j, uint256 dy) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) GetDx(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "get_dx", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDx is a free data retrieval call binding the contract method 0x37ed3a7a.
//
// Solidity: function get_dx(uint256 i, uint256 j, uint256 dy) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.GetDx(&_TwocryptoOptimized.CallOpts, i, j, dy)
}

// GetDx is a free data retrieval call binding the contract method 0x37ed3a7a.
//
// Solidity: function get_dx(uint256 i, uint256 j, uint256 dy) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.GetDx(&_TwocryptoOptimized.CallOpts, i, j, dy)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.GetDy(&_TwocryptoOptimized.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.GetDy(&_TwocryptoOptimized.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) GetVirtualPrice() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.GetVirtualPrice(&_TwocryptoOptimized.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.GetVirtualPrice(&_TwocryptoOptimized.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) InitialAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "initial_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) InitialAGamma() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.InitialAGamma(&_TwocryptoOptimized.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) InitialAGamma() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.InitialAGamma(&_TwocryptoOptimized.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) InitialAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "initial_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) InitialAGammaTime() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.InitialAGammaTime(&_TwocryptoOptimized.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) InitialAGammaTime() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.InitialAGammaTime(&_TwocryptoOptimized.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) LastPrices(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "last_prices")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) LastPrices() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.LastPrices(&_TwocryptoOptimized.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) LastPrices() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.LastPrices(&_TwocryptoOptimized.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x4d23bfa0.
//
// Solidity: function last_timestamp() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) LastTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "last_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x4d23bfa0.
//
// Solidity: function last_timestamp() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) LastTimestamp() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.LastTimestamp(&_TwocryptoOptimized.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x4d23bfa0.
//
// Solidity: function last_timestamp() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) LastTimestamp() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.LastTimestamp(&_TwocryptoOptimized.CallOpts)
}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) LpPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "lp_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) LpPrice() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.LpPrice(&_TwocryptoOptimized.CallOpts)
}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) LpPrice() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.LpPrice(&_TwocryptoOptimized.CallOpts)
}

// MaTime is a free data retrieval call binding the contract method 0x09c3da6a.
//
// Solidity: function ma_time() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) MaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "ma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaTime is a free data retrieval call binding the contract method 0x09c3da6a.
//
// Solidity: function ma_time() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) MaTime() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.MaTime(&_TwocryptoOptimized.CallOpts)
}

// MaTime is a free data retrieval call binding the contract method 0x09c3da6a.
//
// Solidity: function ma_time() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) MaTime() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.MaTime(&_TwocryptoOptimized.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) MidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) MidFee() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.MidFee(&_TwocryptoOptimized.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) MidFee() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.MidFee(&_TwocryptoOptimized.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Name() (string, error) {
	return _TwocryptoOptimized.Contract.Name(&_TwocryptoOptimized.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Name() (string, error) {
	return _TwocryptoOptimized.Contract.Name(&_TwocryptoOptimized.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.Nonces(&_TwocryptoOptimized.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _TwocryptoOptimized.Contract.Nonces(&_TwocryptoOptimized.CallOpts, arg0)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) OutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) OutFee() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.OutFee(&_TwocryptoOptimized.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) OutFee() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.OutFee(&_TwocryptoOptimized.CallOpts)
}

// PackedFeeParams is a free data retrieval call binding the contract method 0xe3616405.
//
// Solidity: function packed_fee_params() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) PackedFeeParams(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "packed_fee_params")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PackedFeeParams is a free data retrieval call binding the contract method 0xe3616405.
//
// Solidity: function packed_fee_params() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) PackedFeeParams() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.PackedFeeParams(&_TwocryptoOptimized.CallOpts)
}

// PackedFeeParams is a free data retrieval call binding the contract method 0xe3616405.
//
// Solidity: function packed_fee_params() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) PackedFeeParams() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.PackedFeeParams(&_TwocryptoOptimized.CallOpts)
}

// PackedRebalancingParams is a free data retrieval call binding the contract method 0x3dd65478.
//
// Solidity: function packed_rebalancing_params() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) PackedRebalancingParams(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "packed_rebalancing_params")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PackedRebalancingParams is a free data retrieval call binding the contract method 0x3dd65478.
//
// Solidity: function packed_rebalancing_params() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) PackedRebalancingParams() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.PackedRebalancingParams(&_TwocryptoOptimized.CallOpts)
}

// PackedRebalancingParams is a free data retrieval call binding the contract method 0x3dd65478.
//
// Solidity: function packed_rebalancing_params() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) PackedRebalancingParams() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.PackedRebalancingParams(&_TwocryptoOptimized.CallOpts)
}

// Precisions is a free data retrieval call binding the contract method 0x3620604b.
//
// Solidity: function precisions() view returns(uint256[2])
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Precisions(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "precisions")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// Precisions is a free data retrieval call binding the contract method 0x3620604b.
//
// Solidity: function precisions() view returns(uint256[2])
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Precisions() ([2]*big.Int, error) {
	return _TwocryptoOptimized.Contract.Precisions(&_TwocryptoOptimized.CallOpts)
}

// Precisions is a free data retrieval call binding the contract method 0x3620604b.
//
// Solidity: function precisions() view returns(uint256[2])
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Precisions() ([2]*big.Int, error) {
	return _TwocryptoOptimized.Contract.Precisions(&_TwocryptoOptimized.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) PriceOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "price_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) PriceOracle() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.PriceOracle(&_TwocryptoOptimized.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) PriceOracle() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.PriceOracle(&_TwocryptoOptimized.CallOpts)
}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) PriceScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "price_scale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) PriceScale() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.PriceScale(&_TwocryptoOptimized.CallOpts)
}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) PriceScale() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.PriceScale(&_TwocryptoOptimized.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Salt() ([32]byte, error) {
	return _TwocryptoOptimized.Contract.Salt(&_TwocryptoOptimized.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Salt() ([32]byte, error) {
	return _TwocryptoOptimized.Contract.Salt(&_TwocryptoOptimized.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Symbol() (string, error) {
	return _TwocryptoOptimized.Contract.Symbol(&_TwocryptoOptimized.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Symbol() (string, error) {
	return _TwocryptoOptimized.Contract.Symbol(&_TwocryptoOptimized.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) TotalSupply() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.TotalSupply(&_TwocryptoOptimized.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) TotalSupply() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.TotalSupply(&_TwocryptoOptimized.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Version() (string, error) {
	return _TwocryptoOptimized.Contract.Version(&_TwocryptoOptimized.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) Version() (string, error) {
	return _TwocryptoOptimized.Contract.Version(&_TwocryptoOptimized.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) VirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) VirtualPrice() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.VirtualPrice(&_TwocryptoOptimized.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) VirtualPrice() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.VirtualPrice(&_TwocryptoOptimized.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) XcpProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "xcp_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) XcpProfit() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.XcpProfit(&_TwocryptoOptimized.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) XcpProfit() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.XcpProfit(&_TwocryptoOptimized.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCaller) XcpProfitA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwocryptoOptimized.contract.Call(opts, &out, "xcp_profit_a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) XcpProfitA() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.XcpProfitA(&_TwocryptoOptimized.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedCallerSession) XcpProfitA() (*big.Int, error) {
	return _TwocryptoOptimized.Contract.XcpProfitA(&_TwocryptoOptimized.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.AddLiquidity(&_TwocryptoOptimized.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.AddLiquidity(&_TwocryptoOptimized.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) AddLiquidity0(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "add_liquidity0", amounts, min_mint_amount, receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) AddLiquidity0(amounts [2]*big.Int, min_mint_amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.AddLiquidity0(&_TwocryptoOptimized.TransactOpts, amounts, min_mint_amount, receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) AddLiquidity0(amounts [2]*big.Int, min_mint_amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.AddLiquidity0(&_TwocryptoOptimized.TransactOpts, amounts, min_mint_amount, receiver)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x6dbcf350.
//
// Solidity: function apply_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_time) returns()
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) ApplyNewParameters(opts *bind.TransactOpts, _new_mid_fee *big.Int, _new_out_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_time *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "apply_new_parameters", _new_mid_fee, _new_out_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_time)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x6dbcf350.
//
// Solidity: function apply_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_time) returns()
func (_TwocryptoOptimized *TwocryptoOptimizedSession) ApplyNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_time *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.ApplyNewParameters(&_TwocryptoOptimized.TransactOpts, _new_mid_fee, _new_out_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_time)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x6dbcf350.
//
// Solidity: function apply_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_time) returns()
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) ApplyNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_time *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.ApplyNewParameters(&_TwocryptoOptimized.TransactOpts, _new_mid_fee, _new_out_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_time)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.Approve(&_TwocryptoOptimized.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.Approve(&_TwocryptoOptimized.TransactOpts, _spender, _value)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.Exchange(&_TwocryptoOptimized.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.Exchange(&_TwocryptoOptimized.TransactOpts, i, j, dx, min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "exchange0", i, j, dx, min_dy, receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.Exchange0(&_TwocryptoOptimized.TransactOpts, i, j, dx, min_dy, receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.Exchange0(&_TwocryptoOptimized.TransactOpts, i, j, dx, min_dy, receiver)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x29b244bb.
//
// Solidity: function exchange_received(uint256 i, uint256 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) ExchangeReceived(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "exchange_received", i, j, dx, min_dy)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x29b244bb.
//
// Solidity: function exchange_received(uint256 i, uint256 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) ExchangeReceived(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.ExchangeReceived(&_TwocryptoOptimized.TransactOpts, i, j, dx, min_dy)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x29b244bb.
//
// Solidity: function exchange_received(uint256 i, uint256 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) ExchangeReceived(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.ExchangeReceived(&_TwocryptoOptimized.TransactOpts, i, j, dx, min_dy)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0x767691e7.
//
// Solidity: function exchange_received(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) ExchangeReceived0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "exchange_received0", i, j, dx, min_dy, receiver)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0x767691e7.
//
// Solidity: function exchange_received(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) ExchangeReceived0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.ExchangeReceived0(&_TwocryptoOptimized.TransactOpts, i, j, dx, min_dy, receiver)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0x767691e7.
//
// Solidity: function exchange_received(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) ExchangeReceived0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.ExchangeReceived0(&_TwocryptoOptimized.TransactOpts, i, j, dx, min_dy, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.Permit(&_TwocryptoOptimized.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.Permit(&_TwocryptoOptimized.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) RampAGamma(opts *bind.TransactOpts, future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "ramp_A_gamma", future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_TwocryptoOptimized *TwocryptoOptimizedSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.RampAGamma(&_TwocryptoOptimized.TransactOpts, future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.RampAGamma(&_TwocryptoOptimized.TransactOpts, future_A, future_gamma, future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns(uint256[2])
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns(uint256[2])
func (_TwocryptoOptimized *TwocryptoOptimizedSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.RemoveLiquidity(&_TwocryptoOptimized.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns(uint256[2])
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.RemoveLiquidity(&_TwocryptoOptimized.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, address receiver) returns(uint256[2])
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "remove_liquidity0", _amount, min_amounts, receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, address receiver) returns(uint256[2])
func (_TwocryptoOptimized *TwocryptoOptimizedSession) RemoveLiquidity0(_amount *big.Int, min_amounts [2]*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.RemoveLiquidity0(&_TwocryptoOptimized.TransactOpts, _amount, min_amounts, receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, address receiver) returns(uint256[2])
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) RemoveLiquidity0(_amount *big.Int, min_amounts [2]*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.RemoveLiquidity0(&_TwocryptoOptimized.TransactOpts, _amount, min_amounts, receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "remove_liquidity_one_coin", token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.RemoveLiquidityOneCoin(&_TwocryptoOptimized.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.RemoveLiquidityOneCoin(&_TwocryptoOptimized.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x0fbcee6e.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "remove_liquidity_one_coin0", token_amount, i, min_amount, receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x0fbcee6e.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) RemoveLiquidityOneCoin0(token_amount *big.Int, i *big.Int, min_amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.RemoveLiquidityOneCoin0(&_TwocryptoOptimized.TransactOpts, token_amount, i, min_amount, receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x0fbcee6e.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, address receiver) returns(uint256)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) RemoveLiquidityOneCoin0(token_amount *big.Int, i *big.Int, min_amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.RemoveLiquidityOneCoin0(&_TwocryptoOptimized.TransactOpts, token_amount, i, min_amount, receiver)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) StopRampAGamma(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "stop_ramp_A_gamma")
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_TwocryptoOptimized *TwocryptoOptimizedSession) StopRampAGamma() (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.StopRampAGamma(&_TwocryptoOptimized.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) StopRampAGamma() (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.StopRampAGamma(&_TwocryptoOptimized.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.Transfer(&_TwocryptoOptimized.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.Transfer(&_TwocryptoOptimized.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.TransferFrom(&_TwocryptoOptimized.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_TwocryptoOptimized *TwocryptoOptimizedTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TwocryptoOptimized.Contract.TransferFrom(&_TwocryptoOptimized.TransactOpts, _from, _to, _value)
}

// TwocryptoOptimizedAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedAddLiquidityIterator struct {
	Event *TwocryptoOptimizedAddLiquidity // Event containing the contract specifics and raw log

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
func (it *TwocryptoOptimizedAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwocryptoOptimizedAddLiquidity)
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
		it.Event = new(TwocryptoOptimizedAddLiquidity)
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
func (it *TwocryptoOptimizedAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwocryptoOptimizedAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwocryptoOptimizedAddLiquidity represents a AddLiquidity event raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedAddLiquidity struct {
	Provider         common.Address
	TokenAmounts     [2]*big.Int
	Fee              *big.Int
	TokenSupply      *big.Int
	PackedPriceScale *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x7196cbf63df1f2ec20638e683ebe51d18260be510592ee1e2efe3f3cfd4c33e9.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply, uint256 packed_price_scale)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*TwocryptoOptimizedAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedAddLiquidityIterator{contract: _TwocryptoOptimized.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x7196cbf63df1f2ec20638e683ebe51d18260be510592ee1e2efe3f3cfd4c33e9.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply, uint256 packed_price_scale)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *TwocryptoOptimizedAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwocryptoOptimizedAddLiquidity)
				if err := _TwocryptoOptimized.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x7196cbf63df1f2ec20638e683ebe51d18260be510592ee1e2efe3f3cfd4c33e9.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply, uint256 packed_price_scale)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) ParseAddLiquidity(log types.Log) (*TwocryptoOptimizedAddLiquidity, error) {
	event := new(TwocryptoOptimizedAddLiquidity)
	if err := _TwocryptoOptimized.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwocryptoOptimizedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedApprovalIterator struct {
	Event *TwocryptoOptimizedApproval // Event containing the contract specifics and raw log

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
func (it *TwocryptoOptimizedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwocryptoOptimizedApproval)
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
		it.Event = new(TwocryptoOptimizedApproval)
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
func (it *TwocryptoOptimizedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwocryptoOptimizedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwocryptoOptimizedApproval represents a Approval event raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TwocryptoOptimizedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedApprovalIterator{contract: _TwocryptoOptimized.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TwocryptoOptimizedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwocryptoOptimizedApproval)
				if err := _TwocryptoOptimized.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) ParseApproval(log types.Log) (*TwocryptoOptimizedApproval, error) {
	event := new(TwocryptoOptimizedApproval)
	if err := _TwocryptoOptimized.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwocryptoOptimizedClaimAdminFeeIterator is returned from FilterClaimAdminFee and is used to iterate over the raw logs and unpacked data for ClaimAdminFee events raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedClaimAdminFeeIterator struct {
	Event *TwocryptoOptimizedClaimAdminFee // Event containing the contract specifics and raw log

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
func (it *TwocryptoOptimizedClaimAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwocryptoOptimizedClaimAdminFee)
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
		it.Event = new(TwocryptoOptimizedClaimAdminFee)
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
func (it *TwocryptoOptimizedClaimAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwocryptoOptimizedClaimAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwocryptoOptimizedClaimAdminFee represents a ClaimAdminFee event raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedClaimAdminFee struct {
	Admin  common.Address
	Tokens [2]*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimAdminFee is a free log retrieval operation binding the contract event 0x3bbd5f2f4711532d6e9ee88dfdf2f1468e9a4c3ae5e14d2e1a67bf4242d008d0.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256[2] tokens)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) FilterClaimAdminFee(opts *bind.FilterOpts, admin []common.Address) (*TwocryptoOptimizedClaimAdminFeeIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.FilterLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedClaimAdminFeeIterator{contract: _TwocryptoOptimized.contract, event: "ClaimAdminFee", logs: logs, sub: sub}, nil
}

// WatchClaimAdminFee is a free log subscription operation binding the contract event 0x3bbd5f2f4711532d6e9ee88dfdf2f1468e9a4c3ae5e14d2e1a67bf4242d008d0.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256[2] tokens)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) WatchClaimAdminFee(opts *bind.WatchOpts, sink chan<- *TwocryptoOptimizedClaimAdminFee, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.WatchLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwocryptoOptimizedClaimAdminFee)
				if err := _TwocryptoOptimized.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
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

// ParseClaimAdminFee is a log parse operation binding the contract event 0x3bbd5f2f4711532d6e9ee88dfdf2f1468e9a4c3ae5e14d2e1a67bf4242d008d0.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256[2] tokens)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) ParseClaimAdminFee(log types.Log) (*TwocryptoOptimizedClaimAdminFee, error) {
	event := new(TwocryptoOptimizedClaimAdminFee)
	if err := _TwocryptoOptimized.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwocryptoOptimizedNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedNewParametersIterator struct {
	Event *TwocryptoOptimizedNewParameters // Event containing the contract specifics and raw log

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
func (it *TwocryptoOptimizedNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwocryptoOptimizedNewParameters)
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
		it.Event = new(TwocryptoOptimizedNewParameters)
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
func (it *TwocryptoOptimizedNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwocryptoOptimizedNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwocryptoOptimizedNewParameters represents a NewParameters event raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedNewParameters struct {
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaTime             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_time)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) FilterNewParameters(opts *bind.FilterOpts) (*TwocryptoOptimizedNewParametersIterator, error) {

	logs, sub, err := _TwocryptoOptimized.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedNewParametersIterator{contract: _TwocryptoOptimized.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_time)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *TwocryptoOptimizedNewParameters) (event.Subscription, error) {

	logs, sub, err := _TwocryptoOptimized.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwocryptoOptimizedNewParameters)
				if err := _TwocryptoOptimized.contract.UnpackLog(event, "NewParameters", log); err != nil {
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

// ParseNewParameters is a log parse operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_time)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) ParseNewParameters(log types.Log) (*TwocryptoOptimizedNewParameters, error) {
	event := new(TwocryptoOptimizedNewParameters)
	if err := _TwocryptoOptimized.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwocryptoOptimizedRampAgammaIterator is returned from FilterRampAgamma and is used to iterate over the raw logs and unpacked data for RampAgamma events raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedRampAgammaIterator struct {
	Event *TwocryptoOptimizedRampAgamma // Event containing the contract specifics and raw log

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
func (it *TwocryptoOptimizedRampAgammaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwocryptoOptimizedRampAgamma)
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
		it.Event = new(TwocryptoOptimizedRampAgamma)
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
func (it *TwocryptoOptimizedRampAgammaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwocryptoOptimizedRampAgammaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwocryptoOptimizedRampAgamma represents a RampAgamma event raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedRampAgamma struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialGamma *big.Int
	FutureGamma  *big.Int
	InitialTime  *big.Int
	FutureTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRampAgamma is a free log retrieval operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) FilterRampAgamma(opts *bind.FilterOpts) (*TwocryptoOptimizedRampAgammaIterator, error) {

	logs, sub, err := _TwocryptoOptimized.contract.FilterLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedRampAgammaIterator{contract: _TwocryptoOptimized.contract, event: "RampAgamma", logs: logs, sub: sub}, nil
}

// WatchRampAgamma is a free log subscription operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) WatchRampAgamma(opts *bind.WatchOpts, sink chan<- *TwocryptoOptimizedRampAgamma) (event.Subscription, error) {

	logs, sub, err := _TwocryptoOptimized.contract.WatchLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwocryptoOptimizedRampAgamma)
				if err := _TwocryptoOptimized.contract.UnpackLog(event, "RampAgamma", log); err != nil {
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

// ParseRampAgamma is a log parse operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) ParseRampAgamma(log types.Log) (*TwocryptoOptimizedRampAgamma, error) {
	event := new(TwocryptoOptimizedRampAgamma)
	if err := _TwocryptoOptimized.contract.UnpackLog(event, "RampAgamma", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwocryptoOptimizedRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedRemoveLiquidityIterator struct {
	Event *TwocryptoOptimizedRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *TwocryptoOptimizedRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwocryptoOptimizedRemoveLiquidity)
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
		it.Event = new(TwocryptoOptimizedRemoveLiquidity)
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
func (it *TwocryptoOptimizedRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwocryptoOptimizedRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwocryptoOptimizedRemoveLiquidity represents a RemoveLiquidity event raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*TwocryptoOptimizedRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedRemoveLiquidityIterator{contract: _TwocryptoOptimized.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *TwocryptoOptimizedRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwocryptoOptimizedRemoveLiquidity)
				if err := _TwocryptoOptimized.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) ParseRemoveLiquidity(log types.Log) (*TwocryptoOptimizedRemoveLiquidity, error) {
	event := new(TwocryptoOptimizedRemoveLiquidity)
	if err := _TwocryptoOptimized.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwocryptoOptimizedRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedRemoveLiquidityOneIterator struct {
	Event *TwocryptoOptimizedRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *TwocryptoOptimizedRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwocryptoOptimizedRemoveLiquidityOne)
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
		it.Event = new(TwocryptoOptimizedRemoveLiquidityOne)
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
func (it *TwocryptoOptimizedRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwocryptoOptimizedRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwocryptoOptimizedRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedRemoveLiquidityOne struct {
	Provider         common.Address
	TokenAmount      *big.Int
	CoinIndex        *big.Int
	CoinAmount       *big.Int
	ApproxFee        *big.Int
	PackedPriceScale *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0xe200e24d4a4c7cd367dd9befe394dc8a14e6d58c88ff5e2f512d65a9e0aa9c5c.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount, uint256 approx_fee, uint256 packed_price_scale)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*TwocryptoOptimizedRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedRemoveLiquidityOneIterator{contract: _TwocryptoOptimized.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0xe200e24d4a4c7cd367dd9befe394dc8a14e6d58c88ff5e2f512d65a9e0aa9c5c.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount, uint256 approx_fee, uint256 packed_price_scale)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *TwocryptoOptimizedRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwocryptoOptimizedRemoveLiquidityOne)
				if err := _TwocryptoOptimized.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0xe200e24d4a4c7cd367dd9befe394dc8a14e6d58c88ff5e2f512d65a9e0aa9c5c.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount, uint256 approx_fee, uint256 packed_price_scale)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) ParseRemoveLiquidityOne(log types.Log) (*TwocryptoOptimizedRemoveLiquidityOne, error) {
	event := new(TwocryptoOptimizedRemoveLiquidityOne)
	if err := _TwocryptoOptimized.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwocryptoOptimizedStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedStopRampAIterator struct {
	Event *TwocryptoOptimizedStopRampA // Event containing the contract specifics and raw log

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
func (it *TwocryptoOptimizedStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwocryptoOptimizedStopRampA)
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
		it.Event = new(TwocryptoOptimizedStopRampA)
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
func (it *TwocryptoOptimizedStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwocryptoOptimizedStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwocryptoOptimizedStopRampA represents a StopRampA event raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedStopRampA struct {
	CurrentA     *big.Int
	CurrentGamma *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) FilterStopRampA(opts *bind.FilterOpts) (*TwocryptoOptimizedStopRampAIterator, error) {

	logs, sub, err := _TwocryptoOptimized.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedStopRampAIterator{contract: _TwocryptoOptimized.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *TwocryptoOptimizedStopRampA) (event.Subscription, error) {

	logs, sub, err := _TwocryptoOptimized.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwocryptoOptimizedStopRampA)
				if err := _TwocryptoOptimized.contract.UnpackLog(event, "StopRampA", log); err != nil {
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

// ParseStopRampA is a log parse operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) ParseStopRampA(log types.Log) (*TwocryptoOptimizedStopRampA, error) {
	event := new(TwocryptoOptimizedStopRampA)
	if err := _TwocryptoOptimized.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwocryptoOptimizedTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedTokenExchangeIterator struct {
	Event *TwocryptoOptimizedTokenExchange // Event containing the contract specifics and raw log

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
func (it *TwocryptoOptimizedTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwocryptoOptimizedTokenExchange)
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
		it.Event = new(TwocryptoOptimizedTokenExchange)
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
func (it *TwocryptoOptimizedTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwocryptoOptimizedTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwocryptoOptimizedTokenExchange represents a TokenExchange event raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedTokenExchange struct {
	Buyer            common.Address
	SoldId           *big.Int
	TokensSold       *big.Int
	BoughtId         *big.Int
	TokensBought     *big.Int
	Fee              *big.Int
	PackedPriceScale *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0x143f1f8e861fbdeddd5b46e844b7d3ac7b86a122f36e8c463859ee6811b1f29c.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought, uint256 fee, uint256 packed_price_scale)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*TwocryptoOptimizedTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedTokenExchangeIterator{contract: _TwocryptoOptimized.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x143f1f8e861fbdeddd5b46e844b7d3ac7b86a122f36e8c463859ee6811b1f29c.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought, uint256 fee, uint256 packed_price_scale)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *TwocryptoOptimizedTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwocryptoOptimizedTokenExchange)
				if err := _TwocryptoOptimized.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0x143f1f8e861fbdeddd5b46e844b7d3ac7b86a122f36e8c463859ee6811b1f29c.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought, uint256 fee, uint256 packed_price_scale)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) ParseTokenExchange(log types.Log) (*TwocryptoOptimizedTokenExchange, error) {
	event := new(TwocryptoOptimizedTokenExchange)
	if err := _TwocryptoOptimized.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwocryptoOptimizedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedTransferIterator struct {
	Event *TwocryptoOptimizedTransfer // Event containing the contract specifics and raw log

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
func (it *TwocryptoOptimizedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwocryptoOptimizedTransfer)
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
		it.Event = new(TwocryptoOptimizedTransfer)
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
func (it *TwocryptoOptimizedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwocryptoOptimizedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwocryptoOptimizedTransfer represents a Transfer event raised by the TwocryptoOptimized contract.
type TwocryptoOptimizedTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*TwocryptoOptimizedTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &TwocryptoOptimizedTransferIterator{contract: _TwocryptoOptimized.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TwocryptoOptimizedTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _TwocryptoOptimized.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwocryptoOptimizedTransfer)
				if err := _TwocryptoOptimized.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_TwocryptoOptimized *TwocryptoOptimizedFilterer) ParseTransfer(log types.Log) (*TwocryptoOptimizedTransfer, error) {
	event := new(TwocryptoOptimizedTransfer)
	if err := _TwocryptoOptimized.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
