package repositories

import (
	"context"
	"linguanski/src/models"
)

type ClientRepository interface {
	GetAll(ctx context.Context) ([]models.Client, error)
	Save(ctx context.Context, client *models.Client) error
	FindByID(ctx context.Context, id string) (models.Client, error)
	DeleteByID(ctx context.Context, id string) (int64, error)
	Update(ctx context.Context, client models.Client, id string) (int64, error)
}

type ProviderRepository interface {
	GetAll(ctx context.Context) ([]models.Provider, error)
	Save(ctx context.Context, client *models.Provider) error
	FindByID(ctx context.Context, id string) (models.Provider, error)
	DeleteByID(ctx context.Context, id string) (int64, error)
	Update(ctx context.Context, provider models.Provider, id string) (int64, error)
}

