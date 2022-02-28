package service

import "webapi/pkg/model"

var Address = map[int]model.AddressUser{}
var countAddress int

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateAddress(address model.AddressUser) (int, error) {
	countAddress++
	address.Id = countAddress
	Address[address.Id] = address
	return address.Id, nil
}

func (s *Service) OutputAllAddresses() []model.AddressUser {
	var allAddress []model.AddressUser
	for _, address := range Address {
		allAddress = append(allAddress, address)
	}
	return allAddress
}
