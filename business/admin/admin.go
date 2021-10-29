package admin

import (
	"time"
)

//Admin
type Admin struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Password  string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt time.Time
	DeletedBy string
}

//NewAdmin create new Admin
func NewAdmin(
	id string,
	email string,
	firstname string,
	lastname string,
	password string,
	creator string,
	createdAt time.Time) Admin {

	return Admin{
		ID:        id,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  password,
		CreatedAt: createdAt,
		CreatedBy: creator,
	}
}

//ModifyAdmin update existing AdminData
func (oldData *Admin) ModifyAdmin(
	newFirstName,
	newLastName string,
	updatedAt time.Time,
	updater string) Admin {

	return Admin{
		ID:        oldData.ID,
		Email:     oldData.Email,
		FirstName: newFirstName,
		LastName:  newLastName,
		Password:  oldData.Password,
		CreatedAt: oldData.CreatedAt,
		CreatedBy: oldData.CreatedBy,
		UpdatedAt: updatedAt,
		UpdatedBy: updater,
	}
}

// //ModifyAdminToken update existing AdminData
// func (oldData *Admin) ModifyAdminToken(
// 	newToken string,
// 	updatedAt time.Time) Admin {

// 	return Admin{
// 		ID:        oldData.ID,
// 		Email:     oldData.Email,
// 		FirstName: oldData.FirstName,
// 		LastName:  oldData.LastName,
// 		Password:  oldData.Password,
// 		Token:     newToken,
// 		CreatedAt: oldData.CreatedAt,
// 		CreatedBy: oldData.CreatedBy,
// 		UpdatedAt: updatedAt,
// 		UpdatedBy: oldData.ID,
// 	}
// }

//ModifyAdminPassword update existing AdminData
func (oldData *Admin) ModifyAdminPassword(
	newPassword string,
	modifier string,
	updatedAt time.Time) Admin {

	return Admin{
		ID:        oldData.ID,
		Email:     oldData.Email,
		FirstName: oldData.FirstName,
		LastName:  oldData.LastName,
		Password:  newPassword,
		CreatedAt: oldData.CreatedAt,
		CreatedBy: oldData.CreatedBy,
		UpdatedAt: updatedAt,
		UpdatedBy: modifier,
	}
}

func (oldData *Admin) DeleteAdmin(
	deleteAt time.Time,
	deleter string) Admin {

	return Admin{
		ID:        oldData.ID,
		Email:     oldData.Email,
		FirstName: oldData.FirstName,
		LastName:  oldData.LastName,
		Password:  oldData.Password,
		CreatedAt: oldData.CreatedAt,
		CreatedBy: oldData.CreatedBy,
		UpdatedAt: oldData.UpdatedAt,
		UpdatedBy: oldData.UpdatedBy,
		DeletedAt: deleteAt,
		DeletedBy: deleter,
	}
}
