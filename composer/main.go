package composer

import (
	"fmt"

	pb "github.com/GlassProtocol/didar/protos/go"
)

const VERSION = "2021-10-12"

func Genesis(key *pb.Key) *pb.Didar {
	return &pb.Didar{
		Version: VERSION,
		Data: &pb.Didar_Genesis{
			Genesis: &pb.Genesis{
				SigningKey: key,
			},
		},
	}
}

func Didar(id string, appending string, newKeys []*pb.Key, signingKey *pb.Key, metadata map[string]string) *pb.Didar {
	return &pb.Didar{
		Version: VERSION,
		Data: &pb.Didar_DocumentAndAttestation{
			DocumentAndAttestation: &pb.DocumentAndAttestation{
				Document: &pb.Document{
					Context:        []string{"https://www.w3.org/ns/did/v1"},
					Id:             fmt.Sprintf("did:ar:%s", id),
					Authentication: newKeys,
					Metadata:       metadata,
				},
				Attestation: &pb.Attestation{
					Appending:  appending,
					SigningKey: signingKey,
				},
			},
		},
	}
}
