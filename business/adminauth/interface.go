package adminauth

import "AltaStore/business/admin"

//Service outgoing port for Admin
type Service interface {

	//Login If data not found will return nil without error
	AdminLogin(email string, password string) (string, error)

	//Create Token If Failed will return error
	CreateToken(admin *admin.Admin) (*TokenDetails, error)
}

// type Repository interface {
// 	//Insert Token to redis if failed will return error
// 	InsertToken(admin *admin.Admin, td *TokenDetails) error
// }
