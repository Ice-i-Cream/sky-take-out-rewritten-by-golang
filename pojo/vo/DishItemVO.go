package vo

type DishItemVO struct {
	Name        string `json:"name"`
	Copies      int64  `json:"copies"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
