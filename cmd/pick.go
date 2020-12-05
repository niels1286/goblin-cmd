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
)

// pickCmd represents the pick command
var pickCmd = &cobra.Command{
	Use:   "pick",
	Short: "收获",
	Long: `收获一种矿石和goblin,参数说明：
		0：所有（默认)）
		1：黑铁,BlackIron
		2：钨矿,Tungsten
		3：铂金,Platinum
		4：黑耀,Obsidian
		5：钴矿,Titanium
		6：钛金,Cobalt`,
	Run: func(cmd *cobra.Command, args []string) {
		list := service.QueryAccounts()
		for _, nulsAcc := range list {
			if address != "" && address != nulsAcc.Address {
				continue
			}
			acc, _ := account.GetAccountFromPrkey(nulsAcc.PrikeyHex, cfg.ChainId, cfg.AddressPrefix)
			status := queryAccountStatus(acc)
			if status.biPool && (0 == number || 1 == number) {
				ClaimEarned(acc, cfg.BlackIron)
			}
			if status.tuPool && (0 == number || 2 == number) {
				ClaimEarned(acc, cfg.Tungsten)
			}
			if status.plPool && (0 == number || 3 == number) {
				ClaimEarned(acc, cfg.Platinum)
			}
			if status.obPool && (0 == number || 4 == number) {
				ClaimEarned(acc, cfg.Obsidian)
			}
			if status.coPool && (0 == number || 5 == number) {
				ClaimEarned(acc, cfg.Cobalt)
			}
			if status.tiPool && (0 == number || 6 == number) {
				ClaimEarned(acc, cfg.Titanium)
			}
		}
	},
}

type NftStatus struct {
	biPool bool
	tuPool bool
	plPool bool
	obPool bool
	tiPool bool
	coPool bool
}

func queryAccountStatus(acc *account.Account) *NftStatus {
	dataList := queryStatus(acc)
	result := &NftStatus{}
	for _, data := range dataList {
		if data["place"].(string) == cfg.BlackIron {
			result.biPool = true
		} else if data["place"].(string) == cfg.Tungsten {
			result.tuPool = true
		} else if data["place"].(string) == cfg.Platinum {
			result.plPool = true
		} else if data["place"].(string) == cfg.Obsidian {
			result.obPool = true
		} else if data["place"].(string) == cfg.Titanium {
			result.tiPool = true
		} else if data["place"].(string) == cfg.Cobalt {
			result.coPool = true
		}
	}
	return result
}

func ClaimEarned(acc *account.Account, where string) {
	txs.SimpleCallContractTx(cfg.Sdk, acc, "NULSd6HguK98JD4yFfjYkDTq8VXfPYDNeFMiL", cfg.ClaimEarned, [][]string{{where}})
}

func init() {
	rootCmd.AddCommand(pickCmd)
	pickCmd.Flags().StringVarP(&address, "address", "a", "", "单个地址收矿，为空时默认全部账户收矿")
	pickCmd.Flags().Uint8VarP(&number, "num", "n", 0, `收获矿洞编号：
		0：所有（默认)）
		1：黑铁,BlackIron
		2：钨矿,Tungsten
		3：铂金,Platinum
		4：黑耀,Obsidian
		5：钴矿,Titanium
		6：钛金,Cobalt`)
}
