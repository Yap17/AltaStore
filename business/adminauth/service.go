package adminauth

import (
	"AltaStore/business"
	"AltaStore/business/admin"
	"AltaStore/config"
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/google/uuid"
)

//=============== The implementation of those interface put below =======================
type service struct {
	adminService admin.Service
	//authService Repository
}

// //NewService Construct admin service object
// func NewService(adminService admin.Service, authService Repository) Service {
// 	return &service{
// 		adminService, authService,
// 	}
// }

//NewService Construct admin service object
func NewService(adminService admin.Service) Service {
	return &service{
		adminService,
	}
}

//Login by given admin Email and Password, return error if not exist
func (s *service) AdminLogin(adminname string, password string) (string, error) {
	admin, err := s.adminService.FindAdminByEmailAndPassword(adminname, password)
	if err != nil {
		return "", business.ErrUnAuthorized
	}
	td, err := s.CreateToken(admin)
	if err != nil {
		return "", business.ErrUnAuthorized
	}
	// err = s.authService.InsertToken(admin, td)
	// if err != nil {
	// 	return "", business.ErrNotFound
	// }
	return td.AccessToken, nil
}

func (s *service) CreateToken(admin *admin.Admin) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 24).Unix()
	td.AccessUuid = uuid.New().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["userId"] = admin.ID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(config.GetConfig().JwtSecretKey))
	if err != nil {
		return nil, err
	}
	return td, nil
}
