package storage

import "errors"

// определим общие ошибки
var (
	ErrUrlNotFound = errors.New("url not found")
	ErrUrlExists   = errors.New("url exists")
)
