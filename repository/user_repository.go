package repository

import (
	"encoding/json"
	"os"

	"github.com/maulanafahrul/mnc-test/model"
)

type UserRepository interface {
	GetUserByUsername(string) *model.UserModel
	GetUserByUsernameForView(string) *model.UserModel
	Update(*model.UserModel) error
	Delete(*model.UserModel) error
	Create(*model.UserModel) error
}

type userRepositoryImpl struct {
	users []model.UserModel
}

func (r *userRepositoryImpl) GetUserByUsername(username string) *model.UserModel {

	for _, usr := range r.users {
		if usr.Username == username {
			return &usr
		}
	}
	return nil
}
func (r *userRepositoryImpl) GetUserByUsernameForView(username string) *model.UserModel {
	for _, usr := range r.users {
		if usr.Username == username {
			userView := &model.UserModel{
				Id:       usr.Id,
				Username: usr.Username,
			}
			return userView
		}
	}
	return nil
}

func (r *userRepositoryImpl) Update(user *model.UserModel) error {
	// Update the user in the slice
	for i, usr := range r.users {
		if usr.Id == user.Id {
			r.users[i] = *user
			break
		}
	}

	// Open the JSON file
	file, err := os.OpenFile("data/users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the users slice back into the file
	err = json.NewEncoder(file).Encode(r.users)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryImpl) Delete(user *model.UserModel) error {
	// Remove the user from the slice
	for i, u := range r.users {
		if u.Id == user.Id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			break
		}
	}

	// Open the JSON file
	file, err := os.OpenFile("data/users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the users slice back into the file
	err = json.NewEncoder(file).Encode(r.users)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryImpl) Create(user *model.UserModel) error {
	// Check if the user exists in the slice
	found := false
	for i, u := range r.users {
		if u.Id == user.Id {
			r.users[i] = *user
			found = true
			break
		}
	}
	if !found {
		r.users = append(r.users, *user)
	}

	// Open the JSON file
	file, err := os.OpenFile("data/users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(r.users)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository() UserRepository {
	repo := &userRepositoryImpl{}

	// Open the JSON file
	file, err := os.Open("data/users.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	// Decode the file into the users slice
	err = json.NewDecoder(file).Decode(&repo.users)
	if err != nil {
		return nil
	}

	return repo
}
