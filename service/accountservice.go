// @Title
// @Description
// @Author  Niels  2020/11/21
package service

import (
	"fmt"
	"github.com/niels1286/goblin-cmd/rlp"
	"github.com/syndtr/goleveldb/leveldb"
	"sort"
)

type BaseTable struct {
	dbOpener func() *leveldb.DB
}
type NULSAccount struct {
	Id           uint64
	Address      string
	PrikeyHex    string
	AddressBytes []byte
}

func openDB(path string) *leveldb.DB {
	//创建并打开数据库
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		panic(err)
	}
	return db
}

func GetAccountTable() *BaseTable {
	dbOpener := func() *leveldb.DB {
		return openDB(".database/accounts")
	}
	return &BaseTable{dbOpener: dbOpener}
}

func AddAccount(address []byte, acc *NULSAccount) error {

	acc.Id = uint64(len(QueryAccounts()) + 1)
	data, err := rlp.EncodeToBytes(acc)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	table := GetAccountTable()
	db := table.dbOpener()
	defer db.Close() //关闭数据库
	return db.Put(address, data, nil)
}

func RemoveAccount(address []byte) error {
	table := GetAccountTable()
	db := table.dbOpener()
	defer db.Close() //关闭数据库
	return db.Delete(address, nil)
}

func QueryAccounts() []*NULSAccount {
	table := GetAccountTable()
	db := table.dbOpener()
	defer db.Close()
	it := db.NewIterator(nil, nil)
	array := []*NULSAccount{}
	for it.Next() {
		val := it.Value()
		acc := &NULSAccount{}
		rlp.DecodeBytes(val, acc)
		array = append(array, acc)
	}
	sort.Slice(array, func(i, j int) bool {
		return array[i].Id < array[j].Id
	})
	return array
}
