package log

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	mlogger "github.com/masuldev/mlogger/logger"
	"sync"
)

var logger *mlogger.Logger

type Interface interface {
	Debug(message string)
	Info(message string)
	Error(message string)
	Warning(message string)
	Critical(message string)
	Panic(message string)
}

func Init() error {
	newLogger, err := mlogger.NewLogger(nil, 2)
	if err != nil {
		return err
	}

	logger = newLogger
	return nil
}

func Debug(message string) {
	logger.Debug(message)
}

func Info(message string) {
	logger.Info(message)
}

func Error(message string) {
	logger.Error(message)
}

func Warning(message string) {
	logger.Warning(message)
}

func Critical(message string) {
	logger.Critical(message)
}

func Panic(message string) {
	logger.Panic(message)
}

type Config struct {
	Next   func(c *fiber.Ctx) bool
	Logger *mlogger.Logger
}

func NewLogger(config ...Config) fiber.Handler {
	var (
		errPadding = 15
		once       sync.Once
		errHandler fiber.ErrorHandler
		cfg        Config
	)

	if len(config) < 1 {
		cfg = Config{
			Next:   nil,
			Logger: logger,
		}
	} else {
		cfg = config[0]
	}

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		once.Do(func() {
			errHandler = c.App().Config().ErrorHandler
			stack := c.App().Stack()
			for m := range stack {
				for r := range stack[m] {
					if len(stack[m][r].Path) > errPadding {
						errPadding = len(stack[m][r].Path)
					}
				}
			}
			errHandler = c.App().ErrorHandler
		})

		chainErr := c.Next()

		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		if chainErr != nil {
			Error(chainErr.Error())
			return nil
		}

		Info(fmt.Sprintf("[%s] api.request", c.Request().RequestURI()))

		return nil
	}
}
