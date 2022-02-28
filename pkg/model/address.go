package model

type AddressUser struct {
	Id     int
	IdUser int
	City   string `json:"city"`
	Street string `json:"street"`
	House  string `json:"house"`
}
