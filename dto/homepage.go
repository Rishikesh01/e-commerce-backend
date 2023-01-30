package dto

type HomePage struct {
	MainProduct Product   `json:"main"`
	FirstRow    []Product `json:"first_row"`
	SecondRow   []Product `json:"second_row"`
}
