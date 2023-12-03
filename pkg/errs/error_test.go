package errs_test

import (
	"encoding/json"
	"errors"
	"log"
	"testing"

	"github.com/italorfeitosa/go-hands-on-prometheus/pkg/errs"
)

var ErrInvalidField = errors.New("invalid field")
var ErrMinLength = errors.New("min length")
var ErrDatabase = errors.New("database error")

var ErrController = errs.Builder().WithContext("controller")

var ErrInvalidUserCreateRequest = ErrController.
	AsInvalidOperation().
	WithCode("invalid_user").
	WithTemplate("field %s is invalid")

var ErrRepository = errs.Builder().WithContext("repository")

var ErrDatabaseInternals = ErrRepository.
	AsInternal().
	WithCode("database_internals").
	WithTemplate("some error on database happened: %w")

func TestError(t *testing.T) {
	err := createUser()
	bytes, _ := json.MarshalIndent(err, "", "  ")
	log.Print(string(bytes))

	log.Print(errors.Is(err, ErrDatabaseInternals))
}

func createUser() error {
	err := databaseCall()
	return ErrDatabaseInternals.WithArgs(err)
}

func databaseCall() error {
	return ErrDatabase
}
