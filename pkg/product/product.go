package product

import "time"

type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	// Update(*Model)
	GetAll() (Models, error)
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

func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}
