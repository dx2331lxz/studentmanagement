package models

type Course struct {
	Cno     int    `json:"cno,omitempty"`
	Cname   string `json:"cname,omitempty"`
	Cpno    int    `json:"cpno,omitempty"`
	Ccredit int    `json:"ccredit,omitempty"`
}
