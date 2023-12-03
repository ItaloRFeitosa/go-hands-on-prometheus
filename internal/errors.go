package internal

import "github.com/italorfeitosa/go-hands-on-prometheus/pkg/errs"

var (
	ErrDaoLinkContext = errs.Builder().WithContext("dao/link")
	ErrLinkNotFound   = ErrDaoLinkContext.AsNotFound().
				WithCode("link_not_found").
				WithTemplate("link with id '%d' not found")

	ErrDatabaseInternals = ErrDaoLinkContext.AsInternal().
				WithCode("database_error")
)

var (
	ErrHandlerLinkContext = errs.Builder().WithContext("handler/link")
	ErrLinkInvalid        = ErrHandlerLinkContext.AsInvalidParams().
				WithCode("link_invalid_params")
)
