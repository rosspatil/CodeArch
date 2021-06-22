package crypto

import (
	"context"
	"crypto/md5"
	"crypto/sha512"
	"fmt"

	"github.com/rosspatil/codearch/runtime/models"
	"github.com/rosspatil/codearch/runtime/utils"
)

type Hashing struct {
	Field string `json:"field,omitempty"`
}

func (h *Hashing) Exceute(ctx context.Context, c *models.Controller) (interface{}, error) {
	value := utils.ResolveValue(h.Field, c)
	return fmt.Sprintf("%x", md5.New().Sum([]byte(value.(string)))), nil
}

type HashingWithSalt struct {
	Field string `json:"field,omitempty"`
	Salt  string `json:"salt,omitempty"`
}

func (h *HashingWithSalt) Exceute(ctx context.Context, c *models.Controller) (interface{}, error) {
	fieldVal := utils.ResolveValue(h.Field, c)
	field := []byte(fieldVal.(string))
	salt := []byte(h.Salt)
	field = append(field, salt...)
	field = sha512.New().Sum(field)
	return fmt.Sprintf("%x", field), nil
}
