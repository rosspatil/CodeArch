package main

import (
	"context"

	"github.com/rosspatil/codearch/runtime/services"
)

func main() {
	ctx := context.Background()
	services.Init(ctx, "./microservice.json")
}
