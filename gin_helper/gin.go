package gin_helper

import (
	helmet "github.com/danielkov/gin-helmet"
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sudtanj/Go-SwissKnife/env"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	"github.com/tavsec/gin-healthcheck/config"
)

//go:generate ifacemaker -f ./GinHelper.go -s GinHelper[T] -i "IGinHelper[T any]" -p interfaces -o ./interfaces/gin_helper.go
//go:generate mockgen -source=./interfaces/gin_helper.go -destination=./mocks/gin_helper.go -package=mocks

type GinHelper[T env.IConfig] struct {
	config T
	r      *gin.Engine
}

type IGinHelperRouter interface {
	InjectRoute()
}

type IGinHelperRouteInitializer[T IGinHelperRouter] interface {
	InitializeRouter(g *gin.Engine) T
}

func InitializeGin[T env.IConfig, V IGinHelperRouteInitializer[IGinHelperRouter]](config T, router V) {
	ginHelper := NewGinHelper(config)
	handler := router.InitializeRouter(ginHelper.GetEngine())
	handler.InjectRoute()
	ginHelper.Initialize()
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
	pingCheck := checks.NewPingCheck("https://www.google.com", "GET", 1000, nil, nil)
	err := healthcheck.New(g.r, config.DefaultConfig(), []checks.Check{pingCheck})
	if err != nil {
		panic(err)
	}
	g.r.Use(helmet.Default())
	g.r.Use(cors.Default())
	g.r.Use(nice.Recovery(recoveryHandler))

	err = g.r.Run(g.GetAddr())
	if err != nil {
		panic(err)
	}
	return g.r
}

func (g *GinHelper[T]) GetAddr() string {
	return g.config.GetAddr() + ":" + g.config.GetPort()
}
