package handlers

import (
	"perpustakaan/pkg/entities"
	"perpustakaan/pkg/peminjaman"

	"github.com/gofiber/fiber/v2"
)

func LoadAllPeminjaman(service *peminjaman.PeminjamanService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		peminjaman, err := service.GetAllPeminjaman()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to load peminjamans",
			})
		}
		return c.JSON(peminjaman)
	}
}

func BuatPeminjamanBaru(service *peminjaman.PeminjamanService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var peminjaman entities.Peminjaman
		if err := c.BodyParser(&peminjaman); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid Request Body",
			})
		}

		// Append value of TanggalPengembalian (+7 Hari)
		tanggalPengembalian := peminjaman.TanggalPeminjaman.AddDate(0, 0, 7)
		peminjaman.TanggalPengembalian = &tanggalPengembalian

		err := service.CreatePeminjaman(peminjaman)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create peminjaman",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(peminjaman)
	}
}
