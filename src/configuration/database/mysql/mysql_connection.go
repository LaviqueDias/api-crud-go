package mysqldb

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
) 

var db *sql.DB

func Connection() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(localhost:3307)/apicrudgo")
		if err != nil {
			panic(err)
		}
	}
	return db
}

func CloseConnection() {
	if db != nil {
		db.Close()
	}
}