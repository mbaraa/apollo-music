package entities

type Storage struct {
	Size float64 `json:"size"`
	Used float64 `json:"used"`
	Free float64 `json:"free"`
}
