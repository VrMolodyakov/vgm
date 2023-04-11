package token

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
}

type tokenHandler struct {
	accessPair  KeyPair
	refreshPair KeyPair
}

func NewTokenHandler(accessPair KeyPair, refreshPair KeyPair) *tokenHandler {
	return &tokenHandler{accessPair: accessPair, refreshPair: refreshPair}
}

func (t *tokenHandler) CreateAccessToken(ttl time.Duration, payload interface{}, role string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(t.accessPair.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("couldn't parse private key: %w ", err)
	}
	return create(ttl, payload, role, key, false)
}

func (t *tokenHandler) CreateRefreshToken(ttl time.Duration, payload interface{}) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(t.refreshPair.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("couldn't parse private key: %w ", err)
	}
	return createRefreshToken(ttl, payload, key)
}

func create(ttl time.Duration, payload interface{}, role string, key *rsa.PrivateKey, refresh bool) (string, error) {
	now := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["role"] = role
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("couldn't sign key due to %w", err)
	}
	return token, nil
}

func createRefreshToken(ttl time.Duration, payload interface{}, key *rsa.PrivateKey) (string, error) {
	now := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("couldn't sign key due to %w", err)
	}
	return token, nil
}

func (t *tokenHandler) ValidateAccessToken(token string) (interface{}, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(t.accessPair.PublicKey)
	if err != nil {
		return -1, fmt.Errorf("couldn't parse public key: %w ", err)
	}
	return validate(token, key)
}

func (t *tokenHandler) ValidateRefreshToken(token string) error {
	key, err := jwt.ParseRSAPublicKeyFromPEM(t.refreshPair.PublicKey)
	if err != nil {
		return fmt.Errorf("couldn't parse public key: %w ", err)
	}
	_, err = validate(token, key)
	return err

}

func validate(token string, key *rsa.PublicKey) (interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, signed := t.Method.(*jwt.SigningMethodRSA); !signed {
			return nil, fmt.Errorf("unexpected method - %v", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't parse : %w", err)
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token : %w", err)
	}
	return claims["sub"], nil
}
