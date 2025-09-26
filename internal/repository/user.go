package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(email string, password string) (*User, error)
	GetUserById(id uuid.UUID) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]*User, error)
	DeleteUser(id uuid.UUID) (bool, error)
}

type User struct {
	Id       uuid.UUID
	Email    string
	Password string `json:"-"`
}

type InMemoryUser struct {
	Users []*User
	mu    sync.RWMutex
}

func NewInMemoryUser() *InMemoryUser {
	return &InMemoryUser{
		Users: make([]*User, 0),
	}
}

func (mem *InMemoryUser) CreateUser(email string, password string) (*User, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	for _, val := range mem.Users {
		if val.Email == email {
			return nil, errors.New("this user is already registered")
		}
	}
	user := &User{
		Id:       uuid.New(),
		Email:    email,
		Password: password,
	}
	mem.Users = append(mem.Users, user)
	return user, nil
}

func (mem *InMemoryUser) GetUserById(id uuid.UUID) (*User, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	for _, val := range mem.Users {
		if val.Id == id {
			return val, nil
		}
	}
	return nil, errors.New("user not found")
}

func (mem *InMemoryUser) GetUserByEmail(email string) (*User, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	for _, val := range mem.Users {
		if val.Email == email {
			return val, nil
		}
	}
	return nil, errors.New("user not found")
}

func (mem *InMemoryUser) GetAllUsers() ([]*User, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	return mem.Users, nil
}

func (mem *InMemoryUser) DeleteUser(id uuid.UUID) (bool, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	for idx, user := range mem.Users {
		if user.Id == id {
			mem.Users[idx] = mem.Users[len(mem.Users)-1]
			mem.Users = mem.Users[:len(mem.Users)-1]
			return true, nil
		}
	}
	return false, errors.New("user not found")
}
