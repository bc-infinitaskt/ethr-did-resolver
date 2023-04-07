package ethr

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	didcontract "github.com/bc-infinitaskt/ethr-did-resolver/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	diddoc "github.com/hyperledger/aries-framework-go/pkg/doc/did"
	vdrapi "github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr"
	"github.com/hyperledger/aries-framework-go/pkg/kms"
	"github.com/hyperledger/aries-framework-go/pkg/kms/localkms"
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

	contractAbi, err := abi.JSON(strings.NewReader(didcontract.ContractMetaData.ABI))
	if err != nil {
		return nil, errors.Wrap(err, "unable to passer contract abi")
	}

	instance, err := didcontract.NewContract(v.ContractAddress, v.Client)
	if err != nil {
		return nil, errors.Wrap(err, "unable to instance contract")
	}

	// get blockNumber for query range
	identity := common.HexToAddress(parsedDID.MethodSpecificID)
	blockNumber, err := instance.Changed(&bind.CallOpts{}, identity)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get BlockNumber")
	}

	ctx, cancel := context.WithTimeout(context.Background(), v.Timeout)
	defer cancel()

	// prepare for query event logs
	didAttributeChangedEvent := crypto.Keccak256Hash([]byte(DIDAttributeChanged))
	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{v.ContractAddress},
	}

	// get time from BlockByNumber
	blockData, err := v.Client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		return nil, errors.Wrap(err, "unable to query event logs")
	}
	txnTime := time.Unix(int64(blockData.Time()), 0)

	// query event logs
	filterLogs, err := v.Client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "unable to query event logs")
	}

	// Unpack log to DIDAttributeChanged struct
	didAttributeChanged := &didcontract.ContractDIDAttributeChanged{}
	for _, log := range filterLogs {
		switch log.Topics[0].Hex() {
		case didAttributeChangedEvent.Hex():
			if err := contractAbi.UnpackIntoInterface(didAttributeChanged, "DIDAttributeChanged", log.Data); err != nil {
				return nil, errors.Wrap(err, "unable to passer to DIDAttributeChanged")
			}
		}
	}

	//TODO: remove when no need to show log Debug
	didAttributeChangedLog, _ := json.MarshalIndent(didAttributeChanged, "", "  ")
	logger.Debugf("DidAttributeChangedLog: %s", string(didAttributeChangedLog))

	var KID string
	didAttributeChangedName := string(didAttributeChanged.Name[:])
	switch didAttributeChangedName {
	case DIDPubEd25519:
		KID, err = localkms.CreateKID(didAttributeChanged.Value, kms.ED25519Type)
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
