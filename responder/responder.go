package responder

import (
	"github.com/sudtanj/Go-SwissKnife/conditional"
	"time"
)

//go:generate mockgen -source=./responder.go -destination=./mocks/responder.go -package=mocks

type IJSONResponder interface {
	JSON(code int, data interface{})
}

type Respond struct {
	status     string
	systemTime string
	data       interface{}
}

func Json[T any](c IJSONResponder, data T) {
	_, ok := any(data).(error)
	stringCond := conditional.NewConditional[string]()
	message := stringCond.If(!ok, "success", "error")
	c.JSON(
		200,
		Respond{
			status:     message,
			systemTime: time.Now().String(),
			data:       data,
		},
	)
	return
}
