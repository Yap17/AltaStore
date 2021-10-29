package userauth_test

import (
	"AltaStore/business"
	"AltaStore/business/user"
	userMock "AltaStore/business/user/mocks"
	"AltaStore/business/userauth"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id        = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	email     = "email@test.com"
	firstname = "firstname"
	lastname  = "lastname"
	password  = "password"
)

var (
	userAuthService userauth.Service
	userService     userMock.Service

	userData user.User
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
func setup() {
	userData = user.User{
		ID:        id,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  password,
	}
	userAuthService = userauth.NewService(&userService)
}

func TestUserLogin(t *testing.T) {
	t.Run("Expect User not found", func(t *testing.T) {
		userService.On("FindUserByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrUnAuthorized).Once()
		user, err := userService.FindUserByEmailAndPassword(id, password)

		assert.NotNil(t, err)
		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrUnAuthorized)
	})
	t.Run("Expect User not found", func(t *testing.T) {
		userService.On("FindUserByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrUnAuthorized).Once()
		user, err := userAuthService.UserLogin(email, password)

		assert.NotNil(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, err, business.ErrUnAuthorized)
		assert.Equal(t, user, "")
	})
	t.Run("Expect Login Success", func(t *testing.T) {
		userService.On("FindUserByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&userData, nil).Once()
		token, err := userAuthService.UserLogin(email, password)

		assert.NotNil(t, token)
		assert.Nil(t, err)

		// assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestCreateToken(t *testing.T) {
	t.Run("Expect create token success", func(t *testing.T) {
		tokenDetail, err := userAuthService.CreateToken(&userData)

		assert.Nil(t, err)
		assert.NotNil(t, tokenDetail)
	})
}
