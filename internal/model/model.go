package model

type User struct {
	Family       string `json:"family,omitempty"`
	Name         string `json:"name,omitempty"`
	Otch         string `json:"otch,omitempty"`
	Registration string `json:"registration,omitempty"`
}

type Shop struct {
	Title   string `json:"title,omitempty"`
	Address string `json:"address,omitempty"`
	Working string `json:"working,omitempty"`
}
