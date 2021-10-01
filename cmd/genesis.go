/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/GlassProtocol/didar/composer"
	pb "github.com/GlassProtocol/didar/protos/go"
	"github.com/everFinance/goar/types"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		switch viper.GetString("protocol") {
		case "solana":
			g, err := composer.GenesisDoc(&pb.Key{
				KeyType:   pb.KeyType_SOLANA,
				PublicKey: viper.GetString("public-key"),
			})
			if err != nil {
				panic(err)
			}

			gsigned, err := composer.SignGenesis(g)
			if err != nil {
				panic(err)
			}

			jsonBytes, err := marshalOptions.Marshal(gsigned)
			if err != nil {
				panic(err)
			}

			id, err := composer.WriteToArweave(jsonBytes, []types.Tag{
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
			})
			if err != nil {
				panic(err)
			}

			fmt.Printf("\nGENESIS: %s\n", id)

		case "ethereum":
			g, err := composer.GenesisDoc(&pb.Key{
				KeyType:   pb.KeyType_ETHEREUM,
				PublicKey: viper.GetString("public-key"),
			})
			if err != nil {
				panic(err)
			}

			gsigned, err := composer.SignGenesis(g)
			if err != nil {
				panic(err)
			}

			jsonBytes, err := marshalOptions.Marshal(gsigned)
			if err != nil {
				panic(err)
			}

			id, err := composer.WriteToArweave(jsonBytes, []types.Tag{
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
			})
			if err != nil {
				panic(err)
			}

			fmt.Printf("\nGENESIS: %s\n", id)

		// TODO: add arweave
		default:

		}

		fmt.Println("genesis called")
	},
}

func init() {

	rootCmd.AddCommand(genesisCmd)

}
