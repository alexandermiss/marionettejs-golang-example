package repositories

import (
	"context"
    "gorm.io/gorm"
	"github.com/pkg/errors"
	"linguanski/src/models"
	"strconv"
)

type providerRepository struct {
	Db *gorm.DB
}

func NewProviderRepository(db *gorm.DB) ProviderRepository {
	return &providerRepository{
		Db: db,
	}
}

func (a providerRepository) GetAll(ctx context.Context) ([]models.Provider, error) {
	var providers []models.Provider

	result := a.Db.Find(&providers)
	if result.Error != nil {
		return nil, result.Error
	}

	if providers == nil {
		return []models.Provider{}, nil
	}
	return providers, nil
}

func (a providerRepository) Save(ctx context.Context, provider *models.Provider) error {
	if provider == nil {
		return errors.New("Input parameter provider is nil")
	}

	result := a.Db.Create(&provider)
	if result.Error != nil {
		return errors.Wrapf(result.Error, "Failed to insert provider %v", provider)
	}
	return nil
}

func (a providerRepository) FindByID(ctx context.Context, id string) (models.Provider, error) {
	if len(id) == 0 {
		return models.Provider{}, errors.New("Provider id is incorrect")
	}

	index, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.Provider{}, errors.Wrap(err, "Cannot convert provider id")
	}
    var provider models.Provider
    result := a.Db.First(&provider, index)
	if result.Error != nil {
		return models.Provider{}, nil
	}
    
	return provider, err
}

func (a providerRepository) DeleteByID(ctx context.Context, id string) (int64, error) {
	if len(id) == 0 {
		return -1, errors.New("Provider id is incorrect")
	}

	index, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return -1, errors.Wrap(err, "Cannot convert Provider id")
	}

	var provider models.Provider
    result := a.Db.Delete(&provider, index)
	return result.RowsAffected, result.Error
}

func (a providerRepository) Update(ctx context.Context, providerModel models.Provider, id string) (int64, error) {
 
    var provider models.Provider
    first := a.Db.First(&provider, id)
	if first.Error != nil {
		return -1, first.Error
	}

    result := a.Db.Model(&provider).Select("*").Updates(providerModel)
	return result.RowsAffected, result.Error
}

