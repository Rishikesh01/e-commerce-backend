package services

import (
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"math/rand"
)

type HomePageService interface {
	ShowItems() (dto.HomePage, error)
}

type homePageService struct {
	productRepo repository.ProductRepo
}

func NewHomePageService(productRepo repository.ProductRepo) HomePageService {
	return &homePageService{productRepo: productRepo}
}

func (h *homePageService) ShowItems() (dto.HomePage, error) {
	data, err := h.productRepo.FindAll()
	if err != nil {
		return dto.HomePage{}, err
	}
	page := dto.HomePage{}
	for i := 0; i < 6; i++ {
		j := rand.Intn(len(data))
		page.OtherProducts = append(page.OtherProducts, dto.Product{
			ID:     data[j].ID,
			Name:   data[j].Name,
			Price:  data[j].ProductSeller[0].Price,
			Rating: data[j].ProductRating.TotalRatingScore,
		})
		if i == 3 {
			k := rand.Intn(len(data))
			page.MainProduct = dto.Product{
				ID:     data[k].ID,
				Name:   data[k].Name,
				Price:  data[k].ProductSeller[0].Price,
				Rating: data[k].ProductRating.TotalRatingScore,
			}
		}
	}

	return page, nil
}
