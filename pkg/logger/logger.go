package logger

import (
	"go.uber.org/zap"
	"log"
)

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(l)
}

func Init(dev bool) {
	var l *zap.Logger
	var err error
	if dev {
		l, err = zap.NewDevelopment()
	} else {
		l, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(l)
}
