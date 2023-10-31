package helper

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	global "changian.com/jubo/startup"
)

// 初始化 - DB
func InitDB() *sql.DB {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		global.Config.DbUrl,
		global.Config.DbPort,
		global.Config.DbUserName,
		global.Config.DbPassword,
		global.Config.DbName)

	fmt.Println("postgresql connection string : ", connectionString)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)

		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)

		panic(err.Error())
	}

	// 設定空閒時間
	db.SetConnMaxLifetime(10 * time.Second)

	db.SetMaxOpenConns(10)

	db.SetMaxIdleConns(10)

	return db
}
