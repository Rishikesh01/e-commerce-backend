package dto

import "mime/multipart"

type Product struct {
	Id    int64 `json:"id"`
	Image *multipart.File
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
