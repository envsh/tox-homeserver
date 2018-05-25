package store

import "strings"

func IsUniqueConstraintErr(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), "UNIQUE constraint failed:")
}

func IsIndexExistErr(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), "index ") &&
		strings.Contains(err.Error(), " already exists")
}
