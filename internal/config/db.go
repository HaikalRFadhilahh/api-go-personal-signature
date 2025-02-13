package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/pkg"
	_ "github.com/go-sql-driver/mysql"
)

/*
	Struct Database Credentials
*/

type databaseConnection struct {
	databaseConnectionString string
}

/*
Factory Init Database COnnection
*/

func NewDB() *databaseConnection {
	return &databaseConnection{
		databaseConnectionString: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", pkg.GetEnvValue("DB_USERNAME", "root"), pkg.GetEnvValue("DB_PASSWORD", ""), pkg.GetEnvValue("DB_HOST", "127.0.0.1"), pkg.GetEnvValue("DB_PORT", "3306"), pkg.GetEnvValue("DB_NAME", "")),
	}

}

/*
	Function Inti Database Connection
*/

func (d *databaseConnection) InitDB() *sql.DB {
	// Create Database Instance
	db, err := sql.Open("mysql", d.databaseConnectionString)
	if err != nil {
		log.Fatalln("Database Init Connection Error, Error :", err.Error())
	}

	// Check Connection Ping
	if err := db.Ping(); err != nil {
		log.Fatalln("Database Connection Ping Error, Error :", err.Error())
	}

	return db
}

func CloseConnection(db *sql.DB) {
	fmt.Println("Shutdown Database Connection...")
	db.Close()
}
