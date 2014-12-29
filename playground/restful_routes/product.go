package main

type Product struct {
	Id    string `json:"id, omitempty"`
	Title string `json:"title, omitempty"`
}

type Products struct {
	Total int       `json:"total"`
	Data  []Product `json:"data"`
}
