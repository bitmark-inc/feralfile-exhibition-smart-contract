// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feralfilev4_2

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

// FeralfileExhibitionV4Artwork is an auto generated low-level Go binding around an user-defined struct.
type FeralfileExhibitionV4Artwork struct {
	SeriesId *big.Int
	TokenId  *big.Int
}

// FeralfileExhibitionV4MintData is an auto generated low-level Go binding around an user-defined struct.
type FeralfileExhibitionV4MintData struct {
	SeriesId *big.Int
	TokenId  *big.Int
	Owner    common.Address
}

// FeralfileExhibitionV42MintDataWithIndex is an auto generated low-level Go binding around an user-defined struct.
type FeralfileExhibitionV42MintDataWithIndex struct {
	SeriesId   *big.Int
	TokenId    *big.Int
	Owner      common.Address
	TokenIndex *big.Int
}

// IFeralfileSaleDataRevenueShare is an auto generated low-level Go binding around an user-defined struct.
type IFeralfileSaleDataRevenueShare struct {
	Recipient common.Address
	Bps       *big.Int
}

// IFeralfileSaleDataSaleData is an auto generated low-level Go binding around an user-defined struct.
type IFeralfileSaleDataSaleData struct {
	Price              *big.Int
	Cost               *big.Int
	ExpiryTime         *big.Int
	Destination        common.Address
	TokenIds           []*big.Int
	RevenueShares      [][]IFeralfileSaleDataRevenueShare
	PayByVaultContract bool
}

// FeralfileExhibitionV42MetaData contains all meta data concerning the FeralfileExhibitionV42 contract.
var FeralfileExhibitionV42MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"burnable_\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"bridgeable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"costReceiver_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"seriesIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"seriesMaxSupplies_\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"erc20ContractAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdvanceAddressAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FunctionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAddressesAndAmounts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintNotEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIDNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BurnArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BuyArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NewArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OperatorFilterRegistry\",\"outputs\":[{\"internalType\":\"contractIOperatorFilterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"addTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"advances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artworkFileURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burnArtworks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"codeVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"costReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getArtwork\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structFeralfileExhibitionV4.Artwork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"removeTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oldAddresses_\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"newAddresses_\",\"type\":\"address[]\"}],\"name\":\"replaceAdvanceAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selling\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"seriesMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"seriesTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"setAdvanceSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"costReceiver_\",\"type\":\"address\"}],\"name\":\"setCostReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setTokenBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopSaleAndBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"seriesIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"recipientAddresses\",\"type\":\"address[]\"}],\"name\":\"stopSaleAndTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPlaceholderAnimationURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPlaceholderImageURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorFilterRegisterAddress\",\"type\":\"address\"}],\"name\":\"updateOperatorFilterRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIFeralfileVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFeralfileExhibitionV4_2.MintDataWithIndex[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"mintArtworksWithIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structFeralfileExhibitionV4.MintData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"mintArtworks\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"parameters\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"coinAmount\",\"type\":\"uint256\"}],\"name\":\"updateTokenInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"internalType\":\"structIFeralfileSaleData.RevenueShare[][]\",\"name\":\"revenueShares\",\"type\":\"tuple[][]\"},{\"internalType\":\"bool\",\"name\":\"payByVaultContract\",\"type\":\"bool\"}],\"internalType\":\"structIFeralfileSaleData.SaleData\",\"name\":\"saleData_\",\"type\":\"tuple\"}],\"name\":\"buyArtworks\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setArtworkFileURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setTokenPlaceholderImageURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setTokenPlaceholderAnimationURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600880546001600160a01b0319166daaeb6d7670e522a718067333cd4e179055600d805463ff000000191663010000001790553480156200004457600080fd5b5060405162006a0638038062006a068339810160408190526200006791620009ad565b8a8a8a8a8a8a8a8a8a8a89898989898989898989858a8a60006200008c838262000bad565b5060016200009b828262000bad565b505050620000b8620000b2620007cb60201b60201c565b620007cf565b6008546001600160a01b03163b156200014557600854604051633e9f1edf60e11b8152306004820152733cc6cdda760b79bafa08df41ecfa224f810dceb660248201526001600160a01b0390911690637d3e3dbe90604401600060405180830381600087803b1580156200012b57600080fd5b505af115801562000140573d6000803e3d6000fd5b505050505b6001600160a01b038116620001ac5760405162461bcd60e51b815260206004820152602260248201527f45434453415369676e3a207369676e65725f206973207a65726f206164647265604482015261737360f01b60648201526084015b60405180910390fd5b600980546001600160a01b0319166001600160a01b039290921691909117905589516200022a5760405162461bcd60e51b815260206004820152602560248201527f466572616c66696c6545786869626974696f6e56343a206e616d655f20697320604482015264656d70747960d81b6064820152608401620001a3565b60008951116200028d5760405162461bcd60e51b815260206004820152602760248201527f466572616c66696c6545786869626974696f6e56343a2073796d626f6c5f20696044820152667320656d70747960c81b6064820152608401620001a3565b6001600160a01b0385166200030b5760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a207661756c744164647260448201527f6573735f206973207a65726f20616464726573730000000000000000000000006064820152608401620001a3565b6001600160a01b038416620003895760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a20636f7374526563656960448201527f7665725f206973207a65726f20616464726573730000000000000000000000006064820152608401620001a3565b6000835111620003f15760405162461bcd60e51b815260206004820152602c60248201527f466572616c66696c6545786869626974696f6e56343a20636f6e74726163745560448201526b52495f20697320656d70747960a01b6064820152608401620001a3565b6000825111620004575760405162461bcd60e51b815260206004820152602a60248201527f466572616c66696c6545786869626974696f6e56343a207365726965734964736044820152695f20697320656d70747960b01b6064820152608401620001a3565b6000815111620004c55760405162461bcd60e51b815260206004820152603260248201527f466572616c66696c6545786869626974696f6e56343a205f7365726965734d6160448201527178537570706c69657320697320656d70747960701b6064820152608401620001a3565b8051825114620005585760405162461bcd60e51b815260206004820152605160248201527f466572616c66696c6545786869626974696f6e56343a207365726965734d617860448201527f537570706c6965735f20616e64207365726965734964735f206c656e6774687360648201527020617265206e6f74207468652073616d6560781b608482015260a401620001a3565b600d805461ffff191689151561ff001916176101008915150217600160201b600160c01b0319166401000000006001600160a01b038781169190910291909117909155600e80546001600160a01b031916918716919091179055600b620005c0848262000bad565b5060005b82518110156200077e576000620005dd82600162000c8f565b90505b83518110156200069a57838181518110620005ff57620005ff62000cab565b60200260200101518483815181106200061c576200061c62000cab565b602002602001015103620006855760405162461bcd60e51b815260206004820152602960248201527f466572616c66696c6545786869626974696f6e56343a206475706c6963617465604482015268081cd95c9a595cd25960ba1b6064820152608401620001a3565b80620006918162000cc1565b915050620005e0565b506000828281518110620006b257620006b262000cab565b602002602001015111620007185760405162461bcd60e51b815260206004820152602660248201527f466572616c66696c6545786869626974696f6e56343a207a65726f206d617820604482015265737570706c7960d01b6064820152608401620001a3565b8181815181106200072d576200072d62000cab565b6020026020010151600f60008584815181106200074e576200074e62000cab565b60200260200101518152602001908152602001600020819055508080620007759062000cc1565b915050620005c4565b50505050505050505050505050505050505050505080601560006101000a8154816001600160a01b0302191690836001600160a01b03160217905550505050505050505050505062000cdd565b3390565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171562000862576200086262000821565b604052919050565b600082601f8301126200087c57600080fd5b81516001600160401b0381111562000898576200089862000821565b6020620008ae601f8301601f1916820162000837565b8281528582848701011115620008c357600080fd5b60005b83811015620008e3578581018301518282018401528201620008c6565b506000928101909101919091529392505050565b805180151581146200090857600080fd5b919050565b80516001600160a01b03811681146200090857600080fd5b600082601f8301126200093757600080fd5b815160206001600160401b0382111562000955576200095562000821565b8160051b6200096682820162000837565b92835284810182019282810190878511156200098157600080fd5b83870192505b84831015620009a25782518252918301919083019062000987565b979650505050505050565b60008060008060008060008060008060006101608c8e031215620009d057600080fd5b8b516001600160401b03811115620009e757600080fd5b620009f58e828f016200086a565b60208e0151909c5090506001600160401b0381111562000a1457600080fd5b62000a228e828f016200086a565b9a505062000a3360408d01620008f7565b985062000a4360608d01620008f7565b975062000a5360808d016200090d565b965062000a6360a08d016200090d565b955062000a7360c08d016200090d565b60e08d01519095506001600160401b0381111562000a9057600080fd5b62000a9e8e828f016200086a565b6101008e015190955090506001600160401b0381111562000abe57600080fd5b62000acc8e828f0162000925565b6101208e015190945090506001600160401b0381111562000aec57600080fd5b62000afa8e828f0162000925565b92505062000b0c6101408d016200090d565b90509295989b509295989b9093969950565b600181811c9082168062000b3357607f821691505b60208210810362000b5457634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000ba857600081815260208120601f850160051c8101602086101562000b835750805b601f850160051c820191505b8181101562000ba45782815560010162000b8f565b5050505b505050565b81516001600160401b0381111562000bc95762000bc962000821565b62000be18162000bda845462000b1e565b8462000b5a565b602080601f83116001811462000c19576000841562000c005750858301515b600019600386901b1c1916600185901b17855562000ba4565b600085815260208120601f198616915b8281101562000c4a5788860151825594840194600190910190840162000c29565b508582101562000c695787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b8082018082111562000ca55762000ca562000c79565b92915050565b634e487b7160e01b600052603260045260246000fd5b60006001820162000cd65762000cd662000c79565b5060010190565b615d198062000ced6000396000f3fe60806040526004361061039b5760003560e01c806370a08231116101dc578063b88d4fde11610102578063e985e9c5116100a0578063f07e7fd01161006f578063f07e7fd014610b76578063f2fde38b14610b96578063f4e638be14610bb6578063fbfa77cf14610bde57600080fd5b8063e985e9c514610abb578063e9f585e014610b04578063eb5c60f214610b19578063eee608a414610b4657600080fd5b8063c9200e4f116100dc578063c9200e4f14610a51578063ccdae73a14610a71578063dc78ac1c14610a86578063e8a3d48514610aa657600080fd5b8063b88d4fde146109fc578063b9b8311a14610a1c578063c87b56dd14610a3157600080fd5b80638ef79e911161017a578063a07c7ce411610149578063a07c7ce41461098d578063a22cb465146109a7578063ae474c3c146109c7578063b66a0e5d146109e757600080fd5b80638ef79e911461090b5780639095e8761461092b578063926ce44e1461094b57806395d89b411461097857600080fd5b80638462151c116101b65780638462151c1461088057806384f78930146108ad5780638cba1c67146108cd5780638da5cb5b146108ed57600080fd5b806370a082311461081e578063715018a61461083e5780637f06ee061461085357600080fd5b806333e364cb116102c15780634e99b8001161025f57806363e602301161022e57806363e602301461077d57806365a46e08146107be5780636817031b146107de5780636c19e783146107fe57600080fd5b80634e99b80014610714578063530da8ef1461072957806355367ba9146107485780636352211e1461075d57600080fd5b80633e41dd201161029b5780633e41dd201461069e57806341a5626a146106b357806342842e0e146106d35780634bf365df146106f357600080fd5b806333e364cb146106495780633c352b0d1461065e5780633c38598e1461067e57600080fd5b8063167ddf6e1161033957806323aed2281161030857806323aed228146105d857806323b872dd146105f65780632977e4b3146106165780632f745c591461062957600080fd5b8063167ddf6e1461053b57806318160ddd1461057657806321fe0c641461059a578063238ac933146105ba57600080fd5b8063081812fc11610375578063081812fc146104a3578063095ea7b3146104db578063114ba8ee146104fb5780631623528f1461051b57600080fd5b806301ffc9a71461042c578063031205061461046157806306fdde031461048157600080fd5b3661042757600e546001600160a01b031633146104255760405162461bcd60e51b815260206004820152603c60248201527f466572616c66696c6545786869626974696f6e56343a206f6e6c79206163636560448201527f70742066756e642066726f6d207661756c7420636f6e74726163742e0000000060648201526084015b60405180910390fd5b005b600080fd5b34801561043857600080fd5b5061044c61044736600461467f565b610bfe565b60405190151581526020015b60405180910390f35b34801561046d57600080fd5b5061042561047c3660046146b8565b610c50565b34801561048d57600080fd5b50610496610c79565b6040516104589190614723565b3480156104af57600080fd5b506104c36104be366004614736565b610d0b565b6040516001600160a01b039091168152602001610458565b3480156104e757600080fd5b506104256104f636600461474f565b610d32565b34801561050757600080fd5b506104256105163660046146b8565b610d4b565b34801561052757600080fd5b506104256105363660046146b8565b610d75565b34801561054757600080fd5b5061055b610556366004614736565b610e1e565b60408051825181526020928301519281019290925201610458565b34801561058257600080fd5b5061058c600c5481565b604051908152602001610458565b3480156105a657600080fd5b506104256105b536600461484d565b610e81565b3480156105c657600080fd5b506009546001600160a01b03166104c3565b3480156105e457600080fd5b50600d5462010000900460ff1661044c565b34801561060257600080fd5b50610425610611366004614881565b610f6b565b6104256106243660046148cc565b610fbe565b34801561063557600080fd5b5061058c61064436600461474f565b611910565b34801561065557600080fd5b506104256119ba565b34801561066a57600080fd5b5061042561067936600461497e565b611a7d565b34801561068a57600080fd5b50610425610699366004614a2a565b611c1d565b3480156106aa57600080fd5b50610496611c54565b3480156106bf57600080fd5b506104256106ce36600461497e565b611ce2565b3480156106df57600080fd5b506104256106ee366004614881565b611ec6565b3480156106ff57600080fd5b50600d5461044c906301000000900460ff1681565b34801561072057600080fd5b50610496611f13565b34801561073557600080fd5b50600d5461044c90610100900460ff1681565b34801561075457600080fd5b50610425611f20565b34801561076957600080fd5b506104c3610778366004614736565b611fd4565b34801561078957600080fd5b506104966040518060400160405280601581526020017411995c985b199a5b19515e1a1a589a5d1a5bdb958d605a1b81525081565b3480156107ca57600080fd5b506104256107d9366004614a6b565b612009565b3480156107ea57600080fd5b506104256107f93660046146b8565b612307565b34801561080a57600080fd5b506104256108193660046146b8565b61239d565b34801561082a57600080fd5b5061058c6108393660046146b8565b612428565b34801561084a57600080fd5b506104256124ae565b34801561085f57600080fd5b5061058c61086e366004614736565b60009081526010602052604090205490565b34801561088c57600080fd5b506108a061089b3660046146b8565b6124c2565b6040516104589190614b2c565b3480156108b957600080fd5b506104256108c8366004614b70565b61252e565b3480156108d957600080fd5b506104256108e8366004614be4565b61265e565b3480156108f957600080fd5b506006546001600160a01b03166104c3565b34801561091757600080fd5b50610425610926366004614c9d565b6126a8565b34801561093757600080fd5b50610425610946366004614a2a565b612717565b34801561095757600080fd5b5061058c6109663660046146b8565b60146020526000908152604090205481565b34801561098457600080fd5b5061049661274e565b34801561099957600080fd5b50600d5461044c9060ff1681565b3480156109b357600080fd5b506104256109c2366004614cfe565b61275d565b3480156109d357600080fd5b506104256109e2366004614d35565b612771565b3480156109f357600080fd5b50610425612a7c565b348015610a0857600080fd5b50610425610a17366004614db6565b612a99565b348015610a2857600080fd5b50610425612ae7565b348015610a3d57600080fd5b50610496610a4c366004614736565b612b88565b348015610a5d57600080fd5b50610425610a6c366004614a2a565b612bb9565b348015610a7d57600080fd5b50610496612bf0565b348015610a9257600080fd5b50610425610aa13660046146b8565b612bfd565b348015610ab257600080fd5b50610496612c29565b348015610ac757600080fd5b5061044c610ad6366004614e31565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b348015610b1057600080fd5b50610496612c36565b348015610b2557600080fd5b5061058c610b34366004614736565b6000908152600f602052604090205490565b348015610b5257600080fd5b5061044c610b613660046146b8565b60076020526000908152604090205460ff1681565b348015610b8257600080fd5b506008546104c3906001600160a01b031681565b348015610ba257600080fd5b50610425610bb13660046146b8565b612c43565b348015610bc257600080fd5b50600d546104c39064010000000090046001600160a01b031681565b348015610bea57600080fd5b50600e546104c3906001600160a01b031681565b60006001600160e01b031982166380ac58cd60e01b1480610c2f57506001600160e01b03198216635b5e139f60e01b145b80610c4a57506301ffc9a760e01b6001600160e01b03198316145b92915050565b610c58612cbc565b6001600160a01b03166000908152600760205260409020805460ff19169055565b606060008054610c8890614e64565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb490614e64565b8015610d015780601f10610cd657610100808354040283529160200191610d01565b820191906000526020600020905b815481529060010190602001808311610ce457829003601f168201915b5050505050905090565b6000610d1682612d16565b506000908152600460205260409020546001600160a01b031690565b81610d3c81612d3b565b610d468383612e0d565b505050565b610d53612cbc565b600880546001600160a01b0319166001600160a01b0392909216919091179055565b610d7d612cbc565b6001600160a01b038116610df05760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a20636f737452656365696044820152737665725f206973207a65726f206164647265737360601b606482015260840161041c565b600d80546001600160a01b0390921664010000000002640100000000600160c01b0319909216919091179055565b6040805180820190915260008082526020820152610e3b82612f1d565b610e575760405162461bcd60e51b815260040161041c90614e9e565b50600090815260116020908152604091829020825180840190935280548352600101549082015290565b600d5460ff16610ee85760405162461bcd60e51b815260206004820152602c60248201527f466572616c66696c6545786869626974696f6e56343a20746f6b656e2069732060448201526b6e6f74206275726e61626c6560a01b606482015260840161041c565b60005b8151811015610f6757610f1733838381518110610f0a57610f0a614ed5565b6020026020010151612f3a565b610f335760405162461bcd60e51b815260040161041c90614eeb565b610f55828281518110610f4857610f48614ed5565b6020026020010151612fb9565b80610f5f81614f4e565b915050610eeb565b5050565b826001600160a01b0381163314610f8557610f8533612d3b565b306001600160a01b03841603610fad5760405162461bcd60e51b815260040161041c90614f67565b610fb884848461308f565b50505050565b600d5462010000900460ff166110295760405162461bcd60e51b815260206004820152602a60248201527f466572616c66696c6545786869626974696f6e56343a2073616c65206973206e6044820152691bdd081cdd185c9d195960b21b606482015260840161041c565b6110316130c0565b61103a8161313b565b61104a60e0820160c08301614fc4565b6110b957803534146110b45760405162461bcd60e51b815260206004820152602d60248201527f466572616c66696c6545786869626974696f6e56343a20696e76616c6964207060448201526c185e5b595b9d08185b5bdd5b9d609a1b606482015260840161041c565b611122565b600e54604051632eeee16360e01b81526001600160a01b0390911690632eeee163906110ef9087908790879087906004016151c9565b600060405180830381600087803b15801561110957600080fd5b505af115801561111d573d6000803e3d6000fd5b505050505b6000463083604051602001611139939291906151fb565b60405160208183030381529060405280519060200120905061115d81868686613294565b61117a57604051638baa579f60e01b815260040160405180910390fd5b60006020830135833511156111b457611196608084018461522e565b90506111a760208501358535615277565b6111b1919061528a565b90505b6000806000805b6111c8608088018861522e565b90508110156116a15761121e306111e560808a0160608b016146b8565b6111f260808b018b61522e565b8581811061120257611202614ed5565b90506020020135604051806020016040528060008152506132ec565b600061122d60a089018961522e565b8381811061123d5761123d614ed5565b905060200281019061124f91906152ac565b808060200260200160405190810160405280939291908181526020016000905b8282101561129b5761128c604083028601368190038101906152f5565b8152602001906001019061126f565b50505050509050600086905060005b8251811080156112ba5750600082115b1561139f576000601460008584815181106112d7576112d7614ed5565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000205490506000838210156113165781611318565b835b9050611324818961534b565b9750806014600087868151811061133d5761133d614ed5565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002060008282546113789190615277565b9091555061138890508185615277565b93505050808061139790614f4e565b9150506112aa565b50801561149b5760005b82518110156114995760008382815181106113c6576113c6614ed5565b602002602001015160000151905060006127108584815181106113eb576113eb614ed5565b60200260200101516020015185611402919061535e565b61140c919061528a565b600d549091506001600160a01b0364010000000090910481169083160361144057611437818961534b565b97505050611487565b61144a818a61534b565b6040519099506001600160a01b0383169082156108fc029083906000818181858888f19350505050158015611483573d6000803e3d6000fd5b5050505b8061149181614f4e565b9150506113a9565b505b60006016816114ad60808d018d61522e565b878181106114bd576114bd614ed5565b9050602002013581526020019081526020016000206040518060600160405290816000820180546114ed90614e64565b80601f016020809104026020016040519081016040528092919081815260200182805461151990614e64565b80156115665780601f1061153b57610100808354040283529160200191611566565b820191906000526020600020905b81548152906001019060200180831161154957829003601f168201915b5050505050815260200160018201805461157f90614e64565b80601f01602080910402602001604051908101604052809291908181526020018280546115ab90614e64565b80156115f85780601f106115cd576101008083540402835291602001916115f8565b820191906000526020600020905b8154815290600101906020018083116115db57829003601f168201915b50505050508152602001600282015481525050905080604001518561161d919061534b565b945061162c60808b018b61522e565b8581811061163c5761163c614ed5565b905060200201358a606001602081019061165691906146b8565b6001600160a01b03167f0475389cd69b8d3163620b43283bf74e8fc71020c3c6cef2a529b5c405e9687f60405160405180910390a3505050808061169990614f4e565b9150506111bb565b5080156118305760006116b38261331f565b6015546040516370a0823160e01b815230600482015291925082916001600160a01b03909116906370a0823190602401602060405180830381865afa158015611700573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117249190615375565b111561182e576015546000906001600160a01b031663a9059cbb61174e60808b0160608c016146b8565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018590526044016020604051808303816000875af115801561179b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117bf919061538e565b90508061182c5760405162461bcd60e51b815260206004820152603560248201527f466572616c66696c6545786869626974696f6e56343a20466572616c66696c65604482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b606482015260840161041c565b505b505b61183a828461534b565b61184960208801358835615277565b10156118ac5760405162461bcd60e51b815260206004820152602c60248201527f466572616c66696c6545786869626974696f6e56343a20746f74616c2062707360448201526b0206f7665722031302c3030360a41b606482015260840161041c565b60006118b9848835615277565b9050801561190457600d546040516401000000009091046001600160a01b0316906108fc8315029083906000818181858888f19350505050158015611902573d6000803e3d6000fd5b505b50505050505050505050565b600061191b83612428565b821061197d5760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b606482015260840161041c565b6001600160a01b03831660009081526012602052604090208054839081106119a7576119a7614ed5565b9060005260206000200154905092915050565b6119c2612cbc565b600d546301000000900460ff16156119ec5760405162461bcd60e51b815260040161041c906153ab565b600d5462010000900460ff1615611a625760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a205f73656c6c696e6720604482015273726571756972656420746f2062652066616c736560601b606482015260840161041c565b611a6a6130c0565b600d805462ff0000191662010000179055565b611a85612cbc565b828114611aa5576040516313086eff60e21b815260040160405180910390fd5b60005b83811015611c16576000858583818110611ac457611ac4614ed5565b9050602002016020810190611ad991906146b8565b6001600160a01b031603611b0057604051630107349760e51b815260040160405180910390fd5b828282818110611b1257611b12614ed5565b90506020020135600003611b3957604051636745f8fb60e01b815260040160405180910390fd5b600060146000878785818110611b5157611b51614ed5565b9050602002016020810190611b6691906146b8565b6001600160a01b03166001600160a01b03168152602001908152602001600020541115611ba6576040516328547bdf60e01b815260040160405180910390fd5b828282818110611bb857611bb8614ed5565b9050602002013560146000878785818110611bd557611bd5614ed5565b9050602002016020810190611bea91906146b8565b6001600160a01b0316815260208101919091526040016000205580611c0e81614f4e565b915050611aa8565b5050505050565b611c25612cbc565b6000819003611c475760405163683d806b60e11b815260040160405180910390fd5b601a610d4682848361544d565b60198054611c6190614e64565b80601f0160208091040260200160405190810160405280929190818152602001828054611c8d90614e64565b8015611cda5780601f10611caf57610100808354040283529160200191611cda565b820191906000526020600020905b815481529060010190602001808311611cbd57829003601f168201915b505050505081565b611cea612cbc565b828114611d0a576040516313086eff60e21b815260040160405180910390fd5b60005b83811015611c16576000838383818110611d2957611d29614ed5565b9050602002016020810190611d3e91906146b8565b6001600160a01b031603611d6557604051630107349760e51b815260040160405180910390fd5b600060146000858585818110611d7d57611d7d614ed5565b9050602002016020810190611d9291906146b8565b6001600160a01b03166001600160a01b03168152602001908152602001600020541115611dd2576040516328547bdf60e01b815260040160405180910390fd5b60146000868684818110611de857611de8614ed5565b9050602002016020810190611dfd91906146b8565b6001600160a01b03166001600160a01b031681526020019081526020016000205460146000858585818110611e3457611e34614ed5565b9050602002016020810190611e4991906146b8565b6001600160a01b03166001600160a01b031681526020019081526020016000208190555060146000868684818110611e8357611e83614ed5565b9050602002016020810190611e9891906146b8565b6001600160a01b03168152602081019190915260400160009081205580611ebe81614f4e565b915050611d0d565b826001600160a01b0381163314611ee057611ee033612d3b565b306001600160a01b03841603611f085760405162461bcd60e51b815260040161041c90614f67565b610fb88484846133a2565b600a8054611c6190614e64565b611f28612cbc565b600d546301000000900460ff1615611f525760405162461bcd60e51b815260040161041c906153ab565b600d5462010000900460ff16611fc65760405162461bcd60e51b815260206004820152603360248201527f466572616c66696c6545786869626974696f6e56343a205f73656c6c696e6720604482015272726571756972656420746f206265207472756560681b606482015260840161041c565b600d805462ff000019169055565b6000818152600260205260408120546001600160a01b031680610c4a5760405162461bcd60e51b815260040161041c90614e9e565b612011612cbc565b60008251118015612023575060008151115b6120a35760405162461bcd60e51b815260206004820152604560248201527f466572616c66696c6545786869626974696f6e56343a2073657269657349647360448201527f206f7220726563697069656e74416464726573736573206c656e677468206973606482015264207a65726f60d81b608482015260a40161041c565b805182511461212f5760405162461bcd60e51b815260206004820152604c60248201527f466572616c66696c6545786869626974696f6e56343a2073657269657349647360448201527f206c656e67746820697320646966666572656e742066726f6d2072656369706960648201526b656e7441646472657373657360a01b608482015260a40161041c565b612137611f20565b3060008181526012602090815260408083208054825181850281018501909352808352919290919083018282801561218e57602002820191906000526020600020905b81548152602001906001019080831161217a575b5050505050905060005b815181101561228a5760008282815181106121b5576121b5614ed5565b602090810291909101810151600081815260118352604080822081518083019092528054825260010154938101939093529092505b87518161ffff16101561227457878161ffff168151811061220d5761220d614ed5565b6020026020010151826000015103612262576000878261ffff168151811061223757612237614ed5565b6020026020010151905061225c878286604051806020016040528060008152506132ec565b50612274565b8061226c8161550c565b9150506121ea565b505050808061228290614f4e565b915050612198565b5061229482612428565b15610fb85760405162461bcd60e51b815260206004820152603c60248201527f466572616c66696c6545786869626974696f6e56343a20546f6b656e20666f7260448201527f2073616c652062616c616e63652068617320746f206265207a65726f00000000606482015260840161041c565b61230f612cbc565b6001600160a01b03811661237b5760405162461bcd60e51b815260206004820152602d60248201527f466572616c66696c6545786869626974696f6e56343a207661756c745f20697360448201526c207a65726f206164647265737360981b606482015260840161041c565b600e80546001600160a01b0319166001600160a01b0392909216919091179055565b6123a5612cbc565b6001600160a01b0381166124065760405162461bcd60e51b815260206004820152602260248201527f45434453415369676e3a207369676e65725f206973207a65726f206164647265604482015261737360f01b606482015260840161041c565b600980546001600160a01b0319166001600160a01b0392909216919091179055565b60006001600160a01b0382166124925760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b606482015260840161041c565b506001600160a01b031660009081526003602052604090205490565b6124b6612cbc565b6124c060006133bd565b565b6001600160a01b03811660009081526012602090815260409182902080548351818402810184019094528084526060939283018282801561252257602002820191906000526020600020905b81548152602001906001019080831161250e575b50505050509050919050565b3360009081526007602052604090205460ff168061255657506006546001600160a01b031633145b61255f57600080fd5b600d546301000000900460ff166125895760405163447691f760e01b815260040160405180910390fd5b60005b81811015610d46578282828181106125a6576125a6614ed5565b90506080020160600135601760008585858181106125c6576125c6614ed5565b9050608002016020013581526020019081526020016000208190555061264c8383838181106125f7576125f7614ed5565b9050608002016000013584848481811061261357612613614ed5565b9050608002016020013585858581811061262f5761262f614ed5565b905060800201604001602081019061264791906146b8565b61340f565b8061265681614f4e565b91505061258c565b3360009081526007602052604090205460ff168061268657506006546001600160a01b031633145b61268f57600080fd5b6040516369bd111d60e11b815260040160405180910390fd5b6126b0612cbc565b600081511161270b5760405162461bcd60e51b815260206004820152602160248201527f4552433732314d657461646174613a20626173655552495f20697320656d70746044820152607960f81b606482015260840161041c565b600a610f67828261552d565b61271f612cbc565b60008190036127415760405163683d806b60e11b815260040160405180910390fd5b6018610d4682848361544d565b606060018054610c8890614e64565b8161276781612d3b565b610d468383613591565b3360009081526007602052604090205460ff168061279957506006546001600160a01b031633145b6127a257600080fd5b6127ab86612f1d565b6127c8576040516352a7a53160e11b815260040160405180910390fd5b6000868152601660205260408082208151606081019092528054829082906127ef90614e64565b80601f016020809104026020016040519081016040528092919081815260200182805461281b90614e64565b80156128685780601f1061283d57610100808354040283529160200191612868565b820191906000526020600020905b81548152906001019060200180831161284b57829003601f168201915b5050505050815260200160018201805461288190614e64565b80601f01602080910402602001604051908101604052809291908181526020018280546128ad90614e64565b80156128fa5780601f106128cf576101008083540402835291602001916128fa565b820191906000526020600020905b8154815290600101906020018083116128dd57829003601f168201915b505050505081526020016002820154815250509050604051806060016040528087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f88018190048102820181019092528681529181019190879087908190840183828082843760009201829052509385525050506020918201859052898152601690915260409020815181906129af908261552d565b50602082015160018201906129c4908261552d565b506040918201516002909101558101516000906129e19084615277565b90508015612a725760006129f489611fd4565b6015549091506001600160a01b03166340c10f1982612a128561331f565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401600060405180830381600087803b158015612a5857600080fd5b505af1158015612a6c573d6000803e3d6000fd5b50505050505b5050505050505050565b612a84612cbc565b600d805463ff000000191690556124c06119ba565b836001600160a01b0381163314612ab357612ab333612d3b565b306001600160a01b03851603612adb5760405162461bcd60e51b815260040161041c90614f67565b611c168585858561359c565b612aef612cbc565b612af7611f20565b30600090815260126020908152604080832080548251818502810185019093528083529192909190830182828015612b4e57602002820191906000526020600020905b815481526020019060010190808311612b3a575b5050505050905060005b8151811015610f6757612b76828281518110610f4857610f48614ed5565b80612b8081614f4e565b915050612b58565b6060612b9382612f1d565b612bb0576040516352a7a53160e11b815260040160405180910390fd5b610c4a826135ce565b612bc1612cbc565b6000819003612be35760405163683d806b60e11b815260040160405180910390fd5b6019610d4682848361544d565b60188054611c6190614e64565b612c05612cbc565b6001600160a01b03166000908152600760205260409020805460ff19166001179055565b600b8054611c6190614e64565b601a8054611c6190614e64565b612c4b612cbc565b6001600160a01b038116612cb05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161041c565b612cb9816133bd565b50565b6006546001600160a01b031633146124c05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161041c565b612d1f81612f1d565b612cb95760405162461bcd60e51b815260040161041c90614e9e565b6008546001600160a01b03163b15612cb957600854604051633185c44d60e21b81523060048201526001600160a01b0383811660248301529091169063c617113490604401602060405180830381865afa158015612d9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dc1919061538e565b612cb95760405162461bcd60e51b815260206004820152601760248201527f6f70657261746f72206973206e6f7420616c6c6f776564000000000000000000604482015260640161041c565b6000612e1882611fd4565b9050806001600160a01b0316836001600160a01b031603612e855760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b606482015260840161041c565b336001600160a01b0382161480612ea15750612ea18133610ad6565b612f135760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c000000606482015260840161041c565b610d4683836137e2565b6000908152600260205260409020546001600160a01b0316151590565b600080612f4683611fd4565b9050806001600160a01b0316846001600160a01b03161480612f8d57506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b80612fb15750836001600160a01b0316612fa684610d0b565b6001600160a01b0316145b949350505050565b612fc281612f1d565b612fde5760405162461bcd60e51b815260040161041c90614e9e565b60008181526011602090815260408083208151808301835281548082526001928301548286015285526010909352908320805492939192909190613023908490615277565b925050819055506001600c600082825461303d9190615277565b909155505060008281526011602052604081208181556001015561306082613850565b60405182907fbde7938970372996ff103863625e348ef2bf8f38a5b02181be75aafef17c23d590600090a25050565b6130993382612f3a565b6130b55760405162461bcd60e51b815260040161041c90614eeb565b610d468383836138f3565b60006130cb30612428565b905060008111612cb95760405162461bcd60e51b815260206004820152603560248201527f466572616c66696c6545786869626974696f6e56343a204e6f20746f6b656e206044820152741bdddb995908189e481d1a194818dbdb9d1c9858dd605a1b606482015260840161041c565b600061314a608083018361522e565b9050116131a55760405162461bcd60e51b8152602060048201526024808201527f466572616c66696c6553616c65446174613a20746f6b656e49647320697320656044820152636d70747960e01b606482015260840161041c565b6131b260a082018261522e565b90506131c1608083018361522e565b9050146132365760405162461bcd60e51b815260206004820152603d60248201527f466572616c66696c6553616c65446174613a20746f6b656e49647320616e642060448201527f726576656e7565536861726573206c656e677468206d69736d61746368000000606482015260840161041c565b42816040013511612cb95760405162461bcd60e51b815260206004820152602260248201527f466572616c66696c6553616c65446174613a2073616c65206973206578706972604482015261195960f21b606482015260840161041c565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c859052603c812081906132d190848787613a64565b6009546001600160a01b039081169116149695505050505050565b6132f78484846138f3565b61330384848484613a8c565b610fb85760405162461bcd60e51b815260040161041c906155ec565b6015546040805163313ce56760e01b815290516000926001600160a01b03169163313ce5679160048083019260209291908290030181865afa158015613369573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061338d919061563e565b61339890600a61573f565b610c4a908361535e565b610d4683838360405180602001604052806000815250612a99565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000838152600f6020526040902054151561342984613b8a565b604051602001613439919061576a565b604051602081830303815290604052906134665760405162461bcd60e51b815260040161041c9190614723565b506000838152600f6020908152604080832054601090925290912054106134e15760405162461bcd60e51b815260206004820152602960248201527f466572616c66696c6545786869626974696f6e56343a206e6f20736c6f747320604482015268617661696c61626c6560b81b606482015260840161041c565b6001600c60008282546134f4919061534b565b9091555050600083815260106020526040812080546001929061351890849061534b565b90915550506040805180820182528481526020808201858152600086815260119092529290209051815590516001909101556135548183613c1d565b8183826001600160a01b03167f407d7da1d3b2b1871fbfa2b5b1c4657a3cc5711d3023c552798551c7ee301eea60405160405180910390a4505050565b610f67338383613d98565b6135a63383612f3a565b6135c25760405162461bcd60e51b815260040161041c90614eeb565b610fb8848484846132ec565b606060006016600084815260200190815260200160002060405180606001604052908160008201805461360090614e64565b80601f016020809104026020016040519081016040528092919081815260200182805461362c90614e64565b80156136795780601f1061364e57610100808354040283529160200191613679565b820191906000526020600020905b81548152906001019060200180831161365c57829003601f168201915b5050505050815260200160018201805461369290614e64565b80601f01602080910402602001604051908101604052809291908181526020018280546136be90614e64565b801561370b5780601f106136e05761010080835404028352916020019161370b565b820191906000526020600020905b8154815290600101906020018083116136ee57829003601f168201915b50505091835250506002919091015460209182015260008581526017909152604081205491925061373a610c79565b61374d61374884600161534b565b613b8a565b60405160200161375e9291906157c7565b604051602081830303815290604052905060008161377f8560000151613e66565b61378c8660200151613efd565b60405160200161379e93929190615804565b60405160208183030381529060405290506137b881613fd7565b6040516020016137c891906158d4565b604051602081830303815290604052945050505050919050565b600081815260046020526040902080546001600160a01b0319166001600160a01b038416908117909155819061381782611fd4565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600061385b82611fd4565b905061386b816000846001614129565b61387482611fd4565b600083815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526003845282852080546000190190558785526002909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b826001600160a01b031661390682611fd4565b6001600160a01b03161461392c5760405162461bcd60e51b815260040161041c90615919565b6001600160a01b03821661398e5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161041c565b61399b8383836001614129565b826001600160a01b03166139ae82611fd4565b6001600160a01b0316146139d45760405162461bcd60e51b815260040161041c90615919565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6000806000613a7587878787614240565b91509150613a8281614304565b5095945050505050565b60006001600160a01b0384163b15613b8257604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290613ad090339089908890889060040161595e565b6020604051808303816000875af1925050508015613b0b575060408051601f3d908101601f19168201909252613b0891810190615991565b60015b613b68573d808015613b39576040519150601f19603f3d011682016040523d82523d6000602084013e613b3e565b606091505b508051600003613b605760405162461bcd60e51b815260040161041c906155ec565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050612fb1565b506001612fb1565b60606000613b978361444e565b60010190506000816001600160401b03811115613bb657613bb6614779565b6040519080825280601f01601f191660200182016040528015613be0576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084613bea575b509392505050565b6001600160a01b038216613c735760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015260640161041c565b613c7c81612f1d565b15613cc95760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161041c565b613cd7600083836001614129565b613ce081612f1d565b15613d2d5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161041c565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b816001600160a01b0316836001600160a01b031603613df95760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015260640161041c565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b60608151600003613ef95760198054613e7e90614e64565b80601f0160208091040260200160405190810160405280929190818152602001828054613eaa90614e64565b80156125225780601f10613ecc57610100808354040283529160200191612522565b820191906000526020600020905b815481529060010190602001808311613eda5750939695505050505050565b5090565b60608151600003613f155760188054613e7e90614e64565b613fb1613f2183613fd7565b601a8054613f2e90614e64565b80601f0160208091040260200160405190810160405280929190818152602001828054613f5a90614e64565b8015613fa75780601f10613f7c57610100808354040283529160200191613fa7565b820191906000526020600020905b815481529060010190602001808311613f8a57829003601f168201915b5050505050614526565b604051602001613fc191906159ae565b6040516020818303038152906040529050919050565b60608151600003613ff657505060408051602081019091526000815290565b6000604051806060016040528060408152602001615ca46040913990506000600384516002614025919061534b565b61402f919061528a565b61403a90600461535e565b6001600160401b0381111561405157614051614779565b6040519080825280601f01601f19166020018201604052801561407b576020820181803683370190505b509050600182016020820185865187015b808210156140e7576003820191508151603f8160121c168501518453600184019350603f81600c1c168501518453600184019350603f8160061c168501518453600184019350603f811685015184535060018301925061408c565b505060038651066001811461410357600281146141165761411e565b603d6001830353603d600283035361411e565b603d60018303535b509195945050505050565b60018111156141985760405162461bcd60e51b815260206004820152603560248201527f455243373231456e756d657261626c653a20636f6e7365637574697665207472604482015274185b9cd9995c9cc81b9bdd081cdd5c1c1bdc9d1959605a1b606482015260840161041c565b816001600160a01b038516158015906141c35750836001600160a01b0316856001600160a01b031614155b156141d2576141d28582614559565b6001600160a01b038416158015906141fc5750846001600160a01b0316846001600160a01b031614155b15611c16576001600160a01b038416600090815260126020908152604080832080546001810182559084528284208101859055848452601390925290912055611c16565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561427757506000905060036142fb565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156142cb573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166142f4576000600192509250506142fb565b9150600090505b94509492505050565b6000816004811115614318576143186159ec565b036143205750565b6001816004811115614334576143346159ec565b036143815760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161041c565b6002816004811115614395576143956159ec565b036143e25760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161041c565b60038160048111156143f6576143f66159ec565b03612cb95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161041c565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b831061448d5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef810000000083106144b9576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106144d757662386f26fc10000830492506010015b6305f5e10083106144ef576305f5e100830492506008015b612710831061450357612710830492506004015b60648310614515576064830492506002015b600a8310610c4a5760010192915050565b6060614552838360405160200161453e929190615a02565b604051602081830303815290604052613fd7565b9392505050565b6000600161456684612428565b6145709190615277565b600083815260136020526040902054909150808214614617576001600160a01b03841660009081526012602052604081208054849081106145b3576145b3614ed5565b906000526020600020015490508060126000876001600160a01b03166001600160a01b0316815260200190815260200160002083815481106145f7576145f7614ed5565b600091825260208083209091019290925591825260139052604090208190555b60008381526013602090815260408083208390556001600160a01b03871683526012909152902080548061464d5761464d615c8d565b6001900381819060005260206000200160009055905550505050565b6001600160e01b031981168114612cb957600080fd5b60006020828403121561469157600080fd5b813561455281614669565b80356001600160a01b03811681146146b357600080fd5b919050565b6000602082840312156146ca57600080fd5b6145528261469c565b60005b838110156146ee5781810151838201526020016146d6565b50506000910152565b6000815180845261470f8160208601602086016146d3565b601f01601f19169290920160200192915050565b60208152600061455260208301846146f7565b60006020828403121561474857600080fd5b5035919050565b6000806040838503121561476257600080fd5b61476b8361469c565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156147b7576147b7614779565b604052919050565b60006001600160401b038211156147d8576147d8614779565b5060051b60200190565b600082601f8301126147f357600080fd5b81356020614808614803836147bf565b61478f565b82815260059290921b8401810191818101908684111561482757600080fd5b8286015b84811015614842578035835291830191830161482b565b509695505050505050565b60006020828403121561485f57600080fd5b81356001600160401b0381111561487557600080fd5b612fb1848285016147e2565b60008060006060848603121561489657600080fd5b61489f8461469c565b92506148ad6020850161469c565b9150604084013590509250925092565b60ff81168114612cb957600080fd5b600080600080608085870312156148e257600080fd5b843593506020850135925060408501356148fb816148bd565b915060608501356001600160401b0381111561491657600080fd5b850160e0818803121561492857600080fd5b939692955090935050565b60008083601f84011261494557600080fd5b5081356001600160401b0381111561495c57600080fd5b6020830191508360208260051b850101111561497757600080fd5b9250929050565b6000806000806040858703121561499457600080fd5b84356001600160401b03808211156149ab57600080fd5b6149b788838901614933565b909650945060208701359150808211156149d057600080fd5b506149dd87828801614933565b95989497509550505050565b60008083601f8401126149fb57600080fd5b5081356001600160401b03811115614a1257600080fd5b60208301915083602082850101111561497757600080fd5b60008060208385031215614a3d57600080fd5b82356001600160401b03811115614a5357600080fd5b614a5f858286016149e9565b90969095509350505050565b60008060408385031215614a7e57600080fd5b82356001600160401b0380821115614a9557600080fd5b614aa1868387016147e2565b9350602091508185013581811115614ab857600080fd5b85019050601f81018613614acb57600080fd5b8035614ad9614803826147bf565b81815260059190911b82018301908381019088831115614af857600080fd5b928401925b82841015614b1d57614b0e8461469c565b82529284019290840190614afd565b80955050505050509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614b6457835183529284019291840191600101614b48565b50909695505050505050565b60008060208385031215614b8357600080fd5b82356001600160401b0380821115614b9a57600080fd5b818501915085601f830112614bae57600080fd5b813581811115614bbd57600080fd5b8660208260071b8501011115614bd257600080fd5b60209290920196919550909350505050565b60008060208385031215614bf757600080fd5b82356001600160401b0380821115614c0e57600080fd5b818501915085601f830112614c2257600080fd5b813581811115614c3157600080fd5b866020606083028501011115614bd257600080fd5b60006001600160401b03831115614c5f57614c5f614779565b614c72601f8401601f191660200161478f565b9050828152838383011115614c8657600080fd5b828260208301376000602084830101529392505050565b600060208284031215614caf57600080fd5b81356001600160401b03811115614cc557600080fd5b8201601f81018413614cd657600080fd5b612fb184823560208401614c46565b8015158114612cb957600080fd5b80356146b381614ce5565b60008060408385031215614d1157600080fd5b614d1a8361469c565b91506020830135614d2a81614ce5565b809150509250929050565b60008060008060008060808789031215614d4e57600080fd5b8635955060208701356001600160401b0380821115614d6c57600080fd5b614d788a838b016149e9565b90975095506040890135915080821115614d9157600080fd5b50614d9e89828a016149e9565b979a9699509497949695606090950135949350505050565b60008060008060808587031215614dcc57600080fd5b614dd58561469c565b9350614de36020860161469c565b92506040850135915060608501356001600160401b03811115614e0557600080fd5b8501601f81018713614e1657600080fd5b614e2587823560208401614c46565b91505092959194509250565b60008060408385031215614e4457600080fd5b614e4d8361469c565b9150614e5b6020840161469c565b90509250929050565b600181811c90821680614e7857607f821691505b602082108103614e9857634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526018908201527f4552433732313a20696e76616c696420746f6b656e2049440000000000000000604082015260600190565b634e487b7160e01b600052603260045260246000fd5b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600060018201614f6057614f60614f38565b5060010190565b6020808252603e908201527f466572616c66696c6545786869626974696f6e56343a20436f6e74726163742060408201527f69736e277420616c6c6f77656420746f207265636569766520746f6b656e0000606082015260800190565b600060208284031215614fd657600080fd5b813561455281614ce5565b6000808335601e19843603018112614ff857600080fd5b83016020810192503590506001600160401b0381111561501757600080fd5b8060051b360382131561497757600080fd5b8183526000602080850194508260005b8581101561506f576001600160a01b036150528361469c565b168752818301358388015260409687019690910190600101615039565b509495945050505050565b81835260006020808501808196508560051b810191508460005b878110156151015782840389528135601e198836030181126150b557600080fd5b870185810190356001600160401b038111156150d057600080fd5b8060061b36038213156150e257600080fd5b6150ed868284615029565b9a87019a9550505090840190600101615094565b5091979650505050505050565b8035825260208082013590830152604080820135908301526001600160a01b0361513a6060830161469c565b166060830152600061514f6080830183614fe1565b60e06080860181905285018190526101006001600160fb1b0382111561517457600080fd5b8160051b9150818382880137818601925061519260a0860186614fe1565b9250818785030160a08801526151ab828501848361507a565b93505050506151bc60c08401614cf3565b80151560c0860152613c15565b84815283602082015260ff831660408201526080606082015260006151f1608083018461510e565b9695505050505050565b8381526001600160a01b03831660208201526060604082018190526000906152259083018461510e565b95945050505050565b6000808335601e1984360301811261524557600080fd5b8301803591506001600160401b0382111561525f57600080fd5b6020019150600581901b360382131561497757600080fd5b81810381811115610c4a57610c4a614f38565b6000826152a757634e487b7160e01b600052601260045260246000fd5b500490565b6000808335601e198436030181126152c357600080fd5b8301803591506001600160401b038211156152dd57600080fd5b6020019150600681901b360382131561497757600080fd5b60006040828403121561530757600080fd5b604051604081018181106001600160401b038211171561532957615329614779565b6040526153358361469c565b8152602083013560208201528091505092915050565b80820180821115610c4a57610c4a614f38565b8082028115828204841417610c4a57610c4a614f38565b60006020828403121561538757600080fd5b5051919050565b6000602082840312156153a057600080fd5b815161455281614ce5565b60208082526034908201527f466572616c66696c6545786869626974696f6e56343a206d696e7461626c6520604082015273726571756972656420746f2062652066616c736560601b606082015260800190565b601f821115610d4657600081815260208120601f850160051c810160208610156154265750805b601f850160051c820191505b8181101561544557828155600101615432565b505050505050565b6001600160401b0383111561546457615464614779565b615478836154728354614e64565b836153ff565b6000601f8411600181146154ac57600085156154945750838201355b600019600387901b1c1916600186901b178355611c16565b600083815260209020601f19861690835b828110156154dd57868501358255602094850194600190920191016154bd565b50868210156154fa5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b600061ffff80831681810361552357615523614f38565b6001019392505050565b81516001600160401b0381111561554657615546614779565b61555a816155548454614e64565b846153ff565b602080601f83116001811461558f57600084156155775750858301515b600019600386901b1c1916600185901b178555615445565b600085815260208120601f198616915b828110156155be5788860151825594840194600190910190840161559f565b50858210156155dc5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b60006020828403121561565057600080fd5b8151614552816148bd565b600181815b8085111561569657816000190482111561567c5761567c614f38565b8085161561568957918102915b93841c9390800290615660565b509250929050565b6000826156ad57506001610c4a565b816156ba57506000610c4a565b81600181146156d057600281146156da576156f6565b6001915050610c4a565b60ff8411156156eb576156eb614f38565b50506001821b610c4a565b5060208310610133831016604e8410600b8410161715615719575081810a610c4a565b615723838361565b565b806000190482111561573757615737614f38565b029392505050565b600061455260ff84168361569e565b600081516157608185602086016146d3565b9290920192915050565b7f466572616c66696c6545786869626974696f6e56343a2073657269657349642081526e03237b2b9b713ba1032bc34b9ba1d1608d1b6020820152600082516157ba81602f8501602087016146d3565b91909101602f0192915050565b600083516157d98184602088016146d3565b61202360f01b90830190815283516157f88160028401602088016146d3565b01600201949350505050565b683d913730b6b2911d1160b91b815283516000906158298160098501602089016146d3565b7f222c202265787465726e616c5f75726c223a2268747470733a2f2f666572616c600991840191820152733334b6329731b7b69116101134b6b0b3b2911d1160611b6029820152845161588381603d8401602089016146d3565b731116101130b734b6b0ba34b7b72fbab936111d1160611b603d929091019182015283516158b88160518401602088016146d3565b61227d60f01b6051929091019182015260530195945050505050565b7f646174613a6170706c69636174696f6e2f6a736f6e3b6261736536342c00000081526000825161590c81601d8501602087016146d3565b91909101601d0192915050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906151f1908301846146f7565b6000602082840312156159a357600080fd5b815161455281614669565b7519185d184e9d195e1d0bda1d1b5b0ed8985cd94d8d0b60521b8152600082516159df8160168501602087016146d3565b9190910160160192915050565b634e487b7160e01b600052602160045260246000fd5b7f3c21444f43545950452068746d6c3e3c68746d6c206c616e673d22656e223e3c81527f686561643e3c6d65746120636861727365743d225554462d38223e000000000060208201527f3c7363726970743e7661722064656661756c74417274776f726b446174613d22603b82015260008351615a8681605b8501602088016146d3565b7f223b6c657420616c6c6f774f726967696e73203d207b2268747470733a2f2f66605b918401918201527f6572616c66696c652e636f6d223a2021307d3b66756e6374696f6e20696e6974607b8201527f4461746128297b646f63756d656e742e676574456c656d656e74427949642822609b8201527f6d61696e6672616d6522292e636f6e74656e7457696e646f772e706f73744d6560bb8201527f73736167652864656661756c74417274776f726b446174612c20222a22293b7d60db820152681e17b9b1b934b83a1f60b91b60fb8201527f3c2f686561643e3c626f6479207374796c653d226f766572666c6f772d783a686101048201527f696464656e3b70616464696e673a303b6d617267696e3a303b77696474683a206101248201527f313030253b22206f6e6c6f61643d22696e6974446174612829223e3c696672616101448201527f6d652069643d226d61696e6672616d6522207374796c653d22646973706c61796101648201527f3a626c6f636b3b70616464696e673a303b6d617267696e3a303b626f726465726101848201527f3a6e6f6e653b77696474683a313030253b6865696768743a31303076683b22206101a48201526439b9319e9160d91b6101c4820152615225615c646101c983018661574e565b7f223e3c2f696672616d653e3c2f626f64793e3c2f68746d6c3e00000000000000815260190190565b634e487b7160e01b600052603160045260246000fdfe4142434445464748494a4b4c4d4e4f505152535455565758595a6162636465666768696a6b6c6d6e6f707172737475767778797a303132333435363738392b2fa2646970667358221220d29ff0bf949597685d919529cf38a5e92442389a9f4fd9f31230bb706c72d9ac64736f6c63430008110033",
}

// FeralfileExhibitionV42ABI is the input ABI used to generate the binding from.
// Deprecated: Use FeralfileExhibitionV42MetaData.ABI instead.
var FeralfileExhibitionV42ABI = FeralfileExhibitionV42MetaData.ABI

// FeralfileExhibitionV42Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeralfileExhibitionV42MetaData.Bin instead.
var FeralfileExhibitionV42Bin = FeralfileExhibitionV42MetaData.Bin

// DeployFeralfileExhibitionV42 deploys a new Ethereum contract, binding an instance of FeralfileExhibitionV42 to it.
func DeployFeralfileExhibitionV42(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, burnable_ bool, bridgeable_ bool, signer_ common.Address, vault_ common.Address, costReceiver_ common.Address, contractURI_ string, seriesIds_ []*big.Int, seriesMaxSupplies_ []*big.Int, erc20ContractAddress_ common.Address) (common.Address, *types.Transaction, *FeralfileExhibitionV42, error) {
	parsed, err := FeralfileExhibitionV42MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeralfileExhibitionV42Bin), backend, name_, symbol_, burnable_, bridgeable_, signer_, vault_, costReceiver_, contractURI_, seriesIds_, seriesMaxSupplies_, erc20ContractAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeralfileExhibitionV42{FeralfileExhibitionV42Caller: FeralfileExhibitionV42Caller{contract: contract}, FeralfileExhibitionV42Transactor: FeralfileExhibitionV42Transactor{contract: contract}, FeralfileExhibitionV42Filterer: FeralfileExhibitionV42Filterer{contract: contract}}, nil
}

// FeralfileExhibitionV42 is an auto generated Go binding around an Ethereum contract.
type FeralfileExhibitionV42 struct {
	FeralfileExhibitionV42Caller     // Read-only binding to the contract
	FeralfileExhibitionV42Transactor // Write-only binding to the contract
	FeralfileExhibitionV42Filterer   // Log filterer for contract events
}

// FeralfileExhibitionV42Caller is an auto generated read-only Go binding around an Ethereum contract.
type FeralfileExhibitionV42Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV42Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FeralfileExhibitionV42Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV42Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeralfileExhibitionV42Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV42Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeralfileExhibitionV42Session struct {
	Contract     *FeralfileExhibitionV42 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FeralfileExhibitionV42CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeralfileExhibitionV42CallerSession struct {
	Contract *FeralfileExhibitionV42Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// FeralfileExhibitionV42TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeralfileExhibitionV42TransactorSession struct {
	Contract     *FeralfileExhibitionV42Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// FeralfileExhibitionV42Raw is an auto generated low-level Go binding around an Ethereum contract.
type FeralfileExhibitionV42Raw struct {
	Contract *FeralfileExhibitionV42 // Generic contract binding to access the raw methods on
}

// FeralfileExhibitionV42CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeralfileExhibitionV42CallerRaw struct {
	Contract *FeralfileExhibitionV42Caller // Generic read-only contract binding to access the raw methods on
}

// FeralfileExhibitionV42TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeralfileExhibitionV42TransactorRaw struct {
	Contract *FeralfileExhibitionV42Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFeralfileExhibitionV42 creates a new instance of FeralfileExhibitionV42, bound to a specific deployed contract.
func NewFeralfileExhibitionV42(address common.Address, backend bind.ContractBackend) (*FeralfileExhibitionV42, error) {
	contract, err := bindFeralfileExhibitionV42(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42{FeralfileExhibitionV42Caller: FeralfileExhibitionV42Caller{contract: contract}, FeralfileExhibitionV42Transactor: FeralfileExhibitionV42Transactor{contract: contract}, FeralfileExhibitionV42Filterer: FeralfileExhibitionV42Filterer{contract: contract}}, nil
}

// NewFeralfileExhibitionV42Caller creates a new read-only instance of FeralfileExhibitionV42, bound to a specific deployed contract.
func NewFeralfileExhibitionV42Caller(address common.Address, caller bind.ContractCaller) (*FeralfileExhibitionV42Caller, error) {
	contract, err := bindFeralfileExhibitionV42(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42Caller{contract: contract}, nil
}

// NewFeralfileExhibitionV42Transactor creates a new write-only instance of FeralfileExhibitionV42, bound to a specific deployed contract.
func NewFeralfileExhibitionV42Transactor(address common.Address, transactor bind.ContractTransactor) (*FeralfileExhibitionV42Transactor, error) {
	contract, err := bindFeralfileExhibitionV42(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42Transactor{contract: contract}, nil
}

// NewFeralfileExhibitionV42Filterer creates a new log filterer instance of FeralfileExhibitionV42, bound to a specific deployed contract.
func NewFeralfileExhibitionV42Filterer(address common.Address, filterer bind.ContractFilterer) (*FeralfileExhibitionV42Filterer, error) {
	contract, err := bindFeralfileExhibitionV42(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42Filterer{contract: contract}, nil
}

// bindFeralfileExhibitionV42 binds a generic wrapper to an already deployed contract.
func bindFeralfileExhibitionV42(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeralfileExhibitionV42MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralfileExhibitionV42.Contract.FeralfileExhibitionV42Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.FeralfileExhibitionV42Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.FeralfileExhibitionV42Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralfileExhibitionV42.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.contract.Transact(opts, method, params...)
}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xf07e7fd0.
//
// Solidity: function OperatorFilterRegistry() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) OperatorFilterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "OperatorFilterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xf07e7fd0.
//
// Solidity: function OperatorFilterRegistry() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) OperatorFilterRegistry() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.OperatorFilterRegistry(&_FeralfileExhibitionV42.CallOpts)
}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xf07e7fd0.
//
// Solidity: function OperatorFilterRegistry() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) OperatorFilterRegistry() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.OperatorFilterRegistry(&_FeralfileExhibitionV42.CallOpts)
}

// Advances is a free data retrieval call binding the contract method 0x926ce44e.
//
// Solidity: function advances(address ) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Advances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "advances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Advances is a free data retrieval call binding the contract method 0x926ce44e.
//
// Solidity: function advances(address ) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Advances(arg0 common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.Advances(&_FeralfileExhibitionV42.CallOpts, arg0)
}

// Advances is a free data retrieval call binding the contract method 0x926ce44e.
//
// Solidity: function advances(address ) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Advances(arg0 common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.Advances(&_FeralfileExhibitionV42.CallOpts, arg0)
}

// ArtworkFileURI is a free data retrieval call binding the contract method 0xe9f585e0.
//
// Solidity: function artworkFileURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) ArtworkFileURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "artworkFileURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ArtworkFileURI is a free data retrieval call binding the contract method 0xe9f585e0.
//
// Solidity: function artworkFileURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) ArtworkFileURI() (string, error) {
	return _FeralfileExhibitionV42.Contract.ArtworkFileURI(&_FeralfileExhibitionV42.CallOpts)
}

// ArtworkFileURI is a free data retrieval call binding the contract method 0xe9f585e0.
//
// Solidity: function artworkFileURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) ArtworkFileURI() (string, error) {
	return _FeralfileExhibitionV42.Contract.ArtworkFileURI(&_FeralfileExhibitionV42.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.BalanceOf(&_FeralfileExhibitionV42.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.BalanceOf(&_FeralfileExhibitionV42.CallOpts, owner)
}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Bridgeable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "bridgeable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Bridgeable() (bool, error) {
	return _FeralfileExhibitionV42.Contract.Bridgeable(&_FeralfileExhibitionV42.CallOpts)
}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Bridgeable() (bool, error) {
	return _FeralfileExhibitionV42.Contract.Bridgeable(&_FeralfileExhibitionV42.CallOpts)
}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Burnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "burnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Burnable() (bool, error) {
	return _FeralfileExhibitionV42.Contract.Burnable(&_FeralfileExhibitionV42.CallOpts)
}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Burnable() (bool, error) {
	return _FeralfileExhibitionV42.Contract.Burnable(&_FeralfileExhibitionV42.CallOpts)
}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) CodeVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "codeVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) CodeVersion() (string, error) {
	return _FeralfileExhibitionV42.Contract.CodeVersion(&_FeralfileExhibitionV42.CallOpts)
}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) CodeVersion() (string, error) {
	return _FeralfileExhibitionV42.Contract.CodeVersion(&_FeralfileExhibitionV42.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) ContractURI() (string, error) {
	return _FeralfileExhibitionV42.Contract.ContractURI(&_FeralfileExhibitionV42.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) ContractURI() (string, error) {
	return _FeralfileExhibitionV42.Contract.ContractURI(&_FeralfileExhibitionV42.CallOpts)
}

// CostReceiver is a free data retrieval call binding the contract method 0xf4e638be.
//
// Solidity: function costReceiver() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) CostReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "costReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CostReceiver is a free data retrieval call binding the contract method 0xf4e638be.
//
// Solidity: function costReceiver() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) CostReceiver() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.CostReceiver(&_FeralfileExhibitionV42.CallOpts)
}

// CostReceiver is a free data retrieval call binding the contract method 0xf4e638be.
//
// Solidity: function costReceiver() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) CostReceiver() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.CostReceiver(&_FeralfileExhibitionV42.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.GetApproved(&_FeralfileExhibitionV42.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.GetApproved(&_FeralfileExhibitionV42.CallOpts, tokenId)
}

// GetArtwork is a free data retrieval call binding the contract method 0x167ddf6e.
//
// Solidity: function getArtwork(uint256 tokenId) view returns((uint256,uint256))
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) GetArtwork(opts *bind.CallOpts, tokenId *big.Int) (FeralfileExhibitionV4Artwork, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "getArtwork", tokenId)

	if err != nil {
		return *new(FeralfileExhibitionV4Artwork), err
	}

	out0 := *abi.ConvertType(out[0], new(FeralfileExhibitionV4Artwork)).(*FeralfileExhibitionV4Artwork)

	return out0, err

}

// GetArtwork is a free data retrieval call binding the contract method 0x167ddf6e.
//
// Solidity: function getArtwork(uint256 tokenId) view returns((uint256,uint256))
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) GetArtwork(tokenId *big.Int) (FeralfileExhibitionV4Artwork, error) {
	return _FeralfileExhibitionV42.Contract.GetArtwork(&_FeralfileExhibitionV42.CallOpts, tokenId)
}

// GetArtwork is a free data retrieval call binding the contract method 0x167ddf6e.
//
// Solidity: function getArtwork(uint256 tokenId) view returns((uint256,uint256))
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) GetArtwork(tokenId *big.Int) (FeralfileExhibitionV4Artwork, error) {
	return _FeralfileExhibitionV42.Contract.GetArtwork(&_FeralfileExhibitionV42.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FeralfileExhibitionV42.Contract.IsApprovedForAll(&_FeralfileExhibitionV42.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FeralfileExhibitionV42.Contract.IsApprovedForAll(&_FeralfileExhibitionV42.CallOpts, owner, operator)
}

// MintArtworks is a free data retrieval call binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] ) view returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) MintArtworks(opts *bind.CallOpts, arg0 []FeralfileExhibitionV4MintData) error {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "mintArtworks", arg0)

	if err != nil {
		return err
	}

	return err

}

// MintArtworks is a free data retrieval call binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] ) view returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) MintArtworks(arg0 []FeralfileExhibitionV4MintData) error {
	return _FeralfileExhibitionV42.Contract.MintArtworks(&_FeralfileExhibitionV42.CallOpts, arg0)
}

// MintArtworks is a free data retrieval call binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] ) view returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) MintArtworks(arg0 []FeralfileExhibitionV4MintData) error {
	return _FeralfileExhibitionV42.Contract.MintArtworks(&_FeralfileExhibitionV42.CallOpts, arg0)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Mintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "mintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Mintable() (bool, error) {
	return _FeralfileExhibitionV42.Contract.Mintable(&_FeralfileExhibitionV42.CallOpts)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Mintable() (bool, error) {
	return _FeralfileExhibitionV42.Contract.Mintable(&_FeralfileExhibitionV42.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Name() (string, error) {
	return _FeralfileExhibitionV42.Contract.Name(&_FeralfileExhibitionV42.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Name() (string, error) {
	return _FeralfileExhibitionV42.Contract.Name(&_FeralfileExhibitionV42.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Owner() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.Owner(&_FeralfileExhibitionV42.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Owner() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.Owner(&_FeralfileExhibitionV42.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.OwnerOf(&_FeralfileExhibitionV42.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.OwnerOf(&_FeralfileExhibitionV42.CallOpts, tokenId)
}

// Selling is a free data retrieval call binding the contract method 0x23aed228.
//
// Solidity: function selling() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Selling(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "selling")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Selling is a free data retrieval call binding the contract method 0x23aed228.
//
// Solidity: function selling() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Selling() (bool, error) {
	return _FeralfileExhibitionV42.Contract.Selling(&_FeralfileExhibitionV42.CallOpts)
}

// Selling is a free data retrieval call binding the contract method 0x23aed228.
//
// Solidity: function selling() view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Selling() (bool, error) {
	return _FeralfileExhibitionV42.Contract.Selling(&_FeralfileExhibitionV42.CallOpts)
}

// SeriesMaxSupply is a free data retrieval call binding the contract method 0xeb5c60f2.
//
// Solidity: function seriesMaxSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) SeriesMaxSupply(opts *bind.CallOpts, seriesId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "seriesMaxSupply", seriesId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeriesMaxSupply is a free data retrieval call binding the contract method 0xeb5c60f2.
//
// Solidity: function seriesMaxSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SeriesMaxSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.SeriesMaxSupply(&_FeralfileExhibitionV42.CallOpts, seriesId)
}

// SeriesMaxSupply is a free data retrieval call binding the contract method 0xeb5c60f2.
//
// Solidity: function seriesMaxSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) SeriesMaxSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.SeriesMaxSupply(&_FeralfileExhibitionV42.CallOpts, seriesId)
}

// SeriesTotalSupply is a free data retrieval call binding the contract method 0x7f06ee06.
//
// Solidity: function seriesTotalSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) SeriesTotalSupply(opts *bind.CallOpts, seriesId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "seriesTotalSupply", seriesId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeriesTotalSupply is a free data retrieval call binding the contract method 0x7f06ee06.
//
// Solidity: function seriesTotalSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SeriesTotalSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.SeriesTotalSupply(&_FeralfileExhibitionV42.CallOpts, seriesId)
}

// SeriesTotalSupply is a free data retrieval call binding the contract method 0x7f06ee06.
//
// Solidity: function seriesTotalSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) SeriesTotalSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.SeriesTotalSupply(&_FeralfileExhibitionV42.CallOpts, seriesId)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Signer() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.Signer(&_FeralfileExhibitionV42.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Signer() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.Signer(&_FeralfileExhibitionV42.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeralfileExhibitionV42.Contract.SupportsInterface(&_FeralfileExhibitionV42.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeralfileExhibitionV42.Contract.SupportsInterface(&_FeralfileExhibitionV42.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Symbol() (string, error) {
	return _FeralfileExhibitionV42.Contract.Symbol(&_FeralfileExhibitionV42.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Symbol() (string, error) {
	return _FeralfileExhibitionV42.Contract.Symbol(&_FeralfileExhibitionV42.CallOpts)
}

// TokenBaseURI is a free data retrieval call binding the contract method 0x4e99b800.
//
// Solidity: function tokenBaseURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) TokenBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "tokenBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenBaseURI is a free data retrieval call binding the contract method 0x4e99b800.
//
// Solidity: function tokenBaseURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) TokenBaseURI() (string, error) {
	return _FeralfileExhibitionV42.Contract.TokenBaseURI(&_FeralfileExhibitionV42.CallOpts)
}

// TokenBaseURI is a free data retrieval call binding the contract method 0x4e99b800.
//
// Solidity: function tokenBaseURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) TokenBaseURI() (string, error) {
	return _FeralfileExhibitionV42.Contract.TokenBaseURI(&_FeralfileExhibitionV42.CallOpts)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.TokenOfOwnerByIndex(&_FeralfileExhibitionV42.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.TokenOfOwnerByIndex(&_FeralfileExhibitionV42.CallOpts, owner, index)
}

// TokenPlaceholderAnimationURI is a free data retrieval call binding the contract method 0xccdae73a.
//
// Solidity: function tokenPlaceholderAnimationURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) TokenPlaceholderAnimationURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "tokenPlaceholderAnimationURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenPlaceholderAnimationURI is a free data retrieval call binding the contract method 0xccdae73a.
//
// Solidity: function tokenPlaceholderAnimationURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) TokenPlaceholderAnimationURI() (string, error) {
	return _FeralfileExhibitionV42.Contract.TokenPlaceholderAnimationURI(&_FeralfileExhibitionV42.CallOpts)
}

// TokenPlaceholderAnimationURI is a free data retrieval call binding the contract method 0xccdae73a.
//
// Solidity: function tokenPlaceholderAnimationURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) TokenPlaceholderAnimationURI() (string, error) {
	return _FeralfileExhibitionV42.Contract.TokenPlaceholderAnimationURI(&_FeralfileExhibitionV42.CallOpts)
}

// TokenPlaceholderImageURI is a free data retrieval call binding the contract method 0x3e41dd20.
//
// Solidity: function tokenPlaceholderImageURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) TokenPlaceholderImageURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "tokenPlaceholderImageURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenPlaceholderImageURI is a free data retrieval call binding the contract method 0x3e41dd20.
//
// Solidity: function tokenPlaceholderImageURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) TokenPlaceholderImageURI() (string, error) {
	return _FeralfileExhibitionV42.Contract.TokenPlaceholderImageURI(&_FeralfileExhibitionV42.CallOpts)
}

// TokenPlaceholderImageURI is a free data retrieval call binding the contract method 0x3e41dd20.
//
// Solidity: function tokenPlaceholderImageURI() view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) TokenPlaceholderImageURI() (string, error) {
	return _FeralfileExhibitionV42.Contract.TokenPlaceholderImageURI(&_FeralfileExhibitionV42.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) TokenURI(tokenId *big.Int) (string, error) {
	return _FeralfileExhibitionV42.Contract.TokenURI(&_FeralfileExhibitionV42.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _FeralfileExhibitionV42.Contract.TokenURI(&_FeralfileExhibitionV42.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.TokensOfOwner(&_FeralfileExhibitionV42.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.TokensOfOwner(&_FeralfileExhibitionV42.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) TotalSupply() (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.TotalSupply(&_FeralfileExhibitionV42.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) TotalSupply() (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.TotalSupply(&_FeralfileExhibitionV42.CallOpts)
}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Trustees(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "trustees", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Trustees(arg0 common.Address) (bool, error) {
	return _FeralfileExhibitionV42.Contract.Trustees(&_FeralfileExhibitionV42.CallOpts, arg0)
}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Trustees(arg0 common.Address) (bool, error) {
	return _FeralfileExhibitionV42.Contract.Trustees(&_FeralfileExhibitionV42.CallOpts, arg0)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Vault() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.Vault(&_FeralfileExhibitionV42.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Vault() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.Vault(&_FeralfileExhibitionV42.CallOpts)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) AddTrustee(opts *bind.TransactOpts, _trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "addTrustee", _trustee)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) AddTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.AddTrustee(&_FeralfileExhibitionV42.TransactOpts, _trustee)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) AddTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.AddTrustee(&_FeralfileExhibitionV42.TransactOpts, _trustee)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.Approve(&_FeralfileExhibitionV42.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.Approve(&_FeralfileExhibitionV42.TransactOpts, operator, tokenId)
}

// BurnArtworks is a paid mutator transaction binding the contract method 0x21fe0c64.
//
// Solidity: function burnArtworks(uint256[] tokenIds) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) BurnArtworks(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "burnArtworks", tokenIds)
}

// BurnArtworks is a paid mutator transaction binding the contract method 0x21fe0c64.
//
// Solidity: function burnArtworks(uint256[] tokenIds) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) BurnArtworks(tokenIds []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.BurnArtworks(&_FeralfileExhibitionV42.TransactOpts, tokenIds)
}

// BurnArtworks is a paid mutator transaction binding the contract method 0x21fe0c64.
//
// Solidity: function burnArtworks(uint256[] tokenIds) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) BurnArtworks(tokenIds []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.BurnArtworks(&_FeralfileExhibitionV42.TransactOpts, tokenIds)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_) payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) BuyArtworks(opts *bind.TransactOpts, r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "buyArtworks", r_, s_, v_, saleData_)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_) payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) BuyArtworks(r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.BuyArtworks(&_FeralfileExhibitionV42.TransactOpts, r_, s_, v_, saleData_)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_) payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) BuyArtworks(r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.BuyArtworks(&_FeralfileExhibitionV42.TransactOpts, r_, s_, v_, saleData_)
}

// MintArtworksWithIndex is a paid mutator transaction binding the contract method 0x84f78930.
//
// Solidity: function mintArtworksWithIndex((uint256,uint256,address,uint256)[] data) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) MintArtworksWithIndex(opts *bind.TransactOpts, data []FeralfileExhibitionV42MintDataWithIndex) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "mintArtworksWithIndex", data)
}

// MintArtworksWithIndex is a paid mutator transaction binding the contract method 0x84f78930.
//
// Solidity: function mintArtworksWithIndex((uint256,uint256,address,uint256)[] data) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) MintArtworksWithIndex(data []FeralfileExhibitionV42MintDataWithIndex) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.MintArtworksWithIndex(&_FeralfileExhibitionV42.TransactOpts, data)
}

// MintArtworksWithIndex is a paid mutator transaction binding the contract method 0x84f78930.
//
// Solidity: function mintArtworksWithIndex((uint256,uint256,address,uint256)[] data) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) MintArtworksWithIndex(data []FeralfileExhibitionV42MintDataWithIndex) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.MintArtworksWithIndex(&_FeralfileExhibitionV42.TransactOpts, data)
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) PauseSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "pauseSale")
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) PauseSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.PauseSale(&_FeralfileExhibitionV42.TransactOpts)
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) PauseSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.PauseSale(&_FeralfileExhibitionV42.TransactOpts)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) RemoveTrustee(opts *bind.TransactOpts, _trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "removeTrustee", _trustee)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) RemoveTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.RemoveTrustee(&_FeralfileExhibitionV42.TransactOpts, _trustee)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) RemoveTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.RemoveTrustee(&_FeralfileExhibitionV42.TransactOpts, _trustee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) RenounceOwnership() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.RenounceOwnership(&_FeralfileExhibitionV42.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.RenounceOwnership(&_FeralfileExhibitionV42.TransactOpts)
}

// ReplaceAdvanceAddresses is a paid mutator transaction binding the contract method 0x41a5626a.
//
// Solidity: function replaceAdvanceAddresses(address[] oldAddresses_, address[] newAddresses_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) ReplaceAdvanceAddresses(opts *bind.TransactOpts, oldAddresses_ []common.Address, newAddresses_ []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "replaceAdvanceAddresses", oldAddresses_, newAddresses_)
}

// ReplaceAdvanceAddresses is a paid mutator transaction binding the contract method 0x41a5626a.
//
// Solidity: function replaceAdvanceAddresses(address[] oldAddresses_, address[] newAddresses_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) ReplaceAdvanceAddresses(oldAddresses_ []common.Address, newAddresses_ []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.ReplaceAdvanceAddresses(&_FeralfileExhibitionV42.TransactOpts, oldAddresses_, newAddresses_)
}

// ReplaceAdvanceAddresses is a paid mutator transaction binding the contract method 0x41a5626a.
//
// Solidity: function replaceAdvanceAddresses(address[] oldAddresses_, address[] newAddresses_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) ReplaceAdvanceAddresses(oldAddresses_ []common.Address, newAddresses_ []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.ReplaceAdvanceAddresses(&_FeralfileExhibitionV42.TransactOpts, oldAddresses_, newAddresses_)
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) ResumeSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "resumeSale")
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) ResumeSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.ResumeSale(&_FeralfileExhibitionV42.TransactOpts)
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) ResumeSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.ResumeSale(&_FeralfileExhibitionV42.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SafeTransferFrom(&_FeralfileExhibitionV42.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SafeTransferFrom(&_FeralfileExhibitionV42.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SafeTransferFrom0(&_FeralfileExhibitionV42.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SafeTransferFrom0(&_FeralfileExhibitionV42.TransactOpts, from, to, tokenId, data)
}

// SetAdvanceSetting is a paid mutator transaction binding the contract method 0x3c352b0d.
//
// Solidity: function setAdvanceSetting(address[] addresses_, uint256[] amounts_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SetAdvanceSetting(opts *bind.TransactOpts, addresses_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "setAdvanceSetting", addresses_, amounts_)
}

// SetAdvanceSetting is a paid mutator transaction binding the contract method 0x3c352b0d.
//
// Solidity: function setAdvanceSetting(address[] addresses_, uint256[] amounts_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetAdvanceSetting(addresses_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetAdvanceSetting(&_FeralfileExhibitionV42.TransactOpts, addresses_, amounts_)
}

// SetAdvanceSetting is a paid mutator transaction binding the contract method 0x3c352b0d.
//
// Solidity: function setAdvanceSetting(address[] addresses_, uint256[] amounts_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SetAdvanceSetting(addresses_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetAdvanceSetting(&_FeralfileExhibitionV42.TransactOpts, addresses_, amounts_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetApprovalForAll(&_FeralfileExhibitionV42.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetApprovalForAll(&_FeralfileExhibitionV42.TransactOpts, operator, approved)
}

// SetArtworkFileURI is a paid mutator transaction binding the contract method 0x3c38598e.
//
// Solidity: function setArtworkFileURI(string uri) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SetArtworkFileURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "setArtworkFileURI", uri)
}

// SetArtworkFileURI is a paid mutator transaction binding the contract method 0x3c38598e.
//
// Solidity: function setArtworkFileURI(string uri) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetArtworkFileURI(uri string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetArtworkFileURI(&_FeralfileExhibitionV42.TransactOpts, uri)
}

// SetArtworkFileURI is a paid mutator transaction binding the contract method 0x3c38598e.
//
// Solidity: function setArtworkFileURI(string uri) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SetArtworkFileURI(uri string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetArtworkFileURI(&_FeralfileExhibitionV42.TransactOpts, uri)
}

// SetCostReceiver is a paid mutator transaction binding the contract method 0x1623528f.
//
// Solidity: function setCostReceiver(address costReceiver_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SetCostReceiver(opts *bind.TransactOpts, costReceiver_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "setCostReceiver", costReceiver_)
}

// SetCostReceiver is a paid mutator transaction binding the contract method 0x1623528f.
//
// Solidity: function setCostReceiver(address costReceiver_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetCostReceiver(costReceiver_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetCostReceiver(&_FeralfileExhibitionV42.TransactOpts, costReceiver_)
}

// SetCostReceiver is a paid mutator transaction binding the contract method 0x1623528f.
//
// Solidity: function setCostReceiver(address costReceiver_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SetCostReceiver(costReceiver_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetCostReceiver(&_FeralfileExhibitionV42.TransactOpts, costReceiver_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SetSigner(opts *bind.TransactOpts, signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "setSigner", signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetSigner(&_FeralfileExhibitionV42.TransactOpts, signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetSigner(&_FeralfileExhibitionV42.TransactOpts, signer_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SetTokenBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "setTokenBaseURI", baseURI_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetTokenBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetTokenBaseURI(&_FeralfileExhibitionV42.TransactOpts, baseURI_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SetTokenBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetTokenBaseURI(&_FeralfileExhibitionV42.TransactOpts, baseURI_)
}

// SetTokenPlaceholderAnimationURI is a paid mutator transaction binding the contract method 0x9095e876.
//
// Solidity: function setTokenPlaceholderAnimationURI(string uri) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SetTokenPlaceholderAnimationURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "setTokenPlaceholderAnimationURI", uri)
}

// SetTokenPlaceholderAnimationURI is a paid mutator transaction binding the contract method 0x9095e876.
//
// Solidity: function setTokenPlaceholderAnimationURI(string uri) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetTokenPlaceholderAnimationURI(uri string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetTokenPlaceholderAnimationURI(&_FeralfileExhibitionV42.TransactOpts, uri)
}

// SetTokenPlaceholderAnimationURI is a paid mutator transaction binding the contract method 0x9095e876.
//
// Solidity: function setTokenPlaceholderAnimationURI(string uri) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SetTokenPlaceholderAnimationURI(uri string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetTokenPlaceholderAnimationURI(&_FeralfileExhibitionV42.TransactOpts, uri)
}

// SetTokenPlaceholderImageURI is a paid mutator transaction binding the contract method 0xc9200e4f.
//
// Solidity: function setTokenPlaceholderImageURI(string uri) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SetTokenPlaceholderImageURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "setTokenPlaceholderImageURI", uri)
}

// SetTokenPlaceholderImageURI is a paid mutator transaction binding the contract method 0xc9200e4f.
//
// Solidity: function setTokenPlaceholderImageURI(string uri) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetTokenPlaceholderImageURI(uri string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetTokenPlaceholderImageURI(&_FeralfileExhibitionV42.TransactOpts, uri)
}

// SetTokenPlaceholderImageURI is a paid mutator transaction binding the contract method 0xc9200e4f.
//
// Solidity: function setTokenPlaceholderImageURI(string uri) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SetTokenPlaceholderImageURI(uri string) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetTokenPlaceholderImageURI(&_FeralfileExhibitionV42.TransactOpts, uri)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vault_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SetVault(opts *bind.TransactOpts, vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "setVault", vault_)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vault_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetVault(vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetVault(&_FeralfileExhibitionV42.TransactOpts, vault_)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vault_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SetVault(vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetVault(&_FeralfileExhibitionV42.TransactOpts, vault_)
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) StartSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "startSale")
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) StartSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.StartSale(&_FeralfileExhibitionV42.TransactOpts)
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) StartSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.StartSale(&_FeralfileExhibitionV42.TransactOpts)
}

// StopSaleAndBurn is a paid mutator transaction binding the contract method 0xb9b8311a.
//
// Solidity: function stopSaleAndBurn() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) StopSaleAndBurn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "stopSaleAndBurn")
}

// StopSaleAndBurn is a paid mutator transaction binding the contract method 0xb9b8311a.
//
// Solidity: function stopSaleAndBurn() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) StopSaleAndBurn() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.StopSaleAndBurn(&_FeralfileExhibitionV42.TransactOpts)
}

// StopSaleAndBurn is a paid mutator transaction binding the contract method 0xb9b8311a.
//
// Solidity: function stopSaleAndBurn() returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) StopSaleAndBurn() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.StopSaleAndBurn(&_FeralfileExhibitionV42.TransactOpts)
}

// StopSaleAndTransfer is a paid mutator transaction binding the contract method 0x65a46e08.
//
// Solidity: function stopSaleAndTransfer(uint256[] seriesIds, address[] recipientAddresses) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) StopSaleAndTransfer(opts *bind.TransactOpts, seriesIds []*big.Int, recipientAddresses []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "stopSaleAndTransfer", seriesIds, recipientAddresses)
}

// StopSaleAndTransfer is a paid mutator transaction binding the contract method 0x65a46e08.
//
// Solidity: function stopSaleAndTransfer(uint256[] seriesIds, address[] recipientAddresses) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) StopSaleAndTransfer(seriesIds []*big.Int, recipientAddresses []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.StopSaleAndTransfer(&_FeralfileExhibitionV42.TransactOpts, seriesIds, recipientAddresses)
}

// StopSaleAndTransfer is a paid mutator transaction binding the contract method 0x65a46e08.
//
// Solidity: function stopSaleAndTransfer(uint256[] seriesIds, address[] recipientAddresses) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) StopSaleAndTransfer(seriesIds []*big.Int, recipientAddresses []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.StopSaleAndTransfer(&_FeralfileExhibitionV42.TransactOpts, seriesIds, recipientAddresses)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.TransferFrom(&_FeralfileExhibitionV42.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.TransferFrom(&_FeralfileExhibitionV42.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.TransferOwnership(&_FeralfileExhibitionV42.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.TransferOwnership(&_FeralfileExhibitionV42.TransactOpts, newOwner)
}

// UpdateOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x114ba8ee.
//
// Solidity: function updateOperatorFilterRegistry(address operatorFilterRegisterAddress) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) UpdateOperatorFilterRegistry(opts *bind.TransactOpts, operatorFilterRegisterAddress common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "updateOperatorFilterRegistry", operatorFilterRegisterAddress)
}

// UpdateOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x114ba8ee.
//
// Solidity: function updateOperatorFilterRegistry(address operatorFilterRegisterAddress) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) UpdateOperatorFilterRegistry(operatorFilterRegisterAddress common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.UpdateOperatorFilterRegistry(&_FeralfileExhibitionV42.TransactOpts, operatorFilterRegisterAddress)
}

// UpdateOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x114ba8ee.
//
// Solidity: function updateOperatorFilterRegistry(address operatorFilterRegisterAddress) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) UpdateOperatorFilterRegistry(operatorFilterRegisterAddress common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.UpdateOperatorFilterRegistry(&_FeralfileExhibitionV42.TransactOpts, operatorFilterRegisterAddress)
}

// UpdateTokenInformation is a paid mutator transaction binding the contract method 0xae474c3c.
//
// Solidity: function updateTokenInformation(uint256 tokenId, string imageURI, bytes parameters, uint256 coinAmount) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) UpdateTokenInformation(opts *bind.TransactOpts, tokenId *big.Int, imageURI string, parameters []byte, coinAmount *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "updateTokenInformation", tokenId, imageURI, parameters, coinAmount)
}

// UpdateTokenInformation is a paid mutator transaction binding the contract method 0xae474c3c.
//
// Solidity: function updateTokenInformation(uint256 tokenId, string imageURI, bytes parameters, uint256 coinAmount) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) UpdateTokenInformation(tokenId *big.Int, imageURI string, parameters []byte, coinAmount *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.UpdateTokenInformation(&_FeralfileExhibitionV42.TransactOpts, tokenId, imageURI, parameters, coinAmount)
}

// UpdateTokenInformation is a paid mutator transaction binding the contract method 0xae474c3c.
//
// Solidity: function updateTokenInformation(uint256 tokenId, string imageURI, bytes parameters, uint256 coinAmount) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) UpdateTokenInformation(tokenId *big.Int, imageURI string, parameters []byte, coinAmount *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.UpdateTokenInformation(&_FeralfileExhibitionV42.TransactOpts, tokenId, imageURI, parameters, coinAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Receive() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.Receive(&_FeralfileExhibitionV42.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) Receive() (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.Receive(&_FeralfileExhibitionV42.TransactOpts)
}

// FeralfileExhibitionV42ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42ApprovalIterator struct {
	Event *FeralfileExhibitionV42Approval // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV42ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV42Approval)
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
		it.Event = new(FeralfileExhibitionV42Approval)
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
func (it *FeralfileExhibitionV42ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV42ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV42Approval represents a Approval event raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV42ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42ApprovalIterator{contract: _FeralfileExhibitionV42.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV42Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV42Approval)
				if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) ParseApproval(log types.Log) (*FeralfileExhibitionV42Approval, error) {
	event := new(FeralfileExhibitionV42Approval)
	if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV42ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42ApprovalForAllIterator struct {
	Event *FeralfileExhibitionV42ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV42ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV42ApprovalForAll)
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
		it.Event = new(FeralfileExhibitionV42ApprovalForAll)
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
func (it *FeralfileExhibitionV42ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV42ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV42ApprovalForAll represents a ApprovalForAll event raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*FeralfileExhibitionV42ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42ApprovalForAllIterator{contract: _FeralfileExhibitionV42.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV42ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV42ApprovalForAll)
				if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) ParseApprovalForAll(log types.Log) (*FeralfileExhibitionV42ApprovalForAll, error) {
	event := new(FeralfileExhibitionV42ApprovalForAll)
	if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV42BurnArtworkIterator is returned from FilterBurnArtwork and is used to iterate over the raw logs and unpacked data for BurnArtwork events raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42BurnArtworkIterator struct {
	Event *FeralfileExhibitionV42BurnArtwork // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV42BurnArtworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV42BurnArtwork)
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
		it.Event = new(FeralfileExhibitionV42BurnArtwork)
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
func (it *FeralfileExhibitionV42BurnArtworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV42BurnArtworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV42BurnArtwork represents a BurnArtwork event raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42BurnArtwork struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnArtwork is a free log retrieval operation binding the contract event 0xbde7938970372996ff103863625e348ef2bf8f38a5b02181be75aafef17c23d5.
//
// Solidity: event BurnArtwork(uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) FilterBurnArtwork(opts *bind.FilterOpts, tokenId []*big.Int) (*FeralfileExhibitionV42BurnArtworkIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.FilterLogs(opts, "BurnArtwork", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42BurnArtworkIterator{contract: _FeralfileExhibitionV42.contract, event: "BurnArtwork", logs: logs, sub: sub}, nil
}

// WatchBurnArtwork is a free log subscription operation binding the contract event 0xbde7938970372996ff103863625e348ef2bf8f38a5b02181be75aafef17c23d5.
//
// Solidity: event BurnArtwork(uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) WatchBurnArtwork(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV42BurnArtwork, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.WatchLogs(opts, "BurnArtwork", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV42BurnArtwork)
				if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "BurnArtwork", log); err != nil {
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

// ParseBurnArtwork is a log parse operation binding the contract event 0xbde7938970372996ff103863625e348ef2bf8f38a5b02181be75aafef17c23d5.
//
// Solidity: event BurnArtwork(uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) ParseBurnArtwork(log types.Log) (*FeralfileExhibitionV42BurnArtwork, error) {
	event := new(FeralfileExhibitionV42BurnArtwork)
	if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "BurnArtwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV42BuyArtworkIterator is returned from FilterBuyArtwork and is used to iterate over the raw logs and unpacked data for BuyArtwork events raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42BuyArtworkIterator struct {
	Event *FeralfileExhibitionV42BuyArtwork // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV42BuyArtworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV42BuyArtwork)
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
		it.Event = new(FeralfileExhibitionV42BuyArtwork)
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
func (it *FeralfileExhibitionV42BuyArtworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV42BuyArtworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV42BuyArtwork represents a BuyArtwork event raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42BuyArtwork struct {
	Buyer   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyArtwork is a free log retrieval operation binding the contract event 0x0475389cd69b8d3163620b43283bf74e8fc71020c3c6cef2a529b5c405e9687f.
//
// Solidity: event BuyArtwork(address indexed buyer, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) FilterBuyArtwork(opts *bind.FilterOpts, buyer []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV42BuyArtworkIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.FilterLogs(opts, "BuyArtwork", buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42BuyArtworkIterator{contract: _FeralfileExhibitionV42.contract, event: "BuyArtwork", logs: logs, sub: sub}, nil
}

// WatchBuyArtwork is a free log subscription operation binding the contract event 0x0475389cd69b8d3163620b43283bf74e8fc71020c3c6cef2a529b5c405e9687f.
//
// Solidity: event BuyArtwork(address indexed buyer, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) WatchBuyArtwork(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV42BuyArtwork, buyer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.WatchLogs(opts, "BuyArtwork", buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV42BuyArtwork)
				if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "BuyArtwork", log); err != nil {
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

// ParseBuyArtwork is a log parse operation binding the contract event 0x0475389cd69b8d3163620b43283bf74e8fc71020c3c6cef2a529b5c405e9687f.
//
// Solidity: event BuyArtwork(address indexed buyer, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) ParseBuyArtwork(log types.Log) (*FeralfileExhibitionV42BuyArtwork, error) {
	event := new(FeralfileExhibitionV42BuyArtwork)
	if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "BuyArtwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV42NewArtworkIterator is returned from FilterNewArtwork and is used to iterate over the raw logs and unpacked data for NewArtwork events raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42NewArtworkIterator struct {
	Event *FeralfileExhibitionV42NewArtwork // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV42NewArtworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV42NewArtwork)
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
		it.Event = new(FeralfileExhibitionV42NewArtwork)
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
func (it *FeralfileExhibitionV42NewArtworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV42NewArtworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV42NewArtwork represents a NewArtwork event raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42NewArtwork struct {
	Owner    common.Address
	SeriesId *big.Int
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewArtwork is a free log retrieval operation binding the contract event 0x407d7da1d3b2b1871fbfa2b5b1c4657a3cc5711d3023c552798551c7ee301eea.
//
// Solidity: event NewArtwork(address indexed owner, uint256 indexed seriesId, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) FilterNewArtwork(opts *bind.FilterOpts, owner []common.Address, seriesId []*big.Int, tokenId []*big.Int) (*FeralfileExhibitionV42NewArtworkIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var seriesIdRule []interface{}
	for _, seriesIdItem := range seriesId {
		seriesIdRule = append(seriesIdRule, seriesIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.FilterLogs(opts, "NewArtwork", ownerRule, seriesIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42NewArtworkIterator{contract: _FeralfileExhibitionV42.contract, event: "NewArtwork", logs: logs, sub: sub}, nil
}

// WatchNewArtwork is a free log subscription operation binding the contract event 0x407d7da1d3b2b1871fbfa2b5b1c4657a3cc5711d3023c552798551c7ee301eea.
//
// Solidity: event NewArtwork(address indexed owner, uint256 indexed seriesId, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) WatchNewArtwork(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV42NewArtwork, owner []common.Address, seriesId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var seriesIdRule []interface{}
	for _, seriesIdItem := range seriesId {
		seriesIdRule = append(seriesIdRule, seriesIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.WatchLogs(opts, "NewArtwork", ownerRule, seriesIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV42NewArtwork)
				if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "NewArtwork", log); err != nil {
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

// ParseNewArtwork is a log parse operation binding the contract event 0x407d7da1d3b2b1871fbfa2b5b1c4657a3cc5711d3023c552798551c7ee301eea.
//
// Solidity: event NewArtwork(address indexed owner, uint256 indexed seriesId, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) ParseNewArtwork(log types.Log) (*FeralfileExhibitionV42NewArtwork, error) {
	event := new(FeralfileExhibitionV42NewArtwork)
	if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "NewArtwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV42OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42OwnershipTransferredIterator struct {
	Event *FeralfileExhibitionV42OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV42OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV42OwnershipTransferred)
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
		it.Event = new(FeralfileExhibitionV42OwnershipTransferred)
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
func (it *FeralfileExhibitionV42OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV42OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV42OwnershipTransferred represents a OwnershipTransferred event raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeralfileExhibitionV42OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42OwnershipTransferredIterator{contract: _FeralfileExhibitionV42.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV42OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV42OwnershipTransferred)
				if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) ParseOwnershipTransferred(log types.Log) (*FeralfileExhibitionV42OwnershipTransferred, error) {
	event := new(FeralfileExhibitionV42OwnershipTransferred)
	if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV42TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42TransferIterator struct {
	Event *FeralfileExhibitionV42Transfer // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV42TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV42Transfer)
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
		it.Event = new(FeralfileExhibitionV42Transfer)
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
func (it *FeralfileExhibitionV42TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV42TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV42Transfer represents a Transfer event raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV42TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42TransferIterator{contract: _FeralfileExhibitionV42.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV42Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV42Transfer)
				if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) ParseTransfer(log types.Log) (*FeralfileExhibitionV42Transfer, error) {
	event := new(FeralfileExhibitionV42Transfer)
	if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
