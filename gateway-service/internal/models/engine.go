package models

type Engine struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EngineResponse struct {
	Data Engine  `json:"data"`
	Err  *string `json:"err"`
}
