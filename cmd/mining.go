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
	"github.com/niels1286/goblin-cmd/cfg"
	"github.com/niels1286/goblin-cmd/service"
	"github.com/niels1286/goblin-cmd/txs"
	"github.com/niels1286/nuls-go-sdk/account"
	"github.com/spf13/cobra"
	"strconv"
)

// miningCmd represents the mining command
var miningCmd = &cobra.Command{
	Use:   "mining",
	Short: "派出空闲NFT去对应矿洞挖矿",
	Long:  `自动查询所有NFT，对所有空闲NFT根据等级派出到不同的矿洞`,
	Run: func(cmd *cobra.Command, args []string) {
		list := service.QueryAccounts()
		for _, nulsAcc := range list {
			acc, _ := account.GetAccountFromPrkey(nulsAcc.PrikeyHex, cfg.ChainId, cfg.AddressPrefix)
			dataList := queryStatus(acc)
			for _, data := range dataList {
				if data["status"].(string) == "idle" {
					level, _ := strconv.ParseFloat(data["level"].(string), 64)
					tokenId := data["id"].(string)
					where := ""
					if level < 10 {
						where = cfg.BlackIron
					} else if level < 20 {
						where = cfg.Tungsten
					} else if level < 30 {
						where = cfg.Platinum
					} else if level < 40 {
						where = cfg.Obsidian
					} else if level < 60 {
						where = cfg.Cobalt
					} else {
						where = cfg.Titanium
					}
					txs.SimpleCallContractTx(cfg.Sdk, acc, "NULSd6HgoncSA11HYE1nQ2VLVu64XWfGHcsw6", "sendToCave", [][]string{{tokenId}, {where}, {"1"}})
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(miningCmd)
}
