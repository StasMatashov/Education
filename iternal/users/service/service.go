package service

import (
	"fmt"
	"webapi/pkg/model"
)

var m = map[int]model.User{}
var count int

type Repository interface {
	CreateUser(user model.User) (model.User, error)
}

type Service struct {
	repo Repository
}

func NewService(repository Repository) *Service {
	return &Service{repo: repository}
}

func (s *Service) CreateUser(user model.User) (int, error) {
	count++
	user.Id = count
	err := user.CheckName()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = user.CheckPassword()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	m[user.Id] = user
	fmt.Println(m)
	return user.Id, nil
}

func (s *Service) SearchUserByIDs(IDs []int) []model.User {
	var sliceUser []model.User
	for _, abc := range IDs {
		val, err := identification(abc)
		if err != nil {
			fmt.Println("User not found")
			continue
		}
		sliceUser = append(sliceUser, val)
	}
	return sliceUser
}

func identification(id int) (model.User, error) {
	if val, ok := m[id]; ok {
		return val, nil
	}
	return model.User{}, fmt.Errorf("lol")
}

func (s *Service) OutputAllUsers() []model.User {
	var allUsers []model.User
	for _, user := range m {
		allUsers = append(allUsers, user)
	}
	return allUsers
}
