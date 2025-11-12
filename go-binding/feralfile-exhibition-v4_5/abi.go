// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feralfilev4_5

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

// FeralfileExhibitionV45TokenIndex is an auto generated low-level Go binding around an user-defined struct.
type FeralfileExhibitionV45TokenIndex struct {
	TokenId *big.Int
	Index   *big.Int
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

// FeralfileExhibitionV45MetaData contains all meta data concerning the FeralfileExhibitionV45 contract.
var FeralfileExhibitionV45MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"burnable_\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"bridgeable_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"costReceiver_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"seriesIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"seriesMaxSupplies_\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"tokenIdPrefixShard_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdvanceAddressAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAddressesAndAmounts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdvanceAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BurnArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BuyArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NewArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OperatorFilterRegistry\",\"outputs\":[{\"internalType\":\"contractIOperatorFilterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"addTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"advances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burnArtworks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"internalType\":\"structIFeralfileSaleData.RevenueShare[][]\",\"name\":\"revenueShares\",\"type\":\"tuple[][]\"},{\"internalType\":\"bool\",\"name\":\"payByVaultContract\",\"type\":\"bool\"}],\"internalType\":\"structIFeralfileSaleData.SaleData\",\"name\":\"saleData_\",\"type\":\"tuple\"}],\"name\":\"buyArtworks\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"codeVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"costReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getArtwork\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structFeralfileExhibitionV4.Artwork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structFeralfileExhibitionV4.MintData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"mintArtworks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"removeTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oldAddresses_\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"newAddresses_\",\"type\":\"address[]\"}],\"name\":\"replaceAdvanceAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selling\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"seriesMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"}],\"name\":\"seriesTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"setAdvanceSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"costReceiver_\",\"type\":\"address\"}],\"name\":\"setCostReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setTokenBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopSaleAndBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"seriesIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"recipientAddresses\",\"type\":\"address[]\"}],\"name\":\"stopSaleAndTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdPrefixShard\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorFilterRegisterAddress\",\"type\":\"address\"}],\"name\":\"updateOperatorFilterRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIFeralfileVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"tokenIndexesByOwner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structFeralfileExhibitionV4_5.TokenIndex[]\",\"name\":\"tokenIndexes\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getArtworkIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"artworkIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FeralfileExhibitionV45ABI is the input ABI used to generate the binding from.
// Deprecated: Use FeralfileExhibitionV45MetaData.ABI instead.
var FeralfileExhibitionV45ABI = FeralfileExhibitionV45MetaData.ABI

// FeralfileExhibitionV45 is an auto generated Go binding around an Ethereum contract.
type FeralfileExhibitionV45 struct {
	FeralfileExhibitionV45Caller     // Read-only binding to the contract
	FeralfileExhibitionV45Transactor // Write-only binding to the contract
	FeralfileExhibitionV45Filterer   // Log filterer for contract events
}

// FeralfileExhibitionV45Caller is an auto generated read-only Go binding around an Ethereum contract.
type FeralfileExhibitionV45Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV45Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FeralfileExhibitionV45Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV45Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeralfileExhibitionV45Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV45Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeralfileExhibitionV45Session struct {
	Contract     *FeralfileExhibitionV45 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FeralfileExhibitionV45CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeralfileExhibitionV45CallerSession struct {
	Contract *FeralfileExhibitionV45Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// FeralfileExhibitionV45TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeralfileExhibitionV45TransactorSession struct {
	Contract     *FeralfileExhibitionV45Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// FeralfileExhibitionV45Raw is an auto generated low-level Go binding around an Ethereum contract.
type FeralfileExhibitionV45Raw struct {
	Contract *FeralfileExhibitionV45 // Generic contract binding to access the raw methods on
}

// FeralfileExhibitionV45CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeralfileExhibitionV45CallerRaw struct {
	Contract *FeralfileExhibitionV45Caller // Generic read-only contract binding to access the raw methods on
}

// FeralfileExhibitionV45TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeralfileExhibitionV45TransactorRaw struct {
	Contract *FeralfileExhibitionV45Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFeralfileExhibitionV45 creates a new instance of FeralfileExhibitionV45, bound to a specific deployed contract.
func NewFeralfileExhibitionV45(address common.Address, backend bind.ContractBackend) (*FeralfileExhibitionV45, error) {
	contract, err := bindFeralfileExhibitionV45(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45{FeralfileExhibitionV45Caller: FeralfileExhibitionV45Caller{contract: contract}, FeralfileExhibitionV45Transactor: FeralfileExhibitionV45Transactor{contract: contract}, FeralfileExhibitionV45Filterer: FeralfileExhibitionV45Filterer{contract: contract}}, nil
}

// NewFeralfileExhibitionV45Caller creates a new read-only instance of FeralfileExhibitionV45, bound to a specific deployed contract.
func NewFeralfileExhibitionV45Caller(address common.Address, caller bind.ContractCaller) (*FeralfileExhibitionV45Caller, error) {
	contract, err := bindFeralfileExhibitionV45(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45Caller{contract: contract}, nil
}

// NewFeralfileExhibitionV45Transactor creates a new write-only instance of FeralfileExhibitionV45, bound to a specific deployed contract.
func NewFeralfileExhibitionV45Transactor(address common.Address, transactor bind.ContractTransactor) (*FeralfileExhibitionV45Transactor, error) {
	contract, err := bindFeralfileExhibitionV45(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45Transactor{contract: contract}, nil
}

// NewFeralfileExhibitionV45Filterer creates a new log filterer instance of FeralfileExhibitionV45, bound to a specific deployed contract.
func NewFeralfileExhibitionV45Filterer(address common.Address, filterer bind.ContractFilterer) (*FeralfileExhibitionV45Filterer, error) {
	contract, err := bindFeralfileExhibitionV45(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45Filterer{contract: contract}, nil
}

// bindFeralfileExhibitionV45 binds a generic wrapper to an already deployed contract.
func bindFeralfileExhibitionV45(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeralfileExhibitionV45MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralfileExhibitionV45.Contract.FeralfileExhibitionV45Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.FeralfileExhibitionV45Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.FeralfileExhibitionV45Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralfileExhibitionV45.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.contract.Transact(opts, method, params...)
}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xf07e7fd0.
//
// Solidity: function OperatorFilterRegistry() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) OperatorFilterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "OperatorFilterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xf07e7fd0.
//
// Solidity: function OperatorFilterRegistry() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) OperatorFilterRegistry() (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.OperatorFilterRegistry(&_FeralfileExhibitionV45.CallOpts)
}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xf07e7fd0.
//
// Solidity: function OperatorFilterRegistry() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) OperatorFilterRegistry() (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.OperatorFilterRegistry(&_FeralfileExhibitionV45.CallOpts)
}

// Advances is a free data retrieval call binding the contract method 0x926ce44e.
//
// Solidity: function advances(address ) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Advances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "advances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Advances is a free data retrieval call binding the contract method 0x926ce44e.
//
// Solidity: function advances(address ) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Advances(arg0 common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.Advances(&_FeralfileExhibitionV45.CallOpts, arg0)
}

// Advances is a free data retrieval call binding the contract method 0x926ce44e.
//
// Solidity: function advances(address ) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Advances(arg0 common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.Advances(&_FeralfileExhibitionV45.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.BalanceOf(&_FeralfileExhibitionV45.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.BalanceOf(&_FeralfileExhibitionV45.CallOpts, owner)
}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Bridgeable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "bridgeable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Bridgeable() (bool, error) {
	return _FeralfileExhibitionV45.Contract.Bridgeable(&_FeralfileExhibitionV45.CallOpts)
}

// Bridgeable is a free data retrieval call binding the contract method 0x530da8ef.
//
// Solidity: function bridgeable() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Bridgeable() (bool, error) {
	return _FeralfileExhibitionV45.Contract.Bridgeable(&_FeralfileExhibitionV45.CallOpts)
}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Burnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "burnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Burnable() (bool, error) {
	return _FeralfileExhibitionV45.Contract.Burnable(&_FeralfileExhibitionV45.CallOpts)
}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Burnable() (bool, error) {
	return _FeralfileExhibitionV45.Contract.Burnable(&_FeralfileExhibitionV45.CallOpts)
}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) CodeVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "codeVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) CodeVersion() (string, error) {
	return _FeralfileExhibitionV45.Contract.CodeVersion(&_FeralfileExhibitionV45.CallOpts)
}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) CodeVersion() (string, error) {
	return _FeralfileExhibitionV45.Contract.CodeVersion(&_FeralfileExhibitionV45.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) ContractURI() (string, error) {
	return _FeralfileExhibitionV45.Contract.ContractURI(&_FeralfileExhibitionV45.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) ContractURI() (string, error) {
	return _FeralfileExhibitionV45.Contract.ContractURI(&_FeralfileExhibitionV45.CallOpts)
}

// CostReceiver is a free data retrieval call binding the contract method 0xf4e638be.
//
// Solidity: function costReceiver() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) CostReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "costReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CostReceiver is a free data retrieval call binding the contract method 0xf4e638be.
//
// Solidity: function costReceiver() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) CostReceiver() (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.CostReceiver(&_FeralfileExhibitionV45.CallOpts)
}

// CostReceiver is a free data retrieval call binding the contract method 0xf4e638be.
//
// Solidity: function costReceiver() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) CostReceiver() (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.CostReceiver(&_FeralfileExhibitionV45.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.GetApproved(&_FeralfileExhibitionV45.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.GetApproved(&_FeralfileExhibitionV45.CallOpts, tokenId)
}

// GetArtwork is a free data retrieval call binding the contract method 0x167ddf6e.
//
// Solidity: function getArtwork(uint256 tokenId) view returns((uint256,uint256))
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) GetArtwork(opts *bind.CallOpts, tokenId *big.Int) (FeralfileExhibitionV4Artwork, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "getArtwork", tokenId)

	if err != nil {
		return *new(FeralfileExhibitionV4Artwork), err
	}

	out0 := *abi.ConvertType(out[0], new(FeralfileExhibitionV4Artwork)).(*FeralfileExhibitionV4Artwork)

	return out0, err

}

// GetArtwork is a free data retrieval call binding the contract method 0x167ddf6e.
//
// Solidity: function getArtwork(uint256 tokenId) view returns((uint256,uint256))
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) GetArtwork(tokenId *big.Int) (FeralfileExhibitionV4Artwork, error) {
	return _FeralfileExhibitionV45.Contract.GetArtwork(&_FeralfileExhibitionV45.CallOpts, tokenId)
}

// GetArtwork is a free data retrieval call binding the contract method 0x167ddf6e.
//
// Solidity: function getArtwork(uint256 tokenId) view returns((uint256,uint256))
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) GetArtwork(tokenId *big.Int) (FeralfileExhibitionV4Artwork, error) {
	return _FeralfileExhibitionV45.Contract.GetArtwork(&_FeralfileExhibitionV45.CallOpts, tokenId)
}

// GetArtworkIndex is a free data retrieval call binding the contract method 0x2ab9b963.
//
// Solidity: function getArtworkIndex(uint256 tokenId_) view returns(uint256 artworkIndex)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) GetArtworkIndex(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "getArtworkIndex", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetArtworkIndex is a free data retrieval call binding the contract method 0x2ab9b963.
//
// Solidity: function getArtworkIndex(uint256 tokenId_) view returns(uint256 artworkIndex)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) GetArtworkIndex(tokenId_ *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.GetArtworkIndex(&_FeralfileExhibitionV45.CallOpts, tokenId_)
}

// GetArtworkIndex is a free data retrieval call binding the contract method 0x2ab9b963.
//
// Solidity: function getArtworkIndex(uint256 tokenId_) view returns(uint256 artworkIndex)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) GetArtworkIndex(tokenId_ *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.GetArtworkIndex(&_FeralfileExhibitionV45.CallOpts, tokenId_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FeralfileExhibitionV45.Contract.IsApprovedForAll(&_FeralfileExhibitionV45.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FeralfileExhibitionV45.Contract.IsApprovedForAll(&_FeralfileExhibitionV45.CallOpts, owner, operator)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Mintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "mintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Mintable() (bool, error) {
	return _FeralfileExhibitionV45.Contract.Mintable(&_FeralfileExhibitionV45.CallOpts)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Mintable() (bool, error) {
	return _FeralfileExhibitionV45.Contract.Mintable(&_FeralfileExhibitionV45.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Name() (string, error) {
	return _FeralfileExhibitionV45.Contract.Name(&_FeralfileExhibitionV45.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Name() (string, error) {
	return _FeralfileExhibitionV45.Contract.Name(&_FeralfileExhibitionV45.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Owner() (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.Owner(&_FeralfileExhibitionV45.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Owner() (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.Owner(&_FeralfileExhibitionV45.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.OwnerOf(&_FeralfileExhibitionV45.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.OwnerOf(&_FeralfileExhibitionV45.CallOpts, tokenId)
}

// Selling is a free data retrieval call binding the contract method 0x23aed228.
//
// Solidity: function selling() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Selling(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "selling")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Selling is a free data retrieval call binding the contract method 0x23aed228.
//
// Solidity: function selling() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Selling() (bool, error) {
	return _FeralfileExhibitionV45.Contract.Selling(&_FeralfileExhibitionV45.CallOpts)
}

// Selling is a free data retrieval call binding the contract method 0x23aed228.
//
// Solidity: function selling() view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Selling() (bool, error) {
	return _FeralfileExhibitionV45.Contract.Selling(&_FeralfileExhibitionV45.CallOpts)
}

// SeriesMaxSupply is a free data retrieval call binding the contract method 0xeb5c60f2.
//
// Solidity: function seriesMaxSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) SeriesMaxSupply(opts *bind.CallOpts, seriesId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "seriesMaxSupply", seriesId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeriesMaxSupply is a free data retrieval call binding the contract method 0xeb5c60f2.
//
// Solidity: function seriesMaxSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SeriesMaxSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.SeriesMaxSupply(&_FeralfileExhibitionV45.CallOpts, seriesId)
}

// SeriesMaxSupply is a free data retrieval call binding the contract method 0xeb5c60f2.
//
// Solidity: function seriesMaxSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) SeriesMaxSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.SeriesMaxSupply(&_FeralfileExhibitionV45.CallOpts, seriesId)
}

// SeriesTotalSupply is a free data retrieval call binding the contract method 0x7f06ee06.
//
// Solidity: function seriesTotalSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) SeriesTotalSupply(opts *bind.CallOpts, seriesId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "seriesTotalSupply", seriesId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeriesTotalSupply is a free data retrieval call binding the contract method 0x7f06ee06.
//
// Solidity: function seriesTotalSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SeriesTotalSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.SeriesTotalSupply(&_FeralfileExhibitionV45.CallOpts, seriesId)
}

// SeriesTotalSupply is a free data retrieval call binding the contract method 0x7f06ee06.
//
// Solidity: function seriesTotalSupply(uint256 seriesId) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) SeriesTotalSupply(seriesId *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.SeriesTotalSupply(&_FeralfileExhibitionV45.CallOpts, seriesId)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Signer() (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.Signer(&_FeralfileExhibitionV45.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Signer() (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.Signer(&_FeralfileExhibitionV45.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeralfileExhibitionV45.Contract.SupportsInterface(&_FeralfileExhibitionV45.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeralfileExhibitionV45.Contract.SupportsInterface(&_FeralfileExhibitionV45.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Symbol() (string, error) {
	return _FeralfileExhibitionV45.Contract.Symbol(&_FeralfileExhibitionV45.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Symbol() (string, error) {
	return _FeralfileExhibitionV45.Contract.Symbol(&_FeralfileExhibitionV45.CallOpts)
}

// TokenBaseURI is a free data retrieval call binding the contract method 0x4e99b800.
//
// Solidity: function tokenBaseURI() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) TokenBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "tokenBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenBaseURI is a free data retrieval call binding the contract method 0x4e99b800.
//
// Solidity: function tokenBaseURI() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) TokenBaseURI() (string, error) {
	return _FeralfileExhibitionV45.Contract.TokenBaseURI(&_FeralfileExhibitionV45.CallOpts)
}

// TokenBaseURI is a free data retrieval call binding the contract method 0x4e99b800.
//
// Solidity: function tokenBaseURI() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) TokenBaseURI() (string, error) {
	return _FeralfileExhibitionV45.Contract.TokenBaseURI(&_FeralfileExhibitionV45.CallOpts)
}

// TokenIdPrefixShard is a free data retrieval call binding the contract method 0x83aeaa36.
//
// Solidity: function tokenIdPrefixShard() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) TokenIdPrefixShard(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "tokenIdPrefixShard")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenIdPrefixShard is a free data retrieval call binding the contract method 0x83aeaa36.
//
// Solidity: function tokenIdPrefixShard() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) TokenIdPrefixShard() (string, error) {
	return _FeralfileExhibitionV45.Contract.TokenIdPrefixShard(&_FeralfileExhibitionV45.CallOpts)
}

// TokenIdPrefixShard is a free data retrieval call binding the contract method 0x83aeaa36.
//
// Solidity: function tokenIdPrefixShard() view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) TokenIdPrefixShard() (string, error) {
	return _FeralfileExhibitionV45.Contract.TokenIdPrefixShard(&_FeralfileExhibitionV45.CallOpts)
}

// TokenIndexesByOwner is a free data retrieval call binding the contract method 0x29dd4115.
//
// Solidity: function tokenIndexesByOwner(uint256[] tokenIds_) view returns((uint256,uint256)[] tokenIndexes)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) TokenIndexesByOwner(opts *bind.CallOpts, tokenIds_ []*big.Int) ([]FeralfileExhibitionV45TokenIndex, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "tokenIndexesByOwner", tokenIds_)

	if err != nil {
		return *new([]FeralfileExhibitionV45TokenIndex), err
	}

	out0 := *abi.ConvertType(out[0], new([]FeralfileExhibitionV45TokenIndex)).(*[]FeralfileExhibitionV45TokenIndex)

	return out0, err

}

// TokenIndexesByOwner is a free data retrieval call binding the contract method 0x29dd4115.
//
// Solidity: function tokenIndexesByOwner(uint256[] tokenIds_) view returns((uint256,uint256)[] tokenIndexes)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) TokenIndexesByOwner(tokenIds_ []*big.Int) ([]FeralfileExhibitionV45TokenIndex, error) {
	return _FeralfileExhibitionV45.Contract.TokenIndexesByOwner(&_FeralfileExhibitionV45.CallOpts, tokenIds_)
}

// TokenIndexesByOwner is a free data retrieval call binding the contract method 0x29dd4115.
//
// Solidity: function tokenIndexesByOwner(uint256[] tokenIds_) view returns((uint256,uint256)[] tokenIndexes)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) TokenIndexesByOwner(tokenIds_ []*big.Int) ([]FeralfileExhibitionV45TokenIndex, error) {
	return _FeralfileExhibitionV45.Contract.TokenIndexesByOwner(&_FeralfileExhibitionV45.CallOpts, tokenIds_)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.TokenOfOwnerByIndex(&_FeralfileExhibitionV45.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.TokenOfOwnerByIndex(&_FeralfileExhibitionV45.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) TokenURI(tokenId *big.Int) (string, error) {
	return _FeralfileExhibitionV45.Contract.TokenURI(&_FeralfileExhibitionV45.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _FeralfileExhibitionV45.Contract.TokenURI(&_FeralfileExhibitionV45.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.TokensOfOwner(&_FeralfileExhibitionV45.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.TokensOfOwner(&_FeralfileExhibitionV45.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) TotalSupply() (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.TotalSupply(&_FeralfileExhibitionV45.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) TotalSupply() (*big.Int, error) {
	return _FeralfileExhibitionV45.Contract.TotalSupply(&_FeralfileExhibitionV45.CallOpts)
}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Trustees(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "trustees", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Trustees(arg0 common.Address) (bool, error) {
	return _FeralfileExhibitionV45.Contract.Trustees(&_FeralfileExhibitionV45.CallOpts, arg0)
}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Trustees(arg0 common.Address) (bool, error) {
	return _FeralfileExhibitionV45.Contract.Trustees(&_FeralfileExhibitionV45.CallOpts, arg0)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Caller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV45.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Vault() (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.Vault(&_FeralfileExhibitionV45.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45CallerSession) Vault() (common.Address, error) {
	return _FeralfileExhibitionV45.Contract.Vault(&_FeralfileExhibitionV45.CallOpts)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) AddTrustee(opts *bind.TransactOpts, _trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "addTrustee", _trustee)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) AddTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.AddTrustee(&_FeralfileExhibitionV45.TransactOpts, _trustee)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) AddTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.AddTrustee(&_FeralfileExhibitionV45.TransactOpts, _trustee)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.Approve(&_FeralfileExhibitionV45.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.Approve(&_FeralfileExhibitionV45.TransactOpts, operator, tokenId)
}

// BurnArtworks is a paid mutator transaction binding the contract method 0x21fe0c64.
//
// Solidity: function burnArtworks(uint256[] tokenIds) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) BurnArtworks(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "burnArtworks", tokenIds)
}

// BurnArtworks is a paid mutator transaction binding the contract method 0x21fe0c64.
//
// Solidity: function burnArtworks(uint256[] tokenIds) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) BurnArtworks(tokenIds []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.BurnArtworks(&_FeralfileExhibitionV45.TransactOpts, tokenIds)
}

// BurnArtworks is a paid mutator transaction binding the contract method 0x21fe0c64.
//
// Solidity: function burnArtworks(uint256[] tokenIds) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) BurnArtworks(tokenIds []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.BurnArtworks(&_FeralfileExhibitionV45.TransactOpts, tokenIds)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_) payable returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) BuyArtworks(opts *bind.TransactOpts, r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "buyArtworks", r_, s_, v_, saleData_)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_) payable returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) BuyArtworks(r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.BuyArtworks(&_FeralfileExhibitionV45.TransactOpts, r_, s_, v_, saleData_)
}

// BuyArtworks is a paid mutator transaction binding the contract method 0x2977e4b3.
//
// Solidity: function buyArtworks(bytes32 r_, bytes32 s_, uint8 v_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_) payable returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) BuyArtworks(r_ [32]byte, s_ [32]byte, v_ uint8, saleData_ IFeralfileSaleDataSaleData) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.BuyArtworks(&_FeralfileExhibitionV45.TransactOpts, r_, s_, v_, saleData_)
}

// MintArtworks is a paid mutator transaction binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] data) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) MintArtworks(opts *bind.TransactOpts, data []FeralfileExhibitionV4MintData) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "mintArtworks", data)
}

// MintArtworks is a paid mutator transaction binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] data) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) MintArtworks(data []FeralfileExhibitionV4MintData) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.MintArtworks(&_FeralfileExhibitionV45.TransactOpts, data)
}

// MintArtworks is a paid mutator transaction binding the contract method 0x8cba1c67.
//
// Solidity: function mintArtworks((uint256,uint256,address)[] data) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) MintArtworks(data []FeralfileExhibitionV4MintData) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.MintArtworks(&_FeralfileExhibitionV45.TransactOpts, data)
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) PauseSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "pauseSale")
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) PauseSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.PauseSale(&_FeralfileExhibitionV45.TransactOpts)
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) PauseSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.PauseSale(&_FeralfileExhibitionV45.TransactOpts)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) RemoveTrustee(opts *bind.TransactOpts, _trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "removeTrustee", _trustee)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) RemoveTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.RemoveTrustee(&_FeralfileExhibitionV45.TransactOpts, _trustee)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) RemoveTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.RemoveTrustee(&_FeralfileExhibitionV45.TransactOpts, _trustee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) RenounceOwnership() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.RenounceOwnership(&_FeralfileExhibitionV45.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.RenounceOwnership(&_FeralfileExhibitionV45.TransactOpts)
}

// ReplaceAdvanceAddresses is a paid mutator transaction binding the contract method 0x41a5626a.
//
// Solidity: function replaceAdvanceAddresses(address[] oldAddresses_, address[] newAddresses_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) ReplaceAdvanceAddresses(opts *bind.TransactOpts, oldAddresses_ []common.Address, newAddresses_ []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "replaceAdvanceAddresses", oldAddresses_, newAddresses_)
}

// ReplaceAdvanceAddresses is a paid mutator transaction binding the contract method 0x41a5626a.
//
// Solidity: function replaceAdvanceAddresses(address[] oldAddresses_, address[] newAddresses_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) ReplaceAdvanceAddresses(oldAddresses_ []common.Address, newAddresses_ []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.ReplaceAdvanceAddresses(&_FeralfileExhibitionV45.TransactOpts, oldAddresses_, newAddresses_)
}

// ReplaceAdvanceAddresses is a paid mutator transaction binding the contract method 0x41a5626a.
//
// Solidity: function replaceAdvanceAddresses(address[] oldAddresses_, address[] newAddresses_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) ReplaceAdvanceAddresses(oldAddresses_ []common.Address, newAddresses_ []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.ReplaceAdvanceAddresses(&_FeralfileExhibitionV45.TransactOpts, oldAddresses_, newAddresses_)
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) ResumeSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "resumeSale")
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) ResumeSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.ResumeSale(&_FeralfileExhibitionV45.TransactOpts)
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) ResumeSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.ResumeSale(&_FeralfileExhibitionV45.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SafeTransferFrom(&_FeralfileExhibitionV45.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SafeTransferFrom(&_FeralfileExhibitionV45.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SafeTransferFrom0(&_FeralfileExhibitionV45.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SafeTransferFrom0(&_FeralfileExhibitionV45.TransactOpts, from, to, tokenId, data)
}

// SetAdvanceSetting is a paid mutator transaction binding the contract method 0x3c352b0d.
//
// Solidity: function setAdvanceSetting(address[] addresses_, uint256[] amounts_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) SetAdvanceSetting(opts *bind.TransactOpts, addresses_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "setAdvanceSetting", addresses_, amounts_)
}

// SetAdvanceSetting is a paid mutator transaction binding the contract method 0x3c352b0d.
//
// Solidity: function setAdvanceSetting(address[] addresses_, uint256[] amounts_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SetAdvanceSetting(addresses_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetAdvanceSetting(&_FeralfileExhibitionV45.TransactOpts, addresses_, amounts_)
}

// SetAdvanceSetting is a paid mutator transaction binding the contract method 0x3c352b0d.
//
// Solidity: function setAdvanceSetting(address[] addresses_, uint256[] amounts_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) SetAdvanceSetting(addresses_ []common.Address, amounts_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetAdvanceSetting(&_FeralfileExhibitionV45.TransactOpts, addresses_, amounts_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetApprovalForAll(&_FeralfileExhibitionV45.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetApprovalForAll(&_FeralfileExhibitionV45.TransactOpts, operator, approved)
}

// SetCostReceiver is a paid mutator transaction binding the contract method 0x1623528f.
//
// Solidity: function setCostReceiver(address costReceiver_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) SetCostReceiver(opts *bind.TransactOpts, costReceiver_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "setCostReceiver", costReceiver_)
}

// SetCostReceiver is a paid mutator transaction binding the contract method 0x1623528f.
//
// Solidity: function setCostReceiver(address costReceiver_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SetCostReceiver(costReceiver_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetCostReceiver(&_FeralfileExhibitionV45.TransactOpts, costReceiver_)
}

// SetCostReceiver is a paid mutator transaction binding the contract method 0x1623528f.
//
// Solidity: function setCostReceiver(address costReceiver_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) SetCostReceiver(costReceiver_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetCostReceiver(&_FeralfileExhibitionV45.TransactOpts, costReceiver_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) SetSigner(opts *bind.TransactOpts, signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "setSigner", signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetSigner(&_FeralfileExhibitionV45.TransactOpts, signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetSigner(&_FeralfileExhibitionV45.TransactOpts, signer_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) SetTokenBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "setTokenBaseURI", baseURI_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SetTokenBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetTokenBaseURI(&_FeralfileExhibitionV45.TransactOpts, baseURI_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) SetTokenBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetTokenBaseURI(&_FeralfileExhibitionV45.TransactOpts, baseURI_)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vault_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) SetVault(opts *bind.TransactOpts, vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "setVault", vault_)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vault_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) SetVault(vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetVault(&_FeralfileExhibitionV45.TransactOpts, vault_)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vault_) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) SetVault(vault_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.SetVault(&_FeralfileExhibitionV45.TransactOpts, vault_)
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) StartSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "startSale")
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) StartSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.StartSale(&_FeralfileExhibitionV45.TransactOpts)
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) StartSale() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.StartSale(&_FeralfileExhibitionV45.TransactOpts)
}

// StopSaleAndBurn is a paid mutator transaction binding the contract method 0xb9b8311a.
//
// Solidity: function stopSaleAndBurn() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) StopSaleAndBurn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "stopSaleAndBurn")
}

// StopSaleAndBurn is a paid mutator transaction binding the contract method 0xb9b8311a.
//
// Solidity: function stopSaleAndBurn() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) StopSaleAndBurn() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.StopSaleAndBurn(&_FeralfileExhibitionV45.TransactOpts)
}

// StopSaleAndBurn is a paid mutator transaction binding the contract method 0xb9b8311a.
//
// Solidity: function stopSaleAndBurn() returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) StopSaleAndBurn() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.StopSaleAndBurn(&_FeralfileExhibitionV45.TransactOpts)
}

// StopSaleAndTransfer is a paid mutator transaction binding the contract method 0x65a46e08.
//
// Solidity: function stopSaleAndTransfer(uint256[] seriesIds, address[] recipientAddresses) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) StopSaleAndTransfer(opts *bind.TransactOpts, seriesIds []*big.Int, recipientAddresses []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "stopSaleAndTransfer", seriesIds, recipientAddresses)
}

// StopSaleAndTransfer is a paid mutator transaction binding the contract method 0x65a46e08.
//
// Solidity: function stopSaleAndTransfer(uint256[] seriesIds, address[] recipientAddresses) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) StopSaleAndTransfer(seriesIds []*big.Int, recipientAddresses []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.StopSaleAndTransfer(&_FeralfileExhibitionV45.TransactOpts, seriesIds, recipientAddresses)
}

// StopSaleAndTransfer is a paid mutator transaction binding the contract method 0x65a46e08.
//
// Solidity: function stopSaleAndTransfer(uint256[] seriesIds, address[] recipientAddresses) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) StopSaleAndTransfer(seriesIds []*big.Int, recipientAddresses []common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.StopSaleAndTransfer(&_FeralfileExhibitionV45.TransactOpts, seriesIds, recipientAddresses)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.TransferFrom(&_FeralfileExhibitionV45.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.TransferFrom(&_FeralfileExhibitionV45.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.TransferOwnership(&_FeralfileExhibitionV45.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.TransferOwnership(&_FeralfileExhibitionV45.TransactOpts, newOwner)
}

// UpdateOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x114ba8ee.
//
// Solidity: function updateOperatorFilterRegistry(address operatorFilterRegisterAddress) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) UpdateOperatorFilterRegistry(opts *bind.TransactOpts, operatorFilterRegisterAddress common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.Transact(opts, "updateOperatorFilterRegistry", operatorFilterRegisterAddress)
}

// UpdateOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x114ba8ee.
//
// Solidity: function updateOperatorFilterRegistry(address operatorFilterRegisterAddress) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) UpdateOperatorFilterRegistry(operatorFilterRegisterAddress common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.UpdateOperatorFilterRegistry(&_FeralfileExhibitionV45.TransactOpts, operatorFilterRegisterAddress)
}

// UpdateOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x114ba8ee.
//
// Solidity: function updateOperatorFilterRegistry(address operatorFilterRegisterAddress) returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) UpdateOperatorFilterRegistry(operatorFilterRegisterAddress common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.UpdateOperatorFilterRegistry(&_FeralfileExhibitionV45.TransactOpts, operatorFilterRegisterAddress)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV45.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Session) Receive() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.Receive(&_FeralfileExhibitionV45.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45TransactorSession) Receive() (*types.Transaction, error) {
	return _FeralfileExhibitionV45.Contract.Receive(&_FeralfileExhibitionV45.TransactOpts)
}

// FeralfileExhibitionV45ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45ApprovalIterator struct {
	Event *FeralfileExhibitionV45Approval // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV45ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV45Approval)
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
		it.Event = new(FeralfileExhibitionV45Approval)
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
func (it *FeralfileExhibitionV45ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV45ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV45Approval represents a Approval event raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV45ApprovalIterator, error) {

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

	logs, sub, err := _FeralfileExhibitionV45.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45ApprovalIterator{contract: _FeralfileExhibitionV45.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV45Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _FeralfileExhibitionV45.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV45Approval)
				if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) ParseApproval(log types.Log) (*FeralfileExhibitionV45Approval, error) {
	event := new(FeralfileExhibitionV45Approval)
	if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV45ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45ApprovalForAllIterator struct {
	Event *FeralfileExhibitionV45ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV45ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV45ApprovalForAll)
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
		it.Event = new(FeralfileExhibitionV45ApprovalForAll)
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
func (it *FeralfileExhibitionV45ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV45ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV45ApprovalForAll represents a ApprovalForAll event raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*FeralfileExhibitionV45ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FeralfileExhibitionV45.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45ApprovalForAllIterator{contract: _FeralfileExhibitionV45.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV45ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FeralfileExhibitionV45.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV45ApprovalForAll)
				if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) ParseApprovalForAll(log types.Log) (*FeralfileExhibitionV45ApprovalForAll, error) {
	event := new(FeralfileExhibitionV45ApprovalForAll)
	if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV45BurnArtworkIterator is returned from FilterBurnArtwork and is used to iterate over the raw logs and unpacked data for BurnArtwork events raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45BurnArtworkIterator struct {
	Event *FeralfileExhibitionV45BurnArtwork // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV45BurnArtworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV45BurnArtwork)
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
		it.Event = new(FeralfileExhibitionV45BurnArtwork)
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
func (it *FeralfileExhibitionV45BurnArtworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV45BurnArtworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV45BurnArtwork represents a BurnArtwork event raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45BurnArtwork struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnArtwork is a free log retrieval operation binding the contract event 0xbde7938970372996ff103863625e348ef2bf8f38a5b02181be75aafef17c23d5.
//
// Solidity: event BurnArtwork(uint256 indexed tokenId)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) FilterBurnArtwork(opts *bind.FilterOpts, tokenId []*big.Int) (*FeralfileExhibitionV45BurnArtworkIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV45.contract.FilterLogs(opts, "BurnArtwork", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45BurnArtworkIterator{contract: _FeralfileExhibitionV45.contract, event: "BurnArtwork", logs: logs, sub: sub}, nil
}

// WatchBurnArtwork is a free log subscription operation binding the contract event 0xbde7938970372996ff103863625e348ef2bf8f38a5b02181be75aafef17c23d5.
//
// Solidity: event BurnArtwork(uint256 indexed tokenId)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) WatchBurnArtwork(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV45BurnArtwork, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV45.contract.WatchLogs(opts, "BurnArtwork", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV45BurnArtwork)
				if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "BurnArtwork", log); err != nil {
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
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) ParseBurnArtwork(log types.Log) (*FeralfileExhibitionV45BurnArtwork, error) {
	event := new(FeralfileExhibitionV45BurnArtwork)
	if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "BurnArtwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV45BuyArtworkIterator is returned from FilterBuyArtwork and is used to iterate over the raw logs and unpacked data for BuyArtwork events raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45BuyArtworkIterator struct {
	Event *FeralfileExhibitionV45BuyArtwork // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV45BuyArtworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV45BuyArtwork)
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
		it.Event = new(FeralfileExhibitionV45BuyArtwork)
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
func (it *FeralfileExhibitionV45BuyArtworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV45BuyArtworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV45BuyArtwork represents a BuyArtwork event raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45BuyArtwork struct {
	Buyer   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyArtwork is a free log retrieval operation binding the contract event 0x0475389cd69b8d3163620b43283bf74e8fc71020c3c6cef2a529b5c405e9687f.
//
// Solidity: event BuyArtwork(address indexed buyer, uint256 indexed tokenId)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) FilterBuyArtwork(opts *bind.FilterOpts, buyer []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV45BuyArtworkIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV45.contract.FilterLogs(opts, "BuyArtwork", buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45BuyArtworkIterator{contract: _FeralfileExhibitionV45.contract, event: "BuyArtwork", logs: logs, sub: sub}, nil
}

// WatchBuyArtwork is a free log subscription operation binding the contract event 0x0475389cd69b8d3163620b43283bf74e8fc71020c3c6cef2a529b5c405e9687f.
//
// Solidity: event BuyArtwork(address indexed buyer, uint256 indexed tokenId)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) WatchBuyArtwork(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV45BuyArtwork, buyer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV45.contract.WatchLogs(opts, "BuyArtwork", buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV45BuyArtwork)
				if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "BuyArtwork", log); err != nil {
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
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) ParseBuyArtwork(log types.Log) (*FeralfileExhibitionV45BuyArtwork, error) {
	event := new(FeralfileExhibitionV45BuyArtwork)
	if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "BuyArtwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV45NewArtworkIterator is returned from FilterNewArtwork and is used to iterate over the raw logs and unpacked data for NewArtwork events raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45NewArtworkIterator struct {
	Event *FeralfileExhibitionV45NewArtwork // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV45NewArtworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV45NewArtwork)
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
		it.Event = new(FeralfileExhibitionV45NewArtwork)
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
func (it *FeralfileExhibitionV45NewArtworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV45NewArtworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV45NewArtwork represents a NewArtwork event raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45NewArtwork struct {
	Owner    common.Address
	SeriesId *big.Int
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewArtwork is a free log retrieval operation binding the contract event 0x407d7da1d3b2b1871fbfa2b5b1c4657a3cc5711d3023c552798551c7ee301eea.
//
// Solidity: event NewArtwork(address indexed owner, uint256 indexed seriesId, uint256 indexed tokenId)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) FilterNewArtwork(opts *bind.FilterOpts, owner []common.Address, seriesId []*big.Int, tokenId []*big.Int) (*FeralfileExhibitionV45NewArtworkIterator, error) {

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

	logs, sub, err := _FeralfileExhibitionV45.contract.FilterLogs(opts, "NewArtwork", ownerRule, seriesIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45NewArtworkIterator{contract: _FeralfileExhibitionV45.contract, event: "NewArtwork", logs: logs, sub: sub}, nil
}

// WatchNewArtwork is a free log subscription operation binding the contract event 0x407d7da1d3b2b1871fbfa2b5b1c4657a3cc5711d3023c552798551c7ee301eea.
//
// Solidity: event NewArtwork(address indexed owner, uint256 indexed seriesId, uint256 indexed tokenId)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) WatchNewArtwork(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV45NewArtwork, owner []common.Address, seriesId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _FeralfileExhibitionV45.contract.WatchLogs(opts, "NewArtwork", ownerRule, seriesIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV45NewArtwork)
				if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "NewArtwork", log); err != nil {
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
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) ParseNewArtwork(log types.Log) (*FeralfileExhibitionV45NewArtwork, error) {
	event := new(FeralfileExhibitionV45NewArtwork)
	if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "NewArtwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV45OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45OwnershipTransferredIterator struct {
	Event *FeralfileExhibitionV45OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV45OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV45OwnershipTransferred)
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
		it.Event = new(FeralfileExhibitionV45OwnershipTransferred)
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
func (it *FeralfileExhibitionV45OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV45OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV45OwnershipTransferred represents a OwnershipTransferred event raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeralfileExhibitionV45OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralfileExhibitionV45.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45OwnershipTransferredIterator{contract: _FeralfileExhibitionV45.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV45OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralfileExhibitionV45.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV45OwnershipTransferred)
				if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) ParseOwnershipTransferred(log types.Log) (*FeralfileExhibitionV45OwnershipTransferred, error) {
	event := new(FeralfileExhibitionV45OwnershipTransferred)
	if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV45TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45TransferIterator struct {
	Event *FeralfileExhibitionV45Transfer // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV45TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV45Transfer)
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
		it.Event = new(FeralfileExhibitionV45Transfer)
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
func (it *FeralfileExhibitionV45TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV45TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV45Transfer represents a Transfer event raised by the FeralfileExhibitionV45 contract.
type FeralfileExhibitionV45Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV45TransferIterator, error) {

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

	logs, sub, err := _FeralfileExhibitionV45.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV45TransferIterator{contract: _FeralfileExhibitionV45.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV45Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _FeralfileExhibitionV45.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV45Transfer)
				if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_FeralfileExhibitionV45 *FeralfileExhibitionV45Filterer) ParseTransfer(log types.Log) (*FeralfileExhibitionV45Transfer, error) {
	event := new(FeralfileExhibitionV45Transfer)
	if err := _FeralfileExhibitionV45.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
