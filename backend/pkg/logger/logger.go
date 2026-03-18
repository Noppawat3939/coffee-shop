package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var Log *zap.Logger

func Init() {
	logger, err := zap.NewProduction() // prod config
	if err != nil {
		fmt.Println("Error intialize logger:", err)
	}

	Log = logger
}
