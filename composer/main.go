package composer

import (
	"fmt"

	pb "github.com/GlassProtocol/didar/protos/go"
)

func GenesisDoc(key *pb.Key) (*pb.Genesis, error) {
	return &pb.Genesis{
		SigningKey: key,
	}, nil
}

func Document(id string, prevDocId string, newKeys []*pb.Key, signingKey *pb.Key) (*pb.Document, error) {
	doc := &pb.Document{
		Context:        []string{"https://www.w3.org/ns/did/v1"},
		Id:             fmt.Sprintf("did:ar:%s", id),
		Authentication: newKeys,
		Reference: &pb.Reference{
			PreviousDocumentId: prevDocId,
			SigningKey:         signingKey,
		},
	}

	return doc, nil
}
