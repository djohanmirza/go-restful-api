package controller

import (
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/service"
	"github.com/gofiber/fiber/v2"
)

type EmployeeControllerImpl struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		EmployeeService: employeeService,
	}
}

func (controller *EmployeeControllerImpl) Create(c *fiber.Ctx) error {
	employeeCreateRequest := new(web.EmployeeCreateRequest)
	if err := c.BodyParser(employeeCreateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	employeeResponse, err := controller.EmployeeService.Create(c.Context(), *employeeCreateRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   employeeResponse,
	})
}

func (controller *EmployeeControllerImpl) Update(c *fiber.Ctx) error {
	employeeUpdateRequest := new(web.EmployeeUpdateRequest)
	if err := c.BodyParser(employeeUpdateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	employeeResponse, err := controller.EmployeeService.Update(c.Context(), *employeeUpdateRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   employeeResponse,
	})
}

func (controller *EmployeeControllerImpl) Delete(c *fiber.Ctx) error {
	employeeId := c.Params("employeeId")

	err := controller.EmployeeService.Delete(c.Context(), employeeId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "Deleted Successfully",
	})
}

func (controller *EmployeeControllerImpl) FindById(c *fiber.Ctx) error {
	employeeId := c.Params("employeeId")

	employeeResponse, err := controller.EmployeeService.FindById(c.Context(), employeeId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   employeeResponse,
	})
}

func (controller *EmployeeControllerImpl) FindAll(c *fiber.Ctx) error {
	employeeResponses, err := controller.EmployeeService.FindAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   employeeResponses,
	})
}
