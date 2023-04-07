// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDAttributeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDDelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDOwnerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addDelegateSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"changed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"identityOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttributeSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"revokeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"revokeDelegateSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttributeSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"validDelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506124d5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ff5760003560e01c806380b29f7c11610097578063a7068d6611610066578063a7068d66146102b8578063e476af5c146102d4578063f00d4b5d146102f0578063f96d0f9f1461030c576100ff565b806380b29f7c146102345780638733d4e81461025057806393072684146102805780639c2c1b2b1461029c576100ff565b8063240cf1fa116100d3578063240cf1fa1461019c578063622b2a3c146101b857806370ae92d2146101e85780637ad4b0a414610218576100ff565b8062c023da14610104578063022914a7146101205780630d44625b14610150578063123b5e9814610180575b600080fd5b61011e600480360381019061011991906114b9565b61033c565b005b61013a60048036038101906101359190611528565b61034d565b6040516101479190611564565b60405180910390f35b61016a6004803603810190610165919061157f565b610380565b60405161017791906115eb565b60405180910390f35b61019a6004803603810190610195919061166b565b6103b2565b005b6101b660048036038101906101b19190611729565b61045d565b005b6101d260048036038101906101cd919061157f565b610500565b6040516101df91906117bf565b60405180910390f35b61020260048036038101906101fd9190611528565b6105c6565b60405161020f91906115eb565b60405180910390f35b610232600480360381019061022d91906117da565b6105de565b005b61024e6004803603810190610249919061157f565b6105f1565b005b61026a60048036038101906102659190611528565b610602565b6040516102779190611564565b60405180910390f35b61029a6004803603810190610295919061185d565b6106ad565b005b6102b660048036038101906102b191906118ea565b610754565b005b6102d260048036038101906102cd919061198c565b6107ff565b005b6102ee60048036038101906102e991906119f3565b610812565b005b61030a60048036038101906103059190611a9c565b6108b9565b005b61032660048036038101906103219190611528565b6108c8565b60405161033391906115eb565b60405180910390f35b610348833384846108e0565b505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160205282600052604060002060205281600052604060002060205280600052604060002060009250925050505481565b6000601960f81b600060f81b30600360006103cc8d610602565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548b888888604051602001610422989796959493929190611ceb565b6040516020818303038152906040528051906020012090506104538861044b8a8a8a8a87610a38565b868686610b64565b5050505050505050565b6000601960f81b600060f81b30600360006104778b610602565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205489866040516020016104c996959493929190611dd0565b6040516020818303038152906040528051906020012090506104f8866104f28888888887610a38565b84610cc7565b505050505050565b600080600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000856040516020016105559190611e5a565b60405160208183030381529060405280519060200120815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490504281119150509392505050565b60036020528060005260406000206000915090505481565b6105eb8433858585610b64565b50505050565b6105fd83338484610e96565b505050565b6000806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146106a357809150506106a8565b829150505b919050565b6000601960f81b600060f81b30600360006106c78c610602565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548a878760405160200161071b9796959493929190611ec1565b60405160208183030381529060405280519060200120905061074b876107448989898987610a38565b8585610e96565b50505050505050565b6000601960f81b600060f81b306003600061076e8d610602565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548b8888886040516020016107c4989796959493929190611f99565b6040516020818303038152906040528051906020012090506107f5886107ed8a8a8a8a87610a38565b8686866110a5565b5050505050505050565b61080c84338585856110a5565b50505050565b6000601960f81b600060f81b306003600061082c8c610602565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548a87876040516020016108809796959493929190612082565b6040516020818303038152906040528051906020012090506108b0876108a98989898987610a38565b85856108e0565b50505050505050565b6108c4823383610cc7565b5050565b60026020528060005260406000206000915090505481565b83836108eb82610602565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610958576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094f90612167565b60405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff167f18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e485856000600260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040516109e4949392919061220c565b60405180910390a243600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050505050565b60008060018387878760405160008152602001604052604051610a5e9493929190612267565b6020604051602081039080840390855afa158015610a80573d6000803e3d6000fd5b505050602060405103519050610a9587610602565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610b02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af9906122f8565b60405180910390fd5b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190610b5290612347565b91905055508091505095945050505050565b8484610b6f82610602565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610bdc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bd390612167565b60405180910390fd5b8673ffffffffffffffffffffffffffffffffffffffff167f18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e486868642610c22919061238f565b600260008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051610c7294939291906123e5565b60405180910390a243600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050505050565b8282610cd282610602565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610d3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3690612167565b60405180910390fd5b826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508473ffffffffffffffffffffffffffffffffffffffff167f38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a384600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051610e43929190612431565b60405180910390a243600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b8383610ea182610602565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610f0e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f0590612167565b60405180910390fd5b42600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086604051602001610f619190611e5a565b60405160208183030381529060405280519060200120815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508573ffffffffffffffffffffffffffffffffffffffff167f5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7858542600260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051611051949392919061245a565b60405180910390a243600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050505050565b84846110b082610602565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461111d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111490612167565b60405180910390fd5b8242611129919061238f565b600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008760405160200161117b9190611e5a565b60405160208183030381529060405280519060200120815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508673ffffffffffffffffffffffffffffffffffffffff167f5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f786868642611226919061238f565b600260008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051611276949392919061245a565b60405180910390a243600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061130a826112df565b9050919050565b61131a816112ff565b811461132557600080fd5b50565b60008135905061133781611311565b92915050565b6000819050919050565b6113508161133d565b811461135b57600080fd5b50565b60008135905061136d81611347565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6113c68261137d565b810181811067ffffffffffffffff821117156113e5576113e461138e565b5b80604052505050565b60006113f86112cb565b905061140482826113bd565b919050565b600067ffffffffffffffff8211156114245761142361138e565b5b61142d8261137d565b9050602081019050919050565b82818337600083830152505050565b600061145c61145784611409565b6113ee565b90508281526020810184848401111561147857611477611378565b5b61148384828561143a565b509392505050565b600082601f8301126114a05761149f611373565b5b81356114b0848260208601611449565b91505092915050565b6000806000606084860312156114d2576114d16112d5565b5b60006114e086828701611328565b93505060206114f18682870161135e565b925050604084013567ffffffffffffffff811115611512576115116112da565b5b61151e8682870161148b565b9150509250925092565b60006020828403121561153e5761153d6112d5565b5b600061154c84828501611328565b91505092915050565b61155e816112ff565b82525050565b60006020820190506115796000830184611555565b92915050565b600080600060608486031215611598576115976112d5565b5b60006115a686828701611328565b93505060206115b78682870161135e565b92505060406115c886828701611328565b9150509250925092565b6000819050919050565b6115e5816115d2565b82525050565b600060208201905061160060008301846115dc565b92915050565b600060ff82169050919050565b61161c81611606565b811461162757600080fd5b50565b60008135905061163981611613565b92915050565b611648816115d2565b811461165357600080fd5b50565b6000813590506116658161163f565b92915050565b600080600080600080600060e0888a03121561168a576116896112d5565b5b60006116988a828b01611328565b97505060206116a98a828b0161162a565b96505060406116ba8a828b0161135e565b95505060606116cb8a828b0161135e565b94505060806116dc8a828b0161135e565b93505060a088013567ffffffffffffffff8111156116fd576116fc6112da565b5b6117098a828b0161148b565b92505060c061171a8a828b01611656565b91505092959891949750929550565b600080600080600060a08688031215611745576117446112d5565b5b600061175388828901611328565b95505060206117648882890161162a565b94505060406117758882890161135e565b93505060606117868882890161135e565b925050608061179788828901611328565b9150509295509295909350565b60008115159050919050565b6117b9816117a4565b82525050565b60006020820190506117d460008301846117b0565b92915050565b600080600080608085870312156117f4576117f36112d5565b5b600061180287828801611328565b94505060206118138782880161135e565b935050604085013567ffffffffffffffff811115611834576118336112da565b5b6118408782880161148b565b925050606061185187828801611656565b91505092959194509250565b60008060008060008060c0878903121561187a576118796112d5565b5b600061188889828a01611328565b965050602061189989828a0161162a565b95505060406118aa89828a0161135e565b94505060606118bb89828a0161135e565b93505060806118cc89828a0161135e565b92505060a06118dd89828a01611328565b9150509295509295509295565b600080600080600080600060e0888a031215611909576119086112d5565b5b60006119178a828b01611328565b97505060206119288a828b0161162a565b96505060406119398a828b0161135e565b955050606061194a8a828b0161135e565b945050608061195b8a828b0161135e565b93505060a061196c8a828b01611328565b92505060c061197d8a828b01611656565b91505092959891949750929550565b600080600080608085870312156119a6576119a56112d5565b5b60006119b487828801611328565b94505060206119c58782880161135e565b93505060406119d687828801611328565b92505060606119e787828801611656565b91505092959194509250565b60008060008060008060c08789031215611a1057611a0f6112d5565b5b6000611a1e89828a01611328565b9650506020611a2f89828a0161162a565b9550506040611a4089828a0161135e565b9450506060611a5189828a0161135e565b9350506080611a6289828a0161135e565b92505060a087013567ffffffffffffffff811115611a8357611a826112da565b5b611a8f89828a0161148b565b9150509295509295509295565b60008060408385031215611ab357611ab26112d5565b5b6000611ac185828601611328565b9250506020611ad285828601611328565b9150509250929050565b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b6000819050919050565b611b23611b1e82611adc565b611b08565b82525050565b6000819050919050565b6000611b4e611b49611b44846112df565b611b29565b6112df565b9050919050565b6000611b6082611b33565b9050919050565b6000611b7282611b55565b9050919050565b60008160601b9050919050565b6000611b9182611b79565b9050919050565b6000611ba382611b86565b9050919050565b611bbb611bb682611b67565b611b98565b82525050565b6000819050919050565b611bdc611bd7826115d2565b611bc1565b82525050565b611bf3611bee826112ff565b611b98565b82525050565b600081905092915050565b7f7365744174747269627574650000000000000000000000000000000000000000600082015250565b6000611c3a600c83611bf9565b9150611c4582611c04565b600c82019050919050565b6000819050919050565b611c6b611c668261133d565b611c50565b82525050565b600081519050919050565b600081905092915050565b60005b83811015611ca5578082015181840152602081019050611c8a565b83811115611cb4576000848401525b50505050565b6000611cc582611c71565b611ccf8185611c7c565b9350611cdf818560208601611c87565b80840191505092915050565b6000611cf7828b611b12565b600182019150611d07828a611b12565b600182019150611d178289611baa565b601482019150611d278288611bcb565b602082019150611d378287611be2565b601482019150611d4682611c2d565b9150611d528286611c5a565b602082019150611d628285611cba565b9150611d6e8284611bcb565b6020820191508190509998505050505050505050565b7f6368616e67654f776e6572000000000000000000000000000000000000000000600082015250565b6000611dba600b83611bf9565b9150611dc582611d84565b600b82019050919050565b6000611ddc8289611b12565b600182019150611dec8288611b12565b600182019150611dfc8287611baa565b601482019150611e0c8286611bcb565b602082019150611e1c8285611be2565b601482019150611e2b82611dad565b9150611e378284611be2565b601482019150819050979650505050505050565b611e548161133d565b82525050565b6000602082019050611e6f6000830184611e4b565b92915050565b7f7265766f6b6544656c6567617465000000000000000000000000000000000000600082015250565b6000611eab600e83611bf9565b9150611eb682611e75565b600e82019050919050565b6000611ecd828a611b12565b600182019150611edd8289611b12565b600182019150611eed8288611baa565b601482019150611efd8287611bcb565b602082019150611f0d8286611be2565b601482019150611f1c82611e9e565b9150611f288285611c5a565b602082019150611f388284611be2565b60148201915081905098975050505050505050565b7f61646444656c6567617465000000000000000000000000000000000000000000600082015250565b6000611f83600b83611bf9565b9150611f8e82611f4d565b600b82019050919050565b6000611fa5828b611b12565b600182019150611fb5828a611b12565b600182019150611fc58289611baa565b601482019150611fd58288611bcb565b602082019150611fe58287611be2565b601482019150611ff482611f76565b91506120008286611c5a565b6020820191506120108285611be2565b6014820191506120208284611bcb565b6020820191508190509998505050505050505050565b7f7265766f6b654174747269627574650000000000000000000000000000000000600082015250565b600061206c600f83611bf9565b915061207782612036565b600f82019050919050565b600061208e828a611b12565b60018201915061209e8289611b12565b6001820191506120ae8288611baa565b6014820191506120be8287611bcb565b6020820191506120ce8286611be2565b6014820191506120dd8261205f565b91506120e98285611c5a565b6020820191506120f98284611cba565b915081905098975050505050505050565b600082825260208201905092915050565b7f6261645f6163746f720000000000000000000000000000000000000000000000600082015250565b600061215160098361210a565b915061215c8261211b565b602082019050919050565b6000602082019050818103600083015261218081612144565b9050919050565b600082825260208201905092915050565b60006121a382611c71565b6121ad8185612187565b93506121bd818560208601611c87565b6121c68161137d565b840191505092915050565b6000819050919050565b60006121f66121f16121ec846121d1565b611b29565b6115d2565b9050919050565b612206816121db565b82525050565b60006080820190506122216000830187611e4b565b81810360208301526122338186612198565b905061224260408301856121fd565b61224f60608301846115dc565b95945050505050565b61226181611606565b82525050565b600060808201905061227c6000830187611e4b565b6122896020830186612258565b6122966040830185611e4b565b6122a36060830184611e4b565b95945050505050565b7f6261645f7369676e617475726500000000000000000000000000000000000000600082015250565b60006122e2600d8361210a565b91506122ed826122ac565b602082019050919050565b60006020820190508181036000830152612311816122d5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000612352826115d2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361238457612383612318565b5b600182019050919050565b600061239a826115d2565b91506123a5836115d2565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156123da576123d9612318565b5b828201905092915050565b60006080820190506123fa6000830187611e4b565b818103602083015261240c8186612198565b905061241b60408301856115dc565b61242860608301846115dc565b95945050505050565b60006040820190506124466000830185611555565b61245360208301846115dc565b9392505050565b600060808201905061246f6000830187611e4b565b61247c6020830186611555565b61248960408301856115dc565b61249660608301846115dc565b9594505050505056fea2646970667358221220abaabef4d6f956a17bc52787b9bd9fe439da2cbbfe699ea4894e8b68736dc01464736f6c634300080f0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_Contract *ContractCaller) Changed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "changed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_Contract *ContractSession) Changed(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Changed(&_Contract.CallOpts, arg0)
}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Changed(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Changed(&_Contract.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates(address , bytes32 , address ) view returns(uint256)
func (_Contract *ContractCaller) Delegates(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegates", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates(address , bytes32 , address ) view returns(uint256)
func (_Contract *ContractSession) Delegates(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _Contract.Contract.Delegates(&_Contract.CallOpts, arg0, arg1, arg2)
}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates(address , bytes32 , address ) view returns(uint256)
func (_Contract *ContractCallerSession) Delegates(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _Contract.Contract.Delegates(&_Contract.CallOpts, arg0, arg1, arg2)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(address identity) view returns(address)
func (_Contract *ContractCaller) IdentityOwner(opts *bind.CallOpts, identity common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "identityOwner", identity)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(address identity) view returns(address)
func (_Contract *ContractSession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _Contract.Contract.IdentityOwner(&_Contract.CallOpts, identity)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(address identity) view returns(address)
func (_Contract *ContractCallerSession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _Contract.Contract.IdentityOwner(&_Contract.CallOpts, identity)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Contract *ContractCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Contract *ContractSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Nonce(&_Contract.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Nonce(&_Contract.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(address)
func (_Contract *ContractCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(address)
func (_Contract *ContractSession) Owners(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.Owners(&_Contract.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(address)
func (_Contract *ContractCallerSession) Owners(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.Owners(&_Contract.CallOpts, arg0)
}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(address identity, bytes32 delegateType, address delegate) view returns(bool)
func (_Contract *ContractCaller) ValidDelegate(opts *bind.CallOpts, identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "validDelegate", identity, delegateType, delegate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(address identity, bytes32 delegateType, address delegate) view returns(bool)
func (_Contract *ContractSession) ValidDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	return _Contract.Contract.ValidDelegate(&_Contract.CallOpts, identity, delegateType, delegate)
}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(address identity, bytes32 delegateType, address delegate) view returns(bool)
func (_Contract *ContractCallerSession) ValidDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	return _Contract.Contract.ValidDelegate(&_Contract.CallOpts, identity, delegateType, delegate)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(address identity, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_Contract *ContractTransactor) AddDelegate(opts *bind.TransactOpts, identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addDelegate", identity, delegateType, delegate, validity)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(address identity, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_Contract *ContractSession) AddDelegate(identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddDelegate(&_Contract.TransactOpts, identity, delegateType, delegate, validity)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(address identity, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_Contract *ContractTransactorSession) AddDelegate(identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddDelegate(&_Contract.TransactOpts, identity, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_Contract *ContractTransactor) AddDelegateSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addDelegateSigned", identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_Contract *ContractSession) AddDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddDelegateSigned(&_Contract.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_Contract *ContractTransactorSession) AddDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddDelegateSigned(&_Contract.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address identity, address newOwner) returns()
func (_Contract *ContractTransactor) ChangeOwner(opts *bind.TransactOpts, identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "changeOwner", identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address identity, address newOwner) returns()
func (_Contract *ContractSession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeOwner(&_Contract.TransactOpts, identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address identity, address newOwner) returns()
func (_Contract *ContractTransactorSession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeOwner(&_Contract.TransactOpts, identity, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, address newOwner) returns()
func (_Contract *ContractTransactor) ChangeOwnerSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "changeOwnerSigned", identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, address newOwner) returns()
func (_Contract *ContractSession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeOwnerSigned(&_Contract.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, address newOwner) returns()
func (_Contract *ContractTransactorSession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeOwnerSigned(&_Contract.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(address identity, bytes32 name, bytes value) returns()
func (_Contract *ContractTransactor) RevokeAttribute(opts *bind.TransactOpts, identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeAttribute", identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(address identity, bytes32 name, bytes value) returns()
func (_Contract *ContractSession) RevokeAttribute(identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeAttribute(&_Contract.TransactOpts, identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(address identity, bytes32 name, bytes value) returns()
func (_Contract *ContractTransactorSession) RevokeAttribute(identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeAttribute(&_Contract.TransactOpts, identity, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_Contract *ContractTransactor) RevokeAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeAttributeSigned", identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_Contract *ContractSession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeAttributeSigned(&_Contract.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_Contract *ContractTransactorSession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeAttributeSigned(&_Contract.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(address identity, bytes32 delegateType, address delegate) returns()
func (_Contract *ContractTransactor) RevokeDelegate(opts *bind.TransactOpts, identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeDelegate", identity, delegateType, delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(address identity, bytes32 delegateType, address delegate) returns()
func (_Contract *ContractSession) RevokeDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RevokeDelegate(&_Contract.TransactOpts, identity, delegateType, delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(address identity, bytes32 delegateType, address delegate) returns()
func (_Contract *ContractTransactorSession) RevokeDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RevokeDelegate(&_Contract.TransactOpts, identity, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate) returns()
func (_Contract *ContractTransactor) RevokeDelegateSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeDelegateSigned", identity, sigV, sigR, sigS, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate) returns()
func (_Contract *ContractSession) RevokeDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RevokeDelegateSigned(&_Contract.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate) returns()
func (_Contract *ContractTransactorSession) RevokeDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RevokeDelegateSigned(&_Contract.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(address identity, bytes32 name, bytes value, uint256 validity) returns()
func (_Contract *ContractTransactor) SetAttribute(opts *bind.TransactOpts, identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAttribute", identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(address identity, bytes32 name, bytes value, uint256 validity) returns()
func (_Contract *ContractSession) SetAttribute(identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetAttribute(&_Contract.TransactOpts, identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(address identity, bytes32 name, bytes value, uint256 validity) returns()
func (_Contract *ContractTransactorSession) SetAttribute(identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetAttribute(&_Contract.TransactOpts, identity, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value, uint256 validity) returns()
func (_Contract *ContractTransactor) SetAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAttributeSigned", identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value, uint256 validity) returns()
func (_Contract *ContractSession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetAttributeSigned(&_Contract.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value, uint256 validity) returns()
func (_Contract *ContractTransactorSession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetAttributeSigned(&_Contract.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// ContractDIDAttributeChangedIterator is returned from FilterDIDAttributeChanged and is used to iterate over the raw logs and unpacked data for DIDAttributeChanged events raised by the Contract contract.
type ContractDIDAttributeChangedIterator struct {
	Event *ContractDIDAttributeChanged // Event containing the contract specifics and raw log

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
func (it *ContractDIDAttributeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDIDAttributeChanged)
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
		it.Event = new(ContractDIDAttributeChanged)
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
func (it *ContractDIDAttributeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDIDAttributeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDIDAttributeChanged represents a DIDAttributeChanged event raised by the Contract contract.
type ContractDIDAttributeChanged struct {
	Identity       common.Address
	Name           [32]byte
	Value          []byte
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDAttributeChanged is a free log retrieval operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: event DIDAttributeChanged(address indexed identity, bytes32 name, bytes value, uint256 validTo, uint256 previousChange)
func (_Contract *ContractFilterer) FilterDIDAttributeChanged(opts *bind.FilterOpts, identity []common.Address) (*ContractDIDAttributeChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &ContractDIDAttributeChangedIterator{contract: _Contract.contract, event: "DIDAttributeChanged", logs: logs, sub: sub}, nil
}

// WatchDIDAttributeChanged is a free log subscription operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: event DIDAttributeChanged(address indexed identity, bytes32 name, bytes value, uint256 validTo, uint256 previousChange)
func (_Contract *ContractFilterer) WatchDIDAttributeChanged(opts *bind.WatchOpts, sink chan<- *ContractDIDAttributeChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDIDAttributeChanged)
				if err := _Contract.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
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

// ParseDIDAttributeChanged is a log parse operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: event DIDAttributeChanged(address indexed identity, bytes32 name, bytes value, uint256 validTo, uint256 previousChange)
func (_Contract *ContractFilterer) ParseDIDAttributeChanged(log types.Log) (*ContractDIDAttributeChanged, error) {
	event := new(ContractDIDAttributeChanged)
	if err := _Contract.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDIDDelegateChangedIterator is returned from FilterDIDDelegateChanged and is used to iterate over the raw logs and unpacked data for DIDDelegateChanged events raised by the Contract contract.
type ContractDIDDelegateChangedIterator struct {
	Event *ContractDIDDelegateChanged // Event containing the contract specifics and raw log

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
func (it *ContractDIDDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDIDDelegateChanged)
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
		it.Event = new(ContractDIDDelegateChanged)
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
func (it *ContractDIDDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDIDDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDIDDelegateChanged represents a DIDDelegateChanged event raised by the Contract contract.
type ContractDIDDelegateChanged struct {
	Identity       common.Address
	DelegateType   [32]byte
	Delegate       common.Address
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDDelegateChanged is a free log retrieval operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: event DIDDelegateChanged(address indexed identity, bytes32 delegateType, address delegate, uint256 validTo, uint256 previousChange)
func (_Contract *ContractFilterer) FilterDIDDelegateChanged(opts *bind.FilterOpts, identity []common.Address) (*ContractDIDDelegateChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DIDDelegateChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &ContractDIDDelegateChangedIterator{contract: _Contract.contract, event: "DIDDelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDIDDelegateChanged is a free log subscription operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: event DIDDelegateChanged(address indexed identity, bytes32 delegateType, address delegate, uint256 validTo, uint256 previousChange)
func (_Contract *ContractFilterer) WatchDIDDelegateChanged(opts *bind.WatchOpts, sink chan<- *ContractDIDDelegateChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DIDDelegateChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDIDDelegateChanged)
				if err := _Contract.contract.UnpackLog(event, "DIDDelegateChanged", log); err != nil {
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

// ParseDIDDelegateChanged is a log parse operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: event DIDDelegateChanged(address indexed identity, bytes32 delegateType, address delegate, uint256 validTo, uint256 previousChange)
func (_Contract *ContractFilterer) ParseDIDDelegateChanged(log types.Log) (*ContractDIDDelegateChanged, error) {
	event := new(ContractDIDDelegateChanged)
	if err := _Contract.contract.UnpackLog(event, "DIDDelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDIDOwnerChangedIterator is returned from FilterDIDOwnerChanged and is used to iterate over the raw logs and unpacked data for DIDOwnerChanged events raised by the Contract contract.
type ContractDIDOwnerChangedIterator struct {
	Event *ContractDIDOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ContractDIDOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDIDOwnerChanged)
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
		it.Event = new(ContractDIDOwnerChanged)
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
func (it *ContractDIDOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDIDOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDIDOwnerChanged represents a DIDOwnerChanged event raised by the Contract contract.
type ContractDIDOwnerChanged struct {
	Identity       common.Address
	Owner          common.Address
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDOwnerChanged is a free log retrieval operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: event DIDOwnerChanged(address indexed identity, address owner, uint256 previousChange)
func (_Contract *ContractFilterer) FilterDIDOwnerChanged(opts *bind.FilterOpts, identity []common.Address) (*ContractDIDOwnerChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DIDOwnerChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &ContractDIDOwnerChangedIterator{contract: _Contract.contract, event: "DIDOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchDIDOwnerChanged is a free log subscription operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: event DIDOwnerChanged(address indexed identity, address owner, uint256 previousChange)
func (_Contract *ContractFilterer) WatchDIDOwnerChanged(opts *bind.WatchOpts, sink chan<- *ContractDIDOwnerChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DIDOwnerChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDIDOwnerChanged)
				if err := _Contract.contract.UnpackLog(event, "DIDOwnerChanged", log); err != nil {
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

// ParseDIDOwnerChanged is a log parse operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: event DIDOwnerChanged(address indexed identity, address owner, uint256 previousChange)
func (_Contract *ContractFilterer) ParseDIDOwnerChanged(log types.Log) (*ContractDIDOwnerChanged, error) {
	event := new(ContractDIDOwnerChanged)
	if err := _Contract.contract.UnpackLog(event, "DIDOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
