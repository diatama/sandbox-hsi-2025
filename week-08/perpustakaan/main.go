package main

import (
	"perpustakaan/api/routes"
	"perpustakaan/pkg/buku"
	"perpustakaan/pkg/entities"
	"perpustakaan/pkg/peminjaman"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. Konfigurasi database dan insialitasi repository serta service
	dsn := "root:some-strong-password@tcp(127.0.0.1:3306)/hsi_sandbox?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal Terhubung ke Database")
	}

	db.AutoMigrate(&entities.Buku{}, &entities.Peminjaman{})

	// 2. Inisialiasi Repository
	bukuRepository := buku.NewBukuRepository(db)
	peminjamanRepository := peminjaman.NewPeminjamanRepository(db)

	// 3. Inisialisasi service
	bukuService := buku.NewBukuService(bukuRepository)
	peminjamanService := peminjaman.NewPeminjamanService(peminjamanRepository)

	// 4. Inisialisasi Fiber dan setup routes
	app := fiber.New()
	api := app.Group("/api")

	routes.BukuRoutes(api, bukuService)
	routes.PeminjamanRoutes(api, peminjamanService)

	app.Listen(":3000")
}
