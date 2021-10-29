package adminauth_test

import (
	"AltaStore/business"
	"AltaStore/business/admin"
	adminMock "AltaStore/business/admin/mocks"
	"AltaStore/business/adminauth"
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
	adminAuthService adminauth.Service
	adminService     adminMock.Service

	adminData admin.Admin
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
func setup() {
	adminData = admin.Admin{
		ID:        id,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  password,
	}
	adminAuthService = adminauth.NewService(&adminService)
}

func TestAdminLogin(t *testing.T) {
	t.Run("Expect admin not found", func(t *testing.T) {
		adminService.On("FindAdminByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrUnAuthorized).Once()
		admin, err := adminService.FindAdminByEmailAndPassword(id, password)

		assert.NotNil(t, err)
		assert.Nil(t, admin)

		assert.Equal(t, err, business.ErrUnAuthorized)
	})
	t.Run("Expect admin not found", func(t *testing.T) {
		adminService.On("FindAdminByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrUnAuthorized).Once()
		admin, err := adminAuthService.AdminLogin(email, password)

		assert.NotNil(t, err)
		assert.NotNil(t, admin)

		assert.Equal(t, err, business.ErrUnAuthorized)
		assert.Equal(t, admin, "")
	})
	t.Run("Expect Login Success", func(t *testing.T) {
		adminService.On("FindAdminByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		token, err := adminAuthService.AdminLogin(email, password)

		assert.NotNil(t, token)
		assert.Nil(t, err)

		// assert.Equal(t, err, business.ErrUnAuthorized)
	})
}

func TestCreateToken(t *testing.T) {
	t.Run("Expect create token success", func(t *testing.T) {
		tokenDetail, err := adminAuthService.CreateToken(&adminData)

		assert.Nil(t, err)
		assert.NotNil(t, tokenDetail)
	})
}
