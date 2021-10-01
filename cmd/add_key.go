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
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addKeyCmd represents the addKey command
var addKeyCmd = &cobra.Command{
	Use:   "add-key",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch viper.GetString("protocol") {
		case "solana":
			newKeys, err := prompter()
			if err != nil {
				panic(err)
			}

			doc, err := composer.Document(viper.GetString("genesis-id"), viper.GetString("previous-id"), newKeys, &pb.Key{
				KeyType:   pb.KeyType_SOLANA,
				PublicKey: viper.GetString("public-key"),
			})
			if err != nil {
				panic(err)
			}

			docSigned, err := composer.SignDocument(doc)
			if err != nil {
				panic(err)
			}

			jsonBytes, err := marshalOptions.Marshal(docSigned)
			if err != nil {
				panic(err)
			}

			id, err := composer.WriteToArweave(jsonBytes, []types.Tag{
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
				{
					Name:  "Genesis-ID",
					Value: viper.GetString("genesis-id"),
				},
				{
					Name:  "Previous-ID",
					Value: viper.GetString("previous-id"),
				},
				{
					Name:  "Operation",
					Value: viper.GetString("ADD_KEY"),
				},
			})
			if err != nil {
				panic(err)
			}

			fmt.Printf("\nNEW DOC: %s\n", id)

		case "ethereum":
			newKeys, err := prompter()
			if err != nil {
				panic(err)
			}

			doc, err := composer.Document(viper.GetString("genesis-id"), viper.GetString("previous-id"), newKeys, &pb.Key{
				KeyType:   pb.KeyType_ETHEREUM,
				PublicKey: viper.GetString("public-key"),
			})
			if err != nil {
				panic(err)
			}

			docSigned, err := composer.SignDocument(doc)
			if err != nil {
				panic(err)
			}

			jsonBytes, err := marshalOptions.Marshal(docSigned)
			if err != nil {
				panic(err)
			}

			id, err := composer.WriteToArweave(jsonBytes, []types.Tag{
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
				{
					Name:  "Genesis-ID",
					Value: viper.GetString("genesis-id"),
				},
				{
					Name:  "Previous-ID",
					Value: viper.GetString("previous-id"),
				},
				{
					Name:  "Operation",
					Value: viper.GetString("ADD_KEY"),
				},
			})
			if err != nil {
				panic(err)
			}

			fmt.Printf("\nNEW DOC: %s\n", id)
		}
	},
}

func init() {

	addKeyCmd.PersistentFlags().String("genesis-id", "", "genesis tx arweave id")
	addKeyCmd.MarkPersistentFlagRequired("genesis-id")
	viper.BindPFlag("genesis-id", addKeyCmd.PersistentFlags().Lookup("genesis-id"))

	addKeyCmd.PersistentFlags().String("previous-id", "", "the last arweave id for the prior doc")
	addKeyCmd.MarkPersistentFlagRequired("previous-id")
	viper.BindPFlag("previous-id", addKeyCmd.PersistentFlags().Lookup("previous-id"))

	rootCmd.AddCommand(addKeyCmd)

}

func prompter() ([]*pb.Key, error) {

	newKeys := []*pb.Key{}

	for {
		prompt := promptui.Select{
			Label: "Select Operation",
			Items: []string{"Add Key", "Finalize"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			return nil, err
		}

		if result == "Add Key" {
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
				newKeys = append(newKeys, &pb.Key{
					PublicKey: pubKey,
					KeyType:   pb.KeyType_SOLANA,
				})

			case "Ethereum":
				newKeys = append(newKeys, &pb.Key{
					PublicKey: pubKey,
					KeyType:   pb.KeyType_ETHEREUM,
				})

			default:
				fmt.Println("fuck")
			}

		} else {
			break
		}

	}

	return newKeys, nil

}
