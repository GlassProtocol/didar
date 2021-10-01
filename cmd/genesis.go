package cmd

import (
	"errors"
	"fmt"

	"github.com/GlassProtocol/didar/composer"
	pb "github.com/GlassProtocol/didar/protos/go"
	"github.com/everFinance/goar/types"
	"github.com/manifoldco/promptui"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/spf13/cobra"
)

var marshalOptions = protojson.MarshalOptions{
	Multiline:       true,
	EmitUnpopulated: true,
	UseProtoNames:   true,
}

// genesisCmd represents the genesis command
var genesisCmd = &cobra.Command{
	Use:   "genesis",
	Short: "create a genesis transaction",
	Long:  `todo`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := promptGenesis()
		if err != nil {
			panic(err)
		}

		genesis, err := composer.GenesisDoc(key)
		if err != nil {
			panic(err)
		}

		jsonBytes, err := marshalOptions.Marshal(genesis)
		if err != nil {
			panic(err)
		}

		id, err := composer.WriteToArweave(jsonBytes, []types.Tag{
			{
				Name:  "Content-Type",
				Value: "application/json",
			},
			{
				Name:  "Operation",
				Value: "GENESIS",
			},
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("\nGENESIS ID: %s\n", id)
	},
}

func init() {

	rootCmd.AddCommand(genesisCmd)

}

func promptGenesis() (*pb.Key, error) {

	prompt := promptui.Select{
		Label: "Select Protocol",
		Items: []string{"Solana", "Ethereum"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	prompt2 := promptui.Prompt{
		Label: "Public Key",
	}

	pubKey, err := prompt2.Run()
	if err != nil {
		return nil, err
	}

	switch result {
	case "Solana":
		return &pb.Key{
			PublicKey: pubKey,
			KeyType:   pb.KeyType_SOLANA,
		}, nil

	case "Ethereum":
		return &pb.Key{
			PublicKey: pubKey,
			KeyType:   pb.KeyType_ETHEREUM,
		}, nil

	default:
		return nil, errors.New("promptui protocol switch failed")
	}
}
