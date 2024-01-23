// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package english_auction

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

// FeralfileEnglishAuctionAuction is an auto generated low-level Go binding around an user-defined struct.
type FeralfileEnglishAuctionAuction struct {
	Id                *big.Int
	StartAt           *big.Int
	EndAt             *big.Int
	ExtendDuration    *big.Int
	ExtendThreshold   *big.Int
	MinIncreaseFactor *big.Int
	MinIncreaseAmount *big.Int
	MinPrice          *big.Int
	IsSettled         bool
}

// FeralfileEnglishAuctionAuctionStatus is an auto generated low-level Go binding around an user-defined struct.
type FeralfileEnglishAuctionAuctionStatus struct {
	HighestBid FeralfileEnglishAuctionBid
	EndAt      *big.Int
	IsSettled  bool
}

// FeralfileEnglishAuctionBid is an auto generated low-level Go binding around an user-defined struct.
type FeralfileEnglishAuctionBid struct {
	Bidder        common.Address
	Amount        *big.Int
	FromFeralFile bool
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

// FeralfileEnglishAuctionMetaData contains all meta data concerning the FeralfileEnglishAuction contract.
var FeralfileEnglishAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fromFeralFile\",\"type\":\"bool\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extendDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extendThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncreaseFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncreaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSettled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBids\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fromFeralFile\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"ongoingAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extendDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extendThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncreaseFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncreaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSettled\",\"type\":\"bool\"}],\"internalType\":\"structFeralfileEnglishAuction.Auction[]\",\"name\":\"auctions_\",\"type\":\"tuple[]\"}],\"name\":\"registerAuctions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"auctionIDs_\",\"type\":\"uint256[]\"}],\"name\":\"listAuctionStatus\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fromFeralFile\",\"type\":\"bool\"}],\"internalType\":\"structFeralfileEnglishAuction.Bid\",\"name\":\"highestBid\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSettled\",\"type\":\"bool\"}],\"internalType\":\"structFeralfileEnglishAuction.AuctionStatus[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"isValidNewBid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionID_\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionID_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTime_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"}],\"name\":\"placeSignedBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionID_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultAddr_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"internalType\":\"structIFeralfileSaleData.RevenueShare[][]\",\"name\":\"revenueShares\",\"type\":\"tuple[][]\"},{\"internalType\":\"bool\",\"name\":\"payByVaultContract\",\"type\":\"bool\"}],\"internalType\":\"structIFeralfileSaleData.SaleData\",\"name\":\"saleData_\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"}],\"name\":\"settleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionID_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vaultAddr_\",\"type\":\"address\"}],\"name\":\"settleAuctionFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620025123803806200251283398101604081905262000034916200011d565b806200004033620000cd565b6001600160a01b038116620000a65760405162461bcd60e51b815260206004820152602260248201527f45434453415369676e3a207369676e65725f206973207a65726f206164647265604482015261737360f01b606482015260840160405180910390fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055506200014f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156200013057600080fd5b81516001600160a01b03811681146200014857600080fd5b9392505050565b6123b3806200015f6000396000f3fe6080604052600436106100e85760003560e01c80636c19e7831161008a5780639979ef45116100595780639979ef4514610355578063e5de8f5d14610368578063f2fde38b14610388578063f92a6d58146103a857600080fd5b80636c19e783146102d55780636d3ae68a146102f5578063715018a6146103225780638da5cb5b1461033757600080fd5b806321fbf1e6116100c657806321fbf1e6146101b4578063238ac933146101d45780633b8b8bb414610206578063571a26a01461022657600080fd5b806304cf7ec4146100ed5780630d33bcee14610122578063111cb25814610144575b600080fd5b3480156100f957600080fd5b5061010d610108366004611abc565b6103c8565b60405190151581526020015b60405180910390f35b34801561012e57600080fd5b5061014261013d366004611d3d565b610434565b005b34801561015057600080fd5b5061018d61015f366004611abc565b6003602052600090815260409020805460018201546002909201546001600160a01b03909116919060ff1683565b604080516001600160a01b0390941684526020840192909252151590820152606001610119565b3480156101c057600080fd5b506101426101cf366004611e5f565b6106e3565b3480156101e057600080fd5b506001546001600160a01b03165b6040516001600160a01b039091168152602001610119565b34801561021257600080fd5b5061010d610221366004611e8b565b6107b0565b34801561023257600080fd5b5061028f610241366004611abc565b60026020819052600091825260409091208054600182015492820154600383015460048401546005850154600686015460078701546008909701549597969495939492939192909160ff1689565b60408051998a5260208a0198909852968801959095526060870193909352608086019190915260a085015260c084015260e0830152151561010082015261012001610119565b3480156102e157600080fd5b506101426102f0366004611ead565b610875565b34801561030157600080fd5b50610315610310366004611ecf565b610900565b6040516101199190611f0c565b34801561032e57600080fd5b50610142610b24565b34801561034357600080fd5b506000546001600160a01b03166101ee565b610142610363366004611abc565b610b2e565b34801561037457600080fd5b50610142610383366004611f85565b610b3e565b34801561039457600080fd5b506101426103a3366004611ead565b610c79565b3480156103b457600080fd5b506101426103c3366004611fda565b610cef565b6000818152600260205260408120546103fc5760405162461bcd60e51b81526004016103f390612050565b60405180910390fd5b600082815260026020526040902060010154421080159061042e57506000828152600260208190526040909120015442105b92915050565b61043c611205565b8360c001516104bf5760405162461bcd60e51b815260206004820152604360248201527f466572616c66696c65456e676c69736841756374696f6e3a2073616c6544617460448201527f612e70617942795661756c74436f6e74726163742073686f756c64206265207460648201526272756560e81b608482015260a4016103f3565b60008781526003602090815260409182902082516060808201855282546001600160a01b0390811680845260018501549584019590955260029093015460ff1615159482019490945292870151161461059c5760405162461bcd60e51b815260206004820152605360248201527f466572616c66696c65456e676c69736841756374696f6e3a2073616c6544617460448201527f615f2e64657374696e6174696f6e20697320646966666572656e742066726f6d606482015272103434b3b432b9ba103134b2103134b23232b960691b608482015260a4016103f3565b6000888152600260208181526040928390208351610120810185528154815260018201549281019290925291820154928101929092526003810154606083015260048101546080830152600581015460a0830152600681015460c0830152600781015460e08301526008015460ff16151561010082015261061e81838961125f565b604051632977e4b360e01b81526001600160a01b03891690632977e4b390610650908890889088908c90600401612118565b600060405180830381600087803b15801561066a57600080fd5b505af115801561067e573d6000803e3d6000fd5b5050505081600001516001600160a01b0316886001600160a01b03168a7f3e25c3675d003af5184b628dd8bd4775b9b3abc7351c3574c90870ca25b55f1185602001516040516106d091815260200190565b60405180910390a4505050505050505050565b6106eb611205565b60008281526002602081815260408084208151610120810183528154815260018083015482860152828601548285015260038084015460608085019190915260048501546080850152600585015460a0850152600685015460c0850152600785015460e085015260089094015460ff90811615156101008501528a8952908652968490208451938401855280546001600160a01b03168452908101549483019490945292909301549093161515928201929092526107aa82828561125f565b50505050565b60008281526002602081815260408084208151610120810183528154815260018083015482860152828601548285015260038084015460608085019190915260048501546080850152600585015460a0850152600685015460c0850152600785015460e085015260089094015460ff90811615156101008501528a89529086528488208551948501865280546001600160a01b0316855291820154958401959095529094015490921615159082015261086a8282866113c1565b506001949350505050565b61087d611205565b6001600160a01b0381166108de5760405162461bcd60e51b815260206004820152602260248201527f45434453415369676e3a207369676e65725f206973207a65726f206164647265604482015261737360f01b60648201526084016103f3565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b60606000825167ffffffffffffffff81111561091e5761091e611af1565b60405190808252806020026020018201604052801561097d57816020015b6040805160c0810182526000606082018181526080830182905260a083018290528252602080830182905292820152825260001990920191018161093c5790505b50905060005b8351811015610b1d576000600260008684815181106109a4576109a46121e1565b602002602001015181526020019081526020016000206040518061012001604052908160008201548152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820160009054906101000a900460ff161515151581525050905060036000868481518110610a4957610a496121e1565b6020908102919091018101518252818101929092526040908101600020815160608101835281546001600160a01b031681526001820154938101939093526002015460ff161515908201528351849084908110610aa857610aa86121e1565b6020026020010151600001819052508060400151838381518110610ace57610ace6121e1565b60200260200101516020018181525050806101000151838381518110610af657610af66121e1565b60209081029190910101519015156040909101525080610b158161220d565b915050610983565b5092915050565b610b2c611205565b565b610b3b81333460006115d2565b50565b428411610ba35760405162461bcd60e51b815260206004820152602d60248201527f466572616c66696c65456e676c69736841756374696f6e3a207369676e61747560448201526c1c99481a5cc8195e1c1a5c9959609a1b60648201526084016103f3565b604080514660208201523091810191909152606081018890526001600160a01b038716608082015260a0810186905260c0810185905260009060e001604051602081830303815290604052805190602001209050610c03818585856117de565b610c625760405162461bcd60e51b815260206004820152602a60248201527f466572616c66696c65456e676c69736841756374696f6e3a20696e76616c6964604482015269207369676e617475726560b01b60648201526084016103f3565b610c6f88888860016115d2565b5050505050505050565b610c81611205565b6001600160a01b038116610ce65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103f3565b610b3b81611836565b610cf7611205565b60005b8181101561120057600060026000858585818110610d1a57610d1a6121e1565b610120908102929092013583525060208083019390935260409182016000208251918201835280548083526001820154948301949094526002810154928201929092526003820154606082015260048201546080820152600582015460a0820152600682015460c0820152600782015460e082015260089091015460ff161515610100820152915015610df45760405162461bcd60e51b815260206004820152602e602482015260008051602061235e83398151915260448201526d08185b1c9958591e48195e1a5cdd60921b60648201526084016103f3565b6000848484818110610e0857610e086121e1565b905061012002016000013511610e745760405162461bcd60e51b815260206004820152602b60248201527f466572616c66696c65456e676c69736841756374696f6e3a20696e76616c696460448201526a08185d58dd1a5bdb881a5960aa1b60648201526084016103f3565b42848484818110610e8757610e876121e1565b905061012002016020013510158015610ebb575042848484818110610eae57610eae6121e1565b9050610120020160400135115b610f345760405162461bcd60e51b8152602060048201526050602482015260008051602061235e83398151915260448201527f2073746172742074696d6520616e6420656e642074696d652073686f756c642060648201526f626520696e207468652066757475726560801b608482015260a4016103f3565b838383818110610f4657610f466121e1565b9050610120020160200135848484818110610f6357610f636121e1565b905061012002016040013511610fdd5760405162461bcd60e51b81526020600482015260446024820181905260008051602061235e833981519152908201527f20656e642074696d652073686f756c642062652061667465722073746172742060648201526374696d6560e01b608482015260a4016103f3565b6000848484818110610ff157610ff16121e1565b9050610120020160a00135116110735760405162461bcd60e51b815260206004820152604d602482015260008051602061235e83398151915260448201527f206d696e20696e63726561736520666163746f722073686f756c64206265206760648201526c0726561746572207468616e203609c1b608482015260a4016103f3565b6000848484818110611087576110876121e1565b9050610120020160c00135116111095760405162461bcd60e51b815260206004820152604d602482015260008051602061235e83398151915260448201527f206d696e20696e63726561736520616d6f756e742073686f756c64206265206760648201526c0726561746572207468616e203609c1b608482015260a4016103f3565b600084848481811061111d5761111d6121e1565b9050610120020160e00135116111955760405162461bcd60e51b8152602060048201526043602482015260008051602061235e83398151915260448201527f206d696e2070726963652073686f756c6420626520677265617465722074686160648201526206e20360ec1b608482015260a4016103f3565b8383838181106111a7576111a76121e1565b90506101200201600260008686868181106111c4576111c46121e1565b9050610120020160000135815260200190815260200160002081816111e99190612233565b9050505080806111f89061220d565b915050610cfa565b505050565b6000546001600160a01b03163314610b2c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103f3565b82516000036112805760405162461bcd60e51b81526004016103f390612050565b42836040015111156112d55760405162461bcd60e51b815260206004820152602a602482015260008051602061235e833981519152604482015269081b9bdd08195b99195960b21b60648201526084016103f3565b826101000151156112f85760405162461bcd60e51b81526004016103f3906122aa565b81516001600160a01b031661134f5760405162461bcd60e51b815260206004820152601f60248201527f466572616c66696c65456e676c69736841756374696f6e3a206e6f206269640060448201526064016103f3565b825160009081526002602052604090819020600801805460ff19166001179055820151158015611383575060008260200151115b156112005760208201516040516001600160a01b0383169180156108fc02916000818181858888f193505050501580156107aa573d6000803e3d6000fd5b82516113df5760405162461bcd60e51b81526004016103f390612050565b428360200151111580156113f65750826040015142105b61144e5760405162461bcd60e51b8152602060048201526035602482015260008051602061235e833981519152604482015274081b9bdd081cdd185c9d1959081bdc88195b991959605a1b60648201526084016103f3565b826101000151156114715760405162461bcd60e51b81526004016103f3906122aa565b60208201511561154757600060648460a00151846020015161149391906122e8565b61149d91906122ff565b90508360c001518110156114b2575060c08301515b8083602001516114c29190612321565b8210156107aa5760405162461bcd60e51b815260206004820152604760248201527f466572616c66696c65456e676c69736841756374696f6e3a2062696420616d6f60448201527f756e742073686f756c6420666f6c6c6f7720746865206d696e696d756d20696e60648201526618dc995b595b9d60ca1b608482015260a4016103f3565b8260e001518110156112005760405162461bcd60e51b815260206004820152604860248201527f466572616c66696c65456e676c69736841756374696f6e3a2062696420616d6f60448201527f756e742073686f756c642062652067726561746572207468616e206d696e696d606482015267756d20707269636560c01b608482015260a4016103f3565b60008481526002602081815260408084208151610120810183528154815260018083015482860152828601548285015260038084015460608085019190915260048501546080850152600585015460a0850152600685015460c0850152600785015460e085015260089094015460ff90811615156101008501528c8952908652968490208451938401855280546001600160a01b03168452908101549483019490945292909301549093161515928201929092526116918282866113c1565b604080516060810182526001600160a01b038781168252602080830188815287151584860190815260008c8152600390935291859020935184546001600160a01b031916931692909217835590516001830155516002909101805460ff191691151591909117905560808301519083015161170d904290612334565b116117355760608201516117219042612321565b600087815260026020819052604090912001555b806040015115801561174b575060008160200151115b1561178f57805160208201516040516001600160a01b039092169181156108fc0291906000818181858888f1935050505015801561178d573d6000803e3d6000fd5b505b83856001600160a01b0316877ff25203491f62af68d98d27c61e683af917ad41cba414886cc71b4976b363d8d2866040516117ce911515815260200190565b60405180910390a4505050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c859052603c8120819061181b90848787611886565b6001546001600160a01b039081169116149695505050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000806000611897878787876118ae565b915091506118a481611972565b5095945050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156118e55750600090506003611969565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611939573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661196257600060019250925050611969565b9150600090505b94509492505050565b600081600481111561198657611986612347565b0361198e5750565b60018160048111156119a2576119a2612347565b036119ef5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016103f3565b6002816004811115611a0357611a03612347565b03611a505760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016103f3565b6003816004811115611a6457611a64612347565b03610b3b5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016103f3565b600060208284031215611ace57600080fd5b5035919050565b80356001600160a01b0381168114611aec57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611b2a57611b2a611af1565b60405290565b60405160e0810167ffffffffffffffff81118282101715611b2a57611b2a611af1565b604051601f8201601f1916810167ffffffffffffffff81118282101715611b7c57611b7c611af1565b604052919050565b600067ffffffffffffffff821115611b9e57611b9e611af1565b5060051b60200190565b600082601f830112611bb957600080fd5b81356020611bce611bc983611b84565b611b53565b82815260059290921b84018101918181019086841115611bed57600080fd5b8286015b84811015611c085780358352918301918301611bf1565b509695505050505050565b600082601f830112611c2457600080fd5b81356020611c34611bc983611b84565b82815260059290921b84018101918181019086841115611c5357600080fd5b8286015b84811015611c0857803567ffffffffffffffff811115611c775760008081fd5b8701603f81018913611c895760008081fd5b848101356040611c9b611bc983611b84565b82815260069290921b8301810191878101908c841115611cbb5760008081fd5b938201935b83851015611d025782858e031215611cd85760008081fd5b611ce0611b07565b611ce986611ad5565b8152858a01358a82015282529382019390880190611cc0565b875250505092840192508301611c57565b8015158114610b3b57600080fd5b8035611aec81611d13565b803560ff81168114611aec57600080fd5b600080600080600080600060e0888a031215611d5857600080fd5b87359650611d6860208901611ad5565b9550611d7660408901611ad5565b9450606088013567ffffffffffffffff80821115611d9357600080fd5b9089019060e0828c031215611da757600080fd5b611daf611b30565b823581526020830135602082015260408301356040820152611dd360608401611ad5565b6060820152608083013582811115611dea57600080fd5b611df68d828601611ba8565b60808301525060a083013582811115611e0e57600080fd5b611e1a8d828601611c13565b60a083015250611e2c60c08401611d21565b60c08201528096505050506080880135925060a08801359150611e5160c08901611d2c565b905092959891949750929550565b60008060408385031215611e7257600080fd5b82359150611e8260208401611ad5565b90509250929050565b60008060408385031215611e9e57600080fd5b50508035926020909101359150565b600060208284031215611ebf57600080fd5b611ec882611ad5565b9392505050565b600060208284031215611ee157600080fd5b813567ffffffffffffffff811115611ef857600080fd5b611f0484828501611ba8565b949350505050565b602080825282518282018190526000919060409081850190868401855b82811015611f78578151805180516001600160a01b0316865287810151888701528601511515868601528681015160608601528501511515608085015260a09093019290850190600101611f29565b5091979650505050505050565b600080600080600080600060e0888a031215611fa057600080fd5b87359650611fb060208901611ad5565b955060408801359450606088013593506080880135925060a08801359150611e5160c08901611d2c565b60008060208385031215611fed57600080fd5b823567ffffffffffffffff8082111561200557600080fd5b818501915085601f83011261201957600080fd5b81358181111561202857600080fd5b8660206101208302850101111561203e57600080fd5b60209290920196919550909350505050565b6020808252602a9082015260008051602061235e833981519152604082015269081b9bdd08199bdd5b9960b21b606082015260800190565b600081518084526020808501808196508360051b810191508286016000805b8681101561210a578385038a52825180518087529087019087870190845b818110156120f557835180516001600160a01b031684528a01518a840152928901926040909201916001016120c5565b50509a87019a955050918501916001016120a7565b509298975050505050505050565b84815260006020858184015260ff85166040840152608060608401526101608301845160808501528185015160a0850152604085015160c085015260018060a01b0360608601511660e0850152608085015160e0610100860152818151808452610180870191508483019350600092505b808310156121a95783518252928401926001929092019190840190612189565b5060a0870151868203607f190161012088015293506121c88185612088565b935050505060c0840151611c0861014085018215159052565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161221f5761221f6121f7565b5060010190565b6000813561042e81611d13565b813581556020820135600182015560408201356002820155606082013560038201556080820135600482015560a0820135600582015560c0820135600682015560e082013560078201556122a661228d6101008401612226565b6008830160ff1981541660ff8315151681178255505050565b5050565b602080825260309082015260008051602061235e83398151915260408201526f08185b1c9958591e481cd95d1d1b195960821b606082015260800190565b808202811582820484141761042e5761042e6121f7565b60008261231c57634e487b7160e01b600052601260045260246000fd5b500490565b8082018082111561042e5761042e6121f7565b8181038181111561042e5761042e6121f7565b634e487b7160e01b600052602160045260246000fdfe466572616c66696c65456e676c69736841756374696f6e3a2061756374696f6ea2646970667358221220ca68dbd7be8b1035b41989c000ce5689da0acc9fecf6f67b6a03d7ff2184bf4b64736f6c63430008110033",
}

// FeralfileEnglishAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use FeralfileEnglishAuctionMetaData.ABI instead.
var FeralfileEnglishAuctionABI = FeralfileEnglishAuctionMetaData.ABI

// FeralfileEnglishAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeralfileEnglishAuctionMetaData.Bin instead.
var FeralfileEnglishAuctionBin = FeralfileEnglishAuctionMetaData.Bin

// DeployFeralfileEnglishAuction deploys a new Ethereum contract, binding an instance of FeralfileEnglishAuction to it.
func DeployFeralfileEnglishAuction(auth *bind.TransactOpts, backend bind.ContractBackend, signer_ common.Address) (common.Address, *types.Transaction, *FeralfileEnglishAuction, error) {
	parsed, err := FeralfileEnglishAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeralfileEnglishAuctionBin), backend, signer_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeralfileEnglishAuction{FeralfileEnglishAuctionCaller: FeralfileEnglishAuctionCaller{contract: contract}, FeralfileEnglishAuctionTransactor: FeralfileEnglishAuctionTransactor{contract: contract}, FeralfileEnglishAuctionFilterer: FeralfileEnglishAuctionFilterer{contract: contract}}, nil
}

// FeralfileEnglishAuction is an auto generated Go binding around an Ethereum contract.
type FeralfileEnglishAuction struct {
	FeralfileEnglishAuctionCaller     // Read-only binding to the contract
	FeralfileEnglishAuctionTransactor // Write-only binding to the contract
	FeralfileEnglishAuctionFilterer   // Log filterer for contract events
}

// FeralfileEnglishAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeralfileEnglishAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileEnglishAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeralfileEnglishAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileEnglishAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeralfileEnglishAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileEnglishAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeralfileEnglishAuctionSession struct {
	Contract     *FeralfileEnglishAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FeralfileEnglishAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeralfileEnglishAuctionCallerSession struct {
	Contract *FeralfileEnglishAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// FeralfileEnglishAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeralfileEnglishAuctionTransactorSession struct {
	Contract     *FeralfileEnglishAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// FeralfileEnglishAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeralfileEnglishAuctionRaw struct {
	Contract *FeralfileEnglishAuction // Generic contract binding to access the raw methods on
}

// FeralfileEnglishAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeralfileEnglishAuctionCallerRaw struct {
	Contract *FeralfileEnglishAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// FeralfileEnglishAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeralfileEnglishAuctionTransactorRaw struct {
	Contract *FeralfileEnglishAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeralfileEnglishAuction creates a new instance of FeralfileEnglishAuction, bound to a specific deployed contract.
func NewFeralfileEnglishAuction(address common.Address, backend bind.ContractBackend) (*FeralfileEnglishAuction, error) {
	contract, err := bindFeralfileEnglishAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeralfileEnglishAuction{FeralfileEnglishAuctionCaller: FeralfileEnglishAuctionCaller{contract: contract}, FeralfileEnglishAuctionTransactor: FeralfileEnglishAuctionTransactor{contract: contract}, FeralfileEnglishAuctionFilterer: FeralfileEnglishAuctionFilterer{contract: contract}}, nil
}

// NewFeralfileEnglishAuctionCaller creates a new read-only instance of FeralfileEnglishAuction, bound to a specific deployed contract.
func NewFeralfileEnglishAuctionCaller(address common.Address, caller bind.ContractCaller) (*FeralfileEnglishAuctionCaller, error) {
	contract, err := bindFeralfileEnglishAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeralfileEnglishAuctionCaller{contract: contract}, nil
}

// NewFeralfileEnglishAuctionTransactor creates a new write-only instance of FeralfileEnglishAuction, bound to a specific deployed contract.
func NewFeralfileEnglishAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*FeralfileEnglishAuctionTransactor, error) {
	contract, err := bindFeralfileEnglishAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeralfileEnglishAuctionTransactor{contract: contract}, nil
}

// NewFeralfileEnglishAuctionFilterer creates a new log filterer instance of FeralfileEnglishAuction, bound to a specific deployed contract.
func NewFeralfileEnglishAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*FeralfileEnglishAuctionFilterer, error) {
	contract, err := bindFeralfileEnglishAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeralfileEnglishAuctionFilterer{contract: contract}, nil
}

// bindFeralfileEnglishAuction binds a generic wrapper to an already deployed contract.
func bindFeralfileEnglishAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeralfileEnglishAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralfileEnglishAuction.Contract.FeralfileEnglishAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.FeralfileEnglishAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.FeralfileEnglishAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralfileEnglishAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.contract.Transact(opts, method, params...)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint256 id, uint256 startAt, uint256 endAt, uint256 extendDuration, uint256 extendThreshold, uint256 minIncreaseFactor, uint256 minIncreaseAmount, uint256 minPrice, bool isSettled)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                *big.Int
	StartAt           *big.Int
	EndAt             *big.Int
	ExtendDuration    *big.Int
	ExtendThreshold   *big.Int
	MinIncreaseFactor *big.Int
	MinIncreaseAmount *big.Int
	MinPrice          *big.Int
	IsSettled         bool
}, error) {
	var out []interface{}
	err := _FeralfileEnglishAuction.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		Id                *big.Int
		StartAt           *big.Int
		EndAt             *big.Int
		ExtendDuration    *big.Int
		ExtendThreshold   *big.Int
		MinIncreaseFactor *big.Int
		MinIncreaseAmount *big.Int
		MinPrice          *big.Int
		IsSettled         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ExtendDuration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExtendThreshold = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MinIncreaseFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MinIncreaseAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.MinPrice = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.IsSettled = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint256 id, uint256 startAt, uint256 endAt, uint256 extendDuration, uint256 extendThreshold, uint256 minIncreaseFactor, uint256 minIncreaseAmount, uint256 minPrice, bool isSettled)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) Auctions(arg0 *big.Int) (struct {
	Id                *big.Int
	StartAt           *big.Int
	EndAt             *big.Int
	ExtendDuration    *big.Int
	ExtendThreshold   *big.Int
	MinIncreaseFactor *big.Int
	MinIncreaseAmount *big.Int
	MinPrice          *big.Int
	IsSettled         bool
}, error) {
	return _FeralfileEnglishAuction.Contract.Auctions(&_FeralfileEnglishAuction.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint256 id, uint256 startAt, uint256 endAt, uint256 extendDuration, uint256 extendThreshold, uint256 minIncreaseFactor, uint256 minIncreaseAmount, uint256 minPrice, bool isSettled)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCallerSession) Auctions(arg0 *big.Int) (struct {
	Id                *big.Int
	StartAt           *big.Int
	EndAt             *big.Int
	ExtendDuration    *big.Int
	ExtendThreshold   *big.Int
	MinIncreaseFactor *big.Int
	MinIncreaseAmount *big.Int
	MinPrice          *big.Int
	IsSettled         bool
}, error) {
	return _FeralfileEnglishAuction.Contract.Auctions(&_FeralfileEnglishAuction.CallOpts, arg0)
}

// HighestBids is a free data retrieval call binding the contract method 0x111cb258.
//
// Solidity: function highestBids(uint256 ) view returns(address bidder, uint256 amount, bool fromFeralFile)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCaller) HighestBids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bidder        common.Address
	Amount        *big.Int
	FromFeralFile bool
}, error) {
	var out []interface{}
	err := _FeralfileEnglishAuction.contract.Call(opts, &out, "highestBids", arg0)

	outstruct := new(struct {
		Bidder        common.Address
		Amount        *big.Int
		FromFeralFile bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bidder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FromFeralFile = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// HighestBids is a free data retrieval call binding the contract method 0x111cb258.
//
// Solidity: function highestBids(uint256 ) view returns(address bidder, uint256 amount, bool fromFeralFile)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) HighestBids(arg0 *big.Int) (struct {
	Bidder        common.Address
	Amount        *big.Int
	FromFeralFile bool
}, error) {
	return _FeralfileEnglishAuction.Contract.HighestBids(&_FeralfileEnglishAuction.CallOpts, arg0)
}

// HighestBids is a free data retrieval call binding the contract method 0x111cb258.
//
// Solidity: function highestBids(uint256 ) view returns(address bidder, uint256 amount, bool fromFeralFile)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCallerSession) HighestBids(arg0 *big.Int) (struct {
	Bidder        common.Address
	Amount        *big.Int
	FromFeralFile bool
}, error) {
	return _FeralfileEnglishAuction.Contract.HighestBids(&_FeralfileEnglishAuction.CallOpts, arg0)
}

// IsValidNewBid is a free data retrieval call binding the contract method 0x3b8b8bb4.
//
// Solidity: function isValidNewBid(uint256 auctionID_, uint256 amount_) view returns(bool)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCaller) IsValidNewBid(opts *bind.CallOpts, auctionID_ *big.Int, amount_ *big.Int) (bool, error) {
	var out []interface{}
	err := _FeralfileEnglishAuction.contract.Call(opts, &out, "isValidNewBid", auctionID_, amount_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidNewBid is a free data retrieval call binding the contract method 0x3b8b8bb4.
//
// Solidity: function isValidNewBid(uint256 auctionID_, uint256 amount_) view returns(bool)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) IsValidNewBid(auctionID_ *big.Int, amount_ *big.Int) (bool, error) {
	return _FeralfileEnglishAuction.Contract.IsValidNewBid(&_FeralfileEnglishAuction.CallOpts, auctionID_, amount_)
}

// IsValidNewBid is a free data retrieval call binding the contract method 0x3b8b8bb4.
//
// Solidity: function isValidNewBid(uint256 auctionID_, uint256 amount_) view returns(bool)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCallerSession) IsValidNewBid(auctionID_ *big.Int, amount_ *big.Int) (bool, error) {
	return _FeralfileEnglishAuction.Contract.IsValidNewBid(&_FeralfileEnglishAuction.CallOpts, auctionID_, amount_)
}

// ListAuctionStatus is a free data retrieval call binding the contract method 0x6d3ae68a.
//
// Solidity: function listAuctionStatus(uint256[] auctionIDs_) view returns(((address,uint256,bool),uint256,bool)[])
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCaller) ListAuctionStatus(opts *bind.CallOpts, auctionIDs_ []*big.Int) ([]FeralfileEnglishAuctionAuctionStatus, error) {
	var out []interface{}
	err := _FeralfileEnglishAuction.contract.Call(opts, &out, "listAuctionStatus", auctionIDs_)

	if err != nil {
		return *new([]FeralfileEnglishAuctionAuctionStatus), err
	}

	out0 := *abi.ConvertType(out[0], new([]FeralfileEnglishAuctionAuctionStatus)).(*[]FeralfileEnglishAuctionAuctionStatus)

	return out0, err

}

// ListAuctionStatus is a free data retrieval call binding the contract method 0x6d3ae68a.
//
// Solidity: function listAuctionStatus(uint256[] auctionIDs_) view returns(((address,uint256,bool),uint256,bool)[])
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) ListAuctionStatus(auctionIDs_ []*big.Int) ([]FeralfileEnglishAuctionAuctionStatus, error) {
	return _FeralfileEnglishAuction.Contract.ListAuctionStatus(&_FeralfileEnglishAuction.CallOpts, auctionIDs_)
}

// ListAuctionStatus is a free data retrieval call binding the contract method 0x6d3ae68a.
//
// Solidity: function listAuctionStatus(uint256[] auctionIDs_) view returns(((address,uint256,bool),uint256,bool)[])
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCallerSession) ListAuctionStatus(auctionIDs_ []*big.Int) ([]FeralfileEnglishAuctionAuctionStatus, error) {
	return _FeralfileEnglishAuction.Contract.ListAuctionStatus(&_FeralfileEnglishAuction.CallOpts, auctionIDs_)
}

// OngoingAuction is a free data retrieval call binding the contract method 0x04cf7ec4.
//
// Solidity: function ongoingAuction(uint256 id_) view returns(bool)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCaller) OngoingAuction(opts *bind.CallOpts, id_ *big.Int) (bool, error) {
	var out []interface{}
	err := _FeralfileEnglishAuction.contract.Call(opts, &out, "ongoingAuction", id_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OngoingAuction is a free data retrieval call binding the contract method 0x04cf7ec4.
//
// Solidity: function ongoingAuction(uint256 id_) view returns(bool)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) OngoingAuction(id_ *big.Int) (bool, error) {
	return _FeralfileEnglishAuction.Contract.OngoingAuction(&_FeralfileEnglishAuction.CallOpts, id_)
}

// OngoingAuction is a free data retrieval call binding the contract method 0x04cf7ec4.
//
// Solidity: function ongoingAuction(uint256 id_) view returns(bool)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCallerSession) OngoingAuction(id_ *big.Int) (bool, error) {
	return _FeralfileEnglishAuction.Contract.OngoingAuction(&_FeralfileEnglishAuction.CallOpts, id_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileEnglishAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) Owner() (common.Address, error) {
	return _FeralfileEnglishAuction.Contract.Owner(&_FeralfileEnglishAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCallerSession) Owner() (common.Address, error) {
	return _FeralfileEnglishAuction.Contract.Owner(&_FeralfileEnglishAuction.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileEnglishAuction.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) Signer() (common.Address, error) {
	return _FeralfileEnglishAuction.Contract.Signer(&_FeralfileEnglishAuction.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionCallerSession) Signer() (common.Address, error) {
	return _FeralfileEnglishAuction.Contract.Signer(&_FeralfileEnglishAuction.CallOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 auctionID_) payable returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactor) PlaceBid(opts *bind.TransactOpts, auctionID_ *big.Int) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.contract.Transact(opts, "placeBid", auctionID_)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 auctionID_) payable returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) PlaceBid(auctionID_ *big.Int) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.PlaceBid(&_FeralfileEnglishAuction.TransactOpts, auctionID_)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 auctionID_) payable returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactorSession) PlaceBid(auctionID_ *big.Int) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.PlaceBid(&_FeralfileEnglishAuction.TransactOpts, auctionID_)
}

// PlaceSignedBid is a paid mutator transaction binding the contract method 0xe5de8f5d.
//
// Solidity: function placeSignedBid(uint256 auctionID_, address bidder_, uint256 amount_, uint256 expiryTime_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactor) PlaceSignedBid(opts *bind.TransactOpts, auctionID_ *big.Int, bidder_ common.Address, amount_ *big.Int, expiryTime_ *big.Int, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.contract.Transact(opts, "placeSignedBid", auctionID_, bidder_, amount_, expiryTime_, r_, s_, v_)
}

// PlaceSignedBid is a paid mutator transaction binding the contract method 0xe5de8f5d.
//
// Solidity: function placeSignedBid(uint256 auctionID_, address bidder_, uint256 amount_, uint256 expiryTime_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) PlaceSignedBid(auctionID_ *big.Int, bidder_ common.Address, amount_ *big.Int, expiryTime_ *big.Int, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.PlaceSignedBid(&_FeralfileEnglishAuction.TransactOpts, auctionID_, bidder_, amount_, expiryTime_, r_, s_, v_)
}

// PlaceSignedBid is a paid mutator transaction binding the contract method 0xe5de8f5d.
//
// Solidity: function placeSignedBid(uint256 auctionID_, address bidder_, uint256 amount_, uint256 expiryTime_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactorSession) PlaceSignedBid(auctionID_ *big.Int, bidder_ common.Address, amount_ *big.Int, expiryTime_ *big.Int, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.PlaceSignedBid(&_FeralfileEnglishAuction.TransactOpts, auctionID_, bidder_, amount_, expiryTime_, r_, s_, v_)
}

// RegisterAuctions is a paid mutator transaction binding the contract method 0xf92a6d58.
//
// Solidity: function registerAuctions((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)[] auctions_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactor) RegisterAuctions(opts *bind.TransactOpts, auctions_ []FeralfileEnglishAuctionAuction) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.contract.Transact(opts, "registerAuctions", auctions_)
}

// RegisterAuctions is a paid mutator transaction binding the contract method 0xf92a6d58.
//
// Solidity: function registerAuctions((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)[] auctions_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) RegisterAuctions(auctions_ []FeralfileEnglishAuctionAuction) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.RegisterAuctions(&_FeralfileEnglishAuction.TransactOpts, auctions_)
}

// RegisterAuctions is a paid mutator transaction binding the contract method 0xf92a6d58.
//
// Solidity: function registerAuctions((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)[] auctions_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactorSession) RegisterAuctions(auctions_ []FeralfileEnglishAuctionAuction) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.RegisterAuctions(&_FeralfileEnglishAuction.TransactOpts, auctions_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.RenounceOwnership(&_FeralfileEnglishAuction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.RenounceOwnership(&_FeralfileEnglishAuction.TransactOpts)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactor) SetSigner(opts *bind.TransactOpts, signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.contract.Transact(opts, "setSigner", signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.SetSigner(&_FeralfileEnglishAuction.TransactOpts, signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactorSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.SetSigner(&_FeralfileEnglishAuction.TransactOpts, signer_)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x0d33bcee.
//
// Solidity: function settleAuction(uint256 auctionID_, address tokenAddr_, address vaultAddr_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactor) SettleAuction(opts *bind.TransactOpts, auctionID_ *big.Int, tokenAddr_ common.Address, vaultAddr_ common.Address, saleData_ IFeralfileSaleDataSaleData, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.contract.Transact(opts, "settleAuction", auctionID_, tokenAddr_, vaultAddr_, saleData_, r_, s_, v_)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x0d33bcee.
//
// Solidity: function settleAuction(uint256 auctionID_, address tokenAddr_, address vaultAddr_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) SettleAuction(auctionID_ *big.Int, tokenAddr_ common.Address, vaultAddr_ common.Address, saleData_ IFeralfileSaleDataSaleData, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.SettleAuction(&_FeralfileEnglishAuction.TransactOpts, auctionID_, tokenAddr_, vaultAddr_, saleData_, r_, s_, v_)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x0d33bcee.
//
// Solidity: function settleAuction(uint256 auctionID_, address tokenAddr_, address vaultAddr_, (uint256,uint256,uint256,address,uint256[],(address,uint256)[][],bool) saleData_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactorSession) SettleAuction(auctionID_ *big.Int, tokenAddr_ common.Address, vaultAddr_ common.Address, saleData_ IFeralfileSaleDataSaleData, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.SettleAuction(&_FeralfileEnglishAuction.TransactOpts, auctionID_, tokenAddr_, vaultAddr_, saleData_, r_, s_, v_)
}

// SettleAuctionFund is a paid mutator transaction binding the contract method 0x21fbf1e6.
//
// Solidity: function settleAuctionFund(uint256 auctionID_, address vaultAddr_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactor) SettleAuctionFund(opts *bind.TransactOpts, auctionID_ *big.Int, vaultAddr_ common.Address) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.contract.Transact(opts, "settleAuctionFund", auctionID_, vaultAddr_)
}

// SettleAuctionFund is a paid mutator transaction binding the contract method 0x21fbf1e6.
//
// Solidity: function settleAuctionFund(uint256 auctionID_, address vaultAddr_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) SettleAuctionFund(auctionID_ *big.Int, vaultAddr_ common.Address) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.SettleAuctionFund(&_FeralfileEnglishAuction.TransactOpts, auctionID_, vaultAddr_)
}

// SettleAuctionFund is a paid mutator transaction binding the contract method 0x21fbf1e6.
//
// Solidity: function settleAuctionFund(uint256 auctionID_, address vaultAddr_) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactorSession) SettleAuctionFund(auctionID_ *big.Int, vaultAddr_ common.Address) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.SettleAuctionFund(&_FeralfileEnglishAuction.TransactOpts, auctionID_, vaultAddr_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.TransferOwnership(&_FeralfileEnglishAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileEnglishAuction.Contract.TransferOwnership(&_FeralfileEnglishAuction.TransactOpts, newOwner)
}

// FeralfileEnglishAuctionAuctionSettledIterator is returned from FilterAuctionSettled and is used to iterate over the raw logs and unpacked data for AuctionSettled events raised by the FeralfileEnglishAuction contract.
type FeralfileEnglishAuctionAuctionSettledIterator struct {
	Event *FeralfileEnglishAuctionAuctionSettled // Event containing the contract specifics and raw log

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
func (it *FeralfileEnglishAuctionAuctionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileEnglishAuctionAuctionSettled)
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
		it.Event = new(FeralfileEnglishAuctionAuctionSettled)
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
func (it *FeralfileEnglishAuctionAuctionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileEnglishAuctionAuctionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileEnglishAuctionAuctionSettled represents a AuctionSettled event raised by the FeralfileEnglishAuction contract.
type FeralfileEnglishAuctionAuctionSettled struct {
	AuctionId       *big.Int
	ContractAddress common.Address
	Winner          common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettled is a free log retrieval operation binding the contract event 0x3e25c3675d003af5184b628dd8bd4775b9b3abc7351c3574c90870ca25b55f11.
//
// Solidity: event AuctionSettled(uint256 indexed auctionId, address indexed contractAddress, address indexed winner, uint256 amount)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionFilterer) FilterAuctionSettled(opts *bind.FilterOpts, auctionId []*big.Int, contractAddress []common.Address, winner []common.Address) (*FeralfileEnglishAuctionAuctionSettledIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _FeralfileEnglishAuction.contract.FilterLogs(opts, "AuctionSettled", auctionIdRule, contractAddressRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileEnglishAuctionAuctionSettledIterator{contract: _FeralfileEnglishAuction.contract, event: "AuctionSettled", logs: logs, sub: sub}, nil
}

// WatchAuctionSettled is a free log subscription operation binding the contract event 0x3e25c3675d003af5184b628dd8bd4775b9b3abc7351c3574c90870ca25b55f11.
//
// Solidity: event AuctionSettled(uint256 indexed auctionId, address indexed contractAddress, address indexed winner, uint256 amount)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionFilterer) WatchAuctionSettled(opts *bind.WatchOpts, sink chan<- *FeralfileEnglishAuctionAuctionSettled, auctionId []*big.Int, contractAddress []common.Address, winner []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _FeralfileEnglishAuction.contract.WatchLogs(opts, "AuctionSettled", auctionIdRule, contractAddressRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileEnglishAuctionAuctionSettled)
				if err := _FeralfileEnglishAuction.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
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

// ParseAuctionSettled is a log parse operation binding the contract event 0x3e25c3675d003af5184b628dd8bd4775b9b3abc7351c3574c90870ca25b55f11.
//
// Solidity: event AuctionSettled(uint256 indexed auctionId, address indexed contractAddress, address indexed winner, uint256 amount)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionFilterer) ParseAuctionSettled(log types.Log) (*FeralfileEnglishAuctionAuctionSettled, error) {
	event := new(FeralfileEnglishAuctionAuctionSettled)
	if err := _FeralfileEnglishAuction.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileEnglishAuctionNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the FeralfileEnglishAuction contract.
type FeralfileEnglishAuctionNewBidIterator struct {
	Event *FeralfileEnglishAuctionNewBid // Event containing the contract specifics and raw log

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
func (it *FeralfileEnglishAuctionNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileEnglishAuctionNewBid)
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
		it.Event = new(FeralfileEnglishAuctionNewBid)
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
func (it *FeralfileEnglishAuctionNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileEnglishAuctionNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileEnglishAuctionNewBid represents a NewBid event raised by the FeralfileEnglishAuction contract.
type FeralfileEnglishAuctionNewBid struct {
	AuctionId     *big.Int
	Bidder        common.Address
	Amount        *big.Int
	FromFeralFile bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xf25203491f62af68d98d27c61e683af917ad41cba414886cc71b4976b363d8d2.
//
// Solidity: event NewBid(uint256 indexed auctionId, address indexed bidder, uint256 indexed amount, bool fromFeralFile)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionFilterer) FilterNewBid(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address, amount []*big.Int) (*FeralfileEnglishAuctionNewBidIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FeralfileEnglishAuction.contract.FilterLogs(opts, "NewBid", auctionIdRule, bidderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileEnglishAuctionNewBidIterator{contract: _FeralfileEnglishAuction.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xf25203491f62af68d98d27c61e683af917ad41cba414886cc71b4976b363d8d2.
//
// Solidity: event NewBid(uint256 indexed auctionId, address indexed bidder, uint256 indexed amount, bool fromFeralFile)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *FeralfileEnglishAuctionNewBid, auctionId []*big.Int, bidder []common.Address, amount []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FeralfileEnglishAuction.contract.WatchLogs(opts, "NewBid", auctionIdRule, bidderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileEnglishAuctionNewBid)
				if err := _FeralfileEnglishAuction.contract.UnpackLog(event, "NewBid", log); err != nil {
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

// ParseNewBid is a log parse operation binding the contract event 0xf25203491f62af68d98d27c61e683af917ad41cba414886cc71b4976b363d8d2.
//
// Solidity: event NewBid(uint256 indexed auctionId, address indexed bidder, uint256 indexed amount, bool fromFeralFile)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionFilterer) ParseNewBid(log types.Log) (*FeralfileEnglishAuctionNewBid, error) {
	event := new(FeralfileEnglishAuctionNewBid)
	if err := _FeralfileEnglishAuction.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileEnglishAuctionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeralfileEnglishAuction contract.
type FeralfileEnglishAuctionOwnershipTransferredIterator struct {
	Event *FeralfileEnglishAuctionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeralfileEnglishAuctionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileEnglishAuctionOwnershipTransferred)
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
		it.Event = new(FeralfileEnglishAuctionOwnershipTransferred)
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
func (it *FeralfileEnglishAuctionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileEnglishAuctionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileEnglishAuctionOwnershipTransferred represents a OwnershipTransferred event raised by the FeralfileEnglishAuction contract.
type FeralfileEnglishAuctionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeralfileEnglishAuctionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralfileEnglishAuction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileEnglishAuctionOwnershipTransferredIterator{contract: _FeralfileEnglishAuction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeralfileEnglishAuctionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralfileEnglishAuction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileEnglishAuctionOwnershipTransferred)
				if err := _FeralfileEnglishAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeralfileEnglishAuction *FeralfileEnglishAuctionFilterer) ParseOwnershipTransferred(log types.Log) (*FeralfileEnglishAuctionOwnershipTransferred, error) {
	event := new(FeralfileEnglishAuctionOwnershipTransferred)
	if err := _FeralfileEnglishAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
