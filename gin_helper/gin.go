package gin_helper

import (
	"github.com/gin-gonic/gin"
	"github.com/sudtanj/Go-SwissKnife/env"
)

//go:generate ifacemaker -f ./GinHelper.go -s GinHelper[T] -i "IGinHelper[T any]" -p interfaces -o ./interfaces/gin_helper.go
//go:generate mockgen -source=./interfaces/gin_helper.go -destination=./mocks/gin_helper.go -package=mocks

type GinHelper[T env.IConfig] struct {
	config T
	r      *gin.Engine
}

func NewGinHelper[T env.IConfig](config T) *GinHelper[T] {
	return &GinHelper[T]{
		config,
		gin.Default(),
	}
}

func (g *GinHelper[T]) GetEngine() *gin.Engine {
	return g.r
}

func (g *GinHelper[T]) Initialize() *gin.Engine {
	err := g.r.Run(g.GetAddr())
	if err != nil {
		panic("cannot run gin server!")
	}
	return g.r
}

func (g *GinHelper[T]) GetAddr() string {
	return g.config.GetAddr() + ":" + g.config.GetPort()
}
