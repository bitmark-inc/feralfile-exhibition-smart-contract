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

// IFeralfileSaleDataV2SaleDataV2 is an auto generated low-level Go binding around an user-defined struct.
type IFeralfileSaleDataV2SaleDataV2 struct {
	Price              *big.Int
	Cost               *big.Int
	ExpiryTime         *big.Int
	Destination        common.Address
	Nonce              *big.Int
	SeriesID           *big.Int
	Quantity           uint16
	RevenueShares      []IFeralfileSaleDataRevenueShare
	PayByVaultContract bool
}

// FeralfileExhibitionV42MetaData contains all meta data concerning the FeralfileExhibitionV42 contract.
var FeralfileExhibitionV42MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"burnable_\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"bridgeable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"costReceiver_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"seriesIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"seriesMaxSupplies_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"seriesNextPurchasableTokenIds_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdvanceAddressAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FunctionNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAddressesAndAmounts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SaleNotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SeriesLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIDNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TotalBpsOver\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BurnArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BuyArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"BuyArtworkV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NewArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OperatorFilterRegistry\",\"outputs\":[{\"internalType\":\"contractIOperatorFilterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"addTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"advances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burnArtworks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"codeVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"costReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getArtwork\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structFeralfileExhibitionV4.Artwork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structFeralfileExhibitionV4.MintData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"mintArtworks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"removeTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oldAddresses_\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"newAddresses_\",\"type\":\"address[]\"}],\"name\":\"replaceAdvanceAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selling\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"seriesMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"seriesTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"setAdvanceSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"costReceiver_\",\"type\":\"address\"}],\"name\":\"setCostReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setTokenBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopSaleAndBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"seriesIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"recipientAddresses\",\"type\":\"address[]\"}],\"name\":\"stopSaleAndTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorFilterRegisterAddress\",\"type\":\"address\"}],\"name\":\"updateOperatorFilterRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIFeralfileVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultV2\",\"outputs\":[{\"internalType\":\"contractIFeralfileVaultV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"}],\"name\":\"setVaultV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seriesID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"quantity\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"internalType\":\"structIFeralfileSaleData.RevenueShare[]\",\"name\":\"revenueShares\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"payByVaultContract\",\"type\":\"bool\"}],\"internalType\":\"structIFeralfileSaleDataV2.SaleDataV2\",\"name\":\"saleData_\",\"type\":\"tuple\"}],\"name\":\"buyBulkArtworks\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"internalType\":\"structIFeralfileSaleData.RevenueShare[][]\",\"name\":\"revenueShares\",\"type\":\"tuple[][]\"},{\"internalType\":\"bool\",\"name\":\"payByVaultContract\",\"type\":\"bool\"}],\"internalType\":\"structIFeralfileSaleData.SaleData\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"buyArtworks\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600880546001600160a01b0319166daaeb6d7670e522a718067333cd4e179055600d805463ff000000191663010000001790553480156200004457600080fd5b5060405162005533380380620055338339810160408190526200006791620009fe565b8a8a8a8a8a8a8a8a8a8a89898989898989898989858a8a60006200008c838262000c1b565b5060016200009b828262000c1b565b505050620000b8620000b26200081c60201b60201c565b62000820565b6008546001600160a01b03163b156200014557600854604051633e9f1edf60e11b8152306004820152733cc6cdda760b79bafa08df41ecfa224f810dceb660248201526001600160a01b0390911690637d3e3dbe90604401600060405180830381600087803b1580156200012b57600080fd5b505af115801562000140573d6000803e3d6000fd5b505050505b6001600160a01b038116620001ac5760405162461bcd60e51b815260206004820152602260248201527f45434453415369676e3a207369676e65725f206973207a65726f206164647265604482015261737360f01b60648201526084015b60405180910390fd5b600980546001600160a01b0319166001600160a01b039290921691909117905589516200022a5760405162461bcd60e51b815260206004820152602560248201527f466572616c66696c6545786869626974696f6e56343a206e616d655f20697320604482015264656d70747960d81b6064820152608401620001a3565b60008951116200028d5760405162461bcd60e51b815260206004820152602760248201527f466572616c66696c6545786869626974696f6e56343a2073796d626f6c5f20696044820152667320656d70747960c81b6064820152608401620001a3565b6001600160a01b0385166200030b5760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a207661756c744164647260448201527f6573735f206973207a65726f20616464726573730000000000000000000000006064820152608401620001a3565b6001600160a01b038416620003895760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a20636f7374526563656960448201527f7665725f206973207a65726f20616464726573730000000000000000000000006064820152608401620001a3565b6000835111620003f15760405162461bcd60e51b815260206004820152602c60248201527f466572616c66696c6545786869626974696f6e56343a20636f6e74726163745560448201526b52495f20697320656d70747960a01b6064820152608401620001a3565b6000825111620004575760405162461bcd60e51b815260206004820152602a60248201527f466572616c66696c6545786869626974696f6e56343a207365726965734964736044820152695f20697320656d70747960b01b6064820152608401620001a3565b6000815111620004c55760405162461bcd60e51b815260206004820152603260248201527f466572616c66696c6545786869626974696f6e56343a205f7365726965734d6160448201527178537570706c69657320697320656d70747960701b6064820152608401620001a3565b8051825114620005585760405162461bcd60e51b815260206004820152605160248201527f466572616c66696c6545786869626974696f6e56343a207365726965734d617860448201527f537570706c6965735f20616e64207365726965734964735f206c656e6774687360648201527020617265206e6f74207468652073616d6560781b608482015260a401620001a3565b600d805461ffff191689151561ff001916176101008915150217600160201b600160c01b0319166401000000006001600160a01b038781169190910291909117909155600e80546001600160a01b031916918716919091179055600b620005c0848262000c1b565b5060005b82518110156200074457600f6000848381518110620005e757620005e762000ce7565b6020026020010151815260200190815260200160002054600014620006615760405162461bcd60e51b815260206004820152602960248201527f466572616c66696c6545786869626974696f6e56343a206475706c6963617465604482015268081cd95c9a595cd25960ba1b6064820152608401620001a3565b600082828151811062000678576200067862000ce7565b602002602001015111620006de5760405162461bcd60e51b815260206004820152602660248201527f466572616c66696c6545786869626974696f6e56343a207a65726f206d617820604482015265737570706c7960d01b6064820152608401620001a3565b818181518110620006f357620006f362000ce7565b6020026020010151600f600085848151811062000714576200071462000ce7565b602002602001015181526020019081526020016000208190555080806200073b9062000cfd565b915050620005c4565b50505050505050505050505050505050505050505080518351146200077c576040516330fa3f3b60e21b815260040160405180910390fd5b601680546001600160a01b0319166001600160a01b03881617905560005b83518110156200080a57818181518110620007b957620007b962000ce7565b602002602001015160176000868481518110620007da57620007da62000ce7565b60200260200101518152602001908152602001600020819055508080620008019062000cfd565b9150506200079a565b50505050505050505050505062000d25565b3390565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620008b357620008b362000872565b604052919050565b600082601f830112620008cd57600080fd5b81516001600160401b03811115620008e957620008e962000872565b6020620008ff601f8301601f1916820162000888565b82815285828487010111156200091457600080fd5b60005b838110156200093457858101830151828201840152820162000917565b506000928101909101919091529392505050565b805180151581146200095957600080fd5b919050565b80516001600160a01b03811681146200095957600080fd5b600082601f8301126200098857600080fd5b815160206001600160401b03821115620009a657620009a662000872565b8160051b620009b782820162000888565b9283528481018201928281019087851115620009d257600080fd5b83870192505b84831015620009f357825182529183019190830190620009d8565b979650505050505050565b60008060008060008060008060008060006101608c8e03121562000a2157600080fd5b8b516001600160401b0381111562000a3857600080fd5b62000a468e828f01620008bb565b60208e0151909c5090506001600160401b0381111562000a6557600080fd5b62000a738e828f01620008bb565b9a505062000a8460408d0162000948565b985062000a9460608d0162000948565b975062000aa460808d016200095e565b965062000ab460a08d016200095e565b955062000ac460c08d016200095e565b60e08d01519095506001600160401b0381111562000ae157600080fd5b62000aef8e828f01620008bb565b6101008e015190955090506001600160401b0381111562000b0f57600080fd5b62000b1d8e828f0162000976565b6101208e015190945090506001600160401b0381111562000b3d57600080fd5b62000b4b8e828f0162000976565b6101408e015190935090506001600160401b0381111562000b6b57600080fd5b62000b798e828f0162000976565b9150509295989b509295989b9093969950565b600181811c9082168062000ba157607f821691505b60208210810362000bc257634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000c1657600081815260208120601f850160051c8101602086101562000bf15750805b601f850160051c820191505b8181101562000c125782815560010162000bfd565b5050505b505050565b81516001600160401b0381111562000c375762000c3762000872565b62000c4f8162000c48845462000b8c565b8462000bc8565b602080601f83116001811462000c87576000841562000c6e5750858301515b600019600386901b1c1916600185901b17855562000c12565b600085815260208120601f198616915b8281101562000cb85788860151825594840194600190910190840162000c97565b508582101562000cd75787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b60006001820162000d1e57634e487b7160e01b600052601160045260246000fd5b5060010190565b6147fe8062000d356000396000f3fe60806040526004361061036f5760003560e01c80636817031b116101c6578063a74cebab116100f7578063e985e9c511610095578063f07e7fd01161006f578063f07e7fd014610aef578063f2fde38b14610b0f578063f4e638be14610b2f578063fbfa77cf14610b5757600080fd5b8063e985e9c514610a49578063eb5c60f214610a92578063eee608a414610abf57600080fd5b8063b9b8311a116100d1578063b9b8311a146109df578063c87b56dd146109f4578063dc78ac1c14610a14578063e8a3d48514610a3457600080fd5b8063a74cebab1461098a578063b66a0e5d146109aa578063b88d4fde146109bf57600080fd5b80638cba1c6711610164578063926ce44e1161013e578063926ce44e1461090e57806395d89b411461093b578063a07c7ce414610950578063a22cb4651461096a57600080fd5b80638cba1c67146108b05780638da5cb5b146108d05780638ef79e91146108ee57600080fd5b8063715018a6116101a0578063715018a61461080b5780637ecebe00146108205780637f06ee06146108565780638462151c1461088357600080fd5b80636817031b146107b05780636c19e783146107cb57806370a08231146107eb57600080fd5b80632f745c59116102a05780634e99b8001161023e5780635eb9bad6116102185780635eb9bad61461070f5780636352211e1461072f57806363e602301461074f57806365a46e081461079057600080fd5b80634e99b800146106c6578063530da8ef146106db57806355367ba9146106fa57600080fd5b806341a5626a1161027a57806341a5626a1461065257806342842e0e146106725780634bda5d89146106925780634bf365df146106a557600080fd5b80632f745c59146105fd57806333e364cb1461061d5780633c352b0d1461063257600080fd5b8063167ddf6e1161030d578063238ac933116102e7578063238ac9331461058e57806323aed228146105ac57806323b872dd146105ca5780632977e4b3146105ea57600080fd5b8063167ddf6e1461050f57806318160ddd1461054a57806321fe0c641461056e57600080fd5b8063081812fc11610349578063081812fc14610477578063095ea7b3146104af578063114ba8ee146104cf5780631623528f146104ef57600080fd5b806301ffc9a714610400578063031205061461043557806306fdde031461045557600080fd5b366103fb57600e546001600160a01b031633146103f95760405162461bcd60e51b815260206004820152603c60248201527f466572616c66696c6545786869626974696f6e56343a206f6e6c79206163636560448201527f70742066756e642066726f6d207661756c7420636f6e74726163742e0000000060648201526084015b60405180910390fd5b005b600080fd5b34801561040c57600080fd5b5061042061041b3660046138fc565b610b77565b60405190151581526020015b60405180910390f35b34801561044157600080fd5b506103f961045036600461393c565b610bc9565b34801561046157600080fd5b5061046a610bf2565b60405161042c91906139a7565b34801561048357600080fd5b506104976104923660046139ba565b610c84565b6040516001600160a01b03909116815260200161042c565b3480156104bb57600080fd5b506103f96104ca3660046139d3565b610cab565b3480156104db57600080fd5b506103f96104ea36600461393c565b610cc4565b3480156104fb57600080fd5b506103f961050a36600461393c565b610cee565b34801561051b57600080fd5b5061052f61052a3660046139ba565b610d97565b6040805182518152602092830151928101929092520161042c565b34801561055657600080fd5b50610560600c5481565b60405190815260200161042c565b34801561057a57600080fd5b506103f9610589366004613ad1565b610dfa565b34801561059a57600080fd5b506009546001600160a01b0316610497565b3480156105b857600080fd5b50600d5462010000900460ff16610420565b3480156105d657600080fd5b506103f96105e5366004613b05565b610ee4565b6103f96105f8366004613b52565b610f37565b34801561060957600080fd5b506105606106183660046139d3565b610f50565b34801561062957600080fd5b506103f9610ffa565b34801561063e57600080fd5b506103f961064d366004613c02565b6110bd565b34801561065e57600080fd5b506103f961066d366004613c02565b61125d565b34801561067e57600080fd5b506103f961068d366004613b05565b611441565b6103f96106a0366004613c6d565b61148e565b3480156106b157600080fd5b50600d54610420906301000000900460ff1681565b3480156106d257600080fd5b5061046a611a8b565b3480156106e757600080fd5b50600d5461042090610100900460ff1681565b34801561070657600080fd5b506103f9611b19565b34801561071b57600080fd5b50601654610497906001600160a01b031681565b34801561073b57600080fd5b5061049761074a3660046139ba565b611bcd565b34801561075b57600080fd5b5061046a6040518060400160405280601581526020017411995c985b199a5b19515e1a1a589a5d1a5bdb958d605a1b81525081565b34801561079c57600080fd5b506103f96107ab366004613cc8565b611c02565b3480156107bc57600080fd5b506103f96105f836600461393c565b3480156107d757600080fd5b506103f96107e636600461393c565b611f00565b3480156107f757600080fd5b5061056061080636600461393c565b611f8b565b34801561081757600080fd5b506103f9612011565b34801561082c57600080fd5b5061056061083b36600461393c565b6001600160a01b031660009081526015602052604090205490565b34801561086257600080fd5b506105606108713660046139ba565b60009081526010602052604090205490565b34801561088f57600080fd5b506108a361089e36600461393c565b612025565b60405161042c9190613d89565b3480156108bc57600080fd5b506103f96108cb366004613dcd565b612091565b3480156108dc57600080fd5b506006546001600160a01b0316610497565b3480156108fa57600080fd5b506103f9610909366004613e98565b6121c0565b34801561091a57600080fd5b5061056061092936600461393c565b60146020526000908152604090205481565b34801561094757600080fd5b5061046a61222f565b34801561095c57600080fd5b50600d546104209060ff1681565b34801561097657600080fd5b506103f9610985366004613ef9565b61223e565b34801561099657600080fd5b506103f96109a536600461393c565b612252565b3480156109b657600080fd5b506103f96122a3565b3480156109cb57600080fd5b506103f96109da366004613f30565b6122c0565b3480156109eb57600080fd5b506103f961230e565b348015610a0057600080fd5b5061046a610a0f3660046139ba565b6123af565b348015610a2057600080fd5b506103f9610a2f36600461393c565b6124bd565b348015610a4057600080fd5b5061046a6124e9565b348015610a5557600080fd5b50610420610a64366004613fab565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b348015610a9e57600080fd5b50610560610aad3660046139ba565b6000908152600f602052604090205490565b348015610acb57600080fd5b50610420610ada36600461393c565b60076020526000908152604090205460ff1681565b348015610afb57600080fd5b50600854610497906001600160a01b031681565b348015610b1b57600080fd5b506103f9610b2a36600461393c565b6124f6565b348015610b3b57600080fd5b50600d546104979064010000000090046001600160a01b031681565b348015610b6357600080fd5b50600e54610497906001600160a01b031681565b60006001600160e01b031982166380ac58cd60e01b1480610ba857506001600160e01b03198216635b5e139f60e01b145b80610bc357506301ffc9a760e01b6001600160e01b03198316145b92915050565b610bd161256f565b6001600160a01b03166000908152600760205260409020805460ff19169055565b606060008054610c0190613fde565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2d90613fde565b8015610c7a5780601f10610c4f57610100808354040283529160200191610c7a565b820191906000526020600020905b815481529060010190602001808311610c5d57829003601f168201915b5050505050905090565b6000610c8f826125c9565b506000908152600460205260409020546001600160a01b031690565b81610cb5816125ee565b610cbf83836126c0565b505050565b610ccc61256f565b600880546001600160a01b0319166001600160a01b0392909216919091179055565b610cf661256f565b6001600160a01b038116610d695760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a20636f737452656365696044820152737665725f206973207a65726f206164647265737360601b60648201526084016103f0565b600d80546001600160a01b0390921664010000000002640100000000600160c01b0319909216919091179055565b6040805180820190915260008082526020820152610db4826127d0565b610dd05760405162461bcd60e51b81526004016103f090614018565b50600090815260116020908152604091829020825180840190935280548352600101549082015290565b600d5460ff16610e615760405162461bcd60e51b815260206004820152602c60248201527f466572616c66696c6545786869626974696f6e56343a20746f6b656e2069732060448201526b6e6f74206275726e61626c6560a01b60648201526084016103f0565b60005b8151811015610ee057610e9033838381518110610e8357610e8361404f565b60200260200101516127ed565b610eac5760405162461bcd60e51b81526004016103f090614065565b610ece828281518110610ec157610ec161404f565b602002602001015161286c565b80610ed8816140c8565b915050610e64565b5050565b826001600160a01b0381163314610efe57610efe336125ee565b306001600160a01b03841603610f265760405162461bcd60e51b81526004016103f0906140e1565b610f31848484612942565b50505050565b6040516369bd111d60e11b815260040160405180910390fd5b6000610f5b83611f8b565b8210610fbd5760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b60648201526084016103f0565b6001600160a01b0383166000908152601260205260409020805483908110610fe757610fe761404f565b9060005260206000200154905092915050565b61100261256f565b600d546301000000900460ff161561102c5760405162461bcd60e51b81526004016103f09061413e565b600d5462010000900460ff16156110a25760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a205f73656c6c696e6720604482015273726571756972656420746f2062652066616c736560601b60648201526084016103f0565b6110aa612973565b600d805462ff0000191662010000179055565b6110c561256f565b8281146110e5576040516313086eff60e21b815260040160405180910390fd5b60005b838110156112565760008585838181106111045761110461404f565b9050602002016020810190611119919061393c565b6001600160a01b03160361114057604051630107349760e51b815260040160405180910390fd5b8282828181106111525761115261404f565b9050602002013560000361117957604051636745f8fb60e01b815260040160405180910390fd5b6000601460008787858181106111915761119161404f565b90506020020160208101906111a6919061393c565b6001600160a01b03166001600160a01b031681526020019081526020016000205411156111e6576040516328547bdf60e01b815260040160405180910390fd5b8282828181106111f8576111f861404f565b90506020020135601460008787858181106112155761121561404f565b905060200201602081019061122a919061393c565b6001600160a01b031681526020810191909152604001600020558061124e816140c8565b9150506110e8565b5050505050565b61126561256f565b828114611285576040516313086eff60e21b815260040160405180910390fd5b60005b838110156112565760008383838181106112a4576112a461404f565b90506020020160208101906112b9919061393c565b6001600160a01b0316036112e057604051630107349760e51b815260040160405180910390fd5b6000601460008585858181106112f8576112f861404f565b905060200201602081019061130d919061393c565b6001600160a01b03166001600160a01b0316815260200190815260200160002054111561134d576040516328547bdf60e01b815260040160405180910390fd5b601460008686848181106113635761136361404f565b9050602002016020810190611378919061393c565b6001600160a01b03166001600160a01b0316815260200190815260200160002054601460008585858181106113af576113af61404f565b90506020020160208101906113c4919061393c565b6001600160a01b03166001600160a01b0316815260200190815260200160002081905550601460008686848181106113fe576113fe61404f565b9050602002016020810190611413919061393c565b6001600160a01b03168152602081019190915260400160009081205580611439816140c8565b915050611288565b826001600160a01b038116331461145b5761145b336125ee565b306001600160a01b038416036114835760405162461bcd60e51b81526004016103f0906140e1565b610f318484846129ee565b600d5462010000900460ff166114b7576040516316851a3760e11b815260040160405180910390fd5b60006114c230611f8b565b90506114d460e0830160c084016141a4565b61ffff168110156114f857604051632d65aa3b60e11b815260040160405180910390fd5b61150182612a09565b6000463084604051602001611518939291906142f9565b60405160208183030381529060405280519060200120905061153c81878787612a67565b61155957604051638baa579f60e01b815260040160405180910390fd5b61157661156c608085016060860161393c565b8460800135612abf565b6115886101208401610100850161432c565b156115fa5760165460405163cdb1f66360e01b81526001600160a01b039091169063cdb1f663906115c3908990899089908990600401614349565b600060405180830381600087803b1580156115dd57600080fd5b505af11580156115f1573d6000803e3d6000fd5b5050505061161b565b8235341461161b57604051637e2897ef60e11b815260040160405180910390fd5b60208301358335101561164157604051637e2897ef60e11b815260040160405180910390fd5b60006116526020850135853561437b565b60a08501356000908152601760205260408120549192505b61167a60e0870160c088016141a4565b61ffff1681101561176c578161168f816127d0565b6116ac576040516352a7a53160e11b815260040160405180910390fd5b826116b6816140c8565b93503090506116c482611bcd565b6001600160a01b0316146116d8575061166a565b611702306116ec60808a0160608b0161393c565b8360405180602001604052806000815250612b12565b806117136080890160608a0161393c565b6001600160a01b03167fba8636482fa7bb52433c25b1cf79e47571bf179a48a271361b50bb78b1d63d7b896080013560405161175191815260200190565b60405180910390a381611763816140c8565b9250505061166a565b60a08601356000908152601760205260408120839055808061179160e08a018a61438e565b808060200260200160405190810160405280939291908181526020016000905b828210156117dd576117ce604083028601368190038101906143d7565b815260200190600101906117b1565b50505050509050600086905060005b8251811080156117fc5750600082115b156118ef576000601460008584815181106118195761181961404f565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000205490508060000361185657506118dd565b6000838210156118665781611868565b835b9050611874818761442d565b9550806014600087868151811061188d5761188d61404f565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002060008282546118c8919061437b565b909155506118d89050818561437b565b935050505b806118e7816140c8565b9150506117ec565b5080156119eb5760005b82518110156119e95760008382815181106119165761191661404f565b6020026020010151600001519050600061271085848151811061193b5761193b61404f565b602002602001015160200151856119529190614440565b61195c9190614457565b600d549091506001600160a01b0364010000000090910481169083160361199057611987818761442d565b955050506119d7565b61199a818861442d565b6040519097506001600160a01b0383169082156108fc029083906000818181858888f193505050501580156119d3573d6000803e3d6000fd5b5050505b806119e1816140c8565b9150506118f9565b505b6119f5838561442d565b611a0460208c01358c3561437b565b1015611a23576040516372ef2a9d60e01b815260040160405180910390fd5b6000611a30858c3561437b565b90508015611a7b57600d546040516401000000009091046001600160a01b0316906108fc8315029083906000818181858888f19350505050158015611a79573d6000803e3d6000fd5b505b5050505050505050505050505050565b600a8054611a9890613fde565b80601f0160208091040260200160405190810160405280929190818152602001828054611ac490613fde565b8015611b115780601f10611ae657610100808354040283529160200191611b11565b820191906000526020600020905b815481529060010190602001808311611af457829003601f168201915b505050505081565b611b2161256f565b600d546301000000900460ff1615611b4b5760405162461bcd60e51b81526004016103f09061413e565b600d5462010000900460ff16611bbf5760405162461bcd60e51b815260206004820152603360248201527f466572616c66696c6545786869626974696f6e56343a205f73656c6c696e6720604482015272726571756972656420746f206265207472756560681b60648201526084016103f0565b600d805462ff000019169055565b6000818152600260205260408120546001600160a01b031680610bc35760405162461bcd60e51b81526004016103f090614018565b611c0a61256f565b60008251118015611c1c575060008151115b611c9c5760405162461bcd60e51b815260206004820152604560248201527f466572616c66696c6545786869626974696f6e56343a2073657269657349647360448201527f206f7220726563697069656e74416464726573736573206c656e677468206973606482015264207a65726f60d81b608482015260a4016103f0565b8051825114611d285760405162461bcd60e51b815260206004820152604c60248201527f466572616c66696c6545786869626974696f6e56343a2073657269657349647360448201527f206c656e67746820697320646966666572656e742066726f6d2072656369706960648201526b656e7441646472657373657360a01b608482015260a4016103f0565b611d30611b19565b30600081815260126020908152604080832080548251818502810185019093528083529192909190830182828015611d8757602002820191906000526020600020905b815481526020019060010190808311611d73575b5050505050905060005b8151811015611e83576000828281518110611dae57611dae61404f565b602090810291909101810151600081815260118352604080822081518083019092528054825260010154938101939093529092505b87518161ffff161015611e6d57878161ffff1681518110611e0657611e0661404f565b6020026020010151826000015103611e5b576000878261ffff1681518110611e3057611e3061404f565b60200260200101519050611e5587828660405180602001604052806000815250612b12565b50611e6d565b80611e6581614479565b915050611de3565b5050508080611e7b906140c8565b915050611d91565b50611e8d82611f8b565b15610f315760405162461bcd60e51b815260206004820152603c60248201527f466572616c66696c6545786869626974696f6e56343a20546f6b656e20666f7260448201527f2073616c652062616c616e63652068617320746f206265207a65726f0000000060648201526084016103f0565b611f0861256f565b6001600160a01b038116611f695760405162461bcd60e51b815260206004820152602260248201527f45434453415369676e3a207369676e65725f206973207a65726f206164647265604482015261737360f01b60648201526084016103f0565b600980546001600160a01b0319166001600160a01b0392909216919091179055565b60006001600160a01b038216611ff55760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b60648201526084016103f0565b506001600160a01b031660009081526003602052604090205490565b61201961256f565b6120236000612b45565b565b6001600160a01b03811660009081526012602090815260409182902080548351818402810184019094528084526060939283018282801561208557602002820191906000526020600020905b815481526020019060010190808311612071575b50505050509050919050565b3360009081526007602052604090205460ff16806120b957506006546001600160a01b031633145b6120c257600080fd5b600d546301000000900460ff166121395760405162461bcd60e51b815260206004820152603560248201527f466572616c66696c6545786869626974696f6e56343a20636f6e747261637420604482015274191bd95cdb89dd08185b1b1bddc81d1bc81b5a5b9d605a1b60648201526084016103f0565b60005b81811015610cbf576121ae8383838181106121595761215961404f565b905060600201600001358484848181106121755761217561404f565b905060600201602001358585858181106121915761219161404f565b90506060020160400160208101906121a9919061393c565b612b97565b806121b8816140c8565b91505061213c565b6121c861256f565b60008151116122235760405162461bcd60e51b815260206004820152602160248201527f4552433732314d657461646174613a20626173655552495f20697320656d70746044820152607960f81b60648201526084016103f0565b600a610ee082826144e8565b606060018054610c0190613fde565b81612248816125ee565b610cbf8383612d19565b61225a61256f565b6001600160a01b0381166122815760405163e6c4247b60e01b815260040160405180910390fd5b601680546001600160a01b0319166001600160a01b0392909216919091179055565b6122ab61256f565b600d805463ff00000019169055612023610ffa565b836001600160a01b03811633146122da576122da336125ee565b306001600160a01b038516036123025760405162461bcd60e51b81526004016103f0906140e1565b61125685858585612d24565b61231661256f565b61231e611b19565b3060009081526012602090815260408083208054825181850281018501909352808352919290919083018282801561237557602002820191906000526020600020905b815481526020019060010190808311612361575b5050505050905060005b8151811015610ee05761239d828281518110610ec157610ec161404f565b806123a7816140c8565b91505061237f565b60606000600a80546123c090613fde565b90501161241e5760405162461bcd60e51b815260206004820152602660248201527f4552433732314d657461646174613a205f746f6b656e4261736555524920697360448201526520656d70747960d01b60648201526084016103f0565b612427826127d0565b61248b5760405162461bcd60e51b815260206004820152602f60248201527f4552433732314d657461646174613a2055524920717565727920666f72206e6f60448201526e3732bc34b9ba32b73a103a37b5b2b760891b60648201526084016103f0565b600a61249683612d56565b6040516020016124a79291906145a7565b6040516020818303038152906040529050919050565b6124c561256f565b6001600160a01b03166000908152600760205260409020805460ff19166001179055565b600b8054611a9890613fde565b6124fe61256f565b6001600160a01b0381166125635760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103f0565b61256c81612b45565b50565b6006546001600160a01b031633146120235760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103f0565b6125d2816127d0565b61256c5760405162461bcd60e51b81526004016103f090614018565b6008546001600160a01b03163b1561256c57600854604051633185c44d60e21b81523060048201526001600160a01b0383811660248301529091169063c617113490604401602060405180830381865afa158015612650573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612674919061463b565b61256c5760405162461bcd60e51b815260206004820152601760248201527f6f70657261746f72206973206e6f7420616c6c6f77656400000000000000000060448201526064016103f0565b60006126cb82611bcd565b9050806001600160a01b0316836001600160a01b0316036127385760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084016103f0565b336001600160a01b038216148061275457506127548133610a64565b6127c65760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c00000060648201526084016103f0565b610cbf8383612de8565b6000908152600260205260409020546001600160a01b0316151590565b6000806127f983611bcd565b9050806001600160a01b0316846001600160a01b0316148061284057506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b806128645750836001600160a01b031661285984610c84565b6001600160a01b0316145b949350505050565b612875816127d0565b6128915760405162461bcd60e51b81526004016103f090614018565b600081815260116020908152604080832081518083018352815480825260019283015482860152855260109093529083208054929391929091906128d690849061437b565b925050819055506001600c60008282546128f0919061437b565b909155505060008281526011602052604081208181556001015561291382612e56565b60405182907fbde7938970372996ff103863625e348ef2bf8f38a5b02181be75aafef17c23d590600090a25050565b61294c33826127ed565b6129685760405162461bcd60e51b81526004016103f090614065565b610cbf838383612ef9565b600061297e30611f8b565b90506000811161256c5760405162461bcd60e51b815260206004820152603560248201527f466572616c66696c6545786869626974696f6e56343a204e6f20746f6b656e206044820152741bdddb995908189e481d1a194818dbdb9d1c9858dd605a1b60648201526084016103f0565b610cbf838383604051806020016040528060008152506122c0565b4281604001351161256c5760405162461bcd60e51b815260206004820152602260248201527f466572616c66696c6553616c65446174613a2073616c65206973206578706972604482015261195960f21b60648201526084016103f0565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c859052603c81208190612aa49084878761306a565b6009546001600160a01b039081169116149695505050505050565b6001600160a01b0382166000908152601560205260409020805460018101909155818114610cbf576040516301d4b62360e61b81526001600160a01b0384166004820152602481018290526044016103f0565b612b1d848484612ef9565b612b2984848484613092565b610f315760405162461bcd60e51b81526004016103f090614658565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000838152600f60205260409020541515612bb184612d56565b604051602001612bc191906146aa565b60405160208183030381529060405290612bee5760405162461bcd60e51b81526004016103f091906139a7565b506000838152600f602090815260408083205460109092529091205410612c695760405162461bcd60e51b815260206004820152602960248201527f466572616c66696c6545786869626974696f6e56343a206e6f20736c6f747320604482015268617661696c61626c6560b81b60648201526084016103f0565b6001600c6000828254612c7c919061442d565b90915550506000838152601060205260408120805460019290612ca090849061442d565b9091555050604080518082018252848152602080820185815260008681526011909252929020905181559051600190910155612cdc8183613190565b8183826001600160a01b03167f407d7da1d3b2b1871fbfa2b5b1c4657a3cc5711d3023c552798551c7ee301eea60405160405180910390a4505050565b610ee033838361330b565b612d2e33836127ed565b612d4a5760405162461bcd60e51b81526004016103f090614065565b610f3184848484612b12565b60606000612d63836133d9565b60010190506000816001600160401b03811115612d8257612d826139fd565b6040519080825280601f01601f191660200182016040528015612dac576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084612db657509392505050565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190612e1d82611bcd565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000612e6182611bcd565b9050612e718160008460016134b1565b612e7a82611bcd565b600083815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526003845282852080546000190190558785526002909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b826001600160a01b0316612f0c82611bcd565b6001600160a01b031614612f325760405162461bcd60e51b81526004016103f090614707565b6001600160a01b038216612f945760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b60648201526084016103f0565b612fa183838360016134b1565b826001600160a01b0316612fb482611bcd565b6001600160a01b031614612fda5760405162461bcd60e51b81526004016103f090614707565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b600080600061307b878787876135c8565b915091506130888161368c565b5095945050505050565b60006001600160a01b0384163b1561318857604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906130d690339089908890889060040161474c565b6020604051808303816000875af1925050508015613111575060408051601f3d908101601f1916820190925261310e9181019061477f565b60015b61316e573d80801561313f576040519150601f19603f3d011682016040523d82523d6000602084013e613144565b606091505b5080516000036131665760405162461bcd60e51b81526004016103f090614658565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050612864565b506001612864565b6001600160a01b0382166131e65760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f206164647265737360448201526064016103f0565b6131ef816127d0565b1561323c5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016103f0565b61324a6000838360016134b1565b613253816127d0565b156132a05760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016103f0565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b816001600160a01b0316836001600160a01b03160361336c5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c65720000000000000060448201526064016103f0565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106134185772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310613444576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061346257662386f26fc10000830492506010015b6305f5e100831061347a576305f5e100830492506008015b612710831061348e57612710830492506004015b606483106134a0576064830492506002015b600a8310610bc35760010192915050565b60018111156135205760405162461bcd60e51b815260206004820152603560248201527f455243373231456e756d657261626c653a20636f6e7365637574697665207472604482015274185b9cd9995c9cc81b9bdd081cdd5c1c1bdc9d1959605a1b60648201526084016103f0565b816001600160a01b0385161580159061354b5750836001600160a01b0316856001600160a01b031614155b1561355a5761355a85826137d6565b6001600160a01b038416158015906135845750846001600160a01b0316846001600160a01b031614155b15611256576001600160a01b038416600090815260126020908152604080832080546001810182559084528284208101859055848452601390925290912055611256565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156135ff5750600090506003613683565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613653573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661367c57600060019250925050613683565b9150600090505b94509492505050565b60008160048111156136a0576136a061479c565b036136a85750565b60018160048111156136bc576136bc61479c565b036137095760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016103f0565b600281600481111561371d5761371d61479c565b0361376a5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016103f0565b600381600481111561377e5761377e61479c565b0361256c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016103f0565b600060016137e384611f8b565b6137ed919061437b565b600083815260136020526040902054909150808214613894576001600160a01b03841660009081526012602052604081208054849081106138305761383061404f565b906000526020600020015490508060126000876001600160a01b03166001600160a01b0316815260200190815260200160002083815481106138745761387461404f565b600091825260208083209091019290925591825260139052604090208190555b60008381526013602090815260408083208390556001600160a01b0387168352601290915290208054806138ca576138ca6147b2565b6001900381819060005260206000200160009055905550505050565b6001600160e01b03198116811461256c57600080fd5b60006020828403121561390e57600080fd5b8135613919816138e6565b9392505050565b80356001600160a01b038116811461393757600080fd5b919050565b60006020828403121561394e57600080fd5b61391982613920565b60005b8381101561397257818101518382015260200161395a565b50506000910152565b60008151808452613993816020860160208601613957565b601f01601f19169290920160200192915050565b602081526000613919602083018461397b565b6000602082840312156139cc57600080fd5b5035919050565b600080604083850312156139e657600080fd5b6139ef83613920565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715613a3b57613a3b6139fd565b604052919050565b60006001600160401b03821115613a5c57613a5c6139fd565b5060051b60200190565b600082601f830112613a7757600080fd5b81356020613a8c613a8783613a43565b613a13565b82815260059290921b84018101918181019086841115613aab57600080fd5b8286015b84811015613ac65780358352918301918301613aaf565b509695505050505050565b600060208284031215613ae357600080fd5b81356001600160401b03811115613af957600080fd5b61286484828501613a66565b600080600060608486031215613b1a57600080fd5b613b2384613920565b9250613b3160208501613920565b9150604084013590509250925092565b803560ff8116811461393757600080fd5b60008060008060808587031215613b6857600080fd5b8435935060208501359250613b7f60408601613b41565b915060608501356001600160401b03811115613b9a57600080fd5b850160e08188031215613bac57600080fd5b939692955090935050565b60008083601f840112613bc957600080fd5b5081356001600160401b03811115613be057600080fd5b6020830191508360208260051b8501011115613bfb57600080fd5b9250929050565b60008060008060408587031215613c1857600080fd5b84356001600160401b0380821115613c2f57600080fd5b613c3b88838901613bb7565b90965094506020870135915080821115613c5457600080fd5b50613c6187828801613bb7565b95989497509550505050565b60008060008060808587031215613c8357600080fd5b8435935060208501359250613c9a60408601613b41565b915060608501356001600160401b03811115613cb557600080fd5b85016101208188031215613bac57600080fd5b60008060408385031215613cdb57600080fd5b82356001600160401b0380821115613cf257600080fd5b613cfe86838701613a66565b9350602091508185013581811115613d1557600080fd5b85019050601f81018613613d2857600080fd5b8035613d36613a8782613a43565b81815260059190911b82018301908381019088831115613d5557600080fd5b928401925b82841015613d7a57613d6b84613920565b82529284019290840190613d5a565b80955050505050509250929050565b6020808252825182820181905260009190848201906040850190845b81811015613dc157835183529284019291840191600101613da5565b50909695505050505050565b60008060208385031215613de057600080fd5b82356001600160401b0380821115613df757600080fd5b818501915085601f830112613e0b57600080fd5b813581811115613e1a57600080fd5b866020606083028501011115613e2f57600080fd5b60209290920196919550909350505050565b60006001600160401b03831115613e5a57613e5a6139fd565b613e6d601f8401601f1916602001613a13565b9050828152838383011115613e8157600080fd5b828260208301376000602084830101529392505050565b600060208284031215613eaa57600080fd5b81356001600160401b03811115613ec057600080fd5b8201601f81018413613ed157600080fd5b61286484823560208401613e41565b801515811461256c57600080fd5b803561393781613ee0565b60008060408385031215613f0c57600080fd5b613f1583613920565b91506020830135613f2581613ee0565b809150509250929050565b60008060008060808587031215613f4657600080fd5b613f4f85613920565b9350613f5d60208601613920565b92506040850135915060608501356001600160401b03811115613f7f57600080fd5b8501601f81018713613f9057600080fd5b613f9f87823560208401613e41565b91505092959194509250565b60008060408385031215613fbe57600080fd5b613fc783613920565b9150613fd560208401613920565b90509250929050565b600181811c90821680613ff257607f821691505b60208210810361401257634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526018908201527f4552433732313a20696e76616c696420746f6b656e2049440000000000000000604082015260600190565b634e487b7160e01b600052603260045260246000fd5b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b6000600182016140da576140da6140b2565b5060010190565b6020808252603e908201527f466572616c66696c6545786869626974696f6e56343a20436f6e74726163742060408201527f69736e277420616c6c6f77656420746f207265636569766520746f6b656e0000606082015260800190565b60208082526034908201527f466572616c66696c6545786869626974696f6e56343a206d696e7461626c6520604082015273726571756972656420746f2062652066616c736560601b606082015260800190565b803561ffff8116811461393757600080fd5b6000602082840312156141b657600080fd5b61391982614192565b6000808335601e198436030181126141d657600080fd5b83016020810192503590506001600160401b038111156141f557600080fd5b8060061b3603821315613bfb57600080fd5b8183526000602080850194508260005b8581101561424d576001600160a01b0361423083613920565b168752818301358388015260409687019690910190600101614217565b509495945050505050565b80358252602080820135908301526040808201359083015260006101206001600160a01b0361428960608501613920565b1660608501526080830135608085015260a083013560a08501526142af60c08401614192565b61ffff1660c08501526142c560e08401846141bf565b8260e08701526142d88387018284614207565b925050506101006142ea818501613eee565b15159401939093525090919050565b8381526001600160a01b038316602082015260606040820181905260009061432390830184614258565b95945050505050565b60006020828403121561433e57600080fd5b813561391981613ee0565b84815283602082015260ff831660408201526080606082015260006143716080830184614258565b9695505050505050565b81810381811115610bc357610bc36140b2565b6000808335601e198436030181126143a557600080fd5b8301803591506001600160401b038211156143bf57600080fd5b6020019150600681901b3603821315613bfb57600080fd5b6000604082840312156143e957600080fd5b604051604081018181106001600160401b038211171561440b5761440b6139fd565b60405261441783613920565b8152602083013560208201528091505092915050565b80820180821115610bc357610bc36140b2565b8082028115828204841417610bc357610bc36140b2565b60008261447457634e487b7160e01b600052601260045260246000fd5b500490565b600061ffff808316818103614490576144906140b2565b6001019392505050565b601f821115610cbf57600081815260208120601f850160051c810160208610156144c15750805b601f850160051c820191505b818110156144e0578281556001016144cd565b505050505050565b81516001600160401b03811115614501576145016139fd565b6145158161450f8454613fde565b8461449a565b602080601f83116001811461454a57600084156145325750858301515b600019600386901b1c1916600185901b1785556144e0565b600085815260208120601f198616915b828110156145795788860151825594840194600190910190840161455a565b50858210156145975787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008084546145b581613fde565b600182811680156145cd57600181146145e257614611565b60ff1984168752821515830287019450614611565b8860005260208060002060005b858110156146085781548a8201529084019082016145ef565b50505082870194505b50602f60f81b84528651925061462d8382860160208a01613957565b919092010195945050505050565b60006020828403121561464d57600080fd5b815161391981613ee0565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b7f466572616c66696c6545786869626974696f6e56343a2073657269657349642081526e03237b2b9b713ba1032bc34b9ba1d1608d1b6020820152600082516146fa81602f850160208701613957565b91909101602f0192915050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906143719083018461397b565b60006020828403121561479157600080fd5b8151613919816138e6565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603160045260246000fdfea264697066735822122002320218b50a6e9c7c892dd8e30b9bdcfbaf0569f467f944b26ba269ee301a9f64736f6c63430008110033",
}

// FeralfileExhibitionV42ABI is the input ABI used to generate the binding from.
// Deprecated: Use FeralfileExhibitionV42MetaData.ABI instead.
var FeralfileExhibitionV42ABI = FeralfileExhibitionV42MetaData.ABI

// FeralfileExhibitionV42Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeralfileExhibitionV42MetaData.Bin instead.
var FeralfileExhibitionV42Bin = FeralfileExhibitionV42MetaData.Bin

// DeployFeralfileExhibitionV42 deploys a new Ethereum contract, binding an instance of FeralfileExhibitionV42 to it.
func DeployFeralfileExhibitionV42(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, burnable_ bool, bridgeable_ bool, signer_ common.Address, vault_ common.Address, costReceiver_ common.Address, contractURI_ string, seriesIds_ []*big.Int, seriesMaxSupplies_ []*big.Int, seriesNextPurchasableTokenIds_ []*big.Int) (common.Address, *types.Transaction, *FeralfileExhibitionV42, error) {
	parsed, err := FeralfileExhibitionV42MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeralfileExhibitionV42Bin), backend, name_, symbol_, burnable_, bridgeable_, signer_, vault_, costReceiver_, contractURI_, seriesIds_, seriesMaxSupplies_, seriesNextPurchasableTokenIds_)
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

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) Nonces(owner common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.Nonces(&_FeralfileExhibitionV42.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV42.Contract.Nonces(&_FeralfileExhibitionV42.CallOpts, owner)
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

// SetVault is a free data retrieval call binding the contract method 0x6817031b.
//
// Solidity: function setVault(address ) pure returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) SetVault(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "setVault", arg0)

	if err != nil {
		return err
	}

	return err

}

// SetVault is a free data retrieval call binding the contract method 0x6817031b.
//
// Solidity: function setVault(address ) pure returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetVault(arg0 common.Address) error {
	return _FeralfileExhibitionV42.Contract.SetVault(&_FeralfileExhibitionV42.CallOpts, arg0)
}

// SetVault is a free data retrieval call binding the contract method 0x6817031b.
//
// Solidity: function setVault(address ) pure returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) SetVault(arg0 common.Address) error {
	return _FeralfileExhibitionV42.Contract.SetVault(&_FeralfileExhibitionV42.CallOpts, arg0)
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

// VaultV2 is a free data retrieval call binding the contract method 0x5eb9bad6.
//
// Solidity: function vaultV2() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Caller) VaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV42.contract.Call(opts, &out, "vaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultV2 is a free data retrieval call binding the contract method 0x5eb9bad6.
//
// Solidity: function vaultV2() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) VaultV2() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.VaultV2(&_FeralfileExhibitionV42.CallOpts)
}

// VaultV2 is a free data retrieval call binding the contract method 0x5eb9bad6.
//
// Solidity: function vaultV2() view returns(address)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42CallerSession) VaultV2() (common.Address, error) {
	return _FeralfileExhibitionV42.Contract.VaultV2(&_FeralfileExhibitionV42.CallOpts)
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
// Solidity: function buyArtworks(bytes32 , bytes32 , uint8 , (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) ) payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) BuyArtworks(opts *bind.TransactOpts, arg0 [32]byte, arg1 [32]byte, arg2 uint8, arg3 IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "buyArtworks", arg0, arg1, arg2, arg3)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 , bytes32 , uint8 , (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) ) payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) BuyArtworks(arg0 [32]byte, arg1 [32]byte, arg2 uint8, arg3 IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.BuyArtworks(&_FeralfileExhibitionV42.TransactOpts, arg0, arg1, arg2, arg3)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 , bytes32 , uint8 , (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) ) payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) BuyArtworks(arg0 [32]byte, arg1 [32]byte, arg2 uint8, arg3 IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.BuyArtworks(&_FeralfileExhibitionV42.TransactOpts, arg0, arg1, arg2, arg3)
}

// BuyBulkArtworks is a paid mutator transaction binding the contract method 0x4bda5d89.
//
// Solidity: function buyBulkArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256,uint256,uint16,(address,uint256)[],bool) saleData_) payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) BuyBulkArtworks(opts *bind.TransactOpts, r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataV2SaleDataV2) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "buyBulkArtworks", r_, s_, v_, saleData_)
}

// BuyBulkArtworks is a paid mutator transaction binding the contract method 0x4bda5d89.
//
// Solidity: function buyBulkArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256,uint256,uint16,(address,uint256)[],bool) saleData_) payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) BuyBulkArtworks(r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataV2SaleDataV2) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.BuyBulkArtworks(&_FeralfileExhibitionV42.TransactOpts, r_, s_, v_, saleData_)
}

// BuyBulkArtworks is a paid mutator transaction binding the contract method 0x4bda5d89.
//
// Solidity: function buyBulkArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256,uint256,uint16,(address,uint256)[],bool) saleData_) payable returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) BuyBulkArtworks(r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataV2SaleDataV2) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.BuyBulkArtworks(&_FeralfileExhibitionV42.TransactOpts, r_, s_, v_, saleData_)
}

// MintArtworks is a paid mutator transaction binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] data) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) MintArtworks(opts *bind.TransactOpts, data []FeralfileExhibitionV4MintData) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "mintArtworks", data)
}

// MintArtworks is a paid mutator transaction binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] data) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) MintArtworks(data []FeralfileExhibitionV4MintData) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.MintArtworks(&_FeralfileExhibitionV42.TransactOpts, data)
}

// MintArtworks is a paid mutator transaction binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] data) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) MintArtworks(data []FeralfileExhibitionV4MintData) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.MintArtworks(&_FeralfileExhibitionV42.TransactOpts, data)
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

// SetVaultV2 is a paid mutator transaction binding the contract method 0xa74cebab.
//
// Solidity: function setVaultV2(address vault_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Transactor) SetVaultV2(opts *bind.TransactOpts, vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.contract.Transact(opts, "setVaultV2", vault_)
}

// SetVaultV2 is a paid mutator transaction binding the contract method 0xa74cebab.
//
// Solidity: function setVaultV2(address vault_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Session) SetVaultV2(vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetVaultV2(&_FeralfileExhibitionV42.TransactOpts, vault_)
}

// SetVaultV2 is a paid mutator transaction binding the contract method 0xa74cebab.
//
// Solidity: function setVaultV2(address vault_) returns()
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42TransactorSession) SetVaultV2(vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV42.Contract.SetVaultV2(&_FeralfileExhibitionV42.TransactOpts, vault_)
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

// FeralfileExhibitionV42BuyArtworkV2Iterator is returned from FilterBuyArtworkV2 and is used to iterate over the raw logs and unpacked data for BuyArtworkV2 events raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42BuyArtworkV2Iterator struct {
	Event *FeralfileExhibitionV42BuyArtworkV2 // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV42BuyArtworkV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV42BuyArtworkV2)
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
		it.Event = new(FeralfileExhibitionV42BuyArtworkV2)
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
func (it *FeralfileExhibitionV42BuyArtworkV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV42BuyArtworkV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV42BuyArtworkV2 represents a BuyArtworkV2 event raised by the FeralfileExhibitionV42 contract.
type FeralfileExhibitionV42BuyArtworkV2 struct {
	Buyer   common.Address
	TokenId *big.Int
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyArtworkV2 is a free log retrieval operation binding the contract event 0xba8636482fa7bb52433c25b1cf79e47571bf179a48a271361b50bb78b1d63d7b.
//
// Solidity: event BuyArtworkV2(address indexed buyer, uint256 indexed tokenId, uint256 nonce)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) FilterBuyArtworkV2(opts *bind.FilterOpts, buyer []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV42BuyArtworkV2Iterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.FilterLogs(opts, "BuyArtworkV2", buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV42BuyArtworkV2Iterator{contract: _FeralfileExhibitionV42.contract, event: "BuyArtworkV2", logs: logs, sub: sub}, nil
}

// WatchBuyArtworkV2 is a free log subscription operation binding the contract event 0xba8636482fa7bb52433c25b1cf79e47571bf179a48a271361b50bb78b1d63d7b.
//
// Solidity: event BuyArtworkV2(address indexed buyer, uint256 indexed tokenId, uint256 nonce)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) WatchBuyArtworkV2(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV42BuyArtworkV2, buyer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV42.contract.WatchLogs(opts, "BuyArtworkV2", buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV42BuyArtworkV2)
				if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "BuyArtworkV2", log); err != nil {
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

// ParseBuyArtworkV2 is a log parse operation binding the contract event 0xba8636482fa7bb52433c25b1cf79e47571bf179a48a271361b50bb78b1d63d7b.
//
// Solidity: event BuyArtworkV2(address indexed buyer, uint256 indexed tokenId, uint256 nonce)
func (_FeralfileExhibitionV42 *FeralfileExhibitionV42Filterer) ParseBuyArtworkV2(log types.Log) (*FeralfileExhibitionV42BuyArtworkV2, error) {
	event := new(FeralfileExhibitionV42BuyArtworkV2)
	if err := _FeralfileExhibitionV42.contract.UnpackLog(event, "BuyArtworkV2", log); err != nil {
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
