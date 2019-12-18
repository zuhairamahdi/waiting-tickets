package infrastructure

import (
	"fmt"
	"path"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"os"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func OpenDbConnection() *gorm.DB {
	dialect := os.Getenv("DB_DIALECT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	var db *gorm.DB
	var err error
	if dialect == "sqlite3" {
		db, err = gorm.Open("sqlite3", path.Join(".", "app.db"))
	} else {
		// db, err := gorm.Open("mysql", "root:root@localhost/igo_api_shop_gonc?charset=utf8")
		databaseUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable ", host, username, password, dbName)
		db, err = gorm.Open(dialect, databaseUrl)
	}

	if err != nil {
		fmt.Println("db err: ", err)
		os.Exit(-1)
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}

//RemoveDb Delete the database after running testing cases.
func RemoveDb(db *gorm.DB) error {
	db.Close()
	err := os.Remove(path.Join(".", "app.db"))
	return err
}

//GetDb Using this function to get a connection, you can create your connection pool here.
func GetDb() *gorm.DB {
	return DB
}
