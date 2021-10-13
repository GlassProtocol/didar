package composer

import (
	"crypto/ed25519"
	"errors"

	pb "github.com/GlassProtocol/didar/protos/go"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/protobuf/proto"
)

func SignDocument(didar *pb.Didar, privateKey string) error {

	documentAndAttestation := didar.GetDocumentAndAttestation()

	data, err := proto.Marshal(documentAndAttestation)
	if err != nil {
		return err
	}

	switch documentAndAttestation.Attestation.SigningKey.KeyType {
	case pb.KeyType_ETHEREUM:
		pk, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return err
		}

		hash := crypto.Keccak256Hash(data)

		signature, err := crypto.Sign(hash.Bytes(), pk)
		if err != nil {
			return err
		}

		sig := hexutil.Encode(signature)

		documentAndAttestation.Attestation.Signature = sig

		return nil

	case pb.KeyType_SOLANA:

		decode := base58.Decode(privateKey)
		data := []byte{}

		signature := ed25519.Sign(decode, data)

		documentAndAttestation.Attestation.Signature = base58.Encode(signature)

		return nil
	}

	return errors.New("doc signing switch failed")
}
