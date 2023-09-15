package generate

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	errMissingSecret       = errors.New("must provide secret")
	errMissingSubject      = errors.New("must provide subject")
	errMissingAllowedRoles = errors.New("must provide allowed-roles")
	errMissingDefaultRole  = errors.New("must provide default role")
)

type HasuraToken struct {
	Subject              string       `json:"sub"`
	IssuedAtEpochSeconds int          `json:"iat"`
	HasuraClaims         HasuraClaims `json:"https://hasura.io/jwt/claims"`
}

type HasuraClaims struct {
	AllowedRoles []string `json:"x-hasura-allowed-roles"`
	DefaultRole  string   `json:"x-hasura-default-role"`
	UserID       string   `json:"x-hasura-user-id,omitempty"`
}

func NewHasuraToken(cfg Config) HasuraToken {
	return HasuraToken{
		Subject:              cfg.subject,
		IssuedAtEpochSeconds: int(time.Now().Unix()),
		HasuraClaims: HasuraClaims{
			AllowedRoles: cfg.allowedRoles,
			DefaultRole:  cfg.defaultRole,
			UserID:       cfg.userID,
		},
	}
}

func (t HasuraToken) toMapClaims() (jwt.MapClaims, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal HasuraToken: %w", err)
	}

	var mc jwt.MapClaims

	err = json.Unmarshal(b, &mc)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal HasuraToken: %w", err)
	}

	return mc, nil
}

type Config struct {
	secret       string
	subject      string
	allowedRoles []string
	defaultRole  string
	userID       string
}

type option func(c *Config)

func WithSecret(secret string) option {
	return func(c *Config) {
		c.secret = secret
	}
}

func WithSubject(subject string) option {
	return func(c *Config) {
		c.subject = subject
	}
}

func WithAllowedRoles(allowedRoles []string) option {
	return func(c *Config) {
		c.allowedRoles = allowedRoles
	}
}

func WithDefaultRole(defaultRole string) option {
	return func(c *Config) {
		c.defaultRole = defaultRole
	}
}

func WithUserID(userID string) option {
	return func(c *Config) {
		c.userID = userID
	}
}

type CreateTokenOutput struct {
	Token string
}

func CreateToken(opts ...option) (*CreateTokenOutput, error) {
	cfg := Config{}
	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.secret == "" {
		return nil, errMissingSecret
	}

	if cfg.subject == "" {
		return nil, errMissingSubject
	}

	if len(cfg.allowedRoles) == 0 {
		return nil, errMissingAllowedRoles
	}

	if cfg.defaultRole == "" {
		return nil, errMissingDefaultRole
	}

	ht := NewHasuraToken(cfg)

	mc, err := ht.toMapClaims()
	if err != nil {
		return nil, fmt.Errorf("unable to extract claims from token: %w", err)
	}

	// TODO make SigningMethod configurable
	// rs256 is preferred: https://auth0.com/blog/rs256-vs-hs256-whats-the-difference/
	// https://www.scottbrady91.com/jose/jwts-which-signing-algorithm-should-i-use
	// Hasura supports: HS256, HS384, HS512, RS256, RS384, RS512, Ed25519
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)

	// Create and return a complete, signed JWT
	tokenString, err := token.SignedString([]byte(cfg.secret))
	if err != nil {
		return nil, fmt.Errorf("unable to sign token: %w", err)
	}

	return &CreateTokenOutput{Token: tokenString}, nil
}
