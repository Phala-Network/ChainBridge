// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PhalaBTCLottery

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
)

// PhalaBTCLotteryMetaData contains all meta data concerning the PhalaBTCLottery contract.
var PhalaBTCLotteryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_genericHandler\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"roundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"winnerCount\",\"type\":\"uint32\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"roundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddr\",\"type\":\"string\"}],\"name\":\"OpenLottery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"PayloadStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"genericHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"openedBox\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payloadSequence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payloadStorage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAdmin\",\"type\":\"address\"}],\"name\":\"setNFTAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"depositHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"roundId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"}],\"name\":\"isLotteryOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620017eb380380620017eb833981810160405281019062000037919062000118565b60008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000192565b600081519050620001128162000178565b92915050565b6000602082840312156200012b57600080fd5b60006200013b8482850162000101565b91505092915050565b6000620001518262000158565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620001838162000144565b81146200018f57600080fd5b50565b61164980620001a26000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b146101645780639360bfad14610182578063d4e3aade146101a0578063e24e8027146101bc578063f2db3797146101da578063f4190f831461020a576100a9565b806312aad3d4146100ae57806313af4035146100de5780632a52ccfb146100fa57806336954720146101185780635a80d92314610148575b600080fd5b6100c860048036038101906100c39190610e42565b610226565b6040516100d59190611285565b60405180910390f35b6100f860048036038101906100f39190610daf565b610255565b005b610102610399565b60405161010f9190611402565b60405180910390f35b610132600480360381019061012d9190610e19565b61039f565b60405161013f91906112a0565b60405180910390f35b610162600480360381019061015d9190610daf565b61044f565b005b61016c610592565b604051610179919061126a565b60405180910390f35b61018a6105b8565b604051610197919061126a565b60405180910390f35b6101ba60048036038101906101b59190610dd8565b6105dc565b005b6101c4610678565b6040516101d1919061126a565b60405180910390f35b6101f460048036038101906101ef9190610e42565b61069e565b6040516102019190611285565b60405180910390f35b610224600480360381019061021f9190610dd8565b6106f2565b005b60046020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102dc90611382565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610355576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034c90611342565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60035481565b60056020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104475780601f1061041c57610100808354040283529160200191610447565b820191906000526020600020905b81548152906001019060200180831161042a57829003601f168201915b505050505081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d690611382565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561054f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610546906112c2565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610663906112e2565b60405180910390fd5b610675816107ff565b50565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600460008463ffffffff1663ffffffff16815260200190815260200160002060008363ffffffff1663ffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600061070860008361087190919063ffffffff16565b905060008160ff16141561077157600061072c6001846108cf90919063ffffffff16565b905060006107446005856108cf90919063ffffffff16565b9050600061075c6009866108cf90919063ffffffff16565b905061076983838361092d565b5050506107fb565b60018160ff1614156107f95760006107936001846108cf90919063ffffffff16565b905060006107ab6005856108cf90919063ffffffff16565b905060006107c36009866108cf90919063ffffffff16565b905060606107e3600d8363ffffffff1688610a549092919063ffffffff16565b90506107f0848483610b60565b505050506107fa565b5b5b5050565b8060056000600360008154809291906001019190505581526020019081526020016000209080519060200190610836929190610c7f565b507f2216a25e8e0ba35d0db376e249196115dad754519da67974ef6172c2c509381e8160405161086691906112a0565b60405180910390a150565b600060018201835110156108ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b1906113c2565b60405180910390fd5b60008260018501015190508091505092915050565b60006004820183511015610918576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090f906113e2565b60405180910390fd5b60008260048501015190508091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141580156109d5575060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16145b610a14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0b90611302565b60405180910390fd5b7fdb03f7d591b8e310b058b3c3c80e9415acee180707443368d8bae2d553831240838383604051610a479392919061145b565b60405180910390a1505050565b606081601f83011015610a9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9390611322565b60405180910390fd5b81830184511015610ae2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad9906113a2565b60405180910390fd5b6060821560008114610b035760405191506000825260208201604052610b54565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610b415780518352602083019250602081019050610b24565b50868552601f19601f8301166040525050505b50809150509392505050565b600460008463ffffffff1663ffffffff16815260200190815260200160002060008363ffffffff1663ffffffff16815260200190815260200160002060009054906101000a900460ff1615610bea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be190611362565b60405180910390fd5b6001600460008563ffffffff1663ffffffff16815260200190815260200160002060008463ffffffff1663ffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f9cebd1f828553e091f5b72c5bd09ee678b2c0e326f8a3d2f93efb9a63062ee92838383604051610c729392919061141d565b60405180910390a1505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610cc057805160ff1916838001178555610cee565b82800160010185558215610cee579182015b82811115610ced578251825591602001919060010190610cd2565b5b509050610cfb9190610cff565b5090565b5b80821115610d18576000816000905550600101610d00565b5090565b600081359050610d2b816115ce565b92915050565b600082601f830112610d4257600080fd5b8135610d55610d50826114bf565b611492565b91508082526020830160208301858383011115610d7157600080fd5b610d7c83828461157b565b50505092915050565b600081359050610d94816115e5565b92915050565b600081359050610da9816115fc565b92915050565b600060208284031215610dc157600080fd5b6000610dcf84828501610d1c565b91505092915050565b600060208284031215610dea57600080fd5b600082013567ffffffffffffffff811115610e0457600080fd5b610e1084828501610d31565b91505092915050565b600060208284031215610e2b57600080fd5b6000610e3984828501610d85565b91505092915050565b60008060408385031215610e5557600080fd5b6000610e6385828601610d9a565b9250506020610e7485828601610d9a565b9150509250929050565b610e8781611523565b82525050565b610e9681611535565b82525050565b6000610ea7826114eb565b610eb18185611501565b9350610ec181856020860161158a565b610eca816115bd565b840191505092915050565b6000610ee0826114f6565b610eea8185611512565b9350610efa81856020860161158a565b610f03816115bd565b840191505092915050565b6000610f1b601983611512565b91507f496e76616c6964204e46542061646d696e2061646472657373000000000000006000830152602082019050919050565b6000610f5b604383611512565b91507f5065726d697373696f6e2044656e6965643a204d6573736167652053656e646560008301527f722073686f756c642062652047656e6572696348616e646c657220636f6e747260208301527f61637400000000000000000000000000000000000000000000000000000000006040830152606082019050919050565b6000610fe7603383611512565b91507f5065726d697373696f6e2044656e6965643a205478204f726967696e2073686f60008301527f756c64206265204e465420636f6e7472616374000000000000000000000000006020830152604082019050919050565b600061104d600e83611512565b91507f736c6963655f6f766572666c6f770000000000000000000000000000000000006000830152602082019050919050565b600061108d601583611512565b91507f496e76616c6964206f776e6572206164647265737300000000000000000000006000830152602082019050919050565b60006110cd602483611512565b91507f496e76616c69642043616c6c3a2042544320626f7820616c7265616479206f7060008301527f656e6564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611133603d83611512565b91507f5065726d697373696f6e2044656e6965643a204d6573736167652053656e646560008301527f722073686f756c64206265206f776e6572206f6620636f6e74726163740000006020830152604082019050919050565b6000611199601183611512565b91507f736c6963655f6f75744f66426f756e64730000000000000000000000000000006000830152602082019050919050565b60006111d9601383611512565b91507f746f55696e74385f6f75744f66426f756e6473000000000000000000000000006000830152602082019050919050565b6000611219601483611512565b91507f746f55696e7433325f6f75744f66426f756e64730000000000000000000000006000830152602082019050919050565b61125581611561565b82525050565b6112648161156b565b82525050565b600060208201905061127f6000830184610e7e565b92915050565b600060208201905061129a6000830184610e8d565b92915050565b600060208201905081810360008301526112ba8184610e9c565b905092915050565b600060208201905081810360008301526112db81610f0e565b9050919050565b600060208201905081810360008301526112fb81610f4e565b9050919050565b6000602082019050818103600083015261131b81610fda565b9050919050565b6000602082019050818103600083015261133b81611040565b9050919050565b6000602082019050818103600083015261135b81611080565b9050919050565b6000602082019050818103600083015261137b816110c0565b9050919050565b6000602082019050818103600083015261139b81611126565b9050919050565b600060208201905081810360008301526113bb8161118c565b9050919050565b600060208201905081810360008301526113db816111cc565b9050919050565b600060208201905081810360008301526113fb8161120c565b9050919050565b6000602082019050611417600083018461124c565b92915050565b6000606082019050611432600083018661125b565b61143f602083018561125b565b81810360408301526114518184610ed5565b9050949350505050565b6000606082019050611470600083018661125b565b61147d602083018561125b565b61148a604083018461125b565b949350505050565b6000604051905081810181811067ffffffffffffffff821117156114b557600080fd5b8060405250919050565b600067ffffffffffffffff8211156114d657600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061152e82611541565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b82818337600083830152505050565b60005b838110156115a857808201518184015260208101905061158d565b838111156115b7576000848401525b50505050565b6000601f19601f8301169050919050565b6115d781611523565b81146115e257600080fd5b50565b6115ee81611561565b81146115f957600080fd5b50565b6116058161156b565b811461161057600080fd5b5056fea2646970667358221220b05fb57994349b4ecef4ca1ea50d9a92568271e92ba271e3581d8d4a6cfde53564736f6c63430007000033",
}

// PhalaBTCLotteryABI is the input ABI used to generate the binding from.
// Deprecated: Use PhalaBTCLotteryMetaData.ABI instead.
var PhalaBTCLotteryABI = PhalaBTCLotteryMetaData.ABI

// PhalaBTCLotteryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PhalaBTCLotteryMetaData.Bin instead.
var PhalaBTCLotteryBin = PhalaBTCLotteryMetaData.Bin

// DeployPhalaBTCLottery deploys a new Ethereum contract, binding an instance of PhalaBTCLottery to it.
func DeployPhalaBTCLottery(auth *bind.TransactOpts, backend bind.ContractBackend, _genericHandler common.Address) (common.Address, *types.Transaction, *PhalaBTCLottery, error) {
	parsed, err := PhalaBTCLotteryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PhalaBTCLotteryBin), backend, _genericHandler)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PhalaBTCLottery{PhalaBTCLotteryCaller: PhalaBTCLotteryCaller{contract: contract}, PhalaBTCLotteryTransactor: PhalaBTCLotteryTransactor{contract: contract}, PhalaBTCLotteryFilterer: PhalaBTCLotteryFilterer{contract: contract}}, nil
}

// PhalaBTCLottery is an auto generated Go binding around an Ethereum contract.
type PhalaBTCLottery struct {
	PhalaBTCLotteryCaller     // Read-only binding to the contract
	PhalaBTCLotteryTransactor // Write-only binding to the contract
	PhalaBTCLotteryFilterer   // Log filterer for contract events
}

// PhalaBTCLotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PhalaBTCLotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PhalaBTCLotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PhalaBTCLotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PhalaBTCLotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PhalaBTCLotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PhalaBTCLotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PhalaBTCLotterySession struct {
	Contract     *PhalaBTCLottery  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PhalaBTCLotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PhalaBTCLotteryCallerSession struct {
	Contract *PhalaBTCLotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PhalaBTCLotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PhalaBTCLotteryTransactorSession struct {
	Contract     *PhalaBTCLotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PhalaBTCLotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PhalaBTCLotteryRaw struct {
	Contract *PhalaBTCLottery // Generic contract binding to access the raw methods on
}

// PhalaBTCLotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PhalaBTCLotteryCallerRaw struct {
	Contract *PhalaBTCLotteryCaller // Generic read-only contract binding to access the raw methods on
}

// PhalaBTCLotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PhalaBTCLotteryTransactorRaw struct {
	Contract *PhalaBTCLotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPhalaBTCLottery creates a new instance of PhalaBTCLottery, bound to a specific deployed contract.
func NewPhalaBTCLottery(address common.Address, backend bind.ContractBackend) (*PhalaBTCLottery, error) {
	contract, err := bindPhalaBTCLottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLottery{PhalaBTCLotteryCaller: PhalaBTCLotteryCaller{contract: contract}, PhalaBTCLotteryTransactor: PhalaBTCLotteryTransactor{contract: contract}, PhalaBTCLotteryFilterer: PhalaBTCLotteryFilterer{contract: contract}}, nil
}

// NewPhalaBTCLotteryCaller creates a new read-only instance of PhalaBTCLottery, bound to a specific deployed contract.
func NewPhalaBTCLotteryCaller(address common.Address, caller bind.ContractCaller) (*PhalaBTCLotteryCaller, error) {
	contract, err := bindPhalaBTCLottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryCaller{contract: contract}, nil
}

// NewPhalaBTCLotteryTransactor creates a new write-only instance of PhalaBTCLottery, bound to a specific deployed contract.
func NewPhalaBTCLotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*PhalaBTCLotteryTransactor, error) {
	contract, err := bindPhalaBTCLottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryTransactor{contract: contract}, nil
}

// NewPhalaBTCLotteryFilterer creates a new log filterer instance of PhalaBTCLottery, bound to a specific deployed contract.
func NewPhalaBTCLotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*PhalaBTCLotteryFilterer, error) {
	contract, err := bindPhalaBTCLottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryFilterer{contract: contract}, nil
}

// bindPhalaBTCLottery binds a generic wrapper to an already deployed contract.
func bindPhalaBTCLottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PhalaBTCLotteryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PhalaBTCLottery *PhalaBTCLotteryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PhalaBTCLottery.Contract.PhalaBTCLotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PhalaBTCLottery *PhalaBTCLotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.PhalaBTCLotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PhalaBTCLottery *PhalaBTCLotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.PhalaBTCLotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PhalaBTCLottery *PhalaBTCLotteryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PhalaBTCLottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PhalaBTCLottery *PhalaBTCLotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PhalaBTCLottery *PhalaBTCLotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.contract.Transact(opts, method, params...)
}

// GenericHandler is a free data retrieval call binding the contract method 0xe24e8027.
//
// Solidity: function genericHandler() view returns(address)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) GenericHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "genericHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GenericHandler is a free data retrieval call binding the contract method 0xe24e8027.
//
// Solidity: function genericHandler() view returns(address)
func (_PhalaBTCLottery *PhalaBTCLotterySession) GenericHandler() (common.Address, error) {
	return _PhalaBTCLottery.Contract.GenericHandler(&_PhalaBTCLottery.CallOpts)
}

// GenericHandler is a free data retrieval call binding the contract method 0xe24e8027.
//
// Solidity: function genericHandler() view returns(address)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) GenericHandler() (common.Address, error) {
	return _PhalaBTCLottery.Contract.GenericHandler(&_PhalaBTCLottery.CallOpts)
}

// IsLotteryOpened is a free data retrieval call binding the contract method 0xf2db3797.
//
// Solidity: function isLotteryOpened(uint32 roundId, uint32 tokenId) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) IsLotteryOpened(opts *bind.CallOpts, roundId uint32, tokenId uint32) (bool, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "isLotteryOpened", roundId, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLotteryOpened is a free data retrieval call binding the contract method 0xf2db3797.
//
// Solidity: function isLotteryOpened(uint32 roundId, uint32 tokenId) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotterySession) IsLotteryOpened(roundId uint32, tokenId uint32) (bool, error) {
	return _PhalaBTCLottery.Contract.IsLotteryOpened(&_PhalaBTCLottery.CallOpts, roundId, tokenId)
}

// IsLotteryOpened is a free data retrieval call binding the contract method 0xf2db3797.
//
// Solidity: function isLotteryOpened(uint32 roundId, uint32 tokenId) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) IsLotteryOpened(roundId uint32, tokenId uint32) (bool, error) {
	return _PhalaBTCLottery.Contract.IsLotteryOpened(&_PhalaBTCLottery.CallOpts, roundId, tokenId)
}

// NftAdmin is a free data retrieval call binding the contract method 0x9360bfad.
//
// Solidity: function nftAdmin() view returns(address)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) NftAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "nftAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftAdmin is a free data retrieval call binding the contract method 0x9360bfad.
//
// Solidity: function nftAdmin() view returns(address)
func (_PhalaBTCLottery *PhalaBTCLotterySession) NftAdmin() (common.Address, error) {
	return _PhalaBTCLottery.Contract.NftAdmin(&_PhalaBTCLottery.CallOpts)
}

// NftAdmin is a free data retrieval call binding the contract method 0x9360bfad.
//
// Solidity: function nftAdmin() view returns(address)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) NftAdmin() (common.Address, error) {
	return _PhalaBTCLottery.Contract.NftAdmin(&_PhalaBTCLottery.CallOpts)
}

// OpenedBox is a free data retrieval call binding the contract method 0x12aad3d4.
//
// Solidity: function openedBox(uint32 , uint32 ) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) OpenedBox(opts *bind.CallOpts, arg0 uint32, arg1 uint32) (bool, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "openedBox", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OpenedBox is a free data retrieval call binding the contract method 0x12aad3d4.
//
// Solidity: function openedBox(uint32 , uint32 ) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotterySession) OpenedBox(arg0 uint32, arg1 uint32) (bool, error) {
	return _PhalaBTCLottery.Contract.OpenedBox(&_PhalaBTCLottery.CallOpts, arg0, arg1)
}

// OpenedBox is a free data retrieval call binding the contract method 0x12aad3d4.
//
// Solidity: function openedBox(uint32 , uint32 ) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) OpenedBox(arg0 uint32, arg1 uint32) (bool, error) {
	return _PhalaBTCLottery.Contract.OpenedBox(&_PhalaBTCLottery.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PhalaBTCLottery *PhalaBTCLotterySession) Owner() (common.Address, error) {
	return _PhalaBTCLottery.Contract.Owner(&_PhalaBTCLottery.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) Owner() (common.Address, error) {
	return _PhalaBTCLottery.Contract.Owner(&_PhalaBTCLottery.CallOpts)
}

// PayloadSequence is a free data retrieval call binding the contract method 0x2a52ccfb.
//
// Solidity: function payloadSequence() view returns(uint256)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) PayloadSequence(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "payloadSequence")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayloadSequence is a free data retrieval call binding the contract method 0x2a52ccfb.
//
// Solidity: function payloadSequence() view returns(uint256)
func (_PhalaBTCLottery *PhalaBTCLotterySession) PayloadSequence() (*big.Int, error) {
	return _PhalaBTCLottery.Contract.PayloadSequence(&_PhalaBTCLottery.CallOpts)
}

// PayloadSequence is a free data retrieval call binding the contract method 0x2a52ccfb.
//
// Solidity: function payloadSequence() view returns(uint256)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) PayloadSequence() (*big.Int, error) {
	return _PhalaBTCLottery.Contract.PayloadSequence(&_PhalaBTCLottery.CallOpts)
}

// PayloadStorage is a free data retrieval call binding the contract method 0x36954720.
//
// Solidity: function payloadStorage(uint256 ) view returns(bytes)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) PayloadStorage(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "payloadStorage", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PayloadStorage is a free data retrieval call binding the contract method 0x36954720.
//
// Solidity: function payloadStorage(uint256 ) view returns(bytes)
func (_PhalaBTCLottery *PhalaBTCLotterySession) PayloadStorage(arg0 *big.Int) ([]byte, error) {
	return _PhalaBTCLottery.Contract.PayloadStorage(&_PhalaBTCLottery.CallOpts, arg0)
}

// PayloadStorage is a free data retrieval call binding the contract method 0x36954720.
//
// Solidity: function payloadStorage(uint256 ) view returns(bytes)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) PayloadStorage(arg0 *big.Int) ([]byte, error) {
	return _PhalaBTCLottery.Contract.PayloadStorage(&_PhalaBTCLottery.CallOpts, arg0)
}

// DepositHandler is a paid mutator transaction binding the contract method 0xf4190f83.
//
// Solidity: function depositHandler(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactor) DepositHandler(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.contract.Transact(opts, "depositHandler", data)
}

// DepositHandler is a paid mutator transaction binding the contract method 0xf4190f83.
//
// Solidity: function depositHandler(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotterySession) DepositHandler(data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.DepositHandler(&_PhalaBTCLottery.TransactOpts, data)
}

// DepositHandler is a paid mutator transaction binding the contract method 0xf4190f83.
//
// Solidity: function depositHandler(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactorSession) DepositHandler(data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.DepositHandler(&_PhalaBTCLottery.TransactOpts, data)
}

// ExecuteHandler is a paid mutator transaction binding the contract method 0xd4e3aade.
//
// Solidity: function executeHandler(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactor) ExecuteHandler(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.contract.Transact(opts, "executeHandler", data)
}

// ExecuteHandler is a paid mutator transaction binding the contract method 0xd4e3aade.
//
// Solidity: function executeHandler(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotterySession) ExecuteHandler(data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.ExecuteHandler(&_PhalaBTCLottery.TransactOpts, data)
}

// ExecuteHandler is a paid mutator transaction binding the contract method 0xd4e3aade.
//
// Solidity: function executeHandler(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactorSession) ExecuteHandler(data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.ExecuteHandler(&_PhalaBTCLottery.TransactOpts, data)
}

// SetNFTAdmin is a paid mutator transaction binding the contract method 0x5a80d923.
//
// Solidity: function setNFTAdmin(address _nftAdmin) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactor) SetNFTAdmin(opts *bind.TransactOpts, _nftAdmin common.Address) (*types.Transaction, error) {
	return _PhalaBTCLottery.contract.Transact(opts, "setNFTAdmin", _nftAdmin)
}

// SetNFTAdmin is a paid mutator transaction binding the contract method 0x5a80d923.
//
// Solidity: function setNFTAdmin(address _nftAdmin) returns()
func (_PhalaBTCLottery *PhalaBTCLotterySession) SetNFTAdmin(_nftAdmin common.Address) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.SetNFTAdmin(&_PhalaBTCLottery.TransactOpts, _nftAdmin)
}

// SetNFTAdmin is a paid mutator transaction binding the contract method 0x5a80d923.
//
// Solidity: function setNFTAdmin(address _nftAdmin) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactorSession) SetNFTAdmin(_nftAdmin common.Address) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.SetNFTAdmin(&_PhalaBTCLottery.TransactOpts, _nftAdmin)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _PhalaBTCLottery.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PhalaBTCLottery *PhalaBTCLotterySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.SetOwner(&_PhalaBTCLottery.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.SetOwner(&_PhalaBTCLottery.TransactOpts, _owner)
}

// PhalaBTCLotteryNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the PhalaBTCLottery contract.
type PhalaBTCLotteryNewRoundIterator struct {
	Event *PhalaBTCLotteryNewRound // Event containing the contract specifics and raw log

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
func (it *PhalaBTCLotteryNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PhalaBTCLotteryNewRound)
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
		it.Event = new(PhalaBTCLotteryNewRound)
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
func (it *PhalaBTCLotteryNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PhalaBTCLotteryNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PhalaBTCLotteryNewRound represents a NewRound event raised by the PhalaBTCLottery contract.
type PhalaBTCLotteryNewRound struct {
	RoundId     uint32
	TotalCount  uint32
	WinnerCount uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0xdb03f7d591b8e310b058b3c3c80e9415acee180707443368d8bae2d553831240.
//
// Solidity: event NewRound(uint32 roundId, uint32 totalCount, uint32 winnerCount)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) FilterNewRound(opts *bind.FilterOpts) (*PhalaBTCLotteryNewRoundIterator, error) {

	logs, sub, err := _PhalaBTCLottery.contract.FilterLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryNewRoundIterator{contract: _PhalaBTCLottery.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0xdb03f7d591b8e310b058b3c3c80e9415acee180707443368d8bae2d553831240.
//
// Solidity: event NewRound(uint32 roundId, uint32 totalCount, uint32 winnerCount)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *PhalaBTCLotteryNewRound) (event.Subscription, error) {

	logs, sub, err := _PhalaBTCLottery.contract.WatchLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PhalaBTCLotteryNewRound)
				if err := _PhalaBTCLottery.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0xdb03f7d591b8e310b058b3c3c80e9415acee180707443368d8bae2d553831240.
//
// Solidity: event NewRound(uint32 roundId, uint32 totalCount, uint32 winnerCount)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) ParseNewRound(log types.Log) (*PhalaBTCLotteryNewRound, error) {
	event := new(PhalaBTCLotteryNewRound)
	if err := _PhalaBTCLottery.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PhalaBTCLotteryOpenLotteryIterator is returned from FilterOpenLottery and is used to iterate over the raw logs and unpacked data for OpenLottery events raised by the PhalaBTCLottery contract.
type PhalaBTCLotteryOpenLotteryIterator struct {
	Event *PhalaBTCLotteryOpenLottery // Event containing the contract specifics and raw log

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
func (it *PhalaBTCLotteryOpenLotteryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PhalaBTCLotteryOpenLottery)
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
		it.Event = new(PhalaBTCLotteryOpenLottery)
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
func (it *PhalaBTCLotteryOpenLotteryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PhalaBTCLotteryOpenLotteryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PhalaBTCLotteryOpenLottery represents a OpenLottery event raised by the PhalaBTCLottery contract.
type PhalaBTCLotteryOpenLottery struct {
	RoundId uint32
	TokenId uint32
	BtcAddr string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpenLottery is a free log retrieval operation binding the contract event 0x9cebd1f828553e091f5b72c5bd09ee678b2c0e326f8a3d2f93efb9a63062ee92.
//
// Solidity: event OpenLottery(uint32 roundId, uint32 tokenId, string btcAddr)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) FilterOpenLottery(opts *bind.FilterOpts) (*PhalaBTCLotteryOpenLotteryIterator, error) {

	logs, sub, err := _PhalaBTCLottery.contract.FilterLogs(opts, "OpenLottery")
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryOpenLotteryIterator{contract: _PhalaBTCLottery.contract, event: "OpenLottery", logs: logs, sub: sub}, nil
}

// WatchOpenLottery is a free log subscription operation binding the contract event 0x9cebd1f828553e091f5b72c5bd09ee678b2c0e326f8a3d2f93efb9a63062ee92.
//
// Solidity: event OpenLottery(uint32 roundId, uint32 tokenId, string btcAddr)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) WatchOpenLottery(opts *bind.WatchOpts, sink chan<- *PhalaBTCLotteryOpenLottery) (event.Subscription, error) {

	logs, sub, err := _PhalaBTCLottery.contract.WatchLogs(opts, "OpenLottery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PhalaBTCLotteryOpenLottery)
				if err := _PhalaBTCLottery.contract.UnpackLog(event, "OpenLottery", log); err != nil {
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

// ParseOpenLottery is a log parse operation binding the contract event 0x9cebd1f828553e091f5b72c5bd09ee678b2c0e326f8a3d2f93efb9a63062ee92.
//
// Solidity: event OpenLottery(uint32 roundId, uint32 tokenId, string btcAddr)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) ParseOpenLottery(log types.Log) (*PhalaBTCLotteryOpenLottery, error) {
	event := new(PhalaBTCLotteryOpenLottery)
	if err := _PhalaBTCLottery.contract.UnpackLog(event, "OpenLottery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PhalaBTCLotteryPayloadStoredIterator is returned from FilterPayloadStored and is used to iterate over the raw logs and unpacked data for PayloadStored events raised by the PhalaBTCLottery contract.
type PhalaBTCLotteryPayloadStoredIterator struct {
	Event *PhalaBTCLotteryPayloadStored // Event containing the contract specifics and raw log

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
func (it *PhalaBTCLotteryPayloadStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PhalaBTCLotteryPayloadStored)
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
		it.Event = new(PhalaBTCLotteryPayloadStored)
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
func (it *PhalaBTCLotteryPayloadStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PhalaBTCLotteryPayloadStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PhalaBTCLotteryPayloadStored represents a PayloadStored event raised by the PhalaBTCLottery contract.
type PhalaBTCLotteryPayloadStored struct {
	Payload []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayloadStored is a free log retrieval operation binding the contract event 0x2216a25e8e0ba35d0db376e249196115dad754519da67974ef6172c2c509381e.
//
// Solidity: event PayloadStored(bytes payload)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) FilterPayloadStored(opts *bind.FilterOpts) (*PhalaBTCLotteryPayloadStoredIterator, error) {

	logs, sub, err := _PhalaBTCLottery.contract.FilterLogs(opts, "PayloadStored")
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryPayloadStoredIterator{contract: _PhalaBTCLottery.contract, event: "PayloadStored", logs: logs, sub: sub}, nil
}

// WatchPayloadStored is a free log subscription operation binding the contract event 0x2216a25e8e0ba35d0db376e249196115dad754519da67974ef6172c2c509381e.
//
// Solidity: event PayloadStored(bytes payload)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) WatchPayloadStored(opts *bind.WatchOpts, sink chan<- *PhalaBTCLotteryPayloadStored) (event.Subscription, error) {

	logs, sub, err := _PhalaBTCLottery.contract.WatchLogs(opts, "PayloadStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PhalaBTCLotteryPayloadStored)
				if err := _PhalaBTCLottery.contract.UnpackLog(event, "PayloadStored", log); err != nil {
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

// ParsePayloadStored is a log parse operation binding the contract event 0x2216a25e8e0ba35d0db376e249196115dad754519da67974ef6172c2c509381e.
//
// Solidity: event PayloadStored(bytes payload)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) ParsePayloadStored(log types.Log) (*PhalaBTCLotteryPayloadStored, error) {
	event := new(PhalaBTCLotteryPayloadStored)
	if err := _PhalaBTCLottery.contract.UnpackLog(event, "PayloadStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
