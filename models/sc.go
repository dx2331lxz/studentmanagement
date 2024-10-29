package models

type SC struct {
	Sno   int    `json:"sno,omitempty"`
	Cno   int    `json:"cno,omitempty"`
	Grade string `json:"grade,omitempty"`
}
