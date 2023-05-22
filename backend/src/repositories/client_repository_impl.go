package repositories

import (
	"context"
    "gorm.io/gorm"
	"github.com/pkg/errors"
	"linguanski/src/models"
	"strconv"
)

type clientRepository struct {
	Db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{
		Db: db,
	}
}

func (a clientRepository) GetAll(ctx context.Context) ([]models.Client, error) {
	var clients []models.Client

	result := a.Db.Preload("Provider").Find(&clients)
	if result.Error != nil {
		return nil, result.Error
	}

	if clients == nil {
		return []models.Client{}, nil
	}
	return clients, nil
}

func (a clientRepository) Save(ctx context.Context, client *models.Client) error {
	if client == nil {
		return errors.New("Input parameter client is nil")
	}

	result := a.Db.Create(&client)
	if result.Error != nil {
		return errors.Wrapf(result.Error, "Failed to insert client %v", client)
	}
	return nil
}

func (a clientRepository) FindByID(ctx context.Context, id string) (models.Client, error) {
	if len(id) == 0 {
		return models.Client{}, errors.New("Client id is incorrect")
	}

	index, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.Client{}, errors.Wrap(err, "Cannot convert client id")
	}
    var client models.Client
    result := a.Db.Preload("Provider").First(&client, index)
	if result.Error != nil {
		return models.Client{}, nil
	}
    
	return client, err
}

func (a clientRepository) DeleteByID(ctx context.Context, id string) (int64, error) {
	if len(id) == 0 {
		return -1, errors.New("Client id is incorrect")
	}

	index, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return -1, errors.Wrap(err, "Cannot convert Client id")
	}

	var client models.Client
    result := a.Db.Delete(&client, index)
	return result.RowsAffected, result.Error
}

func (a clientRepository) Update(ctx context.Context, clientModel models.Client, id string) (int64, error) {
 
    var client models.Client
    first := a.Db.First(&client, id)
	if first.Error != nil {
		return -1, first.Error
	}

    result := a.Db.Model(&client).Select("*").Updates(clientModel)
	return result.RowsAffected, result.Error
}

