package jwt

import (
	"context"
	"crypto/rsa"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

const (
	viewerNameClaim = "viewer_name"
	tokenContextKey = "Authorization"
)

type Provider struct {
	privateKey      *rsa.PrivateKey
	sessionLiveTime time.Duration
}

func NewProvider(
	privateKey *rsa.PrivateKey,
	sessionLiveTime time.Duration,
) *Provider {
	return &Provider{
		privateKey:      privateKey,
		sessionLiveTime: sessionLiveTime,
	}
}

func (p *Provider) GenerateToken(ctx context.Context, viewerName string) (string, error) {
	now := time.Now()
	expireAt := now.Add(p.sessionLiveTime)

	claims := jwt.MapClaims{
		viewerNameClaim: viewerName,
		"exp":           expireAt.Unix(),
		"iat":           now.Unix(),
		"nbf":           now.Add(-1 * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(p.privateKey)
}

func (p *Provider) IsAuthorized(ctx context.Context) bool {
	token, ok := ctx.Value(tokenContextKey).(string)
	if !ok || token == "" {
		log.Debug("invalid token context value")
		return false
	}

	token = strings.TrimPrefix(token, "Bearer ")

	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return &p.privateKey.PublicKey, nil
	})
	if err != nil {
		log.WithError(err).Error("parse jwt error")
		return false
	}

	if !tkn.Valid {
		return false
	}

	return true
}
