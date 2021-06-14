package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
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
		m := map[string]interface{}{}
		extractPathParams(c, m)
		extractQueryParams(c, m)
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
			fmt.Println(err)
			c.JSONP(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSONP(s.RespnseCode, resp)
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
	s.m = gabs.New()
	s.m.SetP(req, "request")
	for _, step := range s.Steps {
		switch step.Type {
		case Load:
			err := step.Load.Execute(ctx, s.m)
			if err != nil {
				return nil, err
			}
		case Store:
			err := step.Store.Execute(ctx, s.m)
			if err != nil {
				return nil, err
			}
		case CustomCode:
			err := step.CustomeCode.Execute(ctx, s.m)
			if err != nil {
				return nil, err
			}
		}
	}
	return s.m.Path(s.Response).Data(), nil
}
