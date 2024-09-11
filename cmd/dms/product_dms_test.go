package dms

import (
	"context"
	"fmt"
	"github.com/hiennq12/project_backend/cmd/struct_model"
	"log"
	"testing"
	"time"
)

func TestGetProducts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &struct_model.ProductsRequest{
		//ProductId: 5,
		//ProductIds: nil,
	}

	response, err := GetProducts(ctx, req)
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	fmt.Println("====Data: ", response)
}

func TestInsertProducts(t *testing.T) {
	req := make([]*struct_model.InsertProductsRequest, 0)
	req = append(req, &struct_model.InsertProductsRequest{
		UserId:        1,
		CategoryId:    1,
		ProductName:   "Nguoi ru ngu1",
		Description:   "tam bo qua1",
		Price:         200000,
		Condition:     "con moi1",
		Location:      "Ha Noi, Viet Nam",
		StockQuantity: 1,
		Weight:        0.3,
		Dimensions:    "13x20cm",
		SKU:           "DonatoCarrisi_NguoiRuNgu1",
		Brand:         "Sach tieng viet",
		Warranty:      "Khong co",
		IsNegotiable:  true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		IsActive:      true,
	}, &struct_model.InsertProductsRequest{
		UserId:        1,
		CategoryId:    1,
		ProductName:   "Nguoi ru ngu2",
		Description:   "tam bo qua2",
		Price:         200000,
		Condition:     "con moi1",
		Location:      "Ha Noi, Viet Nam",
		StockQuantity: 1,
		Weight:        0.3,
		Dimensions:    "13x20cm",
		SKU:           "DonatoCarrisi_NguoiRuNgu2",
		Brand:         "Sach tieng viet",
		Warranty:      "Khong co",
		IsNegotiable:  true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		IsActive:      true,
	})

	response, err := InsertProducts(context.Background(), req)
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	// ignore column : product_id
	// bat buoc phai co: user_id, category_id, product_name, price, condition, stock_quantity, sku
	//product_id,user_id,category_id,product_name,description,price,condition,location,stock_quantity,weight,dimensions,sku,brand,warranty,is_negotiable,created_at,updated_at,is_active
	fmt.Println("====Data: ", response)
}
