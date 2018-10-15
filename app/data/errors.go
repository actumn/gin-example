package data

import (
	"errors"
	"github.com/appleboy/gin-jwt"
)

var (
	ErrDBConnection     = errors.New("database connection error")
	ErrResourceNotFound = errors.New("resource not found")
	ErrFailedAuth       = jwt.ErrFailedAuthentication
	ErrForbidden        = jwt.ErrForbidden
)
