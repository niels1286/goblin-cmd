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
	"encoding/json"
	"fmt"
	"github.com/niels1286/goblin-cmd/cfg"
	"github.com/niels1286/goblin-cmd/service"
	"github.com/niels1286/nuls-go-sdk/account"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"strconv"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "查询所有账户状态",
	Long:  `查询每个账户的相关资产余额，待领取数量，nft信息`,
	Run: func(cmd *cobra.Command, args []string) {
		accountList := service.QueryAccounts()
		var sum0, sum1, sum2, sum3, sum4, sum11, sum21, sum31, sum41, sum5, sum51, sum6, sum61 float64
		var nulsSum, gobSum, biSum, tuSum, plSum, obSum, coSum, tiSum float64
		for index, acco := range accountList {
			if address != "" && address != acco.Address {
				continue
			}

			addr := acco.Address
			acc, _ := account.GetAccountFromPrkey(acco.PrikeyHex, cfg.ChainId, cfg.AddressPrefix)
			result := fmt.Sprintf("%d", index+1) + ", "
			nInfo, _ := cfg.Sdk.GetBalance(acc.Address, cfg.ChainId, cfg.AssetsId)
			nuls := decimal.NewFromBigInt(nInfo.Balance, -8).Round(2).String()
			gInfo, _ := cfg.Sdk.GetNRC20Balance(acc.Address, cfg.NRC20_Contract_Address_Goblin)
			goblin := decimal.NewFromBigInt(gInfo.Amount, -8).Round(2).String()
			bInfo, _ := cfg.Sdk.GetNRC20Balance(acc.Address, cfg.NRC20_Contract_Address_BlackIron)
			blackIron := decimal.NewFromBigInt(bInfo.Amount, -8).Round(2).String()
			tInfo, _ := cfg.Sdk.GetNRC20Balance(acc.Address, cfg.NRC20_Contract_Address_Tungsten)
			tungsten := decimal.NewFromBigInt(tInfo.Amount, -8).Round(2).String()
			pInfo, _ := cfg.Sdk.GetNRC20Balance(acc.Address, cfg.NRC20_Contract_Address_Platinum)
			platinum := decimal.NewFromBigInt(pInfo.Amount, -8).Round(2).String()
			oInfo, _ := cfg.Sdk.GetNRC20Balance(acc.Address, cfg.NRC20_Contract_Address_Obsidian)
			obsidian := decimal.NewFromBigInt(oInfo.Amount, -8).Round(2).String()

			coInfo, _ := cfg.Sdk.GetNRC20Balance(acc.Address, cfg.NRC20_Contract_Address_Cobalt)
			cobalt := decimal.NewFromBigInt(coInfo.Amount, -8).Round(2).String()
			tiInfo, _ := cfg.Sdk.GetNRC20Balance(acc.Address, cfg.NRC20_Contract_Address_Titanium)
			titanium := decimal.NewFromBigInt(tiInfo.Amount, -8).Round(2).String()

			result += acc.Address + " ,  nuls(" + nuls + "), gob(" + goblin + "), bla(" + blackIron + "), tun(" + tungsten +
				"), pla(" + platinum + "), obs(" + obsidian + "), cob( " + cobalt + "), tit(" + titanium + ")"
			val, _ := strconv.ParseFloat(nuls, 64)
			nulsSum += val
			val, _ = strconv.ParseFloat(goblin, 64)
			gobSum += val
			val, _ = strconv.ParseFloat(blackIron, 64)
			biSum += val
			val, _ = strconv.ParseFloat(tungsten, 64)
			tuSum += val
			val, _ = strconv.ParseFloat(platinum, 64)
			plSum += val
			val, _ = strconv.ParseFloat(obsidian, 64)
			obSum += val
			val, _ = strconv.ParseFloat(cobalt, 64)
			coSum += val
			val, _ = strconv.ParseFloat(titanium, 64)
			tiSum += val

			i := ""
			swap := ""
			bi := ""
			t := ""
			p := ""
			o := ""
			ti := ""
			co := ""

			nfts := queryStatus(acc)
			for _, nft := range nfts {
				if nft["status"] == "idle" {
					i += nft["id"].(string) + "(" + nft["level"].(string) + "), "
					continue
				} else if nft["status"] == "minelp" {
					swap += nft["id"].(string) + "(" + nft["level"].(string) + "), "
					continue
				}
				where := nft["place"]
				if where == "BlackIron" {
					bi += nft["id"].(string) + "(" + nft["level"].(string) + "), "
				} else if where == "Tungsten" {
					t += nft["id"].(string) + "(" + nft["level"].(string) + "), "
				} else if where == "Platinum" {
					p += nft["id"].(string) + "(" + nft["level"].(string) + "), "
				} else if where == "Obsidian" {
					o += nft["id"].(string) + "(" + nft["level"].(string) + "), "
				} else if where == cfg.Titanium {
					ti += nft["id"].(string) + "(" + nft["level"].(string) + "), "
				} else if where == cfg.Cobalt {
					co += nft["id"].(string) + "(" + nft["level"].(string) + "), "
				}
			}
			if len(swap) > 0 {
				val := earnedOfMineLP(addr)
				sum, _ := strconv.ParseFloat(val, 64)
				sum0 += sum
				result += "\n流动性池： " + swap + getSpace(len(swap)) + "待收获-gob:" + val
			}
			if len(ti) > 0 {
				val1, val2 := earnedOf(cfg.Titanium, addr)
				suma, _ := strconv.ParseFloat(val1, 64)
				sum6 += suma
				sum, _ := strconv.ParseFloat(val2, 64)
				sum61 += sum
				result += "\n钛金矿池： " + ti + getSpace(len(ti)) + "待收获-gob:" + val1 + ", tit:" + val2
			}
			if len(co) > 0 {
				val1, val2 := earnedOf(cfg.Cobalt, addr)
				suma, _ := strconv.ParseFloat(val1, 64)
				sum5 += suma
				sum, _ := strconv.ParseFloat(val2, 64)
				sum51 += sum
				result += "\n钴矿矿池： " + co + getSpace(len(co)) + "待收获-gob:" + val1 + ", cob:" + val2
			}
			if len(o) > 0 {
				val1, val2 := earnedOf(cfg.Obsidian, addr)
				suma, _ := strconv.ParseFloat(val1, 64)
				sum4 += suma
				sum, _ := strconv.ParseFloat(val2, 64)
				sum41 += sum
				result += "\n黑耀矿池： " + o + getSpace(len(o)) + "待收获-gob:" + val1 + ", obs:" + val2
			}
			if len(p) > 0 {
				val1, val2 := earnedOf(cfg.Platinum, addr)
				suma, _ := strconv.ParseFloat(val1, 64)
				sum3 += suma
				sum, _ := strconv.ParseFloat(val2, 64)
				sum31 += sum
				result += "\n铂金矿池： " + p + getSpace(len(p)) + "待收获-gob:" + val1 + ", pla:" + val2
			}
			if len(t) > 0 {
				val1, val2 := earnedOf(cfg.Tungsten, addr)
				suma, _ := strconv.ParseFloat(val1, 64)
				sum2 += suma
				sum, _ := strconv.ParseFloat(val2, 64)
				sum21 += sum
				result += "\n钨矿矿池： " + t + getSpace(len(t)) + "待收获-gob:" + val1 + ", tun:" + val2
			}
			if len(bi) > 0 {
				val1, val2 := earnedOf(cfg.BlackIron, addr)
				suma, _ := strconv.ParseFloat(val1, 64)
				sum1 += suma
				sum, _ := strconv.ParseFloat(val2, 64)
				sum11 += sum
				result += "\n黑铁矿池： " + bi + getSpace(len(bi)) + "待收获-gob:" + val1 + ", bla:" + val2
			}
			if len(i) > 0 {
				result += "\n空闲休息： " + i
			}
			fmt.Println(result)
		}
		if address == "" {
			fmt.Println("---------------------------------------------------待收取汇总------------------------------------------------------")
			fmt.Println("流动性池：Goblin = " + fmt.Sprintf("%f", sum0))
			fmt.Println("黑铁矿池：Goblin = " + fmt.Sprintf("%f", sum1) + ", 额外奖励：" + fmt.Sprintf("%f", sum11))
			fmt.Println("钨矿矿池：Goblin = " + fmt.Sprintf("%f", sum2) + ", 额外奖励：" + fmt.Sprintf("%f", sum21))
			fmt.Println("铂金矿池：Goblin = " + fmt.Sprintf("%f", sum3) + ", 额外奖励：" + fmt.Sprintf("%f", sum31))
			fmt.Println("黑耀矿池：Goblin = " + fmt.Sprintf("%f", sum4) + ", 额外奖励：" + fmt.Sprintf("%f", sum41))
			fmt.Println("钴矿矿池：Goblin = " + fmt.Sprintf("%f", sum5) + ", 额外奖励：" + fmt.Sprintf("%f", sum51))
			fmt.Println("钛金矿池：Goblin = " + fmt.Sprintf("%f", sum6) + ", 额外奖励：" + fmt.Sprintf("%f", sum61))
			fmt.Println("所有账户余额：Nuls(" + fmt.Sprintf("%f", nulsSum) +
				"), Goblin(" + fmt.Sprintf("%f", gobSum) +
				"), Bla(" + fmt.Sprintf("%f", biSum) +
				"), Tun(" + fmt.Sprintf("%f", tuSum) +
				"), Pla(" + fmt.Sprintf("%f", plSum) +
				"), Obs(" + fmt.Sprintf("%f", obSum) +
				"), Cob(" + fmt.Sprintf("%f", coSum) +
				"), Tit(" + fmt.Sprintf("%f", tiSum) + ")")
		}
	},
}

func getSpace(length int) string {
	need := 130 - length
	space := ""
	for i := 0; i < need; i++ {
		space += " "
	}

	return space
}

func queryStatus(acc *account.Account) []map[string]interface{} {
	result, err := cfg.Sdk.SCMethodInvokeView(cfg.ChainId, "NULSd6Hgkfuo4hKAN5ysFRVP8gynyYD5uRGGE", "getRolesList", "", [][]string{{acc.Address}})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	listStr := result["result"].(string)
	list := []map[string]interface{}{}
	json.Unmarshal([]byte(listStr), &list)
	for _, item := range list {
		item["id"] = item["tokenId"]

		//todo where

	}
	return list
}

func earnedOf(where string, addr string) (string, string) {
	result, err := cfg.Sdk.SCMethodInvokeView(cfg.ChainId, "NULSd6HguK98JD4yFfjYkDTq8VXfPYDNeFMiL", cfg.EarnedOf, "", [][]string{{where}, {addr}})
	if err != nil {
		fmt.Println(err.Error())
		return "", ""
	}
	resultStr := result["result"].(string)
	list := []string{}
	json.Unmarshal([]byte(resultStr), &list)
	val1, _ := mathutils.StringToBigInt(list[0])
	val2, _ := mathutils.StringToBigInt(list[1])
	return decimal.NewFromBigInt(val1, -8).Round(2).String(),
		decimal.NewFromBigInt(val2, -8).Round(2).String()
}

func earnedOfMineLP(addr string) string {
	result, err := cfg.Sdk.SCMethodInvokeView(cfg.ChainId, cfg.MineLP2_Contract_Address, cfg.EarnedOfLP, "", [][]string{{addr}})
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	resultStr := result["result"].(string)
	val1, _ := mathutils.StringToBigInt(resultStr)
	return decimal.NewFromBigInt(val1, -8).Round(2).String()
}

func init() {
	rootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringVarP(&address, "address", "a", "", "查询指定账户的信息")
}
