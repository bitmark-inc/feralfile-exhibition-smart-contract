// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package airdropv1

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

// FeralFileAirdropV1MetaData contains all meta data concerning the FeralFileAirdropV1 contract.
var FeralFileAirdropV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumFeralFileAirdropV1.Type\",\"name\":\"tokenType_\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"burnable_\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"bridgeable_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"addTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"airdroppedAddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"codeVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"removeTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenType\",\"outputs\":[{\"internalType\":\"enumFeralFileAirdropV1.Type\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"end\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"to_\",\"type\":\"address[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"}],\"name\":\"burnRemaining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002676380380620026768339810160408190526200003491620001f6565b836200004081620000b7565b506200004c33620000c9565b60076200005a84826200032d565b506006805486919060ff1916600183818111156200007c576200007c620003f9565b02179055506006805462ffff0019166101009315159390930262ff00001916929092176201000091151591909102179055506200040f915050565b6002620000c582826200032d565b5050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200014357600080fd5b81516001600160401b03808211156200016057620001606200011b565b604051601f8301601f19908116603f011681019082821181831017156200018b576200018b6200011b565b81604052838152602092508683858801011115620001a857600080fd5b600091505b83821015620001cc5785820183015181830184015290820190620001ad565b600093810190920192909252949350505050565b80518015158114620001f157600080fd5b919050565b600080600080600060a086880312156200020f57600080fd5b8551600281106200021f57600080fd5b60208701519095506001600160401b03808211156200023d57600080fd5b6200024b89838a0162000131565b955060408801519150808211156200026257600080fd5b50620002718882890162000131565b9350506200028260608701620001e0565b91506200029260808701620001e0565b90509295509295909350565b600181811c90821680620002b357607f821691505b602082108103620002d457634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200032857600081815260208120601f850160051c81016020861015620003035750805b601f850160051c820191505b8181101562000324578281556001016200030f565b5050505b505050565b81516001600160401b038111156200034957620003496200011b565b62000361816200035a84546200029e565b84620002da565b602080601f831160018114620003995760008415620003805750858301515b600019600386901b1c1916600185901b17855562000324565b600085815260208120601f198616915b82811015620003ca57888601518255948401946001909101908401620003a9565b5085821015620003e95787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b612257806200041f6000396000f3fe608060405234801561001057600080fd5b50600436106101ce5760003560e01c806396d6708211610104578063dc78ac1c116100a2578063efbe1c1c11610071578063efbe1c1c14610444578063f23a6e611461044c578063f242432a14610478578063f2fde38b1461048b57600080fd5b8063dc78ac1c146103ca578063e8a3d485146103dd578063e985e9c5146103e5578063eee608a41461042157600080fd5b8063a22cb465116100de578063a22cb46514610386578063a4fd6f5614610399578063b390c0ab146103a4578063bdf7a8e6146103b757600080fd5b806396d670821461033e578063985573c214610351578063a07c7ce41461037457600080fd5b806330fa738c1161017157806363e602301161014b57806363e60230146102d7578063715018a6146103085780638da5cb5b14610310578063938e3d7b1461032b57600080fd5b806330fa738c1461028a5780634e1273f4146102a4578063530da8ef146102c457600080fd5b806303120506116101ad57806303120506146102315780630e89341c146102445780631b2ef1ca146102645780632eb2c2d61461027757600080fd5b8062fdd58e146101d357806301ffc9a7146101f957806302fe53051461021c575b600080fd5b6101e66101e136600461177c565b61049e565b6040519081526020015b60405180910390f35b61020c6102073660046117bc565b610537565b60405190151581526020016101f0565b61022f61022a366004611881565b610587565b005b61022f61023f3660046118d2565b61059b565b6102576102523660046118ed565b6105c4565b6040516101f0919061194c565b61022f61027236600461195f565b610658565b61022f610285366004611a36565b6106d4565b6006546102979060ff1681565b6040516101f09190611af6565b6102b76102b2366004611b1e565b610720565b6040516101f09190611c24565b60065461020c9062010000900460ff1681565b61025760405180604001604052806012815260200171466572616c46696c6541697264726f70563160701b81525081565b61022f61084a565b6003546040516001600160a01b0390911681526020016101f0565b61022f610339366004611881565b61085e565b61022f61034c3660046118ed565b61089f565b61020c61035f3660046118d2565b60056020526000908152604090205460ff1681565b60065461020c90610100900460ff1681565b61022f610394366004611c37565b6108bb565b60085460ff1661020c565b61022f6103b236600461195f565b6108c6565b61022f6103c5366004611c73565b610931565b61022f6103d83660046118d2565b610aed565b610257610b19565b61020c6103f3366004611cf2565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b61020c61042f3660046118d2565b60046020526000908152604090205460ff1681565b61022f610bab565b61045f61045a366004611d25565b610be5565b6040516001600160e01b031990911681526020016101f0565b61022f610486366004611d25565b610c6b565b61022f6104993660046118d2565b610cb0565b60006001600160a01b03831661050e5760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a2061646472657373207a65726f206973206e6f742061207660448201526930b634b21037bbb732b960b11b60648201526084015b60405180910390fd5b506000818152602081815260408083206001600160a01b03861684529091529020545b92915050565b60006001600160e01b03198216636cdb3d1360e11b148061056857506001600160e01b031982166303a24d0760e21b145b8061053157506301ffc9a760e01b6001600160e01b0319831614610531565b61058f610d26565b61059881610d80565b50565b6105a3610d26565b6001600160a01b03166000908152600460205260409020805460ff19169055565b6060600280546105d390611d8a565b80601f01602080910402602001604051908101604052809291908181526020018280546105ff90611d8a565b801561064c5780601f106106215761010080835404028352916020019161064c565b820191906000526020600020905b81548152906001019060200180831161062f57829003601f168201915b50505050509050919050565b3360009081526004602052604090205460ff168061068057506003546001600160a01b031633145b61068957600080fd5b60085460ff16156106ac5760405162461bcd60e51b815260040161050590611dc4565b6106b581610d8c565b6106d030838360405180602001604052806000815250610e36565b5050565b6001600160a01b0385163314806106f057506106f085336103f3565b61070c5760405162461bcd60e51b815260040161050590611e05565b6107198585858585610f4a565b5050505050565b606081518351146107855760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610505565b6000835167ffffffffffffffff8111156107a1576107a16117e0565b6040519080825280602002602001820160405280156107ca578160200160208202803683370190505b50905060005b8451811015610842576108158582815181106107ee576107ee611e53565b602002602001015185838151811061080857610808611e53565b602002602001015161049e565b82828151811061082757610827611e53565b602090810291909101015261083b81611e7f565b90506107d0565b509392505050565b610852610d26565b61085c6000611127565b565b610866610d26565b60076108728282611ee3565b506040517fa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad96290600090a150565b6108a7610d26565b61059830826108b6308561049e565b611179565b6106d03383836112f5565b600654610100900460ff1661091d5760405162461bcd60e51b815260206004820181905260248201527f466572616c46696c6541697264726f7056313a206e6f74206275726e61626c656044820152606401610505565b61092681610d8c565b6106d0338383611179565b3360009081526004602052604090205460ff168061095957506003546001600160a01b031633145b61096257600080fd5b60085460ff16156109855760405162461bcd60e51b815260040161050590611dc4565b61098e81610d8c565b60005b81811015610ae757600560008484848181106109af576109af611e53565b90506020020160208101906109c491906118d2565b6001600160a01b0316815260208101919091526040016000205460ff1615610a3d5760405162461bcd60e51b815260206004820152602660248201527f466572616c46696c6541697264726f7056313a20616c726561647920616972646044820152651c9bdc1c195960d21b6064820152608401610505565b610a8030848484818110610a5357610a53611e53565b9050602002016020810190610a6891906118d2565b866001604051806020016040528060008152506113d5565b600160056000858585818110610a9857610a98611e53565b9050602002016020810190610aad91906118d2565b6001600160a01b031681526020810191909152604001600020805460ff191691151591909117905580610adf81611e7f565b915050610991565b50505050565b610af5610d26565b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b606060078054610b2890611d8a565b80601f0160208091040260200160405190810160405280929190818152602001828054610b5490611d8a565b8015610ba15780601f10610b7657610100808354040283529160200191610ba1565b820191906000526020600020905b815481529060010190602001808311610b8457829003601f168201915b5050505050905090565b610bb3610d26565b60085460ff1615610bd65760405162461bcd60e51b815260040161050590611dc4565b6008805460ff19166001179055565b60006001600160a01b03851615610c595760405162461bcd60e51b815260206004820152603260248201527f466572616c46696c6541697264726f7056313a206e6f7420616c6c6f77656420604482015271746f2073656e6420746f6b656e206261636b60701b6064820152608401610505565b5063f23a6e6160e01b95945050505050565b6001600160a01b038516331480610c875750610c8785336103f3565b610ca35760405162461bcd60e51b815260040161050590611e05565b61071985858585856113d5565b610cb8610d26565b6001600160a01b038116610d1d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610505565b61059881611127565b6003546001600160a01b0316331461085c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610505565b60026106d08282611ee3565b600060065460ff166001811115610da557610da5611ae0565b148015610db25750600081115b80610dde5750600160065460ff166001811115610dd157610dd1611ae0565b148015610dde5750806001145b6105985760405162461bcd60e51b815260206004820152602360248201527f466572616c46696c6541697264726f7056313a20616d6f756e74206d69736d616044820152620e8c6d60eb1b6064820152608401610505565b6001600160a01b038416610e965760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610505565b336000610ea2856114ff565b90506000610eaf856114ff565b90506000868152602081815260408083206001600160a01b038b16845290915281208054879290610ee1908490611fa3565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610f418360008989898961154a565b50505050505050565b8151835114610fac5760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610505565b6001600160a01b038416610fd25760405162461bcd60e51b815260040161050590611fb6565b3360005b84518110156110b9576000858281518110610ff357610ff3611e53565b60200260200101519050600085838151811061101157611011611e53565b602090810291909101810151600084815280835260408082206001600160a01b038e1683529093529190912054909150818110156110615760405162461bcd60e51b815260040161050590611ffb565b6000838152602081815260408083206001600160a01b038e8116855292528083208585039055908b1682528120805484929061109e908490611fa3565b92505081905550505050806110b290611e7f565b9050610fd6565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051611109929190612045565b60405180910390a461111f8187878787876116a5565b505050505050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0383166111db5760405162461bcd60e51b815260206004820152602360248201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260448201526265737360e81b6064820152608401610505565b3360006111e7846114ff565b905060006111f4846114ff565b60408051602080820183526000918290528882528181528282206001600160a01b038b168352905220549091508481101561127d5760405162461bcd60e51b8152602060048201526024808201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604482015263616e636560e01b6064820152608401610505565b6000868152602081815260408083206001600160a01b038b81168086529184528285208a8703905582518b81529384018a90529092908816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4604080516020810190915260009052610f41565b816001600160a01b0316836001600160a01b0316036113685760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610505565b6001600160a01b03838116600081815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0384166113fb5760405162461bcd60e51b815260040161050590611fb6565b336000611407856114ff565b90506000611414856114ff565b90506000868152602081815260408083206001600160a01b038c168452909152902054858110156114575760405162461bcd60e51b815260040161050590611ffb565b6000878152602081815260408083206001600160a01b038d8116855292528083208985039055908a16825281208054889290611494908490611fa3565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46114f4848a8a8a8a8a61154a565b505050505050505050565b6040805160018082528183019092526060916000919060208083019080368337019050509050828160008151811061153957611539611e53565b602090810291909101015292915050565b6001600160a01b0384163b1561111f5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e619061158e9089908990889088908890600401612073565b6020604051808303816000875af19250505080156115c9575060408051601f3d908101601f191682019092526115c6918101906120b8565b60015b611675576115d56120d5565b806308c379a00361160e57506115e96120f1565b806115f45750611610565b8060405162461bcd60e51b8152600401610505919061194c565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e2d455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610505565b6001600160e01b0319811663f23a6e6160e01b14610f415760405162461bcd60e51b81526004016105059061217b565b6001600160a01b0384163b1561111f5760405163bc197c8160e01b81526001600160a01b0385169063bc197c81906116e990899089908890889088906004016121c3565b6020604051808303816000875af1925050508015611724575060408051601f3d908101601f19168201909252611721918101906120b8565b60015b611730576115d56120d5565b6001600160e01b0319811663bc197c8160e01b14610f415760405162461bcd60e51b81526004016105059061217b565b80356001600160a01b038116811461177757600080fd5b919050565b6000806040838503121561178f57600080fd5b61179883611760565b946020939093013593505050565b6001600160e01b03198116811461059857600080fd5b6000602082840312156117ce57600080fd5b81356117d9816117a6565b9392505050565b634e487b7160e01b600052604160045260246000fd5b601f8201601f1916810167ffffffffffffffff8111828210171561181c5761181c6117e0565b6040525050565b600067ffffffffffffffff83111561183d5761183d6117e0565b604051611854601f8501601f1916602001826117f6565b80915083815284848401111561186957600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561189357600080fd5b813567ffffffffffffffff8111156118aa57600080fd5b8201601f810184136118bb57600080fd5b6118ca84823560208401611823565b949350505050565b6000602082840312156118e457600080fd5b6117d982611760565b6000602082840312156118ff57600080fd5b5035919050565b6000815180845260005b8181101561192c57602081850181015186830182015201611910565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006117d96020830184611906565b6000806040838503121561197257600080fd5b50508035926020909101359150565b600067ffffffffffffffff82111561199b5761199b6117e0565b5060051b60200190565b600082601f8301126119b657600080fd5b813560206119c382611981565b6040516119d082826117f6565b83815260059390931b85018201928281019150868411156119f057600080fd5b8286015b84811015611a0b57803583529183019183016119f4565b509695505050505050565b600082601f830112611a2757600080fd5b6117d983833560208501611823565b600080600080600060a08688031215611a4e57600080fd5b611a5786611760565b9450611a6560208701611760565b9350604086013567ffffffffffffffff80821115611a8257600080fd5b611a8e89838a016119a5565b94506060880135915080821115611aa457600080fd5b611ab089838a016119a5565b93506080880135915080821115611ac657600080fd5b50611ad388828901611a16565b9150509295509295909350565b634e487b7160e01b600052602160045260246000fd5b6020810160028310611b1857634e487b7160e01b600052602160045260246000fd5b91905290565b60008060408385031215611b3157600080fd5b823567ffffffffffffffff80821115611b4957600080fd5b818501915085601f830112611b5d57600080fd5b81356020611b6a82611981565b604051611b7782826117f6565b83815260059390931b8501820192828101915089841115611b9757600080fd5b948201945b83861015611bbc57611bad86611760565b82529482019490820190611b9c565b96505086013592505080821115611bd257600080fd5b50611bdf858286016119a5565b9150509250929050565b600081518084526020808501945080840160005b83811015611c1957815187529582019590820190600101611bfd565b509495945050505050565b6020815260006117d96020830184611be9565b60008060408385031215611c4a57600080fd5b611c5383611760565b915060208301358015158114611c6857600080fd5b809150509250929050565b600080600060408486031215611c8857600080fd5b83359250602084013567ffffffffffffffff80821115611ca757600080fd5b818601915086601f830112611cbb57600080fd5b813581811115611cca57600080fd5b8760208260051b8501011115611cdf57600080fd5b6020830194508093505050509250925092565b60008060408385031215611d0557600080fd5b611d0e83611760565b9150611d1c60208401611760565b90509250929050565b600080600080600060a08688031215611d3d57600080fd5b611d4686611760565b9450611d5460208701611760565b93506040860135925060608601359150608086013567ffffffffffffffff811115611d7e57600080fd5b611ad388828901611a16565b600181811c90821680611d9e57607f821691505b602082108103611dbe57634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526021908201527f466572616c46696c6541697264726f7056313a20616c726561647920656e64656040820152601960fa1b606082015260800190565b6020808252602e908201527f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60408201526d195c881bdc88185c1c1c9bdd995960921b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611e9157611e91611e69565b5060010190565b601f821115611ede57600081815260208120601f850160051c81016020861015611ebf5750805b601f850160051c820191505b8181101561111f57828155600101611ecb565b505050565b815167ffffffffffffffff811115611efd57611efd6117e0565b611f1181611f0b8454611d8a565b84611e98565b602080601f831160018114611f465760008415611f2e5750858301515b600019600386901b1c1916600185901b17855561111f565b600085815260208120601f198616915b82811015611f7557888601518255948401946001909101908401611f56565b5085821015611f935787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b8082018082111561053157610531611e69565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b6040815260006120586040830185611be9565b828103602084015261206a8185611be9565b95945050505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190526000906120ad90830184611906565b979650505050505050565b6000602082840312156120ca57600080fd5b81516117d9816117a6565b600060033d11156120ee5760046000803e5060005160e01c5b90565b600060443d10156120ff5790565b6040516003193d81016004833e81513d67ffffffffffffffff816024840111818411171561212f57505050505090565b82850191508151818111156121475750505050505090565b843d87010160208285010111156121615750505050505090565b612170602082860101876117f6565b509095945050505050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b6001600160a01b0386811682528516602082015260a0604082018190526000906121ef90830186611be9565b82810360608401526122018186611be9565b905082810360808401526122158185611906565b9897505050505050505056fea2646970667358221220f52c7869d614dbda9c3c5b52ad5c19ae8c928176e117be1544e90fd35acb019864736f6c63430008110033",
}

// FeralFileAirdropV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use FeralFileAirdropV1MetaData.ABI instead.
var FeralFileAirdropV1ABI = FeralFileAirdropV1MetaData.ABI

// FeralFileAirdropV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeralFileAirdropV1MetaData.Bin instead.
var FeralFileAirdropV1Bin = FeralFileAirdropV1MetaData.Bin

// DeployFeralFileAirdropV1 deploys a new Ethereum contract, binding an instance of FeralFileAirdropV1 to it.
func DeployFeralFileAirdropV1(auth *bind.TransactOpts, backend bind.ContractBackend, tokenType_ uint8, tokenURI_ string, contractURI_ string, burnable_ bool, bridgeable_ bool) (common.Address, *types.Transaction, *FeralFileAirdropV1, error) {
	parsed, err := FeralFileAirdropV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeralFileAirdropV1Bin), backend, tokenType_, tokenURI_, contractURI_, burnable_, bridgeable_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeralFileAirdropV1{FeralFileAirdropV1Caller: FeralFileAirdropV1Caller{contract: contract}, FeralFileAirdropV1Transactor: FeralFileAirdropV1Transactor{contract: contract}, FeralFileAirdropV1Filterer: FeralFileAirdropV1Filterer{contract: contract}}, nil
}

// FeralFileAirdropV1 is an auto generated Go binding around an Ethereum contract.
type FeralFileAirdropV1 struct {
	FeralFileAirdropV1Caller     // Read-only binding to the contract
	FeralFileAirdropV1Transactor // Write-only binding to the contract
	FeralFileAirdropV1Filterer   // Log filterer for contract events
}

// FeralFileAirdropV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type FeralFileAirdropV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralFileAirdropV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FeralFileAirdropV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralFileAirdropV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeralFileAirdropV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralFileAirdropV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeralFileAirdropV1Session struct {
	Contract     *FeralFileAirdropV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FeralFileAirdropV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeralFileAirdropV1CallerSession struct {
	Contract *FeralFileAirdropV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// FeralFileAirdropV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeralFileAirdropV1TransactorSession struct {
	Contract     *FeralFileAirdropV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// FeralFileAirdropV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type FeralFileAirdropV1Raw struct {
	Contract *FeralFileAirdropV1 // Generic contract binding to access the raw methods on
}

// FeralFileAirdropV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeralFileAirdropV1CallerRaw struct {
	Contract *FeralFileAirdropV1Caller // Generic read-only contract binding to access the raw methods on
}

// FeralFileAirdropV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeralFileAirdropV1TransactorRaw struct {
	Contract *FeralFileAirdropV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFeralFileAirdropV1 creates a new instance of FeralFileAirdropV1, bound to a specific deployed contract.
func NewFeralFileAirdropV1(address common.Address, backend bind.ContractBackend) (*FeralFileAirdropV1, error) {
	contract, err := bindFeralFileAirdropV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeralFileAirdropV1{FeralFileAirdropV1Caller: FeralFileAirdropV1Caller{contract: contract}, FeralFileAirdropV1Transactor: FeralFileAirdropV1Transactor{contract: contract}, FeralFileAirdropV1Filterer: FeralFileAirdropV1Filterer{contract: contract}}, nil
}

// NewFeralFileAirdropV1Caller creates a new read-only instance of FeralFileAirdropV1, bound to a specific deployed contract.
func NewFeralFileAirdropV1Caller(address common.Address, caller bind.ContractCaller) (*FeralFileAirdropV1Caller, error) {
	contract, err := bindFeralFileAirdropV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeralFileAirdropV1Caller{contract: contract}, nil
}

// NewFeralFileAirdropV1Transactor creates a new write-only instance of FeralFileAirdropV1, bound to a specific deployed contract.
func NewFeralFileAirdropV1Transactor(address common.Address, transactor bind.ContractTransactor) (*FeralFileAirdropV1Transactor, error) {
	contract, err := bindFeralFileAirdropV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeralFileAirdropV1Transactor{contract: contract}, nil
}

// NewFeralFileAirdropV1Filterer creates a new log filterer instance of FeralFileAirdropV1, bound to a specific deployed contract.
func NewFeralFileAirdropV1Filterer(address common.Address, filterer bind.ContractFilterer) (*FeralFileAirdropV1Filterer, error) {
	contract, err := bindFeralFileAirdropV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeralFileAirdropV1Filterer{contract: contract}, nil
}

// bindFeralFileAirdropV1 binds a generic wrapper to an already deployed contract.
func bindFeralFileAirdropV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeralFileAirdropV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralFileAirdropV1 *FeralFileAirdropV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralFileAirdropV1.Contract.FeralFileAirdropV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralFileAirdropV1 *FeralFileAirdropV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.FeralFileAirdropV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralFileAirdropV1 *FeralFileAirdropV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.FeralFileAirdropV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralFileAirdropV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.contract.Transact(opts, method, params...)
}

// AirdroppedAddresses is a free data retrieval call binding the contract method 0x985573c2.
//
// Solidity: function airdroppedAddresses(address ) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) AirdroppedAddresses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "airdroppedAddresses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AirdroppedAddresses is a free data retrieval call binding the contract method 0x985573c2.
//
// Solidity: function airdroppedAddresses(address ) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) AirdroppedAddresses(arg0 common.Address) (bool, error) {
	return _FeralFileAirdropV1.Contract.AirdroppedAddresses(&_FeralFileAirdropV1.CallOpts, arg0)
}

// AirdroppedAddresses is a free data retrieval call binding the contract method 0x985573c2.
//
// Solidity: function airdroppedAddresses(address ) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) AirdroppedAddresses(arg0 common.Address) (bool, error) {
	return _FeralFileAirdropV1.Contract.AirdroppedAddresses(&_FeralFileAirdropV1.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _FeralFileAirdropV1.Contract.BalanceOf(&_FeralFileAirdropV1.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _FeralFileAirdropV1.Contract.BalanceOf(&_FeralFileAirdropV1.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _FeralFileAirdropV1.Contract.BalanceOfBatch(&_FeralFileAirdropV1.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _FeralFileAirdropV1.Contract.BalanceOfBatch(&_FeralFileAirdropV1.CallOpts, accounts, ids)
}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) Bridgeable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "bridgeable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) Bridgeable() (bool, error) {
	return _FeralFileAirdropV1.Contract.Bridgeable(&_FeralFileAirdropV1.CallOpts)
}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) Bridgeable() (bool, error) {
	return _FeralFileAirdropV1.Contract.Bridgeable(&_FeralFileAirdropV1.CallOpts)
}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) Burnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "burnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) Burnable() (bool, error) {
	return _FeralFileAirdropV1.Contract.Burnable(&_FeralFileAirdropV1.CallOpts)
}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) Burnable() (bool, error) {
	return _FeralFileAirdropV1.Contract.Burnable(&_FeralFileAirdropV1.CallOpts)
}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) CodeVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "codeVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) CodeVersion() (string, error) {
	return _FeralFileAirdropV1.Contract.CodeVersion(&_FeralFileAirdropV1.CallOpts)
}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) CodeVersion() (string, error) {
	return _FeralFileAirdropV1.Contract.CodeVersion(&_FeralFileAirdropV1.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) ContractURI() (string, error) {
	return _FeralFileAirdropV1.Contract.ContractURI(&_FeralFileAirdropV1.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) ContractURI() (string, error) {
	return _FeralFileAirdropV1.Contract.ContractURI(&_FeralFileAirdropV1.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _FeralFileAirdropV1.Contract.IsApprovedForAll(&_FeralFileAirdropV1.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _FeralFileAirdropV1.Contract.IsApprovedForAll(&_FeralFileAirdropV1.CallOpts, account, operator)
}

// IsEnded is a free data retrieval call binding the contract method 0xa4fd6f56.
//
// Solidity: function isEnded() view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) IsEnded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "isEnded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnded is a free data retrieval call binding the contract method 0xa4fd6f56.
//
// Solidity: function isEnded() view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) IsEnded() (bool, error) {
	return _FeralFileAirdropV1.Contract.IsEnded(&_FeralFileAirdropV1.CallOpts)
}

// IsEnded is a free data retrieval call binding the contract method 0xa4fd6f56.
//
// Solidity: function isEnded() view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) IsEnded() (bool, error) {
	return _FeralFileAirdropV1.Contract.IsEnded(&_FeralFileAirdropV1.CallOpts)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address from_, uint256 , uint256 , bytes ) pure returns(bytes4)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, from_ common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "onERC1155Received", arg0, from_, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address from_, uint256 , uint256 , bytes ) pure returns(bytes4)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) OnERC1155Received(arg0 common.Address, from_ common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _FeralFileAirdropV1.Contract.OnERC1155Received(&_FeralFileAirdropV1.CallOpts, arg0, from_, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address from_, uint256 , uint256 , bytes ) pure returns(bytes4)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) OnERC1155Received(arg0 common.Address, from_ common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _FeralFileAirdropV1.Contract.OnERC1155Received(&_FeralFileAirdropV1.CallOpts, arg0, from_, arg2, arg3, arg4)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) Owner() (common.Address, error) {
	return _FeralFileAirdropV1.Contract.Owner(&_FeralFileAirdropV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) Owner() (common.Address, error) {
	return _FeralFileAirdropV1.Contract.Owner(&_FeralFileAirdropV1.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeralFileAirdropV1.Contract.SupportsInterface(&_FeralFileAirdropV1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeralFileAirdropV1.Contract.SupportsInterface(&_FeralFileAirdropV1.CallOpts, interfaceId)
}

// TokenType is a free data retrieval call binding the contract method 0x30fa738c.
//
// Solidity: function tokenType() view returns(uint8)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) TokenType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "tokenType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenType is a free data retrieval call binding the contract method 0x30fa738c.
//
// Solidity: function tokenType() view returns(uint8)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) TokenType() (uint8, error) {
	return _FeralFileAirdropV1.Contract.TokenType(&_FeralFileAirdropV1.CallOpts)
}

// TokenType is a free data retrieval call binding the contract method 0x30fa738c.
//
// Solidity: function tokenType() view returns(uint8)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) TokenType() (uint8, error) {
	return _FeralFileAirdropV1.Contract.TokenType(&_FeralFileAirdropV1.CallOpts)
}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) Trustees(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "trustees", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) Trustees(arg0 common.Address) (bool, error) {
	return _FeralFileAirdropV1.Contract.Trustees(&_FeralFileAirdropV1.CallOpts, arg0)
}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) Trustees(arg0 common.Address) (bool, error) {
	return _FeralFileAirdropV1.Contract.Trustees(&_FeralFileAirdropV1.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _FeralFileAirdropV1.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) Uri(arg0 *big.Int) (string, error) {
	return _FeralFileAirdropV1.Contract.Uri(&_FeralFileAirdropV1.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_FeralFileAirdropV1 *FeralFileAirdropV1CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _FeralFileAirdropV1.Contract.Uri(&_FeralFileAirdropV1.CallOpts, arg0)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) AddTrustee(opts *bind.TransactOpts, _trustee common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "addTrustee", _trustee)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) AddTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.AddTrustee(&_FeralFileAirdropV1.TransactOpts, _trustee)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) AddTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.AddTrustee(&_FeralFileAirdropV1.TransactOpts, _trustee)
}

// Airdrop is a paid mutator transaction binding the contract method 0xbdf7a8e6.
//
// Solidity: function airdrop(uint256 tokenID_, address[] to_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) Airdrop(opts *bind.TransactOpts, tokenID_ *big.Int, to_ []common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "airdrop", tokenID_, to_)
}

// Airdrop is a paid mutator transaction binding the contract method 0xbdf7a8e6.
//
// Solidity: function airdrop(uint256 tokenID_, address[] to_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) Airdrop(tokenID_ *big.Int, to_ []common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.Airdrop(&_FeralFileAirdropV1.TransactOpts, tokenID_, to_)
}

// Airdrop is a paid mutator transaction binding the contract method 0xbdf7a8e6.
//
// Solidity: function airdrop(uint256 tokenID_, address[] to_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) Airdrop(tokenID_ *big.Int, to_ []common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.Airdrop(&_FeralFileAirdropV1.TransactOpts, tokenID_, to_)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 tokenID_, uint256 amount_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) Burn(opts *bind.TransactOpts, tokenID_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "burn", tokenID_, amount_)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 tokenID_, uint256 amount_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) Burn(tokenID_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.Burn(&_FeralFileAirdropV1.TransactOpts, tokenID_, amount_)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 tokenID_, uint256 amount_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) Burn(tokenID_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.Burn(&_FeralFileAirdropV1.TransactOpts, tokenID_, amount_)
}

// BurnRemaining is a paid mutator transaction binding the contract method 0x96d67082.
//
// Solidity: function burnRemaining(uint256 tokenID_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) BurnRemaining(opts *bind.TransactOpts, tokenID_ *big.Int) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "burnRemaining", tokenID_)
}

// BurnRemaining is a paid mutator transaction binding the contract method 0x96d67082.
//
// Solidity: function burnRemaining(uint256 tokenID_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) BurnRemaining(tokenID_ *big.Int) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.BurnRemaining(&_FeralFileAirdropV1.TransactOpts, tokenID_)
}

// BurnRemaining is a paid mutator transaction binding the contract method 0x96d67082.
//
// Solidity: function burnRemaining(uint256 tokenID_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) BurnRemaining(tokenID_ *big.Int) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.BurnRemaining(&_FeralFileAirdropV1.TransactOpts, tokenID_)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) End(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "end")
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) End() (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.End(&_FeralFileAirdropV1.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) End() (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.End(&_FeralFileAirdropV1.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenID_, uint256 amount_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) Mint(opts *bind.TransactOpts, tokenID_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "mint", tokenID_, amount_)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenID_, uint256 amount_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) Mint(tokenID_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.Mint(&_FeralFileAirdropV1.TransactOpts, tokenID_, amount_)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenID_, uint256 amount_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) Mint(tokenID_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.Mint(&_FeralFileAirdropV1.TransactOpts, tokenID_, amount_)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) RemoveTrustee(opts *bind.TransactOpts, _trustee common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "removeTrustee", _trustee)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) RemoveTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.RemoveTrustee(&_FeralFileAirdropV1.TransactOpts, _trustee)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) RemoveTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.RemoveTrustee(&_FeralFileAirdropV1.TransactOpts, _trustee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.RenounceOwnership(&_FeralFileAirdropV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.RenounceOwnership(&_FeralFileAirdropV1.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.SafeBatchTransferFrom(&_FeralFileAirdropV1.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.SafeBatchTransferFrom(&_FeralFileAirdropV1.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.SafeTransferFrom(&_FeralFileAirdropV1.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.SafeTransferFrom(&_FeralFileAirdropV1.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.SetApprovalForAll(&_FeralFileAirdropV1.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.SetApprovalForAll(&_FeralFileAirdropV1.TransactOpts, operator, approved)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) SetContractURI(opts *bind.TransactOpts, contractURI_ string) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "setContractURI", contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.SetContractURI(&_FeralFileAirdropV1.TransactOpts, contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.SetContractURI(&_FeralFileAirdropV1.TransactOpts, contractURI_)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string uri_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) SetURI(opts *bind.TransactOpts, uri_ string) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "setURI", uri_)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string uri_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) SetURI(uri_ string) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.SetURI(&_FeralFileAirdropV1.TransactOpts, uri_)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string uri_) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) SetURI(uri_ string) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.SetURI(&_FeralFileAirdropV1.TransactOpts, uri_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.TransferOwnership(&_FeralFileAirdropV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralFileAirdropV1 *FeralFileAirdropV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralFileAirdropV1.Contract.TransferOwnership(&_FeralFileAirdropV1.TransactOpts, newOwner)
}

// FeralFileAirdropV1ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1ApprovalForAllIterator struct {
	Event *FeralFileAirdropV1ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FeralFileAirdropV1ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralFileAirdropV1ApprovalForAll)
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
		it.Event = new(FeralFileAirdropV1ApprovalForAll)
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
func (it *FeralFileAirdropV1ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralFileAirdropV1ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralFileAirdropV1ApprovalForAll represents a ApprovalForAll event raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*FeralFileAirdropV1ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FeralFileAirdropV1.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FeralFileAirdropV1ApprovalForAllIterator{contract: _FeralFileAirdropV1.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FeralFileAirdropV1ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FeralFileAirdropV1.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralFileAirdropV1ApprovalForAll)
				if err := _FeralFileAirdropV1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) ParseApprovalForAll(log types.Log) (*FeralFileAirdropV1ApprovalForAll, error) {
	event := new(FeralFileAirdropV1ApprovalForAll)
	if err := _FeralFileAirdropV1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralFileAirdropV1ContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1ContractURIUpdatedIterator struct {
	Event *FeralFileAirdropV1ContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *FeralFileAirdropV1ContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralFileAirdropV1ContractURIUpdated)
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
		it.Event = new(FeralFileAirdropV1ContractURIUpdated)
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
func (it *FeralFileAirdropV1ContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralFileAirdropV1ContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralFileAirdropV1ContractURIUpdated represents a ContractURIUpdated event raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1ContractURIUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*FeralFileAirdropV1ContractURIUpdatedIterator, error) {

	logs, sub, err := _FeralFileAirdropV1.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &FeralFileAirdropV1ContractURIUpdatedIterator{contract: _FeralFileAirdropV1.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *FeralFileAirdropV1ContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _FeralFileAirdropV1.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralFileAirdropV1ContractURIUpdated)
				if err := _FeralFileAirdropV1.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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

// ParseContractURIUpdated is a log parse operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) ParseContractURIUpdated(log types.Log) (*FeralFileAirdropV1ContractURIUpdated, error) {
	event := new(FeralFileAirdropV1ContractURIUpdated)
	if err := _FeralFileAirdropV1.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralFileAirdropV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1OwnershipTransferredIterator struct {
	Event *FeralFileAirdropV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeralFileAirdropV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralFileAirdropV1OwnershipTransferred)
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
		it.Event = new(FeralFileAirdropV1OwnershipTransferred)
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
func (it *FeralFileAirdropV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralFileAirdropV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralFileAirdropV1OwnershipTransferred represents a OwnershipTransferred event raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeralFileAirdropV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralFileAirdropV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeralFileAirdropV1OwnershipTransferredIterator{contract: _FeralFileAirdropV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeralFileAirdropV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralFileAirdropV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralFileAirdropV1OwnershipTransferred)
				if err := _FeralFileAirdropV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) ParseOwnershipTransferred(log types.Log) (*FeralFileAirdropV1OwnershipTransferred, error) {
	event := new(FeralFileAirdropV1OwnershipTransferred)
	if err := _FeralFileAirdropV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralFileAirdropV1TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1TransferBatchIterator struct {
	Event *FeralFileAirdropV1TransferBatch // Event containing the contract specifics and raw log

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
func (it *FeralFileAirdropV1TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralFileAirdropV1TransferBatch)
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
		it.Event = new(FeralFileAirdropV1TransferBatch)
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
func (it *FeralFileAirdropV1TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralFileAirdropV1TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralFileAirdropV1TransferBatch represents a TransferBatch event raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*FeralFileAirdropV1TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeralFileAirdropV1.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeralFileAirdropV1TransferBatchIterator{contract: _FeralFileAirdropV1.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *FeralFileAirdropV1TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeralFileAirdropV1.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralFileAirdropV1TransferBatch)
				if err := _FeralFileAirdropV1.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) ParseTransferBatch(log types.Log) (*FeralFileAirdropV1TransferBatch, error) {
	event := new(FeralFileAirdropV1TransferBatch)
	if err := _FeralFileAirdropV1.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralFileAirdropV1TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1TransferSingleIterator struct {
	Event *FeralFileAirdropV1TransferSingle // Event containing the contract specifics and raw log

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
func (it *FeralFileAirdropV1TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralFileAirdropV1TransferSingle)
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
		it.Event = new(FeralFileAirdropV1TransferSingle)
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
func (it *FeralFileAirdropV1TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralFileAirdropV1TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralFileAirdropV1TransferSingle represents a TransferSingle event raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*FeralFileAirdropV1TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeralFileAirdropV1.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeralFileAirdropV1TransferSingleIterator{contract: _FeralFileAirdropV1.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *FeralFileAirdropV1TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeralFileAirdropV1.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralFileAirdropV1TransferSingle)
				if err := _FeralFileAirdropV1.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) ParseTransferSingle(log types.Log) (*FeralFileAirdropV1TransferSingle, error) {
	event := new(FeralFileAirdropV1TransferSingle)
	if err := _FeralFileAirdropV1.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralFileAirdropV1URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1URIIterator struct {
	Event *FeralFileAirdropV1URI // Event containing the contract specifics and raw log

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
func (it *FeralFileAirdropV1URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralFileAirdropV1URI)
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
		it.Event = new(FeralFileAirdropV1URI)
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
func (it *FeralFileAirdropV1URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralFileAirdropV1URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralFileAirdropV1URI represents a URI event raised by the FeralFileAirdropV1 contract.
type FeralFileAirdropV1URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*FeralFileAirdropV1URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FeralFileAirdropV1.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &FeralFileAirdropV1URIIterator{contract: _FeralFileAirdropV1.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *FeralFileAirdropV1URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FeralFileAirdropV1.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralFileAirdropV1URI)
				if err := _FeralFileAirdropV1.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_FeralFileAirdropV1 *FeralFileAirdropV1Filterer) ParseURI(log types.Log) (*FeralFileAirdropV1URI, error) {
	event := new(FeralFileAirdropV1URI)
	if err := _FeralFileAirdropV1.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
