package composer

import (
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
)

func WriteToArweave(wallet *goar.Wallet, data []byte, tags []types.Tag) (string, error) {
	return wallet.SendData(
		data, // Data bytes
		tags,
	)
}
