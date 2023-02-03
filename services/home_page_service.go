package services

import (
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"math/rand"
	"sort"
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
	data, err := h.productRepo.FindAllLimitRelations()
	if err != nil {
		return dto.HomePage{}, err
	}
	page := dto.HomePage{}
	for i := 0; i < 6; i++ {
		if i == 3 {
			k := rand.Intn(len(data))
			sort.Slice(data[k].ProductSeller, func(i, j int) bool {
				return data[k].ProductSeller[i].Price < data[k].ProductSeller[j].Price
			})
			page.MainProduct = dto.DisplayProduct{
				ID:       data[k].ID,
				Name:     data[k].Name,
				SellerID: data[k].ProductSeller[0].SellerID,
				Img:      data[k].PicturePath,
				Price:    data[k].ProductSeller[0].Price,
				Rating:   data[k].ProductRating.TotalRatingScore,
			}
		} else {
			if i%2 == 0 {
				jm := rand.Intn(len(data))
				sort.Slice(data[jm].ProductSeller, func(i, j int) bool {
					return data[jm].ProductSeller[i].Price < data[jm].ProductSeller[j].Price
				})
				page.SecondRow = append(page.SecondRow, dto.DisplayProduct{
					ID:       data[jm].ID,
					Name:     data[jm].Name,
					SellerID: data[jm].ProductSeller[0].SellerID,
					Img:      data[jm].PicturePath,
					Price:    data[jm].ProductSeller[0].Price,
					Rating:   data[jm].ProductRating.TotalRatingScore,
				})
			} else {
				jm := rand.Intn(len(data))
				sort.Slice(data[jm].ProductSeller, func(i, j int) bool {
					return data[jm].ProductSeller[i].Price < data[jm].ProductSeller[j].Price
				})
				page.FirstRow = append(page.SecondRow, dto.DisplayProduct{
					ID:       data[jm].ID,
					Name:     data[jm].Name,
					SellerID: data[jm].ProductSeller[0].SellerID,
					Img:      data[jm].PicturePath,
					Price:    data[jm].ProductSeller[0].Price,
					Rating:   data[jm].ProductRating.TotalRatingScore,
				})
			}
		}
	}

	return page, nil
}
