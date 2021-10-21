package auth

import (
	"AltaStore/business"
	"AltaStore/business/user"
	"AltaStore/config"
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/google/uuid"
)

//=============== The implementation of those interface put below =======================
type service struct {
	userService user.Service
	//authService Repository
}

// //NewService Construct user service object
// func NewService(userService user.Service, authService Repository) Service {
// 	return &service{
// 		userService, authService,
// 	}
// }

//NewService Construct user service object
func NewService(userService user.Service) Service {
	return &service{
		userService,
	}
}

//Login by given user Email and Password, return error if not exist
func (s *service) Login(username string, password string) (string, error) {
	user, err := s.userService.FindUserByEmailAndPassword(username, password)
	if err != nil {
		return "", business.ErrNotFound
	}
	td, err := s.CreateToken(user)
	if err != nil {
		return "", err
	}
	// err = s.authService.InsertToken(user, td)
	// if err != nil {
	// 	return "", business.ErrNotFound
	// }
	return td.AccessToken, nil
}

func (s *service) CreateToken(user *user.User) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 24).Unix()
	td.AccessUuid = uuid.New().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = user.ID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(config.GetConfig().JwtSecretKey))
	if err != nil {
		return nil, err
	}
	return td, nil
}
