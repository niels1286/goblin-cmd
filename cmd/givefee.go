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
	"github.com/niels1286/nuls-go-sdk"
	"github.com/niels1286/nuls-go-sdk/account"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/spf13/cobra"
	"math/big"
	"time"
)

// givefeeCmd represents the givefee command
var givefeeCmd = &cobra.Command{
	Use:   "givefee",
	Short: "分发手续费",
	Long:  `从指定的用户，分发NULS到其他用户，保证每个用户至少有5个NULS`,
	Run: func(cmd *cobra.Command, args []string) {
		list := service.QueryAccounts()
		var mainAccount *service.NULSAccount
		for _, nulsAcc := range list {
			if address == nulsAcc.Address {
				mainAccount = nulsAcc
				break
			}
		}
		for _, nulsAcc := range list {
			if address == nulsAcc.Address {
				continue
			}
			accountInfo, _ := cfg.Sdk.GetBalance(nulsAcc.Address, cfg.ChainId, cfg.AssetsId)
			if accountInfo.Balance.Cmp(big.NewInt(100000000)) > 0 {
				continue
			}
			val := big.NewInt(300000000)
			transfer(cfg.Sdk, mainAccount, nulsAcc, val.Sub(val, accountInfo.Balance))
		}
	},
}

func transfer(sdk *nuls.NulsSdk, from *service.NULSAccount, to *service.NULSAccount, amount *big.Int) {
	accountInfo, _ := sdk.GetBalance(from.Address, cfg.ChainId, cfg.AssetsId)
	if accountInfo == nil {
		fmt.Println("接口请求失败！")
		return
	}
	if accountInfo.Balance.Cmp(big.NewInt(100000000+amount.Int64())) < 0 {
		fmt.Println("主账户余额不足！")
		return
	}
	tx := txprotocal.Transaction{
		TxType:   txprotocal.TX_TYPE_TRANSFER,
		Time:     uint32(time.Now().Unix()),
		Remark:   nil,
		Extend:   nil,
		CoinData: nil,
		SignData: nil,
	}
	fromAmount := big.NewInt(100000)
	fromAmount = fromAmount.Add(fromAmount, amount)
	coinData := txprotocal.CoinData{
		Froms: []txprotocal.CoinFrom{{
			Coin: txprotocal.Coin{
				Address:       from.AddressBytes,
				AssetsChainId: cfg.ChainId,
				AssetsId:      cfg.AssetsId,
				Amount:        fromAmount,
			},
			Nonce:  accountInfo.Nonce,
			Locked: 0,
		}},
		Tos: []txprotocal.CoinTo{{
			Coin: txprotocal.Coin{
				Address:       to.AddressBytes,
				AssetsChainId: cfg.ChainId,
				AssetsId:      cfg.AssetsId,
				Amount:        amount,
			},
			LockValue: 0,
		}},
	}
	realAcc, _ := account.GetAccountFromPrkey(from.PrikeyHex, cfg.ChainId, cfg.AddressPrefix)
	tx.CoinData, _ = coinData.Serialize()
	hash, _ := tx.GetHash().Serialize()
	signValue, _ := realAcc.Sign(hash)

	txSign := txprotocal.CommonSignData{
		Signatures: []txprotocal.P2PHKSignature{{
			SignValue: signValue,
			PublicKey: realAcc.GetPubKeyBytes(true),
		}},
	}
	tx.SignData, _ = txSign.Serialize()
	resultBytes, _ := tx.Serialize()
	bcdResult, err := sdk.BroadcastTx(resultBytes)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(realAcc.Address + " ： " + bcdResult)
	}
}

func init() {
	rootCmd.AddCommand(givefeeCmd)
	givefeeCmd.Flags().StringVarP(&address, "address", "a", "", "指定分发主账户")
	givefeeCmd.MarkFlagRequired("address")
}
