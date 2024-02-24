package repository

import (
	"github.com/Cr4z1k/MEDODS-test-task/internal/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tokens interface {
	UpdateRefTokenInfo(user core.UserToken) error
	CheckRefresh(objectID primitive.ObjectID, refToken []byte) (core.UserToken, error)
	GetUserByGuid(user core.UserRefToken) (core.UserToken, error)
}

type Repository struct {
	Tokens
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{Tokens: NewTokensRepo(db)}
}
