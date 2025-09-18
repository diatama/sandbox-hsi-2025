package routes

import (
	"perpustakaan/api/handlers"
	"perpustakaan/pkg/buku"

	"github.com/gofiber/fiber/v2"
)

func BukuRoutes(api fiber.Router, bukuService *buku.BukuService) {
	api.Get("/buku", handlers.LoadAllBuku(bukuService))
	api.Post("/buku", handlers.BuatBukuBaru(bukuService))
}
