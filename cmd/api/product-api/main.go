package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hiennq12/project_backend/cmd/api/product-api/app/handlers"
	"log"
)

type User struct {
	Name  string
	Email string
}

func main() {
	//r := mux.NewRouter()
	//r.HandleFunc("/api/hello", handlers.Hello).Methods(http.MethodGet)
	//log.Fatal(http.ListenAndServe(":8000", r))

	// create new app
	app := fiber.New()

	app.Use(cors.New())

	//get product
	app.Get("/api/products", handlers.GetProducts)

	// listen in port 8000
	log.Fatalf("Error: %s", app.Listen("localhost:8000"))
}

//func structToMap(objStruct interface{}) map[string]interface{} {
//	res := make(map[string]interface{})
//	value := reflect.ValueOf(objStruct)
//	if value.Kind() == reflect.Ptr {
//		value = value.Elem()
//	}
//	typeF := value.Type()
//	for i := 0; i < value.NumField(); i++ {
//		fieldName := strings.ToLower(typeF.Field(i).Name)
//		fieldValueKind := value.Field(i).Kind()
//		var fieldValue interface{}
//		if fieldValueKind == reflect.Struct {
//			fieldValue = structToMap(value.Field(i).Interface())
//		} else {
//			fieldValue = value.Field(i).Interface()
//		}
//		res[fieldName] = fieldValue
//	}
//
//	return res
//}
