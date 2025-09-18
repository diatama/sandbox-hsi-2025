package handlers

import (
	"perpustakaan/pkg/buku"
	"perpustakaan/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func LoadAllBuku(service *buku.BukuService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bukus, err := service.GetAllBuku()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to load books",
			})
		}
		return c.JSON(bukus)
	}
}

func BuatBukuBaru(service *buku.BukuService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var buku entities.Buku
		if err := c.BodyParser(&buku); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		err := service.CreateBuku(buku)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create book",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(buku)
	}
}
