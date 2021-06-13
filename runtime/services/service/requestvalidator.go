package service

import (
	"context"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func validateRequest(ctx context.Context, validation string, req interface{}) error {
	if validation == "" {
		return nil
	}
	schemaLoader := gojsonschema.NewStringLoader(validation)
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		return err
	}
	rs, err := schema.Validate(gojsonschema.NewGoLoader(req))
	if err != nil {
		return err
	}
	if !rs.Valid() {
		return fmt.Errorf("%v", rs.Errors())
	}
	return nil
}
