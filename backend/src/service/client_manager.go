package service

import (
	"context"
	"github.com/pkg/errors"
	"linguanski/src/models"
	"linguanski/src/repositories"
)

type ClientManager interface {
	GetAllClient(ctx context.Context) ([]models.Client, error)
	Save(ctx context.Context, client *models.Client) error
	FindByID(ctx context.Context, id string) (models.Client, error)
	DeleteByID(ctx context.Context, id string) (int64, error)
	UpdateClient(ctx context.Context, client models.Client, id string) (int64, error)
}

type clientManager struct {
	clientRepository repositories.ClientRepository
}

func NewClientManager(clientRepository repositories.ClientRepository) ClientManager {
	return &clientManager{
		clientRepository: clientRepository,
	}
}

func (m *clientManager) GetAllClient(ctx context.Context) ([]models.Client, error) {
	clients, err := m.clientRepository.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get all clients")
	}
	return clients, nil
}

func (m *clientManager) Save(ctx context.Context, client *models.Client) error {
	err := m.clientRepository.Save(ctx, client)
	if err != nil {
		return errors.Wrap(err, "Failed to save client")
	}
	return nil
}

func (m *clientManager) FindByID(ctx context.Context, id string) (models.Client, error) {
	client, err := m.clientRepository.FindByID(ctx, id)
	if err != nil {
		return models.Client{}, errors.Wrap(err, "Failed to find client")
	}
	return client, nil
}

func (m *clientManager) DeleteByID(ctx context.Context, id string) (int64, error) {
	rows, err := m.clientRepository.DeleteByID(ctx, id)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to delete client")
	}
	return rows, err
}

func (m *clientManager) UpdateClient(ctx context.Context, client models.Client, id string) (int64, error) {
	rows, err := m.clientRepository.Update(ctx, client, id)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to update client")
	}
	return rows, err
}

