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
	"reflect"
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

func InsertProducts(ctx context.Context, req []*struct_model.InsertProductsRequest) (*ResponeInsert, error) {
	_, err := ConnectDbPostgreSQL()
	if err != nil {
		return nil, err
	}

	tableName := "products"
	ignoreColumns := []string{"id"}
	columns := dms_utils.GetQueryColumnList(ignoreColumns, &struct_model.Product{})

	qb := squirrel.Insert(tableName).Columns(columns).PlaceholderFormat(squirrel.Dollar)

	for _, val := range req {
		argsVal := make([]interface{}, 0)

		v := reflect.ValueOf(val) // lay duoc con tro
		v = reflect.Indirect(v)   // lay gia tri struct tu con tro
		if v.Kind() != reflect.Struct {
			return nil, errors.New("not struct")
		}
		for i := 0; i < v.NumField(); i++ {
			// Get the field by index
			field := v.Field(i)
			// Get the field name
			fieldName := v.Type().Field(i).Name
			// Print the field name and value
			fmt.Printf("Field %d: %s = %v\n", i, fieldName, field.Interface())
			argsVal = append(argsVal, field.Interface())
		}
		qb = qb.Values(argsVal...)
	}
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	response, err := connectDB.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	lastInsertId, _ := response.LastInsertId()
	rowEffect, _ := response.RowsAffected()
	return &ResponeInsert{
		LastInsertId: int(lastInsertId),
		RowEffect:    int(rowEffect),
	}, nil
}

type ResponeInsert struct {
	LastInsertId int
	RowEffect    int
}

type ProductsResponse struct {
}

func GetProducts(ctx context.Context, req *struct_model.ProductsRequest) ([]struct_model.Product, error) {
	_, err := ConnectDbPostgreSQL()
	if err != nil {
		log.Fatal("error: ", err.Error())
		return nil, err
	}

	limit := uint64(10)
	if req.Limit > 0 {
		limit = req.Limit
	}

	qb := squirrel.Select("*").From("products").Limit(limit).
		PlaceholderFormat(squirrel.Dollar)

	if req.ProductId > 0 {
		qb = qb.Where(squirrel.Eq{"id": req.ProductId})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := connectDB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	products := make([]struct_model.Product, 0)
	if err = dms_utils.ScanRowsToStruct(rows, &products); err != nil {
		return nil, fmt.Errorf("failed to map rows to struct: %w", err)
	}

	return products, nil
}
