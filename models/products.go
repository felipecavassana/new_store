package models

import "github.com/felipecavassana/new_store/db"

type Item struct {
	Id          int
	Name        string
	Description string
	Amount      float64
	Quantity    int
}

func SearchAllItems() []Item {
	db := db.ConnectDB()

	allItemsFromDB, err := db.Query("select * from new_store.products")
	if err != nil {
		panic(err.Error())
	}

	p := Item{}
	itemList := []Item{}

	for allItemsFromDB.Next() {
		var id, quantity int
		var name, description string
		var amount float64

		err = allItemsFromDB.Scan(&id, &name, &description, &amount, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Amount = amount
		p.Quantity = quantity

		itemList = append(itemList, p)
	}

	defer db.Close()
	return itemList
}

func NewItem(name, description string, amount float64, quantity int) {
	db := db.ConnectDB()

	insertIntoDB, err := db.Prepare("insert into new_store.products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertIntoDB.Exec(name, description, amount, quantity)
	defer db.Close()
}

func DeleteItem(id string) {
	db := db.ConnectDB()

	deleteItem, err := db.Prepare("delete from new_store.products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteItem.Exec(id)
	defer db.Close()
}

func EditItem(id string) Item {
	db := db.ConnectDB()

	selectItem, err := db.Query("select * from new_store.products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	itemToUpdate := Item{}

	for selectItem.Next() {
		var id, quantity int
		var name, description string
		var amount float64

		err = selectItem.Scan(&id, &name, &description, &amount, &quantity)
		if err != nil {
			panic(err.Error())
		}
		itemToUpdate.Id = id
		itemToUpdate.Name = name
		itemToUpdate.Description = description
		itemToUpdate.Amount = amount
		itemToUpdate.Quantity = quantity
	}

	defer db.Close()
	return itemToUpdate
}

func UpdateItem(id int, name, description string, amount float64, quantity int) {
	db := db.ConnectDB()

	UpdateItem, err := db.Prepare("update new_store.products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	UpdateItem.Exec(name, description, amount, quantity, id)
	defer db.Close()
}
