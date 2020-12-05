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

	"github.com/spf13/cobra"
)

// autoCmd represents the auto command
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "全自动升级挖矿",
	Long:  `自动分配手续费、根据查询的情况，自动领取奖励，按照预设逻辑自动升级NFT，并排除到对应矿池挖矿`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("逻辑优化中，暂不提供")
		//todo 检查所有地址，余额存在低于1NULS的，则从最多nuls的地址进行手续费发放
		//todo 收获
		//todo 根据规则排序一个带升级nft列表
		//todo 根据顺序进行、轨迹、遣回、升级、派出操作
		//todo 余额不足时完成
		//todo 自动归集goblin到一个地址上，并自动出售
		//todo 未来作为守护进程，定期每小时执行升级操作，每天执行卖出操作
	},
}

func init() {
	rootCmd.AddCommand(autoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// autoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// autoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
