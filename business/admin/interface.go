package admin

//Service outgoing port for Admin
type Service interface {
	//InsertAdmin Insert new Admin into storage
	InsertAdmin(insertAdminSpec InsertAdminSpec) error

	//FindAdminByAdminnameAndPassword If data not found will return nil
	FindAdminByEmailAndPassword(email string, password string) (*Admin, error)

	FindAdminByEmail(email string) (*Admin, error)

	//FindAdminByID If data not found will return nil without error
	FindAdminByID(id string) (*Admin, error)

	//UpdateAdminPaasword if data not found or old password wrong will return error
	UpdateAdminPassword(id string, password, oldPassword string, modifier string) error

	//UpdateAdminToken if data not found  will return error
	//UpdateAdminToken(id string, token string) error

	//UpdateAdmin if data not found will return error
	UpdateAdmin(id string, updateAdminSpec UpdateAdminSpec, modifier string) error

	//DeleteAdmin if data not found will return error
	DeleteAdmin(id string, modifier string) error
}

//Repository ingoing port for Admin
type Repository interface {
	//InsertAdmin Insert new Admin into storage
	InsertAdmin(Admin Admin) error

	//FindAdminByAdminnameAndPassword If data not found will return nil
	FindAdminByEmailAndPassword(email string, password string) (*Admin, error)

	FindAdminByEmail(email string) (*Admin, error)

	//FindAdminByID If data not found will return nil without error
	FindAdminByID(id string) (*Admin, error)

	//UpdateAdmin if data not found will return error
	UpdateAdmin(Admin Admin) error

	//UpdateAdmin if data not found will return error
	UpdateAdminPassword(Admin Admin) error

	//DeleteAdmin if data not found will return error
	DeleteAdmin(Admin Admin) error
}
