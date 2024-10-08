package test

import (
	"database/sql"
	"fmt"
	"github.com/hiennq12/project_backend/cmd/struct_model"
	"github.com/hiennq12/project_backend/config"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestConnectDbPostgreSQL(t *testing.T) {
	dataConfig, err := config.ReadConfigDb("../config/connect_database.yaml")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error when read file config database. Detail: [%v]", err.Error()))
		return
	}

	if dataConfig == nil || dataConfig.Database == (struct_model.DatabaseConfig{}) { // check struct zero value
		log.Fatal(fmt.Sprintf("Data config is empty"))
		return
	}

	dataSource := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		dataConfig.Database.Host, dataConfig.Database.Port, dataConfig.Database.User,
		dataConfig.Database.Password, dataConfig.Database.Dbname, dataConfig.Database.Sslmode)
	connectPostgreSQL, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error when connect to postgreSQL. Detail: [%v]", err.Error()))
	}

	defer connectPostgreSQL.Close()

	err = connectPostgreSQL.Ping()
	if err != nil {
		log.Fatalf(fmt.Sprintf("Test ping db error. Detail: [%v]", err.Error()))
	}

	log.Printf("Connect database succcess!!")
}
