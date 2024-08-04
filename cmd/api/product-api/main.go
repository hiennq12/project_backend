package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"reflect"
	"strings"
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

	// send string, method get, endpoint /hello
	app.Get("/", func(ctx *fiber.Ctx) error {
		err := ctx.SendString("Hello Home page!!!")
		return err
	})

	// send string, method get, endpoint /hello
	app.Get("/hello", func(ctx *fiber.Ctx) error {
		err := ctx.SendString("Hello Everyone!!!")
		return err
	})

	app.Get("/user/1", func(ctx *fiber.Ctx) error {
		user := User{
			Name:  "Hien Ahihi",
			Email: "simsonahihi@gmail.com",
		}

		userInfo := structToMap(user)

		userJson, err := json.Marshal(userInfo)
		if err != nil {
			log.Fatalf("Error when marshal map to string, detail [%v]", err.Error())
		}
		err = ctx.SendString(string(userJson))
		return err
	})

	// listen in port 8000
	log.Fatalf("Error: %s", app.Listen("localhost:8000"))
}

func structToMap(objStruct interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	value := reflect.ValueOf(objStruct)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	typeF := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fieldName := strings.ToLower(typeF.Field(i).Name)
		fieldValueKind := value.Field(i).Kind()
		var fieldValue interface{}
		if fieldValueKind == reflect.Struct {
			fieldValue = structToMap(value.Field(i).Interface())
		} else {
			fieldValue = value.Field(i).Interface()
		}
		res[fieldName] = fieldValue
	}

	return res
}
