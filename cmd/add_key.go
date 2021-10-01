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
	"errors"
	"fmt"

	"github.com/GlassProtocol/didar/composer"
	pb "github.com/GlassProtocol/didar/protos/go"
	"github.com/everFinance/goar/types"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
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
		signingKey, privateKey, err := promptSigningKeys()
		if err != nil {
			panic(err)
		}

		genID, err := promptGenesisID()
		if err != nil {
			panic(err)
		}

		prevID, err := promptPreviousID()
		if err != nil {
			panic(err)
		}

		newKeys, err := promptNewKeys()
		if err != nil {
			panic(err)
		}

		newKeys = append(newKeys, signingKey)

		doc, err := composer.Document(genID, prevID, newKeys, signingKey)
		if err != nil {
			panic(err)
		}

		docSigned, err := composer.SignDocument(doc, privateKey)
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
				Value: genID,
			},
			{
				Name:  "Previous-ID",
				Value: prevID,
			},
			{
				Name:  "Operation",
				Value: "ADD_KEY",
			},
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("\nNEW DOC: %s\n", id)

	},
}

func init() {

	rootCmd.AddCommand(addKeyCmd)

}

func promptGenesisID() (string, error) {
	pubKeyPrompt := promptui.Prompt{
		Label: "Genesis ID",
	}

	return pubKeyPrompt.Run()
}

func promptPreviousID() (string, error) {
	pubKeyPrompt := promptui.Prompt{
		Label: "Previous ID (Genesis if that was last)",
	}

	return pubKeyPrompt.Run()
}

func promptSigningKeys() (*pb.Key, string, error) {

	prompt := promptui.Select{
		Label: "Select Protocol For Signing Keys (same as genesis)",
		Items: []string{"Solana", "Ethereum"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return nil, "", err
	}

	pubKeyPrompt := promptui.Prompt{
		Label: "Public Key",
	}

	pubKey, err := pubKeyPrompt.Run()
	if err != nil {
		return nil, "", err
	}

	privateKeyPrompt := promptui.Prompt{
		Label: "Private Key",
		Mask:  '*',
	}

	privKey, err := privateKeyPrompt.Run()
	if err != nil {
		return nil, "", err
	}

	switch result {
	case "Solana":
		return &pb.Key{
			PublicKey: pubKey,
			KeyType:   pb.KeyType_SOLANA,
		}, privKey, nil

	case "Ethereum":
		return &pb.Key{
			PublicKey: pubKey,
			KeyType:   pb.KeyType_ETHEREUM,
		}, privKey, nil

	default:
		return nil, "", errors.New("promptui protocol switch failed")
	}
}

func promptNewKeys() ([]*pb.Key, error) {

	newKeys := []*pb.Key{}

	for {
		prompt := promptui.Select{
			Label: "Select Operation (Your signing key has already been added)",
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
