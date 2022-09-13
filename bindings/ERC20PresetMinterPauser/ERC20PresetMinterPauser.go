// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20PresetMinterPauser

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

// ERC20PresetMinterPauserMetaData contains all meta data concerning the ERC20PresetMinterPauser contract.
var ERC20PresetMinterPauserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002baf38038062002baf833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660018202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c3578082015181840152602081019050620000a6565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011557600080fd5b838201915060208201858111156200012c57600080fd5b82518660018202830111640100000000821117156200014a57600080fd5b8083526020830192505050908051906020019080838360005b838110156200018057808201518184015260208101905062000163565b50505050905090810190601f168015620001ae5780820380516001836020036101000a031916815260200191505b5060405250505081818160049080519060200190620001cf92919062000466565b508060059080519060200190620001e892919062000466565b506012600660006101000a81548160ff021916908360ff16021790555050506000600660016101000a81548160ff021916908315150217905550620002466000801b6200023a620002d060201b60201c565b620002d860201b60201c565b620002877f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a66200027b620002d060201b60201c565b620002d860201b60201c565b620002c87f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a620002bc620002d060201b60201c565b620002d860201b60201c565b50506200050c565b600033905090565b620002ea8282620002ee60201b60201c565b5050565b6200031c816000808581526020019081526020016000206000016200039160201b6200130f1790919060201c565b156200038d5762000332620002d060201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000620003c1836000018373ffffffffffffffffffffffffffffffffffffffff1660001b620003c960201b60201c565b905092915050565b6000620003dd83836200044360201b60201c565b620004385782600001829080600181540180825580915050600190039060005260206000200160009091909190915055826000018054905083600101600084815260200190815260200160002081905550600190506200043d565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004a957805160ff1916838001178555620004da565b82800160010185558215620004da579182015b82811115620004d9578251825591602001919060010190620004bc565b5b509050620004e99190620004ed565b5090565b5b8082111562000508576000816000905550600101620004ee565b5090565b612693806200051c6000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c806370a08231116100f9578063a457c2d711610097578063d539139311610071578063d539139314610861578063d547741f1461087f578063dd62ed3e146108cd578063e63ab1e914610945576101a9565b8063a457c2d714610757578063a9059cbb146107bb578063ca15c8731461081f576101a9565b80639010d07c116100d35780639010d07c146105f057806391d148541461065257806395d89b41146106b6578063a217fddf14610739576101a9565b806370a082311461054057806379cc6790146105985780638456cb59146105e6576101a9565b8063313ce567116101665780633f4ba83a116101405780633f4ba83a1461049a57806340c10f19146104a457806342966c68146104f25780635c975abb14610520576101a9565b8063313ce567146103c757806336568abe146103e85780633950935114610436576101a9565b806306fdde03146101ae578063095ea7b31461023157806318160ddd1461029557806323b872dd146102b3578063248a9ca3146103375780632f2ff15d14610379575b600080fd5b6101b6610963565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101f65780820151818401526020810190506101db565b50505050905090810190601f1680156102235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61027d6004803603604081101561024757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a05565b60405180821515815260200191505060405180910390f35b61029d610a23565b6040518082815260200191505060405180910390f35b61031f600480360360608110156102c957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a2d565b60405180821515815260200191505060405180910390f35b6103636004803603602081101561034d57600080fd5b8101908080359060200190929190505050610b06565b6040518082815260200191505060405180910390f35b6103c56004803603604081101561038f57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b25565b005b6103cf610bae565b604051808260ff16815260200191505060405180910390f35b610434600480360360408110156103fe57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bc5565b005b6104826004803603604081101561044c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c5e565b60405180821515815260200191505060405180910390f35b6104a2610d11565b005b6104f0600480360360408110156104ba57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610da1565b005b61051e6004803603602081101561050857600080fd5b8101908080359060200190929190505050610e35565b005b610528610e49565b60405180821515815260200191505060405180910390f35b6105826004803603602081101561055657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e60565b6040518082815260200191505060405180910390f35b6105e4600480360360408110156105ae57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ea9565b005b6105ee610f0b565b005b6106266004803603604081101561060657600080fd5b810190808035906020019092919080359060200190929190505050610f9b565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61069e6004803603604081101561066857600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fcc565b60405180821515815260200191505060405180910390f35b6106be610ffd565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106fe5780820151818401526020810190506106e3565b50505050905090810190601f16801561072b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61074161109f565b6040518082815260200191505060405180910390f35b6107a36004803603604081101561076d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110a6565b60405180821515815260200191505060405180910390f35b610807600480360360408110156107d157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611173565b60405180821515815260200191505060405180910390f35b61084b6004803603602081101561083557600080fd5b8101908080359060200190929190505050611191565b6040518082815260200191505060405180910390f35b6108696111b7565b6040518082815260200191505060405180910390f35b6108cb6004803603604081101561089557600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111db565b005b61092f600480360360408110156108e357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611264565b6040518082815260200191505060405180910390f35b61094d6112eb565b6040518082815260200191505060405180910390f35b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109fb5780601f106109d0576101008083540402835291602001916109fb565b820191906000526020600020905b8154815290600101906020018083116109de57829003601f168201915b5050505050905090565b6000610a19610a1261133f565b8484611347565b6001905092915050565b6000600354905090565b6000610a3a84848461153e565b610afb84610a4661133f565b610af6856040518060600160405280602881526020016124bd60289139600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610aac61133f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118039092919063ffffffff16565b611347565b600190509392505050565b6000806000838152602001908152602001600020600201549050919050565b610b4b60008084815260200190815260200160002060020154610b4661133f565b610fcc565b610ba0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001806123bb602f913960400191505060405180910390fd5b610baa82826118bd565b5050565b6000600660009054906101000a900460ff16905090565b610bcd61133f565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c50576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180612605602f913960400191505060405180910390fd5b610c5a8282611950565b5050565b6000610d07610c6b61133f565b84610d028560026000610c7c61133f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546119e390919063ffffffff16565b611347565b6001905092915050565b610d427f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610d3d61133f565b610fcc565b610d97576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603981526020018061240c6039913960400191505060405180910390fd5b610d9f611a6b565b565b610dd27f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6610dcd61133f565b610fcc565b610e27576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806124e56036913960400191505060405180910390fd5b610e318282611b56565b5050565b610e46610e4061133f565b82611d1f565b50565b6000600660019054906101000a900460ff16905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000610ee88260405180606001604052806024815260200161251b60249139610ed986610ed461133f565b611264565b6118039092919063ffffffff16565b9050610efc83610ef661133f565b83611347565b610f068383611d1f565b505050565b610f3c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610f3761133f565b610fcc565b610f91576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260378152602001806125a96037913960400191505060405180910390fd5b610f99611ee5565b565b6000610fc482600080868152602001908152602001600020600001611fd190919063ffffffff16565b905092915050565b6000610ff582600080868152602001908152602001600020600001611feb90919063ffffffff16565b905092915050565b606060058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110955780601f1061106a57610100808354040283529160200191611095565b820191906000526020600020905b81548152906001019060200180831161107857829003601f168201915b5050505050905090565b6000801b81565b60006111696110b361133f565b84611164856040518060600160405280602581526020016125e060259139600260006110dd61133f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118039092919063ffffffff16565b611347565b6001905092915050565b600061118761118061133f565b848461153e565b6001905092915050565b60006111b060008084815260200190815260200160002060000161201b565b9050919050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b611201600080848152602001908152602001600020600201546111fc61133f565b610fcc565b611256576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603081526020018061248d6030913960400191505060405180910390fd5b6112608282611950565b5050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6000611337836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612030565b905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156113cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806125856024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611453576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124456022913960400191505060405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156115c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806125606025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561164a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806123986023913960400191505060405180910390fd5b6116558383836120a0565b6116c18160405180606001604052806026815260200161246760269139600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118039092919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061175681600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546119e390919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b60008383111582906118b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561187557808201518184015260208101905061185a565b50505050905090810190601f1680156118a25780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082840390509392505050565b6118e48160008085815260200190815260200160002060000161130f90919063ffffffff16565b1561194c576118f161133f565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b611977816000808581526020019081526020016000206000016120b090919063ffffffff16565b156119df5761198461133f565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b600080828401905083811015611a61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b611a73610e49565b611ae5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b6000600660016101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa611b2961133f565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611bf9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b611c05600083836120a0565b611c1a816003546119e390919063ffffffff16565b600381905550611c7281600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546119e390919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611da5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061253f6021913960400191505060405180910390fd5b611db1826000836120a0565b611e1d816040518060600160405280602281526020016123ea60229139600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118039092919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611e75816003546120e090919063ffffffff16565b600381905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b611eed610e49565b15611f60576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6001600660016101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611fa461133f565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b6000611fe08360000183612163565b60001c905092915050565b6000612013836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6121e6565b905092915050565b600061202982600001612209565b9050919050565b600061203c83836121e6565b61209557826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905061209a565b600090505b92915050565b6120ab83838361221a565b505050565b60006120d8836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612288565b905092915050565b600082821115612158576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b818303905092915050565b6000818360000180549050116121c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806123766022913960400191505060405180910390fd5b8260000182815481106121d357fe5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b612225838383612370565b61222d610e49565b15612283576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612634602a913960400191505060405180910390fd5b505050565b6000808360010160008481526020019081526020016000205490506000811461236457600060018203905060006001866000018054905003905060008660000182815481106122d357fe5b90600052602060002001549050808760000184815481106122f057fe5b906000526020600020018190555060018301876001016000838152602001908152602001600020819055508660000180548061232857fe5b6001900381819060005260206000200160009055905586600101600087815260200190815260200160002060009055600194505050505061236a565b60009150505b92915050565b50505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e647345524332303a207472616e7366657220746f20746865207a65726f2061646472657373416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e7445524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332305072657365744d696e7465725061757365723a206d75737420686176652070617573657220726f6c6520746f20756e706175736545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332305072657365744d696e7465725061757365723a206d7573742068617665206d696e74657220726f6c6520746f206d696e7445524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332305072657365744d696e7465725061757365723a206d75737420686176652070617573657220726f6c6520746f20706175736545524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c6645524332305061757361626c653a20746f6b656e207472616e73666572207768696c6520706175736564a2646970667358221220b33d720ce3ca6045271795ef12721e18f6b7fac6505de16837324fac90df37f864736f6c63430007000033",
}

// ERC20PresetMinterPauserABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20PresetMinterPauserMetaData.ABI instead.
var ERC20PresetMinterPauserABI = ERC20PresetMinterPauserMetaData.ABI

// ERC20PresetMinterPauserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20PresetMinterPauserMetaData.Bin instead.
var ERC20PresetMinterPauserBin = ERC20PresetMinterPauserMetaData.Bin

// DeployERC20PresetMinterPauser deploys a new Ethereum contract, binding an instance of ERC20PresetMinterPauser to it.
func DeployERC20PresetMinterPauser(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *ERC20PresetMinterPauser, error) {
	parsed, err := ERC20PresetMinterPauserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20PresetMinterPauserBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20PresetMinterPauser{ERC20PresetMinterPauserCaller: ERC20PresetMinterPauserCaller{contract: contract}, ERC20PresetMinterPauserTransactor: ERC20PresetMinterPauserTransactor{contract: contract}, ERC20PresetMinterPauserFilterer: ERC20PresetMinterPauserFilterer{contract: contract}}, nil
}

// ERC20PresetMinterPauser is an auto generated Go binding around an Ethereum contract.
type ERC20PresetMinterPauser struct {
	ERC20PresetMinterPauserCaller     // Read-only binding to the contract
	ERC20PresetMinterPauserTransactor // Write-only binding to the contract
	ERC20PresetMinterPauserFilterer   // Log filterer for contract events
}

// ERC20PresetMinterPauserCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PresetMinterPauserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PresetMinterPauserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20PresetMinterPauserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PresetMinterPauserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20PresetMinterPauserSession struct {
	Contract     *ERC20PresetMinterPauser // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20PresetMinterPauserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20PresetMinterPauserCallerSession struct {
	Contract *ERC20PresetMinterPauserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ERC20PresetMinterPauserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20PresetMinterPauserTransactorSession struct {
	Contract     *ERC20PresetMinterPauserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ERC20PresetMinterPauserRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20PresetMinterPauserRaw struct {
	Contract *ERC20PresetMinterPauser // Generic contract binding to access the raw methods on
}

// ERC20PresetMinterPauserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserCallerRaw struct {
	Contract *ERC20PresetMinterPauserCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20PresetMinterPauserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserTransactorRaw struct {
	Contract *ERC20PresetMinterPauserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20PresetMinterPauser creates a new instance of ERC20PresetMinterPauser, bound to a specific deployed contract.
func NewERC20PresetMinterPauser(address common.Address, backend bind.ContractBackend) (*ERC20PresetMinterPauser, error) {
	contract, err := bindERC20PresetMinterPauser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauser{ERC20PresetMinterPauserCaller: ERC20PresetMinterPauserCaller{contract: contract}, ERC20PresetMinterPauserTransactor: ERC20PresetMinterPauserTransactor{contract: contract}, ERC20PresetMinterPauserFilterer: ERC20PresetMinterPauserFilterer{contract: contract}}, nil
}

// NewERC20PresetMinterPauserCaller creates a new read-only instance of ERC20PresetMinterPauser, bound to a specific deployed contract.
func NewERC20PresetMinterPauserCaller(address common.Address, caller bind.ContractCaller) (*ERC20PresetMinterPauserCaller, error) {
	contract, err := bindERC20PresetMinterPauser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserCaller{contract: contract}, nil
}

// NewERC20PresetMinterPauserTransactor creates a new write-only instance of ERC20PresetMinterPauser, bound to a specific deployed contract.
func NewERC20PresetMinterPauserTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20PresetMinterPauserTransactor, error) {
	contract, err := bindERC20PresetMinterPauser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserTransactor{contract: contract}, nil
}

// NewERC20PresetMinterPauserFilterer creates a new log filterer instance of ERC20PresetMinterPauser, bound to a specific deployed contract.
func NewERC20PresetMinterPauserFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20PresetMinterPauserFilterer, error) {
	contract, err := bindERC20PresetMinterPauser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserFilterer{contract: contract}, nil
}

// bindERC20PresetMinterPauser binds a generic wrapper to an already deployed contract.
func bindERC20PresetMinterPauser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20PresetMinterPauserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PresetMinterPauser.Contract.ERC20PresetMinterPauserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.ERC20PresetMinterPauserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.ERC20PresetMinterPauserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PresetMinterPauser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauser.Contract.DEFAULTADMINROLE(&_ERC20PresetMinterPauser.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauser.Contract.DEFAULTADMINROLE(&_ERC20PresetMinterPauser.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) MINTERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauser.Contract.MINTERROLE(&_ERC20PresetMinterPauser.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) MINTERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauser.Contract.MINTERROLE(&_ERC20PresetMinterPauser.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauser.Contract.PAUSERROLE(&_ERC20PresetMinterPauser.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauser.Contract.PAUSERROLE(&_ERC20PresetMinterPauser.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauser.Contract.Allowance(&_ERC20PresetMinterPauser.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauser.Contract.Allowance(&_ERC20PresetMinterPauser.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauser.Contract.BalanceOf(&_ERC20PresetMinterPauser.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauser.Contract.BalanceOf(&_ERC20PresetMinterPauser.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Decimals() (uint8, error) {
	return _ERC20PresetMinterPauser.Contract.Decimals(&_ERC20PresetMinterPauser.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) Decimals() (uint8, error) {
	return _ERC20PresetMinterPauser.Contract.Decimals(&_ERC20PresetMinterPauser.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20PresetMinterPauser.Contract.GetRoleAdmin(&_ERC20PresetMinterPauser.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20PresetMinterPauser.Contract.GetRoleAdmin(&_ERC20PresetMinterPauser.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ERC20PresetMinterPauser.Contract.GetRoleMember(&_ERC20PresetMinterPauser.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ERC20PresetMinterPauser.Contract.GetRoleMember(&_ERC20PresetMinterPauser.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ERC20PresetMinterPauser.Contract.GetRoleMemberCount(&_ERC20PresetMinterPauser.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ERC20PresetMinterPauser.Contract.GetRoleMemberCount(&_ERC20PresetMinterPauser.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20PresetMinterPauser.Contract.HasRole(&_ERC20PresetMinterPauser.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20PresetMinterPauser.Contract.HasRole(&_ERC20PresetMinterPauser.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Name() (string, error) {
	return _ERC20PresetMinterPauser.Contract.Name(&_ERC20PresetMinterPauser.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) Name() (string, error) {
	return _ERC20PresetMinterPauser.Contract.Name(&_ERC20PresetMinterPauser.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Paused() (bool, error) {
	return _ERC20PresetMinterPauser.Contract.Paused(&_ERC20PresetMinterPauser.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) Paused() (bool, error) {
	return _ERC20PresetMinterPauser.Contract.Paused(&_ERC20PresetMinterPauser.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Symbol() (string, error) {
	return _ERC20PresetMinterPauser.Contract.Symbol(&_ERC20PresetMinterPauser.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) Symbol() (string, error) {
	return _ERC20PresetMinterPauser.Contract.Symbol(&_ERC20PresetMinterPauser.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauser.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) TotalSupply() (*big.Int, error) {
	return _ERC20PresetMinterPauser.Contract.TotalSupply(&_ERC20PresetMinterPauser.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20PresetMinterPauser.Contract.TotalSupply(&_ERC20PresetMinterPauser.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Approve(&_ERC20PresetMinterPauser.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Approve(&_ERC20PresetMinterPauser.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Burn(&_ERC20PresetMinterPauser.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Burn(&_ERC20PresetMinterPauser.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.BurnFrom(&_ERC20PresetMinterPauser.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.BurnFrom(&_ERC20PresetMinterPauser.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.DecreaseAllowance(&_ERC20PresetMinterPauser.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.DecreaseAllowance(&_ERC20PresetMinterPauser.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.GrantRole(&_ERC20PresetMinterPauser.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.GrantRole(&_ERC20PresetMinterPauser.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.IncreaseAllowance(&_ERC20PresetMinterPauser.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.IncreaseAllowance(&_ERC20PresetMinterPauser.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Mint(&_ERC20PresetMinterPauser.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Mint(&_ERC20PresetMinterPauser.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Pause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Pause(&_ERC20PresetMinterPauser.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Pause(&_ERC20PresetMinterPauser.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.RenounceRole(&_ERC20PresetMinterPauser.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.RenounceRole(&_ERC20PresetMinterPauser.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.RevokeRole(&_ERC20PresetMinterPauser.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.RevokeRole(&_ERC20PresetMinterPauser.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Transfer(&_ERC20PresetMinterPauser.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Transfer(&_ERC20PresetMinterPauser.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.TransferFrom(&_ERC20PresetMinterPauser.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.TransferFrom(&_ERC20PresetMinterPauser.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserSession) Unpause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Unpause(&_ERC20PresetMinterPauser.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauser.Contract.Unpause(&_ERC20PresetMinterPauser.TransactOpts)
}

// ERC20PresetMinterPauserApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserApprovalIterator struct {
	Event *ERC20PresetMinterPauserApproval // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserApproval)
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
		it.Event = new(ERC20PresetMinterPauserApproval)
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
func (it *ERC20PresetMinterPauserApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserApproval represents a Approval event raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20PresetMinterPauserApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20PresetMinterPauser.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserApprovalIterator{contract: _ERC20PresetMinterPauser.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20PresetMinterPauser.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserApproval)
				if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) ParseApproval(log types.Log) (*ERC20PresetMinterPauserApproval, error) {
	event := new(ERC20PresetMinterPauserApproval)
	if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserPausedIterator struct {
	Event *ERC20PresetMinterPauserPaused // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserPaused)
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
		it.Event = new(ERC20PresetMinterPauserPaused)
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
func (it *ERC20PresetMinterPauserPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserPaused represents a Paused event raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20PresetMinterPauserPausedIterator, error) {

	logs, sub, err := _ERC20PresetMinterPauser.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserPausedIterator{contract: _ERC20PresetMinterPauser.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserPaused) (event.Subscription, error) {

	logs, sub, err := _ERC20PresetMinterPauser.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserPaused)
				if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) ParsePaused(log types.Log) (*ERC20PresetMinterPauserPaused, error) {
	event := new(ERC20PresetMinterPauserPaused)
	if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserRoleAdminChangedIterator struct {
	Event *ERC20PresetMinterPauserRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserRoleAdminChanged)
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
		it.Event = new(ERC20PresetMinterPauserRoleAdminChanged)
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
func (it *ERC20PresetMinterPauserRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ERC20PresetMinterPauserRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ERC20PresetMinterPauser.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserRoleAdminChangedIterator{contract: _ERC20PresetMinterPauser.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ERC20PresetMinterPauser.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserRoleAdminChanged)
				if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) ParseRoleAdminChanged(log types.Log) (*ERC20PresetMinterPauserRoleAdminChanged, error) {
	event := new(ERC20PresetMinterPauserRoleAdminChanged)
	if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserRoleGrantedIterator struct {
	Event *ERC20PresetMinterPauserRoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserRoleGranted)
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
		it.Event = new(ERC20PresetMinterPauserRoleGranted)
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
func (it *ERC20PresetMinterPauserRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserRoleGranted represents a RoleGranted event raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20PresetMinterPauserRoleGrantedIterator, error) {

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

	logs, sub, err := _ERC20PresetMinterPauser.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserRoleGrantedIterator{contract: _ERC20PresetMinterPauser.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20PresetMinterPauser.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserRoleGranted)
				if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) ParseRoleGranted(log types.Log) (*ERC20PresetMinterPauserRoleGranted, error) {
	event := new(ERC20PresetMinterPauserRoleGranted)
	if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserRoleRevokedIterator struct {
	Event *ERC20PresetMinterPauserRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserRoleRevoked)
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
		it.Event = new(ERC20PresetMinterPauserRoleRevoked)
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
func (it *ERC20PresetMinterPauserRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserRoleRevoked represents a RoleRevoked event raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20PresetMinterPauserRoleRevokedIterator, error) {

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

	logs, sub, err := _ERC20PresetMinterPauser.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserRoleRevokedIterator{contract: _ERC20PresetMinterPauser.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20PresetMinterPauser.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserRoleRevoked)
				if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) ParseRoleRevoked(log types.Log) (*ERC20PresetMinterPauserRoleRevoked, error) {
	event := new(ERC20PresetMinterPauserRoleRevoked)
	if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserTransferIterator struct {
	Event *ERC20PresetMinterPauserTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserTransfer)
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
		it.Event = new(ERC20PresetMinterPauserTransfer)
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
func (it *ERC20PresetMinterPauserTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserTransfer represents a Transfer event raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20PresetMinterPauserTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20PresetMinterPauser.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserTransferIterator{contract: _ERC20PresetMinterPauser.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20PresetMinterPauser.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserTransfer)
				if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) ParseTransfer(log types.Log) (*ERC20PresetMinterPauserTransfer, error) {
	event := new(ERC20PresetMinterPauserTransfer)
	if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserUnpausedIterator struct {
	Event *ERC20PresetMinterPauserUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserUnpaused)
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
		it.Event = new(ERC20PresetMinterPauserUnpaused)
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
func (it *ERC20PresetMinterPauserUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserUnpaused represents a Unpaused event raised by the ERC20PresetMinterPauser contract.
type ERC20PresetMinterPauserUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20PresetMinterPauserUnpausedIterator, error) {

	logs, sub, err := _ERC20PresetMinterPauser.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserUnpausedIterator{contract: _ERC20PresetMinterPauser.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20PresetMinterPauser.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserUnpaused)
				if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PresetMinterPauser *ERC20PresetMinterPauserFilterer) ParseUnpaused(log types.Log) (*ERC20PresetMinterPauserUnpaused, error) {
	event := new(ERC20PresetMinterPauserUnpaused)
	if err := _ERC20PresetMinterPauser.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
