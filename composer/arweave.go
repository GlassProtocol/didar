package composer

import (
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/spf13/viper"
)

func WriteToArweave(data []byte, tags []types.Tag) (string, error) {
	w, err := goar.NewWalletFromPath(viper.GetString("arweave-key"), "https://arweave.net")
	if err != nil {
		return "", err
	}

	return w.SendData(
		data, // Data bytes
		tags,
	)
}
