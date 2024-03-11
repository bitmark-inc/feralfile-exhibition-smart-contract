// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownerdata

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

// OwnerDataData is an auto generated low-level Go binding around an user-defined struct.
type OwnerDataData struct {
	Owner    common.Address
	DataHash []byte
	Metadata string
}

// OwnerDataSignature is an auto generated low-level Go binding around an user-defined struct.
type OwnerDataSignature struct {
	OwnerSign   []byte
	ExpiryBlock *big.Int
	R           [32]byte
	S           [32]byte
	V           uint8
}

// OwnerDataMetaData contains all meta data concerning the OwnerData contract.
var OwnerDataMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"serviceFeeReceiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicToken_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyServiceFeeReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerAndSenderMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerDataAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentRequiredForPublicToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotTheOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrusteeIsZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structOwnerData.Data\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"DataAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"DataRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structOwnerData.Data\",\"name\":\"data_\",\"type\":\"tuple\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structOwnerData.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"getByOwner\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structOwnerData.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes_\",\"type\":\"uint256[]\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"serviceFee_\",\"type\":\"uint256\"}],\"name\":\"setServiceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"serviceFeeReceiver_\",\"type\":\"address\"}],\"name\":\"setServiceFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structOwnerData.Data\",\"name\":\"data_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"ownerSign\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiryBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structOwnerData.Signature\",\"name\":\"signature_\",\"type\":\"tuple\"}],\"name\":\"signedAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620025963803806200259683398101604081905262000034916200013b565b6200003f33620000ce565b6001600160a01b03841662000067576040516337eefd4960e21b815260040160405180910390fd5b6001600160a01b0383166200008f5760405163270b560760e01b815260040160405180910390fd5b600180546001600160a01b039586166001600160a01b031991821617909155600280549490951693169290921790925560039190915560805262000183565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146200013657600080fd5b919050565b600080600080608085870312156200015257600080fd5b6200015d856200011e565b93506200016d602086016200011e565b6040860151606090960151949790965092505050565b6080516123e9620001ad6000396000818161015c01528181610b700152610c7501526123e96000f3fe6080604052600436106100e85760003560e01c8063715018a61161008a578063d60d907811610059578063d60d907814610284578063e8ff97f9146102a4578063f2fde38b146102c4578063fe18a4a9146102e457600080fd5b8063715018a61461021b57806372f9bf17146102305780638abdf5aa146102505780638da5cb5b1461026657600080fd5b80635cdf76f8116100c65780635cdf76f81461018c57806361937303146101ae5780636c19e783146101db5780636ca5df7f146101fb57600080fd5b8063238ac933146100ed5780632e06ca061461012a5780634ee3dc911461014a575b600080fd5b3480156100f957600080fd5b5060015461010d906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561013657600080fd5b5060025461010d906001600160a01b031681565b34801561015657600080fd5b5061017e7f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610121565b34801561019857600080fd5b506101ac6101a736600461197e565b6102f7565b005b3480156101ba57600080fd5b506101ce6101c93660046119ac565b610304565b6040516101219190611a3e565b3480156101e757600080fd5b506101ac6101f6366004611ada565b61068b565b34801561020757600080fd5b506101ce610216366004611af7565b6106dc565b34801561022757600080fd5b506101ac6109de565b34801561023c57600080fd5b506101ac61024b366004611ada565b6109f2565b34801561025c57600080fd5b5061017e60035481565b34801561027257600080fd5b506000546001600160a01b031661010d565b34801561029057600080fd5b506101ac61029f366004611b32565b610a43565b3480156102b057600080fd5b506101ac6102bf366004611bd5565b610b54565b3480156102d057600080fd5b506101ac6102df366004611ada565b610bf5565b6101ac6102f2366004611c59565b610c73565b6102ff610d16565b600355565b60606000848460405160200161031b929190611cb1565b6040516020818303038152906040529050600060048260405161033e9190611cdd565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156104d857600084815260209081902060408051606081019091526003850290910180546001600160a01b0316825260018101805492939192918401916103b590611cf9565b80601f01602080910402602001604051908101604052809291908181526020018280546103e190611cf9565b801561042e5780601f106104035761010080835404028352916020019161042e565b820191906000526020600020905b81548152906001019060200180831161041157829003601f168201915b5050505050815260200160028201805461044790611cf9565b80601f016020809104026020016040519081016040528092919081815260200182805461047390611cf9565b80156104c05780601f10610495576101008083540402835291602001916104c0565b820191906000526020600020905b8154815290600101906020018083116104a357829003601f168201915b5050505050815250508152602001906001019061036c565b505050509050600081516001600160401b038111156104f9576104f9611d2d565b60405190808252806020026020018201604052801561053257816020015b61051f611954565b8152602001906001900390816105175790505b5090506000805b83518110156105d057866001600160a01b031684828151811061055e5761055e611d43565b6020026020010151600001516001600160a01b0316036105be5783818151811061058a5761058a611d43565b60200260200101518383815181106105a4576105a4611d43565b602002602001018190525081806105ba90611d6f565b9250505b806105c881611d6f565b915050610539565b506000816001600160401b038111156105eb576105eb611d2d565b60405190808252806020026020018201604052801561062457816020015b610611611954565b8152602001906001900390816106095790505b50905060005b8281101561067c5783818151811061064457610644611d43565b602002602001015182828151811061065e5761065e611d43565b6020026020010181905250808061067490611d6f565b91505061062a565b509450505050505b9392505050565b610693610d16565b6001600160a01b0381166106ba576040516337eefd4960e21b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6060600085856040516020016106f3929190611cb1565b604051602081830303815290604052905060006004826040516107169190611cdd565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156108b057600084815260209081902060408051606081019091526003850290910180546001600160a01b03168252600181018054929391929184019161078d90611cf9565b80601f01602080910402602001604051908101604052809291908181526020018280546107b990611cf9565b80156108065780601f106107db57610100808354040283529160200191610806565b820191906000526020600020905b8154815290600101906020018083116107e957829003601f168201915b5050505050815260200160028201805461081f90611cf9565b80601f016020809104026020016040519081016040528092919081815260200182805461084b90611cf9565b80156108985780601f1061086d57610100808354040283529160200191610898565b820191906000526020600020905b81548152906001019060200180831161087b57829003601f168201915b50505050508152505081526020019060010190610744565b50505050905080518511156108fa5760408051600080825260208201909252906108f0565b6108dd611954565b8152602001906001900390816108d55790505b50925050506109d6565b80516109068587611d88565b111561091c578481516109199190611d9b565b93505b6000846001600160401b0381111561093657610936611d2d565b60405190808252806020026020018201604052801561096f57816020015b61095c611954565b8152602001906001900390816109545790505b50905060005b858110156109d057826109888289611d88565b8151811061099857610998611d43565b60200260200101518282815181106109b2576109b2611d43565b602002602001018190525080806109c890611d6f565b915050610975565b50925050505b949350505050565b6109e6610d16565b6109f06000610d70565b565b6109fa610d16565b6001600160a01b038116610a215760405163270b560760e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b610a4b610d16565b60008484604051602001610a60929190611cb1565b60405160208183030381529060405290506000600482604051610a839190611cdd565b9081526020016040518091039020905060005b83811015610b075760408051600081526020810190915282868684818110610ac057610ac0611d43565b9050602002013581548110610ad757610ad7611d43565b90600052602060002090600302016001019081610af49190611e09565b5080610aff81611d6f565b915050610a96565b5084866001600160a01b03167f391d2ecd7061332e1d3e3e01a6476694150240003fc50b5cb752f19e0adc35748686604051610b44929190611ec2565b60405180910390a3505050505050565b610b5d81610dc0565b6000610b6c6020840184611ada565b90507f0000000000000000000000000000000000000000000000000000000000000000841480610be057610bdd610ba38480611efb565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610eb592505050565b91505b610bed8287878785610f0e565b505050505050565b610bfd610d16565b6001600160a01b038116610c675760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b610c7081610d70565b50565b7f00000000000000000000000000000000000000000000000000000000000000008214808015610ca4575060035434105b15610cc25760405163aae6efb560e01b815260040160405180910390fd5b610ccf3385858585610f0e565b3415610d10576002546040516001600160a01b03909116903480156108fc02916000818181858888f19350505050158015610d0e573d6000803e3d6000fd5b505b50505050565b6000546001600160a01b031633146109f05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610c5e565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b8060200135431115610de557604051638baa579f60e01b815260040160405180910390fd5b60004630610df38480611efb565b8560200135604051602001610e0c959493929190611f6a565b6040516020818303038152906040528051906020012090506000610e7f610e60837f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b610e7060a0860160808701611fa7565b856040013586606001356110e4565b6001549091506001600160a01b03808316911614610eb057604051638baa579f60e01b815260040160405180910390fd5b505050565b6000806040518060600160405280602c8152602001612388602c9139610eda3061110c565b604051602001610eeb929190611fca565b6040516020818303038152906040529050610684610f0882611128565b84611163565b6001600160a01b038516610f256020840184611ada565b6001600160a01b031614610f4c57604051630daf43a160e41b815260040160405180910390fd5b610f596020830183611efb565b9050600003610f7b57604051630e52390960e41b815260040160405180910390fd5b60008484604051602001610f90929190611cb1565b60405160208183030381529060405290508161106857610fdf600582604051610fb99190611cdd565b9081526020016040518091039020846000016020810190610fda9190611ada565b611187565b15610ffd57604051631c872d1360e31b815260040160405180910390fd5b611014858561100f6020870187611ada565b6111a9565b611031576040516311fbe4a960e01b815260040160405180910390fd5b6110666110416020850185611ada565b6005836040516110519190611cdd565b90815260405190819003602001902090611398565b505b6004816040516110789190611cdd565b90815260405160209181900382019020805460018101825560009182529190208491600302016110a882826120cd565b505083856001600160a01b03167f7fb141aff157414e71b891d33d9759e7c1cc7678d08625f20c79bd081f9ef05885604051610b44919061221b565b60008060006110f5878787876113ad565b9150915061110281611471565b5095945050505050565b60606111226001600160a01b03831660146115bb565b92915050565b60006111348251611756565b82604051602001611146929190612290565b604051602081830303815290604052805190602001209050919050565b600080600061117285856117e8565b9150915061117f81611471565b509392505050565b6001600160a01b03811660009081526001830160205260408120541515610684565b6040516301ffc9a760e01b8152636cdb3d1360e11b60048201526000906001600160a01b038516906301ffc9a790602401602060405180830381865afa1580156111f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061121b91906122eb565b1561129b57604051627eeac760e11b81526001600160a01b038381166004830152602482018590526000919086169062fdd58e90604401602060405180830381865afa15801561126f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611293919061230d565b119050610684565b6040516301ffc9a760e01b81526380ac58cd60e01b60048201526001600160a01b038516906301ffc9a790602401602060405180830381865afa1580156112e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130a91906122eb565b1561138e576040516331a9108f60e11b8152600481018490526001600160a01b038084169190861690636352211e90602401602060405180830381865afa158015611359573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137d9190612326565b6001600160a01b0316149050610684565b5060009392505050565b6000610684836001600160a01b03841661182d565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156113e45750600090506003611468565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611438573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661146157600060019250925050611468565b9150600090505b94509492505050565b600081600481111561148557611485612343565b0361148d5750565b60018160048111156114a1576114a1612343565b036114ee5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610c5e565b600281600481111561150257611502612343565b0361154f5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610c5e565b600381600481111561156357611563612343565b03610c705760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610c5e565b606060006115ca836002612359565b6115d5906002611d88565b6001600160401b038111156115ec576115ec611d2d565b6040519080825280601f01601f191660200182016040528015611616576020820181803683370190505b509050600360fc1b8160008151811061163157611631611d43565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061166057611660611d43565b60200101906001600160f81b031916908160001a9053506000611684846002612359565b61168f906001611d88565b90505b6001811115611707576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106116c3576116c3611d43565b1a60f81b8282815181106116d9576116d9611d43565b60200101906001600160f81b031916908160001a90535060049490941c9361170081612370565b9050611692565b5083156106845760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610c5e565b606060006117638361187c565b60010190506000816001600160401b0381111561178257611782611d2d565b6040519080825280601f01601f1916602001820160405280156117ac576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a85049450846117b657509392505050565b600080825160410361181e5760208301516040840151606085015160001a611812878285856113ad565b94509450505050611826565b506000905060025b9250929050565b600081815260018301602052604081205461187457508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611122565b506000611122565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106118bb5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef810000000083106118e7576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061190557662386f26fc10000830492506010015b6305f5e100831061191d576305f5e100830492506008015b612710831061193157612710830492506004015b60648310611943576064830492506002015b600a83106111225760010192915050565b604051806060016040528060006001600160a01b0316815260200160608152602001606081525090565b60006020828403121561199057600080fd5b5035919050565b6001600160a01b0381168114610c7057600080fd5b6000806000606084860312156119c157600080fd5b83356119cc81611997565b92506020840135915060408401356119e381611997565b809150509250925092565b60005b83811015611a095781810151838201526020016119f1565b50506000910152565b60008151808452611a2a8160208601602086016119ee565b601f01601f19169290920160200192915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015611acc57888303603f19018552815180516001600160a01b0316845287810151606089860181905290611a9e82870182611a12565b91505087820151915084810388860152611ab88183611a12565b968901969450505090860190600101611a65565b509098975050505050505050565b600060208284031215611aec57600080fd5b813561068481611997565b60008060008060808587031215611b0d57600080fd5b8435611b1881611997565b966020860135965060408601359560600135945092505050565b60008060008060608587031215611b4857600080fd5b8435611b5381611997565b93506020850135925060408501356001600160401b0380821115611b7657600080fd5b818701915087601f830112611b8a57600080fd5b813581811115611b9957600080fd5b8860208260051b8501011115611bae57600080fd5b95989497505060200194505050565b600060608284031215611bcf57600080fd5b50919050565b60008060008060808587031215611beb57600080fd5b8435611bf681611997565b93506020850135925060408501356001600160401b0380821115611c1957600080fd5b611c2588838901611bbd565b93506060870135915080821115611c3b57600080fd5b50850160a08188031215611c4e57600080fd5b939692955090935050565b600080600060608486031215611c6e57600080fd5b8335611c7981611997565b92506020840135915060408401356001600160401b03811115611c9b57600080fd5b611ca786828701611bbd565b9150509250925092565b60609290921b6bffffffffffffffffffffffff19168252601f60fa1b6014830152601582015260350190565b60008251611cef8184602087016119ee565b9190910192915050565b600181811c90821680611d0d57607f821691505b602082108103611bcf57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611d8157611d81611d59565b5060010190565b8082018082111561112257611122611d59565b8181038181111561112257611122611d59565b601f821115610eb057600081815260208120601f850160051c81016020861015611dd55750805b601f850160051c820191505b81811015610bed57828155600101611de1565b600019600383901b1c191660019190911b1790565b81516001600160401b03811115611e2257611e22611d2d565b611e3681611e308454611cf9565b84611dae565b602080601f831160018114611e655760008415611e535750858301515b611e5d8582611df4565b865550610bed565b600085815260208120601f198616915b82811015611e9457888601518255948401946001909101908401611e75565b5085821015611eb25787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6020808252810182905260006001600160fb1b03831115611ee257600080fd5b8260051b80856040850137919091016040019392505050565b6000808335601e19843603018112611f1257600080fd5b8301803591506001600160401b03821115611f2c57600080fd5b60200191503681900382131561182657600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8581526001600160a01b0385166020820152608060408201819052600090611f959083018587611f41565b90508260608301529695505050505050565b600060208284031215611fb957600080fd5b813560ff8116811461068457600080fd5b60008351611fdc8184602088016119ee565b600160fd1b9083019081528351611ffa8160018401602088016119ee565b601760f91b60019290910191820152600201949350505050565b6001600160401b0383111561202b5761202b611d2d565b61203f836120398354611cf9565b83611dae565b6000601f84116001811461206d576000851561205b5750838201355b6120658682611df4565b845550610d0e565b600083815260209020601f19861690835b8281101561209e578685013582556020948501946001909201910161207e565b50868210156120bb5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81356120d881611997565b81546001600160a01b0319166001600160a01b03919091161781556001818101602061210685820186611efb565b6001600160401b0381111561211d5761211d611d2d565b6121318161212b8654611cf9565b86611dae565b6000601f82116001811461215f576000831561214d5750838201355b6121578482611df4565b8755506121b4565b600086815260209020601f19841690835b8281101561218d5786850135825593870193908901908701612170565b50848210156121aa5760001960f88660031b161c19848701351681555b50508683881b0186555b505050505050506121c86040830183611efb565b610d10818360028601612014565b6000808335601e198436030181126121ed57600080fd5b83016020810192503590506001600160401b0381111561220c57600080fd5b80360382131561182657600080fd5b602081526000823561222c81611997565b6001600160a01b031660208381019190915261224a908401846121d6565b6060604085015261225f608085018284611f41565b91505061226f60408501856121d6565b848303601f19016060860152612286838284611f41565b9695505050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a0000000000008152600083516122c881601a8501602088016119ee565b8351908301906122df81601a8401602088016119ee565b01601a01949350505050565b6000602082840312156122fd57600080fd5b8151801515811461068457600080fd5b60006020828403121561231f57600080fd5b5051919050565b60006020828403121561233857600080fd5b815161068481611997565b634e487b7160e01b600052602160045260246000fd5b808202811582820484141761112257611122611d59565b60008161237f5761237f611d59565b50600019019056fe417574686f72697a6520746f20777269746520796f7572206461746120746f2074686520636f6e7472616374a2646970667358221220460b884331a11357ee904024972be1a3bdaa9953c6dd8fc7eb5b20a455acec0064736f6c63430008110033",
}

// OwnerDataABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnerDataMetaData.ABI instead.
var OwnerDataABI = OwnerDataMetaData.ABI

// OwnerDataBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OwnerDataMetaData.Bin instead.
var OwnerDataBin = OwnerDataMetaData.Bin

// DeployOwnerData deploys a new Ethereum contract, binding an instance of OwnerData to it.
func DeployOwnerData(auth *bind.TransactOpts, backend bind.ContractBackend, signer_ common.Address, serviceFeeReceiver_ common.Address, serviceFee_ *big.Int, publicToken_ *big.Int) (common.Address, *types.Transaction, *OwnerData, error) {
	parsed, err := OwnerDataMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnerDataBin), backend, signer_, serviceFeeReceiver_, serviceFee_, publicToken_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OwnerData{OwnerDataCaller: OwnerDataCaller{contract: contract}, OwnerDataTransactor: OwnerDataTransactor{contract: contract}, OwnerDataFilterer: OwnerDataFilterer{contract: contract}}, nil
}

// OwnerData is an auto generated Go binding around an Ethereum contract.
type OwnerData struct {
	OwnerDataCaller     // Read-only binding to the contract
	OwnerDataTransactor // Write-only binding to the contract
	OwnerDataFilterer   // Log filterer for contract events
}

// OwnerDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnerDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnerDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnerDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnerDataSession struct {
	Contract     *OwnerData        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnerDataCallerSession struct {
	Contract *OwnerDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OwnerDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnerDataTransactorSession struct {
	Contract     *OwnerDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OwnerDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnerDataRaw struct {
	Contract *OwnerData // Generic contract binding to access the raw methods on
}

// OwnerDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnerDataCallerRaw struct {
	Contract *OwnerDataCaller // Generic read-only contract binding to access the raw methods on
}

// OwnerDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnerDataTransactorRaw struct {
	Contract *OwnerDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnerData creates a new instance of OwnerData, bound to a specific deployed contract.
func NewOwnerData(address common.Address, backend bind.ContractBackend) (*OwnerData, error) {
	contract, err := bindOwnerData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnerData{OwnerDataCaller: OwnerDataCaller{contract: contract}, OwnerDataTransactor: OwnerDataTransactor{contract: contract}, OwnerDataFilterer: OwnerDataFilterer{contract: contract}}, nil
}

// NewOwnerDataCaller creates a new read-only instance of OwnerData, bound to a specific deployed contract.
func NewOwnerDataCaller(address common.Address, caller bind.ContractCaller) (*OwnerDataCaller, error) {
	contract, err := bindOwnerData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerDataCaller{contract: contract}, nil
}

// NewOwnerDataTransactor creates a new write-only instance of OwnerData, bound to a specific deployed contract.
func NewOwnerDataTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerDataTransactor, error) {
	contract, err := bindOwnerData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerDataTransactor{contract: contract}, nil
}

// NewOwnerDataFilterer creates a new log filterer instance of OwnerData, bound to a specific deployed contract.
func NewOwnerDataFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerDataFilterer, error) {
	contract, err := bindOwnerData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerDataFilterer{contract: contract}, nil
}

// bindOwnerData binds a generic wrapper to an already deployed contract.
func bindOwnerData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnerDataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnerData *OwnerDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerData.Contract.OwnerDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnerData *OwnerDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerData.Contract.OwnerDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnerData *OwnerDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerData.Contract.OwnerDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnerData *OwnerDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnerData *OwnerDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnerData *OwnerDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerData.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6ca5df7f.
//
// Solidity: function get(address contractAddress_, uint256 tokenID_, uint256 startIndex, uint256 count) view returns((address,bytes,string)[])
func (_OwnerData *OwnerDataCaller) Get(opts *bind.CallOpts, contractAddress_ common.Address, tokenID_ *big.Int, startIndex *big.Int, count *big.Int) ([]OwnerDataData, error) {
	var out []interface{}
	err := _OwnerData.contract.Call(opts, &out, "get", contractAddress_, tokenID_, startIndex, count)

	if err != nil {
		return *new([]OwnerDataData), err
	}

	out0 := *abi.ConvertType(out[0], new([]OwnerDataData)).(*[]OwnerDataData)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6ca5df7f.
//
// Solidity: function get(address contractAddress_, uint256 tokenID_, uint256 startIndex, uint256 count) view returns((address,bytes,string)[])
func (_OwnerData *OwnerDataSession) Get(contractAddress_ common.Address, tokenID_ *big.Int, startIndex *big.Int, count *big.Int) ([]OwnerDataData, error) {
	return _OwnerData.Contract.Get(&_OwnerData.CallOpts, contractAddress_, tokenID_, startIndex, count)
}

// Get is a free data retrieval call binding the contract method 0x6ca5df7f.
//
// Solidity: function get(address contractAddress_, uint256 tokenID_, uint256 startIndex, uint256 count) view returns((address,bytes,string)[])
func (_OwnerData *OwnerDataCallerSession) Get(contractAddress_ common.Address, tokenID_ *big.Int, startIndex *big.Int, count *big.Int) ([]OwnerDataData, error) {
	return _OwnerData.Contract.Get(&_OwnerData.CallOpts, contractAddress_, tokenID_, startIndex, count)
}

// GetByOwner is a free data retrieval call binding the contract method 0x61937303.
//
// Solidity: function getByOwner(address contractAddress_, uint256 tokenID_, address owner_) view returns((address,bytes,string)[])
func (_OwnerData *OwnerDataCaller) GetByOwner(opts *bind.CallOpts, contractAddress_ common.Address, tokenID_ *big.Int, owner_ common.Address) ([]OwnerDataData, error) {
	var out []interface{}
	err := _OwnerData.contract.Call(opts, &out, "getByOwner", contractAddress_, tokenID_, owner_)

	if err != nil {
		return *new([]OwnerDataData), err
	}

	out0 := *abi.ConvertType(out[0], new([]OwnerDataData)).(*[]OwnerDataData)

	return out0, err

}

// GetByOwner is a free data retrieval call binding the contract method 0x61937303.
//
// Solidity: function getByOwner(address contractAddress_, uint256 tokenID_, address owner_) view returns((address,bytes,string)[])
func (_OwnerData *OwnerDataSession) GetByOwner(contractAddress_ common.Address, tokenID_ *big.Int, owner_ common.Address) ([]OwnerDataData, error) {
	return _OwnerData.Contract.GetByOwner(&_OwnerData.CallOpts, contractAddress_, tokenID_, owner_)
}

// GetByOwner is a free data retrieval call binding the contract method 0x61937303.
//
// Solidity: function getByOwner(address contractAddress_, uint256 tokenID_, address owner_) view returns((address,bytes,string)[])
func (_OwnerData *OwnerDataCallerSession) GetByOwner(contractAddress_ common.Address, tokenID_ *big.Int, owner_ common.Address) ([]OwnerDataData, error) {
	return _OwnerData.Contract.GetByOwner(&_OwnerData.CallOpts, contractAddress_, tokenID_, owner_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnerData *OwnerDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnerData.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnerData *OwnerDataSession) Owner() (common.Address, error) {
	return _OwnerData.Contract.Owner(&_OwnerData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnerData *OwnerDataCallerSession) Owner() (common.Address, error) {
	return _OwnerData.Contract.Owner(&_OwnerData.CallOpts)
}

// PublicToken is a free data retrieval call binding the contract method 0x4ee3dc91.
//
// Solidity: function publicToken() view returns(uint256)
func (_OwnerData *OwnerDataCaller) PublicToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OwnerData.contract.Call(opts, &out, "publicToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicToken is a free data retrieval call binding the contract method 0x4ee3dc91.
//
// Solidity: function publicToken() view returns(uint256)
func (_OwnerData *OwnerDataSession) PublicToken() (*big.Int, error) {
	return _OwnerData.Contract.PublicToken(&_OwnerData.CallOpts)
}

// PublicToken is a free data retrieval call binding the contract method 0x4ee3dc91.
//
// Solidity: function publicToken() view returns(uint256)
func (_OwnerData *OwnerDataCallerSession) PublicToken() (*big.Int, error) {
	return _OwnerData.Contract.PublicToken(&_OwnerData.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint256)
func (_OwnerData *OwnerDataCaller) ServiceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OwnerData.contract.Call(opts, &out, "serviceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint256)
func (_OwnerData *OwnerDataSession) ServiceFee() (*big.Int, error) {
	return _OwnerData.Contract.ServiceFee(&_OwnerData.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint256)
func (_OwnerData *OwnerDataCallerSession) ServiceFee() (*big.Int, error) {
	return _OwnerData.Contract.ServiceFee(&_OwnerData.CallOpts)
}

// ServiceFeeReceiver is a free data retrieval call binding the contract method 0x2e06ca06.
//
// Solidity: function serviceFeeReceiver() view returns(address)
func (_OwnerData *OwnerDataCaller) ServiceFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnerData.contract.Call(opts, &out, "serviceFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceFeeReceiver is a free data retrieval call binding the contract method 0x2e06ca06.
//
// Solidity: function serviceFeeReceiver() view returns(address)
func (_OwnerData *OwnerDataSession) ServiceFeeReceiver() (common.Address, error) {
	return _OwnerData.Contract.ServiceFeeReceiver(&_OwnerData.CallOpts)
}

// ServiceFeeReceiver is a free data retrieval call binding the contract method 0x2e06ca06.
//
// Solidity: function serviceFeeReceiver() view returns(address)
func (_OwnerData *OwnerDataCallerSession) ServiceFeeReceiver() (common.Address, error) {
	return _OwnerData.Contract.ServiceFeeReceiver(&_OwnerData.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_OwnerData *OwnerDataCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnerData.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_OwnerData *OwnerDataSession) Signer() (common.Address, error) {
	return _OwnerData.Contract.Signer(&_OwnerData.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_OwnerData *OwnerDataCallerSession) Signer() (common.Address, error) {
	return _OwnerData.Contract.Signer(&_OwnerData.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xfe18a4a9.
//
// Solidity: function add(address contractAddress_, uint256 tokenID_, (address,bytes,string) data_) payable returns()
func (_OwnerData *OwnerDataTransactor) Add(opts *bind.TransactOpts, contractAddress_ common.Address, tokenID_ *big.Int, data_ OwnerDataData) (*types.Transaction, error) {
	return _OwnerData.contract.Transact(opts, "add", contractAddress_, tokenID_, data_)
}

// Add is a paid mutator transaction binding the contract method 0xfe18a4a9.
//
// Solidity: function add(address contractAddress_, uint256 tokenID_, (address,bytes,string) data_) payable returns()
func (_OwnerData *OwnerDataSession) Add(contractAddress_ common.Address, tokenID_ *big.Int, data_ OwnerDataData) (*types.Transaction, error) {
	return _OwnerData.Contract.Add(&_OwnerData.TransactOpts, contractAddress_, tokenID_, data_)
}

// Add is a paid mutator transaction binding the contract method 0xfe18a4a9.
//
// Solidity: function add(address contractAddress_, uint256 tokenID_, (address,bytes,string) data_) payable returns()
func (_OwnerData *OwnerDataTransactorSession) Add(contractAddress_ common.Address, tokenID_ *big.Int, data_ OwnerDataData) (*types.Transaction, error) {
	return _OwnerData.Contract.Add(&_OwnerData.TransactOpts, contractAddress_, tokenID_, data_)
}

// Remove is a paid mutator transaction binding the contract method 0xd60d9078.
//
// Solidity: function remove(address contractAddress_, uint256 tokenID_, uint256[] indexes_) returns()
func (_OwnerData *OwnerDataTransactor) Remove(opts *bind.TransactOpts, contractAddress_ common.Address, tokenID_ *big.Int, indexes_ []*big.Int) (*types.Transaction, error) {
	return _OwnerData.contract.Transact(opts, "remove", contractAddress_, tokenID_, indexes_)
}

// Remove is a paid mutator transaction binding the contract method 0xd60d9078.
//
// Solidity: function remove(address contractAddress_, uint256 tokenID_, uint256[] indexes_) returns()
func (_OwnerData *OwnerDataSession) Remove(contractAddress_ common.Address, tokenID_ *big.Int, indexes_ []*big.Int) (*types.Transaction, error) {
	return _OwnerData.Contract.Remove(&_OwnerData.TransactOpts, contractAddress_, tokenID_, indexes_)
}

// Remove is a paid mutator transaction binding the contract method 0xd60d9078.
//
// Solidity: function remove(address contractAddress_, uint256 tokenID_, uint256[] indexes_) returns()
func (_OwnerData *OwnerDataTransactorSession) Remove(contractAddress_ common.Address, tokenID_ *big.Int, indexes_ []*big.Int) (*types.Transaction, error) {
	return _OwnerData.Contract.Remove(&_OwnerData.TransactOpts, contractAddress_, tokenID_, indexes_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnerData *OwnerDataTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerData.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnerData *OwnerDataSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnerData.Contract.RenounceOwnership(&_OwnerData.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnerData *OwnerDataTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnerData.Contract.RenounceOwnership(&_OwnerData.TransactOpts)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x5cdf76f8.
//
// Solidity: function setServiceFee(uint256 serviceFee_) returns()
func (_OwnerData *OwnerDataTransactor) SetServiceFee(opts *bind.TransactOpts, serviceFee_ *big.Int) (*types.Transaction, error) {
	return _OwnerData.contract.Transact(opts, "setServiceFee", serviceFee_)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x5cdf76f8.
//
// Solidity: function setServiceFee(uint256 serviceFee_) returns()
func (_OwnerData *OwnerDataSession) SetServiceFee(serviceFee_ *big.Int) (*types.Transaction, error) {
	return _OwnerData.Contract.SetServiceFee(&_OwnerData.TransactOpts, serviceFee_)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x5cdf76f8.
//
// Solidity: function setServiceFee(uint256 serviceFee_) returns()
func (_OwnerData *OwnerDataTransactorSession) SetServiceFee(serviceFee_ *big.Int) (*types.Transaction, error) {
	return _OwnerData.Contract.SetServiceFee(&_OwnerData.TransactOpts, serviceFee_)
}

// SetServiceFeeReceiver is a paid mutator transaction binding the contract method 0x72f9bf17.
//
// Solidity: function setServiceFeeReceiver(address serviceFeeReceiver_) returns()
func (_OwnerData *OwnerDataTransactor) SetServiceFeeReceiver(opts *bind.TransactOpts, serviceFeeReceiver_ common.Address) (*types.Transaction, error) {
	return _OwnerData.contract.Transact(opts, "setServiceFeeReceiver", serviceFeeReceiver_)
}

// SetServiceFeeReceiver is a paid mutator transaction binding the contract method 0x72f9bf17.
//
// Solidity: function setServiceFeeReceiver(address serviceFeeReceiver_) returns()
func (_OwnerData *OwnerDataSession) SetServiceFeeReceiver(serviceFeeReceiver_ common.Address) (*types.Transaction, error) {
	return _OwnerData.Contract.SetServiceFeeReceiver(&_OwnerData.TransactOpts, serviceFeeReceiver_)
}

// SetServiceFeeReceiver is a paid mutator transaction binding the contract method 0x72f9bf17.
//
// Solidity: function setServiceFeeReceiver(address serviceFeeReceiver_) returns()
func (_OwnerData *OwnerDataTransactorSession) SetServiceFeeReceiver(serviceFeeReceiver_ common.Address) (*types.Transaction, error) {
	return _OwnerData.Contract.SetServiceFeeReceiver(&_OwnerData.TransactOpts, serviceFeeReceiver_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_OwnerData *OwnerDataTransactor) SetSigner(opts *bind.TransactOpts, signer_ common.Address) (*types.Transaction, error) {
	return _OwnerData.contract.Transact(opts, "setSigner", signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_OwnerData *OwnerDataSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _OwnerData.Contract.SetSigner(&_OwnerData.TransactOpts, signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_OwnerData *OwnerDataTransactorSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _OwnerData.Contract.SetSigner(&_OwnerData.TransactOpts, signer_)
}

// SignedAdd is a paid mutator transaction binding the contract method 0xe8ff97f9.
//
// Solidity: function signedAdd(address contractAddress_, uint256 tokenID_, (address,bytes,string) data_, (bytes,uint256,bytes32,bytes32,uint8) signature_) returns()
func (_OwnerData *OwnerDataTransactor) SignedAdd(opts *bind.TransactOpts, contractAddress_ common.Address, tokenID_ *big.Int, data_ OwnerDataData, signature_ OwnerDataSignature) (*types.Transaction, error) {
	return _OwnerData.contract.Transact(opts, "signedAdd", contractAddress_, tokenID_, data_, signature_)
}

// SignedAdd is a paid mutator transaction binding the contract method 0xe8ff97f9.
//
// Solidity: function signedAdd(address contractAddress_, uint256 tokenID_, (address,bytes,string) data_, (bytes,uint256,bytes32,bytes32,uint8) signature_) returns()
func (_OwnerData *OwnerDataSession) SignedAdd(contractAddress_ common.Address, tokenID_ *big.Int, data_ OwnerDataData, signature_ OwnerDataSignature) (*types.Transaction, error) {
	return _OwnerData.Contract.SignedAdd(&_OwnerData.TransactOpts, contractAddress_, tokenID_, data_, signature_)
}

// SignedAdd is a paid mutator transaction binding the contract method 0xe8ff97f9.
//
// Solidity: function signedAdd(address contractAddress_, uint256 tokenID_, (address,bytes,string) data_, (bytes,uint256,bytes32,bytes32,uint8) signature_) returns()
func (_OwnerData *OwnerDataTransactorSession) SignedAdd(contractAddress_ common.Address, tokenID_ *big.Int, data_ OwnerDataData, signature_ OwnerDataSignature) (*types.Transaction, error) {
	return _OwnerData.Contract.SignedAdd(&_OwnerData.TransactOpts, contractAddress_, tokenID_, data_, signature_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnerData *OwnerDataTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnerData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnerData *OwnerDataSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnerData.Contract.TransferOwnership(&_OwnerData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnerData *OwnerDataTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnerData.Contract.TransferOwnership(&_OwnerData.TransactOpts, newOwner)
}

// OwnerDataDataAddedIterator is returned from FilterDataAdded and is used to iterate over the raw logs and unpacked data for DataAdded events raised by the OwnerData contract.
type OwnerDataDataAddedIterator struct {
	Event *OwnerDataDataAdded // Event containing the contract specifics and raw log

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
func (it *OwnerDataDataAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerDataDataAdded)
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
		it.Event = new(OwnerDataDataAdded)
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
func (it *OwnerDataDataAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerDataDataAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerDataDataAdded represents a DataAdded event raised by the OwnerData contract.
type OwnerDataDataAdded struct {
	ContractAddress common.Address
	TokenID         *big.Int
	Data            OwnerDataData
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDataAdded is a free log retrieval operation binding the contract event 0x7fb141aff157414e71b891d33d9759e7c1cc7678d08625f20c79bd081f9ef058.
//
// Solidity: event DataAdded(address indexed contractAddress, uint256 indexed tokenID, (address,bytes,string) data)
func (_OwnerData *OwnerDataFilterer) FilterDataAdded(opts *bind.FilterOpts, contractAddress []common.Address, tokenID []*big.Int) (*OwnerDataDataAddedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _OwnerData.contract.FilterLogs(opts, "DataAdded", contractAddressRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &OwnerDataDataAddedIterator{contract: _OwnerData.contract, event: "DataAdded", logs: logs, sub: sub}, nil
}

// WatchDataAdded is a free log subscription operation binding the contract event 0x7fb141aff157414e71b891d33d9759e7c1cc7678d08625f20c79bd081f9ef058.
//
// Solidity: event DataAdded(address indexed contractAddress, uint256 indexed tokenID, (address,bytes,string) data)
func (_OwnerData *OwnerDataFilterer) WatchDataAdded(opts *bind.WatchOpts, sink chan<- *OwnerDataDataAdded, contractAddress []common.Address, tokenID []*big.Int) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _OwnerData.contract.WatchLogs(opts, "DataAdded", contractAddressRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerDataDataAdded)
				if err := _OwnerData.contract.UnpackLog(event, "DataAdded", log); err != nil {
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

// ParseDataAdded is a log parse operation binding the contract event 0x7fb141aff157414e71b891d33d9759e7c1cc7678d08625f20c79bd081f9ef058.
//
// Solidity: event DataAdded(address indexed contractAddress, uint256 indexed tokenID, (address,bytes,string) data)
func (_OwnerData *OwnerDataFilterer) ParseDataAdded(log types.Log) (*OwnerDataDataAdded, error) {
	event := new(OwnerDataDataAdded)
	if err := _OwnerData.contract.UnpackLog(event, "DataAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerDataDataRemovedIterator is returned from FilterDataRemoved and is used to iterate over the raw logs and unpacked data for DataRemoved events raised by the OwnerData contract.
type OwnerDataDataRemovedIterator struct {
	Event *OwnerDataDataRemoved // Event containing the contract specifics and raw log

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
func (it *OwnerDataDataRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerDataDataRemoved)
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
		it.Event = new(OwnerDataDataRemoved)
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
func (it *OwnerDataDataRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerDataDataRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerDataDataRemoved represents a DataRemoved event raised by the OwnerData contract.
type OwnerDataDataRemoved struct {
	ContractAddress common.Address
	TokenID         *big.Int
	Indexes         []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDataRemoved is a free log retrieval operation binding the contract event 0x391d2ecd7061332e1d3e3e01a6476694150240003fc50b5cb752f19e0adc3574.
//
// Solidity: event DataRemoved(address indexed contractAddress, uint256 indexed tokenID, uint256[] indexes)
func (_OwnerData *OwnerDataFilterer) FilterDataRemoved(opts *bind.FilterOpts, contractAddress []common.Address, tokenID []*big.Int) (*OwnerDataDataRemovedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _OwnerData.contract.FilterLogs(opts, "DataRemoved", contractAddressRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &OwnerDataDataRemovedIterator{contract: _OwnerData.contract, event: "DataRemoved", logs: logs, sub: sub}, nil
}

// WatchDataRemoved is a free log subscription operation binding the contract event 0x391d2ecd7061332e1d3e3e01a6476694150240003fc50b5cb752f19e0adc3574.
//
// Solidity: event DataRemoved(address indexed contractAddress, uint256 indexed tokenID, uint256[] indexes)
func (_OwnerData *OwnerDataFilterer) WatchDataRemoved(opts *bind.WatchOpts, sink chan<- *OwnerDataDataRemoved, contractAddress []common.Address, tokenID []*big.Int) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _OwnerData.contract.WatchLogs(opts, "DataRemoved", contractAddressRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerDataDataRemoved)
				if err := _OwnerData.contract.UnpackLog(event, "DataRemoved", log); err != nil {
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

// ParseDataRemoved is a log parse operation binding the contract event 0x391d2ecd7061332e1d3e3e01a6476694150240003fc50b5cb752f19e0adc3574.
//
// Solidity: event DataRemoved(address indexed contractAddress, uint256 indexed tokenID, uint256[] indexes)
func (_OwnerData *OwnerDataFilterer) ParseDataRemoved(log types.Log) (*OwnerDataDataRemoved, error) {
	event := new(OwnerDataDataRemoved)
	if err := _OwnerData.contract.UnpackLog(event, "DataRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerDataOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnerData contract.
type OwnerDataOwnershipTransferredIterator struct {
	Event *OwnerDataOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnerDataOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerDataOwnershipTransferred)
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
		it.Event = new(OwnerDataOwnershipTransferred)
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
func (it *OwnerDataOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerDataOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerDataOwnershipTransferred represents a OwnershipTransferred event raised by the OwnerData contract.
type OwnerDataOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnerData *OwnerDataFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnerDataOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnerData.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnerDataOwnershipTransferredIterator{contract: _OwnerData.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnerData *OwnerDataFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnerDataOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnerData.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerDataOwnershipTransferred)
				if err := _OwnerData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnerData *OwnerDataFilterer) ParseOwnershipTransferred(log types.Log) (*OwnerDataOwnershipTransferred, error) {
	event := new(OwnerDataOwnershipTransferred)
	if err := _OwnerData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
