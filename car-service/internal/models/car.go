package models

type Car struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Brand       string `json:"brand"`
	UserID      int    `json:"user_id"`
	EngineID    int    `json:"engine_id"`
}
