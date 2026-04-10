// Package utils provides several helper function for golang.
package utils

import (
	"regexp"

	"github.com/google/uuid"
)

// NewId returns a new UUID in version 4.
func NewId() string {
	return uuid.New().String()
}

// IsId returns true if passed id is a UUID version 4.
func IsId(id string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(id)
}
