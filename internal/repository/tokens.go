package repository

import (
	"context"
	"time"

	"github.com/Cr4z1k/MEDODS-test-task/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	refTTL = time.Hour * 24 * 30
)

type TokensRepo struct {
	db *mongo.Collection
}

func NewTokensRepo(db *mongo.Collection) *TokensRepo {
	return &TokensRepo{db: db}
}

func (r *TokensRepo) GetTokens(guid, refToken string) error {
	userToken := core.UserToken{
		Guid:         guid,
		RefreshToken: refToken,
		ExpiresAt:    time.Now().Add(refTTL),
	}

	filter := bson.D{{Key: "guid", Value: guid}}

	err := r.db.FindOne(context.TODO(), filter).Err()
	if err == mongo.ErrNoDocuments {
		return insertUserToken(r, userToken)
	} else if err != nil {
		return err
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "ref_token", Value: userToken.RefreshToken},
			{Key: "exp_at", Value: userToken.ExpiresAt},
		}},
	}

	if _, err := r.db.UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}

	return nil
}

func insertUserToken(r *TokensRepo, userToken core.UserToken) error {
	_, err := r.db.InsertOne(context.TODO(), userToken)
	if err != nil {
		return err
	}

	return nil
}
