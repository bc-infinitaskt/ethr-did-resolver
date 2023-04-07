package ethr

import "github.com/ethereum/go-ethereum/common"

const (
	DIDAttributeChanged = "DIDAttributeChanged(address,bytes32,bytes,uint,uint)"

	schemaV1 = "https://w3id.org/did/v1"

	DIDPubEd25519   = "did/pub/Ed25519/veriKey/base58"
	DIDPubSecp256k1 = "did/pub/Secp256k1/veriKey/base58"
	DIDPubX25519    = "did/pub/X25519/veriKey/base58"

	ed25519VerificationKey2018        = "Ed25519VerificationKey2018"
	ecdsaSecp256k1VerificationKey2019 = "EcdsaSecp256k1VerificationKey2019"
	x25519KeyAgreementKey2019         = "X25519KeyAgreementKey2019"
)

var EtherDIDKeyTypes = map[string]string{
	DIDPubEd25519:   ed25519VerificationKey2018,
	DIDPubSecp256k1: ecdsaSecp256k1VerificationKey2019,
	DIDPubX25519:    x25519KeyAgreementKey2019,
}
var EmptyContractAddress = common.HexToAddress("0x00")
