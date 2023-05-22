package service

import (
	"context"
	"github.com/pkg/errors"
	"linguanski/src/models"
	"linguanski/src/repositories"
)


type ProviderManager interface {
	GetAll(ctx context.Context) ([]models.Provider, error)
	Save(ctx context.Context, provider *models.Provider) error
	FindByID(ctx context.Context, id string) (models.Provider, error)
	DeleteByID(ctx context.Context, id string) (int64, error)
	Update(ctx context.Context, provider models.Provider, id string) (int64, error)
}

type providerManager struct {
	providerRepository repositories.ProviderRepository
}

func NewProviderManager(providerRepository repositories.ProviderRepository) ProviderManager {
	return &providerManager{
		providerRepository: providerRepository,
	}
}

func (m *providerManager) GetAll(ctx context.Context) ([]models.Provider, error) {
	providers, err := m.providerRepository.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get all providers")
	}
	return providers, nil
}

func (m *providerManager) Save(ctx context.Context, provider *models.Provider) error {
	err := m.providerRepository.Save(ctx, provider)
	if err != nil {
		return errors.Wrap(err, "Failed to save provider")
	}
	return nil
}

func (m *providerManager) FindByID(ctx context.Context, id string) (models.Provider, error) {
	provider, err := m.providerRepository.FindByID(ctx, id)
	if err != nil {
		return models.Provider{}, errors.Wrap(err, "Failed to find provider")
	}
	return provider, nil
}

func (m *providerManager) DeleteByID(ctx context.Context, id string) (int64, error) {
	rows, err := m.providerRepository.DeleteByID(ctx, id)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to delete provider")
	}
	return rows, err
}

func (m *providerManager) Update(ctx context.Context, provider models.Provider, id string) (int64, error) {
	rows, err := m.providerRepository.Update(ctx, provider, id)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to update provider")
	}
	return rows, err
}

