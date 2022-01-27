package storage

import "fmt"

type StorageErr struct {
	internal error
}

func (e *StorageErr) Error() string {
	return fmt.Sprintf("song not found: %v", e.internal)
}

func (e *StorageErr) Unwrap() error {
	return e.internal
}
