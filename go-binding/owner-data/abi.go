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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"serviceFeeReceiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicToken_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyServiceFeeReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerAndSenderMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerDataAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentRequiredForPublicToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotTheOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrusteeIsZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structOwnerData.Data\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"DataAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"DataRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structOwnerData.Data\",\"name\":\"data_\",\"type\":\"tuple\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structOwnerData.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"getByOwner\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structOwnerData.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes_\",\"type\":\"uint256[]\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"serviceFee_\",\"type\":\"uint256\"}],\"name\":\"setServiceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"publicToken_\",\"type\":\"uint256\"}],\"name\":\"setPublicToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"serviceFeeReceiver_\",\"type\":\"address\"}],\"name\":\"setServiceFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structOwnerData.Data\",\"name\":\"data_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"ownerSign\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiryBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structOwnerData.Signature\",\"name\":\"signature_\",\"type\":\"tuple\"}],\"name\":\"signedAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620025a9380380620025a983398101604081905262000034916200013b565b6200003f33620000ce565b6001600160a01b03841662000067576040516337eefd4960e21b815260040160405180910390fd5b6001600160a01b0383166200008f5760405163270b560760e01b815260040160405180910390fd5b600280546001600160a01b039586166001600160a01b031991821617909155600380549490951693169290921790925560049190915560015562000183565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146200013657600080fd5b919050565b600080600080608085870312156200015257600080fd5b6200015d856200011e565b93506200016d602086016200011e565b6040860151606090960151949790965092505050565b61241680620001936000396000f3fe6080604052600436106100f35760003560e01c8063715018a61161008a578063d60d907811610059578063d60d907814610291578063e8ff97f9146102b1578063f2fde38b146102d1578063fe18a4a9146102f157600080fd5b8063715018a61461022857806372f9bf171461023d5780638abdf5aa1461025d5780638da5cb5b1461027357600080fd5b80635cdf76f8116100c65780635cdf76f81461019b57806361937303146101bb5780636c19e783146101e85780636ca5df7f1461020857600080fd5b8063238ac933146100f85780632990b640146101355780632e06ca06146101575780634ee3dc9114610177575b600080fd5b34801561010457600080fd5b50600254610118906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561014157600080fd5b506101556101503660046119ab565b610304565b005b34801561016357600080fd5b50600354610118906001600160a01b031681565b34801561018357600080fd5b5061018d60015481565b60405190815260200161012c565b3480156101a757600080fd5b506101556101b63660046119ab565b610311565b3480156101c757600080fd5b506101db6101d63660046119d9565b61031e565b60405161012c9190611a6b565b3480156101f457600080fd5b50610155610203366004611b07565b6106ba565b34801561021457600080fd5b506101db610223366004611b24565b61070b565b34801561023457600080fd5b50610155610a24565b34801561024957600080fd5b50610155610258366004611b07565b610a38565b34801561026957600080fd5b5061018d60045481565b34801561027f57600080fd5b506000546001600160a01b0316610118565b34801561029d57600080fd5b506101556102ac366004611b5f565b610a89565b3480156102bd57600080fd5b506101556102cc366004611c02565b610bbc565b3480156102dd57600080fd5b506101556102ec366004611b07565b610c40565b6101556102ff366004611c86565b610cbe565b61030c610d43565b600155565b610319610d43565b600455565b606060008484604051602001610335929190611cde565b604051602081830303815290604052905060006005826040516103589190611d0a565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156104f257600084815260209081902060408051606081019091526003850290910180546001600160a01b0316825260018101805492939192918401916103cf90611d26565b80601f01602080910402602001604051908101604052809291908181526020018280546103fb90611d26565b80156104485780601f1061041d57610100808354040283529160200191610448565b820191906000526020600020905b81548152906001019060200180831161042b57829003601f168201915b5050505050815260200160028201805461046190611d26565b80601f016020809104026020016040519081016040528092919081815260200182805461048d90611d26565b80156104da5780601f106104af576101008083540402835291602001916104da565b820191906000526020600020905b8154815290600101906020018083116104bd57829003601f168201915b50505050508152505081526020019060010190610386565b505050509050600081516001600160401b0381111561051357610513611d5a565b60405190808252806020026020018201604052801561054c57816020015b610539611981565b8152602001906001900390816105315790505b5090506000805b83518110156105ea57866001600160a01b031684828151811061057857610578611d70565b6020026020010151600001516001600160a01b0316036105d8578381815181106105a4576105a4611d70565b60200260200101518383815181106105be576105be611d70565b602002602001018190525081806105d490611d9c565b9250505b806105e281611d9c565b915050610553565b506000816001600160401b0381111561060557610605611d5a565b60405190808252806020026020018201604052801561063e57816020015b61062b611981565b8152602001906001900390816106235790505b50905060005b828110156106ab578381610659600186611db5565b6106639190611db5565b8151811061067357610673611d70565b602002602001015182828151811061068d5761068d611d70565b602002602001018190525080806106a390611d9c565b915050610644565b509450505050505b9392505050565b6106c2610d43565b6001600160a01b0381166106e9576040516337eefd4960e21b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b606060008585604051602001610722929190611cde565b604051602081830303815290604052905060006005826040516107459190611d0a565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156108df57600084815260209081902060408051606081019091526003850290910180546001600160a01b0316825260018101805492939192918401916107bc90611d26565b80601f01602080910402602001604051908101604052809291908181526020018280546107e890611d26565b80156108355780601f1061080a57610100808354040283529160200191610835565b820191906000526020600020905b81548152906001019060200180831161081857829003601f168201915b5050505050815260200160028201805461084e90611d26565b80601f016020809104026020016040519081016040528092919081815260200182805461087a90611d26565b80156108c75780601f1061089c576101008083540402835291602001916108c7565b820191906000526020600020905b8154815290600101906020018083116108aa57829003601f168201915b50505050508152505081526020019060010190610773565b5050825192935050508086111561092c576040805160008082526020820190925290610921565b61090e611981565b8152602001906001900390816109065790505b509350505050610a1c565b806109378688611dc8565b111561094a576109478682611db5565b94505b6000856001600160401b0381111561096457610964611d5a565b60405190808252806020026020018201604052801561099d57816020015b61098a611981565b8152602001906001900390816109825790505b50905060005b86811015610a15578388826109b9600187611db5565b6109c39190611db5565b6109cd9190611db5565b815181106109dd576109dd611d70565b60200260200101518282815181106109f7576109f7611d70565b60200260200101819052508080610a0d90611d9c565b9150506109a3565b5093505050505b949350505050565b610a2c610d43565b610a366000610d9d565b565b610a40610d43565b6001600160a01b038116610a675760405163270b560760e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b610a91610d43565b8260015414610ab357604051630e52390960e41b815260040160405180910390fd5b60008484604051602001610ac8929190611cde565b60405160208183030381529060405290506000600582604051610aeb9190611d0a565b9081526020016040518091039020905060005b83811015610b6f5760408051600081526020810190915282868684818110610b2857610b28611d70565b9050602002013581548110610b3f57610b3f611d70565b90600052602060002090600302016001019081610b5c9190611e36565b5080610b6781611d9c565b915050610afe565b5084866001600160a01b03167f391d2ecd7061332e1d3e3e01a6476694150240003fc50b5cb752f19e0adc35748686604051610bac929190611eef565b60405180910390a3505050505050565b610bc581610ded565b6000610bd46020840184611b07565b600154909150841480610c2b57610c28610bee8480611f28565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ee292505050565b91505b610c388287878785610f3b565b505050505050565b610c48610d43565b6001600160a01b038116610cb25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b610cbb81610d9d565b50565b6001548214808015610cd1575060045434105b15610cef5760405163aae6efb560e01b815260040160405180910390fd5b610cfc3385858585610f3b565b3415610d3d576003546040516001600160a01b03909116903480156108fc02916000818181858888f19350505050158015610d3b573d6000803e3d6000fd5b505b50505050565b6000546001600160a01b03163314610a365760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610ca9565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b8060200135431115610e1257604051638baa579f60e01b815260040160405180910390fd5b60004630610e208480611f28565b8560200135604051602001610e39959493929190611f97565b6040516020818303038152906040528051906020012090506000610eac610e8d837f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b610e9d60a0860160808701611fd4565b85604001358660600135611111565b6002549091506001600160a01b03808316911614610edd57604051638baa579f60e01b815260040160405180910390fd5b505050565b6000806040518060600160405280602c81526020016123b5602c9139610f0730611139565b604051602001610f18929190611ff7565b60405160208183030381529060405290506106b3610f3582611155565b84611190565b6001600160a01b038516610f526020840184611b07565b6001600160a01b031614610f7957604051630daf43a160e41b815260040160405180910390fd5b610f866020830183611f28565b9050600003610fa857604051630e52390960e41b815260040160405180910390fd5b60008484604051602001610fbd929190611cde565b6040516020818303038152906040529050816110955761100c600682604051610fe69190611d0a565b90815260200160405180910390208460000160208101906110079190611b07565b6111b4565b1561102a57604051631c872d1360e31b815260040160405180910390fd5b611041858561103c6020870187611b07565b6111d6565b61105e576040516311fbe4a960e01b815260040160405180910390fd5b61109361106e6020850185611b07565b60068360405161107e9190611d0a565b908152604051908190036020019020906113c5565b505b6005816040516110a59190611d0a565b90815260405160209181900382019020805460018101825560009182529190208491600302016110d582826120fa565b505083856001600160a01b03167f7fb141aff157414e71b891d33d9759e7c1cc7678d08625f20c79bd081f9ef05885604051610bac9190612248565b6000806000611122878787876113da565b9150915061112f8161149e565b5095945050505050565b606061114f6001600160a01b03831660146115e8565b92915050565b60006111618251611783565b826040516020016111739291906122bd565b604051602081830303815290604052805190602001209050919050565b600080600061119f8585611815565b915091506111ac8161149e565b509392505050565b6001600160a01b038116600090815260018301602052604081205415156106b3565b6040516301ffc9a760e01b8152636cdb3d1360e11b60048201526000906001600160a01b038516906301ffc9a790602401602060405180830381865afa158015611224573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112489190612318565b156112c857604051627eeac760e11b81526001600160a01b038381166004830152602482018590526000919086169062fdd58e90604401602060405180830381865afa15801561129c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c0919061233a565b1190506106b3565b6040516301ffc9a760e01b81526380ac58cd60e01b60048201526001600160a01b038516906301ffc9a790602401602060405180830381865afa158015611313573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113379190612318565b156113bb576040516331a9108f60e11b8152600481018490526001600160a01b038084169190861690636352211e90602401602060405180830381865afa158015611386573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113aa9190612353565b6001600160a01b03161490506106b3565b5060009392505050565b60006106b3836001600160a01b03841661185a565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156114115750600090506003611495565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611465573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661148e57600060019250925050611495565b9150600090505b94509492505050565b60008160048111156114b2576114b2612370565b036114ba5750565b60018160048111156114ce576114ce612370565b0361151b5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610ca9565b600281600481111561152f5761152f612370565b0361157c5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610ca9565b600381600481111561159057611590612370565b03610cbb5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610ca9565b606060006115f7836002612386565b611602906002611dc8565b6001600160401b0381111561161957611619611d5a565b6040519080825280601f01601f191660200182016040528015611643576020820181803683370190505b509050600360fc1b8160008151811061165e5761165e611d70565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061168d5761168d611d70565b60200101906001600160f81b031916908160001a90535060006116b1846002612386565b6116bc906001611dc8565b90505b6001811115611734576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106116f0576116f0611d70565b1a60f81b82828151811061170657611706611d70565b60200101906001600160f81b031916908160001a90535060049490941c9361172d8161239d565b90506116bf565b5083156106b35760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610ca9565b60606000611790836118a9565b60010190506000816001600160401b038111156117af576117af611d5a565b6040519080825280601f01601f1916602001820160405280156117d9576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a85049450846117e357509392505050565b600080825160410361184b5760208301516040840151606085015160001a61183f878285856113da565b94509450505050611853565b506000905060025b9250929050565b60008181526001830160205260408120546118a15750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561114f565b50600061114f565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106118e85772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611914576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061193257662386f26fc10000830492506010015b6305f5e100831061194a576305f5e100830492506008015b612710831061195e57612710830492506004015b60648310611970576064830492506002015b600a831061114f5760010192915050565b604051806060016040528060006001600160a01b0316815260200160608152602001606081525090565b6000602082840312156119bd57600080fd5b5035919050565b6001600160a01b0381168114610cbb57600080fd5b6000806000606084860312156119ee57600080fd5b83356119f9816119c4565b9250602084013591506040840135611a10816119c4565b809150509250925092565b60005b83811015611a36578181015183820152602001611a1e565b50506000910152565b60008151808452611a57816020860160208601611a1b565b601f01601f19169290920160200192915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015611af957888303603f19018552815180516001600160a01b0316845287810151606089860181905290611acb82870182611a3f565b91505087820151915084810388860152611ae58183611a3f565b968901969450505090860190600101611a92565b509098975050505050505050565b600060208284031215611b1957600080fd5b81356106b3816119c4565b60008060008060808587031215611b3a57600080fd5b8435611b45816119c4565b966020860135965060408601359560600135945092505050565b60008060008060608587031215611b7557600080fd5b8435611b80816119c4565b93506020850135925060408501356001600160401b0380821115611ba357600080fd5b818701915087601f830112611bb757600080fd5b813581811115611bc657600080fd5b8860208260051b8501011115611bdb57600080fd5b95989497505060200194505050565b600060608284031215611bfc57600080fd5b50919050565b60008060008060808587031215611c1857600080fd5b8435611c23816119c4565b93506020850135925060408501356001600160401b0380821115611c4657600080fd5b611c5288838901611bea565b93506060870135915080821115611c6857600080fd5b50850160a08188031215611c7b57600080fd5b939692955090935050565b600080600060608486031215611c9b57600080fd5b8335611ca6816119c4565b92506020840135915060408401356001600160401b03811115611cc857600080fd5b611cd486828701611bea565b9150509250925092565b60609290921b6bffffffffffffffffffffffff19168252601f60fa1b6014830152601582015260350190565b60008251611d1c818460208701611a1b565b9190910192915050565b600181811c90821680611d3a57607f821691505b602082108103611bfc57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611dae57611dae611d86565b5060010190565b8181038181111561114f5761114f611d86565b8082018082111561114f5761114f611d86565b601f821115610edd57600081815260208120601f850160051c81016020861015611e025750805b601f850160051c820191505b81811015610c3857828155600101611e0e565b600019600383901b1c191660019190911b1790565b81516001600160401b03811115611e4f57611e4f611d5a565b611e6381611e5d8454611d26565b84611ddb565b602080601f831160018114611e925760008415611e805750858301515b611e8a8582611e21565b865550610c38565b600085815260208120601f198616915b82811015611ec157888601518255948401946001909101908401611ea2565b5085821015611edf5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6020808252810182905260006001600160fb1b03831115611f0f57600080fd5b8260051b80856040850137919091016040019392505050565b6000808335601e19843603018112611f3f57600080fd5b8301803591506001600160401b03821115611f5957600080fd5b60200191503681900382131561185357600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8581526001600160a01b0385166020820152608060408201819052600090611fc29083018587611f6e565b90508260608301529695505050505050565b600060208284031215611fe657600080fd5b813560ff811681146106b357600080fd5b60008351612009818460208801611a1b565b600160fd1b9083019081528351612027816001840160208801611a1b565b601760f91b60019290910191820152600201949350505050565b6001600160401b0383111561205857612058611d5a565b61206c836120668354611d26565b83611ddb565b6000601f84116001811461209a57600085156120885750838201355b6120928682611e21565b845550610d3b565b600083815260209020601f19861690835b828110156120cb57868501358255602094850194600190920191016120ab565b50868210156120e85760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b8135612105816119c4565b81546001600160a01b0319166001600160a01b03919091161781556001818101602061213385820186611f28565b6001600160401b0381111561214a5761214a611d5a565b61215e816121588654611d26565b86611ddb565b6000601f82116001811461218c576000831561217a5750838201355b6121848482611e21565b8755506121e1565b600086815260209020601f19841690835b828110156121ba578685013582559387019390890190870161219d565b50848210156121d75760001960f88660031b161c19848701351681555b50508683881b0186555b505050505050506121f56040830183611f28565b610d3d818360028601612041565b6000808335601e1984360301811261221a57600080fd5b83016020810192503590506001600160401b0381111561223957600080fd5b80360382131561185357600080fd5b6020815260008235612259816119c4565b6001600160a01b031660208381019190915261227790840184612203565b6060604085015261228c608085018284611f6e565b91505061229c6040850185612203565b848303601f190160608601526122b3838284611f6e565b9695505050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a0000000000008152600083516122f581601a850160208801611a1b565b83519083019061230c81601a840160208801611a1b565b01601a01949350505050565b60006020828403121561232a57600080fd5b815180151581146106b357600080fd5b60006020828403121561234c57600080fd5b5051919050565b60006020828403121561236557600080fd5b81516106b3816119c4565b634e487b7160e01b600052602160045260246000fd5b808202811582820484141761114f5761114f611d86565b6000816123ac576123ac611d86565b50600019019056fe417574686f72697a6520746f20777269746520796f7572206461746120746f2074686520636f6e7472616374a2646970667358221220baca6b31fbed1a38d222a84959f4c602da2be437a7f9b98e8b7cbe53cd9085e464736f6c63430008110033",
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

// SetPublicToken is a paid mutator transaction binding the contract method 0x2990b640.
//
// Solidity: function setPublicToken(uint256 publicToken_) returns()
func (_OwnerData *OwnerDataTransactor) SetPublicToken(opts *bind.TransactOpts, publicToken_ *big.Int) (*types.Transaction, error) {
	return _OwnerData.contract.Transact(opts, "setPublicToken", publicToken_)
}

// SetPublicToken is a paid mutator transaction binding the contract method 0x2990b640.
//
// Solidity: function setPublicToken(uint256 publicToken_) returns()
func (_OwnerData *OwnerDataSession) SetPublicToken(publicToken_ *big.Int) (*types.Transaction, error) {
	return _OwnerData.Contract.SetPublicToken(&_OwnerData.TransactOpts, publicToken_)
}

// SetPublicToken is a paid mutator transaction binding the contract method 0x2990b640.
//
// Solidity: function setPublicToken(uint256 publicToken_) returns()
func (_OwnerData *OwnerDataTransactorSession) SetPublicToken(publicToken_ *big.Int) (*types.Transaction, error) {
	return _OwnerData.Contract.SetPublicToken(&_OwnerData.TransactOpts, publicToken_)
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
