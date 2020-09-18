package util

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// Error log the error to console and return it with status code.
func Error(code codes.Code, message string, err error) error {
	log.Printf("%s: [%v] : %v", message, code, err)

	formattedMessage := message + ":%v"

	return status.Errorf(code, formattedMessage, err)
}
