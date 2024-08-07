package dms

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"log"
)

type TestRequest struct {
	Id   int64
	Name string
}

type TestResponse struct {
	LastInsertId int64
	RowEffect    int64
}

func InsertDataToTestTable(req *TestRequest) (*TestResponse, error) {
	connect, err := ConnectDbPostgreSQL()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func(connect *sql.DB) {
		err = connect.Close()
		if err != nil {
			log.Printf("Error when close connect postgresSQL")
		}
	}(connect)

	if req == nil || len(req.Name) < 1 {
		return nil, errors.New("require data insert")
	}

	query, args, err := squirrel.Insert("test").Columns("name_book").
		Values(req.Name).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//query = `INSERT INTO test (name_book) VALUES ($1)`
	//response, err := connect.Exec(query, req.Name)

	fmt.Println("Query:", query) // In câu lệnh SQL để kiểm tra
	fmt.Println("Args:", args)   // In đối số để kiểm tra

	response, err := connect.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	lastInsertId, _ := response.LastInsertId()
	rowEffect, _ := response.RowsAffected()
	return &TestResponse{
		LastInsertId: lastInsertId,
		RowEffect:    rowEffect,
	}, nil
}
