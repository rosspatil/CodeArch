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
	Fields string
}

func (h *Hashing) Exceute(ctx context.Context, c *models.Controller) (interface{}, error) {
	value := utils.ResolveValue(h.Fields, c)
	return fmt.Sprintf("%x", md5.New().Sum([]byte(value.(string)))), nil
}

type HashingWithSalt struct {
	Fields string
	Salt   string
}

func (h *HashingWithSalt) Exceute(ctx context.Context, c *models.Controller) (interface{}, error) {
	fieldVal := utils.ResolveValue(h.Fields, c)
	saltVal := utils.ResolveValue(h.Salt, c)
	field := []byte(fieldVal.(string))
	salt := []byte(saltVal.(string))
	field = append(field, salt...)
	field = sha512.New().Sum(field)
	return fmt.Sprintf("%x", field), nil
}
