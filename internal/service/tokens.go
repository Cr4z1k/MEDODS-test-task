package service

import (
	"encoding/base64"
	"encoding/hex"
	"time"

	"github.com/Cr4z1k/MEDODS-test-task/internal/core"
	"github.com/Cr4z1k/MEDODS-test-task/internal/repository"
	"github.com/Cr4z1k/MEDODS-test-task/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	ttl = time.Minute * 15
)

type TokensService struct {
	r repository.Tokens
	t auth.TokenManager
}

func NewTokenService(r repository.Tokens, t auth.TokenManager) *TokensService {
	return &TokensService{
		r: r,
		t: t,
	}
}

func (s *TokensService) GetTokens(guid string) (core.Tokens, error) {
	Refuser := core.UserRefToken{
		Guid: guid,
	}

	user, err := s.r.GetUserByGuid(Refuser)
	if err != nil {
		return core.Tokens{}, err
	}

	tokens, err := s.generateAndUpdateNewTokens(user)
	if err != nil {
		return core.Tokens{}, err
	}

	return tokens, nil
}

func (s *TokensService) RefreshTokens(refTokenBase64 string) (core.Tokens, error) {
	refTokenBytes, err := base64.StdEncoding.DecodeString(refTokenBase64)
	if err != nil {
		return core.Tokens{}, err
	}

	objID := refTokenBytes[20:]

	hexObjID, err := hex.DecodeString(string(objID))
	if err != nil {
		return core.Tokens{}, err
	}

	objectID, err := primitive.ObjectIDFromHex(string(hexObjID))
	if err != nil {
		return core.Tokens{}, err
	}

	user, err := s.r.CheckRefresh(objectID, refTokenBytes)
	if err != nil {
		return core.Tokens{}, err
	}

	tokens, err := s.generateAndUpdateNewTokens(user)
	if err != nil {
		return core.Tokens{}, err
	}

	return tokens, nil
}

func (s *TokensService) generateAndUpdateNewTokens(user core.UserToken) (core.Tokens, error) {
	refToken, err := s.t.NewRefreshToken(user.ObjectID.(primitive.ObjectID).Hex())
	if err != nil {
		return core.Tokens{}, err
	}

	accToken, err := s.t.NewAccessToken(user.Guid, ttl)
	if err != nil {
		return core.Tokens{}, err
	}

	hashedRefToken, err := bcrypt.GenerateFromPassword([]byte(refToken), bcrypt.DefaultCost)
	if err != nil {
		return core.Tokens{}, err
	}

	user.RefreshToken = string(hashedRefToken)

	err = s.r.UpdateRefTokenInfo(user)
	if err != nil {
		return core.Tokens{}, err
	}

	refTokenBase64 := base64.StdEncoding.EncodeToString([]byte(refToken))

	tokens := core.Tokens{
		AccToken: accToken,
		RefToken: refTokenBase64,
	}

	return tokens, nil
}
