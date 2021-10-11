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

func SignDocument(doc *pb.Document, privateKey string) (*pb.Document, error) {
	switch doc.Reference.SigningKey.KeyType {
	case pb.KeyType_ETHEREUM:
		pk, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return nil, err
		}

		data := []byte{}

		for _, x := range doc.Authentication {
			tmp, err := proto.Marshal(x)
			if err != nil {
				return nil, err
			}
			data = append(data, tmp...)
		}

		for _, m := range doc.Metadata {
			tmp, err := proto.Marshal(m)
			if err != nil {
				return nil, err
			}
			data = append(data, tmp...)
		}

		hash := crypto.Keccak256Hash(append(data, []byte(doc.Id)...))

		signature, err := crypto.Sign(hash.Bytes(), pk)
		if err != nil {
			return nil, err
		}

		sig := hexutil.Encode(signature)

		doc.Reference.Signature = sig

		return doc, nil

	case pb.KeyType_SOLANA:

		decode := base58.Decode(privateKey)
		data := []byte{}

		for _, x := range doc.Authentication {
			tmp, err := proto.Marshal(x)
			if err != nil {
				return nil, err
			}
			data = append(data, tmp...)
		}
		signature := ed25519.Sign(decode, append(data, []byte(doc.Id)...))

		doc.Reference.Signature = base58.Encode(signature)

		return doc, nil

		// TODO: implement arweave signing
		// case pb.KeyType_ARWEAVE:

	}
	return nil, errors.New("doc signing switch failed")
}
