// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todolist

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

// TodolistMetaData contains all meta data concerning the Todolist contract.
var TodolistMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"createTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"toggleCompleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600080553480156200001557600080fd5b506200005c6040518060400160405280600b81526020017f73616d706c65207461736b0000000000000000000000000000000000000000008152506200006260201b60201c565b62000615565b600080815480929190620000769062000170565b919050555060405180606001604052806000548152602001828152602001600015158152506001600080548152602001908152602001600020600082015181600001556020820151816001019081620000d091906200042d565b5060408201518160020160006101000a81548160ff0219169083151502179055509050507f05d0fb833127fc08168556d0e7ca9554fc3f6bc843b3b7d2bf1c35aea6bab6606000548260006040516200012c93929190620005d1565b60405180910390a150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b60006200017d8262000166565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620001b257620001b162000137565b5b600182019050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200023f57607f821691505b602082108103620002555762000254620001f7565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620002bf7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000280565b620002cb868362000280565b95508019841693508086168417925050509392505050565b6000819050919050565b60006200030e62000308620003028462000166565b620002e3565b62000166565b9050919050565b6000819050919050565b6200032a83620002ed565b62000342620003398262000315565b8484546200028d565b825550505050565b600090565b620003596200034a565b620003668184846200031f565b505050565b5b818110156200038e57620003826000826200034f565b6001810190506200036c565b5050565b601f821115620003dd57620003a7816200025b565b620003b28462000270565b81016020851015620003c2578190505b620003da620003d18562000270565b8301826200036b565b50505b505050565b600082821c905092915050565b60006200040260001984600802620003e2565b1980831691505092915050565b60006200041d8383620003ef565b9150826002028217905092915050565b6200043882620001bd565b67ffffffffffffffff811115620004545762000453620001c8565b5b62000460825462000226565b6200046d82828562000392565b600060209050601f831160018114620004a5576000841562000490578287015190505b6200049c85826200040f565b8655506200050c565b601f198416620004b5866200025b565b60005b82811015620004df57848901518255600182019150602085019450602081019050620004b8565b86831015620004ff5784890151620004fb601f891682620003ef565b8355505b6001600288020188555050505b505050505050565b6200051f8162000166565b82525050565b600082825260208201905092915050565b60005b838110156200055657808201518184015260208101905062000539565b60008484015250505050565b6000601f19601f8301169050919050565b60006200058082620001bd565b6200058c818562000525565b93506200059e81856020860162000536565b620005a98162000562565b840191505092915050565b60008115159050919050565b620005cb81620005b4565b82525050565b6000606082019050620005e8600083018662000514565b8181036020830152620005fc818562000573565b90506200060d6040830184620005c0565b949350505050565b610aaf80620006256000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063111002aa14610051578063455f50241461006d5780638d97767214610089578063b6cb58a5146100bb575b600080fd5b61006b6004803603810190610066919061054d565b6100d9565b005b610087600480360381019061008291906105cc565b6101a8565b005b6100a3600480360381019061009e91906105cc565b61032e565b6040516100b2939291906106a2565b60405180910390f35b6100c36103ed565b6040516100d091906106e0565b60405180910390f35b6000808154809291906100eb9061072a565b919050555060405180606001604052806000548152602001828152602001600015158152506001600080548152602001908152602001600020600082015181600001556020820151816001019081610143919061097e565b5060408201518160020160006101000a81548160ff0219169083151502179055509050507f05d0fb833127fc08168556d0e7ca9554fc3f6bc843b3b7d2bf1c35aea6bab66060005482600060405161019d939291906106a2565b60405180910390a150565b600060016000838152602001908152602001600020604051806060016040529081600082015481526020016001820180546101e2906107a1565b80601f016020809104026020016040519081016040528092919081815260200182805461020e906107a1565b801561025b5780601f106102305761010080835404028352916020019161025b565b820191906000526020600020905b81548152906001019060200180831161023e57829003601f168201915b505050505081526020016002820160009054906101000a900460ff161515151581525050905080604001511581604001901515908115158152505080600160008481526020019081526020016000206000820151816000015560208201518160010190816102c9919061097e565b5060408201518160020160006101000a81548160ff0219169083151502179055509050507fe21fa966ca5cd02748c0752352d18c48165e61cb55b4c29cccf924b5a95fcff1828260400151604051610322929190610a50565b60405180910390a15050565b6001602052806000526040600020600091509050806000015490806001018054610357906107a1565b80601f0160208091040260200160405190810160405280929190818152602001828054610383906107a1565b80156103d05780601f106103a5576101008083540402835291602001916103d0565b820191906000526020600020905b8154815290600101906020018083116103b357829003601f168201915b5050505050908060020160009054906101000a900460ff16905083565b60005481565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61045a82610411565b810181811067ffffffffffffffff8211171561047957610478610422565b5b80604052505050565b600061048c6103f3565b90506104988282610451565b919050565b600067ffffffffffffffff8211156104b8576104b7610422565b5b6104c182610411565b9050602081019050919050565b82818337600083830152505050565b60006104f06104eb8461049d565b610482565b90508281526020810184848401111561050c5761050b61040c565b5b6105178482856104ce565b509392505050565b600082601f83011261053457610533610407565b5b81356105448482602086016104dd565b91505092915050565b600060208284031215610563576105626103fd565b5b600082013567ffffffffffffffff81111561058157610580610402565b5b61058d8482850161051f565b91505092915050565b6000819050919050565b6105a981610596565b81146105b457600080fd5b50565b6000813590506105c6816105a0565b92915050565b6000602082840312156105e2576105e16103fd565b5b60006105f0848285016105b7565b91505092915050565b61060281610596565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610642578082015181840152602081019050610627565b60008484015250505050565b600061065982610608565b6106638185610613565b9350610673818560208601610624565b61067c81610411565b840191505092915050565b60008115159050919050565b61069c81610687565b82525050565b60006060820190506106b760008301866105f9565b81810360208301526106c9818561064e565b90506106d86040830184610693565b949350505050565b60006020820190506106f560008301846105f9565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061073582610596565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610767576107666106fb565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806107b957607f821691505b6020821081036107cc576107cb610772565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026108347fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826107f7565b61083e86836107f7565b95508019841693508086168417925050509392505050565b6000819050919050565b600061087b61087661087184610596565b610856565b610596565b9050919050565b6000819050919050565b61089583610860565b6108a96108a182610882565b848454610804565b825550505050565b600090565b6108be6108b1565b6108c981848461088c565b505050565b5b818110156108ed576108e26000826108b6565b6001810190506108cf565b5050565b601f82111561093257610903816107d2565b61090c846107e7565b8101602085101561091b578190505b61092f610927856107e7565b8301826108ce565b50505b505050565b600082821c905092915050565b600061095560001984600802610937565b1980831691505092915050565b600061096e8383610944565b9150826002028217905092915050565b61098782610608565b67ffffffffffffffff8111156109a05761099f610422565b5b6109aa82546107a1565b6109b58282856108f1565b600060209050601f8311600181146109e857600084156109d6578287015190505b6109e08582610962565b865550610a48565b601f1984166109f6866107d2565b60005b82811015610a1e578489015182556001820191506020850194506020810190506109f9565b86831015610a3b5784890151610a37601f891682610944565b8355505b6001600288020188555050505b505050505050565b6000604082019050610a6560008301856105f9565b610a726020830184610693565b939250505056fea2646970667358221220ebdc7946a9c8417bcd2c02192521e8a72c15bbfc96c4152fff383d765d09de9d64736f6c63430008130033",
}

// TodolistABI is the input ABI used to generate the binding from.
// Deprecated: Use TodolistMetaData.ABI instead.
var TodolistABI = TodolistMetaData.ABI

// TodolistBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodolistMetaData.Bin instead.
var TodolistBin = TodolistMetaData.Bin

// DeployTodolist deploys a new Ethereum contract, binding an instance of Todolist to it.
func DeployTodolist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todolist, error) {
	parsed, err := TodolistMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodolistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todolist{TodolistCaller: TodolistCaller{contract: contract}, TodolistTransactor: TodolistTransactor{contract: contract}, TodolistFilterer: TodolistFilterer{contract: contract}}, nil
}

// Todolist is an auto generated Go binding around an Ethereum contract.
type Todolist struct {
	TodolistCaller     // Read-only binding to the contract
	TodolistTransactor // Write-only binding to the contract
	TodolistFilterer   // Log filterer for contract events
}

// TodolistCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodolistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodolistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodolistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodolistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodolistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodolistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodolistSession struct {
	Contract     *Todolist         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodolistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodolistCallerSession struct {
	Contract *TodolistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TodolistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodolistTransactorSession struct {
	Contract     *TodolistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TodolistRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodolistRaw struct {
	Contract *Todolist // Generic contract binding to access the raw methods on
}

// TodolistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodolistCallerRaw struct {
	Contract *TodolistCaller // Generic read-only contract binding to access the raw methods on
}

// TodolistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodolistTransactorRaw struct {
	Contract *TodolistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodolist creates a new instance of Todolist, bound to a specific deployed contract.
func NewTodolist(address common.Address, backend bind.ContractBackend) (*Todolist, error) {
	contract, err := bindTodolist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todolist{TodolistCaller: TodolistCaller{contract: contract}, TodolistTransactor: TodolistTransactor{contract: contract}, TodolistFilterer: TodolistFilterer{contract: contract}}, nil
}

// NewTodolistCaller creates a new read-only instance of Todolist, bound to a specific deployed contract.
func NewTodolistCaller(address common.Address, caller bind.ContractCaller) (*TodolistCaller, error) {
	contract, err := bindTodolist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodolistCaller{contract: contract}, nil
}

// NewTodolistTransactor creates a new write-only instance of Todolist, bound to a specific deployed contract.
func NewTodolistTransactor(address common.Address, transactor bind.ContractTransactor) (*TodolistTransactor, error) {
	contract, err := bindTodolist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodolistTransactor{contract: contract}, nil
}

// NewTodolistFilterer creates a new log filterer instance of Todolist, bound to a specific deployed contract.
func NewTodolistFilterer(address common.Address, filterer bind.ContractFilterer) (*TodolistFilterer, error) {
	contract, err := bindTodolist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodolistFilterer{contract: contract}, nil
}

// bindTodolist binds a generic wrapper to an already deployed contract.
func bindTodolist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TodolistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todolist *TodolistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todolist.Contract.TodolistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todolist *TodolistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todolist.Contract.TodolistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todolist *TodolistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todolist.Contract.TodolistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todolist *TodolistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todolist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todolist *TodolistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todolist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todolist *TodolistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todolist.Contract.contract.Transact(opts, method, params...)
}

// TaskCount is a free data retrieval call binding the contract method 0xb6cb58a5.
//
// Solidity: function taskCount() view returns(uint256)
func (_Todolist *TodolistCaller) TaskCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Todolist.contract.Call(opts, &out, "taskCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskCount is a free data retrieval call binding the contract method 0xb6cb58a5.
//
// Solidity: function taskCount() view returns(uint256)
func (_Todolist *TodolistSession) TaskCount() (*big.Int, error) {
	return _Todolist.Contract.TaskCount(&_Todolist.CallOpts)
}

// TaskCount is a free data retrieval call binding the contract method 0xb6cb58a5.
//
// Solidity: function taskCount() view returns(uint256)
func (_Todolist *TodolistCallerSession) TaskCount() (*big.Int, error) {
	return _Todolist.Contract.TaskCount(&_Todolist.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, string content, bool completed)
func (_Todolist *TodolistCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Content   string
	Completed bool
}, error) {
	var out []interface{}
	err := _Todolist.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Content   string
		Completed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Content = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Completed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, string content, bool completed)
func (_Todolist *TodolistSession) Tasks(arg0 *big.Int) (struct {
	Id        *big.Int
	Content   string
	Completed bool
}, error) {
	return _Todolist.Contract.Tasks(&_Todolist.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, string content, bool completed)
func (_Todolist *TodolistCallerSession) Tasks(arg0 *big.Int) (struct {
	Id        *big.Int
	Content   string
	Completed bool
}, error) {
	return _Todolist.Contract.Tasks(&_Todolist.CallOpts, arg0)
}

// CreateTask is a paid mutator transaction binding the contract method 0x111002aa.
//
// Solidity: function createTask(string _content) returns()
func (_Todolist *TodolistTransactor) CreateTask(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Todolist.contract.Transact(opts, "createTask", _content)
}

// CreateTask is a paid mutator transaction binding the contract method 0x111002aa.
//
// Solidity: function createTask(string _content) returns()
func (_Todolist *TodolistSession) CreateTask(_content string) (*types.Transaction, error) {
	return _Todolist.Contract.CreateTask(&_Todolist.TransactOpts, _content)
}

// CreateTask is a paid mutator transaction binding the contract method 0x111002aa.
//
// Solidity: function createTask(string _content) returns()
func (_Todolist *TodolistTransactorSession) CreateTask(_content string) (*types.Transaction, error) {
	return _Todolist.Contract.CreateTask(&_Todolist.TransactOpts, _content)
}

// ToggleCompleted is a paid mutator transaction binding the contract method 0x455f5024.
//
// Solidity: function toggleCompleted(uint256 _id) returns()
func (_Todolist *TodolistTransactor) ToggleCompleted(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todolist.contract.Transact(opts, "toggleCompleted", _id)
}

// ToggleCompleted is a paid mutator transaction binding the contract method 0x455f5024.
//
// Solidity: function toggleCompleted(uint256 _id) returns()
func (_Todolist *TodolistSession) ToggleCompleted(_id *big.Int) (*types.Transaction, error) {
	return _Todolist.Contract.ToggleCompleted(&_Todolist.TransactOpts, _id)
}

// ToggleCompleted is a paid mutator transaction binding the contract method 0x455f5024.
//
// Solidity: function toggleCompleted(uint256 _id) returns()
func (_Todolist *TodolistTransactorSession) ToggleCompleted(_id *big.Int) (*types.Transaction, error) {
	return _Todolist.Contract.ToggleCompleted(&_Todolist.TransactOpts, _id)
}

// TodolistTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the Todolist contract.
type TodolistTaskCompletedIterator struct {
	Event *TodolistTaskCompleted // Event containing the contract specifics and raw log

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
func (it *TodolistTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TodolistTaskCompleted)
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
		it.Event = new(TodolistTaskCompleted)
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
func (it *TodolistTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TodolistTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TodolistTaskCompleted represents a TaskCompleted event raised by the Todolist contract.
type TodolistTaskCompleted struct {
	Id        *big.Int
	Completed bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0xe21fa966ca5cd02748c0752352d18c48165e61cb55b4c29cccf924b5a95fcff1.
//
// Solidity: event TaskCompleted(uint256 id, bool completed)
func (_Todolist *TodolistFilterer) FilterTaskCompleted(opts *bind.FilterOpts) (*TodolistTaskCompletedIterator, error) {

	logs, sub, err := _Todolist.contract.FilterLogs(opts, "TaskCompleted")
	if err != nil {
		return nil, err
	}
	return &TodolistTaskCompletedIterator{contract: _Todolist.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0xe21fa966ca5cd02748c0752352d18c48165e61cb55b4c29cccf924b5a95fcff1.
//
// Solidity: event TaskCompleted(uint256 id, bool completed)
func (_Todolist *TodolistFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *TodolistTaskCompleted) (event.Subscription, error) {

	logs, sub, err := _Todolist.contract.WatchLogs(opts, "TaskCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TodolistTaskCompleted)
				if err := _Todolist.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0xe21fa966ca5cd02748c0752352d18c48165e61cb55b4c29cccf924b5a95fcff1.
//
// Solidity: event TaskCompleted(uint256 id, bool completed)
func (_Todolist *TodolistFilterer) ParseTaskCompleted(log types.Log) (*TodolistTaskCompleted, error) {
	event := new(TodolistTaskCompleted)
	if err := _Todolist.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TodolistTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the Todolist contract.
type TodolistTaskCreatedIterator struct {
	Event *TodolistTaskCreated // Event containing the contract specifics and raw log

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
func (it *TodolistTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TodolistTaskCreated)
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
		it.Event = new(TodolistTaskCreated)
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
func (it *TodolistTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TodolistTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TodolistTaskCreated represents a TaskCreated event raised by the Todolist contract.
type TodolistTaskCreated struct {
	Id        *big.Int
	Content   string
	Completed bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x05d0fb833127fc08168556d0e7ca9554fc3f6bc843b3b7d2bf1c35aea6bab660.
//
// Solidity: event TaskCreated(uint256 id, string content, bool completed)
func (_Todolist *TodolistFilterer) FilterTaskCreated(opts *bind.FilterOpts) (*TodolistTaskCreatedIterator, error) {

	logs, sub, err := _Todolist.contract.FilterLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return &TodolistTaskCreatedIterator{contract: _Todolist.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x05d0fb833127fc08168556d0e7ca9554fc3f6bc843b3b7d2bf1c35aea6bab660.
//
// Solidity: event TaskCreated(uint256 id, string content, bool completed)
func (_Todolist *TodolistFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *TodolistTaskCreated) (event.Subscription, error) {

	logs, sub, err := _Todolist.contract.WatchLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TodolistTaskCreated)
				if err := _Todolist.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0x05d0fb833127fc08168556d0e7ca9554fc3f6bc843b3b7d2bf1c35aea6bab660.
//
// Solidity: event TaskCreated(uint256 id, string content, bool completed)
func (_Todolist *TodolistFilterer) ParseTaskCreated(log types.Log) (*TodolistTaskCreated, error) {
	event := new(TodolistTaskCreated)
	if err := _Todolist.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
