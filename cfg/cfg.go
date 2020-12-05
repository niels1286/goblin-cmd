// @Title
// @Description
// @Author  Niels  2020/11/21
package cfg

import "github.com/niels1286/nuls-go-sdk"

const ChainId = uint16(1)
const AssetsId = uint16(1)
const AddressPrefix = "NULS"

var Sdk = nuls.NewNulsSdk("https://api.nuls.io/jsonrpc/", "https://public1.nuls.io/", ChainId)

//nft合约地址
const NFT_Contract_Address = "NULSd6HgvBGqSQBr49QmB9BJia4RnzsAWpjtE"

//矿池合约地址
//const Mine_Contract_Address = "NULSd6HguK98JD4yFfjYkDTq8VXfPYDNeFMiL"
const Mine_Contract_Address = "NULSd6HgzwFcYoCRXvK9ddcHnNzbff8CYdonW"
const MineLP2_Contract_Address = "NULSd6Hgr5k9ic8kxr9UnPmiCqaHxcypimHex"
const NRC20_Contract_Address_BlackIron = "NULSd6HgnScefpS1jGFvJZeNPnFRtAebwVpJr"
const NRC20_Contract_Address_Tungsten = "NULSd6HgngWZE3u8WRe4QxCVcUiyGtDP7cGEU"
const NRC20_Contract_Address_Goblin = "NULSd6HgwJmD4SC1NAJXu8tC6NKsWs99P2jpw"
const NRC20_Contract_Address_Platinum = "NULSd6Hgocjz53UsmRkGJuQiXz993S3vosu81"
const NRC20_Contract_Address_Obsidian = "NULSd6Hgyb8Tfy9vQcf1Z4QGqdRFkzfYR1M5o"
const NRC20_Contract_Address_Titanium = "NULSd6HgwpusTTJMofoykdQt8qpWdQKCGKMSK"
const NRC20_Contract_Address_Cobalt = "NULSd6HgkFgtEhbvtFo7pUVRbD5GdHNGkgFUj"

//黑铁
const BlackIron = "BlackIron"

//钨矿
const Tungsten = "Tungsten"
const Platinum = "Platinum"
const Obsidian = "Obsidian"
const Titanium = "Titanium"
const Cobalt = "Cobalt"

//收获
const ClaimEarned = "claimEarned"

//升级
const LevelUp = "levelUp"

//遣回
const Remand = "remand"

//查询收获数量
const EarnedOf = "earnedOf"
const EarnedOfLP = "earned"

//派出
const EnterMineCave = "enterMineCave"
