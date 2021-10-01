package composer

import (
	"crypto/ed25519"
	"errors"

	pb "github.com/GlassProtocol/didar/protos/go"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
)

func SignGenesis(genesis *pb.Genesis) (*pb.Genesis, error) {

	switch genesis.SigningKey.KeyType {
	case pb.KeyType_ETHEREUM:
		pk, err := crypto.HexToECDSA(viper.GetString("private-key"))
		if err != nil {
			return nil, err
		}

		skBytes, err := proto.Marshal(genesis.SigningKey)
		if err != nil {
			return nil, err
		}

		hash := crypto.Keccak256Hash(append(skBytes, []byte(genesis.Nonce)...))

		signature, err := crypto.Sign(hash.Bytes(), pk)
		if err != nil {
			return nil, err
		}

		sig := hexutil.Encode(signature)

		genesis.NonceSignature = sig

		return genesis, nil

	case pb.KeyType_SOLANA:

		decode := base58.Decode(viper.GetString("private-key"))

		skBytes, err := proto.Marshal(genesis.SigningKey)
		if err != nil {
			return nil, err
		}

		signature := ed25519.Sign(decode, append(skBytes, []byte(genesis.Nonce)...))

		genesis.NonceSignature = base58.Encode(signature)

		base58.Encode(signature)

	// TODO: implement arweave signing
	// case pb.KeyType_ARWEAVE:

	default:
		return nil, errors.New("key is of unknown type")
	}

	return nil, errors.New("genesis signing switch failed")
}

func SignDocument(doc *pb.Document) (*pb.Document, error) {
	switch doc.Reference.SigningKey.KeyType {
	case pb.KeyType_ETHEREUM:
		pk, err := crypto.HexToECDSA(viper.GetString("private-key"))
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

		hash := crypto.Keccak256Hash(append(data, []byte(doc.Id)...))

		signature, err := crypto.Sign(hash.Bytes(), pk)
		if err != nil {
			return nil, err
		}

		sig := hexutil.Encode(signature)

		doc.Reference.Signature = sig

		return doc, nil

	case pb.KeyType_SOLANA:

		decode := base58.Decode(viper.GetString("private-key"))
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
