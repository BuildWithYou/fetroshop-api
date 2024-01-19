package errorhelper

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func GetStackTrace(err error) string { // TODO: fix this
	errWithStackTrace, ok := errors.Cause(err).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := errWithStackTrace.StackTrace()
	jsonString, err := json.Marshal(st)
	if err != nil {
		return "Failed to get stack trace with error: " + err.Error()
	}
	return string(jsonString)
}

func Error400(msg string) error {
	return fiber.NewError(fiber.StatusBadRequest, msg)
}

func Error401(msg string) error {
	return fiber.NewError(fiber.StatusUnauthorized, msg)
}

func Error500(msg string) error {
	return fiber.NewError(fiber.StatusInternalServerError, msg)
}

func ErrorCustom(errorCode int, msg string) error {
	return fiber.NewError(errorCode, msg)
}
