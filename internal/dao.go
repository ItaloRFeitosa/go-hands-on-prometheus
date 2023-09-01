package internal

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"gorm.io/gorm"
)

type LinkDAO interface {
	Get(ctx context.Context, id int64) (*Link, error)
	Save(ctx context.Context, link *Link) error
}

type LinkMemoryDAO struct {
	links sync.Map

	idgen atomic.Int64
}

func NewLinkMemoryDAO() *LinkMemoryDAO {
	return &LinkMemoryDAO{
		links: sync.Map{},
	}
}

func (repo *LinkMemoryDAO) Get(ctx context.Context, id int64) (*Link, error) {
	link, ok := repo.links.Load(id)
	if !ok {
		return nil, fmt.Errorf("link not found")
	}

	return link.(*Link), nil
}

func (repo *LinkMemoryDAO) Save(ctx context.Context, link *Link) error {
	if link.ID == 0 {
		link.ID = repo.idgen.Add(1)
	}

	repo.links.Store(link.ID, link)
	return nil
}

type LinkPostgresDAO struct {
	db *gorm.DB
}

func NewLinkPostgresDAO(db *gorm.DB) *LinkPostgresDAO {
	return &LinkPostgresDAO{
		db,
	}
}

func (repo *LinkPostgresDAO) Get(ctx context.Context, id int64) (*Link, error) {
	link := new(Link)
	err := repo.db.WithContext(ctx).First(link, "id = ?", id).Error
	return link, err
}

func (repo *LinkPostgresDAO) Save(ctx context.Context, link *Link) error {
	return repo.db.WithContext(ctx).Save(link).Error
}
