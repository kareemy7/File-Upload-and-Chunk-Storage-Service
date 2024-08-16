package initializers

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

var DB *leveldb.DB

func ConnectToDB() {
	var err error
	DB, err = leveldb.OpenFile("./data", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CloseDB() {
	DB.Close()
}
