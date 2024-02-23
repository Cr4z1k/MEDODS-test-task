package repository

import "go.mongodb.org/mongo-driver/mongo"

type Tokens interface {
	GetTokens(guid, refToken string) error
}

type Repository struct {
	Tokens
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{Tokens: NewTokensRepo(db)}
}
