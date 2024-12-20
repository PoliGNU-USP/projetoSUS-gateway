package models

type Address struct {
	ID      int    `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}
