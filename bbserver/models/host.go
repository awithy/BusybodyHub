package models

type Host struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type Hosts []Host
