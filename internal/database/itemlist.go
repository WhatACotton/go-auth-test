package database

import (
	"log"
	"unify/internal/models"
)

const (
	Expired   = 0
	Available = 1
	Invaild   = 2
)

func PostItemList(newItem models.ItemList) (Itemlist []models.ItemList) {
	// データベースのハンドルを取得する
	db := ConnectSQL()
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("INSERT INTO itemlist VALUES(?,?,?)")
	if err != nil {
		//log.Fatal(err)
		panic(err.Error())

	}
	defer ins.Close()

	// SQLの実行
	_, err = ins.Exec(
		newItem.ItemId,
		newItem.InfoId,
		newItem.StatusCode,
	)
	if err != nil {
		//log.Fatal(err)
		panic(err.Error())

	}
	return GetItemList()
}

func GetItemList() (Itemlist []models.ItemList) {
	// データベースのハンドルを取得する
	db := ConnectSQL()
	defer db.Close()

	// SQLの実行
	rows, err := db.Query("SELECT * FROM itemlist ")
	if err != nil {

		log.Fatal(err)
	}
	defer rows.Close()
	var resultItem models.ItemList
	var resultItemList []models.ItemList
	// SQLの実行
	for rows.Next() {
		err := rows.Scan(&resultItem.ItemId, &resultItem.InfoId, &resultItem.StatusCode)
		if err != nil {
			panic(err.Error())
		}
		resultItemList = append(resultItemList, resultItem)
	}
	return resultItemList
}

func UpdateItemList(ItemId string, InfoId string) {
	ItemList := new(models.ItemList)
	ItemList.InfoId = InfoId
	ItemList.ItemId = "default"
	ItemList.StatusCode = Expired
	PostItemList(*ItemList)
	// データベースのハンドルを取得する
	db := ConnectSQL()
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("UPDATE itemlist SET InfoId = ? WHERE ItemId = ?")
	if err != nil {
		panic(err.Error())
	}
	defer ins.Close()

	// SQLの実行
	_, err = ins.Exec(InfoId, ItemId)
	if err != nil {
		panic(err.Error())
	}
}

func ChangeItemStatus(ItemId string, StatusCode int) {
	// データベースのハンドルを取得する
	db := ConnectSQL()
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("UPDATE itemlist SET StatusCode = ? WHERE ItemId = ?")
	if err != nil {
		panic(err.Error())
	}
	defer ins.Close()

	// SQLの実行
	_, err = ins.Exec(StatusCode, ItemId)
	if err != nil {
		panic(err.Error())
	}
}

// 基本的にItemInfoは削除しない方針でいくので使わない。
func DeleteItemList(id string) (Itemlist []models.ItemList) {
	// データベースのハンドルを取得する
	db := ConnectSQL()
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("DELETE FROM itemlist WHERE InfoId = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// SQLの実行
	_, err = ins.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	return GetItemList()
}
