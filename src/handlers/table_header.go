package handlers

import (
	"github.com/gofiber/fiber/v2"
	"min-selhoz-backend/src/forms"
	"min-selhoz-backend/src/services"
)

type TableHeaderInterface interface {
	ListByTableId() fiber.Handler
	Create() fiber.Handler
}

type TableHeaderHandler struct {
	s services.TableHeaderInterface
}

func (h TableHeaderHandler) ListByTableId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tableID := ctx.Params("table_id")

		tableHeaderResp, err := h.s.ListByTableId(tableID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(fiber.StatusOK).JSON(tableHeaderResp)
	}
}

func (h TableHeaderHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		form := new(forms.TableHeader)
		err := ctx.BodyParser(&form)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		id, err := h.s.Create(form)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})

	}
}

func NewTableHeaderHandler(s services.TableHeaderInterface) TableHeaderInterface {
	return &TableHeaderHandler{s}
}
