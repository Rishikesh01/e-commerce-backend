package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRatingRepo interface {
	FindByID(id uuid.UUID) (*model.ProductRating, error)
	SaveWithTrackRating(existingModel *model.ProductRating, rating *model.TrackRating) error
}

type productRatingRepo struct {
	db *gorm.DB
}

func NewProductRatingRepo(db *gorm.DB) ProductRatingRepo {
	return &productRatingRepo{db: db}
}

func (p *productRatingRepo) SaveWithTrackRating(existingModel *model.ProductRating, rating *model.TrackRating) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(existingModel).Error; err != nil {
			return err
		}

		if err := tx.Save(rating).Error; err != nil {
			return err
		}
		return nil
	})
}

func (p *productRatingRepo) FindByID(ID uuid.UUID) (*model.ProductRating, error) {
	var m *model.ProductRating
	if err := p.db.Where("product_id=?", ID).FirstOrCreate(m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
