package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"path/filepath"
	"product/internal/model/web"
	"product/internal/services"
	"strconv"
	"strings"
	"time"
)

type ProductControllerImpl struct {
	Service   services.ProductService
	Validator *validator.Validate
}

func NewProductController(service services.ProductService, validator *validator.Validate) ProductController {
	return &ProductControllerImpl{Service: service, Validator: validator}
}

func (controller ProductControllerImpl) Create(ctx *fiber.Ctx) error {
	productRequest := web.ProductCreateRequest{}
	if err := ctx.BodyParser(&productRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to parse JSON",
		})
	}

	imgProd, err := ctx.FormFile("image")
	if err != nil {
		panic(err)
	}

	fileExt := filepath.Ext(imgProd.Filename)

	fileName := strings.TrimSuffix(imgProd.Filename, fileExt)

	currentTime := time.Now().Format("20060102150405")
	fileNameWithTime := fmt.Sprintf("%s-%s", fileName, currentTime)
	newFileName := fmt.Sprintf("%s%s", fileNameWithTime, fileExt)

	imgProd.Filename = newFileName
	productRequest.Image = imgProd.Filename
	err = ctx.SaveFile(imgProd, fmt.Sprintf("./public/product/%s", imgProd.Filename))

	if err := controller.Validator.Struct(productRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	productResponse, err := controller.Service.Create(productRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "product has been created",
		Data:   productResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

// clear
func (controller ProductControllerImpl) Update(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	productRequest := web.ProductUpdateRequest{}
	productRequest.Id = id

	if err := ctx.BodyParser(&productRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to parse JSON",
		})
	}

	if err := controller.Validator.Struct(productRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	productResponse, err := controller.Service.Update(productRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "SERVER INTERNAL ERROR",
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "product has been updated",
		Data:   productResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// clear
func (controller ProductControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	productFounded, err := controller.Service.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "data not found",
		})
	}

	controller.Service.Delete(productFounded.Id)
	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "product has been deleted",
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// clear
func (controller ProductControllerImpl) FindById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	productFounded, err := controller.Service.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "data not found",
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "product has been found",
		Data:   productFounded,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// clear
func (controller ProductControllerImpl) FindAll(ctx *fiber.Ctx) error {
	products, err := controller.Service.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "SERVER INTERNAL ERROR",
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all product",
		Data:   products,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
