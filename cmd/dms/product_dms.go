package dms

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/hiennq12/project_backend/cmd/struct_model"
	"github.com/hiennq12/project_backend/utils/dms-utils"
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

func InsertProduct() {

}

type ProductsRequest struct {
	ProductId  int64   `json:"product_id"`
	ProductIds []int64 `json:"product_ids"`
}

type ProductsResponse struct {
}

func GetProducts(ctx context.Context, req *ProductsRequest) ([]struct_model.Product, error) {
	_, err := ConnectDbPostgreSQL()
	if err != nil {
		log.Fatal("error: ", err.Error())
		return nil, err
	}

	query, args, err := squirrel.Select("*").From("products").
		Where(squirrel.Eq{"id": req.ProductId}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := connectDB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := make([]struct_model.Product, 0)
	if err = dms_utils.ScanRowsToStruct(rows, &products); err != nil {
		return nil, fmt.Errorf("failed to map rows to struct: %w", err)
	}

	return products, nil
}
