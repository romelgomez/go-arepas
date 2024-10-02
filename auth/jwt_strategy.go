package auth

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
)

type JWTStrategy struct {
	IssuerURL string
	JWKSURL   string
	Algorithm string
}

func NewJWTStrategy(issuerURL, jwksURL, algorithm string) *JWTStrategy {
	return &JWTStrategy{
		IssuerURL: issuerURL,
		JWKSURL:   jwksURL,
		Algorithm: algorithm,
	}
}

func (s *JWTStrategy) GetJWKS() (jwk.Set, error) {
	resp, err := http.Get(s.JWKSURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()

	jwks, err := jwk.ParseReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWKS: %w", err)
	}

	return jwks, nil
}

func (s *JWTStrategy) ValidateToken(tokenStr string) (*jwt.Token, error) {
	jwks, err := s.GetJWKS()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("missing kid in token header")
		}

		key, found := jwks.LookupKeyID(kid)
		if !found {
			return nil, fmt.Errorf("unable to find key")
		}

		// Extract the public key
		var pubKey interface{}
		if err := key.Raw(&pubKey); err != nil {
			return nil, fmt.Errorf("failed to get public key: %w", err)
		}

		return pubKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
