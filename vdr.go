package ethr

import (
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/hyperledger/aries-framework-go/pkg/common/log"
	"github.com/hyperledger/aries-framework-go/pkg/doc/did"
	vdrapi "github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr"
)

const (
	// DIDMethod did method.
	MethodEthr = "ethr"

	DefaultTimeout = time.Second * 60
)

var logger = log.New("vdr/ethr")

type VDR struct {
	MethodName      string
	Client          *ethclient.Client
	ContractAddress *common.Address
	Timeout         time.Duration
}

func New(methodName string, opts ...Option) (*VDR, error) {
	vdri := &VDR{MethodName: methodName, Timeout: DefaultTimeout}
	if methodName == "" {
		methodName = MethodEthr
	}
	for _, opt := range opts {
		opt(vdri)
	}
	if vdri.Client == nil {
		return nil, errors.New("an ethereum client must be set with an option to New")
	}
	if vdri.ContractAddress == nil {
		return nil, errors.New("a contract address must be set with an option to New")
	}
	return vdri, nil
}

func (v *VDR) Accept(method string, opts ...vdrapi.DIDMethodOption) bool {
	return method == v.MethodName
}

func (v *VDR) Update(did *did.Doc, opts ...vdrapi.DIDMethodOption) error {
	return fmt.Errorf("ethr vdr is not supported, use an out-of-band method")
}

func (v *VDR) Deactivate(did string, opts ...vdrapi.DIDMethodOption) error {
	return fmt.Errorf("ethr vdr is not supported, use an out-of-band method")
}

func (v *VDR) Close() error {
	v.Client.Close()
	return nil
}
