// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zkblob

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

// ZkBlobMetaData contains all meta data concerning the ZkBlob contract.
var ZkBlobMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"}],\"name\":\"RepeatedBatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"}],\"name\":\"WrongBatchNum\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"WrongSignature\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"postBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerAddress\",\"type\":\"address\"}],\"name\":\"setSequencerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"verifyBlob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610fa7380380610fa783398101604081905261002e91610120565b600180546001600160a01b0319166001600160a01b0383161790556100535f33610083565b61007d7fac4f1890dc96c9a02330d1fa696648a38f3b282d2449c2d8e6f10507488c84c882610083565b5061014d565b5f828152602081815260408083206001600160a01b038516845290915290205460ff1661011c575f828152602081815260408083206001600160a01b03851684529091529020805460ff191660011790556100db3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b5f60208284031215610130575f80fd5b81516001600160a01b0381168114610146575f80fd5b9392505050565b610e4d8061015a5f395ff3fe608060405234801561000f575f80fd5b50600436106100cb575f3560e01c80634842855c11610088578063a217fddf11610063578063a217fddf146101c7578063c2b40ae4146101ce578063d547741f146101ed578063d5cb711b14610200575f80fd5b80634842855c1461018d57806391d14854146101a15780639755fd6c146101b4575f80fd5b806301ffc9a7146100cf5780632349ac23146100f7578063248a9ca31461010c5780632cee6a6b1461013c5780632f2ff15d1461016757806336568abe1461017a575b5f80fd5b6100e26100dd366004610ab3565b610213565b60405190151581526020015b60405180910390f35b61010a610105366004610aee565b610249565b005b61012e61011a366004610bab565b5f9081526020819052604090206001015490565b6040519081526020016100ee565b60015461014f906001600160a01b031681565b6040516001600160a01b0390911681526020016100ee565b61010a610175366004610bdd565b610317565b61010a610188366004610bdd565b610340565b61012e5f80516020610df883398151915281565b6100e26101af366004610bdd565b6103be565b6100e26101c2366004610c07565b6103e6565b61012e5f81565b61012e6101dc366004610bab565b60026020525f908152604090205481565b61010a6101fb366004610bdd565b61044d565b61010a61020e366004610c83565b610471565b5f6001600160e01b03198216637965db0b60e01b148061024357506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f80516020610df8833981519152610260816104c8565b5f8481526002602052604090205415610294576040516309d8854d60e21b8152600481018590526024015b60405180910390fd5b604080516020808201879052818301869052825180830384018152606090920190925280519101205f6102c782856104d5565b6001549091506001600160a01b03808316911614610302576040516220a52160e11b81526001600160a01b038216600482015260240161028b565b5050505f928352506002602052604090912055565b5f82815260208190526040902060010154610331816104c8565b61033b83836104f7565b505050565b6001600160a01b03811633146103b05760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b606482015260840161028b565b6103ba828261057a565b5050565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b5f8481526002602052604081205480610402575f915050610445565b6104418484808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152508592508991506105de9050565b9150505b949350505050565b5f82815260208190526040902060010154610467816104c8565b61033b838361057a565b5f61047b816104c8565b600180546001600160a01b0319166001600160a01b0384169081179091556104b1905f80516020610df88339815191529061057a565b6103ba5f80516020610df8833981519152836104f7565b6104d281336105f3565b50565b5f805f6104e2858561064c565b915091506104ef8161068e565b509392505050565b61050182826103be565b6103ba575f828152602081815260408083206001600160a01b03851684529091529020805460ff191660011790556105363390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b61058482826103be565b156103ba575f828152602081815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b5f826105ea85846107d7565b14949350505050565b6105fd82826103be565b6103ba5761060a8161081b565b61061583602061082d565b604051602001610626929190610cbe565b60408051601f198184030181529082905262461bcd60e51b825261028b91600401610d32565b5f808251604103610680576020830151604084015160608501515f1a610674878285856109ca565b94509450505050610687565b505f905060025b9250929050565b5f8160048111156106a1576106a1610d64565b036106a95750565b60018160048111156106bd576106bd610d64565b0361070a5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161028b565b600281600481111561071e5761071e610d64565b0361076b5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161028b565b600381600481111561077f5761077f610d64565b036104d25760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161028b565b5f81815b84518110156104ef57610807828683815181106107fa576107fa610d78565b6020026020010151610a87565b91508061081381610da0565b9150506107db565b60606102436001600160a01b03831660145b60605f61083b836002610db8565b610846906002610dcf565b67ffffffffffffffff81111561085e5761085e610ada565b6040519080825280601f01601f191660200182016040528015610888576020820181803683370190505b509050600360fc1b815f815181106108a2576108a2610d78565b60200101906001600160f81b03191690815f1a905350600f60fb1b816001815181106108d0576108d0610d78565b60200101906001600160f81b03191690815f1a9053505f6108f2846002610db8565b6108fd906001610dcf565b90505b6001811115610974576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061093157610931610d78565b1a60f81b82828151811061094757610947610d78565b60200101906001600160f81b03191690815f1a90535060049490941c9361096d81610de2565b9050610900565b5083156109c35760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161028b565b9392505050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156109ff57505f90506003610a7e565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610a50573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b038116610a78575f60019250925050610a7e565b91505f90505b94509492505050565b5f818310610aa1575f8281526020849052604090206109c3565b5f8381526020839052604090206109c3565b5f60208284031215610ac3575f80fd5b81356001600160e01b0319811681146109c3575f80fd5b634e487b7160e01b5f52604160045260245ffd5b5f805f60608486031215610b00575f80fd5b8335925060208401359150604084013567ffffffffffffffff80821115610b25575f80fd5b818601915086601f830112610b38575f80fd5b813581811115610b4a57610b4a610ada565b604051601f8201601f19908116603f01168101908382118183101715610b7257610b72610ada565b81604052828152896020848701011115610b8a575f80fd5b826020860160208301375f6020848301015280955050505050509250925092565b5f60208284031215610bbb575f80fd5b5035919050565b80356001600160a01b0381168114610bd8575f80fd5b919050565b5f8060408385031215610bee575f80fd5b82359150610bfe60208401610bc2565b90509250929050565b5f805f8060608587031215610c1a575f80fd5b8435935060208501359250604085013567ffffffffffffffff80821115610c3f575f80fd5b818701915087601f830112610c52575f80fd5b813581811115610c60575f80fd5b8860208260051b8501011115610c74575f80fd5b95989497505060200194505050565b5f60208284031215610c93575f80fd5b6109c382610bc2565b5f5b83811015610cb6578181015183820152602001610c9e565b50505f910152565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081525f8351610cf5816017850160208801610c9c565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351610d26816028840160208801610c9c565b01602801949350505050565b602081525f8251806020840152610d50816040850160208701610c9c565b601f01601f19169190910160400192915050565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f60018201610db157610db1610d8c565b5060010190565b808202811582820484141761024357610243610d8c565b8082018082111561024357610243610d8c565b5f81610df057610df0610d8c565b505f19019056feac4f1890dc96c9a02330d1fa696648a38f3b282d2449c2d8e6f10507488c84c8a264697066735822122034565dc1b6312e737a599e0eb06e99607cd9fc84b46d87445b0ec31b786fb61e64736f6c63430008140033",
}

// ZkBlobABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkBlobMetaData.ABI instead.
var ZkBlobABI = ZkBlobMetaData.ABI

// ZkBlobBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZkBlobMetaData.Bin instead.
var ZkBlobBin = ZkBlobMetaData.Bin

// DeployZkBlob deploys a new Ethereum contract, binding an instance of ZkBlob to it.
func DeployZkBlob(auth *bind.TransactOpts, backend bind.ContractBackend, _sequencerAddress common.Address) (common.Address, *types.Transaction, *ZkBlob, error) {
	parsed, err := ZkBlobMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZkBlobBin), backend, _sequencerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZkBlob{ZkBlobCaller: ZkBlobCaller{contract: contract}, ZkBlobTransactor: ZkBlobTransactor{contract: contract}, ZkBlobFilterer: ZkBlobFilterer{contract: contract}}, nil
}

// ZkBlob is an auto generated Go binding around an Ethereum contract.
type ZkBlob struct {
	ZkBlobCaller     // Read-only binding to the contract
	ZkBlobTransactor // Write-only binding to the contract
	ZkBlobFilterer   // Log filterer for contract events
}

// ZkBlobCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkBlobCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBlobTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkBlobTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBlobFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkBlobFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBlobSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkBlobSession struct {
	Contract     *ZkBlob           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkBlobCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkBlobCallerSession struct {
	Contract *ZkBlobCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZkBlobTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkBlobTransactorSession struct {
	Contract     *ZkBlobTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkBlobRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkBlobRaw struct {
	Contract *ZkBlob // Generic contract binding to access the raw methods on
}

// ZkBlobCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkBlobCallerRaw struct {
	Contract *ZkBlobCaller // Generic read-only contract binding to access the raw methods on
}

// ZkBlobTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkBlobTransactorRaw struct {
	Contract *ZkBlobTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkBlob creates a new instance of ZkBlob, bound to a specific deployed contract.
func NewZkBlob(address common.Address, backend bind.ContractBackend) (*ZkBlob, error) {
	contract, err := bindZkBlob(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkBlob{ZkBlobCaller: ZkBlobCaller{contract: contract}, ZkBlobTransactor: ZkBlobTransactor{contract: contract}, ZkBlobFilterer: ZkBlobFilterer{contract: contract}}, nil
}

// NewZkBlobCaller creates a new read-only instance of ZkBlob, bound to a specific deployed contract.
func NewZkBlobCaller(address common.Address, caller bind.ContractCaller) (*ZkBlobCaller, error) {
	contract, err := bindZkBlob(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkBlobCaller{contract: contract}, nil
}

// NewZkBlobTransactor creates a new write-only instance of ZkBlob, bound to a specific deployed contract.
func NewZkBlobTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkBlobTransactor, error) {
	contract, err := bindZkBlob(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkBlobTransactor{contract: contract}, nil
}

// NewZkBlobFilterer creates a new log filterer instance of ZkBlob, bound to a specific deployed contract.
func NewZkBlobFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkBlobFilterer, error) {
	contract, err := bindZkBlob(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkBlobFilterer{contract: contract}, nil
}

// bindZkBlob binds a generic wrapper to an already deployed contract.
func bindZkBlob(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZkBlobMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkBlob *ZkBlobRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkBlob.Contract.ZkBlobCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkBlob *ZkBlobRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBlob.Contract.ZkBlobTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkBlob *ZkBlobRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkBlob.Contract.ZkBlobTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkBlob *ZkBlobCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkBlob.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkBlob *ZkBlobTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBlob.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkBlob *ZkBlobTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkBlob.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZkBlob.Contract.DEFAULTADMINROLE(&_ZkBlob.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZkBlob.Contract.DEFAULTADMINROLE(&_ZkBlob.CallOpts)
}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobCaller) SEQUENCERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "SEQUENCER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobSession) SEQUENCERROLE() ([32]byte, error) {
	return _ZkBlob.Contract.SEQUENCERROLE(&_ZkBlob.CallOpts)
}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobCallerSession) SEQUENCERROLE() ([32]byte, error) {
	return _ZkBlob.Contract.SEQUENCERROLE(&_ZkBlob.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZkBlob *ZkBlobCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZkBlob *ZkBlobSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZkBlob.Contract.GetRoleAdmin(&_ZkBlob.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZkBlob *ZkBlobCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZkBlob.Contract.GetRoleAdmin(&_ZkBlob.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZkBlob *ZkBlobCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZkBlob *ZkBlobSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZkBlob.Contract.HasRole(&_ZkBlob.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZkBlob *ZkBlobCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZkBlob.Contract.HasRole(&_ZkBlob.CallOpts, role, account)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_ZkBlob *ZkBlobCaller) Roots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_ZkBlob *ZkBlobSession) Roots(arg0 *big.Int) ([32]byte, error) {
	return _ZkBlob.Contract.Roots(&_ZkBlob.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_ZkBlob *ZkBlobCallerSession) Roots(arg0 *big.Int) ([32]byte, error) {
	return _ZkBlob.Contract.Roots(&_ZkBlob.CallOpts, arg0)
}

// SequencerAddress is a free data retrieval call binding the contract method 0x2cee6a6b.
//
// Solidity: function sequencerAddress() view returns(address)
func (_ZkBlob *ZkBlobCaller) SequencerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "sequencerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerAddress is a free data retrieval call binding the contract method 0x2cee6a6b.
//
// Solidity: function sequencerAddress() view returns(address)
func (_ZkBlob *ZkBlobSession) SequencerAddress() (common.Address, error) {
	return _ZkBlob.Contract.SequencerAddress(&_ZkBlob.CallOpts)
}

// SequencerAddress is a free data retrieval call binding the contract method 0x2cee6a6b.
//
// Solidity: function sequencerAddress() view returns(address)
func (_ZkBlob *ZkBlobCallerSession) SequencerAddress() (common.Address, error) {
	return _ZkBlob.Contract.SequencerAddress(&_ZkBlob.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZkBlob *ZkBlobCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZkBlob *ZkBlobSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZkBlob.Contract.SupportsInterface(&_ZkBlob.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZkBlob *ZkBlobCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZkBlob.Contract.SupportsInterface(&_ZkBlob.CallOpts, interfaceId)
}

// VerifyBlob is a free data retrieval call binding the contract method 0x9755fd6c.
//
// Solidity: function verifyBlob(uint256 batchNum, bytes32 hash, bytes32[] merkleProof) view returns(bool)
func (_ZkBlob *ZkBlobCaller) VerifyBlob(opts *bind.CallOpts, batchNum *big.Int, hash [32]byte, merkleProof [][32]byte) (bool, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "verifyBlob", batchNum, hash, merkleProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlob is a free data retrieval call binding the contract method 0x9755fd6c.
//
// Solidity: function verifyBlob(uint256 batchNum, bytes32 hash, bytes32[] merkleProof) view returns(bool)
func (_ZkBlob *ZkBlobSession) VerifyBlob(batchNum *big.Int, hash [32]byte, merkleProof [][32]byte) (bool, error) {
	return _ZkBlob.Contract.VerifyBlob(&_ZkBlob.CallOpts, batchNum, hash, merkleProof)
}

// VerifyBlob is a free data retrieval call binding the contract method 0x9755fd6c.
//
// Solidity: function verifyBlob(uint256 batchNum, bytes32 hash, bytes32[] merkleProof) view returns(bool)
func (_ZkBlob *ZkBlobCallerSession) VerifyBlob(batchNum *big.Int, hash [32]byte, merkleProof [][32]byte) (bool, error) {
	return _ZkBlob.Contract.VerifyBlob(&_ZkBlob.CallOpts, batchNum, hash, merkleProof)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.GrantRole(&_ZkBlob.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.GrantRole(&_ZkBlob.TransactOpts, role, account)
}

// PostBatch is a paid mutator transaction binding the contract method 0x2349ac23.
//
// Solidity: function postBatch(uint256 batchNum, bytes32 root, bytes signature) returns()
func (_ZkBlob *ZkBlobTransactor) PostBatch(opts *bind.TransactOpts, batchNum *big.Int, root [32]byte, signature []byte) (*types.Transaction, error) {
	return _ZkBlob.contract.Transact(opts, "postBatch", batchNum, root, signature)
}

// PostBatch is a paid mutator transaction binding the contract method 0x2349ac23.
//
// Solidity: function postBatch(uint256 batchNum, bytes32 root, bytes signature) returns()
func (_ZkBlob *ZkBlobSession) PostBatch(batchNum *big.Int, root [32]byte, signature []byte) (*types.Transaction, error) {
	return _ZkBlob.Contract.PostBatch(&_ZkBlob.TransactOpts, batchNum, root, signature)
}

// PostBatch is a paid mutator transaction binding the contract method 0x2349ac23.
//
// Solidity: function postBatch(uint256 batchNum, bytes32 root, bytes signature) returns()
func (_ZkBlob *ZkBlobTransactorSession) PostBatch(batchNum *big.Int, root [32]byte, signature []byte) (*types.Transaction, error) {
	return _ZkBlob.Contract.PostBatch(&_ZkBlob.TransactOpts, batchNum, root, signature)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.RenounceRole(&_ZkBlob.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.RenounceRole(&_ZkBlob.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.RevokeRole(&_ZkBlob.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.RevokeRole(&_ZkBlob.TransactOpts, role, account)
}

// SetSequencerAddress is a paid mutator transaction binding the contract method 0xd5cb711b.
//
// Solidity: function setSequencerAddress(address _sequencerAddress) returns()
func (_ZkBlob *ZkBlobTransactor) SetSequencerAddress(opts *bind.TransactOpts, _sequencerAddress common.Address) (*types.Transaction, error) {
	return _ZkBlob.contract.Transact(opts, "setSequencerAddress", _sequencerAddress)
}

// SetSequencerAddress is a paid mutator transaction binding the contract method 0xd5cb711b.
//
// Solidity: function setSequencerAddress(address _sequencerAddress) returns()
func (_ZkBlob *ZkBlobSession) SetSequencerAddress(_sequencerAddress common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.SetSequencerAddress(&_ZkBlob.TransactOpts, _sequencerAddress)
}

// SetSequencerAddress is a paid mutator transaction binding the contract method 0xd5cb711b.
//
// Solidity: function setSequencerAddress(address _sequencerAddress) returns()
func (_ZkBlob *ZkBlobTransactorSession) SetSequencerAddress(_sequencerAddress common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.SetSequencerAddress(&_ZkBlob.TransactOpts, _sequencerAddress)
}

// ZkBlobRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZkBlob contract.
type ZkBlobRoleAdminChangedIterator struct {
	Event *ZkBlobRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZkBlobRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBlobRoleAdminChanged)
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
		it.Event = new(ZkBlobRoleAdminChanged)
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
func (it *ZkBlobRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBlobRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBlobRoleAdminChanged represents a RoleAdminChanged event raised by the ZkBlob contract.
type ZkBlobRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZkBlob *ZkBlobFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZkBlobRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZkBlob.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZkBlobRoleAdminChangedIterator{contract: _ZkBlob.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZkBlob *ZkBlobFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZkBlobRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZkBlob.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBlobRoleAdminChanged)
				if err := _ZkBlob.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZkBlob *ZkBlobFilterer) ParseRoleAdminChanged(log types.Log) (*ZkBlobRoleAdminChanged, error) {
	event := new(ZkBlobRoleAdminChanged)
	if err := _ZkBlob.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBlobRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZkBlob contract.
type ZkBlobRoleGrantedIterator struct {
	Event *ZkBlobRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZkBlobRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBlobRoleGranted)
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
		it.Event = new(ZkBlobRoleGranted)
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
func (it *ZkBlobRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBlobRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBlobRoleGranted represents a RoleGranted event raised by the ZkBlob contract.
type ZkBlobRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZkBlobRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkBlob.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZkBlobRoleGrantedIterator{contract: _ZkBlob.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZkBlobRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkBlob.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBlobRoleGranted)
				if err := _ZkBlob.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) ParseRoleGranted(log types.Log) (*ZkBlobRoleGranted, error) {
	event := new(ZkBlobRoleGranted)
	if err := _ZkBlob.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBlobRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZkBlob contract.
type ZkBlobRoleRevokedIterator struct {
	Event *ZkBlobRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZkBlobRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBlobRoleRevoked)
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
		it.Event = new(ZkBlobRoleRevoked)
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
func (it *ZkBlobRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBlobRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBlobRoleRevoked represents a RoleRevoked event raised by the ZkBlob contract.
type ZkBlobRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZkBlobRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkBlob.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZkBlobRoleRevokedIterator{contract: _ZkBlob.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZkBlobRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkBlob.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBlobRoleRevoked)
				if err := _ZkBlob.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) ParseRoleRevoked(log types.Log) (*ZkBlobRoleRevoked, error) {
	event := new(ZkBlobRoleRevoked)
	if err := _ZkBlob.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
