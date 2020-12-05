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
	"math/big"

	"github.com/spf13/cobra"
)

// gatherCmd represents the gather command
var gatherCmd = &cobra.Command{
	Use:   "gather",
	Short: "资产归集",
	Long:  `将指定资产归集到一个地址上，0：全部资产，1：黑铁，2：钨矿，3：铂金，4：黑曜，5：钴矿，6：钛金，7：Goblin`,
	Run: func(cmd *cobra.Command, args []string) {
		list := service.QueryAccounts()
		for _, nulsAcc := range list {
			realAcc, _ := account.GetAccountFromPrkey(nulsAcc.PrikeyHex, cfg.ChainId, cfg.AddressPrefix)
			if realAcc.Address == address {
				continue
			}
			if 0 == number || 1 == number {
				realPick(realAcc, cfg.NRC20_Contract_Address_BlackIron)
			}
			if 0 == number || 2 == number {
				realPick(realAcc, cfg.NRC20_Contract_Address_Tungsten)
			}
			if 0 == number || 3 == number {
				realPick(realAcc, cfg.NRC20_Contract_Address_Platinum)
			}
			if 0 == number || 4 == number {
				realPick(realAcc, cfg.NRC20_Contract_Address_Obsidian)
			}
			if 0 == number || 5 == number {
				realPick(realAcc, cfg.NRC20_Contract_Address_Cobalt)
			}
			if 0 == number || 6 == number {
				realPick(realAcc, cfg.NRC20_Contract_Address_Titanium)
			}
			if 0 == number || 7 == number {
				realPick(realAcc, cfg.NRC20_Contract_Address_Goblin)
			}

		}
	},
}

func realPick(realAcc *account.Account, contract string) {
	balanceInfo, _ := cfg.Sdk.GetNRC20Balance(realAcc.Address, contract)
	if balanceInfo == nil || balanceInfo.Amount.Cmp(big.NewInt(100000000)) < 0 {
		return
	}
	balanceStr := balanceInfo.Amount.String()
	txs.SimpleCallContractTx(cfg.Sdk, realAcc, contract, "transfer", [][]string{{address}, {balanceStr}})

}

func init() {
	rootCmd.AddCommand(gatherCmd)
	gatherCmd.Flags().StringVarP(&address, "address", "a", "", "目标地址")
	gatherCmd.MarkFlagRequired("address")
	gatherCmd.Flags().Uint8VarP(&number, "num", "n", 0, `收获矿洞编号：
		0：所有（默认)）
		1：黑铁,BlackIron
		2：钨矿,Tungsten
		3：铂金,Platinum
		4：黑耀,Obsidian
		5：钴矿,Titanium
		6：钛金,Cobalt
		7：Goblin`)
}
