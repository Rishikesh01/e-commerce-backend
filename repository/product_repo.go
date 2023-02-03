package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepo interface {
	Save(product *model.Product) error
	FindByLikeName(name string) ([]model.Product, error)
	FindByID(id uuid.UUID) (*model.Product, error)
	FindAll() ([]model.Product, error)
	FindAllLimitRelations() ([]model.Product, error)
	FindByIDLimitRelations(id uuid.UUID) (*model.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}

func (p *productRepo) Save(product *model.Product) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(product).Error; err != nil {
			return err
		}

		if err := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Save(&product.ProductSeller).Error; err != nil {
			return err
		}
		return nil
	})
}

func (p *productRepo) FindAll() ([]model.Product, error) {
	var products []model.Product
	if err := p.db.Model(&model.Product{}).Preload("ProductSeller").
		Preload("ProductRating").Find(&products).Error; err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return products, nil
}

func (p *productRepo) FindAllLimitRelations() ([]model.Product, error) {
	var products []model.Product

	if err := p.db.Debug().Model(&model.Product{}).
		Preload("ProductSeller", func(db *gorm.DB) *gorm.DB {
			return db.Limit(3)
		}).
		Preload("ProductRating").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
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
	if err := p.db.Model(&model.Product{}).Preload("ProductSeller").Preload("ProductRating").Where("id=?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepo) FindByIDLimitRelations(id uuid.UUID) (*model.Product, error) {
	var product model.Product

	if err := p.db.Model(&model.Product{}).Raw("select * from products pd join( select * from product_sellers ps limit 3) as p on pd.id = p.product_id where id =?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
