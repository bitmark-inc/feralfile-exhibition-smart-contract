// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feralfilev4_4

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

// FeralfileExhibitionV44RendererTokenData is an auto generated low-level Go binding around an user-defined struct.
type FeralfileExhibitionV44RendererTokenData struct {
	ImageURI   string
	TextureURI string
	TokenName  string
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

// FeralfileExhibitionV44MetaData contains all meta data concerning the FeralfileExhibitionV44 contract.
var FeralfileExhibitionV44MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"burnable_\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"bridgeable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"costReceiver_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"seriesIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"seriesMaxSupplies_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdvanceAddressAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrEmptyArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrEmptyBytes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrEmptyString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrRendererBlobTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"ErrSeriesDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrSeriesHasNoRenderer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ErrTokenDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrUnsupportedCharacters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAddressesAndAmounts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WriteError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BurnArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BuyArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"DeleteSeriesRenderer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NewArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"textureURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structFeralfileExhibitionV4_4.RendererTokenData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"NewRendererTokenData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NewSeriesName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"pointers\",\"type\":\"address[]\"}],\"name\":\"NewSeriesRenderer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHUNK_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OperatorFilterRegistry\",\"outputs\":[{\"internalType\":\"contractIOperatorFilterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RENDERER_BLOB_MAX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"addTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"advances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burnArtworks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"internalType\":\"structIFeralfileSaleData.RevenueShare[][]\",\"name\":\"revenueShares\",\"type\":\"tuple[][]\"},{\"internalType\":\"bool\",\"name\":\"payByVaultContract\",\"type\":\"bool\"}],\"internalType\":\"structIFeralfileSaleData.SaleData\",\"name\":\"saleData_\",\"type\":\"tuple\"}],\"name\":\"buyArtworks\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"codeVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"costReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getArtwork\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structFeralfileExhibitionV4.Artwork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structFeralfileExhibitionV4.MintData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"mintArtworks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"removeTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rendererTokenData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"textureURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oldAddresses_\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"newAddresses_\",\"type\":\"address[]\"}],\"name\":\"replaceAdvanceAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selling\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"seriesMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seriesNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"seriesTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"setAdvanceSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"costReceiver_\",\"type\":\"address\"}],\"name\":\"setCostReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setTokenBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopSaleAndBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"seriesIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"recipientAddresses\",\"type\":\"address[]\"}],\"name\":\"stopSaleAndTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorFilterRegisterAddress\",\"type\":\"address\"}],\"name\":\"updateOperatorFilterRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIFeralfileVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blob\",\"type\":\"bytes\"}],\"name\":\"setSeriesRenderer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"getSeriesRenderer\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"seriesIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"names\",\"type\":\"string[]\"}],\"name\":\"setSeriesNames\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"textureURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"internalType\":\"structFeralfileExhibitionV4_4.RendererTokenData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"setRendererTokenData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600880546001600160a01b0319166daaeb6d7670e522a718067333cd4e179055600d805463ff000000191663010000001790553480156200004457600080fd5b5060405162006c2c38038062006c2c83398101604081905262000067916200094b565b8989898989898989898989898989898989898989858a8a60006200008c838262000b1a565b5060016200009b828262000b1a565b505050620000b8620000b26200076960201b60201c565b6200076d565b6008546001600160a01b03163b156200014557600854604051633e9f1edf60e11b8152306004820152733cc6cdda760b79bafa08df41ecfa224f810dceb660248201526001600160a01b0390911690637d3e3dbe90604401600060405180830381600087803b1580156200012b57600080fd5b505af115801562000140573d6000803e3d6000fd5b505050505b6001600160a01b038116620001ac5760405162461bcd60e51b815260206004820152602260248201527f45434453415369676e3a207369676e65725f206973207a65726f206164647265604482015261737360f01b60648201526084015b60405180910390fd5b600980546001600160a01b0319166001600160a01b039290921691909117905589516200022a5760405162461bcd60e51b815260206004820152602560248201527f466572616c66696c6545786869626974696f6e56343a206e616d655f20697320604482015264656d70747960d81b6064820152608401620001a3565b60008951116200028d5760405162461bcd60e51b815260206004820152602760248201527f466572616c66696c6545786869626974696f6e56343a2073796d626f6c5f20696044820152667320656d70747960c81b6064820152608401620001a3565b6001600160a01b0385166200030b5760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a207661756c744164647260448201527f6573735f206973207a65726f20616464726573730000000000000000000000006064820152608401620001a3565b6001600160a01b038416620003895760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a20636f7374526563656960448201527f7665725f206973207a65726f20616464726573730000000000000000000000006064820152608401620001a3565b6000835111620003f15760405162461bcd60e51b815260206004820152602c60248201527f466572616c66696c6545786869626974696f6e56343a20636f6e74726163745560448201526b52495f20697320656d70747960a01b6064820152608401620001a3565b6000825111620004575760405162461bcd60e51b815260206004820152602a60248201527f466572616c66696c6545786869626974696f6e56343a207365726965734964736044820152695f20697320656d70747960b01b6064820152608401620001a3565b6000815111620004c55760405162461bcd60e51b815260206004820152603260248201527f466572616c66696c6545786869626974696f6e56343a205f7365726965734d6160448201527178537570706c69657320697320656d70747960701b6064820152608401620001a3565b8051825114620005585760405162461bcd60e51b815260206004820152605160248201527f466572616c66696c6545786869626974696f6e56343a207365726965734d617860448201527f537570706c6965735f20616e64207365726965734964735f206c656e6774687360648201527020617265206e6f74207468652073616d6560781b608482015260a401620001a3565b600d805461ffff191689151561ff001916176101008915150217600160201b600160c01b0319166401000000006001600160a01b038781169190910291909117909155600e80546001600160a01b031916918716919091179055600b620005c0848262000b1a565b5060005b82518110156200074457600f6000848381518110620005e757620005e762000be6565b6020026020010151815260200190815260200160002054600014620006615760405162461bcd60e51b815260206004820152602960248201527f466572616c66696c6545786869626974696f6e56343a206475706c6963617465604482015268081cd95c9a595cd25960ba1b6064820152608401620001a3565b600082828151811062000678576200067862000be6565b602002602001015111620006de5760405162461bcd60e51b815260206004820152602660248201527f466572616c66696c6545786869626974696f6e56343a207a65726f206d617820604482015265737570706c7960d01b6064820152608401620001a3565b818181518110620006f357620006f362000be6565b6020026020010151600f600085848151811062000714576200071462000be6565b602002602001015181526020019081526020016000208190555080806200073b9062000bfc565b915050620005c4565b5050505050505050505050505050505050505050505050505050505050505062000c24565b3390565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620008005762000800620007bf565b604052919050565b600082601f8301126200081a57600080fd5b81516001600160401b03811115620008365762000836620007bf565b60206200084c601f8301601f19168201620007d5565b82815285828487010111156200086157600080fd5b60005b838110156200088157858101830151828201840152820162000864565b506000928101909101919091529392505050565b80518015158114620008a657600080fd5b919050565b80516001600160a01b0381168114620008a657600080fd5b600082601f830112620008d557600080fd5b815160206001600160401b03821115620008f357620008f3620007bf565b8160051b62000904828201620007d5565b92835284810182019282810190878511156200091f57600080fd5b83870192505b84831015620009405782518252918301919083019062000925565b979650505050505050565b6000806000806000806000806000806101408b8d0312156200096c57600080fd5b8a516001600160401b03808211156200098457600080fd5b620009928e838f0162000808565b9b5060208d0151915080821115620009a957600080fd5b620009b78e838f0162000808565b9a50620009c760408e0162000895565b9950620009d760608e0162000895565b9850620009e760808e01620008ab565b9750620009f760a08e01620008ab565b965062000a0760c08e01620008ab565b955060e08d015191508082111562000a1e57600080fd5b62000a2c8e838f0162000808565b94506101008d015191508082111562000a4457600080fd5b62000a528e838f01620008c3565b93506101208d015191508082111562000a6a57600080fd5b5062000a798d828e01620008c3565b9150509295989b9194979a5092959850565b600181811c9082168062000aa057607f821691505b60208210810362000ac157634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000b1557600081815260208120601f850160051c8101602086101562000af05750805b601f850160051c820191505b8181101562000b115782815560010162000afc565b5050505b505050565b81516001600160401b0381111562000b365762000b36620007bf565b62000b4e8162000b47845462000a8b565b8462000ac7565b602080601f83116001811462000b86576000841562000b6d5750858301515b600019600386901b1c1916600185901b17855562000b11565b600085815260208120601f198616915b8281101562000bb75788860151825594840194600190910190840162000b96565b508582101562000bd65787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b60006001820162000c1d57634e487b7160e01b600052601160045260246000fd5b5060010190565b615ff88062000c346000396000f3fe60806040526004361061039b5760003560e01c806370a08231116101dc578063b88d4fde11610102578063eb5c60f2116100a0578063f4e638be1161006f578063f4e638be14610b93578063fa35e4e214610bbb578063fbfa77cf14610bdb578063fe43658d14610bfb57600080fd5b8063eb5c60f214610af6578063eee608a414610b23578063f07e7fd014610b53578063f2fde38b14610b7357600080fd5b8063dc78ac1c116100dc578063dc78ac1c14610a62578063e8a3d48514610a82578063e91e13a914610a97578063e985e9c514610aad57600080fd5b8063b88d4fde14610a0d578063b9b8311a14610a2d578063c87b56dd14610a4257600080fd5b80638ef79e911161017a578063a07c7ce411610149578063a07c7ce41461099e578063a22cb465146109b8578063ab221792146109d8578063b66a0e5d146109f857600080fd5b80638ef79e9114610925578063926ce44e1461094557806395d89b41146109725780639cd7d65c1461098757600080fd5b80637f06ee06116101b65780637f06ee061461088d5780638462151c146108ba5780638cba1c67146108e75780638da5cb5b1461090757600080fd5b806370a0823114610838578063715018a614610858578063771ac3031461086d57600080fd5b806332f37c2e116102c1578063504c05ea1161025f57806363e602301161022e57806363e602301461079757806365a46e08146107d85780636817031b146107f85780636c19e7831461081857600080fd5b8063504c05ea14610714578063530da8ef1461074357806355367ba9146107625780636352211e1461077757600080fd5b806341a5626a1161029b57806341a5626a1461069e57806342842e0e146106be5780634bf365df146106de5780634e99b800146106ff57600080fd5b806332f37c2e1461064957806333e364cb146106695780633c352b0d1461067e57600080fd5b8063167ddf6e1161033957806323aed2281161030857806323aed228146105d857806323b872dd146105f65780632977e4b3146106165780632f745c591461062957600080fd5b8063167ddf6e1461053b57806318160ddd1461057657806321fe0c641461059a578063238ac933146105ba57600080fd5b8063081812fc11610375578063081812fc146104a3578063095ea7b3146104db578063114ba8ee146104fb5780631623528f1461051b57600080fd5b806301ffc9a71461042c578063031205061461046157806306fdde031461048157600080fd5b3661042757600e546001600160a01b031633146104255760405162461bcd60e51b815260206004820152603c60248201527f466572616c66696c6545786869626974696f6e56343a206f6e6c79206163636560448201527f70742066756e642066726f6d207661756c7420636f6e74726163742e0000000060648201526084015b60405180910390fd5b005b600080fd5b34801561043857600080fd5b5061044c6104473660046149dd565b610c1b565b60405190151581526020015b60405180910390f35b34801561046d57600080fd5b5061042561047c366004614a1d565b610c6d565b34801561048d57600080fd5b50610496610c96565b6040516104589190614a88565b3480156104af57600080fd5b506104c36104be366004614a9b565b610d28565b6040516001600160a01b039091168152602001610458565b3480156104e757600080fd5b506104256104f6366004614ab4565b610d4f565b34801561050757600080fd5b50610425610516366004614a1d565b610d68565b34801561052757600080fd5b50610425610536366004614a1d565b610d92565b34801561054757600080fd5b5061055b610556366004614a9b565b610e3b565b60408051825181526020928301519281019290925201610458565b34801561058257600080fd5b5061058c600c5481565b604051908152602001610458565b3480156105a657600080fd5b506104256105b5366004614bb2565b610e9e565b3480156105c657600080fd5b506009546001600160a01b03166104c3565b3480156105e457600080fd5b50600d5462010000900460ff1661044c565b34801561060257600080fd5b50610425610611366004614be6565b610f88565b610425610624366004614c22565b610fdb565b34801561063557600080fd5b5061058c610644366004614ab4565b611618565b34801561065557600080fd5b50610425610664366004614cda565b6116c2565b34801561067557600080fd5b50610425611851565b34801561068a57600080fd5b50610425610699366004614cda565b611914565b3480156106aa57600080fd5b506104256106b9366004614cda565b611aad565b3480156106ca57600080fd5b506104256106d9366004614be6565b611c91565b3480156106ea57600080fd5b50600d5461044c906301000000900460ff1681565b34801561070b57600080fd5b50610496611cde565b34801561072057600080fd5b5061073461072f366004614a9b565b611d6c565b60405161045893929190614d45565b34801561074f57600080fd5b50600d5461044c90610100900460ff1681565b34801561076e57600080fd5b50610425611f26565b34801561078357600080fd5b506104c3610792366004614a9b565b611fda565b3480156107a357600080fd5b506104966040518060400160405280601581526020017411995c985b199a5b19515e1a1a589a5d1a5bdb958d605a1b81525081565b3480156107e457600080fd5b506104256107f3366004614d7e565b61200f565b34801561080457600080fd5b50610425610813366004614a1d565b61230d565b34801561082457600080fd5b50610425610833366004614a1d565b6123a3565b34801561084457600080fd5b5061058c610853366004614a1d565b61242e565b34801561086457600080fd5b506104256124b4565b34801561087957600080fd5b50610425610888366004614e3f565b6124c8565b34801561089957600080fd5b5061058c6108a8366004614a9b565b60009081526010602052604090205490565b3480156108c657600080fd5b506108da6108d5366004614a1d565b6126e5565b6040516104589190614eba565b3480156108f357600080fd5b50610425610902366004614efe565b612751565b34801561091357600080fd5b506006546001600160a01b03166104c3565b34801561093157600080fd5b50610425610940366004614ff2565b612880565b34801561095157600080fd5b5061058c610960366004614a1d565b60146020526000908152604090205481565b34801561097e57600080fd5b506104966128ef565b34801561099357600080fd5b5061058c6201194081565b3480156109aa57600080fd5b50600d5461044c9060ff1681565b3480156109c457600080fd5b506104256109d336600461503f565b6128fe565b3480156109e457600080fd5b506104966109f3366004614a9b565b612912565b348015610a0457600080fd5b5061042561295f565b348015610a1957600080fd5b50610425610a28366004615076565b61297c565b348015610a3957600080fd5b506104256129ca565b348015610a4e57600080fd5b50610496610a5d366004614a9b565b612a6b565b348015610a6e57600080fd5b50610425610a7d366004614a1d565b612d7d565b348015610a8e57600080fd5b50610496612da9565b348015610aa357600080fd5b5061058c615dc081565b348015610ab957600080fd5b5061044c610ac83660046150f1565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b348015610b0257600080fd5b5061058c610b11366004614a9b565b6000908152600f602052604090205490565b348015610b2f57600080fd5b5061044c610b3e366004614a1d565b60076020526000908152604090205460ff1681565b348015610b5f57600080fd5b506008546104c3906001600160a01b031681565b348015610b7f57600080fd5b50610425610b8e366004614a1d565b612db6565b348015610b9f57600080fd5b50600d546104c39064010000000090046001600160a01b031681565b348015610bc757600080fd5b50610425610bd6366004614cda565b612e2f565b348015610be757600080fd5b50600e546104c3906001600160a01b031681565b348015610c0757600080fd5b50610496610c16366004614a9b565b613029565b60006001600160e01b031982166380ac58cd60e01b1480610c4c57506001600160e01b03198216635b5e139f60e01b145b80610c6757506301ffc9a760e01b6001600160e01b03198316145b92915050565b610c75613042565b6001600160a01b03166000908152600760205260409020805460ff19169055565b606060008054610ca590615124565b80601f0160208091040260200160405190810160405280929190818152602001828054610cd190615124565b8015610d1e5780601f10610cf357610100808354040283529160200191610d1e565b820191906000526020600020905b815481529060010190602001808311610d0157829003601f168201915b5050505050905090565b6000610d338261309c565b506000908152600460205260409020546001600160a01b031690565b81610d59816130c1565b610d638383613193565b505050565b610d70613042565b600880546001600160a01b0319166001600160a01b0392909216919091179055565b610d9a613042565b6001600160a01b038116610e0d5760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a20636f737452656365696044820152737665725f206973207a65726f206164647265737360601b606482015260840161041c565b600d80546001600160a01b0390921664010000000002640100000000600160c01b0319909216919091179055565b6040805180820190915260008082526020820152610e58826132a3565b610e745760405162461bcd60e51b815260040161041c90615158565b50600090815260116020908152604091829020825180840190935280548352600101549082015290565b600d5460ff16610f055760405162461bcd60e51b815260206004820152602c60248201527f466572616c66696c6545786869626974696f6e56343a20746f6b656e2069732060448201526b6e6f74206275726e61626c6560a01b606482015260840161041c565b60005b8151811015610f8457610f3433838381518110610f2757610f2761518f565b60200260200101516132c0565b610f505760405162461bcd60e51b815260040161041c906151a5565b610f72828281518110610f6557610f6561518f565b602002602001015161333f565b80610f7c81615208565b915050610f08565b5050565b826001600160a01b0381163314610fa257610fa2336130c1565b306001600160a01b03841603610fca5760405162461bcd60e51b815260040161041c90615221565b610fd5848484613415565b50505050565b600d5462010000900460ff166110465760405162461bcd60e51b815260206004820152602a60248201527f466572616c66696c6545786869626974696f6e56343a2073616c65206973206e6044820152691bdd081cdd185c9d195960b21b606482015260840161041c565b61104e613446565b611057816134c1565b61106760e0820160c0830161527e565b6110d657803534146110d15760405162461bcd60e51b815260206004820152602d60248201527f466572616c66696c6545786869626974696f6e56343a20696e76616c6964207060448201526c185e5b595b9d08185b5bdd5b9d609a1b606482015260840161041c565b61113f565b600e54604051632eeee16360e01b81526001600160a01b0390911690632eeee1639061110c908790879087908790600401615487565b600060405180830381600087803b15801561112657600080fd5b505af115801561113a573d6000803e3d6000fd5b505050505b6000463083604051602001611156939291906154af565b60405160208183030381529060405280519060200120905061117a8186868661361a565b61119757604051638baa579f60e01b815260040160405180910390fd5b60006020830135833511156111d1576111b360808401846154e2565b90506111c46020850135853561552b565b6111ce919061553e565b90505b60008060005b6111e460808701876154e2565b90508110156115385761123a306112016080890160608a01614a1d565b61120e60808a018a6154e2565b8581811061121e5761121e61518f565b9050602002013560405180602001604052806000815250613672565b600061124960a08801886154e2565b838181106112595761125961518f565b905060200281019061126b9190615560565b808060200260200160405190810160405280939291908181526020016000905b828210156112b7576112a8604083028601368190038101906155a9565b8152602001906001019061128b565b50505050509050600085905060005b8251811080156112d65750600082115b156113bb576000601460008584815181106112f3576112f361518f565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000205490506000838210156113325781611334565b835b905061134081886155ff565b965080601460008786815181106113595761135961518f565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000206000828254611394919061552b565b909155506113a49050818561552b565b9350505080806113b390615208565b9150506112c6565b5080156114b75760005b82518110156114b55760008382815181106113e2576113e261518f565b602002602001015160000151905060006127108584815181106114075761140761518f565b6020026020010151602001518561141e9190615612565b611428919061553e565b600d549091506001600160a01b0364010000000090910481169083160361145c5761145381886155ff565b965050506114a3565b61146681896155ff565b6040519098506001600160a01b0383169082156108fc029083906000818181858888f1935050505015801561149f573d6000803e3d6000fd5b5050505b806114ad81615208565b9150506113c5565b505b6114c460808901896154e2565b848181106114d4576114d461518f565b905060200201358860600160208101906114ee9190614a1d565b6001600160a01b03167f0475389cd69b8d3163620b43283bf74e8fc71020c3c6cef2a529b5c405e9687f60405160405180910390a35050808061153090615208565b9150506111d7565b5061154381836155ff565b6115526020870135873561552b565b10156115b55760405162461bcd60e51b815260206004820152602c60248201527f466572616c66696c6545786869626974696f6e56343a20746f74616c2062707360448201526b0206f7665722031302c3030360a41b606482015260840161041c565b60006115c283873561552b565b9050801561160d57600d546040516401000000009091046001600160a01b0316906108fc8315029083906000818181858888f1935050505015801561160b573d6000803e3d6000fd5b505b505050505050505050565b60006116238361242e565b82106116855760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b606482015260840161041c565b6001600160a01b03831660009081526012602052604090208054839081106116af576116af61518f565b9060005260206000200154905092915050565b3360009081526007602052604090205460ff16806116ea57506006546001600160a01b031633145b6116f357600080fd5b8281146117135760405163ab3a7f9960e01b815260040160405180910390fd5b60005b8181101561175b576117538383838181106117335761173361518f565b90506020028101906117459190615629565b61174e90615649565b6136a5565b600101611716565b506117668484613708565b60005b8381101561184a578282828181106117835761178361518f565b90506020028101906117959190615629565b601760008787858181106117ab576117ab61518f565b90506020020135815260200190815260200160002081816117cc9190615855565b9050508484828181106117e1576117e161518f565b905060200201357fac24f5d052fdb3e47d0446513e40bbf8afcfd59e1bd778f60c1f53da4212ead684848481811061181b5761181b61518f565b905060200281019061182d9190615629565b60405161183a91906159bc565b60405180910390a2600101611769565b5050505050565b611859613042565b600d546301000000900460ff16156118835760405162461bcd60e51b815260040161041c90615a32565b600d5462010000900460ff16156118f95760405162461bcd60e51b815260206004820152603460248201527f466572616c66696c6545786869626974696f6e56343a205f73656c6c696e6720604482015273726571756972656420746f2062652066616c736560601b606482015260840161041c565b611901613446565b600d805462ff0000191662010000179055565b61191c613042565b82811461193c576040516313086eff60e21b815260040160405180910390fd5b60005b8381101561184a57600085858381811061195b5761195b61518f565b90506020020160208101906119709190614a1d565b6001600160a01b03160361199757604051630107349760e51b815260040160405180910390fd5b8282828181106119a9576119a961518f565b905060200201356000036119d057604051636745f8fb60e01b815260040160405180910390fd5b6000601460008787858181106119e8576119e861518f565b90506020020160208101906119fd9190614a1d565b6001600160a01b03166001600160a01b03168152602001908152602001600020541115611a3d576040516328547bdf60e01b815260040160405180910390fd5b828282818110611a4f57611a4f61518f565b9050602002013560146000878785818110611a6c57611a6c61518f565b9050602002016020810190611a819190614a1d565b6001600160a01b0316815260208101919091526040016000205580611aa581615208565b91505061193f565b611ab5613042565b828114611ad5576040516313086eff60e21b815260040160405180910390fd5b60005b8381101561184a576000838383818110611af457611af461518f565b9050602002016020810190611b099190614a1d565b6001600160a01b031603611b3057604051630107349760e51b815260040160405180910390fd5b600060146000858585818110611b4857611b4861518f565b9050602002016020810190611b5d9190614a1d565b6001600160a01b03166001600160a01b03168152602001908152602001600020541115611b9d576040516328547bdf60e01b815260040160405180910390fd5b60146000868684818110611bb357611bb361518f565b9050602002016020810190611bc89190614a1d565b6001600160a01b03166001600160a01b031681526020019081526020016000205460146000858585818110611bff57611bff61518f565b9050602002016020810190611c149190614a1d565b6001600160a01b03166001600160a01b031681526020019081526020016000208190555060146000868684818110611c4e57611c4e61518f565b9050602002016020810190611c639190614a1d565b6001600160a01b03168152602081019190915260400160009081205580611c8981615208565b915050611ad8565b826001600160a01b0381163314611cab57611cab336130c1565b306001600160a01b03841603611cd35760405162461bcd60e51b815260040161041c90615221565b610fd58484846137fa565b600a8054611ceb90615124565b80601f0160208091040260200160405190810160405280929190818152602001828054611d1790615124565b8015611d645780601f10611d3957610100808354040283529160200191611d64565b820191906000526020600020905b815481529060010190602001808311611d4757829003601f168201915b505050505081565b601760205260009081526040902080548190611d8790615124565b80601f0160208091040260200160405190810160405280929190818152602001828054611db390615124565b8015611e005780601f10611dd557610100808354040283529160200191611e00565b820191906000526020600020905b815481529060010190602001808311611de357829003601f168201915b505050505090806001018054611e1590615124565b80601f0160208091040260200160405190810160405280929190818152602001828054611e4190615124565b8015611e8e5780601f10611e6357610100808354040283529160200191611e8e565b820191906000526020600020905b815481529060010190602001808311611e7157829003601f168201915b505050505090806002018054611ea390615124565b80601f0160208091040260200160405190810160405280929190818152602001828054611ecf90615124565b8015611f1c5780601f10611ef157610100808354040283529160200191611f1c565b820191906000526020600020905b815481529060010190602001808311611eff57829003601f168201915b5050505050905083565b611f2e613042565b600d546301000000900460ff1615611f585760405162461bcd60e51b815260040161041c90615a32565b600d5462010000900460ff16611fcc5760405162461bcd60e51b815260206004820152603360248201527f466572616c66696c6545786869626974696f6e56343a205f73656c6c696e6720604482015272726571756972656420746f206265207472756560681b606482015260840161041c565b600d805462ff000019169055565b6000818152600260205260408120546001600160a01b031680610c675760405162461bcd60e51b815260040161041c90615158565b612017613042565b60008251118015612029575060008151115b6120a95760405162461bcd60e51b815260206004820152604560248201527f466572616c66696c6545786869626974696f6e56343a2073657269657349647360448201527f206f7220726563697069656e74416464726573736573206c656e677468206973606482015264207a65726f60d81b608482015260a40161041c565b80518251146121355760405162461bcd60e51b815260206004820152604c60248201527f466572616c66696c6545786869626974696f6e56343a2073657269657349647360448201527f206c656e67746820697320646966666572656e742066726f6d2072656369706960648201526b656e7441646472657373657360a01b608482015260a40161041c565b61213d611f26565b3060008181526012602090815260408083208054825181850281018501909352808352919290919083018282801561219457602002820191906000526020600020905b815481526020019060010190808311612180575b5050505050905060005b81518110156122905760008282815181106121bb576121bb61518f565b602090810291909101810151600081815260118352604080822081518083019092528054825260010154938101939093529092505b87518161ffff16101561227a57878161ffff16815181106122135761221361518f565b6020026020010151826000015103612268576000878261ffff168151811061223d5761223d61518f565b6020026020010151905061226287828660405180602001604052806000815250613672565b5061227a565b8061227281615a86565b9150506121f0565b505050808061228890615208565b91505061219e565b5061229a8261242e565b15610fd55760405162461bcd60e51b815260206004820152603c60248201527f466572616c66696c6545786869626974696f6e56343a20546f6b656e20666f7260448201527f2073616c652062616c616e63652068617320746f206265207a65726f00000000606482015260840161041c565b612315613042565b6001600160a01b0381166123815760405162461bcd60e51b815260206004820152602d60248201527f466572616c66696c6545786869626974696f6e56343a207661756c745f20697360448201526c207a65726f206164647265737360981b606482015260840161041c565b600e80546001600160a01b0319166001600160a01b0392909216919091179055565b6123ab613042565b6001600160a01b03811661240c5760405162461bcd60e51b815260206004820152602260248201527f45434453415369676e3a207369676e65725f206973207a65726f206164647265604482015261737360f01b606482015260840161041c565b600980546001600160a01b0319166001600160a01b0392909216919091179055565b60006001600160a01b0382166124985760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b606482015260840161041c565b506001600160a01b031660009081526003602052604090205490565b6124bc613042565b6124c66000613815565b565b3360009081526007602052604090205460ff16806124f057506006546001600160a01b031633145b6124f957600080fd5b6000838152600f6020526040902054839061252a57604051630e17b51760e41b81526004810182905260240161041c565b600082900361254c576040516328f5a3ff60e01b815260040160405180910390fd5b6201194082111561257057604051635d3c714b60e11b815260040160405180910390fd5b600084815260156020526040902054156125c757600084815260156020526040812061259b91614995565b60405184907f5642c40d50e2073be9e5183e69e9c92f89328520b1a18fcd0d109b25434bd9d290600090a25b60005b828110156126945760006125de828561552b565b9050615dc08111156125ef5750615dc05b600061264886848761260186836155ff565b9261260e93929190615aa7565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061386792505050565b60008881526015602090815260408220805460018101825590835291200180546001600160a01b0319166001600160a01b038316179055905061268b82846155ff565b925050506125ca565b847f0bfaeb43f2350d013eaed4cf979fcf8c33f3471f9a6a69b312882034d3134bbb601560008881526020019081526020016000206040516126d69190615ad1565b60405180910390a25050505050565b6001600160a01b03811660009081526012602090815260409182902080548351818402810184019094528084526060939283018282801561274557602002820191906000526020600020905b815481526020019060010190808311612731575b50505050509050919050565b3360009081526007602052604090205460ff168061277957506006546001600160a01b031633145b61278257600080fd5b600d546301000000900460ff166127f95760405162461bcd60e51b815260206004820152603560248201527f466572616c66696c6545786869626974696f6e56343a20636f6e747261637420604482015274191bd95cdb89dd08185b1b1bddc81d1bc81b5a5b9d605a1b606482015260840161041c565b60005b81811015610d635761286e8383838181106128195761281961518f565b905060600201600001358484848181106128355761283561518f565b905060600201602001358585858181106128515761285161518f565b90506060020160400160208101906128699190614a1d565b6138c6565b8061287881615208565b9150506127fc565b612888613042565b60008151116128e35760405162461bcd60e51b815260206004820152602160248201527f4552433732314d657461646174613a20626173655552495f20697320656d70746044820152607960f81b606482015260840161041c565b600a610f848282615b15565b606060018054610ca590615124565b81612908816130c1565b610d638383613a48565b60608161292d816000908152600f6020526040902054151590565b61294d57604051630e17b51760e41b81526004810182905260240161041c565b61295683613a53565b91505b50919050565b612967613042565b600d805463ff000000191690556124c6611851565b836001600160a01b038116331461299657612996336130c1565b306001600160a01b038516036129be5760405162461bcd60e51b815260040161041c90615221565b61184a85858585613b33565b6129d2613042565b6129da611f26565b30600090815260126020908152604080832080548251818502810185019093528083529192909190830182828015612a3157602002820191906000526020600020905b815481526020019060010190808311612a1d575b5050505050905060005b8151811015610f8457612a59828281518110610f6557610f6561518f565b80612a6381615208565b915050612a3b565b6060612a76826132a3565b612a965760405163401e26ff60e01b81526004810183905260240161041c565b600082815260116020526040902054612abd81600090815260156020526040902054151590565b612aca5761295683613b65565b6000612ad582613a53565b9050600060176000868152602001908152602001600020604051806060016040529081600082018054612b0790615124565b80601f0160208091040260200160405190810160405280929190818152602001828054612b3390615124565b8015612b805780601f10612b5557610100808354040283529160200191612b80565b820191906000526020600020905b815481529060010190602001808311612b6357829003601f168201915b50505050508152602001600182018054612b9990615124565b80601f0160208091040260200160405190810160405280929190818152602001828054612bc590615124565b8015612c125780601f10612be757610100808354040283529160200191612c12565b820191906000526020600020905b815481529060010190602001808311612bf557829003601f168201915b50505050508152602001600282018054612c2b90615124565b80601f0160208091040260200160405190810160405280929190818152602001828054612c5790615124565b8015612ca45780601f10612c7957610100808354040283529160200191612ca4565b820191906000526020600020905b815481529060010190602001808311612c8757829003601f168201915b50505050508152505090506000601660008581526020019081526020016000208260400151604051602001612cda929190615c41565b60408051601f198184030181529082905260208401518451636b985f3160e01b845291935073__RendererStorageV0_____________________92636b985f3192612d2e9288929091908790600401615c73565b600060405180830381865af4158015612d4b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612d739190810190615cfb565b9695505050505050565b612d85613042565b6001600160a01b03166000908152600760205260409020805460ff19166001179055565b600b8054611ceb90615124565b612dbe613042565b6001600160a01b038116612e235760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161041c565b612e2c81613815565b50565b3360009081526007602052604090205460ff1680612e5757506006546001600160a01b031633145b612e6057600080fd5b828114612e805760405163ab3a7f9960e01b815260040160405180910390fd5b60005b83811015612f3a57828282818110612e9d57612e9d61518f565b9050602002810190612eaf91906156f3565b9050600003612ed157604051634f45b8fd60e01b815260040160405180910390fd5b612f32838383818110612ee657612ee661518f565b9050602002810190612ef891906156f3565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613c7392505050565b600101612e83565b50612f458484613d03565b60005b8381101561184a57828282818110612f6257612f6261518f565b9050602002810190612f7491906156f3565b60166000888886818110612f8a57612f8a61518f565b9050602002013581526020019081526020016000209182612fac92919061579c565b50848482818110612fbf57612fbf61518f565b905060200201357fa4154606c7e38a6a299a31a214bf3468d191c48e6308a1e849932d2c5aba56b9848484818110612ff957612ff961518f565b905060200281019061300b91906156f3565b604051613019929190615d43565b60405180910390a2600101612f48565b60166020526000908152604090208054611ceb90615124565b6006546001600160a01b031633146124c65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161041c565b6130a5816132a3565b612e2c5760405162461bcd60e51b815260040161041c90615158565b6008546001600160a01b03163b15612e2c57600854604051633185c44d60e21b81523060048201526001600160a01b0383811660248301529091169063c617113490604401602060405180830381865afa158015613123573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131479190615d57565b612e2c5760405162461bcd60e51b815260206004820152601760248201527f6f70657261746f72206973206e6f7420616c6c6f776564000000000000000000604482015260640161041c565b600061319e82611fda565b9050806001600160a01b0316836001600160a01b03160361320b5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b606482015260840161041c565b336001600160a01b038216148061322757506132278133610ac8565b6132995760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c000000606482015260840161041c565b610d638383613dee565b6000908152600260205260409020546001600160a01b0316151590565b6000806132cc83611fda565b9050806001600160a01b0316846001600160a01b0316148061331357506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b806133375750836001600160a01b031661332c84610d28565b6001600160a01b0316145b949350505050565b613348816132a3565b6133645760405162461bcd60e51b815260040161041c90615158565b600081815260116020908152604080832081518083018352815480825260019283015482860152855260109093529083208054929391929091906133a990849061552b565b925050819055506001600c60008282546133c3919061552b565b90915550506000828152601160205260408120818155600101556133e682613e5c565b60405182907fbde7938970372996ff103863625e348ef2bf8f38a5b02181be75aafef17c23d590600090a25050565b61341f33826132c0565b61343b5760405162461bcd60e51b815260040161041c906151a5565b610d63838383613eff565b60006134513061242e565b905060008111612e2c5760405162461bcd60e51b815260206004820152603560248201527f466572616c66696c6545786869626974696f6e56343a204e6f20746f6b656e206044820152741bdddb995908189e481d1a194818dbdb9d1c9858dd605a1b606482015260840161041c565b60006134d060808301836154e2565b90501161352b5760405162461bcd60e51b8152602060048201526024808201527f466572616c66696c6553616c65446174613a20746f6b656e49647320697320656044820152636d70747960e01b606482015260840161041c565b61353860a08201826154e2565b905061354760808301836154e2565b9050146135bc5760405162461bcd60e51b815260206004820152603d60248201527f466572616c66696c6553616c65446174613a20746f6b656e49647320616e642060448201527f726576656e7565536861726573206c656e677468206d69736d61746368000000606482015260840161041c565b42816040013511612e2c5760405162461bcd60e51b815260206004820152602260248201527f466572616c66696c6553616c65446174613a2073616c65206973206578706972604482015261195960f21b606482015260840161041c565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c859052603c8120819061365790848787614070565b6009546001600160a01b039081169116149695505050505050565b61367d848484613eff565b61368984848484614098565b610fd55760405162461bcd60e51b815260040161041c90615d74565b80515115806136b75750602081015151155b806136c55750604081015151155b156136e357604051634f45b8fd60e01b815260040160405180910390fd5b80516136ee90613c73565b6136fb8160200151613c73565b612e2c8160400151613c73565b600081900361372a576040516316ee9d3b60e11b815260040160405180910390fd5b60005b81811015610d635761375683838381811061374a5761374a61518f565b905060200201356132a3565b61378f5782828281811061376c5761376c61518f565b9050602002013560405163401e26ff60e01b815260040161041c91815260200190565b6137d5601160008585858181106137a8576137a861518f565b90506020020135815260200190815260200160002060000154600090815260156020526040902054151590565b6137f25760405163466fcd5560e11b815260040160405180910390fd5b60010161372d565b610d638383836040518060200160405280600081525061297c565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000806138928360405160200161387e9190615dc6565b604051602081830303815290604052614196565b90508051602082016000f091506001600160a01b0382166129595760405163046a55db60e11b815260040160405180910390fd5b6000838152600f602052604090205415156138e0846141ac565b6040516020016138f09190615dec565b6040516020818303038152906040529061391d5760405162461bcd60e51b815260040161041c9190614a88565b506000838152600f6020908152604080832054601090925290912054106139985760405162461bcd60e51b815260206004820152602960248201527f466572616c66696c6545786869626974696f6e56343a206e6f20736c6f747320604482015268617661696c61626c6560b81b606482015260840161041c565b6001600c60008282546139ab91906155ff565b909155505060008381526010602052604081208054600192906139cf9084906155ff565b9091555050604080518082018252848152602080820185815260008681526011909252929020905181559051600190910155613a0b818361423f565b8183826001600160a01b03167f407d7da1d3b2b1871fbfa2b5b1c4657a3cc5711d3023c552798551c7ee301eea60405160405180910390a4505050565b610f843383836143ba565b600081815260156020908152604091829020805483518184028101840190945280845260609392830182828015613ab357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613a95575b5050604051630b09269f60e41b815273__LibBytes______________________________9463b09269f09450613aee93509150600401615e49565b600060405180830381865af4158015613b0b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c679190810190615cfb565b613b3d33836132c0565b613b595760405162461bcd60e51b815260040161041c906151a5565b610fd584848484613672565b60606000600a8054613b7690615124565b905011613bd45760405162461bcd60e51b815260206004820152602660248201527f4552433732314d657461646174613a205f746f6b656e4261736555524920697360448201526520656d70747960d01b606482015260840161041c565b613bdd826132a3565b613c415760405162461bcd60e51b815260206004820152602f60248201527f4552433732314d657461646174613a2055524920717565727920666f72206e6f60448201526e3732bc34b9ba32b73a103a37b5b2b760891b606482015260840161041c565b600a613c4c836141ac565b604051602001613c5d929190615e8a565b6040516020818303038152906040529050919050565b8060005b8151811015610d63576000828281518110613c9457613c9461518f565b01602001516001600160f81b0319169050601160f91b811480613cc45750601760fa1b6001600160f81b03198216145b80613cdc5750600160fd1b6001600160f81b03198216105b15613cfa5760405163a2a50ced60e01b815260040160405180910390fd5b50600101613c77565b6000819003613d25576040516316ee9d3b60e11b815260040160405180910390fd5b60005b81811015610d6357613d60838383818110613d4557613d4561518f565b905060200201356000908152600f6020526040902054151590565b613d9957828282818110613d7657613d7661518f565b90506020020135604051630e17b51760e41b815260040161041c91815260200190565b613dc9838383818110613dae57613dae61518f565b90506020020135600090815260156020526040902054151590565b613de65760405163466fcd5560e11b815260040160405180910390fd5b600101613d28565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190613e2382611fda565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000613e6782611fda565b9050613e77816000846001614488565b613e8082611fda565b600083815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526003845282852080546000190190558785526002909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b826001600160a01b0316613f1282611fda565b6001600160a01b031614613f385760405162461bcd60e51b815260040161041c90615eb0565b6001600160a01b038216613f9a5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161041c565b613fa78383836001614488565b826001600160a01b0316613fba82611fda565b6001600160a01b031614613fe05760405162461bcd60e51b815260040161041c90615eb0565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b60008060006140818787878761459f565b9150915061408e81614663565b5095945050505050565b60006001600160a01b0384163b1561418e57604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906140dc903390899088908890600401615ef5565b6020604051808303816000875af1925050508015614117575060408051601f3d908101601f1916820190925261411491810190615f28565b60015b614174573d808015614145576040519150601f19603f3d011682016040523d82523d6000602084013e61414a565b606091505b50805160000361416c5760405162461bcd60e51b815260040161041c90615d74565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050613337565b506001613337565b6060815182604051602001613c5d929190615f45565b606060006141b9836147ad565b60010190506000816001600160401b038111156141d8576141d8614ade565b6040519080825280601f01601f191660200182016040528015614202576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461420c575b509392505050565b6001600160a01b0382166142955760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015260640161041c565b61429e816132a3565b156142eb5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161041c565b6142f9600083836001614488565b614302816132a3565b1561434f5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161041c565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b816001600160a01b0316836001600160a01b03160361441b5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015260640161041c565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b60018111156144f75760405162461bcd60e51b815260206004820152603560248201527f455243373231456e756d657261626c653a20636f6e7365637574697665207472604482015274185b9cd9995c9cc81b9bdd081cdd5c1c1bdc9d1959605a1b606482015260840161041c565b816001600160a01b038516158015906145225750836001600160a01b0316856001600160a01b031614155b15614531576145318582614885565b6001600160a01b0384161580159061455b5750846001600160a01b0316846001600160a01b031614155b1561184a576001600160a01b03841660009081526012602090815260408083208054600181018255908452828420810185905584845260139092529091205561184a565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156145d6575060009050600361465a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561462a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166146535760006001925092505061465a565b9150600090505b94509492505050565b600081600481111561467757614677615f96565b0361467f5750565b600181600481111561469357614693615f96565b036146e05760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161041c565b60028160048111156146f4576146f4615f96565b036147415760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161041c565b600381600481111561475557614755615f96565b03612e2c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161041c565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106147ec5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310614818576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061483657662386f26fc10000830492506010015b6305f5e100831061484e576305f5e100830492506008015b612710831061486257612710830492506004015b60648310614874576064830492506002015b600a8310610c675760010192915050565b600060016148928461242e565b61489c919061552b565b600083815260136020526040902054909150808214614943576001600160a01b03841660009081526012602052604081208054849081106148df576148df61518f565b906000526020600020015490508060126000876001600160a01b03166001600160a01b0316815260200190815260200160002083815481106149235761492361518f565b600091825260208083209091019290925591825260139052604090208190555b60008381526013602090815260408083208390556001600160a01b03871683526012909152902080548061497957614979615fac565b6001900381819060005260206000200160009055905550505050565b5080546000825590600052602060002090810190612e2c91905b808211156149c357600081556001016149af565b5090565b6001600160e01b031981168114612e2c57600080fd5b6000602082840312156149ef57600080fd5b81356149fa816149c7565b9392505050565b80356001600160a01b0381168114614a1857600080fd5b919050565b600060208284031215614a2f57600080fd5b6149fa82614a01565b60005b83811015614a53578181015183820152602001614a3b565b50506000910152565b60008151808452614a74816020860160208601614a38565b601f01601f19169290920160200192915050565b6020815260006149fa6020830184614a5c565b600060208284031215614aad57600080fd5b5035919050565b60008060408385031215614ac757600080fd5b614ad083614a01565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715614b1c57614b1c614ade565b604052919050565b60006001600160401b03821115614b3d57614b3d614ade565b5060051b60200190565b600082601f830112614b5857600080fd5b81356020614b6d614b6883614b24565b614af4565b82815260059290921b84018101918181019086841115614b8c57600080fd5b8286015b84811015614ba75780358352918301918301614b90565b509695505050505050565b600060208284031215614bc457600080fd5b81356001600160401b03811115614bda57600080fd5b61333784828501614b47565b600080600060608486031215614bfb57600080fd5b614c0484614a01565b9250614c1260208501614a01565b9150604084013590509250925092565b60008060008060808587031215614c3857600080fd5b8435935060208501359250604085013560ff81168114614c5757600080fd5b915060608501356001600160401b03811115614c7257600080fd5b850160e08188031215614c8457600080fd5b939692955090935050565b60008083601f840112614ca157600080fd5b5081356001600160401b03811115614cb857600080fd5b6020830191508360208260051b8501011115614cd357600080fd5b9250929050565b60008060008060408587031215614cf057600080fd5b84356001600160401b0380821115614d0757600080fd5b614d1388838901614c8f565b90965094506020870135915080821115614d2c57600080fd5b50614d3987828801614c8f565b95989497509550505050565b606081526000614d586060830186614a5c565b8281036020840152614d6a8186614a5c565b90508281036040840152612d738185614a5c565b60008060408385031215614d9157600080fd5b82356001600160401b0380821115614da857600080fd5b614db486838701614b47565b9350602091508185013581811115614dcb57600080fd5b85019050601f81018613614dde57600080fd5b8035614dec614b6882614b24565b81815260059190911b82018301908381019088831115614e0b57600080fd5b928401925b82841015614e3057614e2184614a01565b82529284019290840190614e10565b80955050505050509250929050565b600080600060408486031215614e5457600080fd5b8335925060208401356001600160401b0380821115614e7257600080fd5b818601915086601f830112614e8657600080fd5b813581811115614e9557600080fd5b876020828501011115614ea757600080fd5b6020830194508093505050509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614ef257835183529284019291840191600101614ed6565b50909695505050505050565b60008060208385031215614f1157600080fd5b82356001600160401b0380821115614f2857600080fd5b818501915085601f830112614f3c57600080fd5b813581811115614f4b57600080fd5b866020606083028501011115614f6057600080fd5b60209290920196919550909350505050565b60006001600160401b03821115614f8b57614f8b614ade565b50601f01601f191660200190565b6000614fa7614b6884614f72565b9050828152838383011115614fbb57600080fd5b828260208301376000602084830101529392505050565b600082601f830112614fe357600080fd5b6149fa83833560208501614f99565b60006020828403121561500457600080fd5b81356001600160401b0381111561501a57600080fd5b61333784828501614fd2565b8015158114612e2c57600080fd5b8035614a1881615026565b6000806040838503121561505257600080fd5b61505b83614a01565b9150602083013561506b81615026565b809150509250929050565b6000806000806080858703121561508c57600080fd5b61509585614a01565b93506150a360208601614a01565b92506040850135915060608501356001600160401b038111156150c557600080fd5b8501601f810187136150d657600080fd5b6150e587823560208401614f99565b91505092959194509250565b6000806040838503121561510457600080fd5b61510d83614a01565b915061511b60208401614a01565b90509250929050565b600181811c9082168061513857607f821691505b60208210810361295957634e487b7160e01b600052602260045260246000fd5b60208082526018908201527f4552433732313a20696e76616c696420746f6b656e2049440000000000000000604082015260600190565b634e487b7160e01b600052603260045260246000fd5b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60006001820161521a5761521a6151f2565b5060010190565b6020808252603e908201527f466572616c66696c6545786869626974696f6e56343a20436f6e74726163742060408201527f69736e277420616c6c6f77656420746f207265636569766520746f6b656e0000606082015260800190565b60006020828403121561529057600080fd5b81356149fa81615026565b6000808335601e198436030181126152b257600080fd5b83016020810192503590506001600160401b038111156152d157600080fd5b8060051b3603821315614cd357600080fd5b8183526000602080850194508260005b85811015615329576001600160a01b0361530c83614a01565b1687528183013583880152604096870196909101906001016152f3565b509495945050505050565b60008383855260208086019550808560051b8301018460005b878110156153bf57848303601f19018952813536889003601e1901811261537357600080fd5b870184810190356001600160401b0381111561538e57600080fd5b8060061b36038213156153a057600080fd5b6153ab8582846152e3565b9a86019a945050509083019060010161534d565b5090979650505050505050565b8035825260208082013590830152604080820135908301526001600160a01b036153f860608301614a01565b166060830152600061540d608083018361529b565b60e06080860181905285018190526101006001600160fb1b0382111561543257600080fd5b8160051b9150818382880137818601925061545060a086018661529b565b9250818785030160a08801526154698285018483615334565b935050505061547a60c08401615034565b80151560c0860152614237565b84815283602082015260ff83166040820152608060608201526000612d7360808301846153cc565b8381526001600160a01b03831660208201526060604082018190526000906154d9908301846153cc565b95945050505050565b6000808335601e198436030181126154f957600080fd5b8301803591506001600160401b0382111561551357600080fd5b6020019150600581901b3603821315614cd357600080fd5b81810381811115610c6757610c676151f2565b60008261555b57634e487b7160e01b600052601260045260246000fd5b500490565b6000808335601e1984360301811261557757600080fd5b8301803591506001600160401b0382111561559157600080fd5b6020019150600681901b3603821315614cd357600080fd5b6000604082840312156155bb57600080fd5b604051604081018181106001600160401b03821117156155dd576155dd614ade565b6040526155e983614a01565b8152602083013560208201528091505092915050565b80820180821115610c6757610c676151f2565b8082028115828204841417610c6757610c676151f2565b60008235605e1983360301811261563f57600080fd5b9190910192915050565b60006060823603121561565b57600080fd5b604051606081016001600160401b03828210818311171561567e5761567e614ade565b81604052843591508082111561569357600080fd5b61569f36838701614fd2565b835260208501359150808211156156b557600080fd5b6156c136838701614fd2565b602084015260408501359150808211156156da57600080fd5b506156e736828601614fd2565b60408301525092915050565b6000808335601e1984360301811261570a57600080fd5b8301803591506001600160401b0382111561572457600080fd5b602001915036819003821315614cd357600080fd5b601f821115610d6357600081815260208120601f850160051c810160208610156157605750805b601f850160051c820191505b8181101561577f5782815560010161576c565b505050505050565b600019600383901b1c191660019190911b1790565b6001600160401b038311156157b3576157b3614ade565b6157c7836157c18354615124565b83615739565b6000601f8411600181146157f557600085156157e35750838201355b6157ed8682615787565b84555061184a565b600083815260209020601f19861690835b828110156158265786850135825560209485019460019092019101615806565b50868210156158435760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b61585f82836156f3565b6001600160401b0381111561587657615876614ade565b61588a816158848554615124565b85615739565b6000601f8211600181146158b857600083156158a65750838201355b6158b08482615787565b865550615912565b600085815260209020601f19841690835b828110156158e957868501358255602094850194600190920191016158c9565b50848210156159065760001960f88660031b161c19848701351681555b505060018360011b0185555b5050505061592360208301836156f3565b61593181836001860161579c565b505061594060408301836156f3565b610fd581836002860161579c565b6000808335601e1984360301811261596557600080fd5b83016020810192503590506001600160401b0381111561598457600080fd5b803603821315614cd357600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6020815260006159cc838461594e565b606060208501526159e1608085018284615993565b9150506159f1602085018561594e565b601f1980868503016040870152615a09848385615993565b9350615a18604088018861594e565b935091508086850301606087015250612d73838383615993565b60208082526034908201527f466572616c66696c6545786869626974696f6e56343a206d696e7461626c6520604082015273726571756972656420746f2062652066616c736560601b606082015260800190565b600061ffff808316818103615a9d57615a9d6151f2565b6001019392505050565b60008085851115615ab757600080fd5b83861115615ac457600080fd5b5050820193919092039150565b6020808252825482820181905260008481528281209092916040850190845b81811015614ef25783546001600160a01b031683526001938401939285019201615af0565b81516001600160401b03811115615b2e57615b2e614ade565b615b4281615b3c8454615124565b84615739565b602080601f831160018114615b715760008415615b5f5750858301515b615b698582615787565b86555061577f565b600085815260208120601f198616915b82811015615ba057888601518255948401946001909101908401615b81565b5085821015615bbe5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008154615bdb81615124565b60018281168015615bf35760018114615c0857615c37565b60ff1984168752821515830287019450615c37565b8560005260208060002060005b85811015615c2e5781548a820152908401908201615c15565b50505082870194505b5050505092915050565b6000615c4d8285615bce565b600160fd1b81528351615c67816001840160208801614a38565b01600101949350505050565b608081526000615c866080830187614a5c565b8281036020840152615c988187614a5c565b90508281036040840152615cac8186614a5c565b90508281036060840152615cc08185614a5c565b979650505050505050565b6000615cd9614b6884614f72565b9050828152838383011115615ced57600080fd5b6149fa836020830184614a38565b600060208284031215615d0d57600080fd5b81516001600160401b03811115615d2357600080fd5b8201601f81018413615d3457600080fd5b61333784825160208401615ccb565b602081526000613337602083018486615993565b600060208284031215615d6957600080fd5b81516149fa81615026565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6000815260008251615ddf816001850160208701614a38565b9190910160010192915050565b7f466572616c66696c6545786869626974696f6e56343a2073657269657349642081526e03237b2b9b713ba1032bc34b9ba1d1608d1b602082015260008251615e3c81602f850160208701614a38565b91909101602f0192915050565b6020808252825182820181905260009190848201906040850190845b81811015614ef25783516001600160a01b031683529284019291840191600101615e65565b6000615e968285615bce565b602f60f81b81528351615c67816001840160208801614a38565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090612d7390830184614a5c565b600060208284031215615f3a57600080fd5b81516149fa816149c7565b606360f81b815260e083901b6001600160e01b03191660018201526880600e6000396000f360b81b60058201528151600090615f8881600e850160208701614a38565b91909101600e019392505050565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603160045260246000fdfea2646970667358221220f317666b5a2d22cd7785fa4a96f3cf1db617e9944626fda8cc6b2570f77095c664736f6c63430008110033",
}

// FeralfileExhibitionV44ABI is the input ABI used to generate the binding from.
// Deprecated: Use FeralfileExhibitionV44MetaData.ABI instead.
var FeralfileExhibitionV44ABI = FeralfileExhibitionV44MetaData.ABI

// FeralfileExhibitionV44Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeralfileExhibitionV44MetaData.Bin instead.
var FeralfileExhibitionV44Bin = FeralfileExhibitionV44MetaData.Bin

// DeployFeralfileExhibitionV44 deploys a new Ethereum contract, binding an instance of FeralfileExhibitionV44 to it.
func DeployFeralfileExhibitionV44(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, burnable_ bool, bridgeable_ bool, signer_ common.Address, vault_ common.Address, costReceiver_ common.Address, contractURI_ string, seriesIds_ []*big.Int, seriesMaxSupplies_ []*big.Int) (common.Address, *types.Transaction, *FeralfileExhibitionV44, error) {
	parsed, err := FeralfileExhibitionV44MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeralfileExhibitionV44Bin), backend, name_, symbol_, burnable_, bridgeable_, signer_, vault_, costReceiver_, contractURI_, seriesIds_, seriesMaxSupplies_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeralfileExhibitionV44{FeralfileExhibitionV44Caller: FeralfileExhibitionV44Caller{contract: contract}, FeralfileExhibitionV44Transactor: FeralfileExhibitionV44Transactor{contract: contract}, FeralfileExhibitionV44Filterer: FeralfileExhibitionV44Filterer{contract: contract}}, nil
}

// FeralfileExhibitionV44 is an auto generated Go binding around an Ethereum contract.
type FeralfileExhibitionV44 struct {
	FeralfileExhibitionV44Caller     // Read-only binding to the contract
	FeralfileExhibitionV44Transactor // Write-only binding to the contract
	FeralfileExhibitionV44Filterer   // Log filterer for contract events
}

// FeralfileExhibitionV44Caller is an auto generated read-only Go binding around an Ethereum contract.
type FeralfileExhibitionV44Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV44Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FeralfileExhibitionV44Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV44Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeralfileExhibitionV44Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV44Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeralfileExhibitionV44Session struct {
	Contract     *FeralfileExhibitionV44 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FeralfileExhibitionV44CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeralfileExhibitionV44CallerSession struct {
	Contract *FeralfileExhibitionV44Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// FeralfileExhibitionV44TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeralfileExhibitionV44TransactorSession struct {
	Contract     *FeralfileExhibitionV44Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// FeralfileExhibitionV44Raw is an auto generated low-level Go binding around an Ethereum contract.
type FeralfileExhibitionV44Raw struct {
	Contract *FeralfileExhibitionV44 // Generic contract binding to access the raw methods on
}

// FeralfileExhibitionV44CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeralfileExhibitionV44CallerRaw struct {
	Contract *FeralfileExhibitionV44Caller // Generic read-only contract binding to access the raw methods on
}

// FeralfileExhibitionV44TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeralfileExhibitionV44TransactorRaw struct {
	Contract *FeralfileExhibitionV44Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFeralfileExhibitionV44 creates a new instance of FeralfileExhibitionV44, bound to a specific deployed contract.
func NewFeralfileExhibitionV44(address common.Address, backend bind.ContractBackend) (*FeralfileExhibitionV44, error) {
	contract, err := bindFeralfileExhibitionV44(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44{FeralfileExhibitionV44Caller: FeralfileExhibitionV44Caller{contract: contract}, FeralfileExhibitionV44Transactor: FeralfileExhibitionV44Transactor{contract: contract}, FeralfileExhibitionV44Filterer: FeralfileExhibitionV44Filterer{contract: contract}}, nil
}

// NewFeralfileExhibitionV44Caller creates a new read-only instance of FeralfileExhibitionV44, bound to a specific deployed contract.
func NewFeralfileExhibitionV44Caller(address common.Address, caller bind.ContractCaller) (*FeralfileExhibitionV44Caller, error) {
	contract, err := bindFeralfileExhibitionV44(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44Caller{contract: contract}, nil
}

// NewFeralfileExhibitionV44Transactor creates a new write-only instance of FeralfileExhibitionV44, bound to a specific deployed contract.
func NewFeralfileExhibitionV44Transactor(address common.Address, transactor bind.ContractTransactor) (*FeralfileExhibitionV44Transactor, error) {
	contract, err := bindFeralfileExhibitionV44(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44Transactor{contract: contract}, nil
}

// NewFeralfileExhibitionV44Filterer creates a new log filterer instance of FeralfileExhibitionV44, bound to a specific deployed contract.
func NewFeralfileExhibitionV44Filterer(address common.Address, filterer bind.ContractFilterer) (*FeralfileExhibitionV44Filterer, error) {
	contract, err := bindFeralfileExhibitionV44(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44Filterer{contract: contract}, nil
}

// bindFeralfileExhibitionV44 binds a generic wrapper to an already deployed contract.
func bindFeralfileExhibitionV44(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeralfileExhibitionV44MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralfileExhibitionV44.Contract.FeralfileExhibitionV44Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.FeralfileExhibitionV44Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.FeralfileExhibitionV44Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralfileExhibitionV44.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.contract.Transact(opts, method, params...)
}

// CHUNKSIZE is a free data retrieval call binding the contract method 0xe91e13a9.
//
// Solidity: function CHUNK_SIZE() view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) CHUNKSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "CHUNK_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHUNKSIZE is a free data retrieval call binding the contract method 0xe91e13a9.
//
// Solidity: function CHUNK_SIZE() view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) CHUNKSIZE() (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.CHUNKSIZE(&_FeralfileExhibitionV44.CallOpts)
}

// CHUNKSIZE is a free data retrieval call binding the contract method 0xe91e13a9.
//
// Solidity: function CHUNK_SIZE() view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) CHUNKSIZE() (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.CHUNKSIZE(&_FeralfileExhibitionV44.CallOpts)
}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xf07e7fd0.
//
// Solidity: function OperatorFilterRegistry() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) OperatorFilterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "OperatorFilterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xf07e7fd0.
//
// Solidity: function OperatorFilterRegistry() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) OperatorFilterRegistry() (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.OperatorFilterRegistry(&_FeralfileExhibitionV44.CallOpts)
}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xf07e7fd0.
//
// Solidity: function OperatorFilterRegistry() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) OperatorFilterRegistry() (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.OperatorFilterRegistry(&_FeralfileExhibitionV44.CallOpts)
}

// RENDERERBLOBMAXSIZE is a free data retrieval call binding the contract method 0x9cd7d65c.
//
// Solidity: function RENDERER_BLOB_MAX_SIZE() view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) RENDERERBLOBMAXSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "RENDERER_BLOB_MAX_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RENDERERBLOBMAXSIZE is a free data retrieval call binding the contract method 0x9cd7d65c.
//
// Solidity: function RENDERER_BLOB_MAX_SIZE() view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) RENDERERBLOBMAXSIZE() (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.RENDERERBLOBMAXSIZE(&_FeralfileExhibitionV44.CallOpts)
}

// RENDERERBLOBMAXSIZE is a free data retrieval call binding the contract method 0x9cd7d65c.
//
// Solidity: function RENDERER_BLOB_MAX_SIZE() view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) RENDERERBLOBMAXSIZE() (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.RENDERERBLOBMAXSIZE(&_FeralfileExhibitionV44.CallOpts)
}

// Advances is a free data retrieval call binding the contract method 0x926ce44e.
//
// Solidity: function advances(address ) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Advances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "advances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Advances is a free data retrieval call binding the contract method 0x926ce44e.
//
// Solidity: function advances(address ) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Advances(arg0 common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.Advances(&_FeralfileExhibitionV44.CallOpts, arg0)
}

// Advances is a free data retrieval call binding the contract method 0x926ce44e.
//
// Solidity: function advances(address ) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Advances(arg0 common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.Advances(&_FeralfileExhibitionV44.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.BalanceOf(&_FeralfileExhibitionV44.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.BalanceOf(&_FeralfileExhibitionV44.CallOpts, owner)
}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Bridgeable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "bridgeable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Bridgeable() (bool, error) {
	return _FeralfileExhibitionV44.Contract.Bridgeable(&_FeralfileExhibitionV44.CallOpts)
}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Bridgeable() (bool, error) {
	return _FeralfileExhibitionV44.Contract.Bridgeable(&_FeralfileExhibitionV44.CallOpts)
}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Burnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "burnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Burnable() (bool, error) {
	return _FeralfileExhibitionV44.Contract.Burnable(&_FeralfileExhibitionV44.CallOpts)
}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Burnable() (bool, error) {
	return _FeralfileExhibitionV44.Contract.Burnable(&_FeralfileExhibitionV44.CallOpts)
}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) CodeVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "codeVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) CodeVersion() (string, error) {
	return _FeralfileExhibitionV44.Contract.CodeVersion(&_FeralfileExhibitionV44.CallOpts)
}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) CodeVersion() (string, error) {
	return _FeralfileExhibitionV44.Contract.CodeVersion(&_FeralfileExhibitionV44.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) ContractURI() (string, error) {
	return _FeralfileExhibitionV44.Contract.ContractURI(&_FeralfileExhibitionV44.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) ContractURI() (string, error) {
	return _FeralfileExhibitionV44.Contract.ContractURI(&_FeralfileExhibitionV44.CallOpts)
}

// CostReceiver is a free data retrieval call binding the contract method 0xf4e638be.
//
// Solidity: function costReceiver() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) CostReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "costReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CostReceiver is a free data retrieval call binding the contract method 0xf4e638be.
//
// Solidity: function costReceiver() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) CostReceiver() (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.CostReceiver(&_FeralfileExhibitionV44.CallOpts)
}

// CostReceiver is a free data retrieval call binding the contract method 0xf4e638be.
//
// Solidity: function costReceiver() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) CostReceiver() (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.CostReceiver(&_FeralfileExhibitionV44.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.GetApproved(&_FeralfileExhibitionV44.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.GetApproved(&_FeralfileExhibitionV44.CallOpts, tokenId)
}

// GetArtwork is a free data retrieval call binding the contract method 0x167ddf6e.
//
// Solidity: function getArtwork(uint256 tokenId) view returns((uint256,uint256))
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) GetArtwork(opts *bind.CallOpts, tokenId *big.Int) (FeralfileExhibitionV4Artwork, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "getArtwork", tokenId)

	if err != nil {
		return *new(FeralfileExhibitionV4Artwork), err
	}

	out0 := *abi.ConvertType(out[0], new(FeralfileExhibitionV4Artwork)).(*FeralfileExhibitionV4Artwork)

	return out0, err

}

// GetArtwork is a free data retrieval call binding the contract method 0x167ddf6e.
//
// Solidity: function getArtwork(uint256 tokenId) view returns((uint256,uint256))
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) GetArtwork(tokenId *big.Int) (FeralfileExhibitionV4Artwork, error) {
	return _FeralfileExhibitionV44.Contract.GetArtwork(&_FeralfileExhibitionV44.CallOpts, tokenId)
}

// GetArtwork is a free data retrieval call binding the contract method 0x167ddf6e.
//
// Solidity: function getArtwork(uint256 tokenId) view returns((uint256,uint256))
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) GetArtwork(tokenId *big.Int) (FeralfileExhibitionV4Artwork, error) {
	return _FeralfileExhibitionV44.Contract.GetArtwork(&_FeralfileExhibitionV44.CallOpts, tokenId)
}

// GetSeriesRenderer is a free data retrieval call binding the contract method 0xab221792.
//
// Solidity: function getSeriesRenderer(uint256 seriesId) view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) GetSeriesRenderer(opts *bind.CallOpts, seriesId *big.Int) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "getSeriesRenderer", seriesId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSeriesRenderer is a free data retrieval call binding the contract method 0xab221792.
//
// Solidity: function getSeriesRenderer(uint256 seriesId) view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) GetSeriesRenderer(seriesId *big.Int) (string, error) {
	return _FeralfileExhibitionV44.Contract.GetSeriesRenderer(&_FeralfileExhibitionV44.CallOpts, seriesId)
}

// GetSeriesRenderer is a free data retrieval call binding the contract method 0xab221792.
//
// Solidity: function getSeriesRenderer(uint256 seriesId) view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) GetSeriesRenderer(seriesId *big.Int) (string, error) {
	return _FeralfileExhibitionV44.Contract.GetSeriesRenderer(&_FeralfileExhibitionV44.CallOpts, seriesId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FeralfileExhibitionV44.Contract.IsApprovedForAll(&_FeralfileExhibitionV44.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FeralfileExhibitionV44.Contract.IsApprovedForAll(&_FeralfileExhibitionV44.CallOpts, owner, operator)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Mintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "mintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Mintable() (bool, error) {
	return _FeralfileExhibitionV44.Contract.Mintable(&_FeralfileExhibitionV44.CallOpts)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Mintable() (bool, error) {
	return _FeralfileExhibitionV44.Contract.Mintable(&_FeralfileExhibitionV44.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Name() (string, error) {
	return _FeralfileExhibitionV44.Contract.Name(&_FeralfileExhibitionV44.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Name() (string, error) {
	return _FeralfileExhibitionV44.Contract.Name(&_FeralfileExhibitionV44.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Owner() (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.Owner(&_FeralfileExhibitionV44.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Owner() (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.Owner(&_FeralfileExhibitionV44.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.OwnerOf(&_FeralfileExhibitionV44.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.OwnerOf(&_FeralfileExhibitionV44.CallOpts, tokenId)
}

// RendererTokenData is a free data retrieval call binding the contract method 0x504c05ea.
//
// Solidity: function rendererTokenData(uint256 ) view returns(string imageURI, string textureURI, string tokenName)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) RendererTokenData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ImageURI   string
	TextureURI string
	TokenName  string
}, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "rendererTokenData", arg0)

	outstruct := new(struct {
		ImageURI   string
		TextureURI string
		TokenName  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ImageURI = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TextureURI = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.TokenName = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// RendererTokenData is a free data retrieval call binding the contract method 0x504c05ea.
//
// Solidity: function rendererTokenData(uint256 ) view returns(string imageURI, string textureURI, string tokenName)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) RendererTokenData(arg0 *big.Int) (struct {
	ImageURI   string
	TextureURI string
	TokenName  string
}, error) {
	return _FeralfileExhibitionV44.Contract.RendererTokenData(&_FeralfileExhibitionV44.CallOpts, arg0)
}

// RendererTokenData is a free data retrieval call binding the contract method 0x504c05ea.
//
// Solidity: function rendererTokenData(uint256 ) view returns(string imageURI, string textureURI, string tokenName)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) RendererTokenData(arg0 *big.Int) (struct {
	ImageURI   string
	TextureURI string
	TokenName  string
}, error) {
	return _FeralfileExhibitionV44.Contract.RendererTokenData(&_FeralfileExhibitionV44.CallOpts, arg0)
}

// Selling is a free data retrieval call binding the contract method 0x23aed228.
//
// Solidity: function selling() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Selling(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "selling")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Selling is a free data retrieval call binding the contract method 0x23aed228.
//
// Solidity: function selling() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Selling() (bool, error) {
	return _FeralfileExhibitionV44.Contract.Selling(&_FeralfileExhibitionV44.CallOpts)
}

// Selling is a free data retrieval call binding the contract method 0x23aed228.
//
// Solidity: function selling() view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Selling() (bool, error) {
	return _FeralfileExhibitionV44.Contract.Selling(&_FeralfileExhibitionV44.CallOpts)
}

// SeriesMaxSupply is a free data retrieval call binding the contract method 0xeb5c60f2.
//
// Solidity: function seriesMaxSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) SeriesMaxSupply(opts *bind.CallOpts, seriesId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "seriesMaxSupply", seriesId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeriesMaxSupply is a free data retrieval call binding the contract method 0xeb5c60f2.
//
// Solidity: function seriesMaxSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SeriesMaxSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.SeriesMaxSupply(&_FeralfileExhibitionV44.CallOpts, seriesId)
}

// SeriesMaxSupply is a free data retrieval call binding the contract method 0xeb5c60f2.
//
// Solidity: function seriesMaxSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) SeriesMaxSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.SeriesMaxSupply(&_FeralfileExhibitionV44.CallOpts, seriesId)
}

// SeriesNames is a free data retrieval call binding the contract method 0xfe43658d.
//
// Solidity: function seriesNames(uint256 ) view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) SeriesNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "seriesNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SeriesNames is a free data retrieval call binding the contract method 0xfe43658d.
//
// Solidity: function seriesNames(uint256 ) view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SeriesNames(arg0 *big.Int) (string, error) {
	return _FeralfileExhibitionV44.Contract.SeriesNames(&_FeralfileExhibitionV44.CallOpts, arg0)
}

// SeriesNames is a free data retrieval call binding the contract method 0xfe43658d.
//
// Solidity: function seriesNames(uint256 ) view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) SeriesNames(arg0 *big.Int) (string, error) {
	return _FeralfileExhibitionV44.Contract.SeriesNames(&_FeralfileExhibitionV44.CallOpts, arg0)
}

// SeriesTotalSupply is a free data retrieval call binding the contract method 0x7f06ee06.
//
// Solidity: function seriesTotalSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) SeriesTotalSupply(opts *bind.CallOpts, seriesId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "seriesTotalSupply", seriesId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeriesTotalSupply is a free data retrieval call binding the contract method 0x7f06ee06.
//
// Solidity: function seriesTotalSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SeriesTotalSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.SeriesTotalSupply(&_FeralfileExhibitionV44.CallOpts, seriesId)
}

// SeriesTotalSupply is a free data retrieval call binding the contract method 0x7f06ee06.
//
// Solidity: function seriesTotalSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) SeriesTotalSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.SeriesTotalSupply(&_FeralfileExhibitionV44.CallOpts, seriesId)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Signer() (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.Signer(&_FeralfileExhibitionV44.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Signer() (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.Signer(&_FeralfileExhibitionV44.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeralfileExhibitionV44.Contract.SupportsInterface(&_FeralfileExhibitionV44.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeralfileExhibitionV44.Contract.SupportsInterface(&_FeralfileExhibitionV44.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Symbol() (string, error) {
	return _FeralfileExhibitionV44.Contract.Symbol(&_FeralfileExhibitionV44.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Symbol() (string, error) {
	return _FeralfileExhibitionV44.Contract.Symbol(&_FeralfileExhibitionV44.CallOpts)
}

// TokenBaseURI is a free data retrieval call binding the contract method 0x4e99b800.
//
// Solidity: function tokenBaseURI() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) TokenBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "tokenBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenBaseURI is a free data retrieval call binding the contract method 0x4e99b800.
//
// Solidity: function tokenBaseURI() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) TokenBaseURI() (string, error) {
	return _FeralfileExhibitionV44.Contract.TokenBaseURI(&_FeralfileExhibitionV44.CallOpts)
}

// TokenBaseURI is a free data retrieval call binding the contract method 0x4e99b800.
//
// Solidity: function tokenBaseURI() view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) TokenBaseURI() (string, error) {
	return _FeralfileExhibitionV44.Contract.TokenBaseURI(&_FeralfileExhibitionV44.CallOpts)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.TokenOfOwnerByIndex(&_FeralfileExhibitionV44.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.TokenOfOwnerByIndex(&_FeralfileExhibitionV44.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) TokenURI(tokenId *big.Int) (string, error) {
	return _FeralfileExhibitionV44.Contract.TokenURI(&_FeralfileExhibitionV44.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _FeralfileExhibitionV44.Contract.TokenURI(&_FeralfileExhibitionV44.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.TokensOfOwner(&_FeralfileExhibitionV44.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.TokensOfOwner(&_FeralfileExhibitionV44.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) TotalSupply() (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.TotalSupply(&_FeralfileExhibitionV44.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) TotalSupply() (*big.Int, error) {
	return _FeralfileExhibitionV44.Contract.TotalSupply(&_FeralfileExhibitionV44.CallOpts)
}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Trustees(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "trustees", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Trustees(arg0 common.Address) (bool, error) {
	return _FeralfileExhibitionV44.Contract.Trustees(&_FeralfileExhibitionV44.CallOpts, arg0)
}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Trustees(arg0 common.Address) (bool, error) {
	return _FeralfileExhibitionV44.Contract.Trustees(&_FeralfileExhibitionV44.CallOpts, arg0)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Caller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV44.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Vault() (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.Vault(&_FeralfileExhibitionV44.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44CallerSession) Vault() (common.Address, error) {
	return _FeralfileExhibitionV44.Contract.Vault(&_FeralfileExhibitionV44.CallOpts)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) AddTrustee(opts *bind.TransactOpts, _trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "addTrustee", _trustee)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) AddTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.AddTrustee(&_FeralfileExhibitionV44.TransactOpts, _trustee)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) AddTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.AddTrustee(&_FeralfileExhibitionV44.TransactOpts, _trustee)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.Approve(&_FeralfileExhibitionV44.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.Approve(&_FeralfileExhibitionV44.TransactOpts, operator, tokenId)
}

// BurnArtworks is a paid mutator transaction binding the contract method 0x21fe0c64.
//
// Solidity: function burnArtworks(uint256[] tokenIds) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) BurnArtworks(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "burnArtworks", tokenIds)
}

// BurnArtworks is a paid mutator transaction binding the contract method 0x21fe0c64.
//
// Solidity: function burnArtworks(uint256[] tokenIds) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) BurnArtworks(tokenIds []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.BurnArtworks(&_FeralfileExhibitionV44.TransactOpts, tokenIds)
}

// BurnArtworks is a paid mutator transaction binding the contract method 0x21fe0c64.
//
// Solidity: function burnArtworks(uint256[] tokenIds) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) BurnArtworks(tokenIds []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.BurnArtworks(&_FeralfileExhibitionV44.TransactOpts, tokenIds)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_) payable returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) BuyArtworks(opts *bind.TransactOpts, r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "buyArtworks", r_, s_, v_, saleData_)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_) payable returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) BuyArtworks(r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.BuyArtworks(&_FeralfileExhibitionV44.TransactOpts, r_, s_, v_, saleData_)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_) payable returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) BuyArtworks(r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.BuyArtworks(&_FeralfileExhibitionV44.TransactOpts, r_, s_, v_, saleData_)
}

// MintArtworks is a paid mutator transaction binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] data) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) MintArtworks(opts *bind.TransactOpts, data []FeralfileExhibitionV4MintData) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "mintArtworks", data)
}

// MintArtworks is a paid mutator transaction binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] data) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) MintArtworks(data []FeralfileExhibitionV4MintData) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.MintArtworks(&_FeralfileExhibitionV44.TransactOpts, data)
}

// MintArtworks is a paid mutator transaction binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] data) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) MintArtworks(data []FeralfileExhibitionV4MintData) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.MintArtworks(&_FeralfileExhibitionV44.TransactOpts, data)
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) PauseSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "pauseSale")
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) PauseSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.PauseSale(&_FeralfileExhibitionV44.TransactOpts)
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) PauseSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.PauseSale(&_FeralfileExhibitionV44.TransactOpts)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) RemoveTrustee(opts *bind.TransactOpts, _trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "removeTrustee", _trustee)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) RemoveTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.RemoveTrustee(&_FeralfileExhibitionV44.TransactOpts, _trustee)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) RemoveTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.RemoveTrustee(&_FeralfileExhibitionV44.TransactOpts, _trustee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) RenounceOwnership() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.RenounceOwnership(&_FeralfileExhibitionV44.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.RenounceOwnership(&_FeralfileExhibitionV44.TransactOpts)
}

// ReplaceAdvanceAddresses is a paid mutator transaction binding the contract method 0x41a5626a.
//
// Solidity: function replaceAdvanceAddresses(address[] oldAddresses_, address[] newAddresses_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) ReplaceAdvanceAddresses(opts *bind.TransactOpts, oldAddresses_ []common.Address, newAddresses_ []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "replaceAdvanceAddresses", oldAddresses_, newAddresses_)
}

// ReplaceAdvanceAddresses is a paid mutator transaction binding the contract method 0x41a5626a.
//
// Solidity: function replaceAdvanceAddresses(address[] oldAddresses_, address[] newAddresses_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) ReplaceAdvanceAddresses(oldAddresses_ []common.Address, newAddresses_ []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.ReplaceAdvanceAddresses(&_FeralfileExhibitionV44.TransactOpts, oldAddresses_, newAddresses_)
}

// ReplaceAdvanceAddresses is a paid mutator transaction binding the contract method 0x41a5626a.
//
// Solidity: function replaceAdvanceAddresses(address[] oldAddresses_, address[] newAddresses_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) ReplaceAdvanceAddresses(oldAddresses_ []common.Address, newAddresses_ []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.ReplaceAdvanceAddresses(&_FeralfileExhibitionV44.TransactOpts, oldAddresses_, newAddresses_)
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) ResumeSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "resumeSale")
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) ResumeSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.ResumeSale(&_FeralfileExhibitionV44.TransactOpts)
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) ResumeSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.ResumeSale(&_FeralfileExhibitionV44.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SafeTransferFrom(&_FeralfileExhibitionV44.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SafeTransferFrom(&_FeralfileExhibitionV44.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SafeTransferFrom0(&_FeralfileExhibitionV44.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SafeTransferFrom0(&_FeralfileExhibitionV44.TransactOpts, from, to, tokenId, data)
}

// SetAdvanceSetting is a paid mutator transaction binding the contract method 0x3c352b0d.
//
// Solidity: function setAdvanceSetting(address[] addresses_, uint256[] amounts_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SetAdvanceSetting(opts *bind.TransactOpts, addresses_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "setAdvanceSetting", addresses_, amounts_)
}

// SetAdvanceSetting is a paid mutator transaction binding the contract method 0x3c352b0d.
//
// Solidity: function setAdvanceSetting(address[] addresses_, uint256[] amounts_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SetAdvanceSetting(addresses_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetAdvanceSetting(&_FeralfileExhibitionV44.TransactOpts, addresses_, amounts_)
}

// SetAdvanceSetting is a paid mutator transaction binding the contract method 0x3c352b0d.
//
// Solidity: function setAdvanceSetting(address[] addresses_, uint256[] amounts_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SetAdvanceSetting(addresses_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetAdvanceSetting(&_FeralfileExhibitionV44.TransactOpts, addresses_, amounts_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetApprovalForAll(&_FeralfileExhibitionV44.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetApprovalForAll(&_FeralfileExhibitionV44.TransactOpts, operator, approved)
}

// SetCostReceiver is a paid mutator transaction binding the contract method 0x1623528f.
//
// Solidity: function setCostReceiver(address costReceiver_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SetCostReceiver(opts *bind.TransactOpts, costReceiver_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "setCostReceiver", costReceiver_)
}

// SetCostReceiver is a paid mutator transaction binding the contract method 0x1623528f.
//
// Solidity: function setCostReceiver(address costReceiver_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SetCostReceiver(costReceiver_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetCostReceiver(&_FeralfileExhibitionV44.TransactOpts, costReceiver_)
}

// SetCostReceiver is a paid mutator transaction binding the contract method 0x1623528f.
//
// Solidity: function setCostReceiver(address costReceiver_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SetCostReceiver(costReceiver_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetCostReceiver(&_FeralfileExhibitionV44.TransactOpts, costReceiver_)
}

// SetRendererTokenData is a paid mutator transaction binding the contract method 0x32f37c2e.
//
// Solidity: function setRendererTokenData(uint256[] tokenIds, (string,string,string)[] data) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SetRendererTokenData(opts *bind.TransactOpts, tokenIds []*big.Int, data []FeralfileExhibitionV44RendererTokenData) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "setRendererTokenData", tokenIds, data)
}

// SetRendererTokenData is a paid mutator transaction binding the contract method 0x32f37c2e.
//
// Solidity: function setRendererTokenData(uint256[] tokenIds, (string,string,string)[] data) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SetRendererTokenData(tokenIds []*big.Int, data []FeralfileExhibitionV44RendererTokenData) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetRendererTokenData(&_FeralfileExhibitionV44.TransactOpts, tokenIds, data)
}

// SetRendererTokenData is a paid mutator transaction binding the contract method 0x32f37c2e.
//
// Solidity: function setRendererTokenData(uint256[] tokenIds, (string,string,string)[] data) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SetRendererTokenData(tokenIds []*big.Int, data []FeralfileExhibitionV44RendererTokenData) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetRendererTokenData(&_FeralfileExhibitionV44.TransactOpts, tokenIds, data)
}

// SetSeriesNames is a paid mutator transaction binding the contract method 0xfa35e4e2.
//
// Solidity: function setSeriesNames(uint256[] seriesIds, string[] names) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SetSeriesNames(opts *bind.TransactOpts, seriesIds []*big.Int, names []string) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "setSeriesNames", seriesIds, names)
}

// SetSeriesNames is a paid mutator transaction binding the contract method 0xfa35e4e2.
//
// Solidity: function setSeriesNames(uint256[] seriesIds, string[] names) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SetSeriesNames(seriesIds []*big.Int, names []string) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetSeriesNames(&_FeralfileExhibitionV44.TransactOpts, seriesIds, names)
}

// SetSeriesNames is a paid mutator transaction binding the contract method 0xfa35e4e2.
//
// Solidity: function setSeriesNames(uint256[] seriesIds, string[] names) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SetSeriesNames(seriesIds []*big.Int, names []string) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetSeriesNames(&_FeralfileExhibitionV44.TransactOpts, seriesIds, names)
}

// SetSeriesRenderer is a paid mutator transaction binding the contract method 0x771ac303.
//
// Solidity: function setSeriesRenderer(uint256 seriesId, bytes blob) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SetSeriesRenderer(opts *bind.TransactOpts, seriesId *big.Int, blob []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "setSeriesRenderer", seriesId, blob)
}

// SetSeriesRenderer is a paid mutator transaction binding the contract method 0x771ac303.
//
// Solidity: function setSeriesRenderer(uint256 seriesId, bytes blob) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SetSeriesRenderer(seriesId *big.Int, blob []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetSeriesRenderer(&_FeralfileExhibitionV44.TransactOpts, seriesId, blob)
}

// SetSeriesRenderer is a paid mutator transaction binding the contract method 0x771ac303.
//
// Solidity: function setSeriesRenderer(uint256 seriesId, bytes blob) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SetSeriesRenderer(seriesId *big.Int, blob []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetSeriesRenderer(&_FeralfileExhibitionV44.TransactOpts, seriesId, blob)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SetSigner(opts *bind.TransactOpts, signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "setSigner", signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetSigner(&_FeralfileExhibitionV44.TransactOpts, signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetSigner(&_FeralfileExhibitionV44.TransactOpts, signer_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SetTokenBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "setTokenBaseURI", baseURI_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SetTokenBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetTokenBaseURI(&_FeralfileExhibitionV44.TransactOpts, baseURI_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SetTokenBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetTokenBaseURI(&_FeralfileExhibitionV44.TransactOpts, baseURI_)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vault_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) SetVault(opts *bind.TransactOpts, vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "setVault", vault_)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vault_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) SetVault(vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetVault(&_FeralfileExhibitionV44.TransactOpts, vault_)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vault_) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) SetVault(vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.SetVault(&_FeralfileExhibitionV44.TransactOpts, vault_)
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) StartSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "startSale")
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) StartSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.StartSale(&_FeralfileExhibitionV44.TransactOpts)
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) StartSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.StartSale(&_FeralfileExhibitionV44.TransactOpts)
}

// StopSaleAndBurn is a paid mutator transaction binding the contract method 0xb9b8311a.
//
// Solidity: function stopSaleAndBurn() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) StopSaleAndBurn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "stopSaleAndBurn")
}

// StopSaleAndBurn is a paid mutator transaction binding the contract method 0xb9b8311a.
//
// Solidity: function stopSaleAndBurn() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) StopSaleAndBurn() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.StopSaleAndBurn(&_FeralfileExhibitionV44.TransactOpts)
}

// StopSaleAndBurn is a paid mutator transaction binding the contract method 0xb9b8311a.
//
// Solidity: function stopSaleAndBurn() returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) StopSaleAndBurn() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.StopSaleAndBurn(&_FeralfileExhibitionV44.TransactOpts)
}

// StopSaleAndTransfer is a paid mutator transaction binding the contract method 0x65a46e08.
//
// Solidity: function stopSaleAndTransfer(uint256[] seriesIds, address[] recipientAddresses) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) StopSaleAndTransfer(opts *bind.TransactOpts, seriesIds []*big.Int, recipientAddresses []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "stopSaleAndTransfer", seriesIds, recipientAddresses)
}

// StopSaleAndTransfer is a paid mutator transaction binding the contract method 0x65a46e08.
//
// Solidity: function stopSaleAndTransfer(uint256[] seriesIds, address[] recipientAddresses) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) StopSaleAndTransfer(seriesIds []*big.Int, recipientAddresses []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.StopSaleAndTransfer(&_FeralfileExhibitionV44.TransactOpts, seriesIds, recipientAddresses)
}

// StopSaleAndTransfer is a paid mutator transaction binding the contract method 0x65a46e08.
//
// Solidity: function stopSaleAndTransfer(uint256[] seriesIds, address[] recipientAddresses) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) StopSaleAndTransfer(seriesIds []*big.Int, recipientAddresses []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.StopSaleAndTransfer(&_FeralfileExhibitionV44.TransactOpts, seriesIds, recipientAddresses)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.TransferFrom(&_FeralfileExhibitionV44.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.TransferFrom(&_FeralfileExhibitionV44.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.TransferOwnership(&_FeralfileExhibitionV44.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.TransferOwnership(&_FeralfileExhibitionV44.TransactOpts, newOwner)
}

// UpdateOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x114ba8ee.
//
// Solidity: function updateOperatorFilterRegistry(address operatorFilterRegisterAddress) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) UpdateOperatorFilterRegistry(opts *bind.TransactOpts, operatorFilterRegisterAddress common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.Transact(opts, "updateOperatorFilterRegistry", operatorFilterRegisterAddress)
}

// UpdateOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x114ba8ee.
//
// Solidity: function updateOperatorFilterRegistry(address operatorFilterRegisterAddress) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) UpdateOperatorFilterRegistry(operatorFilterRegisterAddress common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.UpdateOperatorFilterRegistry(&_FeralfileExhibitionV44.TransactOpts, operatorFilterRegisterAddress)
}

// UpdateOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x114ba8ee.
//
// Solidity: function updateOperatorFilterRegistry(address operatorFilterRegisterAddress) returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) UpdateOperatorFilterRegistry(operatorFilterRegisterAddress common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.UpdateOperatorFilterRegistry(&_FeralfileExhibitionV44.TransactOpts, operatorFilterRegisterAddress)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV44.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Session) Receive() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.Receive(&_FeralfileExhibitionV44.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44TransactorSession) Receive() (*types.Transaction, error) {
	return _FeralfileExhibitionV44.Contract.Receive(&_FeralfileExhibitionV44.TransactOpts)
}

// FeralfileExhibitionV44ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44ApprovalIterator struct {
	Event *FeralfileExhibitionV44Approval // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44Approval)
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
		it.Event = new(FeralfileExhibitionV44Approval)
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
func (it *FeralfileExhibitionV44ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44Approval represents a Approval event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV44ApprovalIterator, error) {

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

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44ApprovalIterator{contract: _FeralfileExhibitionV44.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44Approval)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseApproval(log types.Log) (*FeralfileExhibitionV44Approval, error) {
	event := new(FeralfileExhibitionV44Approval)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV44ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44ApprovalForAllIterator struct {
	Event *FeralfileExhibitionV44ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44ApprovalForAll)
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
		it.Event = new(FeralfileExhibitionV44ApprovalForAll)
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
func (it *FeralfileExhibitionV44ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44ApprovalForAll represents a ApprovalForAll event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*FeralfileExhibitionV44ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44ApprovalForAllIterator{contract: _FeralfileExhibitionV44.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44ApprovalForAll)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseApprovalForAll(log types.Log) (*FeralfileExhibitionV44ApprovalForAll, error) {
	event := new(FeralfileExhibitionV44ApprovalForAll)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV44BurnArtworkIterator is returned from FilterBurnArtwork and is used to iterate over the raw logs and unpacked data for BurnArtwork events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44BurnArtworkIterator struct {
	Event *FeralfileExhibitionV44BurnArtwork // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44BurnArtworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44BurnArtwork)
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
		it.Event = new(FeralfileExhibitionV44BurnArtwork)
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
func (it *FeralfileExhibitionV44BurnArtworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44BurnArtworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44BurnArtwork represents a BurnArtwork event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44BurnArtwork struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnArtwork is a free log retrieval operation binding the contract event 0xbde7938970372996ff103863625e348ef2bf8f38a5b02181be75aafef17c23d5.
//
// Solidity: event BurnArtwork(uint256 indexed tokenId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterBurnArtwork(opts *bind.FilterOpts, tokenId []*big.Int) (*FeralfileExhibitionV44BurnArtworkIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "BurnArtwork", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44BurnArtworkIterator{contract: _FeralfileExhibitionV44.contract, event: "BurnArtwork", logs: logs, sub: sub}, nil
}

// WatchBurnArtwork is a free log subscription operation binding the contract event 0xbde7938970372996ff103863625e348ef2bf8f38a5b02181be75aafef17c23d5.
//
// Solidity: event BurnArtwork(uint256 indexed tokenId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchBurnArtwork(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44BurnArtwork, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "BurnArtwork", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44BurnArtwork)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "BurnArtwork", log); err != nil {
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
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseBurnArtwork(log types.Log) (*FeralfileExhibitionV44BurnArtwork, error) {
	event := new(FeralfileExhibitionV44BurnArtwork)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "BurnArtwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV44BuyArtworkIterator is returned from FilterBuyArtwork and is used to iterate over the raw logs and unpacked data for BuyArtwork events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44BuyArtworkIterator struct {
	Event *FeralfileExhibitionV44BuyArtwork // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44BuyArtworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44BuyArtwork)
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
		it.Event = new(FeralfileExhibitionV44BuyArtwork)
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
func (it *FeralfileExhibitionV44BuyArtworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44BuyArtworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44BuyArtwork represents a BuyArtwork event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44BuyArtwork struct {
	Buyer   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyArtwork is a free log retrieval operation binding the contract event 0x0475389cd69b8d3163620b43283bf74e8fc71020c3c6cef2a529b5c405e9687f.
//
// Solidity: event BuyArtwork(address indexed buyer, uint256 indexed tokenId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterBuyArtwork(opts *bind.FilterOpts, buyer []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV44BuyArtworkIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "BuyArtwork", buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44BuyArtworkIterator{contract: _FeralfileExhibitionV44.contract, event: "BuyArtwork", logs: logs, sub: sub}, nil
}

// WatchBuyArtwork is a free log subscription operation binding the contract event 0x0475389cd69b8d3163620b43283bf74e8fc71020c3c6cef2a529b5c405e9687f.
//
// Solidity: event BuyArtwork(address indexed buyer, uint256 indexed tokenId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchBuyArtwork(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44BuyArtwork, buyer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "BuyArtwork", buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44BuyArtwork)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "BuyArtwork", log); err != nil {
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
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseBuyArtwork(log types.Log) (*FeralfileExhibitionV44BuyArtwork, error) {
	event := new(FeralfileExhibitionV44BuyArtwork)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "BuyArtwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV44DeleteSeriesRendererIterator is returned from FilterDeleteSeriesRenderer and is used to iterate over the raw logs and unpacked data for DeleteSeriesRenderer events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44DeleteSeriesRendererIterator struct {
	Event *FeralfileExhibitionV44DeleteSeriesRenderer // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44DeleteSeriesRendererIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44DeleteSeriesRenderer)
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
		it.Event = new(FeralfileExhibitionV44DeleteSeriesRenderer)
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
func (it *FeralfileExhibitionV44DeleteSeriesRendererIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44DeleteSeriesRendererIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44DeleteSeriesRenderer represents a DeleteSeriesRenderer event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44DeleteSeriesRenderer struct {
	SeriesId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteSeriesRenderer is a free log retrieval operation binding the contract event 0x5642c40d50e2073be9e5183e69e9c92f89328520b1a18fcd0d109b25434bd9d2.
//
// Solidity: event DeleteSeriesRenderer(uint256 indexed seriesId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterDeleteSeriesRenderer(opts *bind.FilterOpts, seriesId []*big.Int) (*FeralfileExhibitionV44DeleteSeriesRendererIterator, error) {

	var seriesIdRule []interface{}
	for _, seriesIdItem := range seriesId {
		seriesIdRule = append(seriesIdRule, seriesIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "DeleteSeriesRenderer", seriesIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44DeleteSeriesRendererIterator{contract: _FeralfileExhibitionV44.contract, event: "DeleteSeriesRenderer", logs: logs, sub: sub}, nil
}

// WatchDeleteSeriesRenderer is a free log subscription operation binding the contract event 0x5642c40d50e2073be9e5183e69e9c92f89328520b1a18fcd0d109b25434bd9d2.
//
// Solidity: event DeleteSeriesRenderer(uint256 indexed seriesId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchDeleteSeriesRenderer(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44DeleteSeriesRenderer, seriesId []*big.Int) (event.Subscription, error) {

	var seriesIdRule []interface{}
	for _, seriesIdItem := range seriesId {
		seriesIdRule = append(seriesIdRule, seriesIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "DeleteSeriesRenderer", seriesIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44DeleteSeriesRenderer)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "DeleteSeriesRenderer", log); err != nil {
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

// ParseDeleteSeriesRenderer is a log parse operation binding the contract event 0x5642c40d50e2073be9e5183e69e9c92f89328520b1a18fcd0d109b25434bd9d2.
//
// Solidity: event DeleteSeriesRenderer(uint256 indexed seriesId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseDeleteSeriesRenderer(log types.Log) (*FeralfileExhibitionV44DeleteSeriesRenderer, error) {
	event := new(FeralfileExhibitionV44DeleteSeriesRenderer)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "DeleteSeriesRenderer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV44NewArtworkIterator is returned from FilterNewArtwork and is used to iterate over the raw logs and unpacked data for NewArtwork events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44NewArtworkIterator struct {
	Event *FeralfileExhibitionV44NewArtwork // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44NewArtworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44NewArtwork)
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
		it.Event = new(FeralfileExhibitionV44NewArtwork)
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
func (it *FeralfileExhibitionV44NewArtworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44NewArtworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44NewArtwork represents a NewArtwork event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44NewArtwork struct {
	Owner    common.Address
	SeriesId *big.Int
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewArtwork is a free log retrieval operation binding the contract event 0x407d7da1d3b2b1871fbfa2b5b1c4657a3cc5711d3023c552798551c7ee301eea.
//
// Solidity: event NewArtwork(address indexed owner, uint256 indexed seriesId, uint256 indexed tokenId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterNewArtwork(opts *bind.FilterOpts, owner []common.Address, seriesId []*big.Int, tokenId []*big.Int) (*FeralfileExhibitionV44NewArtworkIterator, error) {

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

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "NewArtwork", ownerRule, seriesIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44NewArtworkIterator{contract: _FeralfileExhibitionV44.contract, event: "NewArtwork", logs: logs, sub: sub}, nil
}

// WatchNewArtwork is a free log subscription operation binding the contract event 0x407d7da1d3b2b1871fbfa2b5b1c4657a3cc5711d3023c552798551c7ee301eea.
//
// Solidity: event NewArtwork(address indexed owner, uint256 indexed seriesId, uint256 indexed tokenId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchNewArtwork(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44NewArtwork, owner []common.Address, seriesId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "NewArtwork", ownerRule, seriesIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44NewArtwork)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "NewArtwork", log); err != nil {
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
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseNewArtwork(log types.Log) (*FeralfileExhibitionV44NewArtwork, error) {
	event := new(FeralfileExhibitionV44NewArtwork)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "NewArtwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV44NewRendererTokenDataIterator is returned from FilterNewRendererTokenData and is used to iterate over the raw logs and unpacked data for NewRendererTokenData events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44NewRendererTokenDataIterator struct {
	Event *FeralfileExhibitionV44NewRendererTokenData // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44NewRendererTokenDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44NewRendererTokenData)
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
		it.Event = new(FeralfileExhibitionV44NewRendererTokenData)
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
func (it *FeralfileExhibitionV44NewRendererTokenDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44NewRendererTokenDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44NewRendererTokenData represents a NewRendererTokenData event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44NewRendererTokenData struct {
	TokenId *big.Int
	Data    FeralfileExhibitionV44RendererTokenData
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewRendererTokenData is a free log retrieval operation binding the contract event 0xac24f5d052fdb3e47d0446513e40bbf8afcfd59e1bd778f60c1f53da4212ead6.
//
// Solidity: event NewRendererTokenData(uint256 indexed tokenId, (string,string,string) data)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterNewRendererTokenData(opts *bind.FilterOpts, tokenId []*big.Int) (*FeralfileExhibitionV44NewRendererTokenDataIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "NewRendererTokenData", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44NewRendererTokenDataIterator{contract: _FeralfileExhibitionV44.contract, event: "NewRendererTokenData", logs: logs, sub: sub}, nil
}

// WatchNewRendererTokenData is a free log subscription operation binding the contract event 0xac24f5d052fdb3e47d0446513e40bbf8afcfd59e1bd778f60c1f53da4212ead6.
//
// Solidity: event NewRendererTokenData(uint256 indexed tokenId, (string,string,string) data)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchNewRendererTokenData(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44NewRendererTokenData, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "NewRendererTokenData", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44NewRendererTokenData)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "NewRendererTokenData", log); err != nil {
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

// ParseNewRendererTokenData is a log parse operation binding the contract event 0xac24f5d052fdb3e47d0446513e40bbf8afcfd59e1bd778f60c1f53da4212ead6.
//
// Solidity: event NewRendererTokenData(uint256 indexed tokenId, (string,string,string) data)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseNewRendererTokenData(log types.Log) (*FeralfileExhibitionV44NewRendererTokenData, error) {
	event := new(FeralfileExhibitionV44NewRendererTokenData)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "NewRendererTokenData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV44NewSeriesNameIterator is returned from FilterNewSeriesName and is used to iterate over the raw logs and unpacked data for NewSeriesName events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44NewSeriesNameIterator struct {
	Event *FeralfileExhibitionV44NewSeriesName // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44NewSeriesNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44NewSeriesName)
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
		it.Event = new(FeralfileExhibitionV44NewSeriesName)
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
func (it *FeralfileExhibitionV44NewSeriesNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44NewSeriesNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44NewSeriesName represents a NewSeriesName event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44NewSeriesName struct {
	SeriesId *big.Int
	Name     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewSeriesName is a free log retrieval operation binding the contract event 0xa4154606c7e38a6a299a31a214bf3468d191c48e6308a1e849932d2c5aba56b9.
//
// Solidity: event NewSeriesName(uint256 indexed seriesId, string name)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterNewSeriesName(opts *bind.FilterOpts, seriesId []*big.Int) (*FeralfileExhibitionV44NewSeriesNameIterator, error) {

	var seriesIdRule []interface{}
	for _, seriesIdItem := range seriesId {
		seriesIdRule = append(seriesIdRule, seriesIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "NewSeriesName", seriesIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44NewSeriesNameIterator{contract: _FeralfileExhibitionV44.contract, event: "NewSeriesName", logs: logs, sub: sub}, nil
}

// WatchNewSeriesName is a free log subscription operation binding the contract event 0xa4154606c7e38a6a299a31a214bf3468d191c48e6308a1e849932d2c5aba56b9.
//
// Solidity: event NewSeriesName(uint256 indexed seriesId, string name)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchNewSeriesName(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44NewSeriesName, seriesId []*big.Int) (event.Subscription, error) {

	var seriesIdRule []interface{}
	for _, seriesIdItem := range seriesId {
		seriesIdRule = append(seriesIdRule, seriesIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "NewSeriesName", seriesIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44NewSeriesName)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "NewSeriesName", log); err != nil {
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

// ParseNewSeriesName is a log parse operation binding the contract event 0xa4154606c7e38a6a299a31a214bf3468d191c48e6308a1e849932d2c5aba56b9.
//
// Solidity: event NewSeriesName(uint256 indexed seriesId, string name)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseNewSeriesName(log types.Log) (*FeralfileExhibitionV44NewSeriesName, error) {
	event := new(FeralfileExhibitionV44NewSeriesName)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "NewSeriesName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV44NewSeriesRendererIterator is returned from FilterNewSeriesRenderer and is used to iterate over the raw logs and unpacked data for NewSeriesRenderer events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44NewSeriesRendererIterator struct {
	Event *FeralfileExhibitionV44NewSeriesRenderer // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44NewSeriesRendererIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44NewSeriesRenderer)
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
		it.Event = new(FeralfileExhibitionV44NewSeriesRenderer)
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
func (it *FeralfileExhibitionV44NewSeriesRendererIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44NewSeriesRendererIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44NewSeriesRenderer represents a NewSeriesRenderer event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44NewSeriesRenderer struct {
	SeriesId *big.Int
	Pointers []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewSeriesRenderer is a free log retrieval operation binding the contract event 0x0bfaeb43f2350d013eaed4cf979fcf8c33f3471f9a6a69b312882034d3134bbb.
//
// Solidity: event NewSeriesRenderer(uint256 indexed seriesId, address[] pointers)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterNewSeriesRenderer(opts *bind.FilterOpts, seriesId []*big.Int) (*FeralfileExhibitionV44NewSeriesRendererIterator, error) {

	var seriesIdRule []interface{}
	for _, seriesIdItem := range seriesId {
		seriesIdRule = append(seriesIdRule, seriesIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "NewSeriesRenderer", seriesIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44NewSeriesRendererIterator{contract: _FeralfileExhibitionV44.contract, event: "NewSeriesRenderer", logs: logs, sub: sub}, nil
}

// WatchNewSeriesRenderer is a free log subscription operation binding the contract event 0x0bfaeb43f2350d013eaed4cf979fcf8c33f3471f9a6a69b312882034d3134bbb.
//
// Solidity: event NewSeriesRenderer(uint256 indexed seriesId, address[] pointers)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchNewSeriesRenderer(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44NewSeriesRenderer, seriesId []*big.Int) (event.Subscription, error) {

	var seriesIdRule []interface{}
	for _, seriesIdItem := range seriesId {
		seriesIdRule = append(seriesIdRule, seriesIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "NewSeriesRenderer", seriesIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44NewSeriesRenderer)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "NewSeriesRenderer", log); err != nil {
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

// ParseNewSeriesRenderer is a log parse operation binding the contract event 0x0bfaeb43f2350d013eaed4cf979fcf8c33f3471f9a6a69b312882034d3134bbb.
//
// Solidity: event NewSeriesRenderer(uint256 indexed seriesId, address[] pointers)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseNewSeriesRenderer(log types.Log) (*FeralfileExhibitionV44NewSeriesRenderer, error) {
	event := new(FeralfileExhibitionV44NewSeriesRenderer)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "NewSeriesRenderer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV44OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44OwnershipTransferredIterator struct {
	Event *FeralfileExhibitionV44OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44OwnershipTransferred)
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
		it.Event = new(FeralfileExhibitionV44OwnershipTransferred)
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
func (it *FeralfileExhibitionV44OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44OwnershipTransferred represents a OwnershipTransferred event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeralfileExhibitionV44OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44OwnershipTransferredIterator{contract: _FeralfileExhibitionV44.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44OwnershipTransferred)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseOwnershipTransferred(log types.Log) (*FeralfileExhibitionV44OwnershipTransferred, error) {
	event := new(FeralfileExhibitionV44OwnershipTransferred)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV44TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44TransferIterator struct {
	Event *FeralfileExhibitionV44Transfer // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV44TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV44Transfer)
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
		it.Event = new(FeralfileExhibitionV44Transfer)
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
func (it *FeralfileExhibitionV44TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV44TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV44Transfer represents a Transfer event raised by the FeralfileExhibitionV44 contract.
type FeralfileExhibitionV44Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV44TransferIterator, error) {

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

	logs, sub, err := _FeralfileExhibitionV44.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV44TransferIterator{contract: _FeralfileExhibitionV44.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV44Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _FeralfileExhibitionV44.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV44Transfer)
				if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_FeralfileExhibitionV44 *FeralfileExhibitionV44Filterer) ParseTransfer(log types.Log) (*FeralfileExhibitionV44Transfer, error) {
	event := new(FeralfileExhibitionV44Transfer)
	if err := _FeralfileExhibitionV44.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
