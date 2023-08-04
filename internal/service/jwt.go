package service

import (
	"crypto/ecdsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"exusiai.dev/roguestats-backend/internal/app/appbundle"
	"exusiai.dev/roguestats-backend/internal/app/appconfig"
)

var (
	ErrExpiredToken  = errors.New("token expired")
	JWTSigningMethod = jwt.SigningMethodES512
)

type JWT struct {
	Config *appconfig.Config

	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewJWT(config *appconfig.Config) (JWT, error) {
	privateKey, err := jwt.ParseECPrivateKeyFromPEM(config.JWTPrivateKey)
	if err != nil {
		return JWT{}, err
	}

	publicKey, err := jwt.ParseECPublicKeyFromPEM(config.JWTPublicKey)
	if err != nil {
		return JWT{}, err
	}

	return JWT{
		Config: config,

		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// Sign signs a JWT token with the given user ID as the `sub` claim.
func (s *JWT) Sign(userId string) (string, error) {
	now := time.Now()

	jwtToken := jwt.NewWithClaims(JWTSigningMethod, jwt.MapClaims{
		"iss": "roguestats/" + appbundle.Version,
		"sub": userId,
		"aud": "roguestats",
		"exp": now.Add(s.Config.JWTExpiration).Unix(),
		"nbf": now.Unix(),
		"iat": now.Unix(),
	})

	return jwtToken.SignedString(s.privateKey)
}

// Verify verifies the token and returns the user ID (`sub` claim) along with the expiration time if the token is valid.
func (s *JWT) Verify(token string) (sub string, t time.Time, err error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return s.publicKey, nil
	}, jwt.WithAudience("roguestats"), jwt.WithValidMethods([]string{JWTSigningMethod.Name}), jwt.WithIssuedAt())
	if err != nil {
		return "", t, err
	}

	if !jwtToken.Valid {
		return "", t, jwt.ErrInvalidKey
	}

	// verify `exp` claim
	claims := jwtToken.Claims.(jwt.MapClaims)
	jwtT, err := claims.GetExpirationTime()
	if err != nil {
		return "", t, err
	}

	if jwtT.Before(time.Now()) {
		return "", t, ErrExpiredToken
	}

	sub, err = claims.GetSubject()
	if sub == "" || err != nil {
		return "", t, jwt.ErrInvalidKey
	}

	t = jwtT.Time

	return sub, t, nil
}
