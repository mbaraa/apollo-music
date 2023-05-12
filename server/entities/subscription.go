package entities

import "time"

type Subscription struct {
	Size        float64   `json:"size"`
	Price       float64   `json:"price"`
	ValidBefore time.Time `json:"validBefore"`
}
