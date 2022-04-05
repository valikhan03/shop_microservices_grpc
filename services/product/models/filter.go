package models

type Filter struct {
	Title    string `json:"Title"`
	Category string `json:"Category"`
	MinPrice uint64 `json:"MinPrice"`
	MaxPrice uint64 `json:"MaxPrice"`
}
