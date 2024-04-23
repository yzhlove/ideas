package main

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"os"
)

func main() {

	handle := slog.NewJSONHandler(os.Stderr, nil)
	log := slog.New(handle)
	log.Info("hello world", slog.Int("count", math.MaxInt))
	log.Warn("this is warning message", slog.String("detail", "test message!"))
	log.Error("print error message", slog.AnyValue(fmt.Errorf("configure [%s] failed! ", "User.csv")))

	if err := getResult(); err != nil {
		log.Error("get result failed.", slog.Any("Error", err))
	}

}

func getResult() error {
	return errors.New("return a error message. ")
}
