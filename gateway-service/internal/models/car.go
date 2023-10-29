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

type CarResponse struct {
	Data Car     `json:"data"`
	Err  *string `json:"err"`
}

type CarsResponse struct {
	Data []Car   `json:"data"`
	Err  *string `json:"err"`
}
