// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package miswap

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

// LibOrderAsset is an auto generated low-level Go binding around an user-defined struct.
type LibOrderAsset struct {
	TokenId    *big.Int
	Collection common.Address
	Amount     *big.Int
}

// LibOrderEditDetail is an auto generated low-level Go binding around an user-defined struct.
type LibOrderEditDetail struct {
	OldOrderKey [32]byte
	NewOrder    LibOrderOrder
}

// LibOrderMatchDetail is an auto generated low-level Go binding around an user-defined struct.
type LibOrderMatchDetail struct {
	SellOrder LibOrderOrder
	BuyOrder  LibOrderOrder
}

// LibOrderOrder is an auto generated low-level Go binding around an user-defined struct.
type LibOrderOrder struct {
	Side     uint8
	SaleKind uint8
	Maker    common.Address
	Nft      LibOrderAsset
	Price    *big.Int
	Expiry   uint64
	Salt     uint64
}

// MiSwapOrderBookMetaData contains all meta data concerning the MiSwapOrderBook contract.
var MiSwapOrderBookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"OrderKey\",\"name\":\"orderKey\",\"type\":\"bytes32\"}],\"name\":\"CannotInsertDuplicateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotInsertEmptyKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotInsertExistingKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRemoveEmptyKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRemoveMissingKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"msg\",\"type\":\"bytes\"}],\"name\":\"BatchMatchInnerError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"OrderKey\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"LogCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"OrderKey\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"name\":\"LogMake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"OrderKey\",\"name\":\"makeOrderKey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"OrderKey\",\"name\":\"takeOrderKey\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structLibOrder.Order\",\"name\":\"makeOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structLibOrder.Order\",\"name\":\"takeOrder\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"fillPrice\",\"type\":\"uint128\"}],\"name\":\"LogMatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"newShare\",\"type\":\"uint128\"}],\"name\":\"LogProtocolShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"OrderKey\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"name\":\"LogSkipOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"OrderKey[]\",\"name\":\"orderKeys\",\"type\":\"bytes32[]\"}],\"name\":\"cancelOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"canceled\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"OrderKey\",\"name\":\"oldOrderKey\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"newOrder\",\"type\":\"tuple\"}],\"internalType\":\"structLibOrder.EditDetail[]\",\"name\":\"editDetails\",\"type\":\"tuple[]\"}],\"name\":\"editOrders\",\"outputs\":[{\"internalType\":\"OrderKey[]\",\"name\":\"newOrderKeys\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"OrderKey\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filledAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newProtocolShare\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"EIP712Name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"EIP712Version\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"newOrders\",\"type\":\"tuple[]\"}],\"name\":\"makeOrders\",\"outputs\":[{\"internalType\":\"OrderKey[]\",\"name\":\"orderKeys\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"buyOrder\",\"type\":\"tuple\"}],\"name\":\"matchOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"matchOrderWithoutPayback\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"costValue\",\"type\":\"uint128\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"buyOrder\",\"type\":\"tuple\"}],\"internalType\":\"structLibOrder.MatchDetail[]\",\"name\":\"matchDetails\",\"type\":\"tuple[]\"}],\"name\":\"matchOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"matched\",\"type\":\"bool[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumLibOrder.Side\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"Price\",\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"orderQueues\",\"outputs\":[{\"internalType\":\"OrderKey\",\"name\":\"head\",\"type\":\"bytes32\"},{\"internalType\":\"OrderKey\",\"name\":\"tail\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"OrderKey\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orders\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"OrderKey\",\"name\":\"next\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumLibOrder.Side\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"priceTrees\",\"outputs\":[{\"internalType\":\"Price\",\"name\":\"root\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolShare\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_protocolShare\",\"type\":\"uint128\"}],\"name\":\"setProtocolShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MiSwapOrderBookABI is the input ABI used to generate the binding from.
// Deprecated: Use MiSwapOrderBookMetaData.ABI instead.
var MiSwapOrderBookABI = MiSwapOrderBookMetaData.ABI

// MiSwapOrderBook is an auto generated Go binding around an Ethereum contract.
type MiSwapOrderBook struct {
	MiSwapOrderBookCaller     // Read-only binding to the contract
	MiSwapOrderBookTransactor // Write-only binding to the contract
	MiSwapOrderBookFilterer   // Log filterer for contract events
}

// MiSwapOrderBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiSwapOrderBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiSwapOrderBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiSwapOrderBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiSwapOrderBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MiSwapOrderBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiSwapOrderBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiSwapOrderBookSession struct {
	Contract     *MiSwapOrderBook  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiSwapOrderBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiSwapOrderBookCallerSession struct {
	Contract *MiSwapOrderBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MiSwapOrderBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiSwapOrderBookTransactorSession struct {
	Contract     *MiSwapOrderBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MiSwapOrderBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type MiSwapOrderBookRaw struct {
	Contract *MiSwapOrderBook // Generic contract binding to access the raw methods on
}

// MiSwapOrderBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiSwapOrderBookCallerRaw struct {
	Contract *MiSwapOrderBookCaller // Generic read-only contract binding to access the raw methods on
}

// MiSwapOrderBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiSwapOrderBookTransactorRaw struct {
	Contract *MiSwapOrderBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiSwapOrderBook creates a new instance of MiSwapOrderBook, bound to a specific deployed contract.
func NewMiSwapOrderBook(address common.Address, backend bind.ContractBackend) (*MiSwapOrderBook, error) {
	contract, err := bindMiSwapOrderBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBook{MiSwapOrderBookCaller: MiSwapOrderBookCaller{contract: contract}, MiSwapOrderBookTransactor: MiSwapOrderBookTransactor{contract: contract}, MiSwapOrderBookFilterer: MiSwapOrderBookFilterer{contract: contract}}, nil
}

// NewMiSwapOrderBookCaller creates a new read-only instance of MiSwapOrderBook, bound to a specific deployed contract.
func NewMiSwapOrderBookCaller(address common.Address, caller bind.ContractCaller) (*MiSwapOrderBookCaller, error) {
	contract, err := bindMiSwapOrderBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookCaller{contract: contract}, nil
}

// NewMiSwapOrderBookTransactor creates a new write-only instance of MiSwapOrderBook, bound to a specific deployed contract.
func NewMiSwapOrderBookTransactor(address common.Address, transactor bind.ContractTransactor) (*MiSwapOrderBookTransactor, error) {
	contract, err := bindMiSwapOrderBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookTransactor{contract: contract}, nil
}

// NewMiSwapOrderBookFilterer creates a new log filterer instance of MiSwapOrderBook, bound to a specific deployed contract.
func NewMiSwapOrderBookFilterer(address common.Address, filterer bind.ContractFilterer) (*MiSwapOrderBookFilterer, error) {
	contract, err := bindMiSwapOrderBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookFilterer{contract: contract}, nil
}

// bindMiSwapOrderBook binds a generic wrapper to an already deployed contract.
func bindMiSwapOrderBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MiSwapOrderBookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiSwapOrderBook *MiSwapOrderBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MiSwapOrderBook.Contract.MiSwapOrderBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiSwapOrderBook *MiSwapOrderBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.MiSwapOrderBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiSwapOrderBook *MiSwapOrderBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.MiSwapOrderBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiSwapOrderBook *MiSwapOrderBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MiSwapOrderBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiSwapOrderBook *MiSwapOrderBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiSwapOrderBook *MiSwapOrderBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MiSwapOrderBook *MiSwapOrderBookCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _MiSwapOrderBook.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MiSwapOrderBook *MiSwapOrderBookSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MiSwapOrderBook.Contract.Eip712Domain(&_MiSwapOrderBook.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MiSwapOrderBook *MiSwapOrderBookCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MiSwapOrderBook.Contract.Eip712Domain(&_MiSwapOrderBook.CallOpts)
}

// FilledAmount is a free data retrieval call binding the contract method 0x0b20b7bc.
//
// Solidity: function filledAmount(bytes32 ) view returns(uint256)
func (_MiSwapOrderBook *MiSwapOrderBookCaller) FilledAmount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MiSwapOrderBook.contract.Call(opts, &out, "filledAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FilledAmount is a free data retrieval call binding the contract method 0x0b20b7bc.
//
// Solidity: function filledAmount(bytes32 ) view returns(uint256)
func (_MiSwapOrderBook *MiSwapOrderBookSession) FilledAmount(arg0 [32]byte) (*big.Int, error) {
	return _MiSwapOrderBook.Contract.FilledAmount(&_MiSwapOrderBook.CallOpts, arg0)
}

// FilledAmount is a free data retrieval call binding the contract method 0x0b20b7bc.
//
// Solidity: function filledAmount(bytes32 ) view returns(uint256)
func (_MiSwapOrderBook *MiSwapOrderBookCallerSession) FilledAmount(arg0 [32]byte) (*big.Int, error) {
	return _MiSwapOrderBook.Contract.FilledAmount(&_MiSwapOrderBook.CallOpts, arg0)
}

// OrderQueues is a free data retrieval call binding the contract method 0xe471ba46.
//
// Solidity: function orderQueues(address , uint8 , uint128 ) view returns(bytes32 head, bytes32 tail)
func (_MiSwapOrderBook *MiSwapOrderBookCaller) OrderQueues(opts *bind.CallOpts, arg0 common.Address, arg1 uint8, arg2 *big.Int) (struct {
	Head [32]byte
	Tail [32]byte
}, error) {
	var out []interface{}
	err := _MiSwapOrderBook.contract.Call(opts, &out, "orderQueues", arg0, arg1, arg2)

	outstruct := new(struct {
		Head [32]byte
		Tail [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Head = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Tail = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// OrderQueues is a free data retrieval call binding the contract method 0xe471ba46.
//
// Solidity: function orderQueues(address , uint8 , uint128 ) view returns(bytes32 head, bytes32 tail)
func (_MiSwapOrderBook *MiSwapOrderBookSession) OrderQueues(arg0 common.Address, arg1 uint8, arg2 *big.Int) (struct {
	Head [32]byte
	Tail [32]byte
}, error) {
	return _MiSwapOrderBook.Contract.OrderQueues(&_MiSwapOrderBook.CallOpts, arg0, arg1, arg2)
}

// OrderQueues is a free data retrieval call binding the contract method 0xe471ba46.
//
// Solidity: function orderQueues(address , uint8 , uint128 ) view returns(bytes32 head, bytes32 tail)
func (_MiSwapOrderBook *MiSwapOrderBookCallerSession) OrderQueues(arg0 common.Address, arg1 uint8, arg2 *big.Int) (struct {
	Head [32]byte
	Tail [32]byte
}, error) {
	return _MiSwapOrderBook.Contract.OrderQueues(&_MiSwapOrderBook.CallOpts, arg0, arg1, arg2)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) order, bytes32 next)
func (_MiSwapOrderBook *MiSwapOrderBookCaller) Orders(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Order LibOrderOrder
	Next  [32]byte
}, error) {
	var out []interface{}
	err := _MiSwapOrderBook.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Order LibOrderOrder
		Next  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Order = *abi.ConvertType(out[0], new(LibOrderOrder)).(*LibOrderOrder)
	outstruct.Next = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) order, bytes32 next)
func (_MiSwapOrderBook *MiSwapOrderBookSession) Orders(arg0 [32]byte) (struct {
	Order LibOrderOrder
	Next  [32]byte
}, error) {
	return _MiSwapOrderBook.Contract.Orders(&_MiSwapOrderBook.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) order, bytes32 next)
func (_MiSwapOrderBook *MiSwapOrderBookCallerSession) Orders(arg0 [32]byte) (struct {
	Order LibOrderOrder
	Next  [32]byte
}, error) {
	return _MiSwapOrderBook.Contract.Orders(&_MiSwapOrderBook.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MiSwapOrderBook *MiSwapOrderBookCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MiSwapOrderBook.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MiSwapOrderBook *MiSwapOrderBookSession) Owner() (common.Address, error) {
	return _MiSwapOrderBook.Contract.Owner(&_MiSwapOrderBook.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MiSwapOrderBook *MiSwapOrderBookCallerSession) Owner() (common.Address, error) {
	return _MiSwapOrderBook.Contract.Owner(&_MiSwapOrderBook.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MiSwapOrderBook *MiSwapOrderBookCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MiSwapOrderBook.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MiSwapOrderBook *MiSwapOrderBookSession) Paused() (bool, error) {
	return _MiSwapOrderBook.Contract.Paused(&_MiSwapOrderBook.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MiSwapOrderBook *MiSwapOrderBookCallerSession) Paused() (bool, error) {
	return _MiSwapOrderBook.Contract.Paused(&_MiSwapOrderBook.CallOpts)
}

// PriceTrees is a free data retrieval call binding the contract method 0x7d910dac.
//
// Solidity: function priceTrees(address , uint8 ) view returns(uint128 root)
func (_MiSwapOrderBook *MiSwapOrderBookCaller) PriceTrees(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (*big.Int, error) {
	var out []interface{}
	err := _MiSwapOrderBook.contract.Call(opts, &out, "priceTrees", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceTrees is a free data retrieval call binding the contract method 0x7d910dac.
//
// Solidity: function priceTrees(address , uint8 ) view returns(uint128 root)
func (_MiSwapOrderBook *MiSwapOrderBookSession) PriceTrees(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _MiSwapOrderBook.Contract.PriceTrees(&_MiSwapOrderBook.CallOpts, arg0, arg1)
}

// PriceTrees is a free data retrieval call binding the contract method 0x7d910dac.
//
// Solidity: function priceTrees(address , uint8 ) view returns(uint128 root)
func (_MiSwapOrderBook *MiSwapOrderBookCallerSession) PriceTrees(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _MiSwapOrderBook.Contract.PriceTrees(&_MiSwapOrderBook.CallOpts, arg0, arg1)
}

// ProtocolShare is a free data retrieval call binding the contract method 0x1103f315.
//
// Solidity: function protocolShare() view returns(uint128)
func (_MiSwapOrderBook *MiSwapOrderBookCaller) ProtocolShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MiSwapOrderBook.contract.Call(opts, &out, "protocolShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolShare is a free data retrieval call binding the contract method 0x1103f315.
//
// Solidity: function protocolShare() view returns(uint128)
func (_MiSwapOrderBook *MiSwapOrderBookSession) ProtocolShare() (*big.Int, error) {
	return _MiSwapOrderBook.Contract.ProtocolShare(&_MiSwapOrderBook.CallOpts)
}

// ProtocolShare is a free data retrieval call binding the contract method 0x1103f315.
//
// Solidity: function protocolShare() view returns(uint128)
func (_MiSwapOrderBook *MiSwapOrderBookCallerSession) ProtocolShare() (*big.Int, error) {
	return _MiSwapOrderBook.Contract.ProtocolShare(&_MiSwapOrderBook.CallOpts)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x21c77c96.
//
// Solidity: function cancelOrders(bytes32[] orderKeys) returns(bool[] canceled)
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) CancelOrders(opts *bind.TransactOpts, orderKeys [][32]byte) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "cancelOrders", orderKeys)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x21c77c96.
//
// Solidity: function cancelOrders(bytes32[] orderKeys) returns(bool[] canceled)
func (_MiSwapOrderBook *MiSwapOrderBookSession) CancelOrders(orderKeys [][32]byte) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.CancelOrders(&_MiSwapOrderBook.TransactOpts, orderKeys)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x21c77c96.
//
// Solidity: function cancelOrders(bytes32[] orderKeys) returns(bool[] canceled)
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) CancelOrders(orderKeys [][32]byte) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.CancelOrders(&_MiSwapOrderBook.TransactOpts, orderKeys)
}

// EditOrders is a paid mutator transaction binding the contract method 0xe8f27e95.
//
// Solidity: function editOrders((bytes32,(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] editDetails) payable returns(bytes32[] newOrderKeys)
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) EditOrders(opts *bind.TransactOpts, editDetails []LibOrderEditDetail) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "editOrders", editDetails)
}

// EditOrders is a paid mutator transaction binding the contract method 0xe8f27e95.
//
// Solidity: function editOrders((bytes32,(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] editDetails) payable returns(bytes32[] newOrderKeys)
func (_MiSwapOrderBook *MiSwapOrderBookSession) EditOrders(editDetails []LibOrderEditDetail) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.EditOrders(&_MiSwapOrderBook.TransactOpts, editDetails)
}

// EditOrders is a paid mutator transaction binding the contract method 0xe8f27e95.
//
// Solidity: function editOrders((bytes32,(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] editDetails) payable returns(bytes32[] newOrderKeys)
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) EditOrders(editDetails []LibOrderEditDetail) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.EditOrders(&_MiSwapOrderBook.TransactOpts, editDetails)
}

// Initialize is a paid mutator transaction binding the contract method 0x5d52957e.
//
// Solidity: function initialize(uint128 newProtocolShare, address newVault, string EIP712Name, string EIP712Version) returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) Initialize(opts *bind.TransactOpts, newProtocolShare *big.Int, newVault common.Address, EIP712Name string, EIP712Version string) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "initialize", newProtocolShare, newVault, EIP712Name, EIP712Version)
}

// Initialize is a paid mutator transaction binding the contract method 0x5d52957e.
//
// Solidity: function initialize(uint128 newProtocolShare, address newVault, string EIP712Name, string EIP712Version) returns()
func (_MiSwapOrderBook *MiSwapOrderBookSession) Initialize(newProtocolShare *big.Int, newVault common.Address, EIP712Name string, EIP712Version string) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.Initialize(&_MiSwapOrderBook.TransactOpts, newProtocolShare, newVault, EIP712Name, EIP712Version)
}

// Initialize is a paid mutator transaction binding the contract method 0x5d52957e.
//
// Solidity: function initialize(uint128 newProtocolShare, address newVault, string EIP712Name, string EIP712Version) returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) Initialize(newProtocolShare *big.Int, newVault common.Address, EIP712Name string, EIP712Version string) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.Initialize(&_MiSwapOrderBook.TransactOpts, newProtocolShare, newVault, EIP712Name, EIP712Version)
}

// MakeOrders is a paid mutator transaction binding the contract method 0x307725d9.
//
// Solidity: function makeOrders((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64)[] newOrders) payable returns(bytes32[] orderKeys)
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) MakeOrders(opts *bind.TransactOpts, newOrders []LibOrderOrder) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "makeOrders", newOrders)
}

// MakeOrders is a paid mutator transaction binding the contract method 0x307725d9.
//
// Solidity: function makeOrders((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64)[] newOrders) payable returns(bytes32[] orderKeys)
func (_MiSwapOrderBook *MiSwapOrderBookSession) MakeOrders(newOrders []LibOrderOrder) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.MakeOrders(&_MiSwapOrderBook.TransactOpts, newOrders)
}

// MakeOrders is a paid mutator transaction binding the contract method 0x307725d9.
//
// Solidity: function makeOrders((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64)[] newOrders) payable returns(bytes32[] orderKeys)
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) MakeOrders(newOrders []LibOrderOrder) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.MakeOrders(&_MiSwapOrderBook.TransactOpts, newOrders)
}

// MatchOrder is a paid mutator transaction binding the contract method 0x882849c9.
//
// Solidity: function matchOrder((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder) payable returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) MatchOrder(opts *bind.TransactOpts, sellOrder LibOrderOrder, buyOrder LibOrderOrder) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "matchOrder", sellOrder, buyOrder)
}

// MatchOrder is a paid mutator transaction binding the contract method 0x882849c9.
//
// Solidity: function matchOrder((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder) payable returns()
func (_MiSwapOrderBook *MiSwapOrderBookSession) MatchOrder(sellOrder LibOrderOrder, buyOrder LibOrderOrder) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.MatchOrder(&_MiSwapOrderBook.TransactOpts, sellOrder, buyOrder)
}

// MatchOrder is a paid mutator transaction binding the contract method 0x882849c9.
//
// Solidity: function matchOrder((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder) payable returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) MatchOrder(sellOrder LibOrderOrder, buyOrder LibOrderOrder) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.MatchOrder(&_MiSwapOrderBook.TransactOpts, sellOrder, buyOrder)
}

// MatchOrderWithoutPayback is a paid mutator transaction binding the contract method 0x366c39ed.
//
// Solidity: function matchOrderWithoutPayback((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder, uint256 msgValue) payable returns(uint128 costValue)
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) MatchOrderWithoutPayback(opts *bind.TransactOpts, sellOrder LibOrderOrder, buyOrder LibOrderOrder, msgValue *big.Int) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "matchOrderWithoutPayback", sellOrder, buyOrder, msgValue)
}

// MatchOrderWithoutPayback is a paid mutator transaction binding the contract method 0x366c39ed.
//
// Solidity: function matchOrderWithoutPayback((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder, uint256 msgValue) payable returns(uint128 costValue)
func (_MiSwapOrderBook *MiSwapOrderBookSession) MatchOrderWithoutPayback(sellOrder LibOrderOrder, buyOrder LibOrderOrder, msgValue *big.Int) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.MatchOrderWithoutPayback(&_MiSwapOrderBook.TransactOpts, sellOrder, buyOrder, msgValue)
}

// MatchOrderWithoutPayback is a paid mutator transaction binding the contract method 0x366c39ed.
//
// Solidity: function matchOrderWithoutPayback((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder, uint256 msgValue) payable returns(uint128 costValue)
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) MatchOrderWithoutPayback(sellOrder LibOrderOrder, buyOrder LibOrderOrder, msgValue *big.Int) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.MatchOrderWithoutPayback(&_MiSwapOrderBook.TransactOpts, sellOrder, buyOrder, msgValue)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xfe971c98.
//
// Solidity: function matchOrders(((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64),(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] matchDetails) payable returns(bool[] matched)
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) MatchOrders(opts *bind.TransactOpts, matchDetails []LibOrderMatchDetail) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "matchOrders", matchDetails)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xfe971c98.
//
// Solidity: function matchOrders(((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64),(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] matchDetails) payable returns(bool[] matched)
func (_MiSwapOrderBook *MiSwapOrderBookSession) MatchOrders(matchDetails []LibOrderMatchDetail) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.MatchOrders(&_MiSwapOrderBook.TransactOpts, matchDetails)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xfe971c98.
//
// Solidity: function matchOrders(((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64),(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] matchDetails) payable returns(bool[] matched)
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) MatchOrders(matchDetails []LibOrderMatchDetail) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.MatchOrders(&_MiSwapOrderBook.TransactOpts, matchDetails)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MiSwapOrderBook *MiSwapOrderBookSession) Pause() (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.Pause(&_MiSwapOrderBook.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) Pause() (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.Pause(&_MiSwapOrderBook.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MiSwapOrderBook *MiSwapOrderBookSession) RenounceOwnership() (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.RenounceOwnership(&_MiSwapOrderBook.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.RenounceOwnership(&_MiSwapOrderBook.TransactOpts)
}

// SetProtocolShare is a paid mutator transaction binding the contract method 0x813e7c2a.
//
// Solidity: function setProtocolShare(uint128 _protocolShare) returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) SetProtocolShare(opts *bind.TransactOpts, _protocolShare *big.Int) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "setProtocolShare", _protocolShare)
}

// SetProtocolShare is a paid mutator transaction binding the contract method 0x813e7c2a.
//
// Solidity: function setProtocolShare(uint128 _protocolShare) returns()
func (_MiSwapOrderBook *MiSwapOrderBookSession) SetProtocolShare(_protocolShare *big.Int) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.SetProtocolShare(&_MiSwapOrderBook.TransactOpts, _protocolShare)
}

// SetProtocolShare is a paid mutator transaction binding the contract method 0x813e7c2a.
//
// Solidity: function setProtocolShare(uint128 _protocolShare) returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) SetProtocolShare(_protocolShare *big.Int) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.SetProtocolShare(&_MiSwapOrderBook.TransactOpts, _protocolShare)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address newVault) returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) SetVault(opts *bind.TransactOpts, newVault common.Address) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "setVault", newVault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address newVault) returns()
func (_MiSwapOrderBook *MiSwapOrderBookSession) SetVault(newVault common.Address) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.SetVault(&_MiSwapOrderBook.TransactOpts, newVault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address newVault) returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) SetVault(newVault common.Address) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.SetVault(&_MiSwapOrderBook.TransactOpts, newVault)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MiSwapOrderBook *MiSwapOrderBookSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.TransferOwnership(&_MiSwapOrderBook.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.TransferOwnership(&_MiSwapOrderBook.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MiSwapOrderBook *MiSwapOrderBookSession) Unpause() (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.Unpause(&_MiSwapOrderBook.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) Unpause() (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.Unpause(&_MiSwapOrderBook.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiSwapOrderBook.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MiSwapOrderBook *MiSwapOrderBookSession) Receive() (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.Receive(&_MiSwapOrderBook.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MiSwapOrderBook *MiSwapOrderBookTransactorSession) Receive() (*types.Transaction, error) {
	return _MiSwapOrderBook.Contract.Receive(&_MiSwapOrderBook.TransactOpts)
}

// MiSwapOrderBookBatchMatchInnerErrorIterator is returned from FilterBatchMatchInnerError and is used to iterate over the raw logs and unpacked data for BatchMatchInnerError events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookBatchMatchInnerErrorIterator struct {
	Event *MiSwapOrderBookBatchMatchInnerError // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookBatchMatchInnerErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookBatchMatchInnerError)
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
		it.Event = new(MiSwapOrderBookBatchMatchInnerError)
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
func (it *MiSwapOrderBookBatchMatchInnerErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookBatchMatchInnerErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookBatchMatchInnerError represents a BatchMatchInnerError event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookBatchMatchInnerError struct {
	Offset *big.Int
	Msg    []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchMatchInnerError is a free log retrieval operation binding the contract event 0x050f709fb65709f10a27682788a9d67fe74de81b310b5c34e3d32f0e2c3ac557.
//
// Solidity: event BatchMatchInnerError(uint256 offset, bytes msg)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterBatchMatchInnerError(opts *bind.FilterOpts) (*MiSwapOrderBookBatchMatchInnerErrorIterator, error) {

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "BatchMatchInnerError")
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookBatchMatchInnerErrorIterator{contract: _MiSwapOrderBook.contract, event: "BatchMatchInnerError", logs: logs, sub: sub}, nil
}

// WatchBatchMatchInnerError is a free log subscription operation binding the contract event 0x050f709fb65709f10a27682788a9d67fe74de81b310b5c34e3d32f0e2c3ac557.
//
// Solidity: event BatchMatchInnerError(uint256 offset, bytes msg)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchBatchMatchInnerError(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookBatchMatchInnerError) (event.Subscription, error) {

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "BatchMatchInnerError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookBatchMatchInnerError)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "BatchMatchInnerError", log); err != nil {
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

// ParseBatchMatchInnerError is a log parse operation binding the contract event 0x050f709fb65709f10a27682788a9d67fe74de81b310b5c34e3d32f0e2c3ac557.
//
// Solidity: event BatchMatchInnerError(uint256 offset, bytes msg)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParseBatchMatchInnerError(log types.Log) (*MiSwapOrderBookBatchMatchInnerError, error) {
	event := new(MiSwapOrderBookBatchMatchInnerError)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "BatchMatchInnerError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiSwapOrderBookEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookEIP712DomainChangedIterator struct {
	Event *MiSwapOrderBookEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookEIP712DomainChanged)
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
		it.Event = new(MiSwapOrderBookEIP712DomainChanged)
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
func (it *MiSwapOrderBookEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookEIP712DomainChanged represents a EIP712DomainChanged event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*MiSwapOrderBookEIP712DomainChangedIterator, error) {

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookEIP712DomainChangedIterator{contract: _MiSwapOrderBook.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookEIP712DomainChanged)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParseEIP712DomainChanged(log types.Log) (*MiSwapOrderBookEIP712DomainChanged, error) {
	event := new(MiSwapOrderBookEIP712DomainChanged)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiSwapOrderBookInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookInitializedIterator struct {
	Event *MiSwapOrderBookInitialized // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookInitialized)
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
		it.Event = new(MiSwapOrderBookInitialized)
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
func (it *MiSwapOrderBookInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookInitialized represents a Initialized event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterInitialized(opts *bind.FilterOpts) (*MiSwapOrderBookInitializedIterator, error) {

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookInitializedIterator{contract: _MiSwapOrderBook.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookInitialized) (event.Subscription, error) {

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookInitialized)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParseInitialized(log types.Log) (*MiSwapOrderBookInitialized, error) {
	event := new(MiSwapOrderBookInitialized)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiSwapOrderBookLogCancelIterator is returned from FilterLogCancel and is used to iterate over the raw logs and unpacked data for LogCancel events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookLogCancelIterator struct {
	Event *MiSwapOrderBookLogCancel // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookLogCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookLogCancel)
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
		it.Event = new(MiSwapOrderBookLogCancel)
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
func (it *MiSwapOrderBookLogCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookLogCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookLogCancel represents a LogCancel event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookLogCancel struct {
	OrderKey [32]byte
	Maker    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogCancel is a free log retrieval operation binding the contract event 0x0ac8bb53fac566d7afc05d8b4df11d7690a7b27bdc40b54e4060f9b21fb849bd.
//
// Solidity: event LogCancel(bytes32 indexed orderKey, address indexed maker)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterLogCancel(opts *bind.FilterOpts, orderKey [][32]byte, maker []common.Address) (*MiSwapOrderBookLogCancelIterator, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "LogCancel", orderKeyRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookLogCancelIterator{contract: _MiSwapOrderBook.contract, event: "LogCancel", logs: logs, sub: sub}, nil
}

// WatchLogCancel is a free log subscription operation binding the contract event 0x0ac8bb53fac566d7afc05d8b4df11d7690a7b27bdc40b54e4060f9b21fb849bd.
//
// Solidity: event LogCancel(bytes32 indexed orderKey, address indexed maker)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchLogCancel(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookLogCancel, orderKey [][32]byte, maker []common.Address) (event.Subscription, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "LogCancel", orderKeyRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookLogCancel)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "LogCancel", log); err != nil {
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

// ParseLogCancel is a log parse operation binding the contract event 0x0ac8bb53fac566d7afc05d8b4df11d7690a7b27bdc40b54e4060f9b21fb849bd.
//
// Solidity: event LogCancel(bytes32 indexed orderKey, address indexed maker)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParseLogCancel(log types.Log) (*MiSwapOrderBookLogCancel, error) {
	event := new(MiSwapOrderBookLogCancel)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "LogCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiSwapOrderBookLogMakeIterator is returned from FilterLogMake and is used to iterate over the raw logs and unpacked data for LogMake events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookLogMakeIterator struct {
	Event *MiSwapOrderBookLogMake // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookLogMakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookLogMake)
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
		it.Event = new(MiSwapOrderBookLogMake)
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
func (it *MiSwapOrderBookLogMakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookLogMakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookLogMake represents a LogMake event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookLogMake struct {
	OrderKey [32]byte
	Side     uint8
	SaleKind uint8
	Maker    common.Address
	Nft      LibOrderAsset
	Price    *big.Int
	Expiry   uint64
	Salt     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogMake is a free log retrieval operation binding the contract event 0xfc37f2ff950f95913eb7182357ba3c14df60ef354bc7d6ab1ba2815f249fffe6.
//
// Solidity: event LogMake(bytes32 orderKey, uint8 indexed side, uint8 indexed saleKind, address indexed maker, (uint256,address,uint96) nft, uint128 price, uint64 expiry, uint64 salt)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterLogMake(opts *bind.FilterOpts, side []uint8, saleKind []uint8, maker []common.Address) (*MiSwapOrderBookLogMakeIterator, error) {

	var sideRule []interface{}
	for _, sideItem := range side {
		sideRule = append(sideRule, sideItem)
	}
	var saleKindRule []interface{}
	for _, saleKindItem := range saleKind {
		saleKindRule = append(saleKindRule, saleKindItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "LogMake", sideRule, saleKindRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookLogMakeIterator{contract: _MiSwapOrderBook.contract, event: "LogMake", logs: logs, sub: sub}, nil
}

// WatchLogMake is a free log subscription operation binding the contract event 0xfc37f2ff950f95913eb7182357ba3c14df60ef354bc7d6ab1ba2815f249fffe6.
//
// Solidity: event LogMake(bytes32 orderKey, uint8 indexed side, uint8 indexed saleKind, address indexed maker, (uint256,address,uint96) nft, uint128 price, uint64 expiry, uint64 salt)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchLogMake(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookLogMake, side []uint8, saleKind []uint8, maker []common.Address) (event.Subscription, error) {

	var sideRule []interface{}
	for _, sideItem := range side {
		sideRule = append(sideRule, sideItem)
	}
	var saleKindRule []interface{}
	for _, saleKindItem := range saleKind {
		saleKindRule = append(saleKindRule, saleKindItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "LogMake", sideRule, saleKindRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookLogMake)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "LogMake", log); err != nil {
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

// ParseLogMake is a log parse operation binding the contract event 0xfc37f2ff950f95913eb7182357ba3c14df60ef354bc7d6ab1ba2815f249fffe6.
//
// Solidity: event LogMake(bytes32 orderKey, uint8 indexed side, uint8 indexed saleKind, address indexed maker, (uint256,address,uint96) nft, uint128 price, uint64 expiry, uint64 salt)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParseLogMake(log types.Log) (*MiSwapOrderBookLogMake, error) {
	event := new(MiSwapOrderBookLogMake)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "LogMake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiSwapOrderBookLogMatchIterator is returned from FilterLogMatch and is used to iterate over the raw logs and unpacked data for LogMatch events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookLogMatchIterator struct {
	Event *MiSwapOrderBookLogMatch // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookLogMatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookLogMatch)
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
		it.Event = new(MiSwapOrderBookLogMatch)
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
func (it *MiSwapOrderBookLogMatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookLogMatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookLogMatch represents a LogMatch event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookLogMatch struct {
	MakeOrderKey [32]byte
	TakeOrderKey [32]byte
	MakeOrder    LibOrderOrder
	TakeOrder    LibOrderOrder
	FillPrice    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogMatch is a free log retrieval operation binding the contract event 0xf629aecab94607bc43ce4aebd564bf6e61c7327226a797b002de724b9944b20e.
//
// Solidity: event LogMatch(bytes32 indexed makeOrderKey, bytes32 indexed takeOrderKey, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) makeOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) takeOrder, uint128 fillPrice)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterLogMatch(opts *bind.FilterOpts, makeOrderKey [][32]byte, takeOrderKey [][32]byte) (*MiSwapOrderBookLogMatchIterator, error) {

	var makeOrderKeyRule []interface{}
	for _, makeOrderKeyItem := range makeOrderKey {
		makeOrderKeyRule = append(makeOrderKeyRule, makeOrderKeyItem)
	}
	var takeOrderKeyRule []interface{}
	for _, takeOrderKeyItem := range takeOrderKey {
		takeOrderKeyRule = append(takeOrderKeyRule, takeOrderKeyItem)
	}

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "LogMatch", makeOrderKeyRule, takeOrderKeyRule)
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookLogMatchIterator{contract: _MiSwapOrderBook.contract, event: "LogMatch", logs: logs, sub: sub}, nil
}

// WatchLogMatch is a free log subscription operation binding the contract event 0xf629aecab94607bc43ce4aebd564bf6e61c7327226a797b002de724b9944b20e.
//
// Solidity: event LogMatch(bytes32 indexed makeOrderKey, bytes32 indexed takeOrderKey, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) makeOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) takeOrder, uint128 fillPrice)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchLogMatch(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookLogMatch, makeOrderKey [][32]byte, takeOrderKey [][32]byte) (event.Subscription, error) {

	var makeOrderKeyRule []interface{}
	for _, makeOrderKeyItem := range makeOrderKey {
		makeOrderKeyRule = append(makeOrderKeyRule, makeOrderKeyItem)
	}
	var takeOrderKeyRule []interface{}
	for _, takeOrderKeyItem := range takeOrderKey {
		takeOrderKeyRule = append(takeOrderKeyRule, takeOrderKeyItem)
	}

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "LogMatch", makeOrderKeyRule, takeOrderKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookLogMatch)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "LogMatch", log); err != nil {
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

// ParseLogMatch is a log parse operation binding the contract event 0xf629aecab94607bc43ce4aebd564bf6e61c7327226a797b002de724b9944b20e.
//
// Solidity: event LogMatch(bytes32 indexed makeOrderKey, bytes32 indexed takeOrderKey, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) makeOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) takeOrder, uint128 fillPrice)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParseLogMatch(log types.Log) (*MiSwapOrderBookLogMatch, error) {
	event := new(MiSwapOrderBookLogMatch)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "LogMatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiSwapOrderBookLogProtocolShareUpdatedIterator is returned from FilterLogProtocolShareUpdated and is used to iterate over the raw logs and unpacked data for LogProtocolShareUpdated events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookLogProtocolShareUpdatedIterator struct {
	Event *MiSwapOrderBookLogProtocolShareUpdated // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookLogProtocolShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookLogProtocolShareUpdated)
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
		it.Event = new(MiSwapOrderBookLogProtocolShareUpdated)
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
func (it *MiSwapOrderBookLogProtocolShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookLogProtocolShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookLogProtocolShareUpdated represents a LogProtocolShareUpdated event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookLogProtocolShareUpdated struct {
	NewShare *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogProtocolShareUpdated is a free log retrieval operation binding the contract event 0x468fe83e6a79b90d6425b0ca7d271b4a27e181df355677b7917626b754dbcc81.
//
// Solidity: event LogProtocolShareUpdated(uint128 indexed newShare)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterLogProtocolShareUpdated(opts *bind.FilterOpts, newShare []*big.Int) (*MiSwapOrderBookLogProtocolShareUpdatedIterator, error) {

	var newShareRule []interface{}
	for _, newShareItem := range newShare {
		newShareRule = append(newShareRule, newShareItem)
	}

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "LogProtocolShareUpdated", newShareRule)
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookLogProtocolShareUpdatedIterator{contract: _MiSwapOrderBook.contract, event: "LogProtocolShareUpdated", logs: logs, sub: sub}, nil
}

// WatchLogProtocolShareUpdated is a free log subscription operation binding the contract event 0x468fe83e6a79b90d6425b0ca7d271b4a27e181df355677b7917626b754dbcc81.
//
// Solidity: event LogProtocolShareUpdated(uint128 indexed newShare)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchLogProtocolShareUpdated(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookLogProtocolShareUpdated, newShare []*big.Int) (event.Subscription, error) {

	var newShareRule []interface{}
	for _, newShareItem := range newShare {
		newShareRule = append(newShareRule, newShareItem)
	}

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "LogProtocolShareUpdated", newShareRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookLogProtocolShareUpdated)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "LogProtocolShareUpdated", log); err != nil {
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

// ParseLogProtocolShareUpdated is a log parse operation binding the contract event 0x468fe83e6a79b90d6425b0ca7d271b4a27e181df355677b7917626b754dbcc81.
//
// Solidity: event LogProtocolShareUpdated(uint128 indexed newShare)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParseLogProtocolShareUpdated(log types.Log) (*MiSwapOrderBookLogProtocolShareUpdated, error) {
	event := new(MiSwapOrderBookLogProtocolShareUpdated)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "LogProtocolShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiSwapOrderBookLogSkipOrderIterator is returned from FilterLogSkipOrder and is used to iterate over the raw logs and unpacked data for LogSkipOrder events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookLogSkipOrderIterator struct {
	Event *MiSwapOrderBookLogSkipOrder // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookLogSkipOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookLogSkipOrder)
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
		it.Event = new(MiSwapOrderBookLogSkipOrder)
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
func (it *MiSwapOrderBookLogSkipOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookLogSkipOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookLogSkipOrder represents a LogSkipOrder event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookLogSkipOrder struct {
	OrderKey [32]byte
	Salt     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSkipOrder is a free log retrieval operation binding the contract event 0x43d1f368251ebe03c021962d50212d072e7ccee5c8ad3f541d93d0dc43bbd420.
//
// Solidity: event LogSkipOrder(bytes32 orderKey, uint64 salt)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterLogSkipOrder(opts *bind.FilterOpts) (*MiSwapOrderBookLogSkipOrderIterator, error) {

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "LogSkipOrder")
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookLogSkipOrderIterator{contract: _MiSwapOrderBook.contract, event: "LogSkipOrder", logs: logs, sub: sub}, nil
}

// WatchLogSkipOrder is a free log subscription operation binding the contract event 0x43d1f368251ebe03c021962d50212d072e7ccee5c8ad3f541d93d0dc43bbd420.
//
// Solidity: event LogSkipOrder(bytes32 orderKey, uint64 salt)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchLogSkipOrder(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookLogSkipOrder) (event.Subscription, error) {

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "LogSkipOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookLogSkipOrder)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "LogSkipOrder", log); err != nil {
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

// ParseLogSkipOrder is a log parse operation binding the contract event 0x43d1f368251ebe03c021962d50212d072e7ccee5c8ad3f541d93d0dc43bbd420.
//
// Solidity: event LogSkipOrder(bytes32 orderKey, uint64 salt)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParseLogSkipOrder(log types.Log) (*MiSwapOrderBookLogSkipOrder, error) {
	event := new(MiSwapOrderBookLogSkipOrder)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "LogSkipOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiSwapOrderBookOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookOwnershipTransferredIterator struct {
	Event *MiSwapOrderBookOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookOwnershipTransferred)
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
		it.Event = new(MiSwapOrderBookOwnershipTransferred)
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
func (it *MiSwapOrderBookOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookOwnershipTransferred represents a OwnershipTransferred event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MiSwapOrderBookOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookOwnershipTransferredIterator{contract: _MiSwapOrderBook.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookOwnershipTransferred)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParseOwnershipTransferred(log types.Log) (*MiSwapOrderBookOwnershipTransferred, error) {
	event := new(MiSwapOrderBookOwnershipTransferred)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiSwapOrderBookPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookPausedIterator struct {
	Event *MiSwapOrderBookPaused // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookPaused)
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
		it.Event = new(MiSwapOrderBookPaused)
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
func (it *MiSwapOrderBookPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookPaused represents a Paused event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterPaused(opts *bind.FilterOpts) (*MiSwapOrderBookPausedIterator, error) {

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookPausedIterator{contract: _MiSwapOrderBook.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookPaused) (event.Subscription, error) {

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookPaused)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParsePaused(log types.Log) (*MiSwapOrderBookPaused, error) {
	event := new(MiSwapOrderBookPaused)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiSwapOrderBookUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MiSwapOrderBook contract.
type MiSwapOrderBookUnpausedIterator struct {
	Event *MiSwapOrderBookUnpaused // Event containing the contract specifics and raw log

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
func (it *MiSwapOrderBookUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiSwapOrderBookUnpaused)
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
		it.Event = new(MiSwapOrderBookUnpaused)
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
func (it *MiSwapOrderBookUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiSwapOrderBookUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiSwapOrderBookUnpaused represents a Unpaused event raised by the MiSwapOrderBook contract.
type MiSwapOrderBookUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MiSwapOrderBookUnpausedIterator, error) {

	logs, sub, err := _MiSwapOrderBook.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MiSwapOrderBookUnpausedIterator{contract: _MiSwapOrderBook.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MiSwapOrderBookUnpaused) (event.Subscription, error) {

	logs, sub, err := _MiSwapOrderBook.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiSwapOrderBookUnpaused)
				if err := _MiSwapOrderBook.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MiSwapOrderBook *MiSwapOrderBookFilterer) ParseUnpaused(log types.Log) (*MiSwapOrderBookUnpaused, error) {
	event := new(MiSwapOrderBookUnpaused)
	if err := _MiSwapOrderBook.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
