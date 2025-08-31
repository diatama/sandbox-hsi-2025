package main

import (
	"fmt"
	"hsiweek05/pegawai"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initData(dbConn *gorm.DB) {
	result := dbConn.Find(&pegawai.Pegawai{})

	if result.RowsAffected == 0 {
		bulkPegawai := []*pegawai.Pegawai{
			{Nama: "Anya Petrova", Posisi: "Chief Synergy Officer", GajiBulanan: 27800000},
			{Nama: "Elara Vex", Posisi: "Data Wrangler", GajiBulanan: 15666000},
			{Nama: "Jax Ryder", Posisi: "Cloud Solutions Architect", GajiBulanan: 25450000},
			{Nama: "Kael Stormwall", Posisi: "UI/UX Alchemist", GajiBulanan: 17100000},
			{Nama: "Zephyr Vale", Posisi: "Quantum Analyst", GajiBulanan: 19250000},
		}

		bulkInsert := dbConn.Create(bulkPegawai)

		if bulkInsert.Error != nil {
			log.Fatal("Gagal melakukan init data pegawai", bulkInsert.Error.Error())
		}
	}
}

func displayAllData(dbConn *gorm.DB) {
	var allPegawai []pegawai.Pegawai
	queryAll := dbConn.Find(&allPegawai)

	if queryAll.Error != nil {
		log.Fatal("Terjadi Kesalahan:", "Tidak dapat melakukan query semua data pegawai.", queryAll.Error)
	} else {
		fmt.Println("# Munculkan semua data")
		fmt.Println("+--------------------------------+")

		for _, row := range allPegawai {
			row.TampilkanInformasi()
		}

		fmt.Println("Total Data:", queryAll.RowsAffected, "baris.")
	}
}

func main() {
	dsn := "db-user:some-strong-password@tcp(127.0.0.1:3306)/hsi_sandbox?charset=utf8mb4&parseTime=True&loc=Local"
	dbConn, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Cek Koneksi DB Gagal
	if dbErr != nil {
		log.Fatal("Koneksi Gagal", "Tidak dapat terhubung ke database.", dbErr.Error())
	}

	// Migrasi struktur data
	dbConn.AutoMigrate(&pegawai.Pegawai{})

	// Init Data
	initData(dbConn)

	// Tampilkan List Data Pegawai
	displayAllData(dbConn)

	// Update Salah Satu Data Pegawai
	fmt.Println("\n# Mengubah salah satu data")
	var updatedData pegawai.Pegawai
	dbConn.First(&updatedData, 2)
	updatedData.GajiBulanan = 20000000
	dbConn.Save(&updatedData)
	fmt.Println("Data pegawai dengan id: 2 berhasil diubah.")
	updatedData.TampilkanInformasi()

	// Hapus Salah Satu Data Pegawai
	fmt.Println("\n# Menghapus salah satu data")
	var deletedData pegawai.Pegawai
	dbConn.First(&deletedData, 3) // Get data by id 3
	dbConn.Delete(&deletedData)
	fmt.Println("Data pegawai dengan id: 3 berhasil dihapus", deletedData)
	displayAllData(dbConn)
}
