package ethr

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/hyperledger/aries-framework-go/pkg/doc/util/jwkkid"

	didcontract "github.com/bc-infinitaskt/ethr-did-resolver/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	diddoc "github.com/hyperledger/aries-framework-go/pkg/doc/did"
	vdrapi "github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr"
	"github.com/hyperledger/aries-framework-go/pkg/kms"
	"github.com/pkg/errors"
)

func (v *VDR) Read(did string, opts ...vdrapi.DIDMethodOption) (*diddoc.DocResolution, error) {
	// split did
	parsedDID, err := diddoc.Parse(did)
	if err != nil {
		return nil, fmt.Errorf("parsing did failed in indy resolver: (%w)", err)
	}

	if parsedDID.Method != v.MethodName {
		return nil, fmt.Errorf("invalid ether method name: %s", parsedDID.MethodSpecificID)
	}

	resOpts := &vdrapi.DIDMethodOpts{}
	for _, opt := range opts {
		opt(resOpts)
	}

	abi, err := didcontract.ContractMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "unable get abi from ContractMetaData")
	}

	instance, err := didcontract.NewContract(v.ContractAddress, v.Client)
	if err != nil {
		return nil, errors.Wrap(err, "unable to instance contract")
	}

	// Get least changed at block no
	ctxc, cancc := context.WithTimeout(context.Background(), v.Timeout)
	defer cancc()
	identity := common.HexToAddress(parsedDID.MethodSpecificID)
	changed, err := instance.Changed(&bind.CallOpts{Context: ctxc}, identity)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get changed at block no")
	}
	logger.Debugf("identity [%s] least changed at block no [%s]", identity.String(), changed.String())

	// Get time of least changed at block no
	ctxb, cancb := context.WithTimeout(context.Background(), v.Timeout)
	defer cancb()
	blockData, err := v.Client.BlockByNumber(ctxb, changed)
	if err != nil {
		return nil, errors.Wrap(err, "unable to query event logs")
	}

	txnTime := time.Unix(int64(blockData.Time()), 0)
	logger.Debugf("identity [%s] least changed at block time [%s]", identity.String(), txnTime.String())

	// Query DIDAttributeChanged from event logs
	ctxFilterLogs, canFilterLogs := context.WithTimeout(context.Background(), v.Timeout)
	defer canFilterLogs()

	didAttributeChanged := new(didcontract.ContractDIDAttributeChanged)
	query := ethereum.FilterQuery{
		FromBlock: changed,
		ToBlock:   changed,
		Addresses: []common.Address{v.ContractAddress},
	}
	filterLogs, err := v.Client.FilterLogs(ctxFilterLogs, query)
	for _, log := range filterLogs {
		switch log.Topics[0].Hex() {
		case FilterDIDAttributeChanged.Hex():
			if err := abi.UnpackIntoInterface(didAttributeChanged, "DIDAttributeChanged", log.Data); err != nil {
				return nil, errors.Wrap(err, "unable to passer to DIDAttributeChanged")
			}
		}
	}

	var KID string
	didAttributeChangedName := string(common.TrimRightZeroes(didAttributeChanged.Name[:]))
	switch didAttributeChangedName {
	case DIDPubEd25519:
		KID, err = jwkkid.CreateKID(didAttributeChanged.Value, kms.ED25519Type)
		if err != nil {
			return nil, fmt.Errorf("unable to create key ID: %v", err)
		}
	//case DIDPubSecp256k1: TODO: will support next version
	//case DIDPubX25519:
	default:
		return nil, fmt.Errorf("not supported VerificationMethod public key type: %s", didAttributeChangedName)
	}

	verificationMethod := diddoc.NewVerificationMethodFromBytes("#"+KID, EtherDIDKeyTypes[didAttributeChangedName], "#id", didAttributeChanged.Value)
	verMethod := diddoc.NewReferencedVerification(verificationMethod, diddoc.Authentication)
	doc := &diddoc.Doc{
		Context:            []string{schemaV1},
		ID:                 did,
		VerificationMethod: []diddoc.VerificationMethod{*verificationMethod},
		Authentication:     []diddoc.Verification{*verMethod},
		Created:            &txnTime,
		Updated:            &txnTime,
	}
	return &diddoc.DocResolution{DIDDocument: doc}, nil
}
