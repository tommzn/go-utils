// Package utils provides several helper function for golang.
package utils

import (
	uuid "github.com/fossoreslp/uuid"
)

// NewId returns a new UUID in version 4.
func NewId() string {
	return uuid.NewV4().String()
}

// NewV7Id returns a version 7 (Timestamp + Random, Recommended Sortable) UUID.
func NewV7Id() string {
	return uuid.NewV7().String()
}

// IsId returns true if passed id is a UUID version 4.
func IsId(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
