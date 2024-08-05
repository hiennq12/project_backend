package config

import (
	"github.com/hiennq12/project_backend/cmd/struct_model"
	"gopkg.in/yaml.v2"
	"os"
)

func ReadConfigDb(configFile string) (*struct_model.ConfigDb, error) {
	dataFile, err := os.ReadFile(configFile)
	if err != nil {
		//log.Fatal(fmt.Sprintf("Error when read file config database. Detail: [%v]"), err.Error())
		return nil, err
	}

	dbConfig := &struct_model.ConfigDb{}
	err = yaml.Unmarshal(dataFile, dbConfig)
	if err != nil {
		return nil, err
	}

	return dbConfig, nil
}
