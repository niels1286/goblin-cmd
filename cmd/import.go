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
	"github.com/niels1286/nuls-go-sdk/account"
	"github.com/spf13/cobra"
	"strings"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "导入一个私钥，存储在数据库中",
	Long: `导入一个私钥，存储在数据库中
			存储的数据会本地加密，只在实际操作中使用，不提供导出功能
			请自行做好备份。`,
	Run: func(cmd *cobra.Command, args []string) {
		if strings.TrimSpace(prikeyHex) == "" {
			fmt.Println("请输入私钥字符串")
			return
		}

		acc, _ := account.GetAccountFromPrkey(prikeyHex, cfg.ChainId, cfg.AddressPrefix)
		nulsAcc := &service.NULSAccount{
			Address:      acc.Address,
			PrikeyHex:    prikeyHex,
			AddressBytes: acc.AddressBytes,
		}

		service.AddAccount(acc.AddressBytes, nulsAcc)
		fmt.Println("导入成功： " + acc.Address)
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringVarP(&prikeyHex, "prikey", "p", "", "私钥")
	importCmd.MarkFlagRequired("prikey")

}
