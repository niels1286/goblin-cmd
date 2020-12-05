// @Title
// @Description
// @Author  Niels  2020/11/17
package txs

import (
	"fmt"
	"github.com/niels1286/goblin-cmd/cfg"
	"github.com/niels1286/nuls-go-sdk"
	"github.com/niels1286/nuls-go-sdk/account"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/niels1286/nuls-go-sdk/tx/txdata"
	"math/big"
	"time"
)

func SimpleCallContractTx(sdk *nuls.NulsSdk, realAcc *account.Account, contractAddress string, methodName string, args [][]string) {
	accountInfo, _ := sdk.GetBalance(realAcc.Address, cfg.ChainId, cfg.AssetsId)
	if accountInfo == nil {
		fmt.Println("接口请求失败！")
		return
	}
	tx := txprotocal.Transaction{
		TxType:   txprotocal.TX_TYPE_CALL_CONTRACT,
		Time:     uint32(time.Now().Unix()),
		Remark:   nil,
		Extend:   nil,
		CoinData: nil,
		SignData: nil,
	}

	ccData := txdata.CallContract{
		Sender:          realAcc.AddressBytes,
		ContractAddress: account.AddressStrToBytes(contractAddress),
		Value:           big.NewInt(0),
		GasLimit:        200000,
		Price:           25,
		MethodName:      methodName,
		MethodDesc:      "",
		ArgsCount:       uint8(len(args)),
		Args:            args,
	}
	tx.Extend, _ = ccData.Serialize()
	coinData := txprotocal.CoinData{
		Froms: []txprotocal.CoinFrom{{
			Coin: txprotocal.Coin{
				Address:       realAcc.AddressBytes,
				AssetsChainId: cfg.ChainId,
				AssetsId:      cfg.AssetsId,
				Amount:        big.NewInt(5100000),
			},
			Nonce:  accountInfo.Nonce,
			Locked: 0,
		}},
		Tos: nil,
	}
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
