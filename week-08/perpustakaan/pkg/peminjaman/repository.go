package peminjaman

import (
	"perpustakaan/pkg/entities"

	"gorm.io/gorm"
)

type PeminjamanRepository struct {
	db *gorm.DB
}

func NewPeminjamanRepository(db *gorm.DB) *PeminjamanRepository {
	return &PeminjamanRepository{
		db: db,
	}
}

func (r *PeminjamanRepository) Create(peminjaman entities.Peminjaman) error {
	return r.db.Create(&peminjaman).Error
}

func (r *PeminjamanRepository) GetAll() ([]entities.Peminjaman, error) {
	var peminjaman []entities.Peminjaman
	err := r.db.Preload("Buku").Find(&peminjaman).Error
	return peminjaman, err
}
