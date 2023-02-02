package invoiceheader

import "time"

type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error
	// Create(*Model)
	// Update(*Model)
	// GetAll() (Models, error)
	// GetByID(uint) (*Model, error)
	// Delete(uint) error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
