package dto

type HomePage struct {
	MainProduct DisplayProduct   `json:"main"`
	FirstRow    []DisplayProduct `json:"first_row"`
	SecondRow   []DisplayProduct `json:"second_row"`
}
