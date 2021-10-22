package auth

import (
	"time"

	"AltaStore/business/user"

	"github.com/go-redis/redis/v7"
)

type TokenDetails struct {
	AccessToken string
	AccessUuid  string
	AtExpires   int64
}

type Repository struct {
	client *redis.Client
}

func NewRepository(client *redis.Client) *Repository {
	return &Repository{client}
}

func (r *Repository) InsertToken(user *user.User, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	now := time.Now()

	errAccess := r.client.Set(td.AccessUuid, user.ID, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	return nil
}
