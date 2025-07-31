package mysqldb

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
) 

var db *sql.DB

func Connection() (*sql.DB, error) {
	
	if db == nil {
		var err error
		db, err = sql.Open("mysql", GetMySQLDSN())
		if err != nil {
			return nil, err
		}
		if _, err := db.Exec("SET time_zone = '-03:00'"); err != nil {
            // só um aviso, não trava a aplicação
            logger.Error("warning: não foi possível ajustar time_zone da sessão: %v", err, zap.String("journey", "connectionToDB"),)
        }
	}
	return db, nil
}

func CloseConnection() {
	if db != nil {
		db.Close()
	}
}

func GetMySQLDSN() string {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=America%%2FSao_Paulo",
		user, pass, host, port, dbname)
}