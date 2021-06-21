package models

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/rosspatil/code-arch/types"
)

type Controller struct {
	*gabs.Container
}

var _ = (types.Controller)(&Controller{})
