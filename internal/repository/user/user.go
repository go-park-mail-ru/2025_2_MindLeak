package user

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrUserExists   = errors.New("this user is already registered")
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, password string, name string) (*User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (bool, error)
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

func (mem *InMemoryUser) CreateUser(ctx context.Context, email string, password string, name string) (*User, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	for _, user := range mem.Users {
		if user.Email == email {
			return nil, fmt.Errorf("%w: %s", ErrUserExists, email)
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

func (mem *InMemoryUser) GetUserById(ctx context.Context, userID uuid.UUID) (*User, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	for i := range mem.Users {
		if mem.Users[i].Id == userID {
			copyUser := mem.Users[i]
			return &copyUser, nil
		}
	}
	return nil, fmt.Errorf("%w: %s", ErrUserNotFound, userID)
}

func (mem *InMemoryUser) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	for i := range mem.Users {
		if mem.Users[i].Email == email {
			copyUser := mem.Users[i]
			return &copyUser, nil
		}
	}
	return nil, fmt.Errorf("%w: %s", ErrUserNotFound, email)
}

func (mem *InMemoryUser) GetAllUsers(ctx context.Context) ([]*User, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()
	usersCopy := make([]*User, len(mem.Users))
	for i := range mem.Users {
		temp := mem.Users[i]
		usersCopy[i] = &temp
	}

	return usersCopy, nil
}

func (mem *InMemoryUser) DeleteUser(ctx context.Context, userID uuid.UUID) (bool, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	for idx, user := range mem.Users {
		if user.Id == userID {
			mem.Users = append(mem.Users[:idx], mem.Users[idx+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("%w: %s", ErrUserNotFound, userID)
}
