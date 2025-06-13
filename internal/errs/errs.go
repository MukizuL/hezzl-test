package errs

import "errors"

var (
	ErrProjectNotFound     = errors.New("project not found")
	ErrGoodsNotFound       = errors.New("goods not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrWrongOrderFormat    = errors.New("invalid order number format")
	ErrConflictOrder       = errors.New("this order has already been uploaded by other user")
	ErrDuplicateOrder      = errors.New("this order has already been uploaded by this user")

	ErrWorkerIsDone = errors.New("worker is done")
)
