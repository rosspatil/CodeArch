package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
	"github.com/rosspatil/codearch/runtime/models"
	"github.com/rosspatil/codearch/runtime/services/customerrors"
	"github.com/rosspatil/codearch/runtime/utils"
)

func Register(ctx context.Context, g *gin.Engine, s Service) {
	switch s.Method {
	case http.MethodGet:
		g.GET(s.Path, Handler(ctx, s, false))
	case http.MethodPut:
		g.PUT(s.Path, Handler(ctx, s, true))
	case http.MethodPost:
		g.POST(s.Path, Handler(ctx, s, true))
	case http.MethodDelete:
		g.DELETE(s.Path, Handler(ctx, s, false))
	}
}

func Handler(ctx context.Context, s Service, readBody bool) func(*gin.Context) {
	return func(c *gin.Context) {
		s.m = new(models.Controller)
		s.m.Container = gabs.New()
		m := map[string]interface{}{}
		extractPathParams(c, m)
		extractQueryParams(c, m)
		for k, v := range c.Request.Header {
			s.m.SetP(fmt.Sprintf("request.headers.%s", strings.ToLower(k)), v[0])
		}
		if readBody {
			ba, err := ioutil.ReadAll(c.Request.Body)
			if err == nil {
				err = json.Unmarshal(ba, &m)
			}
			if err != nil {
				c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		resp, err := s.Execute(ctx, m)
		if err != nil {
			err = s.OverrideErrors.Execute(ctx, s.m, err)
			fmt.Println(err)
			c.JSONP(customerrors.ErrorResponse(err))
			return
		}
		for k, v := range s.Response.Headers {
			c.Header(k, fmt.Sprint(utils.ResolveValue(v, s.m)))
		}
		c.JSONP(s.Response.HttpCode, resp)
	}
}

func extractPathParams(c *gin.Context, m map[string]interface{}) {
	for _, param := range c.Params {
		m[param.Key] = param.Value
	}
}
func extractQueryParams(c *gin.Context, m map[string]interface{}) {
	values := c.Request.URL.Query()
	for k, v := range values {
		m[k] = v[0]
	}
}

func (s Service) Execute(ctx context.Context, req interface{}) (interface{}, error) {
	err := validateRequest(ctx, s.RequestBody, req)
	if err != nil {
		return nil, err
	}

	s.m.SetP(req, "request")
	for _, step := range s.Steps {
		switch step.Type {
		case Load:
			err = step.Load.Execute(ctx, s.m)
		case Store:
			err = step.Store.Execute(ctx, s.m)
		case CustomCode:
			err = step.CustomeCode.Execute(ctx, s.m)
		case Condition:
			err = step.Condition.Execute(ctx, s.m)
		case Crypto:
			err = step.Crypto.Execute(ctx, s.m)
		}
		if err != nil {
			return nil, err
		}
	}
	return s.m.Path(s.Response.Field).Data(), nil
}
