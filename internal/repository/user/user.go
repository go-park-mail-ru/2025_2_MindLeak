package user

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(email string, password string, name string) (*User, error)
	GetUserById(id uuid.UUID) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]*User, error)
	DeleteUser(id uuid.UUID) (bool, error)
}

type User struct {
	Id       uuid.UUID `json:"-"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
}

type InMemoryUser struct {
	Users []User
	mu    sync.RWMutex
}

func NewInMemoryUser() *InMemoryUser {
	return &InMemoryUser{
		Users: make([]User, 0),
	}
}

func (mem *InMemoryUser) CreateUser(email string, password string, name string) (*User, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	for _, user := range mem.Users {
		if user.Email == email {
			return nil, errors.New("this user is already registered")
		}
	}
	user := User{
		Id:       uuid.New(),
		Email:    email,
		Password: password,
		Name:     name,
		Avatar:   "https://sun9-88.userapi.com/s/v1/ig2/P_e5HW2lWX3ZxayBg73NnzbHzyhxFCXtBseRjSrN_NbemNC78OpkeYfJeXcTOXqyR8NhSwizZKqJEq_R8PhQo607.jpg?quality=95&as=32x40,48x60,72x90,108x135,160x200,240x300,360x450,480x600,540x675,640x800,720x900,1080x1350,1280x1600,1440x1800,1620x2025&from=bu&cs=1620x0",
	}
	mem.Users = append(mem.Users, user)
	copyUser := user
	return &copyUser, nil
}

func (mem *InMemoryUser) GetUserById(userID uuid.UUID) (*User, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	for i := range mem.Users {
		if mem.Users[i].Id == userID {
			copyUser := mem.Users[i]
			return &copyUser, nil
		}
	}
	return nil, errors.New("user not found")
}

func (mem *InMemoryUser) GetUserByEmail(email string) (*User, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	for i := range mem.Users {
		if mem.Users[i].Email == email {
			copyUser := mem.Users[i]
			return &copyUser, nil
		}
	}
	return nil, errors.New("user not found")
}

func (mem *InMemoryUser) GetAllUsers() ([]*User, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()
	usersCopy := make([]*User, len(mem.Users))
	for i := range mem.Users {
		temp := mem.Users[i]
		usersCopy[i] = &temp
	}

	return usersCopy, nil
}

func (mem *InMemoryUser) DeleteUser(userID uuid.UUID) (bool, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	for idx, user := range mem.Users {
		if user.Id == userID {
			mem.Users[idx] = mem.Users[len(mem.Users)-1]
			mem.Users = mem.Users[:len(mem.Users)-1]
			return true, nil
		}
	}
	return false, errors.New("user not found")
}
