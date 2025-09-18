package entities

type Buku struct {
	ID          uint `gorm:"primaryKey"`
	Judul       string
	Pengarang   string
	TahunTerbit uint
}
