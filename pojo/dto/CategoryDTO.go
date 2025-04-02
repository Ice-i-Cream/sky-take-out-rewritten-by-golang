package dto

type CategoryDTO struct {
	Id   int         `json:"id"`
	Type interface{} `json:"type"`
	Name string      `json:"name"`
	Sort interface{} `json:"sort"`
}
