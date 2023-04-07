package ethr

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
	"time"
)

type Option func(opts *VDR)

// WithDial url for dial
func WithDial(url string) Option {
	return func(opts *VDR) {
		c, err := ethclient.Dial(url)
		if err != nil {
			panic(fmt.Sprintf("unable to dial: [%s], error: [%s]", url, err.Error()))
		}
		opts.Client = c
	}
}

// WithContractAddress is contractAddress with EIP-1056
func WithContractAddress(contractAddress string) Option {
	return func(opts *VDR) {
		*opts.ContractAddress = common.HexToAddress(contractAddress)
	}
}

// WithTimeout set timeout
func WithTimeout(timeout int64) Option {
	return func(opts *VDR) {
		opts.Timeout = time.Second * time.Duration(timeout)
	}
}

// WithRawDial connect with rawurl example: https://rpc.example:0xf3beac30c498d9e26865f34fcaa57dbb935b0d74
func WithRawDial(rawurl string) Option {
	return func(opts *VDR) {
		params := strings.Split(rawurl, ":")
		paramsLength := len(params)
		if paramsLength <= 2 {
			log.Panicf("can't parse rawurl [%s] params is less than 3", rawurl)
		}

		address := params[paramsLength-1]
		url := strings.Replace(rawurl, address, "", -1)
		WithDial(url)
		WithContractAddress(address)
	}
}
