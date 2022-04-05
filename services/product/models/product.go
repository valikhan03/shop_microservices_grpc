package models

type Product struct {
	Slug        string `json:"Slug"`
	Title       string `json:"Title"`
	Category    string `json:"CategoryID"`
	Price       uint64 `json:"Price"`
	Description string `json:"Description"`
}
