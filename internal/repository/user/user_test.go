package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T, mem *InMemoryUser)
	}{
		{
			name: "CreateUser creates new user",
			run: func(t *testing.T, mem *InMemoryUser) {
				u, err := mem.CreateUser("test@example.com", "password", "TestUser")
				assert.NoError(t, err)
				assert.NotNil(t, u)
				assert.NotEqual(t, uuid.Nil, u.Id)
				assert.Equal(t, "test@example.com", u.Email)
				assert.Equal(t, "password", u.Password)
				assert.Equal(t, "TestUser", u.Name)
				assert.NotEmpty(t, u.Avatar)

				assert.Len(t, mem.Users, 1)
				assert.Equal(t, u.Id, mem.Users[0].Id)
			},
		},
		{
			name: "CreateUser returns error if email already registered",
			run: func(t *testing.T, mem *InMemoryUser) {
				_, _ = mem.CreateUser("test@example.com", "password", "TestUser")
				_, err := mem.CreateUser("test@example.com", "newpassword", "NewUser")
				assert.EqualError(t, err, "this user is already registered")
			},
		},
		{
			name: "GetUserById returns existing user",
			run: func(t *testing.T, mem *InMemoryUser) {
				u, _ := mem.CreateUser("test@example.com", "password", "TestUser")
				got, err := mem.GetUserById(u.Id)
				assert.NoError(t, err)
				assert.Equal(t, u.Id, got.Id)
				assert.Equal(t, u.Email, got.Email)
				assert.Equal(t, u.Password, got.Password)
				assert.Equal(t, u.Name, got.Name)
				assert.Equal(t, u.Avatar, got.Avatar)
			},
		},
		{
			name: "GetUserById returns error if not found",
			run: func(t *testing.T, mem *InMemoryUser) {
				_, err := mem.GetUserById(uuid.New())
				assert.EqualError(t, err, "user not found")
			},
		},
		{
			name: "GetUserByEmail returns existing user",
			run: func(t *testing.T, mem *InMemoryUser) {
				u, _ := mem.CreateUser("test@example.com", "password", "TestUser")
				got, err := mem.GetUserByEmail("test@example.com")
				assert.NoError(t, err)
				assert.Equal(t, u.Id, got.Id)
				assert.Equal(t, u.Email, got.Email)
				assert.Equal(t, u.Password, got.Password)
				assert.Equal(t, u.Name, got.Name)
				assert.Equal(t, u.Avatar, got.Avatar)
			},
		},
		{
			name: "GetUserByEmail returns error if not found",
			run: func(t *testing.T, mem *InMemoryUser) {
				_, err := mem.GetUserByEmail("nonexistent@example.com")
				assert.EqualError(t, err, "user not found")
			},
		},
		{
			name: "GetAllUsers returns all users",
			run: func(t *testing.T, mem *InMemoryUser) {
				u1, _ := mem.CreateUser("test1@example.com", "password1", "TestUser1")
				u2, _ := mem.CreateUser("test2@example.com", "password2", "TestUser2")
				all, err := mem.GetAllUsers()
				assert.NoError(t, err)
				assert.Len(t, all, 2)
				assert.Equal(t, u1.Id, all[0].Id)
				assert.Equal(t, u2.Id, all[1].Id)
			},
		},
		{
			name: "GetAllUsers returns empty slice if no users",
			run: func(t *testing.T, mem *InMemoryUser) {
				all, err := mem.GetAllUsers()
				assert.NoError(t, err)
				assert.Empty(t, all)
			},
		},
		{
			name: "DeleteUser deletes existing user",
			run: func(t *testing.T, mem *InMemoryUser) {
				u, _ := mem.CreateUser("test@example.com", "password", "TestUser")
				ok, err := mem.DeleteUser(u.Id)
				assert.True(t, ok)
				assert.NoError(t, err)
				assert.Empty(t, mem.Users)
			},
		},
		{
			name: "DeleteUser returns error if not found",
			run: func(t *testing.T, mem *InMemoryUser) {
				ok, err := mem.DeleteUser(uuid.New())
				assert.False(t, ok)
				assert.EqualError(t, err, "user not found")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mem := NewInMemoryUser()
			test.run(t, mem)
		})
	}
}
