package users

import (
	"context"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getUserState(src map[string]storage.User) map[string]storage.User {
	result := map[string]storage.User{}
	for key, value := range src {
		result[key] = value
	}
	return result
}

func TestMockUser_CreateUser(t *testing.T) {
	ctx := context.Background()
	usersState := map[string]storage.User{
		"user1": {ID: 1, Login: "user1", HashedPassword: "pass1"},
		"user2": {ID: 2, Login: "user2", HashedPassword: "pass2"},
	}

	tests := []struct {
		name       string
		user       storage.User
		lastUsedID int

		wantUsers      map[string]storage.User
		wantUser       *storage.User
		wantLastUsedID int
		wantError      error
	}{
		{
			name:       "Test 1. User not exist",
			user:       storage.User{Login: "user3", HashedPassword: "pass_3"},
			lastUsedID: 2,
			wantUsers: map[string]storage.User{
				"user1": {ID: 1, Login: "user1", HashedPassword: "pass1"},
				"user2": {ID: 2, Login: "user2", HashedPassword: "pass2"},
				"user3": {ID: 3, Login: "user3", HashedPassword: "pass_3"},
			},
			wantUser:       &storage.User{ID: 3, Login: "user3", HashedPassword: "pass_3"},
			wantLastUsedID: 3,
			wantError:      nil,
		},
		{
			name:           "Test 2. User already exist",
			user:           storage.User{Login: "user2", HashedPassword: "pass2"},
			lastUsedID:     2,
			wantUsers:      getUserState(usersState),
			wantUser:       nil,
			wantLastUsedID: 2,
			wantError:      ErrUserExist,
		},
		{
			name:           "Test 3. User already exist(login same, pass differs)",
			user:           storage.User{Login: "user2", HashedPassword: "pass2"},
			wantUsers:      getUserState(usersState),
			wantUser:       nil,
			wantLastUsedID: 2,
			wantError:      ErrUserExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := MockUser{Users: getUserState(usersState), LastUsedID: tt.lastUsedID}

			gotUser, gotError := rep.CreateUser(ctx, tt.user)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantUsers, rep.Users)
			assert.Equal(t, tt.wantUser, gotUser)
		})
	}
}

func TestMockUser_GetUser(t *testing.T) {
	ctx := context.Background()
	usersState := map[string]storage.User{
		"user1": {Login: "user1", HashedPassword: "pass1"},
		"user2": {Login: "user2", HashedPassword: "pass2"},
	}
	rep := MockUser{}

	tests := []struct {
		name      string
		user      storage.User
		wantUser  *storage.User
		wantError error
	}{
		{
			name:      "Test 1. User exist, only login",
			user:      storage.User{Login: "user2"},
			wantUser:  &storage.User{Login: "user2", HashedPassword: "pass2"},
			wantError: nil,
		},
		{
			name:      "Test 2. User exist, login and password(wrong or correct)",
			user:      storage.User{Login: "user2", HashedPassword: "pass2"},
			wantUser:  &storage.User{Login: "user2", HashedPassword: "pass2"},
			wantError: nil,
		},
		{
			name:      "Test 3. User not exist",
			user:      storage.User{Login: "user3"},
			wantUser:  nil,
			wantError: ErrUserNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep.Users = getUserState(usersState)

			gotUser, gotError := rep.GetUser(ctx, tt.user)
			assert.ErrorIs(t, gotError, tt.wantError)
			assert.Equal(t, tt.wantUser, gotUser)
		})
	}
}
