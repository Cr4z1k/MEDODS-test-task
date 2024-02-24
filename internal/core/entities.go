package core

import "time"

type Tokens struct {
	AccToken string
	RefToken string
}

type UserRefToken struct {
	Guid         string    `bson:"guid"`
	RefreshToken string    `bson:"ref_token"`
	ExpiresAt    time.Time `bson:"exp_at"`
}
