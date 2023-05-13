package internal

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"gorm.io/gorm"
)

type LinkRepository interface {
	Get(ctx context.Context, id int64) (*Link, error)
	Save(ctx context.Context, link *Link) error
}

type LinkMemoryRepository struct {
	links sync.Map

	idgen atomic.Int64
}

func NewMemoryRepository() *LinkMemoryRepository {
	return &LinkMemoryRepository{
		links: sync.Map{},
	}
}

func (repo *LinkMemoryRepository) Get(ctx context.Context, id int64) (*Link, error) {
	link, ok := repo.links.Load(id)
	if !ok {
		return nil, fmt.Errorf("link not found")
	}

	return link.(*Link), nil
}

func (repo *LinkMemoryRepository) Save(ctx context.Context, link *Link) error {
	if link.ID == 0 {
		link.ID = repo.idgen.Add(1)
	}

	repo.links.Store(link.ID, link)
	return nil
}

type LinkPostgresRepository struct {
	db *gorm.DB
}

func NewLinkPostgresRepository(db *gorm.DB) *LinkPostgresRepository {
	return &LinkPostgresRepository{
		db,
	}
}

func (repo *LinkPostgresRepository) Get(ctx context.Context, id int64) (*Link, error) {
	link := new(Link)
	err := repo.db.WithContext(ctx).First(link, "id = ?", id).Error
	return link, err
}

func (repo *LinkPostgresRepository) Save(ctx context.Context, link *Link) error {
	return repo.db.WithContext(ctx).Save(link).Error
}
