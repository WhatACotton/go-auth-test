package database

import (
	"log"
	"unify/internal/models"
)

func PostItemInfo(Item models.ItemInfo) {
	db := ConnectSQL()
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("INSERT INTO item VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		//log.Fatal(err)
		panic(err.Error())
	}
	defer ins.Close()

	// SQLの実行
	_, err = ins.Exec(
		Item.InfoId,
		Item.Price,
		Item.Name,
		Item.Stonesize,
		Item.Minlength,
		Item.Maxlength,
		Item.Decsription,
		Item.Keyword,
	)
	if err != nil {
		//log.Fatal(err)
		panic(err.Error())
	}
}

func GetItemInfo(InfoId string) (returnmodels models.ItemInfo) {
	// データベースのハンドルを取得する
	db := ConnectSQL()
	defer db.Close()

	// SQLの実行
	rows, err := db.Query("SELECT * FROM item WHERE InfoId = ?", InfoId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var resultItem models.ItemInfo
	// SQLの実行
	for rows.Next() {
		err := rows.Scan(
			&resultItem.InfoId,
			&resultItem.Price,
			&resultItem.Name,
			&resultItem.Stonesize,
			&resultItem.Minlength,
			&resultItem.Maxlength,
			&resultItem.Decsription,
			&resultItem.Keyword,
		)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultItem
}

func UpdateItemInfo(Item models.ItemInfo) {
	db := ConnectSQL()
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("UPDATE item SET price = ?, name = ?, stonesize = ?, minlength = ?, maxlength = ?, decsription = ?, keyword = ? WHERE InfoId = ?")
	if err != nil {
		//log.Fatal(err)
		panic(err.Error())
	}
	defer ins.Close()

	// SQLの実行
	_, err = ins.Exec(
		Item.Price,
		Item.Name,
		Item.Stonesize,
		Item.Minlength,
		Item.Maxlength,
		Item.Decsription,
		Item.Keyword,
		Item.InfoId,
	)
	if err != nil {
		//log.Fatal(err)
		panic(err.Error())
	}
}

func GetPrice(InfoId string) (price int) {
	// データベースのハンドルを取得する
	db := ConnectSQL()
	defer db.Close()

	// SQLの実行
	rows, err := db.Query("SELECT price FROM item WHERE InfoId = ?", InfoId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var resultPrice int
	// SQLの実行
	for rows.Next() {
		err := rows.Scan(&resultPrice)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultPrice
}
