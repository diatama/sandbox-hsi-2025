package routes

import (
	"perpustakaan/api/handlers"
	"perpustakaan/pkg/peminjaman"

	"github.com/gofiber/fiber/v2"
)

func PeminjamanRoutes(api fiber.Router, peminjamanService *peminjaman.PeminjamanService) {
	api.Get("/peminjaman", handlers.LoadAllPeminjaman(peminjamanService))
	api.Post("/peminjaman", handlers.BuatPeminjamanBaru(peminjamanService))
}
