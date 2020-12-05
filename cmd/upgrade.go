/*
Copyright © 2020 NielsWang niels@nuls.io

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
	"github.com/niels1286/goblin-cmd/cfg"
	"github.com/niels1286/goblin-cmd/service"
	"github.com/niels1286/goblin-cmd/txs"
	"github.com/niels1286/nuls-go-sdk/account"
	"github.com/spf13/cobra"
	"time"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "升级一个NFT",
	Long:  `自动根据nft状态执行撤回、升级`,
	Run: func(cmd *cobra.Command, args []string) {
		list := service.QueryAccounts()
		var realAcc *account.Account
		for _, nulsAcc := range list {
			if nulsAcc.Address == address {
				realAcc, _ = account.GetAccountFromPrkey(nulsAcc.PrikeyHex, cfg.ChainId, cfg.AddressPrefix)
				break
			}
		}
		dataList := queryStatus(realAcc)
		var caveType string
		var idle bool
		for _, data := range dataList {
			if data["id"] == nftId {
				caveType = data["place"].(string)
				if data["status"].(string) == "idle" {
					idle = true
				} else if caveType == "" && data["status"].(string) == "minelp" {
					fmt.Println("NFT在流动池中，本工具咱不处理")
					return
				}
				break
			}
		}
		if !idle {
			txs.SimpleCallContractTx(cfg.Sdk, realAcc, "NULSd6HgoncSA11HYE1nQ2VLVu64XWfGHcsw6", "dismissFromCave", [][]string{{nftId}, {caveType}, {"1"}})
			time.Sleep(10 * time.Second)
		}
		for i := uint8(0); i < count; i++ {
			txs.SimpleCallContractTx(cfg.Sdk, realAcc, "NULSd6HgvnLGUGJ3JwMUmKJVDZbE4JjJeMbeQ", "upgrade", [][]string{{nftId}})
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	upgradeCmd.Flags().StringVarP(&address, "address", "a", "", "地址")
	upgradeCmd.MarkFlagRequired("address")

	upgradeCmd.Flags().StringVarP(&nftId, "nftId", "i", "", "地精id")
	upgradeCmd.MarkFlagRequired("nftId")

	upgradeCmd.Flags().Uint8VarP(&count, "count", "c", 1, "连续升级次数")

}
