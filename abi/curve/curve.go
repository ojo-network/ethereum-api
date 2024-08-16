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

// CurveMetaData contains all meta data concerning the Curve contract.
var CurveMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchangeUnderlying\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"name\":\"old_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"t\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyNewFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"offpeg_fee_multiplier\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetNewMATime\",\"inputs\":[{\"name\":\"ma_exp_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"D_ma_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_offpeg_fee_multiplier\",\"type\":\"uint256\"},{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"},{\"name\":\"_coins\",\"type\":\"address[]\"},{\"name\":\"_rate_multipliers\",\"type\":\"uint256[]\"},{\"name\":\"_asset_types\",\"type\":\"uint8[]\"},{\"name\":\"_method_ids\",\"type\":\"bytes4[]\"},{\"name\":\"_oracles\",\"type\":\"address[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_received\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_received\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_claim_admin_fees\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_price\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ema_price\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_p\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dx\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A_precise\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"stored_rates\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"dynamic_fee\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A\",\"inputs\":[{\"name\":\"_future_A\",\"type\":\"uint256\"},{\"name\":\"_future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_new_fee\",\"inputs\":[{\"name\":\"_new_fee\",\"type\":\"uint256\"},{\"name\":\"_new_offpeg_fee_multiplier\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_ma_exp_time\",\"inputs\":[{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"},{\"name\":\"_D_ma_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"N_COINS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"offpeg_fee_multiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_exp_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D_ma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_last_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"salt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]}]",
}

// CurveABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveMetaData.ABI instead.
var CurveABI = CurveMetaData.ABI

// Curve is an auto generated Go binding around an Ethereum contract.
type Curve struct {
	CurveCaller     // Read-only binding to the contract
	CurveTransactor // Write-only binding to the contract
	CurveFilterer   // Log filterer for contract events
}

// CurveCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveSession struct {
	Contract     *Curve            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveCallerSession struct {
	Contract *CurveCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CurveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveTransactorSession struct {
	Contract     *CurveTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveRaw struct {
	Contract *Curve // Generic contract binding to access the raw methods on
}

// CurveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveCallerRaw struct {
	Contract *CurveCaller // Generic read-only contract binding to access the raw methods on
}

// CurveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveTransactorRaw struct {
	Contract *CurveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurve creates a new instance of Curve, bound to a specific deployed contract.
func NewCurve(address common.Address, backend bind.ContractBackend) (*Curve, error) {
	contract, err := bindCurve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Curve{CurveCaller: CurveCaller{contract: contract}, CurveTransactor: CurveTransactor{contract: contract}, CurveFilterer: CurveFilterer{contract: contract}}, nil
}

// NewCurveCaller creates a new read-only instance of Curve, bound to a specific deployed contract.
func NewCurveCaller(address common.Address, caller bind.ContractCaller) (*CurveCaller, error) {
	contract, err := bindCurve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveCaller{contract: contract}, nil
}

// NewCurveTransactor creates a new write-only instance of Curve, bound to a specific deployed contract.
func NewCurveTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveTransactor, error) {
	contract, err := bindCurve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveTransactor{contract: contract}, nil
}

// NewCurveFilterer creates a new log filterer instance of Curve, bound to a specific deployed contract.
func NewCurveFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveFilterer, error) {
	contract, err := bindCurve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveFilterer{contract: contract}, nil
}

// bindCurve binds a generic wrapper to an already deployed contract.
func bindCurve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curve *CurveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curve.Contract.CurveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curve *CurveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curve.Contract.CurveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curve *CurveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curve.Contract.CurveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curve *CurveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curve *CurveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curve *CurveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curve.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Curve *CurveCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Curve *CurveSession) A() (*big.Int, error) {
	return _Curve.Contract.A(&_Curve.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Curve *CurveCallerSession) A() (*big.Int, error) {
	return _Curve.Contract.A(&_Curve.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_Curve *CurveCaller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_Curve *CurveSession) APrecise() (*big.Int, error) {
	return _Curve.Contract.APrecise(&_Curve.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_Curve *CurveCallerSession) APrecise() (*big.Int, error) {
	return _Curve.Contract.APrecise(&_Curve.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Curve *CurveCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Curve *CurveSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Curve.Contract.DOMAINSEPARATOR(&_Curve.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Curve *CurveCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Curve.Contract.DOMAINSEPARATOR(&_Curve.CallOpts)
}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_Curve *CurveCaller) DMaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "D_ma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_Curve *CurveSession) DMaTime() (*big.Int, error) {
	return _Curve.Contract.DMaTime(&_Curve.CallOpts)
}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_Curve *CurveCallerSession) DMaTime() (*big.Int, error) {
	return _Curve.Contract.DMaTime(&_Curve.CallOpts)
}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_Curve *CurveCaller) DOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "D_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_Curve *CurveSession) DOracle() (*big.Int, error) {
	return _Curve.Contract.DOracle(&_Curve.CallOpts)
}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_Curve *CurveCallerSession) DOracle() (*big.Int, error) {
	return _Curve.Contract.DOracle(&_Curve.CallOpts)
}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_Curve *CurveCaller) NCOINS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "N_COINS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_Curve *CurveSession) NCOINS() (*big.Int, error) {
	return _Curve.Contract.NCOINS(&_Curve.CallOpts)
}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_Curve *CurveCallerSession) NCOINS() (*big.Int, error) {
	return _Curve.Contract.NCOINS(&_Curve.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_Curve *CurveCaller) AdminBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "admin_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_Curve *CurveSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _Curve.Contract.AdminBalances(&_Curve.CallOpts, arg0)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_Curve *CurveCallerSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _Curve.Contract.AdminBalances(&_Curve.CallOpts, arg0)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Curve *CurveCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Curve *CurveSession) AdminFee() (*big.Int, error) {
	return _Curve.Contract.AdminFee(&_Curve.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Curve *CurveCallerSession) AdminFee() (*big.Int, error) {
	return _Curve.Contract.AdminFee(&_Curve.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Curve *CurveCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Curve *CurveSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Curve.Contract.Allowance(&_Curve.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Curve *CurveCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Curve.Contract.Allowance(&_Curve.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Curve *CurveCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Curve *CurveSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Curve.Contract.BalanceOf(&_Curve.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Curve *CurveCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Curve.Contract.BalanceOf(&_Curve.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_Curve *CurveCaller) Balances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_Curve *CurveSession) Balances(i *big.Int) (*big.Int, error) {
	return _Curve.Contract.Balances(&_Curve.CallOpts, i)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_Curve *CurveCallerSession) Balances(i *big.Int) (*big.Int, error) {
	return _Curve.Contract.Balances(&_Curve.CallOpts, i)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_Curve *CurveCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_Curve *CurveSession) CalcTokenAmount(_amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	return _Curve.Contract.CalcTokenAmount(&_Curve.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_Curve *CurveCallerSession) CalcTokenAmount(_amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	return _Curve.Contract.CalcTokenAmount(&_Curve.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_Curve *CurveCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "calc_withdraw_one_coin", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_Curve *CurveSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Curve.Contract.CalcWithdrawOneCoin(&_Curve.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_Curve *CurveCallerSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Curve.Contract.CalcWithdrawOneCoin(&_Curve.CallOpts, _burn_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Curve *CurveCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Curve *CurveSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Curve.Contract.Coins(&_Curve.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Curve *CurveCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Curve.Contract.Coins(&_Curve.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Curve *CurveCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Curve *CurveSession) Decimals() (uint8, error) {
	return _Curve.Contract.Decimals(&_Curve.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Curve *CurveCallerSession) Decimals() (uint8, error) {
	return _Curve.Contract.Decimals(&_Curve.CallOpts)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_Curve *CurveCaller) DynamicFee(opts *bind.CallOpts, i *big.Int, j *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "dynamic_fee", i, j)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_Curve *CurveSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _Curve.Contract.DynamicFee(&_Curve.CallOpts, i, j)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_Curve *CurveCallerSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _Curve.Contract.DynamicFee(&_Curve.CallOpts, i, j)
}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_Curve *CurveCaller) EmaPrice(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "ema_price", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_Curve *CurveSession) EmaPrice(i *big.Int) (*big.Int, error) {
	return _Curve.Contract.EmaPrice(&_Curve.CallOpts, i)
}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_Curve *CurveCallerSession) EmaPrice(i *big.Int) (*big.Int, error) {
	return _Curve.Contract.EmaPrice(&_Curve.CallOpts, i)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Curve *CurveCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Curve *CurveSession) Fee() (*big.Int, error) {
	return _Curve.Contract.Fee(&_Curve.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Curve *CurveCallerSession) Fee() (*big.Int, error) {
	return _Curve.Contract.Fee(&_Curve.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_Curve *CurveCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_Curve *CurveSession) FutureA() (*big.Int, error) {
	return _Curve.Contract.FutureA(&_Curve.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_Curve *CurveCallerSession) FutureA() (*big.Int, error) {
	return _Curve.Contract.FutureA(&_Curve.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_Curve *CurveCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_Curve *CurveSession) FutureATime() (*big.Int, error) {
	return _Curve.Contract.FutureATime(&_Curve.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_Curve *CurveCallerSession) FutureATime() (*big.Int, error) {
	return _Curve.Contract.FutureATime(&_Curve.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_Curve *CurveCaller) GetBalances(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_balances")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_Curve *CurveSession) GetBalances() ([]*big.Int, error) {
	return _Curve.Contract.GetBalances(&_Curve.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_Curve *CurveCallerSession) GetBalances() ([]*big.Int, error) {
	return _Curve.Contract.GetBalances(&_Curve.CallOpts)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_Curve *CurveCaller) GetDx(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_dx", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_Curve *CurveSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _Curve.Contract.GetDx(&_Curve.CallOpts, i, j, dy)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_Curve *CurveCallerSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _Curve.Contract.GetDx(&_Curve.CallOpts, i, j, dy)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Curve *CurveCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Curve *CurveSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Curve.Contract.GetDy(&_Curve.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Curve *CurveCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Curve.Contract.GetDy(&_Curve.CallOpts, i, j, dx)
}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_Curve *CurveCaller) GetP(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_p", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_Curve *CurveSession) GetP(i *big.Int) (*big.Int, error) {
	return _Curve.Contract.GetP(&_Curve.CallOpts, i)
}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_Curve *CurveCallerSession) GetP(i *big.Int) (*big.Int, error) {
	return _Curve.Contract.GetP(&_Curve.CallOpts, i)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Curve *CurveCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Curve *CurveSession) GetVirtualPrice() (*big.Int, error) {
	return _Curve.Contract.GetVirtualPrice(&_Curve.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Curve *CurveCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Curve.Contract.GetVirtualPrice(&_Curve.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_Curve *CurveCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_Curve *CurveSession) InitialA() (*big.Int, error) {
	return _Curve.Contract.InitialA(&_Curve.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_Curve *CurveCallerSession) InitialA() (*big.Int, error) {
	return _Curve.Contract.InitialA(&_Curve.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_Curve *CurveCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_Curve *CurveSession) InitialATime() (*big.Int, error) {
	return _Curve.Contract.InitialATime(&_Curve.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_Curve *CurveCallerSession) InitialATime() (*big.Int, error) {
	return _Curve.Contract.InitialATime(&_Curve.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_Curve *CurveCaller) LastPrice(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "last_price", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_Curve *CurveSession) LastPrice(i *big.Int) (*big.Int, error) {
	return _Curve.Contract.LastPrice(&_Curve.CallOpts, i)
}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_Curve *CurveCallerSession) LastPrice(i *big.Int) (*big.Int, error) {
	return _Curve.Contract.LastPrice(&_Curve.CallOpts, i)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_Curve *CurveCaller) MaExpTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "ma_exp_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_Curve *CurveSession) MaExpTime() (*big.Int, error) {
	return _Curve.Contract.MaExpTime(&_Curve.CallOpts)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_Curve *CurveCallerSession) MaExpTime() (*big.Int, error) {
	return _Curve.Contract.MaExpTime(&_Curve.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_Curve *CurveCaller) MaLastTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "ma_last_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_Curve *CurveSession) MaLastTime() (*big.Int, error) {
	return _Curve.Contract.MaLastTime(&_Curve.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_Curve *CurveCallerSession) MaLastTime() (*big.Int, error) {
	return _Curve.Contract.MaLastTime(&_Curve.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Curve *CurveCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Curve *CurveSession) Name() (string, error) {
	return _Curve.Contract.Name(&_Curve.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Curve *CurveCallerSession) Name() (string, error) {
	return _Curve.Contract.Name(&_Curve.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Curve *CurveCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Curve *CurveSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Curve.Contract.Nonces(&_Curve.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Curve *CurveCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Curve.Contract.Nonces(&_Curve.CallOpts, arg0)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_Curve *CurveCaller) OffpegFeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "offpeg_fee_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_Curve *CurveSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _Curve.Contract.OffpegFeeMultiplier(&_Curve.CallOpts)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_Curve *CurveCallerSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _Curve.Contract.OffpegFeeMultiplier(&_Curve.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_Curve *CurveCaller) PriceOracle(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "price_oracle", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_Curve *CurveSession) PriceOracle(i *big.Int) (*big.Int, error) {
	return _Curve.Contract.PriceOracle(&_Curve.CallOpts, i)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_Curve *CurveCallerSession) PriceOracle(i *big.Int) (*big.Int, error) {
	return _Curve.Contract.PriceOracle(&_Curve.CallOpts, i)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Curve *CurveCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Curve *CurveSession) Salt() ([32]byte, error) {
	return _Curve.Contract.Salt(&_Curve.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Curve *CurveCallerSession) Salt() ([32]byte, error) {
	return _Curve.Contract.Salt(&_Curve.CallOpts)
}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_Curve *CurveCaller) StoredRates(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "stored_rates")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_Curve *CurveSession) StoredRates() ([]*big.Int, error) {
	return _Curve.Contract.StoredRates(&_Curve.CallOpts)
}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_Curve *CurveCallerSession) StoredRates() ([]*big.Int, error) {
	return _Curve.Contract.StoredRates(&_Curve.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Curve *CurveCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Curve *CurveSession) Symbol() (string, error) {
	return _Curve.Contract.Symbol(&_Curve.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Curve *CurveCallerSession) Symbol() (string, error) {
	return _Curve.Contract.Symbol(&_Curve.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Curve *CurveCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Curve *CurveSession) TotalSupply() (*big.Int, error) {
	return _Curve.Contract.TotalSupply(&_Curve.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Curve *CurveCallerSession) TotalSupply() (*big.Int, error) {
	return _Curve.Contract.TotalSupply(&_Curve.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Curve *CurveCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Curve *CurveSession) Version() (string, error) {
	return _Curve.Contract.Version(&_Curve.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Curve *CurveCallerSession) Version() (string, error) {
	return _Curve.Contract.Version(&_Curve.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_Curve *CurveTransactor) AddLiquidity(opts *bind.TransactOpts, _amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "add_liquidity", _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_Curve *CurveSession) AddLiquidity(_amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.AddLiquidity(&_Curve.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_Curve *CurveTransactorSession) AddLiquidity(_amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.AddLiquidity(&_Curve.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_Curve *CurveTransactor) AddLiquidity0(opts *bind.TransactOpts, _amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "add_liquidity0", _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_Curve *CurveSession) AddLiquidity0(_amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.AddLiquidity0(&_Curve.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_Curve *CurveTransactorSession) AddLiquidity0(_amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.AddLiquidity0(&_Curve.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Curve *CurveTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Curve *CurveSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.Approve(&_Curve.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Curve *CurveTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.Approve(&_Curve.TransactOpts, _spender, _value)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Curve *CurveTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "exchange", i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Curve *CurveSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.Exchange(&_Curve.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Curve *CurveTransactorSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.Exchange(&_Curve.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Curve *CurveTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "exchange0", i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Curve *CurveSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.Exchange0(&_Curve.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Curve *CurveTransactorSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.Exchange0(&_Curve.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Curve *CurveTransactor) ExchangeReceived(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "exchange_received", i, j, _dx, _min_dy)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Curve *CurveSession) ExchangeReceived(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.ExchangeReceived(&_Curve.TransactOpts, i, j, _dx, _min_dy)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Curve *CurveTransactorSession) ExchangeReceived(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.ExchangeReceived(&_Curve.TransactOpts, i, j, _dx, _min_dy)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Curve *CurveTransactor) ExchangeReceived0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "exchange_received0", i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Curve *CurveSession) ExchangeReceived0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.ExchangeReceived0(&_Curve.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Curve *CurveTransactorSession) ExchangeReceived0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.ExchangeReceived0(&_Curve.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_Curve *CurveTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_Curve *CurveSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Curve.Contract.Permit(&_Curve.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_Curve *CurveTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Curve.Contract.Permit(&_Curve.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Curve *CurveTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Curve *CurveSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.RampA(&_Curve.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Curve *CurveTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.RampA(&_Curve.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_Curve *CurveTransactor) RemoveLiquidity(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "remove_liquidity", _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_Curve *CurveSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidity(&_Curve.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_Curve *CurveTransactorSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidity(&_Curve.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_Curve *CurveTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "remove_liquidity0", _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_Curve *CurveSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidity0(&_Curve.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_Curve *CurveTransactorSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidity0(&_Curve.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_Curve *CurveTransactor) RemoveLiquidity1(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "remove_liquidity1", _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_Curve *CurveSession) RemoveLiquidity1(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidity1(&_Curve.TransactOpts, _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_Curve *CurveTransactorSession) RemoveLiquidity1(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidity1(&_Curve.TransactOpts, _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_Curve *CurveTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, _amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "remove_liquidity_imbalance", _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_Curve *CurveSession) RemoveLiquidityImbalance(_amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidityImbalance(&_Curve.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_Curve *CurveTransactorSession) RemoveLiquidityImbalance(_amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidityImbalance(&_Curve.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_Curve *CurveTransactor) RemoveLiquidityImbalance0(opts *bind.TransactOpts, _amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "remove_liquidity_imbalance0", _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_Curve *CurveSession) RemoveLiquidityImbalance0(_amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidityImbalance0(&_Curve.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_Curve *CurveTransactorSession) RemoveLiquidityImbalance0(_amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidityImbalance0(&_Curve.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_Curve *CurveTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "remove_liquidity_one_coin", _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_Curve *CurveSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidityOneCoin(&_Curve.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_Curve *CurveTransactorSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidityOneCoin(&_Curve.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_Curve *CurveTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "remove_liquidity_one_coin0", _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_Curve *CurveSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidityOneCoin0(&_Curve.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_Curve *CurveTransactorSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Curve.Contract.RemoveLiquidityOneCoin0(&_Curve.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_Curve *CurveTransactor) SetMaExpTime(opts *bind.TransactOpts, _ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "set_ma_exp_time", _ma_exp_time, _D_ma_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_Curve *CurveSession) SetMaExpTime(_ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.SetMaExpTime(&_Curve.TransactOpts, _ma_exp_time, _D_ma_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_Curve *CurveTransactorSession) SetMaExpTime(_ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.SetMaExpTime(&_Curve.TransactOpts, _ma_exp_time, _D_ma_time)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_Curve *CurveTransactor) SetNewFee(opts *bind.TransactOpts, _new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "set_new_fee", _new_fee, _new_offpeg_fee_multiplier)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_Curve *CurveSession) SetNewFee(_new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.SetNewFee(&_Curve.TransactOpts, _new_fee, _new_offpeg_fee_multiplier)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_Curve *CurveTransactorSession) SetNewFee(_new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.SetNewFee(&_Curve.TransactOpts, _new_fee, _new_offpeg_fee_multiplier)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Curve *CurveTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Curve *CurveSession) StopRampA() (*types.Transaction, error) {
	return _Curve.Contract.StopRampA(&_Curve.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Curve *CurveTransactorSession) StopRampA() (*types.Transaction, error) {
	return _Curve.Contract.StopRampA(&_Curve.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Curve *CurveTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Curve *CurveSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.Transfer(&_Curve.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Curve *CurveTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.Transfer(&_Curve.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Curve *CurveTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Curve *CurveSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.TransferFrom(&_Curve.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Curve *CurveTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.TransferFrom(&_Curve.TransactOpts, _from, _to, _value)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Curve *CurveTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Curve *CurveSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Curve.Contract.WithdrawAdminFees(&_Curve.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Curve *CurveTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Curve.Contract.WithdrawAdminFees(&_Curve.TransactOpts)
}

// CurveAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Curve contract.
type CurveAddLiquidityIterator struct {
	Event *CurveAddLiquidity // Event containing the contract specifics and raw log

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
func (it *CurveAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveAddLiquidity)
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
		it.Event = new(CurveAddLiquidity)
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
func (it *CurveAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveAddLiquidity represents a AddLiquidity event raised by the Curve contract.
type CurveAddLiquidity struct {
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
func (_Curve *CurveFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurveAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curve.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveAddLiquidityIterator{contract: _Curve.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_Curve *CurveFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *CurveAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curve.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveAddLiquidity)
				if err := _Curve.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_Curve *CurveFilterer) ParseAddLiquidity(log types.Log) (*CurveAddLiquidity, error) {
	event := new(CurveAddLiquidity)
	if err := _Curve.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveApplyNewFeeIterator is returned from FilterApplyNewFee and is used to iterate over the raw logs and unpacked data for ApplyNewFee events raised by the Curve contract.
type CurveApplyNewFeeIterator struct {
	Event *CurveApplyNewFee // Event containing the contract specifics and raw log

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
func (it *CurveApplyNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveApplyNewFee)
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
		it.Event = new(CurveApplyNewFee)
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
func (it *CurveApplyNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveApplyNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveApplyNewFee represents a ApplyNewFee event raised by the Curve contract.
type CurveApplyNewFee struct {
	Fee                 *big.Int
	OffpegFeeMultiplier *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterApplyNewFee is a free log retrieval operation binding the contract event 0x750d10a7f37466ce785ee6bcb604aac543358db42afbcc332a3c12a49c80bf6d.
//
// Solidity: event ApplyNewFee(uint256 fee, uint256 offpeg_fee_multiplier)
func (_Curve *CurveFilterer) FilterApplyNewFee(opts *bind.FilterOpts) (*CurveApplyNewFeeIterator, error) {

	logs, sub, err := _Curve.contract.FilterLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return &CurveApplyNewFeeIterator{contract: _Curve.contract, event: "ApplyNewFee", logs: logs, sub: sub}, nil
}

// WatchApplyNewFee is a free log subscription operation binding the contract event 0x750d10a7f37466ce785ee6bcb604aac543358db42afbcc332a3c12a49c80bf6d.
//
// Solidity: event ApplyNewFee(uint256 fee, uint256 offpeg_fee_multiplier)
func (_Curve *CurveFilterer) WatchApplyNewFee(opts *bind.WatchOpts, sink chan<- *CurveApplyNewFee) (event.Subscription, error) {

	logs, sub, err := _Curve.contract.WatchLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveApplyNewFee)
				if err := _Curve.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
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
func (_Curve *CurveFilterer) ParseApplyNewFee(log types.Log) (*CurveApplyNewFee, error) {
	event := new(CurveApplyNewFee)
	if err := _Curve.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Curve contract.
type CurveApprovalIterator struct {
	Event *CurveApproval // Event containing the contract specifics and raw log

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
func (it *CurveApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveApproval)
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
		it.Event = new(CurveApproval)
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
func (it *CurveApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveApproval represents a Approval event raised by the Curve contract.
type CurveApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Curve *CurveFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CurveApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Curve.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CurveApprovalIterator{contract: _Curve.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Curve *CurveFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CurveApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Curve.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveApproval)
				if err := _Curve.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Curve *CurveFilterer) ParseApproval(log types.Log) (*CurveApproval, error) {
	event := new(CurveApproval)
	if err := _Curve.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the Curve contract.
type CurveRampAIterator struct {
	Event *CurveRampA // Event containing the contract specifics and raw log

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
func (it *CurveRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveRampA)
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
		it.Event = new(CurveRampA)
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
func (it *CurveRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveRampA represents a RampA event raised by the Curve contract.
type CurveRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_Curve *CurveFilterer) FilterRampA(opts *bind.FilterOpts) (*CurveRampAIterator, error) {

	logs, sub, err := _Curve.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &CurveRampAIterator{contract: _Curve.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_Curve *CurveFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *CurveRampA) (event.Subscription, error) {

	logs, sub, err := _Curve.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveRampA)
				if err := _Curve.contract.UnpackLog(event, "RampA", log); err != nil {
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
func (_Curve *CurveFilterer) ParseRampA(log types.Log) (*CurveRampA, error) {
	event := new(CurveRampA)
	if err := _Curve.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Curve contract.
type CurveRemoveLiquidityIterator struct {
	Event *CurveRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *CurveRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveRemoveLiquidity)
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
		it.Event = new(CurveRemoveLiquidity)
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
func (it *CurveRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveRemoveLiquidity represents a RemoveLiquidity event raised by the Curve contract.
type CurveRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts []*big.Int
	Fees         []*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x347ad828e58cbe534d8f6b67985d791360756b18f0d95fd9f197a66cc46480ea.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 token_supply)
func (_Curve *CurveFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurveRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curve.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveRemoveLiquidityIterator{contract: _Curve.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x347ad828e58cbe534d8f6b67985d791360756b18f0d95fd9f197a66cc46480ea.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 token_supply)
func (_Curve *CurveFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *CurveRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curve.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveRemoveLiquidity)
				if err := _Curve.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_Curve *CurveFilterer) ParseRemoveLiquidity(log types.Log) (*CurveRemoveLiquidity, error) {
	event := new(CurveRemoveLiquidity)
	if err := _Curve.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the Curve contract.
type CurveRemoveLiquidityImbalanceIterator struct {
	Event *CurveRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *CurveRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveRemoveLiquidityImbalance)
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
		it.Event = new(CurveRemoveLiquidityImbalance)
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
func (it *CurveRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the Curve contract.
type CurveRemoveLiquidityImbalance struct {
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
func (_Curve *CurveFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*CurveRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curve.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveRemoveLiquidityImbalanceIterator{contract: _Curve.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_Curve *CurveFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *CurveRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curve.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveRemoveLiquidityImbalance)
				if err := _Curve.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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
func (_Curve *CurveFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*CurveRemoveLiquidityImbalance, error) {
	event := new(CurveRemoveLiquidityImbalance)
	if err := _Curve.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the Curve contract.
type CurveRemoveLiquidityOneIterator struct {
	Event *CurveRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *CurveRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveRemoveLiquidityOne)
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
		it.Event = new(CurveRemoveLiquidityOne)
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
func (it *CurveRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the Curve contract.
type CurveRemoveLiquidityOne struct {
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
func (_Curve *CurveFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*CurveRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curve.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveRemoveLiquidityOneIterator{contract: _Curve.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x6f48129db1f37ccb9cc5dd7e119cb32750cabdf75b48375d730d26ce3659bbe1.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, int128 token_id, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_Curve *CurveFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *CurveRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curve.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveRemoveLiquidityOne)
				if err := _Curve.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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
func (_Curve *CurveFilterer) ParseRemoveLiquidityOne(log types.Log) (*CurveRemoveLiquidityOne, error) {
	event := new(CurveRemoveLiquidityOne)
	if err := _Curve.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveSetNewMATimeIterator is returned from FilterSetNewMATime and is used to iterate over the raw logs and unpacked data for SetNewMATime events raised by the Curve contract.
type CurveSetNewMATimeIterator struct {
	Event *CurveSetNewMATime // Event containing the contract specifics and raw log

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
func (it *CurveSetNewMATimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveSetNewMATime)
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
		it.Event = new(CurveSetNewMATime)
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
func (it *CurveSetNewMATimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveSetNewMATimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveSetNewMATime represents a SetNewMATime event raised by the Curve contract.
type CurveSetNewMATime struct {
	MaExpTime *big.Int
	DMaTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetNewMATime is a free log retrieval operation binding the contract event 0x68dc4e067dff1862b896b7a0faf55f97df1a60d0aaa79481b69d675f2026a28c.
//
// Solidity: event SetNewMATime(uint256 ma_exp_time, uint256 D_ma_time)
func (_Curve *CurveFilterer) FilterSetNewMATime(opts *bind.FilterOpts) (*CurveSetNewMATimeIterator, error) {

	logs, sub, err := _Curve.contract.FilterLogs(opts, "SetNewMATime")
	if err != nil {
		return nil, err
	}
	return &CurveSetNewMATimeIterator{contract: _Curve.contract, event: "SetNewMATime", logs: logs, sub: sub}, nil
}

// WatchSetNewMATime is a free log subscription operation binding the contract event 0x68dc4e067dff1862b896b7a0faf55f97df1a60d0aaa79481b69d675f2026a28c.
//
// Solidity: event SetNewMATime(uint256 ma_exp_time, uint256 D_ma_time)
func (_Curve *CurveFilterer) WatchSetNewMATime(opts *bind.WatchOpts, sink chan<- *CurveSetNewMATime) (event.Subscription, error) {

	logs, sub, err := _Curve.contract.WatchLogs(opts, "SetNewMATime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveSetNewMATime)
				if err := _Curve.contract.UnpackLog(event, "SetNewMATime", log); err != nil {
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
func (_Curve *CurveFilterer) ParseSetNewMATime(log types.Log) (*CurveSetNewMATime, error) {
	event := new(CurveSetNewMATime)
	if err := _Curve.contract.UnpackLog(event, "SetNewMATime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the Curve contract.
type CurveStopRampAIterator struct {
	Event *CurveStopRampA // Event containing the contract specifics and raw log

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
func (it *CurveStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStopRampA)
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
		it.Event = new(CurveStopRampA)
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
func (it *CurveStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStopRampA represents a StopRampA event raised by the Curve contract.
type CurveStopRampA struct {
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_Curve *CurveFilterer) FilterStopRampA(opts *bind.FilterOpts) (*CurveStopRampAIterator, error) {

	logs, sub, err := _Curve.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &CurveStopRampAIterator{contract: _Curve.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_Curve *CurveFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *CurveStopRampA) (event.Subscription, error) {

	logs, sub, err := _Curve.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStopRampA)
				if err := _Curve.contract.UnpackLog(event, "StopRampA", log); err != nil {
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
func (_Curve *CurveFilterer) ParseStopRampA(log types.Log) (*CurveStopRampA, error) {
	event := new(CurveStopRampA)
	if err := _Curve.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Curve contract.
type CurveTokenExchangeIterator struct {
	Event *CurveTokenExchange // Event containing the contract specifics and raw log

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
func (it *CurveTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveTokenExchange)
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
		it.Event = new(CurveTokenExchange)
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
func (it *CurveTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveTokenExchange represents a TokenExchange event raised by the Curve contract.
type CurveTokenExchange struct {
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
func (_Curve *CurveFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*CurveTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Curve.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurveTokenExchangeIterator{contract: _Curve.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Curve *CurveFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CurveTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Curve.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveTokenExchange)
				if err := _Curve.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_Curve *CurveFilterer) ParseTokenExchange(log types.Log) (*CurveTokenExchange, error) {
	event := new(CurveTokenExchange)
	if err := _Curve.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveTokenExchangeUnderlyingIterator is returned from FilterTokenExchangeUnderlying and is used to iterate over the raw logs and unpacked data for TokenExchangeUnderlying events raised by the Curve contract.
type CurveTokenExchangeUnderlyingIterator struct {
	Event *CurveTokenExchangeUnderlying // Event containing the contract specifics and raw log

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
func (it *CurveTokenExchangeUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveTokenExchangeUnderlying)
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
		it.Event = new(CurveTokenExchangeUnderlying)
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
func (it *CurveTokenExchangeUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveTokenExchangeUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveTokenExchangeUnderlying represents a TokenExchangeUnderlying event raised by the Curve contract.
type CurveTokenExchangeUnderlying struct {
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
func (_Curve *CurveFilterer) FilterTokenExchangeUnderlying(opts *bind.FilterOpts, buyer []common.Address) (*CurveTokenExchangeUnderlyingIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Curve.contract.FilterLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurveTokenExchangeUnderlyingIterator{contract: _Curve.contract, event: "TokenExchangeUnderlying", logs: logs, sub: sub}, nil
}

// WatchTokenExchangeUnderlying is a free log subscription operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Curve *CurveFilterer) WatchTokenExchangeUnderlying(opts *bind.WatchOpts, sink chan<- *CurveTokenExchangeUnderlying, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Curve.contract.WatchLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveTokenExchangeUnderlying)
				if err := _Curve.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
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
func (_Curve *CurveFilterer) ParseTokenExchangeUnderlying(log types.Log) (*CurveTokenExchangeUnderlying, error) {
	event := new(CurveTokenExchangeUnderlying)
	if err := _Curve.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Curve contract.
type CurveTransferIterator struct {
	Event *CurveTransfer // Event containing the contract specifics and raw log

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
func (it *CurveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveTransfer)
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
		it.Event = new(CurveTransfer)
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
func (it *CurveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveTransfer represents a Transfer event raised by the Curve contract.
type CurveTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Curve *CurveFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*CurveTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Curve.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CurveTransferIterator{contract: _Curve.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Curve *CurveFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CurveTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Curve.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveTransfer)
				if err := _Curve.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Curve *CurveFilterer) ParseTransfer(log types.Log) (*CurveTransfer, error) {
	event := new(CurveTransfer)
	if err := _Curve.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
