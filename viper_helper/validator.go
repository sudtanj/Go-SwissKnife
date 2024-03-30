package viper_helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/sudtanj/Go-SwissKnife/env"
)

func ValidateRequiredConfig[T env.IConfig](config T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(config)
	if err != nil {
		panic(err.Error())
	}
}
