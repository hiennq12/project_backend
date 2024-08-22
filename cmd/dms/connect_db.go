package dms

import (
	"database/sql"
	"fmt"
	"github.com/hiennq12/project_backend/cmd/struct_model"
	"github.com/hiennq12/project_backend/config"
	_ "github.com/lib/pq"
	"log"
)

var connectDB *sql.DB

func ConnectDbPostgreSQL() (*sql.DB, error) {
	dataConfig, err := config.ReadConfigDb("../../config/connect_database.yaml")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error when read file config database. Detail: [%v]", err.Error()))
		return nil, err
	}

	if dataConfig == nil || dataConfig.Database == (struct_model.DatabaseConfig{}) { // check struct zero value
		log.Fatal(fmt.Sprintf("Data config is empty"))
		return nil, err
	}

	dataSource := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		dataConfig.Database.Host, dataConfig.Database.Port, dataConfig.Database.User,
		dataConfig.Database.Password, dataConfig.Database.Dbname, dataConfig.Database.Sslmode)
	connectPostgreSQL, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error when connect to postgreSQL. Detail: [%v]", err.Error()))
		return nil, err
	}
	//defer connectPostgreSQL.Close()

	log.Printf("Connect database succcess!!")
	connectDB = connectPostgreSQL
	return connectPostgreSQL, nil
	//err = connectPostgreSQL.Ping()
	//if err != nil {
	//	log.Fatalf(fmt.Sprintf("Test ping db error. Detail: [%v]", err.Error()))
	//}

}
