package core

import "time"

type UserToken struct {
	Guid         string    `bson:"guid"`
	RefreshToken string    `bson:"ref_token"`
	ExpiresAt    time.Time `bson:"exp_at"`
}
