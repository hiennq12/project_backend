package dms

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestGetProducts(t *testing.T) {
	req := &ProductsRequest{
		ProductId: 5,
		//ProductIds: nil,
	}

	response, err := GetProducts(context.Background(), req)
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	fmt.Println("====Data: ", response)
}
