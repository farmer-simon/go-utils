package egu

import uuid "github.com/satori/go.uuid"

// UUID return unique id
func UUID() string {
	return uuid.NewV4().String()
}


// DefaultString return defaultV if v is empty
func DefaultString(v, defaultV string) string {
	if v == "" {
		return defaultV
	}

	return v
}

