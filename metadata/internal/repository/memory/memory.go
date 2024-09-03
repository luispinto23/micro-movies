package memory

import (
	"context"
	"sync"

	"github.com/luispinto23/micro-movies/metadata/internal/repository"
	model "github.com/luispinto23/micro-movies/metadata/pkg"
)

type Repository struct {
	data map[string]*model.Metadata
	sync.RWMutex
}

func New() *Repository {
	return &Repository{
		data: map[string]*model.Metadata{},
	}
}

func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.RLock()
	defer r.RUnlock()

	r.data[id] = metadata

	return nil
}
