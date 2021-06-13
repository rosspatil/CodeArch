package services

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rosspatil/codearch/runtime/connectors"
	"github.com/rosspatil/codearch/runtime/services/service"
)

type MicroService struct {
	Name       string                 `json:"name,omitempty"`
	Port       string                 `json:"port,omitempty"`
	Connectors []connectors.Connector `json:"connectors,omitempty"`
	Services   []service.Service      `json:"services,omitempty"`
}

func Init(ctx context.Context, path string) {
	ba, _ := ioutil.ReadFile(path)
	ms := &MicroService{}
	err := json.Unmarshal(ba, ms)
	if err != nil {
		panic(err)
	}
	for _, con := range ms.Connectors {
		err := con.Load(ctx)
		if err != nil {
			panic(err)
		}
	}
	g := gin.New()
	g.GET("/health/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	for _, s := range ms.Services {
		service.Register(ctx, g, s)
	}
	if ms.Port == "" {
		ms.Port = "8080"
	}
	panic(g.Run(":" + ms.Port))
}
