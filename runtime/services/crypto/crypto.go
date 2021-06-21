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
	Type            T
	ResultField     string
	Hash            Hashing
	HashingWithSalt HashingWithSalt
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
	}
	if err != nil {
		return err
	}
	m.SetP(data, l.ResultField)
	return nil
}
