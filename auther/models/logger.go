package models

import (
	"log"

	"github.com/hashicorp/go-hclog"
)

type Logger interface {
	StandardLogger(*hclog.StandardLoggerOptions) *log.Logger
	Debug(message string, v ...interface{})
	Info(message string, v ...interface{})
	Error(message string, v ...interface{})
}
