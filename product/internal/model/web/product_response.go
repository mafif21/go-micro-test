package web

import "time"

type ProductResponse struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Quantity   int       `json:"quantity"`
	Price      int       `json:"price"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
