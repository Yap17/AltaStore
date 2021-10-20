package user

import (
	"AltaStore/business"
	"AltaStore/util/validator"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//InsertUserSpec create user spec
type InsertUserSpec struct {
	Email     string `validate:"required"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Password  string `validate:"required"`
}

type UpdateUserSpec struct {
	Email     string `validate:"required"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	HandPhone string
	Address   string
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository Repository
}

//NewService Construct user service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

//InsertUser Create new user and store into database
func (s *service) InsertUser(insertUserSpec InsertUserSpec, createdBy string) error {
	err := validator.GetValidator().Struct(insertUserSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(insertUserSpec.Password), bcrypt.DefaultCost)
	if err != nil {
		return business.ErrInvalidSpec
	}
	var newuuid = uuid.New().String()
	user := NewUser(
		newuuid,
		insertUserSpec.Email,
		insertUserSpec.FirstName,
		insertUserSpec.LastName,
		string(hashedPassword),
		newuuid,
		time.Now(),
	)

	err = s.repository.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

//FindUserByUsernameAndPassword If data not found will return nil
func (s *service) FindUserByEmailAndPassword(email string, password string) (*User, error) {
	return s.repository.FindUserByEmailAndPassword(email, password)
}

//FindUserByID If data not found will return nil without error
func (s *service) FindUserByID(id string) (*User, error) {
	return s.repository.FindUserByID(id)
}

//UpdateUserPaasword if data not found or old password wrong will return error
func (s *service) UpdateUserPassword(id string, newpassword, oldPassword string, updatedBy string) error {

	user, err := s.repository.FindUserByID(id)
	if err != nil {
		return err
	} else if user == nil {
		return business.ErrNotFound
	} else if user.DeletedBy != "" {
		return business.ErrUserDeleted
	} else {
		_, err := s.repository.FindUserByEmailAndPassword(user.Email, oldPassword)
		if err != nil {
			return business.ErrPasswordMisMatch
		}
	}

	modifiedUser := user.ModifyUserPassword(
		newpassword,
		time.Now(),
		updatedBy,
	)

	return s.repository.UpdateUser(modifiedUser)
}

//UpdateUser if data not found will return error
func (s *service) UpdateUser(id string, updateUserSpec UpdateUserSpec, updatedBy string) error {
	user, err := s.repository.FindUserByID(id)
	if err != nil {
		return err
	} else if user == nil {
		return business.ErrNotFound
	} else if user.DeletedBy != "" {
		return business.ErrUserDeleted
	}

	modifiedUser := user.ModifyUser(
		updateUserSpec.Email,
		updateUserSpec.FirstName,
		updateUserSpec.LastName,
		updateUserSpec.HandPhone,
		updateUserSpec.Address,
		time.Now(),
		updatedBy,
	)

	return s.repository.UpdateUser(modifiedUser)
}

//Deleteuser if data not found will return error
func (s *service) DeleteUser(id string, deletedBy string) error {
	user, err := s.repository.FindUserByID(id)
	if err != nil {
		return err
	} else if user == nil {
		return business.ErrNotFound
	}

	deleteUser := user.DeleteUser(
		time.Now(),
		deletedBy,
	)

	return s.repository.DeleteUser(deleteUser)
}
