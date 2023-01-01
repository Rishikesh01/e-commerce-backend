package dto

type HomePage struct {
	MainProduct   Product   `json:"main_product"`
	OtherProducts []Product `json:"other_products"`
}
