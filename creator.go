package ethr

import (
	"fmt"
	"github.com/hyperledger/aries-framework-go/pkg/doc/did"
	vdrapi "github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr"
	"time"
)

// Create builds a new DID Doc (simulate).
func (v *VDR) Create(didDoc *did.Doc, opts ...vdrapi.DIDMethodOption) (*did.DocResolution, error) {
	docOpts := &vdrapi.DIDMethodOpts{Values: make(map[string]interface{})}
	// Apply options
	for _, opt := range opts {
		opt(docOpts)
	}

	docResolution, err := build(didDoc, docOpts)
	if err != nil {
		return nil, fmt.Errorf("create %s DID : %w", v.MethodName, err)
	}

	didDoc = docResolution.DIDDocument
	didKey := "did:ethr:0xb9c5714089478a327f09197987f16f9e5d936e8a" // hardcode for simulate
	didDoc.ID = didKey
	return &did.DocResolution{DIDDocument: didDoc}, nil
}

func build(didDoc *did.Doc, docOpts *vdrapi.DIDMethodOpts) (*did.DocResolution, error) {
	if len(didDoc.VerificationMethod) == 0 && len(didDoc.KeyAgreement) == 0 {
		return nil, fmt.Errorf("verification method and key agreement are empty, at least one should be set")
	}

	mainVM, err := buildDIDVMs(didDoc)
	if err != nil {
		return nil, err
	}

	// Created/Updated time
	t := time.Now()
	assertion := []did.Verification{{
		VerificationMethod: *mainVM,
		Relationship:       did.AssertionMethod,
	}}

	authentication := []did.Verification{{
		VerificationMethod: *mainVM,
		Relationship:       did.Authentication,
	}}

	verificationMethods := []did.VerificationMethod{*mainVM}

	didDoc, err = newDoc(
		verificationMethods,
		did.WithCreatedTime(t),
		did.WithUpdatedTime(t),
		did.WithAuthentication(authentication),
		did.WithAssertion(assertion),
	)
	if err != nil {
		return nil, err
	}

	return &did.DocResolution{DIDDocument: didDoc}, nil
}

func buildDIDVMs(didDoc *did.Doc) (*did.VerificationMethod, error) {
	var verificationMethod *did.VerificationMethod

	if len(didDoc.VerificationMethod) != 0 {
		switch didDoc.VerificationMethod[0].Type {
		case ed25519VerificationKey2018:
			verificationMethod = did.NewVerificationMethodFromBytes(didDoc.VerificationMethod[0].ID, ed25519VerificationKey2018,
				"#id", didDoc.VerificationMethod[0].Value)

		default:
			return nil, fmt.Errorf("not supported VerificationMethod public key type: %s",
				didDoc.VerificationMethod[0].Type)
		}
	}

	return verificationMethod, nil
}

func newDoc(publicKey []did.VerificationMethod, opts ...did.DocOption) (*did.Doc, error) {
	if len(publicKey) == 0 {
		return nil, fmt.Errorf("the did:peer genesis version must include public keys and authentication")
	}

	// build DID Doc
	doc := did.BuildDoc(append([]did.DocOption{did.WithVerificationMethod(publicKey)}, opts...)...)

	// Create a did doc based on the mandatory value: publicKeys & authentication
	if len(doc.Authentication) == 0 || len(doc.VerificationMethod) == 0 {
		return nil, fmt.Errorf("the did must include public keys and authentication")
	}

	return doc, nil
}
