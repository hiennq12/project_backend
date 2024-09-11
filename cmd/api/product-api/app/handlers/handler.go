package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hiennq12/project_backend/cmd/dms"
	"github.com/hiennq12/project_backend/cmd/struct_model"
	"github.com/hiennq12/project_backend/utils/log"
)

func GetProducts(ctx *fiber.Ctx) error {
	req := &struct_model.ProductsRequest{}
	if err := ctx.QueryParser(req); err != nil {
		log.LogErrorWithLine(err)
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Error detail: %v", err.Error()))
	}
	products, err := dms.GetProducts(ctx.Context(), req)
	if err != nil {
		log.LogErrorWithLine(err)
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error detail: %v", err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(products)
}
