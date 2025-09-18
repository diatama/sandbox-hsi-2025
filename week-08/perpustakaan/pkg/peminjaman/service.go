package peminjaman

import "perpustakaan/pkg/entities"

type PeminjamanRepo interface {
	Create(peminjaman entities.Peminjaman) error
	GetAll() ([]entities.Peminjaman, error)
}

type PeminjamanService struct {
	repo PeminjamanRepo
}

func NewPeminjamanService(repo PeminjamanRepo) *PeminjamanService {
	return &PeminjamanService{repo: repo}
}

func (s *PeminjamanService) CreatePeminjaman(peminjaman entities.Peminjaman) error {
	return s.repo.Create(peminjaman)
}

func (s *PeminjamanService) GetAllPeminjaman() ([]entities.Peminjaman, error) {
	return s.repo.GetAll()
}
