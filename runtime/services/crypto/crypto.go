package crypto

import (
	"context"

	"github.com/rosspatil/codearch/runtime/models"
)

type T int

const (
	Hash T = iota
	HashWithSALT
	JWTEncode
	JWTDecode
	Base64Encode
	Base64Decode
	Encrypt
	Decrypt
)

type crypto interface {
	Execute(ctx context.Context, c *models.Controller) (interface{}, error)
}

type Crypto struct {
	Type            T               `json:"type,omitempty"`
	ResultField     string          `json:"result_field,omitempty"`
	Hash            Hashing         `json:"hash,omitempty"`
	HashingWithSalt HashingWithSalt `json:"hashing_with_salt,omitempty"`
	JWTEncode       JWTEncoding     `json:"jwt_encode,omitempty"`
	JWTDecode       JWTDecoding     `json:"jwt_decode,omitempty"`
}

func (l *Crypto) Execute(ctx context.Context, m *models.Controller) error {
	var (
		data interface{}
		err  error
	)
	switch l.Type {
	case Hash:
		data, err = l.Hash.Exceute(ctx, m)
	case HashWithSALT:
		data, err = l.HashingWithSalt.Exceute(ctx, m)
	case JWTEncode:
		data, err = l.JWTEncode.Exceute(ctx, m)
	case JWTDecode:
		err = l.JWTDecode.Exceute(ctx, m)
	}
	if err != nil {
		return err
	}
	if data != nil {
		m.SetP(data, l.ResultField)
	}
	return nil
}
