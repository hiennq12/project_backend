package dms

import (
	"fmt"
	"log"
	"testing"
)

func TestInsertData(t *testing.T) {
	req := &TestRequest{
		Name: "4MK+1",
	}
	res, err := InsertDataToTestTable(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DATA RESPONSE: ", res)
}
