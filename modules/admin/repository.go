package admin

import (
	"AltaStore/business/admin"

	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Implementasi repositori admin
type Repository struct {
	DB *gorm.DB
}

type Admin struct {
	ID        string    `gorm:"type:uuid;primary_key"`
	Email     string    `gorm:"email;index:idx_email,unique;type:varchar(50)"`
	FirstName string    `gorm:"firstname;type:varchar(50)"`
	LastName  string    `gorm:"lastname;type:varchar(50)"`
	Password  string    `gorm:"password;type:varchar(100)"`
	CreatedAt time.Time `gorm:"created_at"`
	CreatedBy string    `gorm:"created_by;type:varchar(50)"`
	UpdatedAt time.Time `gorm:"updated_at"`
	UpdatedBy string    `gorm:"updated_by;type:varchar(50)"`
	DeletedAt time.Time `gorm:"deleted_at"`
	DeletedBy string    `gorm:"deleted_by;type:varchar(50)"`
}

func newAdminTable(admin admin.Admin) *Admin {

	return &Admin{
		admin.ID,
		admin.Email,
		admin.FirstName,
		admin.LastName,
		admin.Password,
		admin.CreatedAt,
		admin.CreatedBy,
		admin.UpdatedAt,
		admin.UpdatedBy,
		admin.DeletedAt,
		admin.DeletedBy,
	}
}

func (col *Admin) Toadmin() admin.Admin {
	var admin admin.Admin

	admin.ID = col.ID
	admin.Email = col.Email
	admin.FirstName = col.FirstName
	admin.LastName = col.LastName
	admin.CreatedAt = col.CreatedAt
	admin.CreatedBy = col.CreatedBy
	admin.UpdatedAt = col.UpdatedAt
	admin.UpdatedBy = col.UpdatedBy
	admin.DeletedAt = col.DeletedAt
	admin.DeletedBy = col.DeletedBy

	return admin
}

// Menghasilkan ORM DB untuk admin repository
func NewDBRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

//InsertAdmin Insert new Admin into storage
func (repo *Repository) InsertAdmin(admin admin.Admin) error {

	adminData := newAdminTable(admin)

	err := repo.DB.Create(adminData).Error
	if err != nil {
		return err
	}

	return nil
}

//FindAdminByEmailAndPassword If data not found will return nil
func (repo *Repository) FindAdminByEmailAndPassword(email string, password string) (*admin.Admin, error) {

	var adminData Admin

	err := repo.DB.Where("email = ?", email).Where("deleted_by = ''").First(&adminData).Error
	if err != nil {
		return nil, err
	}

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword([]byte(adminData.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	Admin := adminData.Toadmin()

	return &Admin, nil
}

//FindAdminByID If data not found will return nil without error
func (repo *Repository) FindAdminByID(id string) (*admin.Admin, error) {

	var adminData Admin

	err := repo.DB.Where("id = ?", id).Where("deleted_by = ''").First(&adminData).Error
	if err != nil {
		return nil, err
	}

	Admin := adminData.Toadmin()

	return &Admin, nil
}

//UpdateAdminPassword if data not found or old password wrong will return error
func (repo *Repository) UpdateAdminPassword(admin admin.Admin) error {
	adminData := newAdminTable(admin)

	err := repo.DB.Model(&adminData).Updates(Admin{
		Password: adminData.Password,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

//UpdateAdmin Update existing Admin in database
func (repo *Repository) UpdateAdmin(admin admin.Admin) error {
	AdminData := newAdminTable(admin)

	err := repo.DB.Model(&AdminData).Updates(Admin{
		Email:     AdminData.Email,
		FirstName: AdminData.FirstName,
		LastName:  AdminData.LastName,
		Password:  AdminData.Password,
		UpdatedAt: AdminData.UpdatedAt,
		UpdatedBy: AdminData.UpdatedBy,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

//DeleteAdmin Set IsDelete true in database
func (repo *Repository) DeleteAdmin(admin admin.Admin) error {
	adminData := newAdminTable(admin)

	err := repo.DB.Model(&adminData).Updates(Admin{
		DeletedBy: adminData.DeletedBy,
		DeletedAt: adminData.DeletedAt,
	}).Error
	if err != nil {
		return err
	}

	return nil
}
