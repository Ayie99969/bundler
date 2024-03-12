// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abivalidator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/ethkit/go-ethereum/event"
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

// IEndorserConstraint is an auto generated low-level Go binding around an user-defined struct.
type IEndorserConstraint struct {
	Slot     [32]byte
	MinValue [32]byte
	MaxValue [32]byte
}

// IEndorserDependency is an auto generated low-level Go binding around an user-defined struct.
type IEndorserDependency struct {
	Addr        common.Address
	Balance     bool
	Code        bool
	Nonce       bool
	AllSlots    bool
	Slots       [][32]byte
	Constraints []IEndorserConstraint
}

// IEndorserGlobalDependency is an auto generated low-level Go binding around an user-defined struct.
type IEndorserGlobalDependency struct {
	BaseFee           bool
	BlobBaseFee       bool
	ChainId           bool
	CoinBase          bool
	Difficulty        bool
	GasLimit          bool
	Number            bool
	Timestamp         bool
	TxOrigin          bool
	TxGasPrice        bool
	MaxBlockNumber    *big.Int
	MaxBlockTimestamp *big.Int
}

// IEndorserOperation is an auto generated low-level Go binding around an user-defined struct.
type IEndorserOperation struct {
	Entrypoint             common.Address
	Data                   []byte
	EndorserCallData       []byte
	FixedGas               *big.Int
	GasLimit               *big.Int
	MaxFeePerGas           *big.Int
	MaxPriorityFeePerGas   *big.Int
	FeeToken               common.Address
	FeeScalingFactor       *big.Int
	FeeNormalizationFactor *big.Int
	HasUntrustedContext    bool
}

// OperationValidatorSimulationResult is an auto generated low-level Go binding around an user-defined struct.
type OperationValidatorSimulationResult struct {
	Paid             bool
	Readiness        bool
	GlobalDependency IEndorserGlobalDependency
	Dependencies     []IEndorserDependency
	Err              []byte
}

// OperationValidatorMetaData contains all meta data concerning the OperationValidator contract.
var OperationValidatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"_executeAndMeasureNoSideEffects\",\"inputs\":[{\"name\":\"_op\",\"type\":\"tuple\",\"internalType\":\"structIEndorser.Operation\",\"components\":[{\"name\":\"entrypoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"endorserCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fixedGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeScalingFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeNormalizationFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hasUntrustedContext\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"simulateOperation\",\"inputs\":[{\"name\":\"_endorser\",\"type\":\"address\",\"internalType\":\"contractEndorser\"},{\"name\":\"_op\",\"type\":\"tuple\",\"internalType\":\"structIEndorser.Operation\",\"components\":[{\"name\":\"entrypoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"endorserCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fixedGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeScalingFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeNormalizationFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hasUntrustedContext\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structOperationValidator.SimulationResult\",\"components\":[{\"name\":\"paid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"readiness\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"globalDependency\",\"type\":\"tuple\",\"internalType\":\"structIEndorser.GlobalDependency\",\"components\":[{\"name\":\"baseFee\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"blobBaseFee\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"coinBase\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"difficulty\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"gasLimit\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"number\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"timestamp\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"txOrigin\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"txGasPrice\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxBlockTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"dependencies\",\"type\":\"tuple[]\",\"internalType\":\"structIEndorser.Dependency[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"code\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nonce\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allSlots\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"slots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"constraints\",\"type\":\"tuple[]\",\"internalType\":\"structIEndorser.Constraint[]\",\"components\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"minValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"BundlerExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BundlerUnderpaid\",\"inputs\":[{\"name\":\"_succeed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_paid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// OperationValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use OperationValidatorMetaData.ABI instead.
var OperationValidatorABI = OperationValidatorMetaData.ABI

// OperationValidator is an auto generated Go binding around an Ethereum contract.
type OperationValidator struct {
	OperationValidatorCaller     // Read-only binding to the contract
	OperationValidatorTransactor // Write-only binding to the contract
	OperationValidatorFilterer   // Log filterer for contract events
}

// OperationValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperationValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperationValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperationValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperationValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperationValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperationValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperationValidatorSession struct {
	Contract     *OperationValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OperationValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperationValidatorCallerSession struct {
	Contract *OperationValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OperationValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperationValidatorTransactorSession struct {
	Contract     *OperationValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OperationValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperationValidatorRaw struct {
	Contract *OperationValidator // Generic contract binding to access the raw methods on
}

// OperationValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperationValidatorCallerRaw struct {
	Contract *OperationValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// OperationValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperationValidatorTransactorRaw struct {
	Contract *OperationValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperationValidator creates a new instance of OperationValidator, bound to a specific deployed contract.
func NewOperationValidator(address common.Address, backend bind.ContractBackend) (*OperationValidator, error) {
	contract, err := bindOperationValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperationValidator{OperationValidatorCaller: OperationValidatorCaller{contract: contract}, OperationValidatorTransactor: OperationValidatorTransactor{contract: contract}, OperationValidatorFilterer: OperationValidatorFilterer{contract: contract}}, nil
}

// NewOperationValidatorCaller creates a new read-only instance of OperationValidator, bound to a specific deployed contract.
func NewOperationValidatorCaller(address common.Address, caller bind.ContractCaller) (*OperationValidatorCaller, error) {
	contract, err := bindOperationValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperationValidatorCaller{contract: contract}, nil
}

// NewOperationValidatorTransactor creates a new write-only instance of OperationValidator, bound to a specific deployed contract.
func NewOperationValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OperationValidatorTransactor, error) {
	contract, err := bindOperationValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperationValidatorTransactor{contract: contract}, nil
}

// NewOperationValidatorFilterer creates a new log filterer instance of OperationValidator, bound to a specific deployed contract.
func NewOperationValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OperationValidatorFilterer, error) {
	contract, err := bindOperationValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperationValidatorFilterer{contract: contract}, nil
}

// bindOperationValidator binds a generic wrapper to an already deployed contract.
func bindOperationValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperationValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperationValidator *OperationValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperationValidator.Contract.OperationValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperationValidator *OperationValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperationValidator.Contract.OperationValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperationValidator *OperationValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperationValidator.Contract.OperationValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperationValidator *OperationValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperationValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperationValidator *OperationValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperationValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperationValidator *OperationValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperationValidator.Contract.contract.Transact(opts, method, params...)
}

// SimulateOperation is a free data retrieval call binding the contract method 0x7b9fe62f.
//
// Solidity: function simulateOperation(address _endorser, (address,bytes,bytes,uint256,uint256,uint256,uint256,address,uint256,uint256,bool) _op) view returns((bool,bool,(bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,uint256,uint256),(address,bool,bool,bool,bool,bytes32[],(bytes32,bytes32,bytes32)[])[],bytes) result)
func (_OperationValidator *OperationValidatorCaller) SimulateOperation(opts *bind.CallOpts, _endorser common.Address, _op IEndorserOperation) (OperationValidatorSimulationResult, error) {
	var out []interface{}
	err := _OperationValidator.contract.Call(opts, &out, "simulateOperation", _endorser, _op)

	if err != nil {
		return *new(OperationValidatorSimulationResult), err
	}

	out0 := *abi.ConvertType(out[0], new(OperationValidatorSimulationResult)).(*OperationValidatorSimulationResult)

	return out0, err

}

// SimulateOperation is a free data retrieval call binding the contract method 0x7b9fe62f.
//
// Solidity: function simulateOperation(address _endorser, (address,bytes,bytes,uint256,uint256,uint256,uint256,address,uint256,uint256,bool) _op) view returns((bool,bool,(bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,uint256,uint256),(address,bool,bool,bool,bool,bytes32[],(bytes32,bytes32,bytes32)[])[],bytes) result)
func (_OperationValidator *OperationValidatorSession) SimulateOperation(_endorser common.Address, _op IEndorserOperation) (OperationValidatorSimulationResult, error) {
	return _OperationValidator.Contract.SimulateOperation(&_OperationValidator.CallOpts, _endorser, _op)
}

// SimulateOperation is a free data retrieval call binding the contract method 0x7b9fe62f.
//
// Solidity: function simulateOperation(address _endorser, (address,bytes,bytes,uint256,uint256,uint256,uint256,address,uint256,uint256,bool) _op) view returns((bool,bool,(bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,uint256,uint256),(address,bool,bool,bool,bool,bytes32[],(bytes32,bytes32,bytes32)[])[],bytes) result)
func (_OperationValidator *OperationValidatorCallerSession) SimulateOperation(_endorser common.Address, _op IEndorserOperation) (OperationValidatorSimulationResult, error) {
	return _OperationValidator.Contract.SimulateOperation(&_OperationValidator.CallOpts, _endorser, _op)
}

// ExecuteAndMeasureNoSideEffects is a paid mutator transaction binding the contract method 0x57de03ee.
//
// Solidity: function _executeAndMeasureNoSideEffects((address,bytes,bytes,uint256,uint256,uint256,uint256,address,uint256,uint256,bool) _op) returns(bool)
func (_OperationValidator *OperationValidatorTransactor) ExecuteAndMeasureNoSideEffects(opts *bind.TransactOpts, _op IEndorserOperation) (*types.Transaction, error) {
	return _OperationValidator.contract.Transact(opts, "_executeAndMeasureNoSideEffects", _op)
}

// ExecuteAndMeasureNoSideEffects is a paid mutator transaction binding the contract method 0x57de03ee.
//
// Solidity: function _executeAndMeasureNoSideEffects((address,bytes,bytes,uint256,uint256,uint256,uint256,address,uint256,uint256,bool) _op) returns(bool)
func (_OperationValidator *OperationValidatorSession) ExecuteAndMeasureNoSideEffects(_op IEndorserOperation) (*types.Transaction, error) {
	return _OperationValidator.Contract.ExecuteAndMeasureNoSideEffects(&_OperationValidator.TransactOpts, _op)
}

// ExecuteAndMeasureNoSideEffects is a paid mutator transaction binding the contract method 0x57de03ee.
//
// Solidity: function _executeAndMeasureNoSideEffects((address,bytes,bytes,uint256,uint256,uint256,uint256,address,uint256,uint256,bool) _op) returns(bool)
func (_OperationValidator *OperationValidatorTransactorSession) ExecuteAndMeasureNoSideEffects(_op IEndorserOperation) (*types.Transaction, error) {
	return _OperationValidator.Contract.ExecuteAndMeasureNoSideEffects(&_OperationValidator.TransactOpts, _op)
}
