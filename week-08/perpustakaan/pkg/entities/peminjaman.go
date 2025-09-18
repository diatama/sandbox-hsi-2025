package entities

import "time"

type Peminjaman struct {
	TanggalPeminjaman   *time.Time
	TanggalPengembalian *time.Time
	NamaPeminjam        string
	BukuYangDipinjam    uint
	Buku                Buku `gorm:"foreignKey:BukuYangDipinjam;references:ID"`
}
