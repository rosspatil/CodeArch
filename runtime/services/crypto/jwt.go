package crypto

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rosspatil/codearch/runtime/models"
	"github.com/rosspatil/codearch/runtime/utils"
)

type JWTEncoding struct {
	Key         string            `json:"key,omitempty"`
	Iss         string            `json:"iss,omitempty"`
	Sub         string            `json:"sub,omitempty"`
	Aud         string            `json:"aud,omitempty"`
	Exp         int64             `json:"exp,omitempty"`
	Alg         string            `json:"alg,omitempty"`
	Claim       map[string]string `json:"claim,omitempty"`
	ResultField string            `json:"result_field,omitempty"`
}

type Claim struct {
	jwt.StandardClaims
	Data map[string]interface{} `json:"data,omitempty"`
}

func (j *JWTEncoding) Exceute(ctx context.Context, m *models.Controller) (string, error) {
	data := map[string]interface{}{}
	for k, v := range j.Claim {
		data[k] = utils.ResolveValue(v, m)
	}
	key, ok := utils.ResolveEnvironmentVariable(j.Key)
	if !ok {
		return "", fmt.Errorf("JWT Signin key not set")
	}
	expiration := j.Exp
	if expiration == 0 {
		expiration = int64(time.Hour) * 24
	}
	t := jwt.NewWithClaims(jwt.GetSigningMethod(j.Alg), Claim{
		StandardClaims: jwt.StandardClaims{
			Audience:  j.Aud,
			ExpiresAt: time.Now().Add(time.Duration(j.Exp)).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    j.Iss,
			Subject:   j.Sub,
		},
		Data: data,
	})
	token, err := t.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

type JWTDecoding struct {
	Key   string `json:"key,omitempty"`
	Field string `json:"field,omitempty"`
}

func (j *JWTDecoding) Exceute(ctx context.Context, m *models.Controller) error {
	key, ok := utils.ResolveEnvironmentVariable(j.Key)
	if !ok {
		return fmt.Errorf("JWT Signin key not set")
	}
	if j.Field == "" {
		j.Field = "request.headers.authorization"
	}
	token := utils.ResolveValue(j.Field, m)
	tokenStr := token.(string)
	claim := &Claim{}
	_, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return err
	}
	m.SetP(claim.Data, "request.auth")
	return nil
}
