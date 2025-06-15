package errs

import "errors"

var (
	ErrProjectNotFound     = errors.New("project not found")
	ErrGoodsNotFound       = errors.New("goods not found")
	ErrInternalServerError = errors.New("internal server error")

	ErrCacheMiss = errors.New("cache miss")

	ErrWorkerIsDone = errors.New("worker is done")
)
