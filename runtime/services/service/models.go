package service

import (
	"context"

	"github.com/rosspatil/codearch/runtime/models"
	"github.com/rosspatil/codearch/runtime/services/condition"
	"github.com/rosspatil/codearch/runtime/services/crypto"
	"github.com/rosspatil/codearch/runtime/services/customcode"
	"github.com/rosspatil/codearch/runtime/services/customerrors"
	"github.com/rosspatil/codearch/runtime/services/load"
	"github.com/rosspatil/codearch/runtime/services/store"
)

type T int

const (
	Load T = iota
	Store
	CustomCode
	Condition
	Crypto
)

type Step struct {
	Type        T                     `json:"type,omitempty"`
	Load        load.Load             `json:"load,omitempty"`
	Store       store.Store           `json:"store,omitempty"`
	CustomeCode customcode.CustomCode `json:"custome_code,omitempty"`
	Condition   condition.Condition   `json:"condition,omitempty"`
	Crypto      crypto.Crypto         `json:"crypto,omitempty"`
}

type Service struct {
	Name           string                      `json:"name,omitempty"`
	Path           string                      `json:"path,omitempty"`
	Method         string                      `json:"method,omitempty"`
	RequestBody    string                      `json:"request_body,omitempty"`
	Steps          []Step                      `json:"steps,omitempty"`
	Response       Response                    `json:"response,omitempty"`
	OverrideErrors customerrors.OverrideErrors `json:"override_errors,omitempty"`
	m              *models.Controller          `json:"m,omitempty"`
}

type Response struct {
	HttpCode int               `json:"http_code,omitempty"`
	Field    string            `json:"field,omitempty"`
	Headers  map[string]string `json:"headers,omitempty"`
}

type Exce interface {
	Executor(context.Context, *models.Controller) error
}
