package service

import (
	"context"

	"github.com/Jeffail/gabs/v2"
	"github.com/rosspatil/codearch/runtime/services/customcode"
	"github.com/rosspatil/codearch/runtime/services/load"
	"github.com/rosspatil/codearch/runtime/services/store"
)

type T int

const (
	Load T = iota
	Store
	CustomCode
)

type Step struct {
	Type        T                     `json:"type,omitempty"`
	Load        load.Load             `json:"load,omitempty"`
	Store       store.Store           `json:"store,omitempty"`
	CustomeCode customcode.CustomCode `json:"custome_code,omitempty"`
}

type Service struct {
	Name        string `json:"name,omitempty"`
	Path        string `json:"path,omitempty"`
	Method      string `json:"method,omitempty"`
	RequestBody string `json:"request_body,omitempty"`
	Steps       []Step `json:"steps,omitempty"`
	Response    string `json:"response,omitempty"`
	RespnseCode int    `json:"respnse_code,omitempty"`
	m           *gabs.Container
}

type Exce interface {
	Executor(context.Context, *gabs.Container) error
}
