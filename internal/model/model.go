package model

type User struct {
	Family       string `json:"family"`
	Name         string `json:"name"`
	Otch         string `json:"otch"`
	Registration string `json:"registration"`
}

type Shop struct {
	Title   string `json:"title"`
	Address string `json:"address"`
	Working string `json:"working"`
}
