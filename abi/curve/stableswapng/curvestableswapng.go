// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvestableswapng

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

// CurveStableSwapNGMetaData contains all meta data concerning the CurveStableSwapNG contract.
var CurveStableSwapNGMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchangeUnderlying\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"name\":\"old_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"t\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyNewFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"offpeg_fee_multiplier\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetNewMATime\",\"inputs\":[{\"name\":\"ma_exp_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"D_ma_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_offpeg_fee_multiplier\",\"type\":\"uint256\"},{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"},{\"name\":\"_coins\",\"type\":\"address[]\"},{\"name\":\"_rate_multipliers\",\"type\":\"uint256[]\"},{\"name\":\"_asset_types\",\"type\":\"uint8[]\"},{\"name\":\"_method_ids\",\"type\":\"bytes4[]\"},{\"name\":\"_oracles\",\"type\":\"address[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_received\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_received\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_claim_admin_fees\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_price\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ema_price\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_p\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dx\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A_precise\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"stored_rates\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"dynamic_fee\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A\",\"inputs\":[{\"name\":\"_future_A\",\"type\":\"uint256\"},{\"name\":\"_future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_new_fee\",\"inputs\":[{\"name\":\"_new_fee\",\"type\":\"uint256\"},{\"name\":\"_new_offpeg_fee_multiplier\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_ma_exp_time\",\"inputs\":[{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"},{\"name\":\"_D_ma_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"N_COINS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"offpeg_fee_multiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_exp_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D_ma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_last_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"salt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]}]",
}

// CurveStableSwapNGABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveStableSwapNGMetaData.ABI instead.
var CurveStableSwapNGABI = CurveStableSwapNGMetaData.ABI

// CurveStableSwapNG is an auto generated Go binding around an Ethereum contract.
type CurveStableSwapNG struct {
	CurveStableSwapNGCaller     // Read-only binding to the contract
	CurveStableSwapNGTransactor // Write-only binding to the contract
	CurveStableSwapNGFilterer   // Log filterer for contract events
}

// CurveStableSwapNGCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveStableSwapNGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveStableSwapNGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveStableSwapNGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveStableSwapNGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveStableSwapNGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveStableSwapNGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveStableSwapNGSession struct {
	Contract     *CurveStableSwapNG // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CurveStableSwapNGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveStableSwapNGCallerSession struct {
	Contract *CurveStableSwapNGCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CurveStableSwapNGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveStableSwapNGTransactorSession struct {
	Contract     *CurveStableSwapNGTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CurveStableSwapNGRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveStableSwapNGRaw struct {
	Contract *CurveStableSwapNG // Generic contract binding to access the raw methods on
}

// CurveStableSwapNGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveStableSwapNGCallerRaw struct {
	Contract *CurveStableSwapNGCaller // Generic read-only contract binding to access the raw methods on
}

// CurveStableSwapNGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveStableSwapNGTransactorRaw struct {
	Contract *CurveStableSwapNGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveStableSwapNG creates a new instance of CurveStableSwapNG, bound to a specific deployed contract.
func NewCurveStableSwapNG(address common.Address, backend bind.ContractBackend) (*CurveStableSwapNG, error) {
	contract, err := bindCurveStableSwapNG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNG{CurveStableSwapNGCaller: CurveStableSwapNGCaller{contract: contract}, CurveStableSwapNGTransactor: CurveStableSwapNGTransactor{contract: contract}, CurveStableSwapNGFilterer: CurveStableSwapNGFilterer{contract: contract}}, nil
}

// NewCurveStableSwapNGCaller creates a new read-only instance of CurveStableSwapNG, bound to a specific deployed contract.
func NewCurveStableSwapNGCaller(address common.Address, caller bind.ContractCaller) (*CurveStableSwapNGCaller, error) {
	contract, err := bindCurveStableSwapNG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGCaller{contract: contract}, nil
}

// NewCurveStableSwapNGTransactor creates a new write-only instance of CurveStableSwapNG, bound to a specific deployed contract.
func NewCurveStableSwapNGTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveStableSwapNGTransactor, error) {
	contract, err := bindCurveStableSwapNG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGTransactor{contract: contract}, nil
}

// NewCurveStableSwapNGFilterer creates a new log filterer instance of CurveStableSwapNG, bound to a specific deployed contract.
func NewCurveStableSwapNGFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveStableSwapNGFilterer, error) {
	contract, err := bindCurveStableSwapNG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGFilterer{contract: contract}, nil
}

// bindCurveStableSwapNG binds a generic wrapper to an already deployed contract.
func bindCurveStableSwapNG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveStableSwapNGMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveStableSwapNG *CurveStableSwapNGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveStableSwapNG.Contract.CurveStableSwapNGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveStableSwapNG *CurveStableSwapNGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.CurveStableSwapNGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveStableSwapNG *CurveStableSwapNGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.CurveStableSwapNGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveStableSwapNG *CurveStableSwapNGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveStableSwapNG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveStableSwapNG *CurveStableSwapNGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveStableSwapNG *CurveStableSwapNGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) A() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.A(&_CurveStableSwapNG.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) A() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.A(&_CurveStableSwapNG.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) APrecise() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.APrecise(&_CurveStableSwapNG.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) APrecise() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.APrecise(&_CurveStableSwapNG.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurveStableSwapNG *CurveStableSwapNGSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CurveStableSwapNG.Contract.DOMAINSEPARATOR(&_CurveStableSwapNG.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CurveStableSwapNG.Contract.DOMAINSEPARATOR(&_CurveStableSwapNG.CallOpts)
}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) DMaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "D_ma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) DMaTime() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.DMaTime(&_CurveStableSwapNG.CallOpts)
}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) DMaTime() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.DMaTime(&_CurveStableSwapNG.CallOpts)
}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) DOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "D_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) DOracle() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.DOracle(&_CurveStableSwapNG.CallOpts)
}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) DOracle() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.DOracle(&_CurveStableSwapNG.CallOpts)
}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) NCOINS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "N_COINS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) NCOINS() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.NCOINS(&_CurveStableSwapNG.CallOpts)
}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) NCOINS() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.NCOINS(&_CurveStableSwapNG.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) AdminBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "admin_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.AdminBalances(&_CurveStableSwapNG.CallOpts, arg0)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.AdminBalances(&_CurveStableSwapNG.CallOpts, arg0)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) AdminFee() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.AdminFee(&_CurveStableSwapNG.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) AdminFee() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.AdminFee(&_CurveStableSwapNG.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.Allowance(&_CurveStableSwapNG.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.Allowance(&_CurveStableSwapNG.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.BalanceOf(&_CurveStableSwapNG.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.BalanceOf(&_CurveStableSwapNG.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) Balances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Balances(i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.Balances(&_CurveStableSwapNG.CallOpts, i)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) Balances(i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.Balances(&_CurveStableSwapNG.CallOpts, i)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) CalcTokenAmount(_amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.CalcTokenAmount(&_CurveStableSwapNG.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) CalcTokenAmount(_amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.CalcTokenAmount(&_CurveStableSwapNG.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "calc_withdraw_one_coin", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.CalcWithdrawOneCoin(&_CurveStableSwapNG.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.CalcWithdrawOneCoin(&_CurveStableSwapNG.CallOpts, _burn_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurveStableSwapNG.Contract.Coins(&_CurveStableSwapNG.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurveStableSwapNG.Contract.Coins(&_CurveStableSwapNG.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Decimals() (uint8, error) {
	return _CurveStableSwapNG.Contract.Decimals(&_CurveStableSwapNG.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) Decimals() (uint8, error) {
	return _CurveStableSwapNG.Contract.Decimals(&_CurveStableSwapNG.CallOpts)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) DynamicFee(opts *bind.CallOpts, i *big.Int, j *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "dynamic_fee", i, j)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.DynamicFee(&_CurveStableSwapNG.CallOpts, i, j)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.DynamicFee(&_CurveStableSwapNG.CallOpts, i, j)
}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) EmaPrice(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "ema_price", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) EmaPrice(i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.EmaPrice(&_CurveStableSwapNG.CallOpts, i)
}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) EmaPrice(i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.EmaPrice(&_CurveStableSwapNG.CallOpts, i)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Fee() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.Fee(&_CurveStableSwapNG.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) Fee() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.Fee(&_CurveStableSwapNG.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) FutureA() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.FutureA(&_CurveStableSwapNG.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) FutureA() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.FutureA(&_CurveStableSwapNG.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) FutureATime() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.FutureATime(&_CurveStableSwapNG.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) FutureATime() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.FutureATime(&_CurveStableSwapNG.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGCaller) GetBalances(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "get_balances")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGSession) GetBalances() ([]*big.Int, error) {
	return _CurveStableSwapNG.Contract.GetBalances(&_CurveStableSwapNG.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) GetBalances() ([]*big.Int, error) {
	return _CurveStableSwapNG.Contract.GetBalances(&_CurveStableSwapNG.CallOpts)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) GetDx(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "get_dx", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.GetDx(&_CurveStableSwapNG.CallOpts, i, j, dy)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.GetDx(&_CurveStableSwapNG.CallOpts, i, j, dy)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.GetDy(&_CurveStableSwapNG.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.GetDy(&_CurveStableSwapNG.CallOpts, i, j, dx)
}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) GetP(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "get_p", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) GetP(i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.GetP(&_CurveStableSwapNG.CallOpts, i)
}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) GetP(i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.GetP(&_CurveStableSwapNG.CallOpts, i)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.GetVirtualPrice(&_CurveStableSwapNG.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.GetVirtualPrice(&_CurveStableSwapNG.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) InitialA() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.InitialA(&_CurveStableSwapNG.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) InitialA() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.InitialA(&_CurveStableSwapNG.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) InitialATime() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.InitialATime(&_CurveStableSwapNG.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) InitialATime() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.InitialATime(&_CurveStableSwapNG.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) LastPrice(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "last_price", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) LastPrice(i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.LastPrice(&_CurveStableSwapNG.CallOpts, i)
}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) LastPrice(i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.LastPrice(&_CurveStableSwapNG.CallOpts, i)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) MaExpTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "ma_exp_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) MaExpTime() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.MaExpTime(&_CurveStableSwapNG.CallOpts)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) MaExpTime() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.MaExpTime(&_CurveStableSwapNG.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) MaLastTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "ma_last_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) MaLastTime() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.MaLastTime(&_CurveStableSwapNG.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) MaLastTime() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.MaLastTime(&_CurveStableSwapNG.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Name() (string, error) {
	return _CurveStableSwapNG.Contract.Name(&_CurveStableSwapNG.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) Name() (string, error) {
	return _CurveStableSwapNG.Contract.Name(&_CurveStableSwapNG.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.Nonces(&_CurveStableSwapNG.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.Nonces(&_CurveStableSwapNG.CallOpts, arg0)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) OffpegFeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "offpeg_fee_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.OffpegFeeMultiplier(&_CurveStableSwapNG.CallOpts)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.OffpegFeeMultiplier(&_CurveStableSwapNG.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) PriceOracle(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "price_oracle", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) PriceOracle(i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.PriceOracle(&_CurveStableSwapNG.CallOpts, i)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) PriceOracle(i *big.Int) (*big.Int, error) {
	return _CurveStableSwapNG.Contract.PriceOracle(&_CurveStableSwapNG.CallOpts, i)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Salt() ([32]byte, error) {
	return _CurveStableSwapNG.Contract.Salt(&_CurveStableSwapNG.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) Salt() ([32]byte, error) {
	return _CurveStableSwapNG.Contract.Salt(&_CurveStableSwapNG.CallOpts)
}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGCaller) StoredRates(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "stored_rates")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGSession) StoredRates() ([]*big.Int, error) {
	return _CurveStableSwapNG.Contract.StoredRates(&_CurveStableSwapNG.CallOpts)
}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) StoredRates() ([]*big.Int, error) {
	return _CurveStableSwapNG.Contract.StoredRates(&_CurveStableSwapNG.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Symbol() (string, error) {
	return _CurveStableSwapNG.Contract.Symbol(&_CurveStableSwapNG.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) Symbol() (string, error) {
	return _CurveStableSwapNG.Contract.Symbol(&_CurveStableSwapNG.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) TotalSupply() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.TotalSupply(&_CurveStableSwapNG.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) TotalSupply() (*big.Int, error) {
	return _CurveStableSwapNG.Contract.TotalSupply(&_CurveStableSwapNG.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurveStableSwapNG *CurveStableSwapNGCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveStableSwapNG.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Version() (string, error) {
	return _CurveStableSwapNG.Contract.Version(&_CurveStableSwapNG.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurveStableSwapNG *CurveStableSwapNGCallerSession) Version() (string, error) {
	return _CurveStableSwapNG.Contract.Version(&_CurveStableSwapNG.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) AddLiquidity(opts *bind.TransactOpts, _amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "add_liquidity", _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) AddLiquidity(_amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.AddLiquidity(&_CurveStableSwapNG.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) AddLiquidity(_amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.AddLiquidity(&_CurveStableSwapNG.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) AddLiquidity0(opts *bind.TransactOpts, _amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "add_liquidity0", _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) AddLiquidity0(_amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.AddLiquidity0(&_CurveStableSwapNG.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) AddLiquidity0(_amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.AddLiquidity0(&_CurveStableSwapNG.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.Approve(&_CurveStableSwapNG.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.Approve(&_CurveStableSwapNG.TransactOpts, _spender, _value)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "exchange", i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.Exchange(&_CurveStableSwapNG.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.Exchange(&_CurveStableSwapNG.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "exchange0", i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.Exchange0(&_CurveStableSwapNG.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.Exchange0(&_CurveStableSwapNG.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) ExchangeReceived(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "exchange_received", i, j, _dx, _min_dy)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) ExchangeReceived(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.ExchangeReceived(&_CurveStableSwapNG.TransactOpts, i, j, _dx, _min_dy)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) ExchangeReceived(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.ExchangeReceived(&_CurveStableSwapNG.TransactOpts, i, j, _dx, _min_dy)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) ExchangeReceived0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "exchange_received0", i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) ExchangeReceived0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.ExchangeReceived0(&_CurveStableSwapNG.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) ExchangeReceived0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.ExchangeReceived0(&_CurveStableSwapNG.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.Permit(&_CurveStableSwapNG.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.Permit(&_CurveStableSwapNG.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveStableSwapNG *CurveStableSwapNGSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RampA(&_CurveStableSwapNG.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RampA(&_CurveStableSwapNG.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) RemoveLiquidity(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "remove_liquidity", _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidity(&_CurveStableSwapNG.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidity(&_CurveStableSwapNG.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "remove_liquidity0", _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidity0(&_CurveStableSwapNG.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidity0(&_CurveStableSwapNG.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) RemoveLiquidity1(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "remove_liquidity1", _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGSession) RemoveLiquidity1(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidity1(&_CurveStableSwapNG.TransactOpts, _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) RemoveLiquidity1(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidity1(&_CurveStableSwapNG.TransactOpts, _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, _amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "remove_liquidity_imbalance", _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) RemoveLiquidityImbalance(_amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidityImbalance(&_CurveStableSwapNG.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) RemoveLiquidityImbalance(_amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidityImbalance(&_CurveStableSwapNG.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) RemoveLiquidityImbalance0(opts *bind.TransactOpts, _amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "remove_liquidity_imbalance0", _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) RemoveLiquidityImbalance0(_amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidityImbalance0(&_CurveStableSwapNG.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) RemoveLiquidityImbalance0(_amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidityImbalance0(&_CurveStableSwapNG.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "remove_liquidity_one_coin", _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidityOneCoin(&_CurveStableSwapNG.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidityOneCoin(&_CurveStableSwapNG.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "remove_liquidity_one_coin0", _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidityOneCoin0(&_CurveStableSwapNG.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.RemoveLiquidityOneCoin0(&_CurveStableSwapNG.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) SetMaExpTime(opts *bind.TransactOpts, _ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "set_ma_exp_time", _ma_exp_time, _D_ma_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_CurveStableSwapNG *CurveStableSwapNGSession) SetMaExpTime(_ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.SetMaExpTime(&_CurveStableSwapNG.TransactOpts, _ma_exp_time, _D_ma_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) SetMaExpTime(_ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.SetMaExpTime(&_CurveStableSwapNG.TransactOpts, _ma_exp_time, _D_ma_time)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) SetNewFee(opts *bind.TransactOpts, _new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "set_new_fee", _new_fee, _new_offpeg_fee_multiplier)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_CurveStableSwapNG *CurveStableSwapNGSession) SetNewFee(_new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.SetNewFee(&_CurveStableSwapNG.TransactOpts, _new_fee, _new_offpeg_fee_multiplier)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) SetNewFee(_new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.SetNewFee(&_CurveStableSwapNG.TransactOpts, _new_fee, _new_offpeg_fee_multiplier)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveStableSwapNG *CurveStableSwapNGSession) StopRampA() (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.StopRampA(&_CurveStableSwapNG.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) StopRampA() (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.StopRampA(&_CurveStableSwapNG.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.Transfer(&_CurveStableSwapNG.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.Transfer(&_CurveStableSwapNG.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.TransferFrom(&_CurveStableSwapNG.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.TransferFrom(&_CurveStableSwapNG.TransactOpts, _from, _to, _value)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveStableSwapNG *CurveStableSwapNGTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveStableSwapNG.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveStableSwapNG *CurveStableSwapNGSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.WithdrawAdminFees(&_CurveStableSwapNG.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveStableSwapNG *CurveStableSwapNGTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurveStableSwapNG.Contract.WithdrawAdminFees(&_CurveStableSwapNG.TransactOpts)
}

// CurveStableSwapNGAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGAddLiquidityIterator struct {
	Event *CurveStableSwapNGAddLiquidity // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGAddLiquidity)
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
		it.Event = new(CurveStableSwapNGAddLiquidity)
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
func (it *CurveStableSwapNGAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGAddLiquidity represents a AddLiquidity event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGAddLiquidity struct {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurveStableSwapNGAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGAddLiquidityIterator{contract: _CurveStableSwapNG.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGAddLiquidity)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseAddLiquidity(log types.Log) (*CurveStableSwapNGAddLiquidity, error) {
	event := new(CurveStableSwapNGAddLiquidity)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGApplyNewFeeIterator is returned from FilterApplyNewFee and is used to iterate over the raw logs and unpacked data for ApplyNewFee events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGApplyNewFeeIterator struct {
	Event *CurveStableSwapNGApplyNewFee // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGApplyNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGApplyNewFee)
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
		it.Event = new(CurveStableSwapNGApplyNewFee)
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
func (it *CurveStableSwapNGApplyNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGApplyNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGApplyNewFee represents a ApplyNewFee event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGApplyNewFee struct {
	Fee                 *big.Int
	OffpegFeeMultiplier *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterApplyNewFee is a free log retrieval operation binding the contract event 0x750d10a7f37466ce785ee6bcb604aac543358db42afbcc332a3c12a49c80bf6d.
//
// Solidity: event ApplyNewFee(uint256 fee, uint256 offpeg_fee_multiplier)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterApplyNewFee(opts *bind.FilterOpts) (*CurveStableSwapNGApplyNewFeeIterator, error) {

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGApplyNewFeeIterator{contract: _CurveStableSwapNG.contract, event: "ApplyNewFee", logs: logs, sub: sub}, nil
}

// WatchApplyNewFee is a free log subscription operation binding the contract event 0x750d10a7f37466ce785ee6bcb604aac543358db42afbcc332a3c12a49c80bf6d.
//
// Solidity: event ApplyNewFee(uint256 fee, uint256 offpeg_fee_multiplier)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchApplyNewFee(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGApplyNewFee) (event.Subscription, error) {

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGApplyNewFee)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseApplyNewFee(log types.Log) (*CurveStableSwapNGApplyNewFee, error) {
	event := new(CurveStableSwapNGApplyNewFee)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGApprovalIterator struct {
	Event *CurveStableSwapNGApproval // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGApproval)
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
		it.Event = new(CurveStableSwapNGApproval)
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
func (it *CurveStableSwapNGApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGApproval represents a Approval event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CurveStableSwapNGApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGApprovalIterator{contract: _CurveStableSwapNG.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGApproval)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseApproval(log types.Log) (*CurveStableSwapNGApproval, error) {
	event := new(CurveStableSwapNGApproval)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGRampAIterator struct {
	Event *CurveStableSwapNGRampA // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGRampA)
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
		it.Event = new(CurveStableSwapNGRampA)
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
func (it *CurveStableSwapNGRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGRampA represents a RampA event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterRampA(opts *bind.FilterOpts) (*CurveStableSwapNGRampAIterator, error) {

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGRampAIterator{contract: _CurveStableSwapNG.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGRampA) (event.Subscription, error) {

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGRampA)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "RampA", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseRampA(log types.Log) (*CurveStableSwapNGRampA, error) {
	event := new(CurveStableSwapNGRampA)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGRemoveLiquidityIterator struct {
	Event *CurveStableSwapNGRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGRemoveLiquidity)
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
		it.Event = new(CurveStableSwapNGRemoveLiquidity)
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
func (it *CurveStableSwapNGRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGRemoveLiquidity represents a RemoveLiquidity event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts []*big.Int
	Fees         []*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x347ad828e58cbe534d8f6b67985d791360756b18f0d95fd9f197a66cc46480ea.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 token_supply)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurveStableSwapNGRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGRemoveLiquidityIterator{contract: _CurveStableSwapNG.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x347ad828e58cbe534d8f6b67985d791360756b18f0d95fd9f197a66cc46480ea.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 token_supply)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGRemoveLiquidity)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseRemoveLiquidity(log types.Log) (*CurveStableSwapNGRemoveLiquidity, error) {
	event := new(CurveStableSwapNGRemoveLiquidity)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGRemoveLiquidityImbalanceIterator struct {
	Event *CurveStableSwapNGRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGRemoveLiquidityImbalance)
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
		it.Event = new(CurveStableSwapNGRemoveLiquidityImbalance)
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
func (it *CurveStableSwapNGRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGRemoveLiquidityImbalance struct {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*CurveStableSwapNGRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGRemoveLiquidityImbalanceIterator{contract: _CurveStableSwapNG.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGRemoveLiquidityImbalance)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*CurveStableSwapNGRemoveLiquidityImbalance, error) {
	event := new(CurveStableSwapNGRemoveLiquidityImbalance)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGRemoveLiquidityOneIterator struct {
	Event *CurveStableSwapNGRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGRemoveLiquidityOne)
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
		it.Event = new(CurveStableSwapNGRemoveLiquidityOne)
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
func (it *CurveStableSwapNGRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGRemoveLiquidityOne struct {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*CurveStableSwapNGRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGRemoveLiquidityOneIterator{contract: _CurveStableSwapNG.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x6f48129db1f37ccb9cc5dd7e119cb32750cabdf75b48375d730d26ce3659bbe1.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, int128 token_id, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGRemoveLiquidityOne)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseRemoveLiquidityOne(log types.Log) (*CurveStableSwapNGRemoveLiquidityOne, error) {
	event := new(CurveStableSwapNGRemoveLiquidityOne)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGSetNewMATimeIterator is returned from FilterSetNewMATime and is used to iterate over the raw logs and unpacked data for SetNewMATime events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGSetNewMATimeIterator struct {
	Event *CurveStableSwapNGSetNewMATime // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGSetNewMATimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGSetNewMATime)
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
		it.Event = new(CurveStableSwapNGSetNewMATime)
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
func (it *CurveStableSwapNGSetNewMATimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGSetNewMATimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGSetNewMATime represents a SetNewMATime event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGSetNewMATime struct {
	MaExpTime *big.Int
	DMaTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetNewMATime is a free log retrieval operation binding the contract event 0x68dc4e067dff1862b896b7a0faf55f97df1a60d0aaa79481b69d675f2026a28c.
//
// Solidity: event SetNewMATime(uint256 ma_exp_time, uint256 D_ma_time)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterSetNewMATime(opts *bind.FilterOpts) (*CurveStableSwapNGSetNewMATimeIterator, error) {

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "SetNewMATime")
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGSetNewMATimeIterator{contract: _CurveStableSwapNG.contract, event: "SetNewMATime", logs: logs, sub: sub}, nil
}

// WatchSetNewMATime is a free log subscription operation binding the contract event 0x68dc4e067dff1862b896b7a0faf55f97df1a60d0aaa79481b69d675f2026a28c.
//
// Solidity: event SetNewMATime(uint256 ma_exp_time, uint256 D_ma_time)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchSetNewMATime(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGSetNewMATime) (event.Subscription, error) {

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "SetNewMATime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGSetNewMATime)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "SetNewMATime", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseSetNewMATime(log types.Log) (*CurveStableSwapNGSetNewMATime, error) {
	event := new(CurveStableSwapNGSetNewMATime)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "SetNewMATime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGStopRampAIterator struct {
	Event *CurveStableSwapNGStopRampA // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGStopRampA)
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
		it.Event = new(CurveStableSwapNGStopRampA)
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
func (it *CurveStableSwapNGStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGStopRampA represents a StopRampA event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGStopRampA struct {
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterStopRampA(opts *bind.FilterOpts) (*CurveStableSwapNGStopRampAIterator, error) {

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGStopRampAIterator{contract: _CurveStableSwapNG.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGStopRampA) (event.Subscription, error) {

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGStopRampA)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "StopRampA", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseStopRampA(log types.Log) (*CurveStableSwapNGStopRampA, error) {
	event := new(CurveStableSwapNGStopRampA)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGTokenExchangeIterator struct {
	Event *CurveStableSwapNGTokenExchange // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGTokenExchange)
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
		it.Event = new(CurveStableSwapNGTokenExchange)
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
func (it *CurveStableSwapNGTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGTokenExchange represents a TokenExchange event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGTokenExchange struct {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*CurveStableSwapNGTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGTokenExchangeIterator{contract: _CurveStableSwapNG.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGTokenExchange)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseTokenExchange(log types.Log) (*CurveStableSwapNGTokenExchange, error) {
	event := new(CurveStableSwapNGTokenExchange)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGTokenExchangeUnderlyingIterator is returned from FilterTokenExchangeUnderlying and is used to iterate over the raw logs and unpacked data for TokenExchangeUnderlying events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGTokenExchangeUnderlyingIterator struct {
	Event *CurveStableSwapNGTokenExchangeUnderlying // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGTokenExchangeUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGTokenExchangeUnderlying)
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
		it.Event = new(CurveStableSwapNGTokenExchangeUnderlying)
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
func (it *CurveStableSwapNGTokenExchangeUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGTokenExchangeUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGTokenExchangeUnderlying represents a TokenExchangeUnderlying event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGTokenExchangeUnderlying struct {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterTokenExchangeUnderlying(opts *bind.FilterOpts, buyer []common.Address) (*CurveStableSwapNGTokenExchangeUnderlyingIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGTokenExchangeUnderlyingIterator{contract: _CurveStableSwapNG.contract, event: "TokenExchangeUnderlying", logs: logs, sub: sub}, nil
}

// WatchTokenExchangeUnderlying is a free log subscription operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchTokenExchangeUnderlying(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGTokenExchangeUnderlying, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGTokenExchangeUnderlying)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseTokenExchangeUnderlying(log types.Log) (*CurveStableSwapNGTokenExchangeUnderlying, error) {
	event := new(CurveStableSwapNGTokenExchangeUnderlying)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStableSwapNGTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CurveStableSwapNG contract.
type CurveStableSwapNGTransferIterator struct {
	Event *CurveStableSwapNGTransfer // Event containing the contract specifics and raw log

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
func (it *CurveStableSwapNGTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStableSwapNGTransfer)
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
		it.Event = new(CurveStableSwapNGTransfer)
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
func (it *CurveStableSwapNGTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStableSwapNGTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStableSwapNGTransfer represents a Transfer event raised by the CurveStableSwapNG contract.
type CurveStableSwapNGTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*CurveStableSwapNGTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CurveStableSwapNGTransferIterator{contract: _CurveStableSwapNG.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CurveStableSwapNGTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CurveStableSwapNG.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStableSwapNGTransfer)
				if err := _CurveStableSwapNG.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CurveStableSwapNG *CurveStableSwapNGFilterer) ParseTransfer(log types.Log) (*CurveStableSwapNGTransfer, error) {
	event := new(CurveStableSwapNGTransfer)
	if err := _CurveStableSwapNG.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
