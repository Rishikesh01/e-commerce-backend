package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepo interface {
	Save(product *model.Product) error
	FindByLikeName(name string) ([]model.Product, error)
	FindByID(id uuid.UUID) (*model.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}

func (p *productRepo) Save(product *model.Product) error {
	return p.db.Save(product).Error
}

func (p *productRepo) FindByLikeName(name string) ([]model.Product, error) {
	var products []model.Product
	if err := p.db.Where("name = ?", name).Find(&products).Error; err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return products, nil
}

func (p *productRepo) FindByID(id uuid.UUID) (*model.Product, error) {
	var product model.Product
	if err := p.db.Where("id=?", id).First(product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
