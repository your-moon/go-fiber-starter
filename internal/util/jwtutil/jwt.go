package jwtutil

import (
	"math/rand"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	bcrypt "golang.org/x/crypto/bcrypt"
)

// access secret key
var accessKey = []byte("bmVnb3RpYXRpb24")

type TokenClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

// ExtractJWTString Get claim from token string
func ExtractJWTString(tokenString string) (*TokenClaims, error) {
	retClaim := &TokenClaims{}
	JwtToken, err := jwt.ParseWithClaims(
		tokenString,
		retClaim,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(accessKey), nil
		},
	)
	if err == nil {
		if !JwtToken.Valid {
			return retClaim, nil
		}
	}
	return retClaim, err
}

func GenerateToken(user TokenClaims) string {
	accessExpTime := time.Now().Add(365 * 24 * time.Hour)
	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: accessExpTime},
		},
	}).SignedString(accessKey)
	return accessToken
}

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
