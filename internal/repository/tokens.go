package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Cr4z1k/MEDODS-test-task/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

func (r *TokensRepo) GetUserByGuid(user core.UserRefToken) (core.UserToken, error) {
	var resultUser core.UserToken

	filter := bson.D{{Key: "guid", Value: user.Guid}}

	result := r.db.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		objectID, err := insertUserToken(r, user)
		if err != nil {
			return core.UserToken{}, err
		}

		resultUser.ObjectID = objectID
		resultUser.Guid = user.Guid

		return resultUser, nil
	} else if result.Err() != nil {
		return core.UserToken{}, result.Err()
	}

	if err := result.Decode(&resultUser); err != nil {
		return core.UserToken{}, err
	}

	return resultUser, nil
}

func (r *TokensRepo) UpdateRefTokenInfo(user core.UserToken) error {
	filter := bson.D{{Key: "_id", Value: user.ObjectID.(primitive.ObjectID)}}

	err := r.db.FindOne(context.TODO(), filter).Err()
	if err != nil {
		return err
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "_id", Value: user.ObjectID},
			{Key: "ref_token", Value: user.RefreshToken},
			{Key: "exp_at", Value: time.Now().Add(refTTL)},
		}},
	}

	if _, err := r.db.UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}

	return nil
}

func (r *TokensRepo) CheckRefresh(objectID primitive.ObjectID, refToken []byte) (core.UserToken, error) {
	filter := bson.D{{Key: "_id", Value: objectID}}

	var user core.UserToken

	if err := r.db.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return core.UserToken{}, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.RefreshToken), refToken)
	if err != nil {
		return core.UserToken{}, errors.New("invalid refresh token")
	}

	if user.ExpiresAt.Before(time.Now()) {
		return core.UserToken{}, errors.New("refresh token has expired")
	}

	return user, nil
}

func insertUserToken(r *TokensRepo, user core.UserRefToken) (interface{}, error) {
	result, err := r.db.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}
