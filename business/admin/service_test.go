package admin_test

import (
	"AltaStore/business"
	"AltaStore/business/admin"
	adminMock "AltaStore/business/admin/mocks"
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
	createdby = "creator"
	updatedby = "modifier"
	deletedby = ""
)

var (
	adminService    admin.Service
	adminRepository adminMock.Repository

	adminData       admin.Admin
	insertAdminData admin.InsertAdminSpec
	updateAdminData admin.UpdateAdminSpec
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

	insertAdminData = admin.InsertAdminSpec{
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  password,
	}

	updateAdminData = admin.UpdateAdminSpec{
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
	}

	adminService = admin.NewService(&adminRepository)
}

func TestFindAdminByID(t *testing.T) {
	t.Run("Expect found the admin", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()

		admin, err := adminService.FindAdminByID(id)

		assert.Nil(t, err)
		assert.NotNil(t, admin)

		assert.Equal(t, id, admin.ID)
		assert.Equal(t, email, admin.Email)
		assert.Equal(t, firstname, admin.FirstName)
		assert.Equal(t, lastname, admin.LastName)
		assert.Equal(t, password, admin.Password)
	})

	t.Run("Expect admin not found", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		admin, err := adminService.FindAdminByID(id)

		assert.NotNil(t, err)
		assert.Nil(t, admin)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindAdminByEmailAndPassword(t *testing.T) {
	t.Run("Expect found the admin", func(t *testing.T) {
		adminRepository.On("FindAdminByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&adminData, nil).Once()

		admin, err := adminService.FindAdminByEmailAndPassword(email, password)

		assert.Nil(t, err)
		assert.NotNil(t, admin)

		assert.Equal(t, id, admin.ID)
		assert.Equal(t, email, admin.Email)
		assert.Equal(t, firstname, admin.FirstName)
		assert.Equal(t, lastname, admin.LastName)
		assert.Equal(t, password, admin.Password)
	})

	t.Run("Expect admin not found", func(t *testing.T) {
		adminRepository.On("FindAdminByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		admin, err := adminService.FindAdminByEmailAndPassword(email, password)

		assert.NotNil(t, err)
		assert.Nil(t, admin)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindAdminByEmail(t *testing.T) {
	t.Run("Expect found the admin", func(t *testing.T) {
		adminRepository.On("FindAdminByEmail", mock.AnythingOfType("string")).Return(&adminData, nil).Once()

		admin, err := adminService.FindAdminByEmail(email)

		assert.Nil(t, err)
		assert.NotNil(t, admin)

		assert.Equal(t, id, admin.ID)
		assert.Equal(t, email, admin.Email)
		assert.Equal(t, firstname, admin.FirstName)
		assert.Equal(t, lastname, admin.LastName)
		assert.Equal(t, password, admin.Password)
	})

	t.Run("Expect admin not found", func(t *testing.T) {
		adminRepository.On("FindAdminByEmail", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		admin, err := adminService.FindAdminByEmail(email)

		assert.NotNil(t, err)
		assert.Nil(t, admin)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestInsertAdmin(t *testing.T) {
	t.Run("Expect admin email exist", func(t *testing.T) {
		adminRepository.On("FindAdminByEmail", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		err := adminService.InsertAdmin(insertAdminData)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrDataExists)
	})
	t.Run("Expect insert admin success", func(t *testing.T) {
		adminRepository.On("FindAdminByEmail", mock.AnythingOfType("string")).Return(nil, nil).Once()
		adminRepository.On("InsertAdmin", mock.AnythingOfType("admin.Admin")).Return(nil).Once()

		err := adminService.InsertAdmin(insertAdminData)

		assert.Nil(t, err)
	})

	t.Run("Expect insert admin failed", func(t *testing.T) {
		adminRepository.On("FindAdminByEmail", mock.AnythingOfType("string")).Return(nil, nil).Once()
		adminRepository.On("InsertAdmin", mock.AnythingOfType("admin.Admin")).Return(business.ErrInternalServer).Once()

		err := adminService.InsertAdmin(insertAdminData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestUpdateAdmin(t *testing.T) {
	t.Run("Expect admin not found", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		err := adminService.UpdateAdmin(id, updateAdminData, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
	t.Run("Expect update admin success", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		adminRepository.On("UpdateAdmin", mock.AnythingOfType("admin.Admin")).Return(nil).Once()

		err := adminService.UpdateAdmin(id, updateAdminData, id)

		assert.Nil(t, err)
	})
	t.Run("Expect update admin failed", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		adminRepository.On("UpdateAdmin", mock.AnythingOfType("admin.Admin")).Return(business.ErrInternalServer).Once()

		err := adminService.UpdateAdmin(id, updateAdminData, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestUpdateAdminPassword(t *testing.T) {
	t.Run("Expect admin not found by id", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		err := adminService.UpdateAdminPassword(id, password, password, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
	t.Run("Expect admin not found by email and password", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		adminRepository.On("FindAdminByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrPasswordMisMatch).Once()
		err := adminService.UpdateAdminPassword(id, password, password, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrPasswordMisMatch)
	})
	t.Run("Expect update admin password success", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		adminRepository.On("FindAdminByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		adminRepository.On("UpdateAdminPassword", mock.AnythingOfType("admin.Admin")).Return(nil).Once()

		err := adminService.UpdateAdminPassword(id, password, password, id)

		assert.Nil(t, err)
	})
	t.Run("Expect update admin password failed", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		adminRepository.On("FindAdminByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		adminRepository.On("UpdateAdminPassword", mock.AnythingOfType("admin.Admin")).Return(business.ErrInternalServer).Once()

		err := adminService.UpdateAdminPassword(id, password, password, email)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestDeleteAdmin(t *testing.T) {
	t.Run("Expect admin not found", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		err := adminService.DeleteAdmin(id, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
	t.Run("Expect delete admin success", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		adminRepository.On("DeleteAdmin", mock.AnythingOfType("admin.Admin")).Return(nil).Once()

		err := adminService.DeleteAdmin(id, id)

		assert.Nil(t, err)
	})
	t.Run("Expect delete admin failed", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		adminRepository.On("DeleteAdmin", mock.AnythingOfType("admin.Admin")).Return(business.ErrInternalServer).Once()

		err := adminService.DeleteAdmin(id, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}
