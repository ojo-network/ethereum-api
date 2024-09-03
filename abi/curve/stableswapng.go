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

// StableSwapNGMetaData contains all meta data concerning the StableSwapNG contract.
var StableSwapNGMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchangeUnderlying\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"name\":\"old_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"t\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyNewFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"offpeg_fee_multiplier\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetNewMATime\",\"inputs\":[{\"name\":\"ma_exp_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"D_ma_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_offpeg_fee_multiplier\",\"type\":\"uint256\"},{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"},{\"name\":\"_coins\",\"type\":\"address[]\"},{\"name\":\"_rate_multipliers\",\"type\":\"uint256[]\"},{\"name\":\"_asset_types\",\"type\":\"uint8[]\"},{\"name\":\"_method_ids\",\"type\":\"bytes4[]\"},{\"name\":\"_oracles\",\"type\":\"address[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_received\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_received\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_claim_admin_fees\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_price\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ema_price\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_p\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dx\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A_precise\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"stored_rates\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"dynamic_fee\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A\",\"inputs\":[{\"name\":\"_future_A\",\"type\":\"uint256\"},{\"name\":\"_future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_new_fee\",\"inputs\":[{\"name\":\"_new_fee\",\"type\":\"uint256\"},{\"name\":\"_new_offpeg_fee_multiplier\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_ma_exp_time\",\"inputs\":[{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"},{\"name\":\"_D_ma_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"N_COINS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"offpeg_fee_multiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_exp_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D_ma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_last_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"salt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]}]",
}

// StableSwapNGABI is the input ABI used to generate the binding from.
// Deprecated: Use StableSwapNGMetaData.ABI instead.
var StableSwapNGABI = StableSwapNGMetaData.ABI

// StableSwapNG is an auto generated Go binding around an Ethereum contract.
type StableSwapNG struct {
	StableSwapNGCaller     // Read-only binding to the contract
	StableSwapNGTransactor // Write-only binding to the contract
	StableSwapNGFilterer   // Log filterer for contract events
}

// StableSwapNGCaller is an auto generated read-only Go binding around an Ethereum contract.
type StableSwapNGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableSwapNGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StableSwapNGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableSwapNGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StableSwapNGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableSwapNGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StableSwapNGSession struct {
	Contract     *StableSwapNG     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StableSwapNGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StableSwapNGCallerSession struct {
	Contract *StableSwapNGCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StableSwapNGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StableSwapNGTransactorSession struct {
	Contract     *StableSwapNGTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StableSwapNGRaw is an auto generated low-level Go binding around an Ethereum contract.
type StableSwapNGRaw struct {
	Contract *StableSwapNG // Generic contract binding to access the raw methods on
}

// StableSwapNGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StableSwapNGCallerRaw struct {
	Contract *StableSwapNGCaller // Generic read-only contract binding to access the raw methods on
}

// StableSwapNGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StableSwapNGTransactorRaw struct {
	Contract *StableSwapNGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStableSwapNG creates a new instance of StableSwapNG, bound to a specific deployed contract.
func NewStableSwapNG(address common.Address, backend bind.ContractBackend) (*StableSwapNG, error) {
	contract, err := bindStableSwapNG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StableSwapNG{StableSwapNGCaller: StableSwapNGCaller{contract: contract}, StableSwapNGTransactor: StableSwapNGTransactor{contract: contract}, StableSwapNGFilterer: StableSwapNGFilterer{contract: contract}}, nil
}

// NewStableSwapNGCaller creates a new read-only instance of StableSwapNG, bound to a specific deployed contract.
func NewStableSwapNGCaller(address common.Address, caller bind.ContractCaller) (*StableSwapNGCaller, error) {
	contract, err := bindStableSwapNG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGCaller{contract: contract}, nil
}

// NewStableSwapNGTransactor creates a new write-only instance of StableSwapNG, bound to a specific deployed contract.
func NewStableSwapNGTransactor(address common.Address, transactor bind.ContractTransactor) (*StableSwapNGTransactor, error) {
	contract, err := bindStableSwapNG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGTransactor{contract: contract}, nil
}

// NewStableSwapNGFilterer creates a new log filterer instance of StableSwapNG, bound to a specific deployed contract.
func NewStableSwapNGFilterer(address common.Address, filterer bind.ContractFilterer) (*StableSwapNGFilterer, error) {
	contract, err := bindStableSwapNG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGFilterer{contract: contract}, nil
}

// bindStableSwapNG binds a generic wrapper to an already deployed contract.
func bindStableSwapNG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StableSwapNGMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableSwapNG *StableSwapNGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StableSwapNG.Contract.StableSwapNGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableSwapNG *StableSwapNGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableSwapNG.Contract.StableSwapNGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableSwapNG *StableSwapNGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableSwapNG.Contract.StableSwapNGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableSwapNG *StableSwapNGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StableSwapNG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableSwapNG *StableSwapNGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableSwapNG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableSwapNG *StableSwapNGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableSwapNG.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) A() (*big.Int, error) {
	return _StableSwapNG.Contract.A(&_StableSwapNG.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) A() (*big.Int, error) {
	return _StableSwapNG.Contract.A(&_StableSwapNG.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) APrecise() (*big.Int, error) {
	return _StableSwapNG.Contract.APrecise(&_StableSwapNG.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) APrecise() (*big.Int, error) {
	return _StableSwapNG.Contract.APrecise(&_StableSwapNG.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StableSwapNG *StableSwapNGCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StableSwapNG *StableSwapNGSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StableSwapNG.Contract.DOMAINSEPARATOR(&_StableSwapNG.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StableSwapNG *StableSwapNGCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StableSwapNG.Contract.DOMAINSEPARATOR(&_StableSwapNG.CallOpts)
}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) DMaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "D_ma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) DMaTime() (*big.Int, error) {
	return _StableSwapNG.Contract.DMaTime(&_StableSwapNG.CallOpts)
}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) DMaTime() (*big.Int, error) {
	return _StableSwapNG.Contract.DMaTime(&_StableSwapNG.CallOpts)
}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) DOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "D_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) DOracle() (*big.Int, error) {
	return _StableSwapNG.Contract.DOracle(&_StableSwapNG.CallOpts)
}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) DOracle() (*big.Int, error) {
	return _StableSwapNG.Contract.DOracle(&_StableSwapNG.CallOpts)
}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) NCOINS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "N_COINS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) NCOINS() (*big.Int, error) {
	return _StableSwapNG.Contract.NCOINS(&_StableSwapNG.CallOpts)
}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) NCOINS() (*big.Int, error) {
	return _StableSwapNG.Contract.NCOINS(&_StableSwapNG.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) AdminBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "admin_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.AdminBalances(&_StableSwapNG.CallOpts, arg0)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.AdminBalances(&_StableSwapNG.CallOpts, arg0)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) AdminFee() (*big.Int, error) {
	return _StableSwapNG.Contract.AdminFee(&_StableSwapNG.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) AdminFee() (*big.Int, error) {
	return _StableSwapNG.Contract.AdminFee(&_StableSwapNG.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StableSwapNG.Contract.Allowance(&_StableSwapNG.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StableSwapNG.Contract.Allowance(&_StableSwapNG.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _StableSwapNG.Contract.BalanceOf(&_StableSwapNG.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _StableSwapNG.Contract.BalanceOf(&_StableSwapNG.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) Balances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) Balances(i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.Balances(&_StableSwapNG.CallOpts, i)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) Balances(i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.Balances(&_StableSwapNG.CallOpts, i)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) CalcTokenAmount(_amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	return _StableSwapNG.Contract.CalcTokenAmount(&_StableSwapNG.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) CalcTokenAmount(_amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	return _StableSwapNG.Contract.CalcTokenAmount(&_StableSwapNG.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "calc_withdraw_one_coin", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.CalcWithdrawOneCoin(&_StableSwapNG.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.CalcWithdrawOneCoin(&_StableSwapNG.CallOpts, _burn_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_StableSwapNG *StableSwapNGCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_StableSwapNG *StableSwapNGSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _StableSwapNG.Contract.Coins(&_StableSwapNG.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_StableSwapNG *StableSwapNGCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _StableSwapNG.Contract.Coins(&_StableSwapNG.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StableSwapNG *StableSwapNGCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StableSwapNG *StableSwapNGSession) Decimals() (uint8, error) {
	return _StableSwapNG.Contract.Decimals(&_StableSwapNG.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StableSwapNG *StableSwapNGCallerSession) Decimals() (uint8, error) {
	return _StableSwapNG.Contract.Decimals(&_StableSwapNG.CallOpts)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) DynamicFee(opts *bind.CallOpts, i *big.Int, j *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "dynamic_fee", i, j)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.DynamicFee(&_StableSwapNG.CallOpts, i, j)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.DynamicFee(&_StableSwapNG.CallOpts, i, j)
}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) EmaPrice(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "ema_price", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) EmaPrice(i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.EmaPrice(&_StableSwapNG.CallOpts, i)
}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) EmaPrice(i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.EmaPrice(&_StableSwapNG.CallOpts, i)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) Fee() (*big.Int, error) {
	return _StableSwapNG.Contract.Fee(&_StableSwapNG.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) Fee() (*big.Int, error) {
	return _StableSwapNG.Contract.Fee(&_StableSwapNG.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) FutureA() (*big.Int, error) {
	return _StableSwapNG.Contract.FutureA(&_StableSwapNG.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) FutureA() (*big.Int, error) {
	return _StableSwapNG.Contract.FutureA(&_StableSwapNG.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) FutureATime() (*big.Int, error) {
	return _StableSwapNG.Contract.FutureATime(&_StableSwapNG.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) FutureATime() (*big.Int, error) {
	return _StableSwapNG.Contract.FutureATime(&_StableSwapNG.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_StableSwapNG *StableSwapNGCaller) GetBalances(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "get_balances")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_StableSwapNG *StableSwapNGSession) GetBalances() ([]*big.Int, error) {
	return _StableSwapNG.Contract.GetBalances(&_StableSwapNG.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_StableSwapNG *StableSwapNGCallerSession) GetBalances() ([]*big.Int, error) {
	return _StableSwapNG.Contract.GetBalances(&_StableSwapNG.CallOpts)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) GetDx(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "get_dx", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.GetDx(&_StableSwapNG.CallOpts, i, j, dy)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.GetDx(&_StableSwapNG.CallOpts, i, j, dy)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.GetDy(&_StableSwapNG.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.GetDy(&_StableSwapNG.CallOpts, i, j, dx)
}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) GetP(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "get_p", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) GetP(i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.GetP(&_StableSwapNG.CallOpts, i)
}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) GetP(i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.GetP(&_StableSwapNG.CallOpts, i)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) GetVirtualPrice() (*big.Int, error) {
	return _StableSwapNG.Contract.GetVirtualPrice(&_StableSwapNG.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _StableSwapNG.Contract.GetVirtualPrice(&_StableSwapNG.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) InitialA() (*big.Int, error) {
	return _StableSwapNG.Contract.InitialA(&_StableSwapNG.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) InitialA() (*big.Int, error) {
	return _StableSwapNG.Contract.InitialA(&_StableSwapNG.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) InitialATime() (*big.Int, error) {
	return _StableSwapNG.Contract.InitialATime(&_StableSwapNG.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) InitialATime() (*big.Int, error) {
	return _StableSwapNG.Contract.InitialATime(&_StableSwapNG.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) LastPrice(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "last_price", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) LastPrice(i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.LastPrice(&_StableSwapNG.CallOpts, i)
}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) LastPrice(i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.LastPrice(&_StableSwapNG.CallOpts, i)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) MaExpTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "ma_exp_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) MaExpTime() (*big.Int, error) {
	return _StableSwapNG.Contract.MaExpTime(&_StableSwapNG.CallOpts)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) MaExpTime() (*big.Int, error) {
	return _StableSwapNG.Contract.MaExpTime(&_StableSwapNG.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) MaLastTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "ma_last_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) MaLastTime() (*big.Int, error) {
	return _StableSwapNG.Contract.MaLastTime(&_StableSwapNG.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) MaLastTime() (*big.Int, error) {
	return _StableSwapNG.Contract.MaLastTime(&_StableSwapNG.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StableSwapNG *StableSwapNGCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StableSwapNG *StableSwapNGSession) Name() (string, error) {
	return _StableSwapNG.Contract.Name(&_StableSwapNG.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StableSwapNG *StableSwapNGCallerSession) Name() (string, error) {
	return _StableSwapNG.Contract.Name(&_StableSwapNG.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _StableSwapNG.Contract.Nonces(&_StableSwapNG.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _StableSwapNG.Contract.Nonces(&_StableSwapNG.CallOpts, arg0)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) OffpegFeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "offpeg_fee_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _StableSwapNG.Contract.OffpegFeeMultiplier(&_StableSwapNG.CallOpts)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _StableSwapNG.Contract.OffpegFeeMultiplier(&_StableSwapNG.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) PriceOracle(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "price_oracle", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) PriceOracle(i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.PriceOracle(&_StableSwapNG.CallOpts, i)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) PriceOracle(i *big.Int) (*big.Int, error) {
	return _StableSwapNG.Contract.PriceOracle(&_StableSwapNG.CallOpts, i)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_StableSwapNG *StableSwapNGCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_StableSwapNG *StableSwapNGSession) Salt() ([32]byte, error) {
	return _StableSwapNG.Contract.Salt(&_StableSwapNG.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_StableSwapNG *StableSwapNGCallerSession) Salt() ([32]byte, error) {
	return _StableSwapNG.Contract.Salt(&_StableSwapNG.CallOpts)
}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_StableSwapNG *StableSwapNGCaller) StoredRates(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "stored_rates")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_StableSwapNG *StableSwapNGSession) StoredRates() ([]*big.Int, error) {
	return _StableSwapNG.Contract.StoredRates(&_StableSwapNG.CallOpts)
}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_StableSwapNG *StableSwapNGCallerSession) StoredRates() ([]*big.Int, error) {
	return _StableSwapNG.Contract.StoredRates(&_StableSwapNG.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StableSwapNG *StableSwapNGCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StableSwapNG *StableSwapNGSession) Symbol() (string, error) {
	return _StableSwapNG.Contract.Symbol(&_StableSwapNG.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StableSwapNG *StableSwapNGCallerSession) Symbol() (string, error) {
	return _StableSwapNG.Contract.Symbol(&_StableSwapNG.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StableSwapNG *StableSwapNGCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StableSwapNG *StableSwapNGSession) TotalSupply() (*big.Int, error) {
	return _StableSwapNG.Contract.TotalSupply(&_StableSwapNG.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StableSwapNG *StableSwapNGCallerSession) TotalSupply() (*big.Int, error) {
	return _StableSwapNG.Contract.TotalSupply(&_StableSwapNG.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StableSwapNG *StableSwapNGCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StableSwapNG.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StableSwapNG *StableSwapNGSession) Version() (string, error) {
	return _StableSwapNG.Contract.Version(&_StableSwapNG.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StableSwapNG *StableSwapNGCallerSession) Version() (string, error) {
	return _StableSwapNG.Contract.Version(&_StableSwapNG.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactor) AddLiquidity(opts *bind.TransactOpts, _amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "add_liquidity", _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_StableSwapNG *StableSwapNGSession) AddLiquidity(_amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.AddLiquidity(&_StableSwapNG.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactorSession) AddLiquidity(_amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.AddLiquidity(&_StableSwapNG.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactor) AddLiquidity0(opts *bind.TransactOpts, _amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "add_liquidity0", _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGSession) AddLiquidity0(_amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.AddLiquidity0(&_StableSwapNG.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactorSession) AddLiquidity0(_amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.AddLiquidity0(&_StableSwapNG.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StableSwapNG *StableSwapNGTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StableSwapNG *StableSwapNGSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.Approve(&_StableSwapNG.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StableSwapNG *StableSwapNGTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.Approve(&_StableSwapNG.TransactOpts, _spender, _value)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "exchange", i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_StableSwapNG *StableSwapNGSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.Exchange(&_StableSwapNG.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactorSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.Exchange(&_StableSwapNG.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "exchange0", i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.Exchange0(&_StableSwapNG.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactorSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.Exchange0(&_StableSwapNG.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactor) ExchangeReceived(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "exchange_received", i, j, _dx, _min_dy)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_StableSwapNG *StableSwapNGSession) ExchangeReceived(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.ExchangeReceived(&_StableSwapNG.TransactOpts, i, j, _dx, _min_dy)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactorSession) ExchangeReceived(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.ExchangeReceived(&_StableSwapNG.TransactOpts, i, j, _dx, _min_dy)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactor) ExchangeReceived0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "exchange_received0", i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGSession) ExchangeReceived0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.ExchangeReceived0(&_StableSwapNG.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactorSession) ExchangeReceived0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.ExchangeReceived0(&_StableSwapNG.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_StableSwapNG *StableSwapNGTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_StableSwapNG *StableSwapNGSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StableSwapNG.Contract.Permit(&_StableSwapNG.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_StableSwapNG *StableSwapNGTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StableSwapNG.Contract.Permit(&_StableSwapNG.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_StableSwapNG *StableSwapNGTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_StableSwapNG *StableSwapNGSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RampA(&_StableSwapNG.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_StableSwapNG *StableSwapNGTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RampA(&_StableSwapNG.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_StableSwapNG *StableSwapNGTransactor) RemoveLiquidity(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "remove_liquidity", _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_StableSwapNG *StableSwapNGSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidity(&_StableSwapNG.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_StableSwapNG *StableSwapNGTransactorSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidity(&_StableSwapNG.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_StableSwapNG *StableSwapNGTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "remove_liquidity0", _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_StableSwapNG *StableSwapNGSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidity0(&_StableSwapNG.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_StableSwapNG *StableSwapNGTransactorSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidity0(&_StableSwapNG.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_StableSwapNG *StableSwapNGTransactor) RemoveLiquidity1(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "remove_liquidity1", _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_StableSwapNG *StableSwapNGSession) RemoveLiquidity1(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidity1(&_StableSwapNG.TransactOpts, _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_StableSwapNG *StableSwapNGTransactorSession) RemoveLiquidity1(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidity1(&_StableSwapNG.TransactOpts, _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, _amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "remove_liquidity_imbalance", _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_StableSwapNG *StableSwapNGSession) RemoveLiquidityImbalance(_amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidityImbalance(&_StableSwapNG.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactorSession) RemoveLiquidityImbalance(_amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidityImbalance(&_StableSwapNG.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactor) RemoveLiquidityImbalance0(opts *bind.TransactOpts, _amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "remove_liquidity_imbalance0", _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGSession) RemoveLiquidityImbalance0(_amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidityImbalance0(&_StableSwapNG.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactorSession) RemoveLiquidityImbalance0(_amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidityImbalance0(&_StableSwapNG.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "remove_liquidity_one_coin", _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_StableSwapNG *StableSwapNGSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidityOneCoin(&_StableSwapNG.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactorSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidityOneCoin(&_StableSwapNG.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "remove_liquidity_one_coin0", _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidityOneCoin0(&_StableSwapNG.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_StableSwapNG *StableSwapNGTransactorSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _StableSwapNG.Contract.RemoveLiquidityOneCoin0(&_StableSwapNG.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_StableSwapNG *StableSwapNGTransactor) SetMaExpTime(opts *bind.TransactOpts, _ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "set_ma_exp_time", _ma_exp_time, _D_ma_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_StableSwapNG *StableSwapNGSession) SetMaExpTime(_ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.SetMaExpTime(&_StableSwapNG.TransactOpts, _ma_exp_time, _D_ma_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_StableSwapNG *StableSwapNGTransactorSession) SetMaExpTime(_ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.SetMaExpTime(&_StableSwapNG.TransactOpts, _ma_exp_time, _D_ma_time)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_StableSwapNG *StableSwapNGTransactor) SetNewFee(opts *bind.TransactOpts, _new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "set_new_fee", _new_fee, _new_offpeg_fee_multiplier)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_StableSwapNG *StableSwapNGSession) SetNewFee(_new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.SetNewFee(&_StableSwapNG.TransactOpts, _new_fee, _new_offpeg_fee_multiplier)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_StableSwapNG *StableSwapNGTransactorSession) SetNewFee(_new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.SetNewFee(&_StableSwapNG.TransactOpts, _new_fee, _new_offpeg_fee_multiplier)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_StableSwapNG *StableSwapNGTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_StableSwapNG *StableSwapNGSession) StopRampA() (*types.Transaction, error) {
	return _StableSwapNG.Contract.StopRampA(&_StableSwapNG.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_StableSwapNG *StableSwapNGTransactorSession) StopRampA() (*types.Transaction, error) {
	return _StableSwapNG.Contract.StopRampA(&_StableSwapNG.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StableSwapNG *StableSwapNGTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StableSwapNG *StableSwapNGSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.Transfer(&_StableSwapNG.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StableSwapNG *StableSwapNGTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.Transfer(&_StableSwapNG.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StableSwapNG *StableSwapNGTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StableSwapNG *StableSwapNGSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.TransferFrom(&_StableSwapNG.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StableSwapNG *StableSwapNGTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StableSwapNG.Contract.TransferFrom(&_StableSwapNG.TransactOpts, _from, _to, _value)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_StableSwapNG *StableSwapNGTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableSwapNG.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_StableSwapNG *StableSwapNGSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _StableSwapNG.Contract.WithdrawAdminFees(&_StableSwapNG.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_StableSwapNG *StableSwapNGTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _StableSwapNG.Contract.WithdrawAdminFees(&_StableSwapNG.TransactOpts)
}

// StableSwapNGAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the StableSwapNG contract.
type StableSwapNGAddLiquidityIterator struct {
	Event *StableSwapNGAddLiquidity // Event containing the contract specifics and raw log

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
func (it *StableSwapNGAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGAddLiquidity)
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
		it.Event = new(StableSwapNGAddLiquidity)
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
func (it *StableSwapNGAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGAddLiquidity represents a AddLiquidity event raised by the StableSwapNG contract.
type StableSwapNGAddLiquidity struct {
	Provider     common.Address
	TokenAmounts []*big.Int
	Fees         []*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNGAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGAddLiquidityIterator{contract: _StableSwapNG.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *StableSwapNGAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGAddLiquidity)
				if err := _StableSwapNG.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) ParseAddLiquidity(log types.Log) (*StableSwapNGAddLiquidity, error) {
	event := new(StableSwapNGAddLiquidity)
	if err := _StableSwapNG.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGApplyNewFeeIterator is returned from FilterApplyNewFee and is used to iterate over the raw logs and unpacked data for ApplyNewFee events raised by the StableSwapNG contract.
type StableSwapNGApplyNewFeeIterator struct {
	Event *StableSwapNGApplyNewFee // Event containing the contract specifics and raw log

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
func (it *StableSwapNGApplyNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGApplyNewFee)
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
		it.Event = new(StableSwapNGApplyNewFee)
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
func (it *StableSwapNGApplyNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGApplyNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGApplyNewFee represents a ApplyNewFee event raised by the StableSwapNG contract.
type StableSwapNGApplyNewFee struct {
	Fee                 *big.Int
	OffpegFeeMultiplier *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterApplyNewFee is a free log retrieval operation binding the contract event 0x750d10a7f37466ce785ee6bcb604aac543358db42afbcc332a3c12a49c80bf6d.
//
// Solidity: event ApplyNewFee(uint256 fee, uint256 offpeg_fee_multiplier)
func (_StableSwapNG *StableSwapNGFilterer) FilterApplyNewFee(opts *bind.FilterOpts) (*StableSwapNGApplyNewFeeIterator, error) {

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return &StableSwapNGApplyNewFeeIterator{contract: _StableSwapNG.contract, event: "ApplyNewFee", logs: logs, sub: sub}, nil
}

// WatchApplyNewFee is a free log subscription operation binding the contract event 0x750d10a7f37466ce785ee6bcb604aac543358db42afbcc332a3c12a49c80bf6d.
//
// Solidity: event ApplyNewFee(uint256 fee, uint256 offpeg_fee_multiplier)
func (_StableSwapNG *StableSwapNGFilterer) WatchApplyNewFee(opts *bind.WatchOpts, sink chan<- *StableSwapNGApplyNewFee) (event.Subscription, error) {

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGApplyNewFee)
				if err := _StableSwapNG.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
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

// ParseApplyNewFee is a log parse operation binding the contract event 0x750d10a7f37466ce785ee6bcb604aac543358db42afbcc332a3c12a49c80bf6d.
//
// Solidity: event ApplyNewFee(uint256 fee, uint256 offpeg_fee_multiplier)
func (_StableSwapNG *StableSwapNGFilterer) ParseApplyNewFee(log types.Log) (*StableSwapNGApplyNewFee, error) {
	event := new(StableSwapNGApplyNewFee)
	if err := _StableSwapNG.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StableSwapNG contract.
type StableSwapNGApprovalIterator struct {
	Event *StableSwapNGApproval // Event containing the contract specifics and raw log

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
func (it *StableSwapNGApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGApproval)
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
		it.Event = new(StableSwapNGApproval)
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
func (it *StableSwapNGApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGApproval represents a Approval event raised by the StableSwapNG contract.
type StableSwapNGApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableSwapNG *StableSwapNGFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StableSwapNGApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGApprovalIterator{contract: _StableSwapNG.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableSwapNG *StableSwapNGFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StableSwapNGApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGApproval)
				if err := _StableSwapNG.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StableSwapNG *StableSwapNGFilterer) ParseApproval(log types.Log) (*StableSwapNGApproval, error) {
	event := new(StableSwapNGApproval)
	if err := _StableSwapNG.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the StableSwapNG contract.
type StableSwapNGRampAIterator struct {
	Event *StableSwapNGRampA // Event containing the contract specifics and raw log

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
func (it *StableSwapNGRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGRampA)
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
		it.Event = new(StableSwapNGRampA)
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
func (it *StableSwapNGRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGRampA represents a RampA event raised by the StableSwapNG contract.
type StableSwapNGRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_StableSwapNG *StableSwapNGFilterer) FilterRampA(opts *bind.FilterOpts) (*StableSwapNGRampAIterator, error) {

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &StableSwapNGRampAIterator{contract: _StableSwapNG.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_StableSwapNG *StableSwapNGFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *StableSwapNGRampA) (event.Subscription, error) {

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGRampA)
				if err := _StableSwapNG.contract.UnpackLog(event, "RampA", log); err != nil {
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

// ParseRampA is a log parse operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_StableSwapNG *StableSwapNGFilterer) ParseRampA(log types.Log) (*StableSwapNGRampA, error) {
	event := new(StableSwapNGRampA)
	if err := _StableSwapNG.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the StableSwapNG contract.
type StableSwapNGRemoveLiquidityIterator struct {
	Event *StableSwapNGRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *StableSwapNGRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGRemoveLiquidity)
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
		it.Event = new(StableSwapNGRemoveLiquidity)
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
func (it *StableSwapNGRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGRemoveLiquidity represents a RemoveLiquidity event raised by the StableSwapNG contract.
type StableSwapNGRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts []*big.Int
	Fees         []*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x347ad828e58cbe534d8f6b67985d791360756b18f0d95fd9f197a66cc46480ea.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNGRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGRemoveLiquidityIterator{contract: _StableSwapNG.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x347ad828e58cbe534d8f6b67985d791360756b18f0d95fd9f197a66cc46480ea.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *StableSwapNGRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGRemoveLiquidity)
				if err := _StableSwapNG.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x347ad828e58cbe534d8f6b67985d791360756b18f0d95fd9f197a66cc46480ea.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) ParseRemoveLiquidity(log types.Log) (*StableSwapNGRemoveLiquidity, error) {
	event := new(StableSwapNGRemoveLiquidity)
	if err := _StableSwapNG.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the StableSwapNG contract.
type StableSwapNGRemoveLiquidityImbalanceIterator struct {
	Event *StableSwapNGRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *StableSwapNGRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGRemoveLiquidityImbalance)
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
		it.Event = new(StableSwapNGRemoveLiquidityImbalance)
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
func (it *StableSwapNGRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the StableSwapNG contract.
type StableSwapNGRemoveLiquidityImbalance struct {
	Provider     common.Address
	TokenAmounts []*big.Int
	Fees         []*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNGRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGRemoveLiquidityImbalanceIterator{contract: _StableSwapNG.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *StableSwapNGRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGRemoveLiquidityImbalance)
				if err := _StableSwapNG.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*StableSwapNGRemoveLiquidityImbalance, error) {
	event := new(StableSwapNGRemoveLiquidityImbalance)
	if err := _StableSwapNG.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the StableSwapNG contract.
type StableSwapNGRemoveLiquidityOneIterator struct {
	Event *StableSwapNGRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *StableSwapNGRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGRemoveLiquidityOne)
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
		it.Event = new(StableSwapNGRemoveLiquidityOne)
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
func (it *StableSwapNGRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the StableSwapNG contract.
type StableSwapNGRemoveLiquidityOne struct {
	Provider    common.Address
	TokenId     *big.Int
	TokenAmount *big.Int
	CoinAmount  *big.Int
	TokenSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x6f48129db1f37ccb9cc5dd7e119cb32750cabdf75b48375d730d26ce3659bbe1.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, int128 token_id, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNGRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGRemoveLiquidityOneIterator{contract: _StableSwapNG.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x6f48129db1f37ccb9cc5dd7e119cb32750cabdf75b48375d730d26ce3659bbe1.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, int128 token_id, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *StableSwapNGRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGRemoveLiquidityOne)
				if err := _StableSwapNG.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x6f48129db1f37ccb9cc5dd7e119cb32750cabdf75b48375d730d26ce3659bbe1.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, int128 token_id, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_StableSwapNG *StableSwapNGFilterer) ParseRemoveLiquidityOne(log types.Log) (*StableSwapNGRemoveLiquidityOne, error) {
	event := new(StableSwapNGRemoveLiquidityOne)
	if err := _StableSwapNG.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGSetNewMATimeIterator is returned from FilterSetNewMATime and is used to iterate over the raw logs and unpacked data for SetNewMATime events raised by the StableSwapNG contract.
type StableSwapNGSetNewMATimeIterator struct {
	Event *StableSwapNGSetNewMATime // Event containing the contract specifics and raw log

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
func (it *StableSwapNGSetNewMATimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGSetNewMATime)
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
		it.Event = new(StableSwapNGSetNewMATime)
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
func (it *StableSwapNGSetNewMATimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGSetNewMATimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGSetNewMATime represents a SetNewMATime event raised by the StableSwapNG contract.
type StableSwapNGSetNewMATime struct {
	MaExpTime *big.Int
	DMaTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetNewMATime is a free log retrieval operation binding the contract event 0x68dc4e067dff1862b896b7a0faf55f97df1a60d0aaa79481b69d675f2026a28c.
//
// Solidity: event SetNewMATime(uint256 ma_exp_time, uint256 D_ma_time)
func (_StableSwapNG *StableSwapNGFilterer) FilterSetNewMATime(opts *bind.FilterOpts) (*StableSwapNGSetNewMATimeIterator, error) {

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "SetNewMATime")
	if err != nil {
		return nil, err
	}
	return &StableSwapNGSetNewMATimeIterator{contract: _StableSwapNG.contract, event: "SetNewMATime", logs: logs, sub: sub}, nil
}

// WatchSetNewMATime is a free log subscription operation binding the contract event 0x68dc4e067dff1862b896b7a0faf55f97df1a60d0aaa79481b69d675f2026a28c.
//
// Solidity: event SetNewMATime(uint256 ma_exp_time, uint256 D_ma_time)
func (_StableSwapNG *StableSwapNGFilterer) WatchSetNewMATime(opts *bind.WatchOpts, sink chan<- *StableSwapNGSetNewMATime) (event.Subscription, error) {

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "SetNewMATime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGSetNewMATime)
				if err := _StableSwapNG.contract.UnpackLog(event, "SetNewMATime", log); err != nil {
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

// ParseSetNewMATime is a log parse operation binding the contract event 0x68dc4e067dff1862b896b7a0faf55f97df1a60d0aaa79481b69d675f2026a28c.
//
// Solidity: event SetNewMATime(uint256 ma_exp_time, uint256 D_ma_time)
func (_StableSwapNG *StableSwapNGFilterer) ParseSetNewMATime(log types.Log) (*StableSwapNGSetNewMATime, error) {
	event := new(StableSwapNGSetNewMATime)
	if err := _StableSwapNG.contract.UnpackLog(event, "SetNewMATime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the StableSwapNG contract.
type StableSwapNGStopRampAIterator struct {
	Event *StableSwapNGStopRampA // Event containing the contract specifics and raw log

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
func (it *StableSwapNGStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGStopRampA)
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
		it.Event = new(StableSwapNGStopRampA)
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
func (it *StableSwapNGStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGStopRampA represents a StopRampA event raised by the StableSwapNG contract.
type StableSwapNGStopRampA struct {
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_StableSwapNG *StableSwapNGFilterer) FilterStopRampA(opts *bind.FilterOpts) (*StableSwapNGStopRampAIterator, error) {

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &StableSwapNGStopRampAIterator{contract: _StableSwapNG.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_StableSwapNG *StableSwapNGFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *StableSwapNGStopRampA) (event.Subscription, error) {

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGStopRampA)
				if err := _StableSwapNG.contract.UnpackLog(event, "StopRampA", log); err != nil {
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

// ParseStopRampA is a log parse operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_StableSwapNG *StableSwapNGFilterer) ParseStopRampA(log types.Log) (*StableSwapNGStopRampA, error) {
	event := new(StableSwapNGStopRampA)
	if err := _StableSwapNG.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the StableSwapNG contract.
type StableSwapNGTokenExchangeIterator struct {
	Event *StableSwapNGTokenExchange // Event containing the contract specifics and raw log

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
func (it *StableSwapNGTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGTokenExchange)
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
		it.Event = new(StableSwapNGTokenExchange)
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
func (it *StableSwapNGTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGTokenExchange represents a TokenExchange event raised by the StableSwapNG contract.
type StableSwapNGTokenExchange struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_StableSwapNG *StableSwapNGFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*StableSwapNGTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGTokenExchangeIterator{contract: _StableSwapNG.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_StableSwapNG *StableSwapNGFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *StableSwapNGTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGTokenExchange)
				if err := _StableSwapNG.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_StableSwapNG *StableSwapNGFilterer) ParseTokenExchange(log types.Log) (*StableSwapNGTokenExchange, error) {
	event := new(StableSwapNGTokenExchange)
	if err := _StableSwapNG.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGTokenExchangeUnderlyingIterator is returned from FilterTokenExchangeUnderlying and is used to iterate over the raw logs and unpacked data for TokenExchangeUnderlying events raised by the StableSwapNG contract.
type StableSwapNGTokenExchangeUnderlyingIterator struct {
	Event *StableSwapNGTokenExchangeUnderlying // Event containing the contract specifics and raw log

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
func (it *StableSwapNGTokenExchangeUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGTokenExchangeUnderlying)
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
		it.Event = new(StableSwapNGTokenExchangeUnderlying)
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
func (it *StableSwapNGTokenExchangeUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGTokenExchangeUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGTokenExchangeUnderlying represents a TokenExchangeUnderlying event raised by the StableSwapNG contract.
type StableSwapNGTokenExchangeUnderlying struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangeUnderlying is a free log retrieval operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_StableSwapNG *StableSwapNGFilterer) FilterTokenExchangeUnderlying(opts *bind.FilterOpts, buyer []common.Address) (*StableSwapNGTokenExchangeUnderlyingIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGTokenExchangeUnderlyingIterator{contract: _StableSwapNG.contract, event: "TokenExchangeUnderlying", logs: logs, sub: sub}, nil
}

// WatchTokenExchangeUnderlying is a free log subscription operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_StableSwapNG *StableSwapNGFilterer) WatchTokenExchangeUnderlying(opts *bind.WatchOpts, sink chan<- *StableSwapNGTokenExchangeUnderlying, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGTokenExchangeUnderlying)
				if err := _StableSwapNG.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
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

// ParseTokenExchangeUnderlying is a log parse operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_StableSwapNG *StableSwapNGFilterer) ParseTokenExchangeUnderlying(log types.Log) (*StableSwapNGTokenExchangeUnderlying, error) {
	event := new(StableSwapNGTokenExchangeUnderlying)
	if err := _StableSwapNG.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNGTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StableSwapNG contract.
type StableSwapNGTransferIterator struct {
	Event *StableSwapNGTransfer // Event containing the contract specifics and raw log

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
func (it *StableSwapNGTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNGTransfer)
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
		it.Event = new(StableSwapNGTransfer)
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
func (it *StableSwapNGTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNGTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNGTransfer represents a Transfer event raised by the StableSwapNG contract.
type StableSwapNGTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_StableSwapNG *StableSwapNGFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*StableSwapNGTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _StableSwapNG.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNGTransferIterator{contract: _StableSwapNG.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_StableSwapNG *StableSwapNGFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StableSwapNGTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _StableSwapNG.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNGTransfer)
				if err := _StableSwapNG.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StableSwapNG *StableSwapNGFilterer) ParseTransfer(log types.Log) (*StableSwapNGTransfer, error) {
	event := new(StableSwapNGTransfer)
	if err := _StableSwapNG.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
